// 树形框
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
	w := window.NewWindow(0, 0, 430, 300, "", 0, xcc.Window_Style_Default)

	// 创建Tree
	tree := widget.NewTree(12, 33, 400, 260, w.Handle)
	// 创建数据适配器, 这个是必须的, 存储数据的
	tree.CreateAdapter()

	// 循环添加数据
	for i := 0; i < 5; i++ {
		// 插入项
		index := tree.InsertItemText(fmt.Sprintf("item%d", i), xcc.XC_ID_ROOT, xcc.XC_ID_LAST)
		// 插入2个子项
		tree.InsertItemText("subitem-1", index, xcc.XC_ID_LAST)
		subitemIndex := tree.InsertItemText("subitem-2", index, xcc.XC_ID_LAST)
		// 给子项2插入2个子项
		tree.InsertItemText("subitem-2-1", subitemIndex, xcc.XC_ID_LAST)
		tree.InsertItemText("subitem-2-2", subitemIndex, xcc.XC_ID_LAST)
	}

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
