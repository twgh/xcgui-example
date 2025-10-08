// 注册元素事件
/*
	炫彩的全部事件都已经在各种类里定义好了，有两种形式，分别以两种格式开头： `AddEvent` 或 `Event` ，一般使用 `AddEvent` 类型的函数即可。

	区别是：由于使用的是 `syscall.NewCallBack` 创建的事件回调函数，该方法限制只能创建 2000 个左右的回调函数，超过就会 panic。当使用 `Event` 类型的函数来注册事件且回调函数是匿名函数时，每次都会创建 1 个新的回调函数，如果不加以控制，就可能会超过 2000 个。而 `AddEvent` 类型的函数会复用创建好的回调函数，可以任意使用匿名函数作为事件回调函数，无需担心超过 2000 个的限制。
*/
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	// 创建一个按钮
	btn := widget.NewButton(50, 50, 120, 40, "button", w.Handle)

	// 注册按钮被单击事件
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		fmt.Println("按钮点击事件")
		return 0
	})

	// 注册鼠标进入事件
	btn.AddEvent_MouseStay(func(hEle int, pbHandled *bool) int {
		fmt.Println("鼠标进入事件")
		return 0
	})

	// 注册鼠标离开事件
	btn.AddEvent_MouseLeave(func(hEle int, hEleStay int, pbHandled *bool) int {
		fmt.Println("鼠标离开事件")
		return 0
	})

	// 注册鼠标滚轮滚动事件
	btn.EnableEvent_XE_MOUSEWHEEL(true)
	btn.AddEvent_MouseWheel(func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		fmt.Println("鼠标滚轮滚动事件", nFlags, pPt.X, pPt.Y)
		return 0
	})

	// 注册鼠标移动事件
	/* btn.AddEvent_MouseMove(func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		fmt.Println("鼠标移动事件", nFlags, pPt.X, pPt.Y)
		return 0
	}) */

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
