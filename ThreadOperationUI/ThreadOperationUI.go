// 炫彩_调用界面线程, 在主线程操作UI
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"strconv"
	"time"
)

var (
	a         *app.App
	w         *window.Window
	btn       *widget.Button
	radioBtn1 *widget.Button
	radioBtn2 *widget.Button

	t = 1 // 方式类型
)

func main() {
	a = app.New(true)
	w = window.New(0, 0, 550, 300, "ThreadOperationUI", 0, xcc.Window_Style_Default)

	// 创建按钮
	btn = widget.NewButton(20, 35, 100, 30, "click", w.Handle)
	btn.Event_BnClick(onBnClick)

	// 单选按钮
	radioBtn1 = widget.NewButton(20, 70, 70, 30, "方式1", w.Handle)
	radioBtn2 = widget.NewButton(100, 70, 70, 30, "方式2", w.Handle)
	radioBtn1.SetTypeEx(xcc.Button_Type_Radio)
	radioBtn2.SetTypeEx(xcc.Button_Type_Radio)
	radioBtn1.SetCheck(true) // 默认选中radioBtn1
	radioBtn1.Event_BUTTON_CHECK1(onBnCheck)
	radioBtn2.Event_BUTTON_CHECK1(onBnCheck)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func onBnClick(pbHandled *bool) int {
	// 禁用按钮
	btn.Enable(false)
	btn.Redraw(true)

	// 另起线程是为了不卡界面
	switch t {
	case 1:
		go updateBtn1() // 第一种方式: xc.XC_CallUiThreadEx
	case 2:
		go updateBtn2() // 第二种方式: xc.XC_CallUiThreader
	}
	return 0
}

// 第一种方式
func updateBtn1() {
	fmt.Println("使用方式1: xc.XC_CallUiThreadEx")
	for i := 0; i < 2010; i++ {
		// 如果直接在非主线程内操作UI, 次数多了程序必将崩溃, 而且你不会知道它在什么时候崩溃.
		// 使用 xc.XC_CallUiThreadEx 这样是在界面线程进行UI操作, 就不会崩溃了.
		xc.XC_CallUiThreadEx(func(data int) int {
			btn.SetText(strconv.Itoa(data))
			btn.SetWidth(i / 5)
			w.Redraw(false)
			return 0
		}, i)
		time.Sleep(time.Millisecond * 1)
	}

	// 解禁按钮
	xc.XC_CallUiThreadEx(func(data int) int {
		btn.Enable(true)
		btn.Redraw(true)
		return 0
	}, 0)
}

func updateBtn2() {
	fmt.Println("使用方式2: xc.XC_CallUiThreader")
	u := updateButton{
		HEle:         btn.Handle,
		RedrawWindow: false,
	}

	for i := 0; i < 2010; i++ {
		// 如果直接在非主线程内操作UI, 次数多了程序必将崩溃, 而且你不会知道它在什么时候崩溃.
		// 使用 xc.XC_CallUiThreader 这样是在界面线程进行UI操作, 就不会崩溃了.
		u.Text = strconv.Itoa(i)
		u.Width = i / 5
		xc.XC_CallUiThreader(u, 0)
		time.Sleep(time.Millisecond * 1)
	}

	// 解禁按钮
	xc.XC_CallUiThreadEx(func(data int) int {
		btn.Enable(true)
		btn.Redraw(true)
		return 0
	}, 0)
}

type updateButton struct {
	HEle         int
	Text         string
	Width        int
	RedrawWindow bool
}

func (u updateButton) UiThreadCallBack(data int) int {
	xc.XBtn_SetText(u.HEle, u.Text)
	xc.XEle_SetWidth(u.HEle, u.Width)
	w.Redraw(u.RedrawWindow)
	return 0
}

// 单选按钮被选择
func onBnCheck(hEle int, bCheck bool, pbHandled *bool) int {
	if bCheck {
		switch hEle {
		case radioBtn1.Handle:
			t = 1
		case radioBtn2.Handle:
			t = 2
		}
	}
	return 0
}
