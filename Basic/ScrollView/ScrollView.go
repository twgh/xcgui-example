// 滚动视图
package main

// 滚动视图(ScrollView)是一个带有滚动条的容器元素, 当内容区域比视口大时,
// 可以通过水平/垂直滚动条滚动查看超出视口的内容,
// 常用于大图查看、画板、自定义可滚动的面板等场景.
//
// 本例演示: 在 ScrollView 里放一个比视口更大的内容区域(一个背景色不同的元素和若干按钮),
// 演示内容尺寸(SetTotalSize)、滚动单位(SetLineSize)、滚动条大小的设置,
// 以及滚动事件和代码控制滚动.

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 560, 480, "ScrollView - 滚动视图", 0, xcc.Window_Style_Default)

	// 说明文本
	widget.NewShapeText(20, 34, 520, 24, "拖动滚动条或滚轮查看超出视口的内容:", w.Handle)

	// 创建滚动视图, 视口大小 400x300
	sv := widget.NewScrollView(20, 68, 400, 300, w.Handle)
	// 设置内容区(视图)大小为 700x520, 比视口大, 因此会出现滚动条
	sv.SetTotalSize(700, 520)
	// 设置滚动单位大小, 每次滚动一行的大小
	sv.SetLineSize(20, 20)
	// 设置滚动条宽度
	sv.SetScrollBarSize(14)
	// 启用自动显示滚动条: 内容超出视口时自动出现, 否则自动隐藏
	sv.EnableAutoShowScrollBar(true)

	// 在滚动视图里创建一个比视口大的内容元素, 用不同背景色区分内容区域
	eleContent := widget.NewElement(0, 0, 700, 520, sv.Handle)
	eleContent.AddBkFill(xcc.Element_State_Flag_Leave, xc.RGBA(226, 235, 244, 255))

	// 在内容区域里放置 6 个按钮, 组成 3x2 网格, 部分按钮在视口外, 需要滚动才能看到
	for i := 1; i <= 6; i++ {
		name := fmt.Sprintf("按钮 %d", i)
		btn := widget.NewButton(int32(30+(i-1)%3*220), int32(30+(i-1)/3*180), 200, 120, name, eleContent.Handle)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			w.MessageBox("提示", fmt.Sprintf("你点击了内容区里的 %s", name), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
			return 0
		})
	}

	// 用形状文本实时显示滚动位置
	stPos := widget.NewShapeText(20, 380, 400, 24, "滚动位置  水平:0  垂直:0", w.Handle)
	stPos.SetTextAlign(xcc.TextAlignFlag_Left)
	updatePos := func() {
		stPos.SetText(fmt.Sprintf("滚动位置  水平:%d  垂直:%d", sv.GetViewPosH(), sv.GetViewPosV()))
		// 形状对象的 Redraw 无参数
		stPos.Redraw()
	}

	// 注册水平滚动事件
	sv.AddEvent_ScrollView_Scroll_H(func(hEle int, pos int32, pbHandled *bool) int {
		updatePos()
		return 0
	})
	// 注册垂直滚动事件
	sv.AddEvent_ScrollView_Scroll_V(func(hEle int, pos int32, pbHandled *bool) int {
		updatePos()
		return 0
	})

	// 用代码控制滚动到四个角
	btnTop := widget.NewButton(20, 414, 90, 30, "滚动到顶部", w.Handle)
	btnTop.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		sv.ScrollTop()
		updatePos()
		return 0
	})
	btnBottom := widget.NewButton(120, 414, 90, 30, "滚动到底部", w.Handle)
	btnBottom.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		sv.ScrollBottom()
		updatePos()
		return 0
	})
	btnLeft := widget.NewButton(220, 414, 90, 30, "滚动到左侧", w.Handle)
	btnLeft.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		sv.ScrollLeft()
		updatePos()
		return 0
	})
	btnRight := widget.NewButton(320, 414, 90, 30, "滚动到右侧", w.Handle)
	btnRight.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		sv.ScrollRight()
		updatePos()
		return 0
	})

	// 演示动态改变内容尺寸, 滚动范围随之改变
	bBig := false
	btnSize := widget.NewButton(440, 414, 100, 30, "扩大内容区", w.Handle)
	btnSize.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		bBig = !bBig
		if bBig {
			sv.SetTotalSize(1000, 700)
			btnSize.SetText("恢复内容区")
		} else {
			sv.SetTotalSize(700, 520)
			btnSize.SetText("扩大内容区")
		}
		// 修改显示后重绘
		btnSize.Redraw(false)
		sv.Redraw(false)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
