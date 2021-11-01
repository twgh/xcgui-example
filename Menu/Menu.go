// 菜单
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w := window.NewWindow(0, 0, 366, 200, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	// 创建一个按钮
	btn := widget.NewButton(50, 50, 70, 30, "menu", w.Handle)
	selected := true // 控制item3是否选中
	btn.Event_BnClick(func(pbHandled *bool) int {
		// 创建菜单
		menu := widget.NewMenu()
		// 一级菜单
		menu.AddItem(201, "item1", 0, xcc.Menu_Item_Flag_Normal)
		menu.AddItem(202, "item2", 0, xcc.Menu_Item_Flag_Normal)
		if selected {
			menu.AddItem(203, "item3", 0, xcc.Menu_Item_Flag_Check)
		} else {
			menu.AddItem(203, "item3", 0, xcc.Menu_Item_Flag_Normal)
		}

		menu.AddItem(204, "", 0, xcc.Menu_Item_Flag_Separator)
		menu.AddItem(205, "Disable", 0, xcc.Menu_Item_Flag_Disable)
		// 二级菜单
		menu.AddItem(206, "item1", 201, xcc.Menu_Item_Flag_Normal)
		menu.AddItem(207, "item2", 201, xcc.Menu_Item_Flag_Normal)

		// 获取按钮坐标
		var r xc.RECT
		btn.GetRect(&r)
		// 转换到屏幕坐标
		pt := xc.POINT{X: r.Left, Y: r.Bottom}
		xc.ClientToScreen(w.Handle, &pt)
		// 弹出菜单
		menu.Popup(w.Handle, int(pt.X), int(pt.Y), btn.Handle, xcc.Menu_Popup_Position_Left_Top)
		return 0
	})

	// 注册菜单被选择事件
	btn.Event_MENU_SELECT(func(nID int, pbHandled *bool) int {
		fmt.Println("菜单被选择:", nID)
		if nID == 203 {
			selected = !selected
		}
		return 0
	})
	// 注册菜单弹出事件
	btn.Event_MENU_POPUP(func(HMENUX int, pbHandled *bool) int {
		fmt.Println("弹出菜单")
		return 0
	})
	// 注册菜单退出事件
	btn.Event_MENU_EXIT(func(pbHandled *bool) int {
		fmt.Println("菜单退出")
		return 0
	})

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
