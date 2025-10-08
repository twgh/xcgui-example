// 窗口简单缓动
package main

import (
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/ease"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 设置UI的最小重绘频率.
	a.SetPaintFrequency(10)

	// 创建窗口
	w := window.New(0, 0, 400, 300, "窗口简单缓动", 0, xcc.Window_Style_Default)
	// 窗口置顶
	w.SetTop(true)
	// 显示窗口
	w.Show(true)

	// 窗口缓动, 自上而下
	rc := w.GetRectDPI()
	for t := 0; t <= 30; t++ {
		v := ease.Bounce(float32(t)/30.0, xcc.Ease_Type_Out)
		y := int32(v * float32(rc.Top))
		w.SetPosition(rc.Left, y).Redraw(true)
		time.Sleep(time.Millisecond * 10)
	}

	a.Run()
	a.Exit()
}
