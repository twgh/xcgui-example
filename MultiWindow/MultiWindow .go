// 多窗口例子，比如你登录后销毁登录窗口载入主窗口
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	a    *app.App
	w1   *window.Window
	w2   *window.Window
	btn1 *widget.Button
	btn2 *widget.Button
)

func main() {
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)

	loadWindow1()

	a.Run()
	a.Exit()
}

func loadWindow1() {
	w1 = window.New(0, 0, 200, 200, "窗口1", 0, xcc.Window_Style_Default)
	btn1 = widget.NewButton(50, 50, 100, 30, "载入窗口2", w1.Handle)
	btn1.Event_BnClick1(onBnClick)
	w1.Show(true)
}

func loadWindow2() {
	w2 = window.New(0, 0, 300, 300, "窗口2", 0, xcc.Window_Style_Default)
	btn2 = widget.NewButton(100, 100, 100, 30, "载入窗口1", w2.Handle)
	btn2.Event_BnClick1(onBnClick)
	w2.Show(true)
}

func onBnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn1.Handle:
		*pbHandled = true // 把单击事件拦截了, 载入新窗口时这是必要的
		w1.CloseWindow()
		loadWindow2()
	case btn2.Handle:
		*pbHandled = true // 把单击事件拦截了, 载入新窗口时这是必要的
		w2.CloseWindow()
		loadWindow1()
	}
	return 0
}
