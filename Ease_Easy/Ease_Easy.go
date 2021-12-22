// 窗口简单缓动
package main

import (
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/ease"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	w := window.NewWindow(0, 0, 400, 300, "", 0, xcc.Window_Style_Default)
	// 显示窗口
	w.ShowWindow(xcc.SW_SHOW)

	a.CallUiThread(func(data int) int {
		// 获取窗口坐标
		var rect xc.RECT
		w.GetRect(&rect)
		// 缓动
		for i := 0; i <= 30; i++ {
			v := ease.Bounce(float32(i)/30.0, xcc.Ease_Type_Out)
			y := int(v * float32(rect.Top))

			w.SetPosition(int(rect.Left), y)
			w.Redraw(true)
			time.Sleep(time.Millisecond * 10)
		}
		return 0
	}, 0)

	a.Run()
	a.Exit()
}
