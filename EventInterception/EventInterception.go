// 事件拦截.
// 一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数.
// 当你想拦截当前事件或不想向后传递，只需要将参数(*pbHnadled=true)即可.
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w := window.NewWindow(0, 0, 430, 300, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	// 创建一个按钮
	btn := widget.NewButton(50, 50, 70, 30, "button", w.Handle)
	// 一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数.
	// 当你想拦截当前事件或不想向后传递，只需要将参数(*pbHnadled=true)即可.
	btn.Event_BnClick(event1)
	btn.Event_BnClick(event2)
	btn.Event_BnClick(event3)

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func event1(pbHandled *bool) int {
	println("event1")
	return 0
}

func event2(pbHandled *bool) int {
	println("event2")
	return 0
}

func event3(pbHandled *bool) int {
	println("event3")
	*pbHandled = true
	return 0
}
