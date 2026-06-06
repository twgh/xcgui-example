// Draw 渐变填充 + 高级设置 — 渐变/裁剪/偏移/焦点框.
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

	w := window.New(0, 0, 650, 430, "Draw 渐变与高级设置", 0, xcc.Window_Style_Default)

	border := w.GetBorderSizeEx()
	offY := border.Top

	w.AddEvent_Paint(func(hWindow int, hDraw int, pbHandled *bool) int {
		*pbHandled = true                  // 拦截窗口原本的绘制
		xc.XWnd_DrawWindow(hWindow, hDraw) // 绘制窗口
		draw := drawx.NewByHandle(hDraw)
		draw.EnableSmoothingMode(true)

		// ====== 渐变 ======

		// 1. 水平渐变
		draw.GradientFill2(
			&xc.RECT{Left: 10, Top: offY + 10, Right: 200, Bottom: offY + 80},
			xc.RGBA(66, 133, 244, 255),
			xc.RGBA(52, 168, 83, 255),
			xcc.GRADIENT_FILL_RECT_H,
		)

		// 2. 垂直渐变
		draw.GradientFill2(
			&xc.RECT{Left: 220, Top: offY + 10, Right: 410, Bottom: offY + 80},
			xc.RGBA(234, 67, 53, 255),
			xc.RGBA(251, 188, 4, 255),
			xcc.GRADIENT_FILL_RECT_V,
		)

		// 3. 水平渐变 + 圆角边框叠层
		draw.GradientFill2(
			&xc.RECT{Left: 430, Top: offY + 10, Right: 640, Bottom: offY + 80},
			xc.RGBA(100, 50, 200, 255),
			xc.RGBA(50, 200, 200, 255),
			xcc.GRADIENT_FILL_RECT_H,
		)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 100))
		draw.DrawRoundRect(&xc.RECT{Left: 430, Top: offY + 10, Right: 640, Bottom: offY + 80}, 10, 10)

		// 4. 四角渐变
		draw.GradientFill4(
			&xc.RECT{Left: 10, Top: offY + 100, Right: 310, Bottom: offY + 250},
			xc.RGBA(255, 0, 0, 200),
			xc.RGBA(0, 255, 0, 200),
			xc.RGBA(0, 0, 255, 200),
			xc.RGBA(255, 255, 0, 200),
			xcc.GRADIENT_FILL_RECT_H,
		)

		// ====== 裁剪区域 ======

		// 5. SetClipRect
		draw.SetClipRect(&xc.RECT{Left: 330, Top: offY + 100, Right: 640, Bottom: offY + 250})
		draw.SetBrushColor(xc.RGBA(66, 133, 244, 100))
		draw.FillRoundRect(&xc.RECT{Left: 300, Top: offY + 80, Right: 670, Bottom: offY + 270}, 20, 20)
		draw.SetBrushColor(xc.RGBA(234, 67, 53, 255))
		draw.DrawLine(320, offY+120, 640, offY+230)
		draw.DrawLine(320, offY+230, 640, offY+120)
		draw.SetBrushColor(xc.RGBA(50, 50, 50, 255))
		draw.TextOutEx(340, offY+155, "SetClipRect 裁剪了边缘内容")
		draw.ClearClip()

		// ====== 偏移 ======

		// 6. SetOffset
		draw.SetOffset(10, -15)
		draw.SetBrushColor(xc.RGBA(52, 168, 83, 180))
		draw.FillRoundRect(&xc.RECT{Left: 10, Top: offY + 280, Right: 200, Bottom: offY + 360}, 8, 8)
		draw.SetBrushColor(xc.RGBA(255, 255, 255, 255))
		draw.TextOutEx(30, offY+305, "偏移绘制")
		draw.SetOffset(0, 0)

		// 7. FocusRect
		draw.SetBrushColor(xc.RGBA(66, 133, 244, 255))
		draw.FocusRect(&xc.RECT{Left: 430, Top: offY + 280, Right: 630, Bottom: offY + 360})

		// 8. 底部分隔
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.Dottedline(10, offY+380, 640, offY+380)

		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
