// 列表, 模板进阶操作.
package main

import (
	_ "embed"
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/tmpl"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed res/list2.zip
	zipData []byte

	a  *app.App
	w  *window.Window
	ls *widget.List
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a = app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 从内存zip加载资源文件
	a.LoadResourceZipMem(zipData, "resource.res", "")
	// 创建窗口
	w = window.New(0, 0, 302, 308, "列表, 模板进阶操作", 0, xcc.Window_Style_Default)

	// 创建List
	ls = widget.NewList(10, 33, 282, 263, w.Handle)

	// List的模板得设置两遍, 一遍是列表项, 一遍是列表头
	var hTemp1, hTemp2 int
	tmpl.LoadZipMemEx(xcc.ListItemTemp_Type_List, zipData, "tmpl_list.xml", "", &hTemp1, &hTemp2)
	fmt.Println(hTemp1, hTemp2)
	if hTemp1 == 0 || hTemp2 == 0 {
		panic("列表头或列表项模板加载失败")
	}
	ls.SetItemTemplate(hTemp1)
	ls.SetItemTemplate(hTemp2)

	// 创建表头数据适配器
	ls.CreateAdapterHeader()
	// 创建数据适配器: 2列
	ls.CreateAdapter(2)
	// 列表_置项默认高度
	ls.SetRowHeightDefault(28, 28)

	// 添加列
	ls.AddColumnText(68, "name1", "状态")
	ls.AddColumnText(192, "name2", "操作")

	// 循环添加数据
	for i := 0; i < 8; i++ {
		// 添加行
		index := ls.AddRowText("")
		// 置行数据
		ls.SetItemText(index, 1, "")
	}

	// 列表_项模板创建完成事件
	ls.Event_LIST_TEMP_CREATE_END(onLIST_TEMP_CREATE_END)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 项模板创建完成事件. 在此事件中获取按钮并注册事件
func onLIST_TEMP_CREATE_END(pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int {
	// 只在创建新模板实例的时候, 给按钮注册事件, 这样是为了避免重复注册事件
	if nFlag == 1 { // 0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		hBtn := ls.GetTemplateObject(pItem.Index, 0, 2) // 前两个参数是项索引和列索引, 第三个参数是项模板里按钮的itemID, 在设计器里是可以自己填的, 必须填了, 这里才能获取
		fmt.Println(xc.XBtn_GetText(hBtn))
		// 注册按钮事件
		xc.XEle_RegEventC1(hBtn, xcc.XE_BNCLICK, onBnClick)

		// 项模板里按钮的itemID是2,3,4
		for i := int32(2); i < 5; i++ {
			hBtn = ls.GetTemplateObject(pItem.Index, 1, i)
			fmt.Println(xc.XBtn_GetText(hBtn))
			xc.XEle_RegEventC1(hBtn, xcc.XE_BNCLICK, onBnClick)
			// 把按钮所在列索引存进去, 可能会用到, 用不到就不存
			xc.XEle_SetUserData(hBtn, 1)
		}
	}
	return 0
}

// 按钮事件
func onBnClick(hEle int, pbHandled *bool) int {
	// 获取按钮所在的行索引
	row := ls.GetRowIndexFromHXCGUI(hEle)
	// 获取按钮所在的列索引
	col := xc.XEle_GetUserData(hEle)

	xc.XC_MessageBox("提示", fmt.Sprintf("你点击了按钮: %s, 行: %d, 列: %d", xc.XBtn_GetText(hEle), row, col), xcc.MessageBox_Flag_Ok, w.GetHWND(), xcc.Window_Style_Default)
	return 0
}
