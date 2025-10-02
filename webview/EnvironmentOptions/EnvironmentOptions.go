// 使用 WebView2 环境选项创建 WebView2 环境.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	w := window.New(0, 0, 1200, 900, "使用 WebView2 环境选项创建 WebView2 环境", 0, xcc.Window_Style_Default)

	// 创建 WebView2 环境选项.
	envOpts, err := edge.CreateEnvironmentOptions()
	if err != nil {
		wapi.MessageBoxW(0, "创建环境选项失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(3)
	}
	defer envOpts.Release()

	// 设置 WebView2 环境的语言
	envOpts.SetLanguage("en-us")

	// 构建浏览器命令行参数
	sb := strings.Builder{}
	// 允许无需用户交互的自动播放
	sb.WriteString("--autoplay-policy=no-user-gesture-required ")
	// 禁用媒体参与度检查，绕过自动播放策略
	sb.WriteString("--disable-features=PreloadMediaEngagementData,MediaEngagementBypassAutoplayPolicies ")
	// 忽略 Web Audio 的自动播放限制
	sb.WriteString("--enable-features=AutoplayIgnoreWebAudio")
	// 设置创建 WebView2 环境时要传递给浏览器进程的其它命令行参数。
	envOpts.SetAdditionalBrowserArguments(sb.String())

	// 获取 WebView2 环境选项2
	envOpts2, err := envOpts.GetICoreWebView2EnvironmentOptions2()
	if err != nil {
		log.Println("获取环境选项2失败: " + err.Error())
	} else {
		// 设置其他进程可以从使用相同用户数据文件夹创建的 WebView2 环境创建 WebView2
		envOpts2.SetExclusiveUserDataFolderAccess(true)
		envOpts2.Release()
	}

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

	// 获取 WebView2 环境选项6
	envOpts6, err := envOpts.GetICoreWebView2EnvironmentOptions6()
	if err != nil {
		log.Println("获取环境选项6失败: " + err.Error())
	} else {
		// 启用浏览器扩展功能
		envOpts6.SetAreBrowserExtensionsEnabled(true)
		envOpts6.Release()
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

	fmt.Println("------------------- WebView2 环境选项 -------------------")
	fmt.Println("语言:", envOpts.MustGetLanguage())
	fmt.Println("命令行参数:", envOpts.MustGetAdditionalBrowserArguments())
	fmt.Println("多进程共享用户数据文件夹:", envOpts2.MustGetExclusiveUserDataFolderAccess())
	fmt.Println("跟踪防护功能:", envOpts5.MustGetEnableTrackingPrevention())
	fmt.Println("浏览器扩展功能:", envOpts6.MustGetAreBrowserExtensionsEnabled())
	fmt.Println("滚动条样式:", envOpts8.MustGetScrollBarStyle())

	// 创建 webview 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder:     os.TempDir(), // 实际应用中应使用自己创建的固定目录，这里用临时目录示例
		EnvironmentOptions: envOpts,
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 创建 webview
	wv, err := edg.NewWebView(w.Handle,
		edge.WithFillParent(true), // WebView填充窗口
		edge.WithDebug(true),      // 启用调试模式
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 导航
	wv.Navigate(" https://www.bing.com/ ")

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
