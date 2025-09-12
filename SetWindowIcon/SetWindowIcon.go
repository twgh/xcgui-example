// 设置窗口相关的图标.
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 设置程序默认窗口图标, 所有未设置图标的窗口, 都将使用此默认图标
	a.SetWindowIcon(imagex.NewByFile("SetWindowIcon\\icon1.ico").Handle)

	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xcgui window", 0, xcc.Window_Style_Default)

	// 单独设置指定窗口图标
	w.SetIcon(imagex.NewByFile("SetWindowIcon\\icon2.ico").Handle)

	// 从文件加载图标.
	// 如果想从内存加载图标, 可以把图标放到资源文件里, 就是生成的那个 syso 文件里, LoadImageW 的参数也要改, 具体看这个api介绍.
	hIcon1 := wapi.LoadImageW(0, common.StrPtr("SetWindowIcon\\icon1.ico"), wapi.IMAGE_ICON, 0, 0, wapi.LR_LOADFROMFILE|wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
	hIcon2 := wapi.LoadImageW(0, common.StrPtr("SetWindowIcon\\icon2.ico"), wapi.IMAGE_ICON, 0, 0, wapi.LR_LOADFROMFILE|wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
	fmt.Println(hIcon1, hIcon2)

	// 通过动态设置可以给不同窗口设置不同图标
	hWnd := w.GetHWND()
	// 设置大图标 (任务栏图标), 这个一般不设置, 因为编译的程序有图标的话就是那个.
	// 像 QQ 打开独立的聊天窗口, 可以设置这个以显示头像在任务栏.
	wapi.SendMessageW(hWnd, wapi.WM_SETICON, wapi.ICON_BIG, hIcon1)

	// 设置小图标 (任务栏预览左上角图标), 这个需要注意, 不设置的话, 就是没有
	wapi.SendMessageW(hWnd, wapi.WM_SETICON, wapi.ICON_SMALL, hIcon2)

	w.Show(true)
	a.Run()
	a.Exit()
}
