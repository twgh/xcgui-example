// 日期时间框
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	w := window.NewWindow(0, 0, 400, 300, "", 0, xcc.Window_Style_Default)

	dt := widget.NewDateTime(20, 50, 120, 26, w.Handle)
	// 0为日期元素, 1为时间元素.
	dt.SetStyle(0)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
