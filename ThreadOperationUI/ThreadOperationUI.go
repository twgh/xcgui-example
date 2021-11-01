// 线程操作UI, 炫彩是不能直接在线程里操作UI的, 所以要使用炫彩_调用界面线程命令来实现
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var w *window.Window

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w = window.NewWindow(0, 0, 466, 300, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	btn := widget.NewButton(30, 40, 70, 30, "click", w.Handle)
	btn.Event_BnClick(func(pbHandled *bool) int {
		// go test(0) // 如果直接这样在线程中操作UI, 那将直接崩溃

		go func() {
			// 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
			a.CallUiThread(test, 0)
		}()
		return 0
	})

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func test(data int) int {
	btn := widget.NewButton(30, 80, 80, 26, "new button", w.Handle)
	btn.Redraw(true)
	return 0
}
