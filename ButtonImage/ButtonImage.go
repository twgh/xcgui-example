// 1. 给窗口添加背景色
// 2. 给按钮加上三种状态下的图片
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/bkmanager"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 2.创建窗口
	w := window.NewWindow(0, 0, 465, 300, "Title", 0, xcc.Window_Style_Simple|xcc.Window_Style_Title|xcc.Window_Style_Drag_Window)
	// 设置窗口透明类型
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	w.SetShadowInfo(8, 255, 10, false, 0)

	// 创建背景管理器
	bkm := bkmanager.NewBkManager()
	// 给整个窗口添加背景色
	bkm.AddFill(xcc.Window_State_Flag_Leave, xc.ABGR(51, 57, 60, 254))
	// 给窗口设置背景管理器
	w.SetBkMagager(bkm.Handle)

	// 创建最小化按钮
	btn_Min := widget.NewButton(397, 8, 30, 30, "", w.Handle)
	btn_Min.SetTypeEx(xcc.Button_Type_Min)
	// 创建结束按钮
	btn_Close := widget.NewButton(427, 8, 30, 30, "", w.Handle)
	btn_Close.SetTypeEx(xcc.Button_Type_Close)

	// 给按钮加上三种状态下的图片, 这里图片使用了相对路径.
	// 请使用go run运行程序, 如果你使用go build运行, 那么请把这里改成`res\button_min.png`和`res\button_close.png`
	setBtnImg(`ButtonImage\res\button_min.png`, btn_Min)
	setBtnImg(`ButtonImage\res\button_close.png`, btn_Close)

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

// 给按钮加上三态图片
func setBtnImg(fileName string, btn *widget.Button) {
	for i := 0; i < 3; i++ {
		x := i * 31
		// 图片_加载从文件指定区域, 加载图片, 指定区域位置及大小
		img := imagex.NewImage_LoadFileRect(fileName, x, 0, 30, 30)

		if img.Handle == 0 {
			fmt.Println("hImg=", img.Handle)
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
