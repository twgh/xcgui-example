// 计算文件MD5.
// 不使用炫彩元素, 直接使用html文件作为窗口内容, 圆角窗口.
package main

import (
	"crypto/md5"
	"embed"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"unsafe"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed res/CalcMD5.xml
	xmlStr string
	//go:embed assets/**
	embedAssets embed.FS // 嵌入 assets 目录以及子目录下的文件, 不包括隐藏文件

	isDebug = false
)

const hostName = "app.example"

type MainWindow struct {
	edg *edge.Edge
	w   *window.Window
	wv  *edge.WebView
}

func NewMainWindow(edg *edge.Edge) *MainWindow {
	m := &MainWindow{edg: edg}
	m.main()
	return m
}

func (m *MainWindow) main() {
	var err error
	// 创建窗口
	// 这个窗口是有特殊设计的, 它是透明的, 这是为了规避当窗口类型为[透明窗口阴影]时的一个问题:
	// 当窗口最小化或最大化时会有一瞬间漏出 webview 后面的炫彩窗口, 表现出来是闪烁了一下, 所以设计为透明就看不到后面的窗口了. 只有网页颜色和炫彩窗口颜色相差很大时才会容易看出此问题, 这是很追求细节的才会注意到的.
	// 这个xml是通用的, 打开稍微修改下即可, content: 标题, rect: 后两个是宽高, 用窗口函数来设置也行. 在设计器里打开的话会更直观.
	w := window.NewByLayoutStringW(xmlStr, 0, 0)
	m.w = w
	// 炫彩窗口圆角8px
	w.SetShadowInfo(8, 255, 8, false, 0)

	// 创建 webview
	m.wv, err = m.edg.NewWebView(w.Handle,
		edge.WithFillParent(true),
		edge.WithAppDrag(true),
		edge.WithStatusBar(false),
		edge.WithZoomControl(false),
		edge.WithDebug(isDebug),
		edge.WithDefaultContextMenus(isDebug),
		edge.WithBrowserAcceleratorKeys(isDebug),
		edge.WithRoundRadius(8), // 圆角8px
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 设置虚拟主机名和嵌入文件系统之间的映射
	edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
	// 启用虚拟主机名和嵌入文件系统之间的映射
	m.wv.EnableVirtualHostNameToEmbedFSMapping(true)

	// 注册事件
	m.regEvent()

	// 绑定函数, 最好在导航之前绑定
	m.bindBasicFuncs()
	m.bindFuncs()

	// 访问 HTML
	m.wv.Navigate(edge.JoinUrlHeader(hostName) + "/CalcMD5.html")
	// 调整窗口布局
	w.AdjustLayout()
}

// 注册事件
func (m *MainWindow) regEvent() {
	var firstLoad = true
	// 导航完成事件
	m.wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		uri := sender.MustGetSource()
		fmt.Println("导航完成:", uri)
		switch uri {
		case edge.JoinUrlHeader(hostName) + "/CalcMD5.html":
			// 在导航完成事件里判断第一次加载完毕时才显示窗口,
			// 这是因为采用嵌入文件系统的方式时, 网页还没加载出来的时候, 会显示 webview 白色的背景,
			// 然后才会加载出网页, 表现出来就是有一瞬间的闪烁, 所以等加载完再显示窗口
			if firstLoad {
				firstLoad = false
				m.w.Show(true)
			}
		}
		return 0
	})

	// 网页消息事件, 用于获取拖拽的文件路径
	m.wv.Event_WebMessageReceived(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2WebMessageReceivedEventArgs) uintptr {
		// 获取网页消息
		webMessage, err := args.TryGetWebMessageAsString()
		if err != nil {
			return 0
		}
		if webMessage != "drag_files" { // 这是前端传过来的
			return 0
		}

		args2, err := args.GetICoreWebView2WebMessageReceivedEventArgs2()
		if err != nil {
			log.Println("GetICoreWebView2WebMessageReceivedEventArgs2 失败: " + err.Error())
			return 0
		}
		defer args2.Release()

		// 获取包含随 Web 消息一起发送的附加对象的对象集合视图
		objs, err := args2.GetAdditionalObjects()
		if err != nil {
			log.Println("GetAdditionalObjects 失败: " + err.Error())
			return 0
		}
		defer objs.Release()

		// 获取集合中的对象数量
		objCount, err := objs.GetCount()
		if err != nil {
			log.Println("GetCount 失败: " + err.Error())
			return 0
		}

		// 遍历集合，查找 File 对象
		for i := uint32(0); i < objCount; i++ {
			obj, err := objs.GetValueAtIndex(i)
			if err != nil {
				log.Println("GetValueAtIndex 失败: " + err.Error())
				continue
			}

			file := new(edge.ICoreWebView2File)
			err = obj.QueryInterface(wapi.NewGUIDPointer(edge.IID_ICoreWebView2File), unsafe.Pointer(&file))
			if err != nil {
				log.Println("QueryInterface 失败: " + err.Error())
				continue
			}
			defer file.Release()

			// 获取文件路径
			filePath, err := file.GetPath()
			if err != nil {
				log.Println("GetPath 失败: " + err.Error())
				continue
			}
			fmt.Println("文件:", filePath)
			// 将路径传进 js 函数
			m.wv.Eval("calculate('" + strings.ReplaceAll(filePath, "\\", "\\\\") + "');")
			break // 目前只处理第一个文件
		}
		return 0
	})
}

// bindBasicFuncs 绑定基本函数.
func (m *MainWindow) bindBasicFuncs() {
	// 绑定 最小化窗口函数
	m.wv.Bind("go.minimizeWindow", func() {
		m.w.ShowWindow(xcc.SW_MINIMIZE)
	})

	// 绑定 切换最大化窗口函数
	m.wv.Bind("go.toggleMaximize", func() {
		m.w.MaxWindow(!m.w.IsMaxWindow())
	})

	// 绑定 关闭窗口函数
	m.wv.Bind("go.closeWindow", func() {
		m.w.CloseWindow()
	})
}

// bindFuncs 绑定函数.
func (m *MainWindow) bindFuncs() {
	// 绑定 openFile
	m.wv.Bind("go.openFile", func() string {
		return wutil.OpenFile(m.w.Handle, []string{"All Files(*.*)", "*.*"}, "")
	})

	// 绑定 calculateMD5
	m.wv.Bind("go.calculateMD5", func(filePath string) string {
		// 判断文件是否存在
		if !xc.PathExists2(filePath) {
			return "错误: 文件不存在"
		}

		var ret string
		// 这是耗时操作, 所以在协程里执行, 不卡界面
		go func() {
			// 读取文件内容
			data, err := os.ReadFile(filePath)
			if err != nil {
				ret = "错误: " + err.Error()
				return
			}

			// 计算MD5
			hash := md5.Sum(data)
			md5Str := hex.EncodeToString(hash[:])
			ret = "文件: " + filePath + "\nMD5: " + md5Str
		}()

		// 等待 md5 计算完成, 这个代码是会不卡界面的阻塞在这里, 等待协程执行完毕
		var msg wapi.MSG
		for ret == "" {
			if wapi.GetMessage(&msg, 0, 0, 0) == 0 {
				break
			}
			wapi.TranslateMessage(&msg)
			wapi.DispatchMessage(&msg)
		}
		return ret
	})
}

func main() {
	checkWebView2()
	app.InitOrExit()

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

	// 初始化界面库
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	NewMainWindow(edg)

	a.Run()
	a.Exit()
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
