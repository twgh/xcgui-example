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
	var err error

	// 创建 WebView
	m.w, m.wv, err = m.edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("计算文件MD5"),       // 炫彩 XML 窗口标题
		edge.WithXmlWindowClassName("CalcMD5"),   // 炫彩 XML 窗口类名
		edge.WithXmlWindowSize(550, 500),         // 炫彩 XML 窗口大小
		edge.WithXmlWindowShadowAngleSize(8),     // 炫彩 XML 窗口阴影圆角大小, 设置后会使窗口变为圆角
		edge.WithFillParent(true),                // 填充父
		edge.WithAppDrag(true),                   // 启用非客户区域支持
		edge.WithStatusBar(false),                // 禁用状态栏
		edge.WithZoomControl(false),              // 禁用缩放控件
		edge.WithDebug(isDebug),                  // 开发者工具
		edge.WithDefaultContextMenus(isDebug),    // 上下文菜单
		edge.WithBrowserAcceleratorKeys(isDebug), // 浏览器快捷键
		edge.WithRoundRadius(8),                  // WebView 圆角8px
		edge.WithAutoFocus(true),                 // 在窗口获得焦点时尝试保持 WebView 的焦点
		// 设置默认背景色为透明
		edge.WithDefaultBackgroundColor(edge.NewColor(255, 255, 255, 0)),
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
	m.regEvents()

	// 绑定函数, 最好在导航之前绑定
	m.bindBasicFuncs()
	m.bindFuncs()

	// 访问 HTML
	m.wv.Navigate(edge.JoinUrlHeader(hostName) + "/CalcMD5.html")
	return m
}

// 注册事件
func (m *MainWindow) regEvents() {
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
				// 使 html 中的输入框获取焦点, 当然你也可以在前端文件中设置焦点
				m.wv.Eval(`document.getElementById('filePath').focus()`)
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

		if objCount == 0 {
			return 0
		}

		var files []*edge.ICoreWebView2File
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

			files = append(files, file)
		}

		if len(files) == 0 {
			return 0
		}

		// 获取文件路径, 目前只处理第一个文件
		filePath, err := files[0].GetPath()
		if err != nil {
			log.Println("GetPath 失败: " + err.Error())
			return 0
		}
		fmt.Println("文件路径:", filePath)

		// 将路径传进 js 函数, 这个路径中的 \ 得转义
		m.wv.Eval(`calculate('` + strings.ReplaceAll(filePath, `\`, `\\`) + `');`)

		// 释放文件对象
		for i := range files {
			files[i].Release()
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
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	NewMainWindow(edg)

	a.Run()
	a.Exit()
}

func createEdge() *edge.Edge {
	// 创建 WebView2 环境选项.
	envOpts, err := edge.CreateEnvironmentOptions()
	if err != nil {
		log.Println("创建 WebView2 环境选项失败: " + err.Error())
	} else {
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

	if envOpts != nil { // 没用了, 直接释放
		envOpts.Release()
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
