// 设置默认字体, 获取字体信息
package main

import (
	"fmt"
	"syscall"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)

	// 创建字体
	f := font.NewFontEX("Microsoft YaHei", 11, xcc.FontStyle_Regular)
	// 设置程序默认字体
	a.SetDefaultFont(f.Handle)

	// 2.创建窗口
	w := window.NewWindow(0, 0, 466, 300, "xc", 0, xcc.Window_Style_Simple|xcc.Window_Style_Btn_Close)

	// 创建一个按钮
	btn := widget.NewButton(30, 50, 150, 30, "GetFontInfo", w.Handle)
	btn.Event_BnClick(func(pbHandled *bool) int {
		// 获取字体信息
		var fontInfo xc.Font_Info_
		f.GetFontInfo(&fontInfo)
		w.MessageBox(fmt.Sprintf("fontName=%s, fontSize=%d, fontStyle=%d", uint16ToString(fontInfo.Name), fontInfo.NSize, fontInfo.NStyle), "Font Info", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Pop)
		return 0
	})

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// uint16到string
func uint16ToString(str [32]uint16) string {
	return syscall.UTF16ToString(str[0:])
}
