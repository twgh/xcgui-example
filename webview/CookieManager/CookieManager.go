// Cookie 管理
package main

// 演示使用 ICoreWebView2CookieManager 对 cookie 进行 创建/查询/修改/删除.

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

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

const (
	hostName = "app.example" // 虚拟主机名
	ckName   = "xcgui"       // 演示用的 cookie 名称
	ckDomain = hostName      // cookie 的域
	ckPath   = "/"           // cookie 的路径
)

// hostURL 虚拟主机名的 URL 头: "http://app.example"
var hostURL = edge.JoinUrlHeader(hostName)

// 全局对象
var (
	edg   *edge.Edge
	w     *window.Window
	wv    *edge.WebView
	ckMgr *edge.ICoreWebView2CookieManager // Cookie 管理器
)

func main() {
	checkWebView2()
	edg = createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口和 WebView
	createWindow()

	// 获取 Cookie 管理器, 用于创建/查询/修改/删除 cookie.
	// 库里也提供了 wv.GetCookies/wv.SetCookie/wv.DeleteCookie 等便捷方法, 内部会自动获取并释放 CookieManager.
	mgr, err := wv.GetCookieManager()
	if err != nil {
		wapi.MessageBoxW(0, "获取 CookieManager 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(3)
	}
	ckMgr = mgr
	defer ckMgr.Release() // COM 对象不再使用时记得释放

	// 设置虚拟主机名和嵌入文件系统之间的映射
	if err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets); err != nil {
		wapi.MessageBoxW(0, "SetVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(4)
	}
	// 启用虚拟主机名和嵌入文件系统之间的映射
	if err = wv.EnableVirtualHostNameToEmbedFSMapping(true); err != nil {
		wapi.MessageBoxW(0, "EnableVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(5)
	}

	// 注册 WebView 事件
	regWebViewEvents()
	// 导航到嵌入的页面
	wv.Navigate(hostURL + "/CookieManager.html")

	w.Show(true)
	a.Run()
	a.Exit()
}

// createWindow 创建窗口: 顶部是炫彩按钮工具栏, 工具栏下方是 WebView.
func createWindow() {
	// 创建窗口
	w = window.New(0, 0, 1000, 680, "Cookie 管理 (WebView2)", 0, xcc.Window_Style_Default)
	// 缩小窗口四边框, Top 为标题栏高度, 默认边框很宽不美观
	w.SetBorderSize(2, 32, 2, 2)
	// 启用窗口布局并设为垂直方向: 工具栏在顶部, WebView 占据剩余空间, 拉伸窗口时自动调整
	w.EnableLayout(true)
	w.EnableHorizon(false)
	w.SetPadding(6, 6, 6, 6)

	// 工具栏: 布局元素, 内部水平排列按钮
	toolBar := widget.NewLayoutEle(0, 0, 0, 36, w.Handle)
	toolBar.LayoutItem_SetWidth(xcc.Layout_Size_Fill, 0) // 宽度填满一行
	toolBar.SetAlignV(xcc.Layout_Align_Center)           // 子元素垂直居中
	toolBar.SetSpace(6)                                  // 子元素间距

	// 添加Cookie按钮: CreateCookie + AddOrUpdateCookie
	btnAdd := widget.NewButton(0, 0, 96, 28, "添加Cookie", toolBar.Handle)
	btnAdd.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		addCookie()
		return 0
	})

	// 修改值按钮: 查询cookie → 修改属性(SetValue) → AddOrUpdateCookie 应用修改
	btnModify := widget.NewButton(0, 0, 96, 28, "修改值", toolBar.Handle)
	btnModify.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		modifyCookie()
		return 0
	})

	// 读取列表按钮: GetCookiesEx 查询指定 URI 的所有 cookie 并读取属性
	btnList := widget.NewButton(0, 0, 96, 28, "读取列表", toolBar.Handle)
	btnList.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		listCookies()
		return 0
	})

	// 删除Cookie按钮: CreateCookie 同名同域同路径的 cookie, 然后 DeleteCookie 匹配删除
	btnDelete := widget.NewButton(0, 0, 96, 28, "删除Cookie", toolBar.Handle)
	btnDelete.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		deleteCookie()
		return 0
	})

	// 清空按钮: DeleteAllCookies 删除当前配置文件下的所有 cookie
	btnClear := widget.NewButton(0, 0, 96, 28, "清空全部", toolBar.Handle)
	btnClear.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		clearCookies()
		return 0
	})

	// WebView 宿主: 布局元素占据工具栏下方剩余空间
	layoutWV := widget.NewLayoutEle(0, 0, 0, 0, w.Handle)
	layoutWV.LayoutItem_SetWidth(xcc.Layout_Size_Fill, 0)    // 宽度填满
	layoutWV.LayoutItem_SetHeight(xcc.Layout_Size_Weight, 1) // 高度占据剩余空间

	// 创建 WebView 到布局元素里
	var err error
	wv, err = edg.NewWebView(layoutWV.Handle,
		edge.WithFillParent(true), // WebView 填充布局元素
		edge.WithDebug(isDebug),
		edge.WithAutoFocus(true),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}
	wv.Show(false) // 首次导航完成后再显示, 避免白屏闪烁

	// 调整布局
	w.AdjustLayout()
}

