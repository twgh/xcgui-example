// 1. 给窗口添加背景色
// 2. 给按钮加上三种状态下的图片
// 纯代码的要记好多api, 还是建议用设计器来做
package main

import (
	_ "embed"
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed res/button_min.png
	img1 []byte
	//go:embed res/button_close.png
	img2 []byte
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 465, 300, "", 0, xcc.Window_Style_Simple|xcc.Window_Style_Title|xcc.Window_Style_Drag_Window)
	// 设置窗口透明类型
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.SetShadowInfo(8, 255, 10, false, 0)
	// 给整个窗口添加背景色
	w.AddBkFill(xcc.Window_State_Flag_Leave, xc.ABGR(51, 57, 60, 254))

	// 创建最小化按钮
	btnMin := widget.NewButton(397, 8, 30, 30, "", w.Handle)
	btnMin.SetTypeEx(xcc.Button_Type_Min)
	// 创建结束按钮
	btnClose := widget.NewButton(427, 8, 30, 30, "", w.Handle)
	btnClose.SetTypeEx(xcc.Button_Type_Close)

	// 给按钮加上三种状态下的图片
	setBtnImg(btnMin, img1)
	setBtnImg(btnClose, img2)

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// 给按钮加上三态图片
func setBtnImg(btn *widget.Button, file []byte) {
	for i := 0; i < 3; i++ {
		x := int32(i * 31)
		// 图片_加载从内存, 指定区域位置及大小
		img := imagex.NewByMemRect(file, x, 0, 30, 30)

		if img.Handle == 0 {
			fmt.Println("Error: hImg=", img.Handle)
			continue
		}

		// 启用图片透明色
		img.EnableTranColor(true)
		// 添加背景图片
		switch i {
		case 0:
			btn.AddBkImage(xcc.Button_State_Flag_Leave, img.Handle)
		case 1:
			btn.AddBkImage(xcc.Button_State_Flag_Stay, img.Handle)
		case 2:
			btn.AddBkImage(xcc.Button_State_Flag_Down, img.Handle)
		}
		// 启用按钮背景透明
		btn.EnableBkTransparent(true)
	}
}
