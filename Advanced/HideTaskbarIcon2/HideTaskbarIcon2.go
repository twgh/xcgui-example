// 任务栏不显示图标, 使用xml创建窗口.
package main

import (
	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed HideTaskbarIcon2.xml
var xmlStr string

func main() {
	// 初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 从xml加载窗口
	w := window.NewByLayoutStringW(xmlStr, 0, 0)

	// 获取窗口当前扩展样式
	exStyle := wapi.GetWindowLongPtrW(w.GetHWND(), wapi.GWL_EXSTYLE)
	// 添加 WS_EX_TOOLWINDOW
	newExStyle := exStyle | int(xcc.WS_EX_TOOLWINDOW)
	// 设置新的窗口扩展样式
	wapi.SetWindowLongPtrW(w.GetHWND(), wapi.GWL_EXSTYLE, newExStyle)

	w.Show(true)
	a.Run()
	a.Exit()
}
