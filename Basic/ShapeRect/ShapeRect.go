// 矩形形状对象
package main

// 演示 ShapeRect: 设置填充色、边框色、圆角矩形、颜色透明度、整体透明度等外观效果.
//
// 形状对象没有事件, 也不具备按钮那样的 leave/stay/down 状态颜色,
// 想要多态效果可以用按钮配合背景管理器实现, 这里重点展示外观设置 API.

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
	w := window.New(0, 0, 560, 440, "ShapeRect", 0, xcc.Window_Style_Default)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// ---------------- 第一行: 填充与边框 ----------------
	// 1. 只有填充色, 未启用边框
	r1 := widget.NewShapeRect(20, 45, 110, 80, w.Handle)
	r1.SetFillColor(xc.RGBA(66, 133, 244, 255))
	r1.EnableBorder(false)

	// 2. 填充色 + 边框色
	r2 := widget.NewShapeRect(150, 45, 110, 80, w.Handle)
	r2.SetFillColor(xc.HexRGB2RGBA("#34A853", 255))
	r2.EnableBorder(true)
	r2.SetBorderColor(xc.RGBA(30, 100, 60, 255))

	// 3. 只有边框, 填充透明, 形成空心矩形
	r3 := widget.NewShapeRect(280, 45, 110, 80, w.Handle)
	r3.EnableFill(false)
	r3.EnableBorder(true)
	r3.SetBorderColor(xc.HexRGB2RGBA("#EA4335", 255))

	// 4. 填充色带透明度(RGBA 第4个参数), 可以透出窗口背景
	r4 := widget.NewShapeRect(410, 45, 110, 80, w.Handle)
	r4.SetFillColor(xc.RGBA(251, 188, 5, 120))

	// ---------------- 第二行: 圆角矩形 ----------------
	// 5. 圆角矩形: 先启用圆角, 再设置圆角宽高
	r5 := widget.NewShapeRect(20, 155, 110, 80, w.Handle)
	r5.EnableRoundAngle(true)
	r5.SetRoundAngle(15, 15)
	r5.SetFillColor(xc.HexRGB2RGBA("#9AA0A6", 255))

	// 6. 大圆角 + 边框, 圆角值越大越圆
	r6 := widget.NewShapeRect(150, 155, 110, 80, w.Handle)
	r6.EnableRoundAngle(true)
	r6.SetRoundAngle(35, 35)
	r6.SetFillColor(xc.RGBA(108, 99, 255, 255))
	r6.EnableBorder(true)
	r6.SetBorderColor(xc.HexRGB2RGBA("#5F5AA2", 255))

	// 7. 圆角为高度一半时, 变成胶囊形状(填充透明+边框)
	r7 := widget.NewShapeRect(280, 155, 110, 80, w.Handle)
	r7.EnableRoundAngle(true)
	r7.SetRoundAngle(40, 40)
	r7.EnableFill(false)
	r7.EnableBorder(true)
	r7.SetBorderColor(xc.RGBA(0, 150, 200, 255))

	// ---------------- 第三行: 胶囊条与整体透明度 ----------------
	// 8. 通栏胶囊条, 常用来做卡片标题背景
	r8 := widget.NewShapeRect(20, 265, 500, 44, w.Handle)
	r8.EnableRoundAngle(true)
	r8.SetRoundAngle(22, 22)
	r8.SetFillColor(xc.HexRGB2RGBA("#F1F3F4", 255))
	r8.EnableBorder(true)
	r8.SetBorderColor(xc.HexRGB2RGBA("#DADCE0", 255))

	// 9. 整体透明度 SetAlpha, 同时作用填充和边框
	r9 := widget.NewShapeRect(20, 330, 240, 80, w.Handle)
	r9.SetFillColor(xc.HexRGB2RGBA("#FF6D00", 255))
	r9.EnableBorder(true)
	r9.SetBorderColor(xc.HexRGB2RGBA("#E65100", 255))
	r9.SetAlpha(128)

	// 10. 不透明度100的对比矩形
	r10 := widget.NewShapeRect(280, 330, 240, 80, w.Handle)
	r10.SetFillColor(xc.HexRGB2RGBA("#FF6D00", 255))
	r10.EnableBorder(true)
	r10.SetBorderColor(xc.HexRGB2RGBA("#E65100", 255))

	w.Show(true)
	a.Run()
	a.Exit()
}
