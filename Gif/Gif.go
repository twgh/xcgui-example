// Gif.
package main

import (
	_ "embed"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed res/bg1.gif
	bg1 []byte
)

var (
	a *app.App
	w *window.Window
)

func main() {
	a = app.New(false)
	w = window.New(0, 0, 600, 400, "xcgui window", 0, xcc.Window_Style_Default)

	shapeGif := widget.NewShapeGif(0, 30, 600, 400, w.Handle)
	shapeGif.SetImage(imagex.NewByMemAdaptive(bg1, 0, 0, 0, 0).Handle)

	w.Show(true)
	a.Run()
	a.Exit()
}
