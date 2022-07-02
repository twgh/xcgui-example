// 内存加载图片
package main

import (
	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//go:embed 1.png
var img1 []byte

func main() {
	a := app.New(true)
	w := window.New(0, 0, 415, 296, "", 0, xcc.Window_Style_Default)

	// 加载图片从内存
	hImg := xc.XImage_LoadMemory(img1)

	// 创建形状对象-图片
	shapePic := widget.NewShapePicture(8, 30, 400, 260, w.Handle)
	shapePic.SetImage(hImg)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
