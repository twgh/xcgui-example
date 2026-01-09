// 网络资源请求和响应接收事件
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	edg   *edge.Edge
	wv    *edge.WebView
	w     *window.Window
	impls = make([]*edge.WebViewEventImpl, 2)
)

func main() {
	checkWebView2()
	edg = createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w = window.New(0, 0, 1400, 900, "网络资源请求和响应接收事件", 0, xcc.Window_Style_Default)

	var err error
	// 创建 WebView
	wv, err = edg.NewWebView(w.Handle,
		edge.WithFillParent(true), // WebView 填充窗口
		edge.WithDebug(true),      // 可打开开发者工具
		edge.WithAutoFocus(true),  // 在窗口获得焦点时尝试保持 WebView 的焦点
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 创建多个 WebView 事件接口实现对象.
	for i := range impls {
		impls[i] = edge.NewWebViewEventImpl(wv)
	}

	// 添加网络资源请求过滤, 这里获取出错是因为运行时版本低
	wv.WebView2_22, err = wv.CoreWebView.GetICoreWebView2_22()
	if err == nil {
		/*
			获取k线数据:
				push2his.eastmoney.com/api/qt/stock/kline/get
			获取个股信息:
				push2.eastmoney.com/api/qt/stock/get
		*/
		wv.WebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds("*push2*.eastmoney.com/api/qt/stock*/get*", edge.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SCRIPT, edge.COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_ALL)
	} else {
		wv.AddWebResourceRequestedFilter("*push2*.eastmoney.com/api/qt/stock*/get*", edge.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SCRIPT)
	}

	// 注册 WebView 事件
	regWebViewEvents()

	// 导航
	wv.Navigate("https://quote.eastmoney.com/concept/sh600000.html")

	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}

func createEdge() *edge.Edge {
	// 创建 WebView 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 实际应用中应使用自己创建的固定目录，这里用临时目录示例
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}
	return edg
}

// 注册 WebView 事件
func regWebViewEvents() {
	// 网络资源请求事件
	wv.Event_WebResourceRequested(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2WebResourceRequestedEventArgs) uintptr {
		// 获取请求
		req, err := args.GetRequest()
		if err != nil {
			log.Println("请求获取失败:", err.Error())
			return 0
		}
		defer req.Release()

		// 获取请求地址
		uri, err := req.GetUri()
		if err != nil {
			log.Println("请求地址获取失败:", err.Error())
			return 0
		}

		// 获取请求方法
		method, err := req.GetMethod()
		if err != nil {
			log.Println("请求方法获取失败:", err.Error())
			return 0
		}

		if method == http.MethodGet {
			switch {
			// 日k线数据查询
			case isKlineDayGet(uri):
				// 改变日k线数据条数
				newUrl := strings.ReplaceAll(uri, "lmt=210", "lmt=570")
				err = req.SetUri(newUrl)
				if err != nil {
					log.Println("日k线数据请求地址设置失败:", err.Error())
					return 0
				}

			// 个股信息查询
			case isStockInfoGet(uri):
				// 给 uri 中的 fields 参数中加上 f189%2C
				newUrl := strings.Replace(uri, "fields=", "fields=f189%2C", 1)
				err = req.SetUri(newUrl)
				if err != nil {
					log.Println("个股信息查询请求地址设置失败", err.Error())
					return 0
				}
			}
		}
		return 0
	})

	// 网络资源响应接收事件
	wv.Event_WebResourceResponseReceived(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr {
		// 获取请求
		req, err := args.GetRequest()
		if err != nil {
			log.Println("请求获取失败:", err.Error())
			return 0
		}
		defer req.Release()

		// 获取请求地址
		uri, err := req.GetUri()
		if err != nil {
			log.Println("请求地址获取失败:", err.Error())
			return 0
		}

		// 获取请求方法
		method, err := req.GetMethod()
		if err != nil {
			log.Println("请求方法获取失败:", err.Error())
			return 0
		}

		if method == http.MethodGet {
			switch {
			// 日k线数据查询
			case isKlineDayGet(uri):
				// 获取响应内容字符串
				getResponseStr(args, "日k线数据响应", func(resStr string) {
					go saveFile(resStr, "日k线数据响应")
				})

			// 个股信息查询
			case isStockInfoGet(uri):
				// 获取响应内容字符串
				getResponseStr(args, "个股信息查询响应", func(resStr string) {
					go saveFile(resStr, "个股信息查询响应")
				})

			// 行业查询
			case isIndustryGet(uri):
				// 获取响应内容字符串
				getResponseStr(args, "行业查询响应", func(resStr string) {
					go saveFile(resStr, "行业查询响应")
				})
			}
		}
		return 0
	})

	// 导航完成事件
	wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		// 获取导航是否成功
		isSuccess, err := args.GetIsSuccess()
		if err != nil {
			log.Println("导航是否成功获取失败", err.Error())
			return 0
		}

		if !isSuccess {
			log.Println("导航失败, 错误状态码:", args.MustGetWebErrorStatus(), ", url:", sender.MustGetSource())
			return 0
		}

		// 设置窗口标题
		w.SetTitle(sender.MustGetDocumentTitle()).Redraw(false)

		// 点击日K选项
		go func() {
			time.Sleep(time.Millisecond * 500)
			for i := 0; i < 5; i++ {
				time.Sleep(time.Millisecond * 300)
				xc.UI(func() {
					clickDayK()
				})
			}
		}()
		return 0
	})
}

