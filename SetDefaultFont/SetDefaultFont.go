package main

import (
	"fmt"
	"syscall"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/fontx"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New("")

	// 创建字体
	font := fontx.NewFontX2("msyh", 12, xcc.FontStyle_Regular)
	// 设置程序默认字体
	a.SetDefaultFont(font.HFontX)

	// 2.创建窗口
	win := window.NewWindow(0, 0, 466, 300, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

	// 设置窗口边框大小
	win.SetBorderSize(1, 30, 1, 1)
	// 窗口置顶
	win.SetTop()
	// 窗口居中
	win.Center()
	// 创建标签_窗口标题
	lbl_Title := widget.NewShapeText(5, 5, 56, 20, "Title", win.Handle)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// 创建结束按钮
	btn_Close := widget.NewButton(436, 0, 30, 30, "X", win.Handle)
	btn_Close.SetTextColor(xc.RGB(255, 255, 255), 255)
	btn_Close.SetType(xcc.Button_Type_Close)
	btn_Close.EnableBkTransparent(true)

	// 获取字体信息
	var fontInfo xc.Font_Info_
	font.GetFontInfo(&fontInfo)
	fmt.Println(uint16ToString(fontInfo.Name), fontInfo.NSize, fontInfo.NStyle)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// uint16到string
func uint16ToString(str [32]uint16) string {
	return syscall.UTF16ToString(str[0:])
}
