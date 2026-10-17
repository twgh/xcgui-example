// 权限请求
package main

// 演示处理 WebView2 的权限请求事件(Event_PermissionRequested),
// 以网页请求地理位置权限为例: Go 侧弹出炫彩消息框询问用户允许还是拒绝, 再把结果设置到事件参数上.
//
// 注意: navigator.geolocation 是需要安全上下文(secure context)的 DOM API,
// 所以页面不能用 NavigateToString 加载(它的来源是 about:blank, 不是安全上下文),
// 否则浏览器会直接拒绝, 页面拿到 PERMISSION_DENIED, 权限请求事件也不会被触发.
// 这里和 EmbedAssets 例子一样, 采用嵌入文件系统 + 虚拟主机名的方式加载页面,
// 注意页面必须用 https:// 加载(edge.JoinUrlHeader 默认拼出来的就是 https://),
// 因为 http:// 的页面不是安全上下文, navigator.geolocation 这类 API 会被直接拒绝.
// 参考: https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/concepts/working-with-local-content

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
	//go:embed assets/**
	embedAssets embed.FS // 嵌入 assets 目录以及子目录下的文件, 不包括隐藏文件
)

const (
	hostName = "permission.example"       // 虚拟主机名
	pageName = "PermissionRequested.html" // 页面文件名
)

func main() {
	checkWebView2()
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建 WebView
	w, wv, err := edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("PermissionRequested 例子 - WebView2 权限请求处理"),
		edge.WithXmlWindowSize(1000, 720),
		edge.WithXmlWindowTitleBar(true), // 使用炫彩窗口标题栏
		edge.WithFillParent(true),        // WebView 填充窗口
		edge.WithDebug(true),             // 可打开开发者工具
		edge.WithAutoFocus(true),         // 在窗口获得焦点时尝试保持 WebView 的焦点
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 注册权限请求事件, 需要在导航前注册
	regPermissionRequested(w, wv)

	// 设置虚拟主机名和嵌入文件系统之间的映射(全局的映射, 所有的 WebView 都可以访问)
	if err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets); err != nil {
		wapi.MessageBoxW(0, "SetVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(3)
	}
	// 启用虚拟主机名和嵌入文件系统之间的映射
	if err = wv.EnableVirtualHostNameToEmbedFSMapping(true); err != nil {
		wapi.MessageBoxW(0, "EnableVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(4)
	}

	// 通过虚拟主机名加载页面(edge.JoinUrlHeader 默认拼出来的就是 https://, 页面来源是安全上下文),
	// 页面里的按钮才能触发地理位置权限请求
	if err = wv.Navigate(edge.JoinUrlHeader(hostName) + "/" + pageName); err != nil {
		wapi.MessageBoxW(0, "导航失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(5)
	}

	// 显示窗口并运行应用
	w.Show(true)
	a.Run()
	a.Exit()
}

// regPermissionRequested 注册权限请求事件.
// 当 WebView 中的内容请求访问特权资源(地理位置, 摄像头, 麦克风等)时触发,
// 在事件里通过 args.SetState 设置允许或拒绝, 不处理则保持默认行为.
func regPermissionRequested(w *window.Window, wv *edge.WebView) {
	_, _ = wv.Event_PermissionRequested(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2PermissionRequestedEventArgs) uintptr {
		// 获取请求的权限类型
		kind, err := args.GetPermissionKind()
		if err != nil {
			log.Println("GetPermissionKind 失败:", err)
			return 0
		}
		// 本例子只演示地理位置权限, 其它权限不做处理(保持默认状态)
		if kind != edge.COREWEBVIEW2_PERMISSION_KIND_GEOLOCATION {
			return 0
		}

		// 获取请求信息: 请求权限的网页来源, 是否由用户主动触发
		uri, _ := args.GetUri()
		isUserInitiated, _ := args.GetIsUserInitiated()
		log.Printf("收到权限请求: kind=%d, uri=%s, isUserInitiated=%v", kind, uri, isUserInitiated)

		// 权限事件回调在 UI 线程触发, 可以直接弹出炫彩消息框询问用户
		ret := w.MessageBox("地理位置权限请求",
			fmt.Sprintf("网页来源: %s\r\n是否由用户主动触发: %v\r\n\r\n网页想获取你的地理位置, 是否允许?", uri, isUserInitiated),
			xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Cancel|xcc.MessageBox_Flag_Icon_Qustion,
			xcc.Window_Style_Modal)

		// 把用户的选择设置到事件参数上, 告诉 WebView2 允许还是拒绝
		if ret == xcc.MessageBox_Flag_Ok {
			_ = args.SetState(edge.COREWEBVIEW2_PERMISSION_STATE_ALLOW)
			log.Println("已允许地理位置权限")
		} else {
			_ = args.SetState(edge.COREWEBVIEW2_PERMISSION_STATE_DENY)
			log.Println("已拒绝地理位置权限")
		}
		return 0
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
