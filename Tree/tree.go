// 树形框
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "Tree", 0, xcc.Window_Style_Default)

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

	// 列表树元素-项模板创建完成.
	// 在此事件里可以对创建完成的项模板进行操作.
	// 你需要理解项模板的概念, Tree里面有一个默认的项模板, 它的每一项都是根据固定的项模板生成的.
	// 如果你想知道默认的项模板是什么结构的, 可以打开设计器, 在一个项目中新建文件, 选择树项模板.
	// 然后你就能理解项模板是什么样子了, 项模板就是各个基础元素组合而成的, 而你可以diy它达成你想要的样子.
	// 这就是项模板存在的意义, 然后可以通过tree.SetItemTemplate相关函数设置你自己的项模板.
	// 其它有项模板的元素还有: List, ListBox, ListView等.
	tree.Event_TREE_TEMP_CREATE_END(func(pItem *xc.Tree_Item_, nFlag int32, pbHandled *bool) int {
		// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		if nFlag == 1 {
			// 获取项模板中(itemID=2)的形状文本句柄.
			// 在默认的项模板里, 文本是一个形状文本元素, 它的itemID就是2.
			hst := tree.GetTemplateObject(pItem.NID, 2)
			// 设置文本字体
			xc.XShapeText_SetFont(hst, font.NewEX("Arial", 12, xcc.FontStyle_Bold).Handle)
			// 设置文本颜色
			xc.XShapeText_SetTextColor(hst, xc.ABGR(255, 34, 33, 255))
		}
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
