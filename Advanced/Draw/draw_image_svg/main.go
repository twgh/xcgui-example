// Draw 图片与SVG绘制 — Image/ImageEx/Adaptive/Tile/MaskEllipse/Svg.
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 650, 320, "Draw 图片与SVG", 0, xcc.Window_Style_Default)
	w.MaxWindow()

	imgData := makeTestImage()
	hImg := imagex.NewByDataRGBA(imgData, 64, 64).Handle
	svgIcon := svg.NewByString(svgHeart)
	defer svgIcon.Destroy()

	border := w.GetBorderSizeEx()
	offY := border.Top

	w.AddEvent_Paint(func(hWindow int, hDraw int, pbHandled *bool) int {
		*pbHandled = true
		xc.XWnd_DrawWindow(hWindow, hDraw)
		draw := drawx.NewByHandle(hDraw)
		draw.EnableSmoothingMode(true)

		y := offY + 20

		// 1. Image 原始大小
		draw.Image(hImg, 20, y)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.TextOutEx(20, y+70, "Image(原始64x64)")
		y += 130

		// 2. ImageEx 指定宽高
		draw.XDraw_ImageEx(hImg, 20, y, 48, 48)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.TextOutEx(20, y+55, "ImageEx(指定宽高48x48)")
		y += 120

		// 3. ImageAdaptive 九宫格
		xc.XImage_SetDrawTypeAdaptive(hImg, 10, 10, 10, 10)
		draw.ImageAdaptive(hImg,
			&xc.RECT{Left: 20, Top: y, Right: 220, Bottom: y + 60}, false)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.TextOutEx(20, y+65, "ImageAdaptive(九宫格)")
		xc.XImage_SetDrawType(hImg, xcc.Image_Draw_Type_Stretch)

		// 4. ImageMaskEllipse 圆形遮罩
		sx := int32(240)
		draw.ImageMaskEllipse(hImg,
			&xc.RECT{Left: sx, Top: offY + 20, Right: sx + 150, Bottom: offY + 20 + 120},
			&xc.RECT{Left: sx, Top: offY + 20, Right: sx + 120, Bottom: offY + 20 + 120},
		)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.TextOutEx(sx, offY+150, "ImageMaskEllipse 圆形遮罩")

		// todo: 报错: 有 BUG ?
		// 5. ImageTile 平铺
		/* tx := int32(430)
		draw.ImageTile(hImg, 0,
			&xc.RECT{Left: tx, Top: offY + 20, Right: tx + 210, Bottom: offY + 20 + 160}, 0)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.TextOutEx(tx, offY+185, "ImageTile(平铺)") */

		// ====== SVG ======
		sx = 20
		y = offY + 380
		draw.DrawSvg(svgIcon.Handle, 400, offY)
		draw.SetBrushColor(xc.RGBA(100, 100, 100, 255))
		draw.TextOutEx(500, offY+60, "Svg原始:")
		sx += 60
		draw.DrawSvgEx(svgIcon.Handle, sx, y, 32, 32)
		draw.TextOutEx(sx, y+40, "32x32")
		sx += 50
		draw.DrawSvgEx(svgIcon.Handle, sx, y, 48, 48)
		draw.TextOutEx(sx, y+55, "48x48")
		sx += 70
		draw.DrawSvgEx(svgIcon.Handle, sx, y, 64, 64)
		draw.TextOutEx(sx, y+70, "64x64")

		return 0
	})

	w.Show()
	a.Run()
	a.Exit()
}

const svgHeart = `<svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg">
<path d="M512 928L164.5 580.5C82.5 498.5 48 396.5 48 298.5C48 154.5 162.5 40 306.5 40C388.5 40 466.5 76 512 136C557.5 76 635.5 40 717.5 40C861.5 40 976 154.5 976 298.5C976 396.5 941.5 498.5 859.5 580.5L512 928Z" fill="#E53935"/>
</svg>`

func makeTestImage() []byte {
	w, h := 64, 64
	data := make([]byte, w*h*4)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i := (y*w + x) * 4
			if (x/16+y/16)%2 == 0 {
				data[i], data[i+1], data[i+2], data[i+3] = 66, 133, 244, 255
			} else {
				data[i], data[i+1], data[i+2], data[i+3] = 234, 67, 53, 255
			}
		}
	}
	return data
}
