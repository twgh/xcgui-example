// 简单 WebView 例子
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	checkWebView2()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 1400, 900, "简单 WebView 例子", 0, xcc.Window_Style_Default)

	// 创建 webview 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 实际应用中应使用自己创建的固定目录，这里用临时目录示例
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 创建 webview
	wv, err := edg.NewWebView(w.Handle,
		edge.WithFillParent(true), // WebView 填充窗口
		edge.WithDebug(true),      // 可打开开发者工具
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 导航
	wv.Navigate("https://www.baidu.com")

	// 显示窗口并运行应用
	w.Show(true)
	a.Run()
	a.Exit()
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
