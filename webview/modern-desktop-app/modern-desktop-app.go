// 现代风格桌面应用
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
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

//go:embed assets/**
var embedAssets embed.FS
var isDebug = true

const hostName = "app.modern"

type MainWindow struct {
	edg *edge.Edge
	w   *window.Window
	wv  *edge.WebView
}

func NewMainWindow(edg *edge.Edge) *MainWindow {
	m := &MainWindow{edg: edg}
	var err error

	// 创建 WebView 窗口
	m.w, m.wv, err = m.edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("现代风格桌面应用"),
		edge.WithXmlWindowSize(1200, 750),
		edge.WithFillParent(true),
		edge.WithDebug(isDebug),                                    // 正式版关闭调试
		edge.WithDefaultContextMenus(isDebug),                      // 禁用右键菜单
		edge.WithBrowserAcceleratorKeys(isDebug),                   // 禁用浏览器快捷键
		edge.WithStatusBar(false),                                  // 禁用状态栏
		edge.WithZoomControl(false),                                // 禁用缩放控制
		edge.WithAutoFocus(true),                                   // 自动聚焦
		edge.WithAppDrag(true),                                     // 允许应用内拖动
		edge.WithDefaultBackgroundColor(edge.NewColor(0, 0, 0, 0)), // 设置背景透明
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 设置窗口最小大小
	m.w.SetMinimumSize(800, 500)

	// 嵌入资源映射
	err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
	if err != nil {
		wapi.MessageBoxW(0, "SetVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}
	err = m.wv.EnableVirtualHostNameToEmbedFSMapping(true)
	if err != nil {
		wapi.MessageBoxW(0, "EnableVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(3)
	}

	// 绑定函数
	m.bindFunctions()

	// 注册事件
	m.regWebViewEvents()

	// 导航到首页
	m.wv.Navigate(edge.JoinUrlHeader(hostName) + "/index.html")
	return m
}

// 注册 WebView 事件
func (m *MainWindow) regWebViewEvents() {
	firstLoad := true
	// 导航完成事件
	m.wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		uri := sender.MustGetSource()
		fmt.Println("导航完成:", uri)

		if firstLoad && uri == edge.JoinUrlHeader(hostName)+"/index.html" {
			firstLoad = false
			m.w.Show(true) // 导航完毕再显示窗口
		}
		return 0
	})
}

// 绑定 Go 函数到 JS
func (m *MainWindow) bindFunctions() {
	// 窗口最小化
	m.wv.Bind("app.minimize", func() {
		m.w.ShowWindow(xcc.SW_SHOWMINIMIZED)
	})

	// 窗口最大化/还原
	m.wv.Bind("app.toggleMaximize", func() {
		m.w.MaxWindow(!m.w.IsMaxWindow())
	})

	// 窗口关闭
	m.wv.Bind("app.close", func() {
		m.w.CloseWindow()
	})

	// 打开设置
	m.wv.Bind("app.openSettings", func() {
		fmt.Println("打开设置")
		m.wv.Eval(`showSettings()`)
	})

	// 显示关于
	m.wv.Bind("app.showAbout", func() {
		wapi.MessageBoxW(m.w.GetHWND(), "现代桌面应用 v1.0.0\n基于 xcgui + WebView2 构建", "关于", wapi.MB_OK|wapi.MB_IconInformation)
	})
}

func main() {
	checkWebView2()
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	NewMainWindow(edg)

	a.Run()
	a.Exit()
}

// 创建 WebView 环境
func createEdge() *edge.Edge {
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(),
		EnvOptions: &edge.EnvOptions{
			DisableTrackingPrevention: true,
			ScrollBarStyle:            edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY,
		},
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}
	return edg
}

// 检查 WebView2 运行时
func checkWebView2() {
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())

	localVersion, err := edge.GetAvailableBrowserVersion()
	if err != nil {
		wapi.MessageBoxW(0, "获取 WebView2 运行时版本号失败: "+err.Error(), "提示", wapi.MB_IconError)
		os.Exit(1)
	}
	if localVersion == "" {
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)

	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的版本!")
	}
}
