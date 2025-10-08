// 列表: 添加行, 删除选中行, 清空行, 排序, 表头表项文本居中, 双击编辑列表项, 显示指定行
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	w    *window.Window
	list *widget.List

	btnAdd   *widget.Button
	btnDel   *widget.Button
	btnClear *widget.Button
	btnJump  *widget.Button

	editLine *widget.Edit
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w = window.New(0, 0, 784, 400, "List", 0, xcc.Window_Style_Default)

	// 创建List
	createList()
	// List添加行
	listAddItem()

	var startX int32 = 10
	btnAdd = widget.NewButton(startX, 35, 100, 30, "添加20行", w.Handle)
	btnAdd.Event_BnClick1(onBnClick)

	startX += 100 + 3
	btnDel = widget.NewButton(startX, 35, 100, 30, "删除选中行", w.Handle)
	btnDel.Event_BnClick1(onBnClick)

	startX += 100 + 3
	btnClear = widget.NewButton(startX, 35, 100, 30, "删除所有行", w.Handle)
	btnClear.Event_BnClick1(onBnClick)

	startX += 100 + 3
	btnJump = widget.NewButton(startX, 35, 100, 30, "跳转指定行", w.Handle)
	btnJump.Event_BnClick1(onBnClick)

	startX += 100 + 3
	editLine = widget.NewEdit(startX, 35, 100, 30, w.Handle)

	w.Show(true)
	a.Run()
	a.Exit()
}

// 按钮单击事件
func onBnClick(hEle int, pbHandled *bool) int {
	xc.XEle_Enable(hEle, false) // 操作前禁用按钮

	switch hEle {
	case btnAdd.Handle:
		listAddItem()
	case btnDel.Handle:
		listDelSelectItem()
	case btnClear.Handle:
		list.DeleteRowAll()
		list.Redraw(true)
	case btnJump.Handle:
		row := xc.Atoi(editLine.GetTextEx()) - 1
		if row > -1 && row < list.GetCount_AD() {
			list.VisibleRow(row)
			list.Redraw(true)
		}
	}

	xc.XEle_Enable(hEle, true) // 操作后解禁按钮
	return 0
}

// 创建List
func createList() {
	// 创建List
	list = widget.NewList(10, 70, 764, 315, w.Handle)
	// 创建表头数据适配器
	list.CreateAdapterHeader()
	// 创建数据适配器: 5列
	list.CreateAdapter(5)

	// 列表_置项默认高度和选中时高度
	list.SetRowHeightDefault(24, 26)
	// 列表_绘制项分割线
	// list.SetDrawRowBkFlags(xcc.List_DrawItemBk_Flag_Line | xcc.List_DrawItemBk_Flag_LineV | xcc.List_DrawItemBk_Flag_Leave | xcc.List_DrawItemBk_Flag_Stay | xcc.List_DrawItemBk_Flag_Select)
	// 表头和表项居中
	listTextAlign()

	// 添加列
	// 如果想要更好看的多功能的List就需要到设计器里设计[列表项模板], 比如说可以在项里添加按钮, 编辑框, 选择框, 组合框等, 可以任意DIY. 可参照例子: List2
	list.AddColumnText(50, "name1", "序号")
	list.AddColumnText(147, "name2", "Column2")
	list.AddColumnText(147, "name3", "Column3")
	list.AddColumnText(147, "name4", "Column4")
	list.AddColumnText(147, "name5", "Column5")

	// 设置序号列可排序, 单击表头时排序
	list.SetSort(0, 0, true)
	// 这里我使用了置属性的方法是为了不新建多个变量, 因为考虑到组件可能会很多, 当然你也可以用变量来控制.
	// 这个置属性你可以理解为就是给元素绑定的map中赋值. 并不是在操作元素的属性.
	// 也是为了演示Set/GetProperty, 这个东西很有用, 比如说你的列表每1行都有隐藏的值, 就可以存在这里.
	list.SetProperty("sortType", "1") // 1是正序, 0是倒序.
	list.SetProperty("sortFlag", "0") // 只是我设定的标记

	// 列表头项被单击事件
	list.Event_LIST_HEADER_CLICK(func(iItem int32, pbHandled *bool) int {
		// 为了记录排序类型
		if iItem == 0 {
			// 下面这个sortFlag只是我设定的1个标记, 意义是让第1次单击表头排序时不设置sortType的值, 因为第1次默认就是正序
			if list.GetProperty("sortFlag") != "1" {
				list.SetProperty("sortFlag", "1")
			} else {
				if list.GetProperty("sortType") == "1" {
					list.SetProperty("sortType", "0")
					fmt.Println("列表当前排序: 倒序")
				} else {
					list.SetProperty("sortType", "1")
					fmt.Println("列表当前排序: 正序")
				}
			}
		}
		return 0
	})

	// 列表_鼠标左键双击事件
	list.Event_LBUTTONDBCLICK(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		// 取鼠标点击的行和列
		var row, column int32
		list.HitTestOffset(pPt, &row, &column)
		fmt.Println("双击行索引:", row, "列索引:", column)
		if row < 0 || column < 0 {
			return 0
		}
		// 取列表行高
		var height int32
		list.GetRowHeight(row, &height, &height)

		// 创建编辑框
		// 获取双击项的布局元素句柄. 列表默认项都是一个布局元素里放一个形状文本, 布局元素itemID: 0, 形状文本itemID: 1
		// 至于我是怎么知道的, 这个是打开设计器创建一个列表项模板文件后, 就知道它里面是什么了
		hLayout := list.GetTemplateObject(row, column, 0)
		if hLayout == 0 {
			return 0
		}
		// 取列宽度
		width := list.GetColumnWidth(column)
		edit := widget.NewEdit(0, 0, width, height, hLayout)
		// 设置编辑框的值为当前列表项内容, 并全选, 置焦点, 文本格式水平垂直居中
		text := list.GetItemText(row, column)
		edit.SetTextAlign(xcc.Edit_TextAlign_Flag_Center | xcc.Edit_TextAlign_Flag_Center_V)
		edit.SetText(text)
		if text != "" {
			edit.SelectAll()
		}
		w.SetFocusEle(edit.Handle)

		// 存储当前列表项的行和列, 下面要用
		edit.SetProperty("row", xc.Itoa(row))
		edit.SetProperty("column", xc.Itoa(column))

		// 编辑框键盘按下事件, 确认修改列表项文本
		edit.Event_KEYDOWN1(onEleKeyDown)
		// 编辑框失去焦点事件, 销毁编辑框
		edit.Event_KILLFOCUS1(onEleKillFocus)
		return 0
	})
}

