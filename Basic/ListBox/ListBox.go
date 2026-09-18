// 列表框
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 430, 500, "ListBox", 0, xcc.Window_Style_Default)

	// 创建形状文本, 显示当前选中项
	label := widget.NewShapeText(12, 32, 400, 22, "当前选中: 未选中", w.Handle)

	// 创建ListBox
	lb := widget.NewListBox(12, 60, 400, 380, w.Handle)

	// 创建数据适配器, 这个必须创建, 存储数据的
	lb.CreateAdapter()

	// 项计数器, 用于给新添加的项命名
	itemCount := 0

	for i := 0; i < 8; i++ {
		// 添加行, 返回项索引
		lb.AddItemText(fmt.Sprintf("item-%d", itemCount))
		itemCount++
	}

	// 更新选中项显示的回调: GetSelectItem 取当前选中索引, GetItemText 取项文本
	updateSelText := func() {
		sel := lb.GetSelectItem() // 返回当前选中项索引, 没有选中时为 -1
		if sel < 0 {
			label.SetText("当前选中: 未选中")
		} else {
			// GetItemText 取指定项的文本, 第2个参数是列索引, 内置模板只有1列, 填0
			itemText := lb.GetItemText(sel, 0)
			label.SetText(fmt.Sprintf("当前选中: 索引=%d, 文本=%s", sel, itemText))
		}
		label.Redraw() // 修改元素显示后必须重绘
	}

	// 注册列表框项选择事件, 选中项改变时更新显示
	lb.AddEvent_ListBox_Select(func(hEle int, iItem int32, pbHandled *bool) int {
		updateSelText()
		return 0
	})

	// 默认选中第0项, 会触发选择事件
	lb.SetSelectItem(0)

	// 创建按钮行: 插入一项
	btnInsert := widget.NewButton(12, 455, 120, 28, "插入到最前面", w.Handle)
	btnInsert.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		itemCount++
		// 在索引0处插入一项, 返回插入后的项索引
		lb.InsertItemText(0, fmt.Sprintf("insert-%d", itemCount))
		lb.Redraw(false) // 列表内容改变后重绘
		return 0
	})

	// 创建按钮行: 追加一项
	btnAppend := widget.NewButton(152, 455, 120, 28, "追加到末尾", w.Handle)
	btnAppend.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		itemCount++
		lb.AddItemText(fmt.Sprintf("append-%d", itemCount))
		lb.Redraw(false)
		return 0
	})

	// 创建按钮行: 删除选中项
	btnDelete := widget.NewButton(292, 455, 120, 28, "删除选中项", w.Handle)
	btnDelete.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		sel := lb.GetSelectItem() // 获取当前选中索引
		if sel < 0 {
			w.MessageBox("提示", "请先选中一项", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
			return 0
		}
		// 按索引删除一项
		lb.DeleteItem(sel)
		lb.Redraw(false)
		updateSelText() // 删除后选中状态会变化, 重新显示选中状态
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
