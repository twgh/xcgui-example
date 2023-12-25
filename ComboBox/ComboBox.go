// 组合框
package main

import (
	"strconv"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "ComboBox", 0, xcc.Window_Style_Default)

	// 创建组合框
	cbb := widget.NewComboBox(24, 50, 100, 30, w.Handle)
	// 创建数据适配器, 这个是必须创建的, 存储数据的
	cbb.CreateAdapter()
	// 组合框加入项
	for i := 1; i <= 5; i++ {
		cbb.AddItemText("item" + strconv.Itoa(i))
	}

	// 组合框选中项
	cbb.SetSelItem(0)
	// 组合框禁止编辑项
	cbb.EnableEdit(false)

	// 创建编辑框
	edit := widget.NewEdit(138, 50, 100, 30, w.Handle)
	edit.SetText("hello")

	// 注册组合框被选择事件
	cbb.Event_ComboBox_Select_End(func(iItem int32, pbHandled *bool) int {
		edit.SetText(cbb.GetItemText(iItem, 0))
		edit.Redraw(false)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