// 元素键盘按下事件
func onEleKeyDown(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
	switch wParam {
	case xcc.VK_Enter: // 回车键
		row := xc.Atoi(xc.XC_GetProperty(hEle, "row"))
		column := xc.Atoi(xc.XC_GetProperty(hEle, "column"))
		list.SetItemText(row, column, xc.XEdit_GetText_Temp(hEle))
		xc.XEle_Destroy(hEle)
		list.RefreshRow(row)
		list.Redraw(true)
	case xcc.VK_Esc: // Esc键
		xc.XEle_Destroy(hEle)
		list.Redraw(true)
	}
	return 0
}

func onEleKillFocus(hEle int, pbHandled *bool) int {
	xc.XEle_Destroy(hEle)
	list.Redraw(true)
	return 0
}

// 表头和表项居中, 纯代码实现需要记一些api, 需要有清晰的思维, 还是用设计器来的简单, 真要写大程序不可能离开设计器的
func listTextAlign() {
	list.Event_LIST_HEADER_TEMP_CREATE_END(func(pItem *xc.List_Header_Item_, pbHandled *bool) int {
		for i := int32(0); i < list.GetColumnCount(); i++ {
			hEle := list.GetHeaderTemplateObject(i, 1)
			if app.IsHXCGUI(hEle, xcc.XC_SHAPE_TEXT) { // 是形状文本
				xc.XShapeText_SetTextAlign(hEle, xcc.TextAlignFlag_Center|xcc.TextAlignFlag_Vcenter)
			}
		}
		return 0
	})

	list.Event_LIST_TEMP_CREATE_END(func(pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int {
		// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		if nFlag == 1 {
			for i := int32(0); i < list.GetColumnCount(); i++ {
				hEle := list.GetTemplateObject(pItem.Index, i, 1)
				if app.IsHXCGUI(hEle, xcc.XC_SHAPE_TEXT) { // 是形状文本
					xc.XShapeText_SetTextAlign(hEle, xcc.TextAlignFlag_Center|xcc.TextAlignFlag_Vcenter)
				}
			}
		}
		return 0
	})
}

// List添加20行
func listAddItem() {
	// 循环添加数据
	for i := 0; i < 20; i++ {
		num := list.GetCount_AD() + 1

		// 添加行
		var index int32
		if list.GetProperty("sortType") == "1" { // 正序
			index = list.AddRowTextEx("name2", fmt.Sprintf("item%d-Column2", num))
		} else { // 倒序
			index = list.InsertRowTextEx(0, "name2", fmt.Sprintf("item%d-Column2", num))
		}
		fmt.Printf("添加行索引: %d\n", index)

		// 置行数据
		// 序号列设置int型的数据才能按数字大小排序
		list.SetItemInt(index, 0, num)
		list.SetItemText(index, 2, fmt.Sprintf("item%d-Column3", num))
		list.SetItemText(index, 3, fmt.Sprintf("item%d-Column4", num))
		list.SetItemText(index, 4, fmt.Sprintf("item%d-Column5", num))
	}

	list.Redraw(true)
}

// List删除选中行
func listDelSelectItem() {
	count := list.GetSelectRowCount()
	if count == 0 {
		w.MessageBox("提示", "你没有选中列表任何行!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Modal)
		return
	}

	var indexArr []int32
	// 取选中行索引数组
	list.GetSelectAll(&indexArr, count)
	// 根据选中行索引数组倒着删, 正着删的话你每删1行下面的行索引就变了
	for i := count - 1; i > -1; i-- {
		list.DeleteRow(indexArr[i])
		fmt.Printf("删除行索引: %d\n", indexArr[i])
	}

	// 重排剩余行序号
	count = list.GetCount_AD()
	if list.GetProperty("sortType") == "1" { // 正序
		for i := int32(0); i < count; i++ {
			list.SetItemInt(i, 0, i+1)
		}
	} else { // 倒序
		for i, num := int32(0), count; i < count; i, num = i+1, num-1 {
			list.SetItemInt(i, 0, num)
		}
	}

	// 刷新列表项数据
	list.RefreshData()
	// 列表重绘
	list.Redraw(true)
}
