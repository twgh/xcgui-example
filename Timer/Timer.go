// 定时器
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"time"
)

var (
	a    *app.App
	w    *window.Window
	text *widget.ShapeText
)

func main() {
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 500, 130, "定时器", 0, xcc.Window_Style_Default)

	text = widget.NewShapeText(50, 50, 400, 50, time.Now().Format("2006-01-02 15:04:05"), w.Handle)
	text.SetFont(font.New(30).Handle)

	// 定时器id是自己定的
	w.SetTimer(1, 1000)
	w.Event_TIMER(func(nIDEvent uint, pbHandled *bool) int {
		switch nIDEvent {
		case 1:
			text.SetText(time.Now().Format("2006-01-02 15:04:05"))
			text.Redraw()
		}
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
