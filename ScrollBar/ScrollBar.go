// 滚动条, 设置背景, 获取滚动条上的三个按钮并加以改变
package main

import (
	"fmt"
	"github.com/twgh/xcgui/xc"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "ScrollBar", 0, xcc.Window_Style_Default)

	// 创建滚动条
	bar1 := widget.NewScrollBar(12, 33, 300, 20, w.Handle)
	bar2 := widget.NewScrollBar(330, 33, 20, 240, w.Handle)

	// 设置为垂直滚动条
	bar2.EnableHorizon(false)
	// 添加背景
	bar2.AddBkFill(xcc.Element_State_Flag_Leave, xc.RGBA(247, 248, 250, 255))

	// 获取滑块按钮
	btnSlider := widget.NewButtonByHandle(bar2.GetButtonSlider())
	btnSlider.AddBkFill(xcc.Button_State_Flag_Leave, xc.RGBA(221, 221, 223, 255))
	btnSlider.AddBkFill(xcc.Button_State_Flag_Stay, xc.RGBA(202, 202, 204, 255))
	btnSlider.AddBkFill(xcc.Button_State_Flag_Down, xc.RGBA(202, 202, 204, 255))
	// 获取滚动条上按钮
	btnUp := widget.NewButtonByHandle(bar2.GetButtonUp())
	btnUp.AddBkFill(xcc.Button_State_Flag_Leave, xc.RGBA(137, 140, 151, 255))
	btnUp.AddBkFill(xcc.Button_State_Flag_Stay, xc.RGBA(255, 135, 250, 255))
	btnUp.AddBkFill(xcc.Button_State_Flag_Down, xc.RGBA(255, 75, 250, 255))
	// 获取滚动条下按钮
	btnDown := widget.NewButtonByHandle(bar2.GetButtonDown())
	btnDown.AddBkFill(xcc.Button_State_Flag_Leave, xc.RGBA(137, 140, 151, 255))
	btnDown.AddBkFill(xcc.Button_State_Flag_Stay, xc.RGBA(255, 135, 250, 255))
	btnDown.AddBkFill(xcc.Button_State_Flag_Down, xc.RGBA(255, 75, 250, 255))
	// 因为上下按钮背景改变了, 你可以自己准备图片设置到按钮上去
	// btnDown.SetIcon()  或  btnDown.AddBkImage()

	// 注册滚动条元素滚动事件
	bar1.Event_SBAR_SCROLL1(SBAR_SCROLL1)
	bar2.Event_SBAR_SCROLL1(SBAR_SCROLL1)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 滚动条元素滚动事件
func SBAR_SCROLL1(hEle int, pos int32, pbHandled *bool) int {
	fmt.Println(pos)
	// 为了鼠标滚轮滚动和点击两端按钮实时显示效果而刷新
	xc.XEle_Redraw(hEle, false)
	return 0
}
