// 注册自定义方案
package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	w  *window.Window
	wv *edge.WebView

	//go:embed CustomSchemeRegistration.html
	indexHtml string
)

func main() {
	checkWebView2()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建 WebView2 环境
	edg := createEdge()

	var err error
	// 创建 WebView
	w, wv, err = edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("注册自定义方案"),
		edge.WithXmlWindowSize(600, 400),
		edge.WithXmlWindowTitleBar(true),
		edge.WithXmlWindowTitleBarBgColor(xc.RGBA(58, 118, 206, 255)),
		edge.WithFillParent(true), // WebView 填充窗口
		edge.WithDebug(true),      // 启用调试模式
		edge.WithAutoFocus(true),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 注册 WebView 事件
	regWebViewEvent()

	// 导航
	wv.NavigateToString(indexHtml)

	// 显示窗口并运行应用
	w.Show(true)
	a.Run()
	a.Exit()
}

// 注册 WebView 事件
func regWebViewEvent() {
	// 导航开始事件
	wv.Event_NavigationStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationStartingEventArgs) uintptr {
		uri, err := args.GetUri()
		if err != nil {
			log.Println("NavigationStarting GetUri 失败:", err)
			return 0
		}
		log.Println("uri:", uri)

		// 1. 检查是否是自定义协议
		if strings.HasPrefix(uri, "myapp://") {
			// 取消导航
			err := args.SetCancel(true)
			if err != nil {
				log.Println("NavigationStarting SetCancel 失败:", err)
				w.MessageBox("错误", "NavigationStarting SetCancel 失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
				return 0
			}

			// 2. 解析 URL
			u, err := url.Parse(uri)
			if err != nil {
				log.Printf("URL 解析失败: %v\n", err)
				return 0
			}

			// 3. 获取路径和查询参数
			path := "/" + u.Host // 例如: "/openFile"
			query := u.Query()   // 类似 map[string][]string

			log.Printf("路径: %s\n", path)
			log.Printf("查询参数: %v\n", query)

			// 4. 根据路径执行不同操作
			switch path {
			case "/openFile":
				// 文件路径参数, 不为空就是打开文件
				// myapp://openFile?path=C:\Windows\system.ini
				filePath := query.Get("path")
				if filePath != "" {
					wapi.ShellExecuteW(0, "open", filePath, "", "", xcc.SW_SHOW)
					return 0
				}

				// 如果 path 参数为空，则弹出文件选择对话框
				filePath = wutil.OpenFile(w.Handle, []string{"All Files(*.*)", "*.*"}, "")
				if filePath == "" {
					return 0
				}

				log.Println("选择的文件路径:", filePath)
				filePath = strings.ReplaceAll(filePath, `\`, `\\`) // 转义反斜杠
				wv.Eval(`document.getElementById('filePath').value = '` + filePath + `'`)

			case "/showSettings":
				log.Println("显示设置窗口")
				app.Alert("提示", "设置窗口")
				// 例如: wv.Navigate("http://app.example/showSettings")

			case "/navigateTo":
				page := query.Get("page")
				if page == "" {
					page = "home"
				}

				switch page {
				case "order":
					id := query.Get("id")
					log.Printf("跳转到页面: %s, id: %s\n", page, id)
					app.Alert("提示", fmt.Sprintf("跳转到页面: %s, id: %s\n", page, id))
					// 例如：wv.Navigate(fmt.Sprintf("http://app.example/%s/%s", page, id))

				default:
					log.Printf("跳转到页面: %s\n", page)
					app.Alert("提示", fmt.Sprintf("跳转到页面: %s\n", page))
					// 例如：wv.Navigate(fmt.Sprintf("http://app.example/%s", page))
				}

			default:
				log.Printf("未知命令: %s\n", path)
			}
		}
		return 0
	})
}

// 创建 WebView2 环境
func createEdge() *edge.Edge {
	customScheme, err := edge.CreateCustomSchemeRegistration("myapp")
	if err != nil {
		wapi.MessageBoxW(0, "创建自定义方案 myapp 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}
	defer customScheme.Release()

	// 创建 WebView 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 实际应用中应使用自己创建的固定目录，这里用临时目录示例
		EnvOptions: &edge.EnvOptions{
			ExclusiveUserDataFolderAccess: true, // 其他进程可以从使用相同用户数据文件夹创建的 WebView2Environment 创建 WebView2，从而共享同一个 WebView 浏览器进程实例
			DisableTrackingPrevention:     true, // 禁用 WebView2 中的跟踪防护功能
			// 滚动条样式
			ScrollBarStyle:            edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY,
			CustomSchemeRegistrations: []*edge.ICoreWebView2CustomSchemeRegistration{customScheme},
		},
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	return edg
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
