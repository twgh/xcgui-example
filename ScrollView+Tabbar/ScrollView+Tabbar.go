// 选择夹: ScrollView+TabBar的使用
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	sv  *widget.ScrollView
	tab *widget.TabBar
)

func main() {
	// 1.初始化UI库
	a := app.New("")
	// 2.创建窗口
	win := window.NewWindow(0, 0, 400, 270, "xcgui", 0, xcc.Xc_Window_Style_Default)
	// 创建结束按钮
	btn_Close := widget.NewButton(355, 5, 40, 20, "close", win.Handle)
	btn_Close.SetType(xcc.Button_Type_Close)

	// 创建选择夹顶部Tab条
	tab = widget.NewTabBar(10, 33, 380, 28, win.Handle)
	tab.AddLabel("Page 1")
	tab.AddLabel("Page 2")

	// 注册选择夹顶部按钮事件
	tab.Event_TABBAR_SELECT(tabBarSelect)

	// 创建滚动视图, 即ScrollView
	sv = widget.NewScrollView(10, 60, 380, 200, win.Handle)
	// 隐藏滚动视图的纵/横滚动条
	sv.ShowSBarH(false)
	sv.ShowSBarV(false)
	// 设置视图内容大小, 视图内容总高度400, 可分成两页, 每页高200
	sv.SetTotalSize(380, 400)
	// 禁用接收鼠标滚轮事件, 防止用户手动滚动视图
	sv.EnableEvent_XE_MOUSEWHEEL(false)

	// 第一页, 从0开始
	widget.NewButton(10, 10, 100, 20, "Button1", sv.Handle)
	widget.NewButton(10, 40, 100, 20, "Button2", sv.Handle)
	widget.NewButton(10, 70, 100, 20, "Button3", sv.Handle)
	// 第二页, 从200开始
	widget.NewButton(10, 200+10, 100, 20, "Button4", sv.Handle)
	widget.NewButton(10, 200+40, 100, 20, "Button5", sv.Handle)
	widget.NewButton(10, 200+70, 100, 20, "Button6", sv.Handle)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func tabBarSelect(iItem int, pbHandled *bool) int {
	switch iItem {
	case 0:
		sv.ScrollPosYV(0)
	case 1:
		sv.ScrollPosYV(200)
	}
	return 0
}
