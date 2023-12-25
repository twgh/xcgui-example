// 菜单
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
	a   *app.App
	w   *window.Window
	btn *widget.Button

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
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w = window.New(0, 0, 400, 300, "Menu", 0, xcc.Window_Style_Default)

	// 创建一个按钮
	btn = widget.NewButton(50, 50, 100, 30, "ShowMenu", w.Handle)
	// 注册按钮被单击事件
	btn.Event_BnClick(onBnClick)

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

// 按钮被单击事件
func onBnClick(pbHandled *bool) int {
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

	// 获取按钮坐标
	var rc xc.RECT
	btn.GetWndClientRectDPI(&rc)
	// 转换到屏幕坐标
	pt := wapi.POINT{X: rc.Left, Y: rc.Bottom}
	wapi.ClientToScreen(w.GetHWND(), &pt)
	// 弹出菜单
	menu.Popup(w.GetHWND(), pt.X, pt.Y, 0, xcc.Menu_Popup_Position_Left_Top)
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
