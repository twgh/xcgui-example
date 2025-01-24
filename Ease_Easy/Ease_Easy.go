// 窗口简单缓动
package main

import (
	"runtime"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/ease"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	a *app.App
	w *window.Window
)

func main() {
	// 这是必要的, 这将保证main函数中对UI库命令的调用是在一个系统线程中执行的。
	// 如果不在一个系统线程中执行, 那程序有很大概率卡死.
	// 因为下面用了time.Sleep(), go的运行时可能会进行调度, 就跳到其他线程了, 所以必须用这个.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	a = app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	a.SetPaintFrequency(10)

	w = window.New(0, 0, 400, 300, "窗口简单缓动", 0, xcc.Window_Style_Default)
	w.Show(true)

	rc := w.GetRectDPI()
	for t := 0; t <= 30; t++ {
		v := ease.Bounce(float32(t)/30.0, xcc.Ease_Type_Out)
		y := int32(v * float32(rc.Top))
		w.SetPosition(rc.Left, y)
		w.Redraw(true)
		time.Sleep(time.Millisecond * 10)
	}

	a.Run()
	a.Exit()
}
