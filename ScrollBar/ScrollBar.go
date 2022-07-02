// 滚动条
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	w := window.New(0, 0, 430, 300, "ScrollBar", 0, xcc.Window_Style_Default)

	// 创建滚动条
	bar1 := widget.NewScrollBar(12, 33, 300, 20, w.Handle)
	bar2 := widget.NewScrollBar(330, 33, 20, 240, w.Handle)

	// 设置为垂直滚动条
	bar2.EnableHorizon(false)

	// 注册滚动条元素滚动事件
	bar1.Event_SBAR_SCROLL1(SBAR_SCROLL1)
	bar2.Event_SBAR_SCROLL1(SBAR_SCROLL1)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 滚动条元素滚动事件
func SBAR_SCROLL1(hEle, pos int, pbHandled *bool) int {
	fmt.Println(pos)
	// 为了鼠标滚轮滚动和点击两端按钮实时显示效果而刷新
	xc.XEle_Redraw(hEle, true)
	return 0
}
