// WebView局部透明, 鼠标穿透
package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	isDebug = false // 是否为调试版

	//go:embed assets/**
	embedAssets embed.FS // 嵌入 assets 目录以及子目录下的文件, 不包括隐藏文件
)

const hostName = "app.example"

type MainWindow struct {
	edg        *edge.Edge
	w          *window.Window
	wv         *edge.WebView
	layContent *widget.LayoutEle
}

func NewMainWindow(edg *edge.Edge) *MainWindow {
	m := &MainWindow{edg: edg}
	var err error

	// 如果要仅显示 WebView 内容, 可以把窗口 style 改为 Window_Style_Center, 只留个居中
	m.w = window.New(0, 0, 900, 900, "局部透明", 0, xcc.Window_Style_Default)

	// 设置为透明窗口
	m.w.SetTransparentType(xcc.Window_Transparent_Shaped)
	m.w.SetTransparentAlpha(255)

	// 创建布局元素
	m.layContent = widget.NewLayoutEle(40, 40, 650, 500, m.w.Handle)

	// 创建 WebView
	m.wv, err = m.edg.NewWebView(m.layContent.Handle,
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
	} else { // 调试版采用文件夹映射的方式方便重载, 注意: 并不会自动热重载, 需要手动刷新页面.
		folderPath, _ := filepath.Abs("./assets")
		fmt.Println("如果报错肯定是这个路径有问题, 得改下, folderPath:", folderPath)
		err = m.wv.SetVirtualHostNameToFolderMapping(hostName,
			folderPath, edge.COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW)
		if err != nil {
			wapi.MessageBoxW(0, "SetVirtualHostNameToFolderMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(5)
		}
	}

	// 注册 WebView 事件
	m.regWebViewEvents()
	// 绑定函数
	m.bindBasicFuncs()
	// 导航到首页
	m.wv.Navigate(edge.JoinUrlHeader(hostName) + "/Transparent.html")
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
		case edge.JoinUrlHeader(hostName) + "/Transparent.html":
			// 在导航完成事件里判断第一次加载完毕时才显示窗口,
			// 这是因为采用嵌入文件系统的方式时, 网页还没加载出来的时候, 会显示webview白色的背景,
			// 然后才会加载出网页, 表现出来就是有一瞬间的闪烁, 所以等加载完再显示窗口
			if firstLoad {
				firstLoad = false

				// 因为要用到 EvalSync 同步执行 js 代码并取回返回值, 而 EvalSync 是依赖于窗口消息循环的,
				// 所以要等窗口的消息循环开始了, 才能调用 EvalSync 函数, 也就是要等 app.Run 已经执行了.
				// 事件函数都是在消息循环中执行的, 所以注册这个窗口绘制事件, 等窗口在绘制的时候再执行 EvalSync 函数.
				// 有时候你可能想要窗口创建完成事件, 可以用 AddEvent_Paint_Display: 窗口绘制完成并且已经显示到屏幕
				m.w.AddEvent_Paint(func(hWindow, hDraw int, pbHandled *bool) int {
					fmt.Println("进入 Paint")
					// 这个事件只用一次, 用完直接移除
					defer m.w.RemoveEvent(xcc.WM_PAINT)

					// 获取 html 中卡片的外边距矩形
					jsonStr, err := m.wv.EvalSync("getCardRectMargin()")
					if err != nil {
						fmt.Println("getCardRectMargin 失败:", err)
						return 0
					}
					fmt.Println("getCardRectMargin jsonStr:", jsonStr)

					rc1, _ := UnmarshalRect(jsonStr)
					fmt.Println("getCardRectMargin 结果:", rc1)

					// 获取 html 中底栏的外边距矩形
					jsonStr, err = m.wv.EvalSync("getStatusBarRectMargin()")
					if err != nil {
						fmt.Println("getStatusBarRectMargin 失败:", err)
						return 0
					}
					fmt.Println("getStatusBarRectMargin jsonStr:", jsonStr)

					rc2, _ := UnmarshalRect(jsonStr)
					fmt.Println("getStatusBarRectMargin 结果:", rc2)

					// 设置为透明窗口后, 整个窗口都会鼠标穿透了, 包括 WebView 不透明的地方也是, 但我们的需求是要 WebView 不透明的地方是可以点击的.
					// 之所以这样是因为炫彩窗口本体现在是全透明的, 没有不透明的地方, WebView 不透明的地方并不是在炫彩窗口上的, 所以我们需要在炫彩窗口上添加不透明的区域,
					// 而且不透明的区域还得正好对上 WebView 不透明的地方, 正好在 WebView 不透明的内容下方, 这样才能实现透明的地方是鼠标穿透的, 不透明的地方是可以点击的.
					{
						// 获取背景管理对象
						bkm := m.layContent.GetBkManagerObj()
						// 添加填充矩形 (对应 card)
						bkm.AddFill(xcc.Element_State_Flag_Leave, xc.RGBA(255, 255, 255, 255), 1)
						// 添加填充矩形 (对应 statusBar)
						bkm.AddFill(xcc.Element_State_Flag_Leave, xc.RGBA(255, 255, 255, 255), 2)

						// 获取填充矩形, 匹配 card
						obj1 := bkm.GetObjectObj(1)
						// SetMargin 在默认对齐下, 4 个参数按 [左边间距, 上边间距, 右边间距, 下边间距] 解释;
						// HTML 端 getCardRectMargin 返回的 4 个值刚好是元素到 layContent 四边的距离, 直接对应.
						obj1.SetMargin(rc1.Left, rc1.Top, rc1.Right, rc1.Bottom)
						// card: border-radius: 16px
						obj1.SetRectRoundAngle(16, 16, 16, 16)

						// 获取填充矩形, 匹配 statusBar
						obj2 := bkm.GetObjectObj(2)
						obj2.SetMargin(rc2.Left, rc2.Top, rc2.Right, rc2.Bottom)
						// info-bar: border-radius: 12px
						obj2.SetRectRoundAngle(12, 12, 12, 12)
					}
					return 0
				})

				m.w.Show(true)
			}
		}
		return 0
	})
}

type Rect struct {
	Left   int32 `json:"left"`
	Top    int32 `json:"top"`
	Right  int32 `json:"right"`
	Bottom int32 `json:"bottom"`
}

func UnmarshalRect(jsonStr string) (Rect, error) {
	var rect Rect
	return rect, json.Unmarshal([]byte(jsonStr), &rect)
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