const prefixPath = `webview\WebResourceRequestedEvent\`

func saveFile(resStr string, resType string) {
	if resStr == "" {
		log.Println("响应内容为空,", resType)
		return
	}

	err := os.WriteFile(prefixPath+resType+`.json`, []byte(resStr), 0666)
	if err != nil {
		log.Println("保存文件失败:", err.Error())
		return
	}
	
	log.Println("保存文件成功:", prefixPath+resType+`.json`)
}

// 点击日k
func clickDayK() {
	clickEle("#app > div > div.maincharts.self_clearfix > div.charts > div > div.timechart_types.self_clearfix > ul > li:nth-child(4)")
}

// 点击元素
func clickEle(selector string) {
	wv.Eval(`document.querySelector('` + selector + `')?.click();`)
}

// 获取响应内容字符串, 在回调函数中处理
func getResponseStr(args *edge.ICoreWebView2WebResourceResponseReceivedEventArgs, resType string, cb func(resStr string)) {
	// 获取响应
	res, err := args.GetResponse()
	if err != nil {
		log.Println("响应获取失败", err.Error())

		// 确保任务能完成，即使获取响应失败也要触发相应的保存函数
		if cb != nil {
			// 传递空字符串给回调函数，让保存函数处理空数据的情况
			go cb("")
		}
		return
	}
	defer res.Release()

	// 获取响应内容
	// 因为这个方法是异步的，如果连续调用, 前面的还没执行完，
	// 后面再次调用就会把前面的回调函数给覆盖掉, 导致数据混乱,
	// 所以要在第一个参数里给每个请求分配不同的*edge.WebViewEventImpl,
	// 它们就会有各自独立的回调函数, 不会相互覆盖.
	err = res.GetContentEx(getImpl(resType), func(errorCode syscall.Errno, content []byte) uintptr {
		if !wapi.IsOK(errorCode) {
			log.Println("响应内容获取失败:", err.Error(), "res_type:", resType)

			// 确保任务能完成，即使获取响应内容失败也要触发相应的保存函数
			if cb != nil {
				// 传递空字符串给回调函数，让保存函数处理空数据的情况
				go cb("")
			}
			return 0
		}

		// 响应内容字符串
		contentStr := string(content)
		// 删除从开头到第一个(之间的文本
		startIndex := strings.Index(contentStr, "(")
		if startIndex != -1 {
			contentStr = contentStr[startIndex+1:]
		}
		// 删除末尾的 );
		contentStr = strings.TrimSuffix(contentStr, ");")

		// 将内容传进回调函数
		if cb != nil {
			cb(contentStr)
		}
		return 0
	})

	if err != nil {
		log.Println("响应内容获取失败:", err.Error(), "res_type:", resType)

		// 确保任务能完成，即使获取响应内容失败也要触发相应的保存函数
		if cb != nil {
			// 传递空字符串给回调函数，让保存函数处理空数据的情况
			go cb("")
		}
	}
}

// 根据 resType 获取获取 WebViewEventImpl
func getImpl(resType string) *edge.WebViewEventImpl {
	switch resType {
	case "日k线数据响应":
		return wv.GetWebViewEventImpl()
	case "个股信息查询响应":
		return impls[0]
	case "行业查询响应":
		return impls[1]
	}
	return nil
}

// 判断是否是行业查询
func isIndustryGet(uri string) bool {
	return strings.Contains(uri, "datacenter-web.eastmoney.com/web/api/data/v1/get") && strings.Contains(uri, "EM2016%2CMAIN_BUSINESS")
}

// 判断是否是日k线数据查询
func isKlineDayGet(uri string) bool {
	return strings.Contains(uri, "push2his.eastmoney.com/api/qt/stock/kline/get") && strings.Contains(uri, "cb=quote_jp") && strings.Contains(uri, "klt=101") && strings.Contains(uri, "fqt=1") && strings.Contains(uri, "lmt=210")
}

// 判断是否是个股信息查询
func isStockInfoGet(uri string) bool {
	return strings.Contains(uri, "push2.eastmoney.com/api/qt/stock/get") && strings.Contains(uri, "fields=f58%2Cf734%2Cf107%2Cf57%2Cf43%2Cf59%2Cf169%2Cf170%2Cf152%2Cf177%2Cf111%2Cf46%2Cf60%2Cf44%2Cf45%2Cf47%2Cf260%2Cf48%2Cf261%2Cf279%2Cf277%2Cf278%2Cf288%2Cf19%2Cf17%2Cf531%2Cf15%2Cf13%2Cf11%2Cf20%2Cf18%2Cf16%2Cf14%2Cf12%2Cf39%2Cf37%2Cf35%2Cf33%2Cf31%2Cf40%2Cf38%2Cf36%2Cf34%2Cf32%2Cf211%2Cf212%2Cf213%2Cf214%2Cf215%2Cf210%2Cf209%2Cf208%2Cf207%2Cf206%2Cf161%2Cf49%2Cf171%2Cf50%2Cf86%2Cf84%2Cf85%2Cf168%2Cf108%2Cf116%2Cf167%2Cf164%2Cf162%2Cf163%2Cf92%2Cf71%2Cf117%2Cf292%2Cf51%2Cf52%2Cf191%2Cf192%2Cf262%2Cf294%2Cf181%2Cf295%2Cf269%2Cf270%2Cf256%2Cf257%2Cf285%2Cf286%2Cf120%2Cf121%2Cf122%2Cf55%2Cf174%2Cf175%2Cf135%2Cf136%2Cf301%2Cf803")
}

func checkWebView2() {
	// 输出本库使用的 WebView2 版本
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())

	// 获取本机已安装的 WebView2 运行时版本
	localVersion, err := edge.GetAvailableBrowserVersion()
	if err != nil {
		wapi.MessageBoxW(0, "获取 WebView2 运行时版本号失败: "+err.Error(), "提示", wapi.MB_IconError)
		os.Exit(1)
	}
	if localVersion == "" {
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序! 下载完后请使用管理员权限运行安装包!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		// 打开 WebView2 运行时下载页面
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)

	// 检查本机版本是否低于库版本
	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}
