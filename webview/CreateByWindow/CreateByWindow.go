// 在窗口中创建 webview, 仍然使用炫彩窗口的标题栏
package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
)

var (
	//go:embed res/CreateByWindow.xml
	xmlStr string
)

func main() {
	checkWebView2()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口从布局文件
	// 这个窗口是有特殊设计的, 它的主体是透明的, 这是为了规避一个问题:
	// 当窗口最小化或最大化时会有一瞬间漏出 webview 后面的炫彩窗口, 表现出来是闪烁了一下, 所以设计为透明就看不到后面的窗口了.
	w := window.NewByLayoutStringW(xmlStr, 0, 0)

	// 创建 webview 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 自己的软件应该在固定位置创建一个自己的目录, 而不是用临时目录
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 创建 webview
	wv, err := edg.NewWebView(w.Handle,
		edge.WithFillParent(true),
		edge.WithDebug(true),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 导航到指定网页
	wv.Navigate("https://www.vben.pro/")

	w.AdjustLayout()
	w.Show(true)
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
