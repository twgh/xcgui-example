// 炫彩输出调试信息
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	// 炫彩_启用debug文件, xcgui_debug.txt
	a.EnableDebugFile(true)

	w := window.NewWindow(0, 0, 400, 300, "", 0, xcc.Window_Style_Default)

	btn := widget.NewButton(20, 35, 70, 30, "Click", w.Handle)
	btn.Event_BnClick(func(pbHandled *bool) int {
		a.DebugToFileInfo("hello")
		a.DebugToFileInfo("word")
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
