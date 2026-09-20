// 圆形形状对象
package main

// 演示 ShapeEllipse: 设置填充色、边框色、圆环(填充透明+边框)等外观效果.
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
	w := window.New(0, 0, 560, 430, "ShapeEllipse", 0, xcc.Window_Style_Default)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// ---------------- 第一行: 填充与边框 ----------------
	// 1. 只有填充色的实心圆, 默认无边框
	e1 := widget.NewShapeEllipse(20, 45, 100, 100, w.Handle)
	e1.SetFillColor(xc.RGBA(66, 133, 244, 255))

	// 2. 填充色 + 边框色
	e2 := widget.NewShapeEllipse(150, 45, 100, 100, w.Handle)
	e2.SetFillColor(xc.HexRGB2RGBA("#34A853", 255))
	e2.EnableBorder(true)
	e2.SetBorderColor(xc.HexRGB2RGBA("#1E7A34", 255))

	// 3. 只有边框, 填充透明, 形成细边框圆环
	e3 := widget.NewShapeEllipse(280, 45, 100, 100, w.Handle)
	e3.EnableFill(false)
	e3.EnableBorder(true)
	e3.SetBorderColor(xc.HexRGB2RGBA("#EA4335", 255))

	// 4. 椭圆: 宽高不相等就是椭圆
	e4 := widget.NewShapeEllipse(410, 45, 130, 100, w.Handle)
	e4.SetFillColor(xc.RGBA(251, 188, 5, 160))
	e4.EnableBorder(true)
	e4.SetBorderColor(xc.RGBA(180, 130, 0, 255))

	// ---------------- 第二行: 粗边框圆环 ----------------
	// 5. 圆环效果: 形状圆没有边框宽度 API, 用多个同心边框圆叠加出粗边框,
	// 这个效果并不好, 边框中间有空隙, 建议用背景对象 / Draw API 来实现.
	for i := 0; i < 5; i++ {
		ring := widget.NewShapeEllipse(int32(20+i), int32(175+i), int32(100-i*2), int32(100-i*2), w.Handle)
		ring.EnableFill(false)
		ring.EnableBorder(true)
		ring.SetBorderColor(xc.HexRGB2RGBA("#9AA0A6", 255))
	}

	// 6. 渐变同心圆: 多个填充透明度递减的圆叠加
	for i := 0; i < 5; i++ {
		c := widget.NewShapeEllipse(int32(160+i*12), int32(187+i*12), int32(76-i*24), int32(76-i*24), w.Handle)
		c.SetFillColor(xc.RGBA(108, 99, 255, 255))
		c.SetAlpha(byte(50 + i*40))
	}

	// 7. 外圈边框圆 + 内部实心圆, 组成指示灯效果
	outside := widget.NewShapeEllipse(420, 175, 100, 100, w.Handle)
	outside.EnableFill(false)
	outside.EnableBorder(true)
	outside.SetBorderColor(xc.HexRGB2RGBA("#34A853", 255))
	inside := widget.NewShapeEllipse(440, 195, 60, 60, w.Handle)
	inside.SetFillColor(xc.HexRGB2RGBA("#34A853", 255))

	// ---------------- 第三行: 圆环与圆点组合 ----------------
	// 8. 一排彩色小圆点, 常用来做状态指示点
	colors := []string{"#EA4335", "#FBBC05", "#34A853", "#4285F4", "#9AA0A6"}
	for i, hex := range colors {
		dot := widget.NewShapeEllipse(int32(20+i*50), 320, 26, 26, w.Handle)
		dot.SetFillColor(xc.HexRGB2RGBA(hex, 255))
	}

	w.Show(true)
	a.Run()
	a.Exit()
}
