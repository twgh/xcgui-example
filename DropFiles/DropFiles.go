// 拖放文件到窗口or元素.
package main

import (
	"fmt"
	"github.com/twgh/xcgui/wapi/wutil"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	a    *app.App
	w    *window.Window
	edit *widget.Edit
)

func main() {
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)

	w = window.New(0, 0, 600, 600, "拖放文件到窗口or元素", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// 创建编辑框.
	edit = widget.NewEdit(15, 40, 570, 300, w.Handle)
	// 编辑框允许多行.
	edit.EnableMultiLine(true)

	// 窗口_启用拖放文件.
	w.EnableDragFiles(true)

	// 注册元素文件拖放事件
	edit.Event_DROPFILES1(onEleDropFiles)

	// 注册窗口文件拖放事件.
	// w.Event_DROPFILES1(onWndDropFiles)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 事件_元素文件拖放.
func onEleDropFiles(hEle int, hDropInfo uintptr, pbHandled *bool) int {
	fmt.Println("***************************************拖放文件到元素***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt xc.POINT
	wapi.DragQueryPoint(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	files := wutil.GetDropFiles(hDropInfo)
	for _, v := range files {
		edit.AddText(v + "\n")
		fmt.Println("文件路径:", v)
	}
	return 0
}

// 事件_窗口文件拖放.
func onWndDropFiles(HXCGUI int, hDropInfo uintptr, pbHandled *bool) int {
	// win7在窗口拖放事件这里利用 [窗口_取鼠标停留元素] 可以实现对元素的拖放事件进行处理, 所以即使不注册元素拖放事件也行, 自己灵活使用..
	hEle := w.GetStayEle() // win10 好像获取不到, 那就还是去注册元素拖放事件, 不注册窗口拖放事件
	fmt.Println("鼠标停留元素句柄:", hEle)
	if hEle == edit.Handle {
		return onEleDropFiles(hEle, hDropInfo, pbHandled)
	}

	fmt.Println("***************************************拖放文件到窗口***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt xc.POINT
	wapi.DragQueryPoint(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	files := wutil.GetDropFiles(hDropInfo)
	for _, v := range files {
		fmt.Println("文件路径:", v)
	}
	return 0
}
