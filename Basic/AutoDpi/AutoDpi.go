// 启用自动DPI的三种方法, 解决高分辨率屏幕界面模糊问题.
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)

	// 三种方法:
	// 1.调用 a.EnableDPI(true)
	// 2.使用程序清单, 看这个: 程序清单方式启用DPI.rar
	// 3.自行调用 Windows 相关api
	a.EnableDPI(true)

	// 使用上面的几种方法之一, 然后调用这个函数启用自动DPI
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "EnableAutoDPI", 0, xcc.Window_Style_Default)

	// 创建按钮
	btn := widget.NewButton(165, 135, 100, 30, "Button", w.Handle)
	// 注册按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		a.MessageBox("提示", btn.GetText(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
