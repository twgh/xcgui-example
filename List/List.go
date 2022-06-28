// 列表
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
	w := window.NewWindow(0, 0, 784, 308, "List", 0, xcc.Window_Style_Default)

	// 创建List
	list := widget.NewList(10, 33, 764, 263, w.Handle)
	// 创建表头数据适配器
	list.CreateAdapterHeader()
	// 创建数据适配器: 5列
	list.CreateAdapter(5)
	// 列表_置项默认高度
	list.SetItemHeightDefault(24, 24)

	// 添加列
	// 如果想要更好看的多功能的List就需要到设计器里设计[列表项模板], 比如说可以在项里添加按钮, 编辑框, 选择框, 组合框等, 可以任意DIY. 可参照例子: List2
	list.AddColumnText(147, "name1", "Column1")
	list.AddColumnText(147, "name2", "Column2")
	list.AddColumnText(147, "name3", "Column3")
	list.AddColumnText(147, "name4", "Column4")
	list.AddColumnText(147, "name5", "Column5")

	// 循环添加数据
	for i := 0; i < 20; i++ {
		// 添加行
		index := list.AddItemText(fmt.Sprintf("Column1-item%d", i))
		fmt.Printf("index: %v\n", index)
		// 置行数据
		list.SetItemText(index, 1, fmt.Sprintf("Column2-item%d", i))
		list.SetItemText(index, 2, fmt.Sprintf("Column3-item%d", i))
		list.SetItemText(index, 3, fmt.Sprintf("Column4-item%d", i))
		list.SetItemText(index, 4, fmt.Sprintf("Column5-item%d", i))
	}

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
