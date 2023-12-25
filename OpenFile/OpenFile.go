// 调用 wapi 打开/保存文件, 浏览文件夹
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	a *app.App
	w *window.Window

	btn1 *widget.Button
	btn2 *widget.Button
	btn3 *widget.Button
	btn4 *widget.Button
)

func main() {
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 430, 300, "", 0, xcc.Window_Style_Default)

	// 创建按钮
	btn1 = widget.NewButton(20, 40, 100, 30, "浏览文件夹", w.Handle)
	btn2 = widget.NewButton(20, 80, 100, 30, "单选打开文件", w.Handle)
	btn3 = widget.NewButton(130, 80, 100, 30, "多选打开文件", w.Handle)
	btn4 = widget.NewButton(20, 120, 100, 30, "保存文件", w.Handle)

	// 注册按钮事件
	btn1.Event_BnClick1(onBnClick)
	btn2.Event_BnClick1(onBnClick)
	btn3.Event_BnClick1(onBnClick)
	btn4.Event_BnClick1(onBnClick)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func onBnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn1.Handle:
		fmt.Println(wutil.OpenDir(w.Handle))
	case btn2.Handle:
		fmt.Println(wutil.OpenFile(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
	case btn3.Handle:
		fmt.Println(wutil.OpenFiles(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
	case btn4.Handle:
		fmt.Println(wutil.SaveFile(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "", "默认文件名.txt"))
	}
	return 0
}
