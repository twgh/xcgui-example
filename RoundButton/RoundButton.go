package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New("")
	// 2.创建窗口
	win := window.NewWindow(0, 0, 466, 300, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

	// 设置窗口边框大小
	win.SetBorderSize(1, 30, 1, 1)
	// 设置窗口透明类型
	win.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口透明度, 不透明
	win.SetTransparentAlpha(255)
	// 设置窗口阴影
	win.SetShadowInfo(10, 255, 10, false, 0)
	// 窗口置顶
	win.SetTop()
	// 窗口居中
	win.Center()
	// 创建标签_窗口标题
	lbl_Title := widget.NewShapeText(15, 15, 56, 20, "Title", win.Handle)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// 创建结束按钮
	btn_Close := widget.NewButton(426, 10, 30, 30, "X", win.Handle)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

	// 创建按钮
	btn := widget.NewButton(30, 50, 70, 30, "Button", win.Handle)
	// 设置按钮字体颜色, 白色
	btn.SetTextColor(xc.RGB(255, 255, 255), 255)
	// 设置按钮圆角
	setBtnRound(btn, 14)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// 设置按钮圆角
func setBtnRound(btn *widget.Button, round int) {
	// 启用按钮背景透明
	btn.EnableBkTransparent(true)
	// 注册按钮绘制事件
	btn.Event_PAINT1(func(hEle int, hDraw int, pbHandled *bool) int {
		// 创建Draw对象
		draw := drawx.NewDrawByHandle(hDraw)
		// 启用平滑模式
		draw.EnableSmoothingMode(true)
		// 设置三种状态下的按钮背景色
		nState := xc.XBtn_GetStateEx(hEle)
		var bgcolor int
		switch nState {
		case xcc.Button_State_Leave:
			bgcolor = xc.RGB(1, 162, 232)
		case xcc.Button_State_Stay:
			bgcolor = xc.RGB(1, 182, 252)
		case xcc.Button_State_Down:
			bgcolor = xc.RGB(1, 122, 192)

		}
		// 设置画刷颜色
		draw.SetBrushColor(bgcolor, 255)

		// 绘制填充圆角矩形
		rc := xc.RECT{}
		rc.Right = int32(xc.XEle_GetWidth(hEle))
		rc.Bottom = int32(xc.XEle_GetHeight(hEle))
		draw.FillRoundRect(&rc, round, round)
		return 0
	})
}
