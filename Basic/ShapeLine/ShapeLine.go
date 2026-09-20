// 直线形状对象
package main

// 演示 ShapeLine: 设置线条颜色, 水平/垂直/斜线, 以及用分隔线组织界面元素.
// 形状对象没有事件, 重点在于演示各种外观设置 API.

import (
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
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 560, 430, "ShapeLine", 0, xcc.Window_Style_Default)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// ---------------- 水平分隔线 ----------------
	// 1. 浅灰色水平线, 最常用的界面分隔线, 两点 y 值相等即为水平线
	widget.NewShapeLine(20, 55, 540, 55, w.Handle).
		SetColor(xc.HexRGB2RGBA("#DADCE0", 255))

	// 2. 分隔线上下各放一个元素, 展示分隔效果
	widget.NewShapeText(20, 70, 200, 24, "用户信息", w.Handle)
	widget.NewShapeText(20, 106, 200, 24, "这是分隔线下方的内容", w.Handle)
	widget.NewShapeLine(20, 98, 540, 98, w.Handle).
		SetColor(xc.HexRGB2RGBA("#9AA0A6", 255))

	// 3. 不同颜色的水平线
	widget.NewShapeLine(20, 145, 540, 145, w.Handle).
		SetColor(xc.RGBA(66, 133, 244, 255))
	widget.NewShapeLine(20, 155, 540, 155, w.Handle).
		SetColor(xc.HexRGB2RGBA("#EA4335", 255))
	widget.NewShapeLine(20, 165, 540, 165, w.Handle).
		SetColor(xc.RGBA(251, 188, 5, 255))

	// 4. 粗线: 线宽没有单独的 API, 用2条相邻的水平线叠加, 这个方法并不好,
	// 如果是水平或垂直线, 建议用 ShapeRect / Element / 背景对象来实现,
	// 如果是斜线, 建议用 Draw API 来实现.
	for i := 0; i < 2; i++ {
		widget.NewShapeLine(20, int32(185+i), 540, int32(185+i), w.Handle).
			SetColor(xc.HexRGB2RGBA("#5F6368", 255))
	}

	// 5. 虚线观感: 用多段短线间隔排列模拟, 常用于表单分组
	for i := 0; i < 6; i++ {
		widget.NewShapeLine(int32(20+i*90), 210, int32(80+i*90), 210, w.Handle).
			SetColor(xc.HexRGB2RGBA("#BDC1C6", 255))
	}

	// ---------------- 垂直线与斜线 ----------------
	// 6. 垂直线: 两点 x 值相等即为垂直线, 用来分隔左右两栏
	widget.NewShapeLine(300, 240, 300, 390, w.Handle).
		SetColor(xc.HexRGB2RGBA("#DADCE0", 255))

	// 左右两栏各放一个按钮, 展示垂直分隔线的分组效果
	widget.NewButton(120, 290, 140, 34, "左侧操作", w.Handle)
	widget.NewButton(360, 290, 140, 34, "右侧操作", w.Handle)

	// 7. 斜线: 两点坐标 x、y 都不相同即为斜线
	widget.NewShapeLine(330, 390, 540, 240, w.Handle).
		SetColor(xc.RGBA(108, 99, 255, 255))
	widget.NewShapeLine(330, 240, 540, 390, w.Handle).
		SetColor(xc.RGBA(52, 168, 83, 128))

	// 8. 用 SetPosition 动态改变直线位置(重新设置两个端点坐标)
	line := widget.NewShapeLine(20, 240, 260, 240, w.Handle)
	line.SetColor(xc.HexRGB2RGBA("#E8710A", 255))
	line.SetPosition(20, 300, 260, 300)

	w.Show(true)
	a.Run()
	a.Exit()
}
