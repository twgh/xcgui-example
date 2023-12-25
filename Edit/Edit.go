// 编辑框
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
	w := window.New(0, 0, 340, 280, "", 0, xcc.Window_Style_Default)

	// 1.普通编辑框
	edit := widget.NewEdit(12, 35, 100, 30, w.Handle)
	edit.SetTextColor(xc.ABGR(236, 64, 122, 255))
	edit.SetText("hello")

	// 2.密码输入框
	editPwd := widget.NewEdit(12, 75, 100, 30, w.Handle)
	editPwd.EnablePassword(true)
	editPwd.SetText("pwd")
	// 这个可以改变密码字符
	editPwd.SetPasswordCharacter('#')

	// 3.多行编辑框
	editMultiLine := widget.NewEdit(12, 115, 300, 100, w.Handle)
	editMultiLine.EnableAutoShowScrollBar(true)
	editMultiLine.EnableMultiLine(true)
	editMultiLine.EnableAutoWrap(true)
	editMultiLine.AddText("你好, 世界")
	// 添加样式
	style1 := editMultiLine.AddStyleEx("Arial", 14, xcc.FontStyle_Bold, xc.ABGR(0, 191, 165, 255), true)
	// 添加带样式的文本
	editMultiLine.AddTextEx("\nhello world", style1)
	// 获取编辑框文本
	fmt.Println("获取编辑框文本:", editMultiLine.GetTextEx())

	// 设置定时器, 循环获取多行编辑框鼠标选中的文本
	editMultiLine.SetXCTimer(111, 1000)
	editMultiLine.Event_XC_TIMER(func(nTimerID int, pbHandled *bool) int {
		if nTimerID == 111 {
			text := editMultiLine.GetSelectTextEx()
			if text != "" {
				fmt.Println("选中的文本:", text)
			}
		}
		return 0
	})

	// 4.只能输入数字的编辑框
	editOnlyNumber := widget.NewEdit(12, 225, 100, 30, w.Handle)
	editOnlyNumber.Event_CHAR(func(wParam, lParam uint, pbHandled *bool) int {
		fmt.Println(wParam)
		if wParam < 58 && wParam > 47 { // 0-9
			return 0
		}

		switch wParam { // 放行删除,复制,撤销,剪切,全选. 有其它需要的自己添加.
		case 8, 3, 26, 24, 1:
			return 0
		}

		*pbHandled = true // 其它的都拦截
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
