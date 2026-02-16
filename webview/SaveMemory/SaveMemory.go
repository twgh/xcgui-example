// 简单 WebView 例子
package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
)

func main() {
	checkWebView2()
	// 创建 WebView 环境
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建 WebView
	w, wv, err := edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("简单 WebView 例子"), // 窗口标题
		edge.WithXmlWindowSize(1400, 900),        // 窗口大小
		edge.WithXmlWindowTitleBar(true),         // 使用炫彩窗口标题栏
		edge.WithFillParent(true),                // WebView 填充窗口
		edge.WithDebug(true),                     // 可打开开发者工具
		edge.WithAutoFocus(true),                 // 在窗口获得焦点时尝试保持 WebView 的焦点
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 节省内存
	saveMemory(w, wv)

	// 导航
	wv.Navigate("https://www.baidu.com")

	w.Show(true)
	a.Run()
	a.Exit()
}

// 节省内存
func saveMemory(w *window.Window, wv *edge.WebView) {
	// 设置内存使用目标级别, 能节省近一半占用内存
	wv.WebView2_19 = wv.CoreWebView.MustGetICoreWebView2_19()
	if wv.WebView2_19 != nil {
		wv.WebView2_19.SetMemoryUsageTargetLevel(edge.COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_LOW)
	}

	// 挂起后, 内存占用会到个位数
	var suspendTimer *time.Timer
	var suspendTime = 10 * time.Second
	// 最小化一段时间后挂起 WebView 以节省内存.
	// 注意: 如果正在导航, 则不应挂起, 否则等导航完成, 恢复时 WebView 将不可见.
	// 这里我没有写判断的代码, 如果要写的话可以结合导航开始/完成事件来判断.
	w.AddEvent_WindProc(func(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
		if wv.Controller == nil { // WebView 未创建或已被销毁
			return 0
		}

		if message == wapi.WM_SIZE {
			switch wParam {
			case wapi.SIZE_MINIMIZED: // 窗口最小化
				// 设置定时器
				if suspendTimer != nil {
					suspendTimer.Stop()
				}
				suspendTimer = time.AfterFunc(suspendTime, func() {
					xc.UI(func() { // 需在 UI 线程操作 WebView
						// 挂起前要隐藏 WebView 才行
						wv.Show(false)
						fmt.Println("开始挂起")

						err := wv.TrySuspend(func(errorCode syscall.Errno, isSuccessful bool) uintptr {
							if !wapi.IsOK(errorCode) || !isSuccessful {
								log.Println("挂起失败, errorCode:", errorCode, ", isSuccessful:", isSuccessful)
							} else {
								fmt.Println("挂起成功")
							}
							return 0
						})
						if err != nil {
							log.Println("执行 TrySuspend 失败, err:", err)
						}
					})
				})
				fmt.Printf("窗口最小化，%v 内恢复可避免挂起\n", suspendTime)

			case wapi.SIZE_RESTORED: // 窗口恢复
				// 取消挂起定时器
				if suspendTimer != nil {
					suspendTimer.Stop()
					suspendTimer = nil
					fmt.Println("已停止挂起定时器")
				}

				if wv.IsSuspended() {
					wv.Resume()
					wv.Show(true)
					fmt.Println("从挂起状态恢复正常")
				}
			}
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
