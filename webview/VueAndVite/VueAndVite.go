// Vue + Vite 桌面应用
// 开发模式: 连接 Vite 开发服务器 (http://localhost:5173)，支持极速热重载
// 生产模式: 嵌入 dist 目录资源，单文件部署
package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	isDebug = true // 是否为调试模式
	host    string

	//go:embed dist/**
	embedAssets embed.FS // 嵌入 dist 目录以及子目录下的文件 (生产模式使用)
)

const hostName = "app.example"

type MainWindow struct {
	edg *edge.Edge
	w   *window.Window
	wv  *edge.WebView
}

func NewMainWindow(edg *edge.Edge) *MainWindow {
	m := &MainWindow{edg: edg}
	var err error

	// 创建 WebView
	m.w, m.wv, err = m.edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("Vue Desktop App"),
		edge.WithXmlWindowClassName("VueViteApp"),
		edge.WithXmlWindowSize(1300, 900),
		edge.WithFillParent(true),
		edge.WithAppDrag(true),
		edge.WithDebug(isDebug),
		edge.WithDefaultContextMenus(isDebug),
		edge.WithBrowserAcceleratorKeys(isDebug),
		edge.WithStatusBar(false),
		edge.WithZoomControl(false),
		edge.WithAutoFocus(true),
		// 设置默认背景色为透明, 这也是当采用嵌入文件系统的方式时防止首次加载会闪烁的一部分
		edge.WithDefaultBackgroundColor(edge.NewColor(0, 0, 0, 0)),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	if !isDebug {
		// 生产模式: 使用嵌入文件系统
		err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
		if err != nil {
			wapi.MessageBoxW(0, "SetVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(6)
		}
		err = m.wv.EnableVirtualHostNameToEmbedFSMapping(true)
		if err != nil {
			wapi.MessageBoxW(0, "EnableVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(5)
		}

		host = edge.JoinUrlHeader(hostName)
		fmt.Println("生产模式: 使用嵌入文件系统")
	} else {
		// 开发模式: 连接 Vite 开发服务器
		host = "http://localhost:5173"
		fmt.Println("开发模式: 连接 Vite 开发服务器", host)
	}

	// 注册 WebView 事件
	m.regWebViewEvents()
	// 绑定函数
	m.bindBasicFuncs()
	// 导航到首页
	m.wv.Navigate(host + "/index.html")
	return m
}

// 注册 WebView 事件
func (m *MainWindow) regWebViewEvents() {
	var firstLoad = true
	// 导航完成事件
	m.wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		uri, err := sender.GetSource()
		if err != nil {
			log.Println("GetSource 失败:", err)
			return 0
		}
		fmt.Println("导航完成:", uri)

		switch uri {
		case host + "/index.html":
			// 在导航完成事件里判断第一次加载完毕时才显示窗口,
			// 这是因为采用嵌入文件系统的方式时, 网页还没加载出来的时候, 会显示webview白色的背景,
			// 然后才会加载出网页, 表现出来就是有一瞬间的闪烁, 所以等加载完再显示窗口
			if firstLoad {
				firstLoad = false
				m.w.Show(true)
			}
		}
		return 0
	})
}

// bindBasicFuncs 绑定基本函数
func (m *MainWindow) bindBasicFuncs() {
	// 绑定最小化窗口函数
	m.wv.Bind("wnd.minimize", func() {
		m.w.ShowWindow(xcc.SW_MINIMIZE)
	})

	// 绑定切换最大化窗口函数
	m.wv.Bind("wnd.toggleMaximize", func() {
		m.w.MaxWindow(!m.w.IsMaxWindow())
	})

	// 绑定关闭窗口函数
	m.wv.Bind("wnd.close", func() {
		m.w.CloseWindow()
	})
}

/* func init() {
	// 当使用嵌入文件系统时，返回资源会使用go内置的 mime.TypeByExtension 查找 MIME 类型,
	// 并设置 Content-Type 头为对应的 MIME 类型.
	// Go 的默认 MIME 表可能不包含某些新扩展名（如 .mjs）。如有需要你可以手动注册：
	mime.AddExtensionType(".mjs", "application/javascript")
	mime.AddExtensionType(".ts", "application/typescript")
	mime.AddExtensionType(".wasm", "application/wasm")
} */

func main() {
	checkWebView2()
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	NewMainWindow(edg)

	a.Run()
	a.Exit()
}

// 创建 WebView2 环境
func createEdge() *edge.Edge {
	// 创建 WebView2 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(),
		EnvOptions: &edge.EnvOptions{
			DisableTrackingPrevention: true,
			ScrollBarStyle:            edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY,
		},
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}
	return edg
}

func checkWebView2() {
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())
	localVersion, err := edge.GetAvailableBrowserVersion()
	if err != nil {
		wapi.MessageBoxW(0, "获取 WebView2 运行时版本号失败: "+err.Error(), "提示", wapi.MB_IconError)
		os.Exit(1)
	}
	if localVersion == "" {
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序! 下载完后请使用管理员权限运行安装包!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)

	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}
