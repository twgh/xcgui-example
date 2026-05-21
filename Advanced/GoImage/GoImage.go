// 从go image创建炫彩图片
package main

import (
	"image"
	"image/color"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 500, 400, "从go image创建炫彩图片", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)

	// 创建2张 RGBA 图片
	imgs := createGoImage()

	// 从go image创建炫彩图片
	img := imagex.NewByGoImage(imgs[0])

	// 创建形状图片
	sp := widget.NewShapePicture(30, 100, 256, 256, w.Handle)
	sp.SetImage(img.Handle)

	// 按钮_修改图片
	widget.NewButton(30, 40, 100, 30, "修改图片", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if w.GetProperty("图片") == "2" {
			w.SetProperty("图片", "1")
			img.ModifyDataGoImage(imgs[0]) // 在炫彩图片句柄不变的情况下, 修改其对应的图片数据
		} else {
			w.SetProperty("图片", "2")
			img.ModifyDataGoImage(imgs[1])
		}

		sp.Redraw()
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}

// 创建2张 RGBA 图片
func createGoImage() []*image.RGBA {
	const w, h = 256, 256
	img1 := image.NewRGBA(image.Rect(0, 0, w, h))
	img2 := image.NewRGBA(image.Rect(0, 0, w, h))

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if x < w/2 {
				img1.Set(x, y, red)  // 左半边：红色
				img2.Set(x, y, blue) // 右半边：蓝色
			} else {
				img1.Set(x, y, blue) // 右半边：蓝色
				img2.Set(x, y, red)  // 左半边：红色
			}
		}
	}

	return []*image.RGBA{img1, img2}
}
