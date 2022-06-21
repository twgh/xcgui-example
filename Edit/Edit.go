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
	w := window.NewWindow(0, 0, 430, 300, "", 0, xcc.Window_Style_Default)

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
	edit_MultiLine := widget.NewEdit(12, 115, 300, 100, w.Handle)
	edit_MultiLine.EnableMultiLine(true)
	edit_MultiLine.AddText("你好, 世界")

	// 添加样式
	style1 := edit_MultiLine.AddStyleEx("Arial", 12, xcc.FontStyle_Bold, xc.ABGR(0, 191, 165, 255), true)
	// 添加带样式的文本
	edit_MultiLine.AddTextEx("\nhello world", style1)

	// 获取编辑框文本
	var s string
	edit_MultiLine.GetText(&s, edit_MultiLine.GetLength()+1) // 长度必须+1
	fmt.Printf("s: %s\n", s)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
