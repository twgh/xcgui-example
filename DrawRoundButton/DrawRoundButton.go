// 绘制圆角按钮
// 自己绘制要记一些api, 还是建议使用设计器来实现
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
	a := app.New(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	// 创建一个按钮
	btn := widget.NewButton(30, 50, 70, 30, "Button", w.Handle)
	// 设置按钮字体颜色, 白色
	btn.SetTextColor(xc.ABGR(255, 255, 255, 254))
	// 设置按钮圆角
	setBtnRound(btn, 14)

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
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
		draw := drawx.NewByHandle(hDraw)
		// 启用平滑模式
		draw.EnableSmoothingMode(true)

		// 设置三种状态下的按钮背景色
		nState := xc.XBtn_GetStateEx(hEle)
		var bgcolor int
		switch nState {
		case xcc.Button_State_Leave:
			bgcolor = xc.ABGR(1, 162, 232, 254)
		case xcc.Button_State_Stay:
			bgcolor = xc.ABGR(1, 182, 252, 254)
		case xcc.Button_State_Down:
			bgcolor = xc.ABGR(1, 122, 192, 254)
		case xcc.Button_State_Disable:
			bgcolor = xc.ABGR(211, 215, 212, 255)
		}
		// 设置画刷颜色
		draw.SetBrushColor(bgcolor)

		// 绘制填充圆角矩形
		rc := xc.RECT{}
		rc.Right = int32(xc.XEle_GetWidth(hEle))
		rc.Bottom = int32(xc.XEle_GetHeight(hEle))
		draw.FillRoundRect(&rc, round, round)
		return 0
	})
}
