// 菜单条.
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	a  *app.App
	w  *window.Window
	mb *widget.MenuBar
)

func main() {
	// 1.初始化UI库
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w = window.New(0, 0, 570, 400, "MenuBar", 0, xcc.Window_Style_Default)
	w.SetBorderSize(1, 30, 1, 1)

	// 创建菜单条
	mb = widget.NewMenuBar(5, 32, w.GetWidth()-10, 30, w.Handle)
	// 菜单条禁用按钮自动宽度
	mb.EnableAutoWidth(false)
	// 菜单条禁用绘制边框
	mb.EnableDrawBorder(false)
	// 菜单条禁用绘制焦点
	mb.EnableDrawFocus(false)

	// 添加菜单条按钮
	mb.AddButton("文件(F)")
	mb.AddButton("编辑(E)")
	mb.AddButton("搜索(S)")
	mb.AddButton("视图(V)")
	mb.AddButton("编码(N)")
	mb.AddButton("语言(L)")
	mb.AddButton("设置(T)")
	mb.AddButton("工具(O)")

	// 取消绘制菜单条按钮边框
	for i := 0; i < mb.GetChildCount(); i++ {
		hele := mb.GetChildByIndex(i)
		xc.XEle_EnableDrawBorder(hele, false)
		xc.XEle_RegEventC1(hele, xcc.XE_BNCLICK, onMenuBarBnClick)
		xc.XEle_RegEventC1(hele, xcc.XE_MENU_SELECT, onMenuSelect)
	}

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func onMenuBarBnClick(hEle int, pbHandled *bool) int {
	fmt.Println(xc.XBtn_GetText(hEle) + "被单击了")
	// 创建菜单
	menu := widget.NewMenu()
	// 一级菜单
	menu.AddItem(10001, "item1", 0, xcc.Menu_Item_Flag_Normal)
	menu.AddItem(10002, "item2", 0, xcc.Menu_Item_Flag_Normal)

	// 获取按钮坐标
	var rc xc.RECT
	xc.XEle_GetWndClientRectDPI(hEle, &rc)
	// 转换到屏幕坐标
	pt := wapi.POINT{X: rc.Left, Y: rc.Bottom}
	wapi.ClientToScreen(w.GetHWND(), &pt)
	// 弹出菜单
	menu.Popup(w.GetHWND(), pt.X, pt.Y, hEle, xcc.Menu_Popup_Position_Left_Top)
	return 0
}

// 菜单被选择事件
func onMenuSelect(hEle int, nID int32, pbHandled *bool) int {
	fmt.Println(xc.XBtn_GetText(hEle)+"下的菜单被选择:", nID)
	return 0
}
