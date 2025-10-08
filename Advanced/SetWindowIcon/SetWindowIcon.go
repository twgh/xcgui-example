// 设置窗口相关的图标.
// 这个例子必须得go build后运行exe才能看出效果, go run是看不出效果的.
// goland是build后再运行的, 可以直接看出效果.
//
//go:generate goversioninfo
package main

import (
	_ "embed"
	"fmt"
	"syscall"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed icon1.ico
	icon1Data []byte
	//go:embed icon2.ico
	icon2Data []byte
)

func main() {
	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	img1 := imagex.NewByMemAdaptive(icon1Data, 0, 0, 0, 0)
	img2 := imagex.NewByMemAdaptive(icon2Data, 0, 0, 0, 0)

	// 设置程序默认窗口图标, 所有未设置图标的窗口, 都将使用此默认图标
	a.SetWindowIcon(img1.Handle)

	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xcgui window", 0, xcc.Window_Style_Default)

	// 单独设置指定窗口图标
	w.SetIcon(img2.Handle)

	// 从syso资源文件中加载图标.
	// 这个例子必须得go build后运行exe才能看出效果, go run是看不出效果的. go build时syso文件会被嵌入程序.
	// https://mcn1fno5w69l.feishu.cn/wiki/KvbpwTHkCibl6mkT8tXcWgD2nbk
	// 如何生成syso文件先看上面的文档.
	// 在versioninfo.json中填的有 "IconPath": "icon1.ico,icon2.ico",
	// 可以填多个图标, 用英文逗号隔开, 第一个图标将作为程序图标.
	// 下面代码里的1和8是图标组索引, 把编译后的程序拖进ResHacker.exe可以看到所有的图标组索引, 目前我也只想到这个办法, 或许有其他办法可以直接查看syso的.
	hInst := wapi.GetModuleHandleW("")
	hIcon1 := wapi.LoadImageW(hInst, uintptr(1), wapi.IMAGE_ICON, 0, 0, wapi.LR_SHARED|wapi.LR_DEFAULTSIZE)
	hIcon2 := wapi.LoadImageW(hInst, uintptr(8), wapi.IMAGE_ICON, 0, 0, wapi.LR_SHARED|wapi.LR_DEFAULTSIZE)
	fmt.Println(hIcon1, hIcon2, syscall.GetLastError())

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
