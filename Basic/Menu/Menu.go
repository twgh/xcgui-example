// 菜单. 默认的菜单并不好看, 好看的可参考 DrawMenu 例子, 但那个不熟练的话又不好理解.
// 用窗口放按钮来做菜单会更方便美化, 也比自绘好理解一些, 也是一个常用的做法.
package main

import (
	"fmt"

	"github.com/twgh/xcgui/wapi"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	item_selected = true // 控制item_select是否选中
)

const (
	item1 = iota + 10000
	subitem1
	subitem2

	item2
	item_select
	item_disable
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 2.创建窗口
	w := window.New(0, 0, 400, 300, "Menu", 0, xcc.Window_Style_Default)

	// 启用窗口布局
	w.EnableLayout(true)
	// 水平居中
	w.SetAlignH(xcc.Layout_Align_Center)

	widget.NewShapeText(0, 0, 400, 30, "点击鼠标右键显示菜单", w.Handle).SetFont(app.NewFont(16).Handle).LayoutItem_SetWidth(xcc.Layout_Size_Auto, 0)

	// 窗口鼠标右键弹起事件
	w.AddEvent_RButtonUp(onWindowRButtonUp)

	// 注册菜单被选择事件
	w.Event_MENU_SELECT(onMenuSelect)
	// 注册菜单弹出事件
	w.Event_MENU_POPUP(onMenuPopup)
	// 注册菜单退出事件
	w.Event_MENU_EXIT(onMenuExit)

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// 窗口鼠标右键弹起事件
func onWindowRButtonUp(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
	// 创建菜单
	menu := widget.NewMenu()
	// 一级菜单
	menu.AddItem(item1, "item1", 0, xcc.Menu_Item_Flag_Normal)
	menu.AddItem(item2, "item2", 0, xcc.Menu_Item_Flag_Normal)
	if item_selected {
		menu.AddItem(item_select, "item_select", 0, xcc.Menu_Item_Flag_Check)
	} else {
		menu.AddItem(item_select, "item_select", 0, xcc.Menu_Item_Flag_Normal)
	}
	menu.AddItem(-1, "", 0, xcc.Menu_Item_Flag_Separator) // 分隔栏
	menu.AddItem(item_disable, "item_disable", 0, xcc.Menu_Item_Flag_Disable)

	// item1的二级菜单
	menu.AddItem(subitem1, "subitem1", item1, xcc.Menu_Item_Flag_Normal)
	menu.AddItem(subitem2, "subitem2", item1, xcc.Menu_Item_Flag_Normal)

	// 获取鼠标光标的位置
	var pt wapi.POINT
	wapi.GetCursorPos(&pt)
	// 弹出菜单
	menu.Popup(xc.XWnd_GetHWND(hWindow), pt.X, pt.Y, 0, xcc.Menu_Popup_Position_Left_Top)
	return 0
}

// 菜单被选择事件
func onMenuSelect(nID int32, pbHandled *bool) int {
	fmt.Println("菜单被选择:", nID)
	if nID == item_select {
		item_selected = !item_selected
	}
	return 0
}

// 菜单弹出事件
func onMenuPopup(HMENUX int, pbHandled *bool) int {
	fmt.Println("弹出菜单")
	return 0
}

// 菜单退出事件
func onMenuExit(pbHandled *bool) int {
	fmt.Println("菜单退出")
	return 0
}
