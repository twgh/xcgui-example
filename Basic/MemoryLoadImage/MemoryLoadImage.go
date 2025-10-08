// 内存加载图片, 窗口设置背景图片
package main

import (
	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed 1.png
var imgBytes []byte

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 415, 296, "", 0, xcc.Window_Style_Default)
	w.SetBorderSize(2, 28, 2, 2)

	// 加载图片从内存, 自适应
	bkImg := imagex.NewByMemAdaptive(imgBytes, 0, 0, 0, 0)
	// 窗口_添加背景图片
	w.AddBkImage(xcc.Window_State_Flag_Body_Leave, bkImg.Handle)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
