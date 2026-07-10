// 设置窗口相关的图标和程序图标
package main

// 在运行前, 可选择阅读<在vscode中实现编译后运行>: https://mcn1fno5w69l.feishu.cn/wiki/I0d9wpUVai4GxikqLxwcftYunWh

import (
	"fmt"
	"syscall"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 从 syso 文件中加载图标.
	// 这个例子必须得 go build 后运行 exe 才能看出效果, go run 是看不出效果的. go build 时 syso 文件会被嵌入程序.
	// 如何生成 syso 文件: https://github.com/tc-hib/go-winres
	hInst := wapi.GetModuleHandleW("")
	hIcon1 := wapi.LoadImageW(hInst, common.StrPtr("ICON1"), wapi.IMAGE_ICON, 0, 0, wapi.LR_SHARED|wapi.LR_DEFAULTSIZE)
	hIcon2 := wapi.LoadImageW(hInst, common.StrPtr("ICON2"), wapi.IMAGE_ICON, 0, 0, wapi.LR_SHARED|wapi.LR_DEFAULTSIZE)

	if hIcon1 == 0 || hIcon2 == 0 {
		fmt.Println("GetLastError:", syscall.GetLastError())
		panic("图标加载失败, 需要先编译再运行! ")
	}

	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 从 HICON 创建 Image 对象
	img1 := imagex.NewByHICON(hIcon1)
	img2 := imagex.NewByHICON(hIcon2)

	// 设置程序默认窗口标题栏上的图标, 所有未设置图标的窗口, 都将使用此默认图标
	a.SetWindowIcon(img1.Handle)

	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xcgui window", 0, xcc.Window_Style_Default)
	widget.NewShapeText(20, 50, 300, 100, "程序运行效果: \n任务栏图标是齿轮\n窗口图标是炸弹\n任务栏预览窗口的图标是炸弹", w.Handle)

	// 单独设置指定窗口标题栏上的图标
	w.SetIcon(img2.Handle)

	// 通过动态设置可以给不同窗口设置不同图标
	// 设置小图标 (任务栏预览左上角图标)
	w.SetSmallIcon(hIcon2)

	// 设置大图标 (任务栏图标和 Alt+Tab页面中的图标).
	// 像 QQ 打开独立的聊天窗口, 可以设置这个以显示头像在任务栏.
	w.SetBigIcon(hIcon1)

	w.Show(true)
	a.Run()
	a.Exit()
}
