// 调用 wapi 打开/保存文件, 浏览文件夹
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
		fmt.Println(wutil.OpenFile(w.Handle, []string{"All Files(*.*)", "*.*", "Text Files(*txt)", "*.txt"}, ""))
	case btn3.Handle:
		arr := wutil.OpenFileEx(wutil.OpenFileOption{
			HwndOwner:    w.GetHWND(),
			Title:        "打开文件 (最多选择2个)",
			Filters:      []string{"All Files(*.*)", "*.*", "Text Files(*txt)", "*.txt"},
			MaxOpenFiles: 2,
			// 打开多个文件时, 需要填这个
			Flags: wapi.OFN_ALLOWMULTISELECT | wapi.OFN_EXPLORER | wapi.OFN_PATHMUTEXIST,
		})

		if arr == nil && wapi.CommDlgExtendedError() == wapi.FNERR_BUFFERTOOSMALL {
			a.Alert("提示", "最多只能选择2个文件")
			return 0
		}

		for i, s := range arr {
			fmt.Printf("第%d个文件: %s\n", i+1, s)
		}
	case btn4.Handle:
		fileName := wutil.SaveFileEx(wutil.OpenFileOption{
			HwndOwner:   w.GetHWND(),
			Title:       "保存文件",
			Filters:     []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"},
			DefDir:      "D:\\",
			DefExt:      "txt",
			DefFileName: "默认文件名.txt",
		})
		fmt.Println(fileName)
	}
	return 0
}
