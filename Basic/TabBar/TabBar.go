// 标签栏/Tab条.
package main

import (
	"fmt"
	"strconv"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	// 2.创建窗口
	w := window.New(0, 0, 570, 400, "TabBar", 0, xcc.Window_Style_Default)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// 创建标签栏
	tb := widget.NewTabBar(5, 32, w.GetWidth()-10, 30, w.Handle)
	// TAB条_启用标签带关闭按钮, 启用关闭标签功能.
	tb.EnableClose(true)

	for i := 1; i < 20; i++ {
		// 添加标签
		tb.AddLabel("标签" + strconv.Itoa(i))
	}

	// 添加TabBar_标签按钮选择改变事件.
	tb.AddEvent_TabBar_Select(func(hEle int, iItem int32, pbHandled *bool) int {
		fmt.Println("选中标签, 索引:", iItem)
		return 0
	})
	// 添加TabBar_标签按钮删除事件.
	tb.AddEvent_TabBar_Delete(func(hEle int, iItem int32, pbHandled *bool) int {
		fmt.Println("删除标签, 索引:", iItem)
		return 0
	})

	// 创建按钮
	cb := widget.NewButton(20, 100, 200, 40, "启用下拉菜单按钮", w.Handle)
	// 设置按钮类型为复选框
	cb.SetTypeEx(xcc.Button_Type_Check)
	// 设置按钮背景透明
	cb.EnableBkTransparent(true)
	// 添加按钮选中事件.
	cb.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		// TAB条_启用下拉菜单按钮.
		tb.EnableDropMenu(bCheck)
		// 因为界面布局发生改变, 所以是要调整窗口布局, 当然这个不用记, 只要是显示错乱了, 那肯定要调整父元素布局了
		w.AdjustLayout()
		w.Redraw(false)
		return 0
	})

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
