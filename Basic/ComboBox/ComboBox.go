// 组合框(下拉列表)
package main

import (
	"strconv"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	
	w := window.New(0, 0, 400, 240, "ComboBox", 0, xcc.Window_Style_Default)

	// 创建组合框
	cbb := widget.NewComboBox(24, 50, 100, 30, w.Handle)
	// 创建数据适配器, 这个是必须创建的, 存储数据的
	cbb.CreateAdapter()
	// 组合框加入项
	for i := 1; i <= 5; i++ {
		index := cbb.AddItemText("item" + strconv.Itoa(i))
		// 有时候会想在每个项里存个值, 比如从数据库拿到分类名和id, 项文本是分类名, 每个项里存的值是分类id
		// 但没有直接给项存值的函数, 下面这是另一种方法, 你理解为给元素内置的一个map设置键和值就行
		id := i * 10 // 从数据库拿到的分类id
		cbb.SetProperty(xc.Itoa(index), strconv.Itoa(id))
	}

	// 组合框选中项
	cbb.SetSelItem(0)
	// 组合框禁止编辑项
	cbb.EnableEdit(false)

	// 创建编辑框
	edit := widget.NewEdit(138, 50, 150, 30, w.Handle)
	edit.SetDefaultText("项文本")
	edit2 := widget.NewEdit(138, 90, 150, 30, w.Handle)
	edit2.SetDefaultText("每个项对应的值")

	// 注册组合框被选择事件
	cbb.Event_ComboBox_Select_End(func(iItem int32, pbHandled *bool) int {
		edit.SetText(cbb.GetItemText(iItem, 0))
		// 如果你是点击按钮取值的场景时, 可用 cbb.GetSelItem() 获取现在选中的项索引
		edit2.SetText(cbb.GetProperty(xc.Itoa(iItem)))
		edit.Redraw(false)
		edit2.Redraw(false)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
