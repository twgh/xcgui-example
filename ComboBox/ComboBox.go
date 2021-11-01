// 组合框操作
package main

import (
	"strconv"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w := window.NewWindow(0, 0, 430, 300, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	// 创建组合框
	cbb := widget.NewComboBox(24, 50, 100, 30, w.Handle)
	// 创建数据适配器
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
	cbb.Event_ComboBox_Select_End(func(iItem int, pbHandled *bool) int {
		edit.SetText(cbb.GetItemText(iItem, 0))
		edit.Redraw(true)
		return 0
	})

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
