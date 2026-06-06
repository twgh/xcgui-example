// Draw 文本绘制 — TextOut / DrawText / 对齐 / 字体 / 下划线.
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 600, 380, "Draw 文本绘制", 0, xcc.Window_Style_Default)

	fontNormal := font.NewEX("微软雅黑", 14, xcc.FontStyle_Regular)
	fontBold := font.NewEX("微软雅黑", 20, xcc.FontStyle_Bold)
	defer fontNormal.Destroy()
	defer fontBold.Destroy()

	border := w.GetBorderSizeEx()
	offY := border.Top

	w.AddEvent_Paint(func(hWindow int, hDraw int, pbHandled *bool) int {
		*pbHandled = true                  // 拦截窗口默认绘制
		xc.XWnd_DrawWindow(hWindow, hDraw) // 绘制窗口
		draw := drawx.NewByHandle(hDraw)

		// ------ 1. TextOutEx 简单文本 ------
		draw.SetBrushColor(xc.RGBA(50, 50, 50, 255))
		draw.TextOutEx(20, offY+20, "TextOutEx: 简单文本输出 (默认字体)")

		// ------ 2. 指定字体 ------
		draw.SetFont(fontNormal.Handle)
		draw.SetBrushColor(xc.RGBA(66, 133, 244, 255))
		draw.TextOutEx(20, offY+50, "TextOutEx: 微软雅黑 14px")

		// ------ 3. 粗体 ------
		draw.SetFont(fontBold.Handle)
		draw.SetBrushColor(xc.RGBA(234, 67, 53, 255))
		draw.TextOutEx(20, offY+80, "TextOutEx: 微软雅黑 20px 粗体")

		// ------ 4. DrawText 居中 ------
		draw.SetFont(fontNormal.Handle)
		draw.SetBrushColor(xc.RGBA(52, 168, 83, 255))
		draw.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)
		draw.DrawText("DrawText 居中对齐,\n支持多行文本自动换行",
			&xc.RECT{Left: 20, Top: offY + 120, Right: 280, Bottom: offY + 180})

		// ------ 5. 左对齐 + 边框参考 ------
		draw.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Top)
		draw.SetBrushColor(xc.RGBA(120, 80, 40, 255))
		draw.DrawRect(&xc.RECT{Left: 300, Top: offY + 120, Right: 580, Bottom: offY + 180})
		draw.SetBrushColor(xc.RGBA(50, 50, 50, 255))
		draw.DrawText("左对齐+顶对齐: 这段文字在\n矩形框内左上角对齐",
			&xc.RECT{Left: 305, Top: offY + 125, Right: 575, Bottom: offY + 175})

		// ------ 6. 下划线文本 ------
		draw.SetFont(fontBold.Handle)
		draw.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Top)
		draw.SetBrushColor(xc.RGBA(180, 50, 120, 255))
		draw.DrawTextUnderline("带下划线的粗体 — DrawTextUnderline",
			&xc.RECT{Left: 20, Top: offY + 210, Right: 580, Bottom: offY + 240},
			xc.RGBA(180, 50, 120, 255))

		// ------ 7. 十六进制颜色 ------
		draw.SetFont(fontNormal.Handle)
		draw.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Top)
		draw.SetBrushColor(xc.HexRGB2RGBA("#1E88E5", 255))
		draw.TextOutEx(20, offY+260, "HexRGB2RGBA(\"#1E88E5\", 255) 设置颜色")

		// ------ 8. 底部分隔虚线 ------
		draw.SetBrushColor(xc.RGBA(150, 150, 150, 100))
		draw.Dottedline(20, offY+300, 580, offY+300)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.SetFont(fontNormal.Handle)
		draw.TextOutEx(20, offY+310, "HexRGB2RGBA(\"#333\", 255) → ")
		draw.SetBrushColor(xc.HexRGB2RGBA("#333333", 255))
		draw.TextOutEx(310, offY+310, "#333333")

		return 0
	})

	w.Show()
	a.Run()
	a.Exit()
}
