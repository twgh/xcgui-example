package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	w := window.NewWindow(0, 0, 484, 308, "List", 0, xcc.Window_Style_Default)

	// 创建List
	list := widget.NewList(10, 33, 464, 263, w.Handle)
	// 创建表头数据适配器
	list.CreateAdapterHeader()
	// 创建数据适配器
	list.CreateAdapter()

	// 添加列
	list.AddColumnText(147, "name1", "Column1")
	list.AddColumnText(147, "name2", "Column2")
	list.AddColumnText(147, "name3", "Column3")

	// 循环添加数据
	for i := 0; i < 20; i++ {
		// 添加行
		index := list.AddItemText(fmt.Sprintf("Column1-item%d", i))
		// 置行数据
		list.SetItemText(index, 1, fmt.Sprintf("Column2-item%d", i))
		list.SetItemText(index, 2, fmt.Sprintf("Column3-item%d", i))
	}

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
