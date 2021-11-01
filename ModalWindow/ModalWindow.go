// 模态窗口
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

	// 创建按钮_模态窗口
	btn := widget.NewButton(30, 50, 100, 30, "ModalWindow", w.Handle)
	// 给按钮绑定事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		// 创建模态窗口
		win_Modal := window.NewModalWindow(300, 200, "ModalWindow", w.GetHWND(), xcc.Window_Style_Modal_Simple|xcc.Window_Style_Drag_Window|xcc.Window_Style_Btn_Close)
		// 显示模态窗口
		win_Modal.DoModal()
		return 0
	})

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
