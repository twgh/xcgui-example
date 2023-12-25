// 注册热键, 全局的或窗口内的.
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

// 注册热键时填的id
const (
	ID_F3 = iota + 1
	ID_F4
)

var (
	a *app.App
	w *window.Window
)

func main() {
	a = app.New(true)
	w = window.New(0, 0, 400, 300, "", 0, xcc.Window_Style_Default)

	// 全局生效
	one()
	// 只在窗口内生效, 窗口内热键
	two()

	w.Show(true)
	a.Run()
	a.Exit()
}

// 全局生效
func one() {
	// 注册热键F3
	if !wapi.RegisterHotKey(w.GetHWND(), ID_F3, 0, xcc.VK_F3) {
		fmt.Println("注册热键F3失败")
	}

	// 注册热键F4
	if !wapi.RegisterHotKey(w.GetHWND(), ID_F4, 0, xcc.VK_F4) {
		fmt.Println("注册热键F4失败")
	}

	w.Event_WINDPROC1(func(hWindow int, message uint32, wParam, lParam uint, pbHandled *bool) int {
		if message == uint32(xcc.WM_HOTKEY) {
			switch wParam {
			case ID_F3:
				fmt.Println("Event_WINDPROC1 F3键被按下")
			case ID_F4:
				fmt.Println("Event_WINDPROC1 F4键被按下")
			}
		}
		return 0
	})
}

// 只在窗口内生效, 窗口内热键
func two() {
	w.Event_KEYDOWN1(func(hWindow int, wParam, lParam uint, pbHandled *bool) int {
		switch wParam {
		case xcc.VK_F5:
			fmt.Println("Event_KEYDOWN1 F5键被按下")
		case xcc.VK_F6:
			fmt.Println("Event_KEYDOWN1 F6键被按下")
		}
		return 0
	})
}
