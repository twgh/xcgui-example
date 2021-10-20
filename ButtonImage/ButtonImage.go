// 给按钮加上三种状态下的图片
package main

import (
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
	a := app.New("")
	// 添加资源搜索目录, 你运行时要改成自己的路径, 也可以使用相对路径
	a.AddFileSearchPath(`D:\GoProject\src\github.com\twgh\xcgui-example\ButtonImage\res`)
	// 2.创建窗口
	win := window.NewWindow(0, 0, 465, 300, "炫彩窗口", 0, xcc.Xc_Window_Style_Default)

	// 设置窗口透明类型
	win.SetTransparentType(xcc.Window_Transparent_Shadow)
	// 设置窗口阴影
	win.SetShadowInfo(10, 255, 10, false, 0)
	// 窗口置顶
	win.SetTop()
	// 窗口居中
	win.Center()

	// 设置窗口背景颜色
	bkm_win := bkmanager.NewBkManager()
	bkm_win.AddFill(xcc.Window_State_Flag_Leave, xc.RGB(51, 57, 60), 255)
	win.SetBkMagager(bkm_win.Handle)

	// 创建标签_窗口标题
	lbl_Title := widget.NewShapeText(15, 15, 56, 20, "Title", win.Handle)
	lbl_Title.SetTextColor(xc.RGB(255, 255, 255), 255)

	// 创建最小化按钮
	btn_Min := widget.NewButton(395, 10, 30, 30, "", win.Handle)
	btn_Min.SetType(xcc.Button_Type_Min)
	// 创建结束按钮
	btn_Close := widget.NewButton(425, 10, 30, 30, "", win.Handle)
	btn_Close.SetType(xcc.Button_Type_Close)
	// 启用按钮背景透明
	btn_Min.EnableBkTransparent(true)
	btn_Close.EnableBkTransparent(true)

	// 给按钮加上三态图片
	setBtnImg(`button_min.png`, btn_Min)
	setBtnImg(`button_close.png`, btn_Close)

	// 3.显示窗口
	win.ShowWindow(xcc.SW_SHOW)
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
		// 启用图片透明色
		img.EnableTranColor(true)
		// 添加背景图片
		btn.AddBkImage(i, img.Handle)
	}
}
