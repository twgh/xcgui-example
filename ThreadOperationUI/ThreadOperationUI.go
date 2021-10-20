// 线程操作UI, 炫彩是不能直接在线程里操作UI的, 所以要使用炫彩_调用界面线程命令来实现
package main

import (
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var win *window.Window

func main() {
	// 1.初始化UI库
	a := app.New("")
	// 2.创建窗口
	win = window.NewWindow(0, 0, 466, 300, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

	// 创建结束按钮
	btn_Close := widget.NewButton(406, 4, 50, 24, "close", win.Handle)
	btn_Close.SetType(xcc.Button_Type_Close)

	go func() {
		time.Sleep(time.Second * 2)
		// 回调函数最好不使用匿名函数
		a.CallUiThread(fun, 0)
	}()

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func fun(data int) int {
	btn := widget.NewButton(10, 33, 70, 26, "button", win.Handle)
	btn.Redraw(true)
	return 0
}
