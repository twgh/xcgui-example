// 注册元素事件
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	// 创建一个按钮
	btn := widget.NewButton(50, 50, 120, 40, "button", w.Handle)

	// 所有的封装好的组件事件都是以Event_开头, 你跳转到事件定义的源码之后, 可以看到这个组件所有的事件
	// 注册按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		println("Event_BnClick")
		return 0
	})

	// 注册鼠标进入事件
	btn.Event_MOUSESTAY(func(pbHandled *bool) int {
		println("Event_MOUSESTAY")
		return 0
	})

	// 注册鼠标离开事件
	btn.Event_MOUSELEAVE(func(hEleStay int, pbHandled *bool) int {
		println("Event_MOUSELEAVE")
		return 0
	})

	// 注册鼠标滚轮滚动事件
	btn.EnableEvent_XE_MOUSEWHEEL(true)
	btn.Event_MOUSEWHEEL(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		println("Event_MOUSEWHEEL", nFlags, pPt.X, pPt.Y)
		return 0
	})

	// 注册鼠标移动事件
	/* 	btn.Event_MOUSEMOVE(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		println("Event_MOUSEMOVE", nFlags, pPt.X, pPt.Y)
		return 0
	}) */

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
