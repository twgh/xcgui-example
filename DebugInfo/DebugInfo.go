// 炫彩资源监视器和DebugInfo
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)

	// 炫彩_显示布局边界
	a.ShowLayoutFrame(true)
	// 显示svg边界
	a.ShowSvgFrame(true)
	// 炫彩_启用资源监视器. 按 Ctrl+~ 键呼出炫彩资源监视器窗口.
	a.EnableResMonitor(true)

	// 创建窗口
	w := window.New(0, 0, 600, 400, "Test", 0, xcc.Window_Style_Default)
	// 窗口启用布局, 布局盒子水平居中
	w.EnableLayout(true).SetAlignH(xcc.Layout_Align_Center)
	// 窗口布局盒子垂直居中
	w.SetAlignV(xcc.Layout_Align_Center)
	// 创建形状文本, 设置字体大小, 形状文本自动调节宽度
	widget.NewShapeText(0, 0, 300, 100, "按 Ctrl+~ 键呼出炫彩资源监视器窗口", w.Handle).SetFont(font.New(20).Handle).LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)

	// 炫彩_启用debug文件. 文本会输出到程序运行目录的xcgui_debug.txt.
	// 这个的主要作用是帮助你排查炫彩的一些弹窗报错.
	// 有时候你的一些操作导致弹窗报错xx句柄无效, 看debug文本就可以排查了.
	a.EnableDebugFile(true)

	// 输出自定义文本到deubg文件
	a.DebugToFileInfo("调试信息DebugToFileInfo")
	xc.XDebug_Print(0, "调试信息0")
	xc.XDebug_Print(1, "调试信息1")

	// 级别大于1会弹出一个错误提示窗
	// xc.XDebug_Print(2, "弹一个错误提示窗, 点否可继续往下运行")

	w.Show(true)
	a.Run()
	a.Exit()
}
