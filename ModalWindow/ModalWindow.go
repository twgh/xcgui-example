// 模态窗口
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "模态窗口", 0, xcc.Window_Style_Default)

	// 创建按钮_模态窗口
	btn := widget.NewButton(30, 50, 100, 30, "ModalWindow", w.Handle)
	// 给按钮绑定事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		// 创建模态窗口
		mw := window.NewModalWindow(300, 200, "ModalWindow", w.GetHWND(), xcc.Window_Style_Modal)
		// 显示模态窗口
		mw.DoModal()
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
