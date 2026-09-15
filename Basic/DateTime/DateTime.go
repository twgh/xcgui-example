// 日期时间框
package main

import (
	"fmt"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 460, 240, "日期时间框", 0, xcc.Window_Style_Default)

	dt0 := widget.NewDateTime(20, 50, 130, 26, w.Handle)
	// 0为日期元素
	dt0.SetStyle(0)

	dt1 := widget.NewDateTime(170, 50, 130, 26, w.Handle)
	// 1为时间元素
	dt1.SetStyle(1)

	// 用形状文本显示当前选择的结果
	stResult := widget.NewShapeText(20, 100, 420, 24, "当前选择:", w.Handle)

	// 读取两个日期时间框的当前值并刷新到界面.
	// GetDateEx 返回 (year, month, day int32), GetTimeEx 返回 (hour, minute, second int32).
	updateText := func() {
		y, m, d := dt0.GetDateEx()
		h, mi, s := dt1.GetTimeEx()
		stResult.SetText(fmt.Sprintf("当前选择日期时间: %d年%d月%d日  %02d:%02d:%02d", y, m, d, h, mi, s))
		// 修改元素显示内容后必须重绘, 否则界面不刷新
		stResult.Redraw()
	}

	// 日期时间元素内容改变事件, 用户在控件上修改日期或时间时触发.
	dt0.AddEvent_DateTime_Change(func(hEle int, pbHandled *bool) int {
		updateText()
		return 0
	})

	// 两个调节器按钮加上点击事件
	widget.NewButtonByHandle(dt1.GetButton(1)).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		updateText()
		return 0
	})
	widget.NewButtonByHandle(dt1.GetButton(2)).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		updateText()
		return 0
	})

	// 按键抬起事件, 当用户在控件上按下键盘按键并抬起时触发.
	dt1.AddEvent_KeyUp(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
		// 只处理数字键盘按键.
		if wParam < xcc.VK_0 && wParam > xcc.VK_9 {
			return 0
		}
		if wParam < xcc.VK_Numpad0 && wParam > xcc.VK_Numpad9 {
			return 0
		}
		updateText()
		return 0
	})

	// 按钮: 用 SetDate/SetTime 把两个控件设置为系统当前日期和时间.
	btnSet := widget.NewButton(20, 150, 130, 30, "设置为现在", w.Handle)
	btnSet.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		n := time.Now()
		// SetDate(年, 月, 日) 设置日期控件的值.
		dt0.SetDate(int32(n.Year()), int32(n.Month()), int32(n.Day()))
		// SetTime(时, 分, 秒) 设置时间控件的值.
		dt1.SetTime(int32(n.Hour()), int32(n.Minute()), int32(n.Second()))
		// SetDate 也会触发内容改变事件, 这里再刷新一次保证显示同步.
		updateText()
		return 0
	})

	// 初始显示一次当前值.
	updateText()

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
