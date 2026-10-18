// 拦截dll内部错误弹窗
package main

// dll 内部错误弹窗通常是因为传入了错误类型的句柄(把按钮句柄传给输入框API), 无效的句柄(元素已不存在)等.

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 拦截 dll 内部错误弹窗, 让 go 侧能拿到报错的 API 名和调用栈
	if err := xc.InstallErrorHook(); err != nil {
		println("安装错误钩子失败:", err.Error())
	}

	w := window.New(0, 0, 400, 240, "测试dll弹窗报错", 0, xcc.Window_Style_Default)

	widget.NewButton(50, 50, 100, 30, "测试报错", w.Handle).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		xc.XEle_Redraw(99999) //故意传错句柄
		return 0
	})

	w.Show()
	a.Run()
	a.Exit()
}
