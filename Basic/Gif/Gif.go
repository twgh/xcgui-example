// Gif.
package main

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed bg1.gif
var bg1 []byte

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 600, 400, "Gif", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// 创建一个布局元素
	lay := widget.NewLayoutEle(0, 0, 0, 0, w.Handle)
	// 设置布局元素填充父
	lay.LayoutItem_SetWidth(xcc.Layout_Size_Fill, -1)
	lay.LayoutItem_SetHeight(xcc.Layout_Size_Fill, -1)
	// 布局盒子_置子项间距.
	lay.SetSpace(2)

	// 按钮_停止/播放. 在布局里x/y坐标是无效的, 所以填什么都无所谓.
	btnStop := widget.NewButton(0, 0, 70, 30, "停止", lay.Handle)
	// 按钮_暂停/继续
	btnPause := widget.NewButton(0, 0, 70, 30, "暂停", lay.Handle)
	// 按钮_销毁
	btnDestroy := widget.NewButton(0, 0, 70, 30, "销毁", lay.Handle)

	// 创建 Gif 播放器.
	gp, err := widget.NewGifPlayer(bytes.NewReader(bg1))
	if err != nil {
		panic(err)
	}
	// 播放 Gif.
	gph := gp.Play(lay.Handle, true, 0, func(h *widget.GifPlayerHandler, frame int) {
		maxFrame := h.GetMaxFrame()
		if frame == maxFrame {
			fmt.Println("完成一次播放")
		}
	})

	// 按钮_停止/播放
	btnStop.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		btnStop.Enable(false).Redraw(false)
		defer btnStop.Enable(true).Redraw(false)

		if gph.IsStopped() {
			gph = gp.Play(lay.Handle, true, 0)
			btnStop.SetText("停止")
		} else {
			gph.Stop()
			btnStop.SetText("播放")
		}
		btnPause.SetText("暂停").Redraw(false)
		return 0
	})

	// 按钮_暂停/继续
	btnPause.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if gph.IsStopped() {
			return 0
		}

		btnPause.Enable(false).Redraw(false)
		defer btnPause.Enable(true).Redraw(false)

		if gph.IsPaused() {
			gph.Resume()
			btnPause.SetText("暂停")
		} else {
			gph.Pause()
			btnPause.SetText("继续")
		}
		return 0
	})

	// 按钮_销毁
	btnDestroy.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		btnPause.Enable(false).Redraw(false)
		defer btnPause.Enable(true).Redraw(false)

		gph.Destroy()

		btnStop.SetText("播放").Redraw(false)
		btnPause.SetText("暂停").Redraw(false)
		return 0
	})

	// 窗口关闭事件
	w.AddEvent_Close(func(hWindow int, pbHandled *bool) int {
		gph.Destroy()
		// 释放 Gif 帧图片
		gp.ReleaseImages()
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
