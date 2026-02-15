// 嵌入资源到程序, 以便打包为单文件, 无需启动 http 服务器.
// 不使用炫彩元素, 直接使用html文件作为窗口内容, html代码由AI生成.
package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	isDebug = false // 是否为调试版

	//go:embed assets/**
	embedAssets embed.FS // 嵌入 assets 目录以及子目录下的文件, 不包括隐藏文件

	/* 如果想包含目录中所有文件, 包括隐藏文件(以 . 或 _ 开头的文件和目录), 可以使用下面的写法.
	//go:embed all:assets/**
	*/
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
		edge.WithXmlWindowTitle("我的应用"),
		edge.WithXmlWindowClassName("EmbedAssets"),
		edge.WithXmlWindowSize(1300, 900),
		edge.WithFillParent(true),
		edge.WithAppDrag(true),
		edge.WithDebug(isDebug),
		edge.WithDefaultContextMenus(isDebug),
		edge.WithBrowserAcceleratorKeys(isDebug),
		edge.WithStatusBar(false),
		edge.WithZoomControl(false),
		edge.WithAutoFocus(true),
		// 设置默认背景色为透明, 这也是防止首次加载会闪烁的一部分
		edge.WithDefaultBackgroundColor(edge.NewColor(0, 0, 0, 0)),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	if !isDebug { // 正式版采用嵌入文件系统的方式
		// 设置虚拟主机名和嵌入文件系统之间的映射
		err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
		if err != nil {
			wapi.MessageBoxW(0, "SetVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(6)
		}
		// 启用虚拟主机名和嵌入文件系统之间的映射
		err = m.wv.EnableVirtualHostNameToEmbedFSMapping(true)
		if err != nil {
			wapi.MessageBoxW(0, "EnableVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(5)
		}
	} else { // 调试版采用文件夹映射的方式方便热重载
		folderPath, _ := filepath.Abs("webview/EmbedAssets/assets")
		err = m.wv.SetVirtualHostNameToFolderMapping(hostName,
			folderPath, edge.COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW)
		if err != nil {
			wapi.MessageBoxW(0, "SetVirtualHostNameToFolderMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(5)
		}
		// 注意: 并不会自动热重载, 需要手动刷新页面.
		// 由于只是简单的静态页面, 可以让 AI 实现一个, 很简单的, 比如用 https://github.com/fsnotify/fsnotify 监听前端文件变化, 然后刷新页面.
		// 或者查看 VueAndVite 例子.
	}

	// 注册 WebView 事件
	m.regWebViewEvents()
	// 绑定函数
	m.bindBasicFuncs()
	// 导航到首页
	m.wv.Navigate(edge.JoinUrlHeader(hostName) + "/EmbedAssets.html")
	return m
}

// 注册 WebView 事件
func (m *MainWindow) regWebViewEvents() {
	var firstLoad = true
	// 导航完成事件
	m.wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		uri, err := sender.GetSource()
		if err != nil {
			log.Println("GetSource 失败: " + err.Error())
			return 0
		}
		fmt.Println("导航完成:", uri)

		switch uri {
		case edge.JoinUrlHeader(hostName) + "/EmbedAssets.html":
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

// bindBasicFuncs 绑定基本函数.
func (m *MainWindow) bindBasicFuncs() {
	// 绑定 最小化窗口函数
	m.wv.Bind("wnd.minimize", func() {
		m.w.ShowWindow(xcc.SW_MINIMIZE)
	})

	// 绑定 切换最大化窗口函数
	m.wv.Bind("wnd.toggleMaximize", func() {
		m.w.MaxWindow(!m.w.IsMaxWindow())
	})

	// 绑定 关闭窗口函数
	m.wv.Bind("wnd.close", func() {
		m.w.CloseWindow()
	})
}

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

func createEdge() *edge.Edge {
	// 创建 webview 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 自己的软件应该在指定位置创建个固定目录
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
