// 任务栏不显示图标, 不是xml加载窗口的情况.
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 1. 使用 xcc.WS_EX_TOOLWINDOW 样式会把窗口创建为工具窗口, 就没有任务栏图标了
	w := window.NewEx(xcc.WS_EX_TOOLWINDOW, 0, "xcgui", 0, 0, 800, 600, "任务栏不显示图标", 0, xcc.Window_Style_Default)

	widget.NewButton(50, 50, 200, 30, "载入没填父句柄的窗口", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		// 如果没填窗口父句柄, 那它就有任务栏图标
		window.NewEx(0, 0, "xcgui2", 0, 0, 400, 300, "没填父句柄的窗口-有任务栏图标", 0, xcc.Window_Style_Default).Show(true)
		*pbHandled = true // 创建窗口后这个事件最好拦截下
		return 0
	})

	widget.NewButton(50, 150, 200, 30, "载入填了父句柄的窗口", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		// 2. 如果窗口父句柄不是桌面句柄也就是填了父句柄, 那它就没有任务栏图标
		window.NewEx(0, 0, "xcgui3", 0, 0, 400, 300, "填了父句柄的窗口-没任务栏图标", w.GetHWND(), xcc.Window_Style_Default).Show(true)
		*pbHandled = true // 创建窗口后这个事件最好拦截下
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
