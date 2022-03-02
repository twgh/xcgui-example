// 拖放文件到窗口or元素.
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/shell32"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	a *app.App
	w *window.Window

	edit *widget.Edit
)

func main() {
	a = app.New(true)
	defer a.Exit()

	w = window.NewWindow(0, 0, 600, 600, "拖放文件到窗口or元素", 0, xcc.Window_Style_Default)
	// 窗口_启用拖放文件.
	w.EnableDragFiles(true)
	// 注册窗口文件拖放事件.
	w.Event_DROPFILES1(onWndDropFiles)

	// 创建编辑框.
	edit = widget.NewEdit(15, 40, 570, 300, w.Handle)
	// 编辑框允许多行.
	edit.EnableMultiLine(true)
	// 注册元素文件拖放事件, 可注册也可不注册, 因为在窗口拖放事件里也可以处理元素的, 根据自己的需求来吧, 很灵活.
	//edit.Event_DROPFILES1(onEleDropFiles)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
}

// 事件_窗口文件拖放.
func onWndDropFiles(HXCGUI, hDropInfo int, pbHandled *bool) int {
	// 在窗口拖放事件这里可以实现对其他元素的拖放事件进行处理, 所以即使不注册元素拖放事件也行, 自己灵活使用..
	// 窗口_取鼠标停留元素.
	hEle := w.GetStayEle()
	fmt.Println("鼠标停留元素句柄:", hEle)
	if hEle == edit.Handle {
		return onEleDropFiles(hEle, hDropInfo, pbHandled)
	}

	fmt.Println("***************************************拖放文件到窗口***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt xc.POINT
	shell32.DragQueryPoint(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	// 循环获取拖放进窗口的所有文件.
	i := 0
	for {
		filePath := ""
		length := shell32.DragQueryFileW(hDropInfo, i, &filePath, 260)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}

		fmt.Println("文件路径:", filePath)
		i++ // 索引+1检索下一个文件
	}

	shell32.DragFinish(hDropInfo)
	return 0
}

// 事件_元素文件拖放.
func onEleDropFiles(HXCGUI, hDropInfo int, pbHandled *bool) int {
	fmt.Println("***************************************拖放文件到元素***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt xc.POINT
	shell32.DragQueryPoint(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	// 循环获取拖放进元素的所有文件.
	i := 0
	for {
		filePath := ""
		length := shell32.DragQueryFileW(hDropInfo, i, &filePath, 260)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}

		edit.AddText(filePath + "\r\n")
		fmt.Println("文件路径:", filePath)
		i++ // 索引+1检索下一个文件
	}

	shell32.DragFinish(hDropInfo)
	return 0
}
