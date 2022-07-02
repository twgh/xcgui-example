// 进度条
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	bar     *widget.ProgressBar
	btn_Add *widget.Button
	btn_Sub *widget.Button
)

func main() {
	a := app.New(true)
	w := window.New(0, 0, 436, 104, "xc", 0, xcc.Window_Style_Default)

	// 创建一个进度条
	bar = widget.NewProgressBar(24, 60, 200, 10, w.Handle)
	// 设置进度条边框大小
	bar.SetBorderSize(1, 1, 1, 1)
	// 设置进度条不显示进度文字
	bar.EnableShowText(false)
	// 设置进度条最大值
	bar.SetRange(100)
	// 设置进度条进度为0
	bar.SetPos(0)

	// 创建按钮_进度加
	btn_Add = widget.NewButton(238, 50, 70, 30, "+", w.Handle)
	btn_Add.Event_BnClick1(onBtnClick)
	// 创建按钮_进度减
	btn_Sub = widget.NewButton(318, 50, 70, 30, "-", w.Handle)
	btn_Sub.Event_BnClick1(onBtnClick)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 事件_按钮被单击
func onBtnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn_Add.Handle:
		bar.SetPos(bar.GetPos() + 10)
	case btn_Sub.Handle:
		bar.SetPos(bar.GetPos() - 10)
	}
	bar.Redraw(true)
	return 0
}
