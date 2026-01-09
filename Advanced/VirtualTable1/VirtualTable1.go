// 虚表. 支持列表, 列表框. 解决超大数据量卡顿问题.
package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	w    *window.Window
	list *widget.List
	btn  *widget.Button
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 2.创建窗口
	w = window.New(0, 0, 800, 600, "虚表", 0, xcc.Window_Style_Default)

	// 创建形状文本
	widget.NewShapeText(15, 35, 60, 30, "虚表行数: ", w.Handle)
	// 创建编辑框
	edit := widget.NewEdit(80, 35, 100, 30, w.Handle)
	edit.SetText("100000")

	// 按钮_应用虚表行数
	btn = widget.NewButton(200, 35, 100, 30, "应用虚表行数", w.Handle)
	btn.Event_BnClick1(func(hEle int, pbHandled *bool) int {
		go initListItems(xc.Atoi(edit.GetText()))
		return 0
	})

	// 通知消息_置持续时间, 设置默认持续时间.
	w.NotifyMsg_SetDuration(1000)
	// 通知消息_置宽度, 设置默认宽度.
	w.NotifyMsg_SetWidth(150)

	// 按钮_修改数据
	btnUpdate := widget.NewButton(320, 35, 100, 30, "修改数据", w.Handle)
	btnUpdate.Event_BnClick1(func(hEle int, pbHandled *bool) int {
		btnUpdate.Enable(false).Redraw(false)
		defer btnUpdate.Enable(true).Redraw(false)

		for i := 0; i < 10; i++ {
			listItems[i].Field1 = "修改: " + strconv.Itoa(rand.Intn(100000000))
		}
		list.RefreshData().Redraw(false)
		// 通知消息_窗口中弹出, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
		w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "", "修改数据成功", 0, xcc.NotifyMsg_Skin_Success, true, true, -1, 34)
		return 0
	})

	// 按钮_删除数据
	btnDelete := widget.NewButton(450, 35, 100, 30, "删除数据", w.Handle)
	btnDelete.Event_BnClick1(func(hEle int, pbHandled *bool) int {
		btnDelete.Enable(false).Redraw(false)
		defer btnDelete.Enable(true).Redraw(false)

		index := 3 // 要删除的行索引
		if len(listItems) > 0 {
			listItems = append(listItems[:index], listItems[index+1:]...)
			// 列表_置虚表行数, 刷新数据
			list.SetVirtualRowCount(int32(len(listItems))).RefreshData().Redraw(false)
			w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "", "删除数据成功", 0, xcc.NotifyMsg_Skin_Success, true, true, -1, 34)
		}
		return 0
	})

	// 按钮_插入数据
	btnInsert := widget.NewButton(580, 35, 100, 30, "插入数据", w.Handle)
	btnInsert.Event_BnClick1(func(hEle int, pbHandled *bool) int {
		btnInsert.Enable(false).Redraw(false)
		defer btnInsert.Enable(true).Redraw(false)

		index := 3 // 要插入的行索引
		if len(listItems) > 0 {
			listItems = append(listItems[:index], append([]*listItem{{
				Index:  "插入: " + strconv.Itoa(rand.Intn(100000000)),
				Field1: "插入: " + strconv.Itoa(rand.Intn(100000000)),
				Field2: "插入: " + strconv.Itoa(rand.Intn(100000000)),
				Field3: "插入: " + strconv.Itoa(rand.Intn(100000000)),
			}}, listItems[index:]...)...)
			// 列表_置虚表行数, 刷新数据
			list.SetVirtualRowCount(int32(len(listItems))).RefreshData().Redraw(false)
			w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "", "插入数据成功", 0, xcc.NotifyMsg_Skin_Success, true, true, -1, 34)
		}
		return 0
	})

	// 创建列表
	list = widget.NewList(15, 75, 770, 500, w.Handle)
	// 列表_创建列表头数据适配器
	list.CreateAdapterHeader()
	// 列表_创建数据适配器
	list.CreateAdapters(4)
	// 启用虚表
	list.EnableVirtualTable(true)

	// 添加表头
	list.AddColumnText2(120, "序号")
	list.AddColumnText2(150, "字段1")
	list.AddColumnText2(150, "字段2")
	list.AddColumnText2(150, "字段3")

	// 初始化数据
	go initListItems(xc.Atoi(edit.GetText()))

	// 列表项模板创建完成事件.
	list.Event_LIST_TEMP_CREATE_END1(func(hList int, pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int {
		// 获取列表项里的形状文本对象, 并设置文本内容
		for i := int32(0); i < list.GetColumnCount(); i++ {
			hEle := list.GetTemplateObject(pItem.Index, i, 1)
			if xc.XC_GetObjectType(hEle) == xcc.XC_SHAPE_TEXT {
				listItem := listItems[pItem.Index]
				switch i {
				case 0:
					xc.XShapeText_SetText(hEle, listItem.Index)
				case 1:
					xc.XShapeText_SetText(hEle, listItem.Field1)
				case 2:
					xc.XShapeText_SetText(hEle, listItem.Field2)
				case 3:
					xc.XShapeText_SetText(hEle, listItem.Field3)
				}
			}
		}
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

var listItems []*listItem

type listItem struct {
	Index  string // 因为不需要排序, 所以直接用了string, 要排序就改int
	Field1 string
	Field2 string
	Field3 string
}

func newListItem(index int32) *listItem {
	num := xc.Itoa(index + 1)
	return &listItem{
		Index:  num,
		Field1: "行" + num + "--字段1",
		Field2: "行" + num + "--字段2",
		Field3: "行" + num + "--字段3",
	}
}

// 初始化数据
func initListItems(count int32) {
	xc.XC_CallUT(func() {
		// 禁用按钮
		btn.Enable(false).Redraw(false)
		// 列表_设置虚表行数
		list.SetVirtualRowCount(count)
	})

	listItems = make([]*listItem, count)
	for i := int32(0); i < count; i++ {
		listItems[i] = newListItem(i)
	}

	xc.XC_CallUT(func() {
		// 列表_刷新项数据
		list.RefreshData().Redraw(false)
		// 解禁按钮
		btn.Enable(true).Redraw(false)
		// 通知消息_窗口中弹出
		w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "", "初始化数据成功", 0, xcc.NotifyMsg_Skin_Success, true, true, -1, 34)
	})
}
