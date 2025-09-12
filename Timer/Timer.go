// 定时器
package main

import (
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 500, 130, "定时器", 0, xcc.Window_Style_Default)

	text := widget.NewShapeText(50, 50, 400, 50, time.Now().Format("2006-01-02 15:04:05"), w.Handle)
	text.SetFont(font.New(30).Handle)

	// 定时器id是自己定的
	w.SetTimer(1, 1000)
	w.AddEvent_Timer(func(hWindow int, nIDEvent uint, pbHandled *bool) int {
		switch nIDEvent {
		case 1:
			text.SetText(time.Now().Format("2006-01-02 15:04:05"))
			text.Redraw()
		}
		return 0
	})
	// 窗口_关闭定时器
	// w.KillTimer(1)

	w.Show(true)
	a.Run()
	a.Exit()
}
