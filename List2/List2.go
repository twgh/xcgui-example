// 列表, 模板进阶操作.
package main

import (
	_ "embed"
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed list2/list2.zip
	zipData []byte

	a  *app.App
	w  *window.Window
	ls *widget.List
)

func main() {
	a = app.New(true)
	// 从内存zip加载资源文件
	a.LoadResourceZipMem(zipData, "resource.res", "")
	w = window.New(0, 0, 302, 308, "列表, 模板进阶操作", 0, xcc.Window_Style_Default)

	// 创建List
	ls = widget.NewList(10, 33, 282, 263, w.Handle)

	// List的模板得设置两遍, 一遍是列表头, 一遍是列表项
	var hTemp int
	if hTemp = xc.XTemp_LoadZipMem(xcc.ListItemTemp_Type_List_Head, zipData, "tmpl_list.xml", ""); hTemp == 0 {
		panic("ListItemTemp_Type_List_Head: hTemp==0")
	}
	ls.SetItemTemplate(hTemp)

	if hTemp = xc.XTemp_LoadZipMem(xcc.ListItemTemp_Type_List_Item, zipData, "tmpl_list.xml", ""); hTemp == 0 {
		panic("ListItemTemp_Type_List_Item: hTemp==0")
	}
	ls.SetItemTemplate(hTemp)

	// 创建表头数据适配器
	ls.CreateAdapterHeader()
	// 创建数据适配器: 2列
	ls.CreateAdapter(2)
	// 列表_置项默认高度
	ls.SetItemHeightDefault(28, 28)

	// 添加列
	ls.AddColumnText(68, "name1", "状态")
	ls.AddColumnText(192, "name2", "操作")

	// 循环添加数据
	for i := 0; i < 8; i++ {
		// 添加行
		index := ls.AddItemText("")
		// 置行数据
		ls.SetItemText(index, 1, "")
	}

	// 注册项模板创建完成事件
	ls.Event_LIST_TEMP_CREATE_END(onLIST_TEMP_CREATE_END)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 项模板创建完成事件. 在此事件中获取按钮并注册事件
func onLIST_TEMP_CREATE_END(pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int {
	// 只在创建新模板实例的时候, 给按钮注册事件, 这样是为了避免重复注册事件
	if nFlag == 1 { // 0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		index := int(pItem.Index)
		hBtn := ls.GetTemplateObject(index, 0, 2) // 前两个参数是项索引和列索引, 第三个参数是项模板里按钮的itemID, 在设计器里是可以自己填的, 必须填了, 这里才能获取
		fmt.Println(xc.XBtn_GetText(hBtn))
		// 注册按钮事件
		xc.XEle_RegEventC1(hBtn, xcc.XE_BNCLICK, onBnClick)

		// 项模板里按钮的itemID是2,3,4
		for i := 2; i < 5; i++ {
			hBtn = ls.GetTemplateObject(index, 1, i)
			fmt.Println(xc.XBtn_GetText(hBtn))
			xc.XEle_RegEventC1(hBtn, xcc.XE_BNCLICK, onBnClick)
		}
	}
	return 0
}

// 按钮事件
func onBnClick(hEle int, pbHandled *bool) int {
	row := ls.GetItemIndexFromHXCGUI(hEle) // 获取项索引
	btnText := xc.XBtn_GetText(hEle)
	var col int // 列索引
	switch btnText {
	case "运行中":
		col = 0
	case "日志", "复制", "删除":
		col = 1
	}

	xc.XC_MessageBox("提示", fmt.Sprintf("你点击了按钮: %s, 行: %d, 列: %d", btnText, row, col), xcc.MessageBox_Flag_Ok, w.GetHWND(), xcc.Window_Style_Default)
	return 0
}
