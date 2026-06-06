// Draw 基本图形 — 矩形/圆角/椭圆/线/弧/多边形/虚线.
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 600, 450, "Draw 基本图形", 0, xcc.Window_Style_Default)

	// 获取标题栏高度, 用于偏移所有绘制坐标
	border := w.GetBorderSizeEx()
	offY := border.Top

	w.AddEvent_Paint(func(hWindow int, hDraw int, pbHandled *bool) int {
		*pbHandled = true
		xc.XWnd_DrawWindow(hWindow, hDraw)
		draw := drawx.NewByHandle(hDraw)
		draw.EnableSmoothingMode(true) // 绘制_启用平滑模式

		// ------ 1. 填充矩形 ------
		draw.SetBrushColor(xc.RGBA(66, 133, 244, 255))
		draw.FillRect(&xc.RECT{Left: 10, Top: offY + 10, Right: 100, Bottom: offY + 70})

		// ------ 2. 矩形边框 ------
		draw.SetBrushColor(xc.RGBA(234, 67, 53, 255))
		draw.SetLineWidth(2)
		draw.DrawRect(&xc.RECT{Left: 120, Top: offY + 10, Right: 210, Bottom: offY + 70})

		// ------ 3. 填充圆角矩形 ------
		draw.SetBrushColor(xc.RGBA(52, 168, 83, 255))
		draw.FillRoundRect(&xc.RECT{Left: 230, Top: offY + 10, Right: 320, Bottom: offY + 70}, 12, 12)

		// ------ 4. 圆角矩形边框 ------
		draw.SetBrushColor(xc.RGBA(251, 188, 4, 255))
		draw.DrawRoundRect(&xc.RECT{Left: 340, Top: offY + 10, Right: 430, Bottom: offY + 70}, 12, 12)

		// ------ 5. 填充椭圆 ------
		draw.SetBrushColor(xc.RGBA(66, 244, 155, 255))
		draw.FillEllipse(&xc.RECT{Left: 450, Top: offY + 10, Right: 530, Bottom: offY + 70})

		// ------ 6. 椭圆边框 ------
		draw.SetBrushColor(xc.RGBA(244, 66, 170, 255))
		draw.SetLineWidth(3)
		draw.DrawEllipse(&xc.RECT{Left: 540, Top: offY + 10, Right: 590, Bottom: offY + 60})

		// ------ 7. 线条 ------
		draw.SetBrushColor(xc.RGBA(0, 0, 0, 200))
		draw.SetLineWidth(2)
		draw.DrawLine(10, offY+90, 590, offY+90)
		draw.SetBrushColor(xc.RGBA(0, 0, 0, 100))
		draw.DrawLine(300, offY+90, 300, offY+300)

		// ------ 8. 圆弧 ------
		draw.SetBrushColor(xc.RGBA(155, 66, 244, 255))
		draw.SetLineWidth(3)
		draw.DrawArc(20, offY+110, 120, 100, 180, 180)

		// ------ 9. 多边形 ------
		draw.SetBrushColor(xc.RGBA(244, 150, 66, 255))
		draw.FillPolygon([]xc.POINT{
			{X: 170, Y: offY + 200}, {X: 230, Y: offY + 140},
			{X: 290, Y: offY + 200}, {X: 260, Y: offY + 260},
			{X: 200, Y: offY + 260},
		})

		// ------ 10. 虚线 ------
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.Dottedline(350, offY+140, 500, offY+140)

		// ------ 11. FillRectColor ------
		draw.FillRectColor(
			&xc.RECT{Left: 350, Top: offY + 160, Right: 500, Bottom: offY + 200},
			xc.RGBA(100, 200, 255, 180))

		return 0
	})

	w.Show()
	a.Run()
	a.Exit()
}
