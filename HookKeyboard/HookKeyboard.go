// 创建全局键盘钩子, 监听键盘消息, 可当热键使用或其他用途.
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 启用自适应DPI
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "全局键盘钩子", 0, xcc.Window_Style_Default)

	widget.NewShapeText(40, 40, 300, 30, "在任何窗口按键都能够监听到", w.Handle)
	widget.NewEdit(40, 80, 300, 30, w.Handle).SetFocus()
	checkBtn := widget.NewButton(40, 120, 300, 30, "拦截A键按下", w.Handle).SetTypeEx(xcc.Button_Type_Check)
	checkBtn.EnableBkTransparent(true)

	kbHook := wutil.NewHookKeyboard(func(nCode int32, wParam xcc.WM_, lParam *wapi.KBDLLHOOKSTRUCT) uintptr {
		if nCode < 0 { // nCode小于0时不应继续处理
			return wutil.CallNextHookEx_Keyboard(nCode, wParam, lParam)
		}

		if wParam == xcc.WM_KEYDOWN { // 键盘按下
			if checkBtn.GetStateEx() == xcc.Button_State_Check {
				if lParam.VkCode == xcc.VK_A {
					fmt.Println("拦截了A键按下, 是不会输入文本框的, 部分程序不会被拦截, 自行研究")
					return 1 // 返回1可拦截, 这时按下A键是不会输入文本框的, 部分程序不会被拦截, 因为它可能进行了特殊处理
				}
			}
			fmt.Printf("按键按下: 虚拟键码=%d, 扫描码=%d\n", lParam.VkCode, lParam.ScanCode)
		} else if wParam == xcc.WM_KEYUP { // 键盘弹起
			fmt.Printf("按键弹起: 虚拟键码=%d, 扫描码=%d\n", lParam.VkCode, lParam.ScanCode)
		}
		return wutil.CallNextHookEx_Keyboard(nCode, wParam, lParam)
	})

	w.Event_CLOSE(func(pbHandled *bool) int {
		kbHook.Unhook()
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
