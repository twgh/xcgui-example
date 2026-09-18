// 多窗口例子
package main

// 窗口1是登录窗口, 登录后销毁登录窗口载入主窗口(窗口2),
// 主窗口点"注销"再回到登录窗口, 形成完整闭环.
//
// 关键点: 在按钮事件里销毁当前窗口并载入新窗口时, 必须 *pbHandled = true 拦截事件

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	loadLoginWindow()

	a.Run()
	a.Exit()
}

// 登录窗口: 用户名、密码两个输入框 + 登录按钮
func loadLoginWindow() {
	w1 := window.New(0, 0, 300, 220, "登录", 0, xcc.Window_Style_Default)

	// 用户名输入框
	editUser := widget.NewEdit(90, 50, 170, 26, w1.Handle)
	editUser.SetDefaultText("用户名")
	// 密码输入框, 启用密码模式, 显示为圆点
	editPwd := widget.NewEdit(90, 90, 170, 26, w1.Handle)
	editPwd.EnablePassword(true)
	editPwd.SetDefaultText("密码")

	// 登录按钮
	btn := widget.NewButton(90, 140, 170, 32, "登录", w1.Handle)
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		// 不做真实校验, 简单判断: 用户名为空就弹框提示, 不进入主窗口
		name := editUser.GetText()
		if len(name) == 0 {
			w1.MessageBox("提示", "用户名不能为空", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
			return 0
		}

		*pbHandled = true // 拦截事件, 载入新窗口时这是必要的
		w1.CloseWindow()  // 销毁登录窗口
		loadMainWindow(name)
		return 0
	})

	w1.Show(true)
}

// 主窗口: 显示欢迎文本 + 注销按钮
func loadMainWindow(name string) {
	w2 := window.New(0, 0, 300, 220, "主窗口", 0, xcc.Window_Style_Default)

	// 形状文本显示欢迎信息
	shape := widget.NewShapeText(50, 70, 200, 24, "欢迎 "+name+" 登录!", w2.Handle)
	shape.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)

	// 注销按钮: 销毁主窗口回到登录窗口, 形成闭环
	btn := widget.NewButton(90, 130, 120, 32, "注销", w2.Handle)
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		*pbHandled = true // 拦截事件, 载入新窗口时这是必要的
		w2.CloseWindow()
		loadLoginWindow()
		return 0
	})

	w2.Show(true)
}
