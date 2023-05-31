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
	w := window.New(0, 0, 430, 300, "", 0, xcc.Window_Style_Default)

	// 1.普通编辑框
	edit := widget.NewEdit(12, 35, 100, 30, w.Handle)
	edit.SetTextColor(xc.ABGR(236, 64, 122, 255))
	edit.SetText("hello")

	// 2.密码输入框
	edit_pwd := widget.NewEdit(12, 75, 100, 30, w.Handle)
	edit_pwd.EnablePassword(true)
	edit_pwd.SetText("pwd")
	// 这个可以改变密码字符
	edit_pwd.SetPasswordCharacter('#')

	// 3.多行编辑框
	edit_MultiLine := widget.NewEdit(12, 115, 300, 70, w.Handle)
	edit_MultiLine.EnableMultiLine(true)
	edit_MultiLine.AddText("你好, 世界")
	// 添加样式
	style1 := edit_MultiLine.AddStyleEx("Arial", 14, xcc.FontStyle_Bold, xc.ABGR(0, 191, 165, 255), true)
	// 添加带样式的文本
	edit_MultiLine.AddTextEx("\nhello world", style1)

	// 获取编辑框文本
	var s string
	edit_MultiLine.GetText(&s, edit_MultiLine.GetLength()+1) // 长度必须+1
	fmt.Printf("s: %s\n\n", s)

	// 或者使用封装好的方法 GetTextEx()
	fmt.Println(edit_MultiLine.GetTextEx())

	// 4.只能输入数字的编辑框
	edit4 := widget.NewEdit(12, 195, 100, 30, w.Handle)
	edit4.Event_CHAR(func(wParam int, lParam int, pbHandled *bool) int {
		fmt.Println(wParam)
		if wParam < 58 && wParam > 47 { // 0-9
			return 0
		}

		switch wParam { // 删除,复制,撤销,剪切,全选. 有其它需要的自己添加.
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
