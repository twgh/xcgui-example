// 滑块条
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 430, 300, "SliderBar", 0, xcc.Window_Style_Default)

	// 创建SliderBar
	sb := widget.NewSliderBar(12, 33, 300, 60, w.Handle)
	// 设置滑动范围
	sb.SetRange(100)

	// 设置滑块按钮高度和宽度
	sb.SetButtonHeight(27)
	sb.SetButtonWidth(27)

	// 启用背景透明
	sb.EnableBkTransparent(true)
	// 禁用绘制焦点
	sb.EnableDrawFocus(false)

	// 创建形状文本, 实时显示滑块当前值
	stValue := widget.NewShapeText(330, 48, 90, 24, "40", w.Handle)

	// 创建ProgressBar, 与滑块联动, 范围和滑块一致
	pb := widget.NewProgressBar(12, 120, 400, 26, w.Handle)
	pb.SetRange(100)        // 与滑块范围保持一致
	pb.EnableShowText(true) // 在进度条上显示进度文本

	// 注册滑块位置改变事件
	sb.AddEvent_SliderBar_Change(func(hEle int, pos int32, pbHandled *bool) int {
		// 把当前值显示到窗口上的形状文本里
		stValue.SetText(fmt.Sprintf("%d", pos))
		stValue.Redraw() // 修改显示后必须重绘
		// 让进度条跟随滑块位置
		pb.SetPos(pos)
		pb.Redraw(false)
		return 0
	})

	// 纵向滑块条
	sb2 := widget.NewSliderBar(312, 150, 30, 120, w.Handle)
	// 设置水平或垂直, true为水平, false为垂直
	sb2.EnableHorizon(false)
	sb2.SetRange(100)
	sb2.SetButtonHeight(16)
	sb2.EnableBkTransparent(true)
	sb2.EnableDrawFocus(false)

	// 设置初始位置
	sb.SetPos(40)
	pb.SetPos(40)

	w.Show(true)
	a.Run()
	a.Exit()
}
