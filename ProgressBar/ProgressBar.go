// 进度条
package main

import (
	_ "embed"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	bar    *widget.ProgressBar
	bar2   *widget.ProgressBar
	btnAdd *widget.Button
	btnSub *widget.Button
)

//go:embed jindu.png
var img []byte

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 436, 450, "ProgressBar", 0, xcc.Window_Style_Default)

	// 创建一个水平进度条
	bar = widget.NewProgressBar(24, 60, 200, 10, w.Handle)
	// 设置进度条边框大小
	bar.SetBorderSize(1, 1, 1, 1)
	// 设置进度条不显示进度文字
	bar.EnableShowText(false)
	// 设置进度条最大值
	bar.SetRange(100)
	// 设置进度条进度
	bar.SetPos(40)
	// 置进度颜色
	bar.SetColorLoad(xc.RGBA(43, 170, 255, 255))
	// 置进度条背景颜色
	bar.AddBkFill(xcc.Element_State_Flag_Leave, xc.RGBA(221, 221, 223, 255))

	bar2 = widget.NewProgressBar(24, 200, 24, 200, w.Handle)
	// 设置为垂直进度条
	bar2.EnableHorizon(false)
	// 设置进度条边框大小
	bar2.SetBorderSize(0, 0, 0, 0)
	// 不显示进度文本
	bar2.EnableShowText(false)
	// 置进度图片
	bar2.SetImageLoad(imagex.NewByMemAdaptive(img, 0, 0, 0, 0).Handle)
	// 置进度条背景颜色
	bar2.AddBkFill(xcc.Element_State_Flag_Leave, xc.RGBA(221, 221, 223, 255))

	// 创建按钮_进度加
	btnAdd = widget.NewButton(238, 50, 70, 30, "+", w.Handle)
	btnAdd.Event_BnClick1(onBtnClick)
	// 创建按钮_进度减
	btnSub = widget.NewButton(318, 50, 70, 30, "-", w.Handle)
	btnSub.Event_BnClick1(onBtnClick)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 事件_按钮被单击
func onBtnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btnAdd.Handle:
		bar.SetPos(bar.GetPos() + 10)
		bar2.SetPos(bar.GetPos() + 10)
	case btnSub.Handle:
		bar.SetPos(bar.GetPos() - 10)
		bar2.SetPos(bar.GetPos() - 10)
	}
	bar.Redraw(true)
	bar2.Redraw(true)
	return 0
}
