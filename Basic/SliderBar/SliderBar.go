// 滑块条.
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
	sb.SetRange(10)

	// 设置滑块按钮高度和宽度
	sb.SetButtonHeight(27)
	sb.SetButtonWidth(27)

	// 启用背景透明
	sb.EnableBkTransparent(true)

	// 注册滑块位置改变事件
	sb.AddEvent_SliderBar_Change(func(hEle int, pos int32, pbHandled *bool) int {
		fmt.Println("滑块位置改变:", pos)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
