// 计算文件MD5.
// 不使用炫彩元素, 直接使用html文件作为窗口内容, 圆角窗口.
package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed assets/CalcMD5.xml
	xmlStr  string
	isDebug = false
)

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
	// 这个窗口是有特殊设计的, 它是透明的, 这是为了规避一个问题:
	// 当窗口最小化或最大化时会有一瞬间漏出 webview 后面的炫彩窗口, 表现出来是闪烁了一下, 所以设计为透明就看不到后面的窗口了.
	// 这个xml是通用的, 打开稍微修改下即可, content: 标题, rect: 后两个是宽高. 在设计器里打开的话会更直观.
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

	folderPath, _ := filepath.Abs("webview\\CalcMD5\\assets")
	fmt.Println("要映射的文件夹路径:", folderPath)

	// 将本地文件夹映射为虚拟域名, 比直接使用file:///访问本地文件更好.
	const hostName = "app.example"
	err = m.wv.SetVirtualHostNameToFolderMapping(hostName, folderPath, edge.COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW)
	if err != nil {
		wapi.MessageBoxW(0, "SetVirtualHostNameToFolderMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(3)
	}

	// 绑定函数, 最好在导航之前绑定
	m.bindBasicFuncs()
	m.bindFuncs()

	// 访问 HTML
	m.wv.Navigate(edge.JoinUrlHeader(hostName) + "/CalcMD5.html")
	// 显示窗口
	w.AdjustLayout()
	w.Show(true)
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

	// 创建 webview 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 自己的软件应该在固定位置创建一个自己的目录, 而不是用临时目录
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
	if ret, _ := edge.CompareBrowserVersions(edge.GetVersion(), localVersion); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}