// regWebViewEvents 注册 WebView 事件.
func regWebViewEvents() {
	firstLoad := true
	// 导航完成事件: 首次加载完成后再显示 WebView, 避免白屏闪烁
	wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		log.Println("导航完成:", sender.MustGetSource())
		if firstLoad {
			firstLoad = false
			// 回调里操作炫彩 UI, 用 xc.UI() 包裹
			xc.UI(func() {
				_ = wv.Show(true)
			})
		}
		return 0
	})
}

// addCookie 演示 CreateCookie + AddOrUpdateCookie.
func addCookie() {
	value := fmt.Sprintf("你好炫彩_%d", time.Now().Unix()) // cookie 值
	// 创建 cookie 对象: 名称/值/域/路径
	cookie, err := ckMgr.CreateCookie(ckName, value, ckDomain, ckPath)
	if err != nil {
		w.MessageBox("错误", "CreateCookie 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return
	}
	defer cookie.Release() // COM 对象用完释放
	// 添加到 WebView: 名称/域/路径都相同时会覆盖旧值
	if err = ckMgr.AddOrUpdateCookie(cookie); err != nil {
		w.MessageBox("错误", "AddOrUpdateCookie 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return
	}
	refreshPage()
}

// modifyCookie 演示 查询cookie → 修改属性(SetValue) → AddOrUpdateCookie 把修改应用到 WebView.
func modifyCookie() {
	// GetCookies 是异步的, 结果通过回调返回; uri 为空字符串时返回同一配置文件下的所有 cookie.
	err := ckMgr.GetCookiesEx(&wv.WebViewEventImpl, hostURL+ckPath, func(errorCode syscall.Errno, cookies *edge.ICoreWebView2CookieList) uintptr {
		if !wapi.IsOK(errorCode) {
			log.Println("GetCookies 失败, errorCode:", errorCode)
			return 0
		}
		// 注意: cookieList 归 WebView2 运行时所有, 仅在本次回调期间有效, 不要对它调用 Release,
		// 否则运行时在回调返回后继续使用它时会崩溃(访问冲突).
		// 只有 GetValueAtIndex 返回的每个 cookie 归调用方所有, 用完才需要 Release.
		count, _ := cookies.GetCount()
		for i := uint32(0); i < count; i++ {
			cookie, err := cookies.GetValueAtIndex(i)
			if err != nil {
				continue
			}
			func() {
				defer cookie.Release() // COM 对象用完释放
				name, _ := cookie.GetName()
				if name != ckName {
					return
				}
				// 修改 cookie 的值
				if err := cookie.SetValue(fmt.Sprintf("新值_%d", time.Now().Unix())); err != nil {
					log.Println("SetValue 失败:", err)
					return
				}
				// 把修改应用到 WebView
				if err := ckMgr.AddOrUpdateCookie(cookie); err != nil {
					log.Println("AddOrUpdateCookie 失败:", err)
					return
				}
				refreshPage()
			}()
		}
		return 0
	})
	if err != nil {
		w.MessageBox("错误", "GetCookiesEx 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
	}
}

// listCookies 演示 GetCookiesEx 查询指定 URI 匹配的所有 cookie, 读取属性后显示到弹窗.
func listCookies() {
	err := ckMgr.GetCookiesEx(&wv.WebViewEventImpl, hostURL+ckPath, func(errorCode syscall.Errno, cookies *edge.ICoreWebView2CookieList) uintptr {
		if !wapi.IsOK(errorCode) {
			log.Println("GetCookies 失败, errorCode:", errorCode)
			return 0
		}
		// cookieList 归 WebView2 运行时所有, 仅在本次回调期间有效, 不要对它调用 Release.
		// 只有 GetValueAtIndex 返回的每个 cookie 归调用方所有, 用完才需要 Release.
		var sb strings.Builder
		count, _ := cookies.GetCount()
		for i := uint32(0); i < count; i++ {
			cookie, err := cookies.GetValueAtIndex(i)
			if err != nil {
				continue
			}
			func() {
				defer cookie.Release() // COM 对象用完释放
				// 读取 cookie 的各项属性
				name, _ := cookie.GetName()
				value, _ := cookie.GetValue()
				domain, _ := cookie.GetDomain()
				path, _ := cookie.GetPath()
				isSession, _ := cookie.GetIsSession()
				expires := "会话 cookie"
				if !isSession {
					exp, _ := cookie.GetExpires()
					expires = "到期: " + time.Unix(int64(exp), 0).Format("2006-01-02 15:04:05")
				}
				fmt.Fprintf(&sb, "%s=%s\r\n域: %s  路径: %s  %s\r\n\r\n", name, value, domain, path, expires)
			}()
		}
		text := strings.TrimSpace(sb.String())
		if text == "" {
			text = "(该 URI 下没有 cookie)"
		}

		// 回调里操作炫彩 UI, 用 xc.UI() 包裹
		xc.UI(func() {
			refreshPage()
			w.MessageBox("当前 cookie 列表 ("+hostURL+ckPath+")", text, xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		})
		return 0
	})
	if err != nil {
		w.MessageBox("错误", "GetCookiesEx 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
	}
}

// deleteCookie 演示 DeleteCookie: 按 名称+域/路径 匹配删除.
func deleteCookie() {
	// DeleteCookie 通过比较 cookie 的名称和域/路径来匹配, 所以先创建一个匹配的 cookie 对象.
	cookie, err := ckMgr.CreateCookie(ckName, "", ckDomain, ckPath)
	if err != nil {
		w.MessageBox("错误", "CreateCookie 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return
	}
	defer cookie.Release() // COM 对象用完释放
	if err = ckMgr.DeleteCookie(cookie); err != nil {
		w.MessageBox("错误", "DeleteCookie 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return
	}

	// DeleteCookie 是异步生效的: 删除操作要过一会才反映到页面的 document.cookie 上,
	// 删完立刻刷新页面的话, 页面读到的往往还是删除前的值(现象就是点击删除后页面没有变化),
	// 所以这里延迟一下再刷新, 保证页面显示的是删除后的结果.
	// 对比: DeleteAllCookies 是立即生效的, 可以直接刷新, 见 clearCookies().
	time.AfterFunc(100*time.Millisecond, refreshPage)
}

// clearCookies 演示 DeleteAllCookies: 删除同一配置文件下的所有 cookie.
func clearCookies() {
	if err := ckMgr.DeleteAllCookies(); err != nil {
		w.MessageBox("错误", "DeleteAllCookies 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return
	}
	// DeleteAllCookies 是立即生效的, 直接刷新页面即可
	refreshPage()
}

// refreshPage 调用页面里的 js 函数, 让页面重新显示 document.cookie.
// WebView 的 Eval 必须在 UI 线程调用, 所以用 xc.UI() 包裹, 因此本函数可以在任意协程中调用.
func refreshPage() {
	xc.UI(func() {
		if wv.CoreWebView == nil {
			return
		}
		if err := wv.Eval("if (window.renderCookie) renderCookie()"); err != nil {
			log.Println("Eval 失败:", err)
		}
	})
}

// 创建 WebView 环境
func createEdge() *edge.Edge {
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 实际应用中应使用自己创建的固定目录
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
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
