// 设置鼠标光标
package main

// 在运行前, 可选择阅读<在vscode中实现编译后运行>: https://mcn1fno5w69l.feishu.cn/wiki/I0d9wpUVai4GxikqLxwcftYunWh

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 500, 400, "设置鼠标光标", 0, xcc.Window_Style_Default)

	// 从syso中的游标加载, 设置窗口鼠标光标
	// 这个例子必须得 go build 后运行 exe 才能看出效果, go run 是看不出效果的. go build 时 syso 文件会被嵌入程序.
	// 如何生成 syso 文件: https://github.com/tc-hib/go-winres
	hCur := wapi.LoadImageW(wapi.GetModuleHandleW(""), common.StrPtr("ARROW"), wapi.IMAGE_CURSOR, 0, 0, wapi.LR_SHARED|wapi.LR_DEFAULTSIZE)
	fmt.Println("窗口鼠标光标句柄:", hCur)
	if hCur != 0 {
		w.SetCursor(hCur)
	}

	// 设置按钮鼠标光标
	rand.Seed(time.Now().UnixNano())
	btn := widget.NewButton(50, 50, 150, 40, "改变按钮鼠标光标", w.Handle)

	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		// 加载系统预定义的游标, 可使用 wapi.IDC_ 系列常量
		// https://learn.microsoft.com/zh-cn/windows/win32/menurc/about-cursors
		var idc int
		if rand.Intn(2) == 1 {
			idc = wapi.IDC_SIZENWSE + rand.Intn(5)
		} else {
			idc = wapi.IDC_ARROW + rand.Intn(5)
		}

		hCur := wapi.LoadImageW(0, uintptr(idc), wapi.IMAGE_CURSOR, 0, 0, wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
		fmt.Println("新的按钮鼠标光标句柄, 点击按钮后移动一下鼠标,", hCur)

		if hCur != 0 {
			btn.SetCursor(hCur) // 设置元素鼠标光标
		}
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
