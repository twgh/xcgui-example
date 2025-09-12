// 发送事件/发送消息.
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	// 2.创建窗口
	w1 := window.New(0, 0, 430, 300, "窗口1", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	w2 := window.New(0, 0, 300, 150, "窗口2", w1.GetHWND(), xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)

	// 创建按钮
	btnCheck := widget.NewButton(10, 30, 200, 30, "窗口2跟随窗口1移动", w1.Handle)
	// 设置按钮类型为复选框
	btnCheck.SetTypeEx(xcc.Button_Type_Check)
	// 设置按钮背景透明
	btnCheck.EnableBkTransparent(true)

	// 注册按钮选中事件
	btnCheck.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		btnCheck.Enable(false).Redraw(false)
		defer btnCheck.Enable(true).Redraw(false)

		if bCheck {
			// 你可能觉得窗口事件不全, 比如没有 Event_Move 这样的函数, 因为不需要, 在消息过程里判断即可.
			// https://mcn1fno5w69l.feishu.cn/wiki/AdQKwX8aXirnpikqiUEccPKDnbd

			// 添加窗口消息过程事件.
			w1.AddEvent_WindProc(func(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
				if message == wapi.WM_MOVE { // 窗口移动消息.
					rc1 := w1.GetRectEx()
					rc2 := w2.GetRectEx()

					width := rc2.Right - rc2.Left
					height := rc2.Bottom - rc2.Top
					rc2.Top = rc1.Top
					rc2.Left = rc1.Right
					rc2.Right = rc2.Left + width
					rc2.Bottom = rc2.Top + height

					w2.SetRect(&rc2).Redraw(true)
				}
				return 0
			})

			// 上面添加完事件, 由于窗口1没有移动, 所以窗口2并不会立即移动, 所以就可以手动触发一下消息,
			// 这会把消息传进窗口消息循环.
			w1.SendMessage(wapi.WM_MOVE, 0, 0)
		} else {
			w1.RemoveEvent(xcc.WM_MOVE)
		}
		return 0
	})

	{
		btn1 := widget.NewButton(12, 80, 100, 30, "按钮1", w1.Handle)
		btn1.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			fmt.Println("按钮1被点击了")
			return 0
		})

		btn2 := widget.NewButton(12, 120, 100, 30, "按钮2", w1.Handle)
		btn2.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			fmt.Println("按钮2被点击了")
			// 元素_发送事件, 实际上也是发送消息.
			btn1.SendEvent(xcc.XE_BNCLICK, 0, 0)
			return 0
		})
	}

	w1.Show(true)
	w2.Show(true)
	a.Run()
	a.Exit()
}
