// 窗口局部透明, 鼠标穿透, 包括布局元素都是局部透明的
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 显示布局对象边界, 方便调试时查看
	a.ShowLayoutFrame(true)

	// 创建窗口, 如果要不显示标题栏, 可以把窗口 style 改为 Window_Style_Center, 只留个居中
	w := window.New(0, 0, 800, 600, "窗口局部透明", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// 设置为透明窗口
	w.SetTransparentType(xcc.Window_Transparent_Shaped)
	w.SetTransparentAlpha(255)

	// 创建布局元素
	width := int32(650)
	height := int32(500)
	layContent := widget.NewLayoutEle(50, 50, width, height, w.Handle)

	// 获取背景管理对象
	bkm := layContent.GetBkManagerObj()
	// 添加填充矩形, 绿色. 如果没用布局元素, 而是直接在窗口上添加的话, Element_State_Flag_Leave 这个要改的
	bkm.AddFill(xcc.Element_State_Flag_Leave, xc.RGBA(0, 255, 0, 255), 1)
	// 添加填充矩形, 蓝色
	bkm.AddFill(xcc.Element_State_Flag_Leave, xc.RGBA(0, 0, 255, 255), 2)

	// 获取填充矩形, 绿色
	obj1 := bkm.GetObjectObj(1)
	// 置外间距, 相当于设置了位置和大小
	obj1.SetMargin(30, 30, 30, 250)
	// 置矩形圆角
	obj1.SetRectRoundAngle(8, 8, 8, 8)

	// 获取填充矩形, 绿色
	obj2 := bkm.GetObjectObj(2)
	// 置外间距, 相当于设置了位置和大小
	obj2.SetMargin(30, height-250+50, 30, 30)
	// 置矩形圆角
	obj2.SetRectRoundAngle(8, 8, 8, 8)

	w.Show()
	a.Run()
	a.Exit()
}
