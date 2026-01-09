// 在窗口中创建 WebView, 仍然使用炫彩窗口的标题栏.
package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/xc"
)

func main() {
	checkWebView2()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建 WebView 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 自己的软件应该在固定位置创建一个自己的目录
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 创建 WebView
	w, wv, err := edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("在窗口中创建 WebView, 仍然使用炫彩窗口的标题栏"),
		// 设置炫彩 XML 窗口是否启用标题栏.
		edge.WithXmlWindowTitleBar(true),
		// 设置炫彩 XML 窗口标题栏背景颜色.
		edge.WithXmlWindowTitleBarBgColor(xc.RGBA(17, 17, 26, 255)),
		edge.WithXmlWindowSize(1400, 900),
		edge.WithFillParent(true),
		edge.WithDebug(true),
		edge.WithAutoFocus(true),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 设置炫彩 XML 窗口标题栏背景颜色.
	// wv.SetXmlWindowTitleBarBgColor(xc.RGBA(87, 161, 162, 255))

	// 导航到指定网页
	wv.Navigate("https://www.vben.pro/")

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
