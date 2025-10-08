// 多窗口例子，比如你登录后销毁登录窗口载入主窗口
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	loadWindow1()

	a.Run()
	a.Exit()
}

func loadWindow1() {
	w1 := window.New(0, 0, 200, 200, "窗口1", 0, xcc.Window_Style_Default)
	btn := widget.NewButton(50, 50, 100, 30, "载入窗口2", w1.Handle)
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		*pbHandled = true // 把事件拦截了, 载入新窗口时这是必要的
		w1.CloseWindow()
		loadWindow2()
		return 0
	})
	w1.Show(true)
}

func loadWindow2() {
	w2 := window.New(0, 0, 300, 300, "窗口2", 0, xcc.Window_Style_Default)
	btn := widget.NewButton(100, 100, 100, 30, "载入窗口1", w2.Handle)
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		*pbHandled = true // 把事件拦截了, 载入新窗口时这是必要的
		w2.CloseWindow()
		loadWindow1()
		return 0
	})
	w2.Show(true)
}
