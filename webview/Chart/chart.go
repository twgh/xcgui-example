// 在窗口中创建多个 webview 元素, 显示图表, Go发送数据到前端以更新图标
package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
)

var (
	//go:embed res/Chart.xml
	xmlStr string
	//go:embed assets/**
	embedAssets embed.FS // 嵌入 assets 目录以及子目录下的文件, 不包括隐藏文件
	isDebug     = true   // 是否为调试版
)

const hostName = "app.example"

func main() {
	checkWebView2()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口从布局文件
	w := window.NewByLayoutStringW(xmlStr, 0, 0)

	// 创建 WebView2 环境选项.
	envOpts, err := edge.CreateEnvironmentOptions()
	if err != nil {
		log.Println("创建 WebView2 环境选项失败: " + err.Error())
	} else {
		defer envOpts.Release()
		// 获取 WebView2 环境选项5
		envOpts5, err := envOpts.GetICoreWebView2EnvironmentOptions5()
		if err != nil {
			log.Println("获取环境选项5失败: " + err.Error())
		} else {
			// 禁用 WebView2 中的跟踪防护功能以提高运行时性能, 仅在 WebView2 中呈现已知安全的内容时可以这样做.
			// 如果 WebView2 被用作具有任意导航功能的“完整浏览器”且需要保护最终用户隐私，那么不应禁用此属性。
			envOpts5.SetEnableTrackingPrevention(false)
			envOpts5.Release()
		}

		// 获取 WebView2 环境选项8
		envOpts8, err := envOpts.GetICoreWebView2EnvironmentOptions8()
		if err != nil {
			log.Println("获取环境选项8失败: " + err.Error())
		} else {
			// 设置滚动条样式
			envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY)
			envOpts8.Release()
		}
	}

	// 创建 webview 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder:     os.TempDir(), // 自己的软件应该在固定位置创建一个自己的目录, 而不是用临时目录
		EnvironmentOptions: envOpts,
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 设置虚拟主机名和嵌入文件系统之间的映射.
	// 这个映射是全局可用的, 所有的 WebView 都可访问, 只需 EnableVirtualHostNameToEmbedFSMapping(true)
	// 如果想要删除映射关系, 请使用 edge.DeleteVirtualHostNameToEmbedFSMapping(hostName)
	edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)

	// 创建 webview 选项
	wvOption := []edge.WebViewOption{
		edge.WithFillParent(true),
		edge.WithStatusBar(false),
		edge.WithDefaultContextMenus(isDebug),
		edge.WithBrowserAcceleratorKeys(isDebug),
		edge.WithDebug(isDebug),
	}

	// 创建 webview1
	createWebView1(edg, wvOption)
	// 创建 webview2
	createWebView2(edg, wvOption)
	// 创建 webview3
	createWebView3(edg, wvOption)

	w.AdjustLayout()
	w.Show(true)
	a.Run()
	a.Exit()
}

func createWebView1(edg *edge.Edge, wvOption []edge.WebViewOption) {
	layout_chart := widget.NewLayoutEleByName("layout_chart1")
	wv, err := edg.NewWebView(layout_chart.Handle, wvOption...)
	if err != nil {
		log.Println("创建 webview1 失败: " + err.Error())
		return
	}
	wv.Show(false)
	wv.EnableVirtualHostNameToEmbedFSMapping(true)

	firstLoad := true
	// 导航完成事件
	wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		uri := sender.MustGetSource()
		switch uri {
		case edge.JoinUrlHeader(hostName) + "/chart1.html":
			if firstLoad { // 首次加载页面
				firstLoad = false
				wv.Show(true) // 首次加载完成才显示
				return 0
			}
		}
		return 0
	})

	wv.Navigate(edge.JoinUrlHeader(hostName) + "/chart1.html")
}

func createWebView2(edg *edge.Edge, wvOption []edge.WebViewOption) {
	layout_chart := widget.NewLayoutEleByName("layout_chart2")
	wv, err := edg.NewWebView(layout_chart.Handle, wvOption...)
	if err != nil {
		log.Println("创建 webview2 失败: " + err.Error())
		return
	}
	wv.Show(false)
	wv.EnableVirtualHostNameToEmbedFSMapping(true)

	firstLoad := true
	// 导航完成事件
	wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		uri := sender.MustGetSource()
		switch uri {
		case edge.JoinUrlHeader(hostName) + "/chart2.html":
			if firstLoad { // 首次加载页面
				firstLoad = false
				wv.Show(true) // 首次加载完成才显示
				return 0
			}
		}
		return 0
	})

	wv.Navigate(edge.JoinUrlHeader(hostName) + "/chart2.html")
}

// DataPoint 定义数据点结构
type DataPoint struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

func createWebView3(edg *edge.Edge, wvOption []edge.WebViewOption) {
	layout_chart := widget.NewLayoutEleByName("layout_chart3")
	wv, err := edg.NewWebView(layout_chart.Handle, wvOption...)
	if err != nil {
		log.Println("创建 webview3 失败: " + err.Error())
		return
	}
	wv.Show(false)
	wv.EnableVirtualHostNameToEmbedFSMapping(true)

	firstLoad := true
	// 导航完成事件
	wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		uri := sender.MustGetSource()
		switch uri {
		case edge.JoinUrlHeader(hostName) + "/chart3.html":
			if firstLoad { // 首次加载页面时发送数据
				firstLoad = false
				// 模拟一些数据
				data := []DataPoint{
					{"A", 30},
					{"B", 45},
					{"C", 60},
					{"D", 25},
					{"E", 80},
				}
				bs, _ := json.Marshal(data)
				wv.PostWebMessageAsJSON(string(bs))
				wv.Show(true) // 首次加载完成才显示
				return 0
			}
		}
		return 0
	})

	wv.Navigate(edge.JoinUrlHeader(hostName) + "/chart3.html")

	rand.Seed(time.Now().UnixNano())
	btn := widget.NewButtonByName("按钮_发送数据")
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		// 生成随机数据
		data := []DataPoint{
			{"A", rand.Float64()*100 + 10},
			{"B", rand.Float64()*100 + 10},
			{"C", rand.Float64()*100 + 10},
			{"D", rand.Float64()*100 + 10},
			{"E", rand.Float64()*100 + 10},
		}
		bs, _ := json.Marshal(data)
		wv.PostWebMessageAsJSON(string(bs))
		return 0
	})
}

func checkWebView2() {
	// 高版本运行时是兼容低版本的。低版本不用高版本的方法也没有问题。
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())
	// 获取本机已安装的 WebView2 运行时版本号
	localVersion, err := edge.GetAvailableBrowserVersion()
	if err != nil {
		wapi.MessageBoxW(0, "获取 WebView2 运行时版本号失败: "+err.Error(), "提示", wapi.MB_IconError)
		os.Exit(1)
	}
	if localVersion == "" {
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序! 下载完后请使用管理员权限运行安装包!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		// 使用默认浏览器打开 WebView2 运行时小型引导程序下载地址.
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)

	// 判断本机 WebView2 运行时版本是否低于本库使用的 WebView2 运行时版本
	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}
