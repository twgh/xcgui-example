// 工具条.
package main

import (
	"fmt"
	"strconv"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	// 2.创建窗口
	w := window.New(0, 0, 570, 400, "ToolBar", 0, xcc.Window_Style_Default)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// 创建工具条
	tb := widget.NewToolBar(5, 32, w.GetWidth()-10, 30, w.Handle)

	// 按钮单击事件
	onBnClick := func(hEle int, pbHandled *bool) int {
		fmt.Println(xc.XBtn_GetText(hEle) + "被单击了")
		return 0
	}

	// 插入元素
	for i := 1; i < 10; i++ {
		// 创建按钮
		btn := widget.NewButton(0, 0, 100, 30, "按钮"+strconv.Itoa(i), tb.Handle)
		// 按钮不绘制边框
		btn.EnableDrawBorder(false)
		// 按钮不绘制焦点
		btn.EnableDrawFocus(false)
		// 注册按钮单击事件
		btn.AddEvent_BnClick(onBnClick)
		// 插入元素到工具条
		tb.InsertEle(btn.Handle, -1)
	}

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
