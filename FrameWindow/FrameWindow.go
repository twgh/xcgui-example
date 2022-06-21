// 框架窗口
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	hPane_left   *widget.Pane
	hPane_right  *widget.Pane
	hPane_bottom *widget.Pane

	hEdit *widget.Edit
)

func main() {
	a := app.New(true)
	w := window.NewFrameWindow(0, 0, 1000, 800, "FrameWindow", 0, xcc.Window_Style_Default)

	// 创建窗格
	hPane_left = widget.NewPane("left", 200, 280, w.Handle)
	hPane_right = widget.NewPane("right", 200, 280, w.Handle)
	hPane_bottom = widget.NewPane("bottom", 770, 170, w.Handle)

	// 添加窗格
	w.AddPane(0, hPane_left.Handle, xcc.Pane_Align_Left)
	w.AddPane(0, hPane_right.Handle, xcc.Pane_Align_Right)
	w.AddPane(0, hPane_bottom.Handle, xcc.Pane_Align_Bottom)

	// 创建编辑框
	hEdit = widget.NewEdit(0, 0, 0, 0, w.Handle)
	hEdit.EnableMultiLine(true)
	// 设置主视图为编辑框
	w.SetView(hEdit.Handle)

	// 窗口_调整布局
	w.AdjustLayout()

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
