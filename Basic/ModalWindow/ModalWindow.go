// 模态窗口
package main

// 用法要点:
//  1. window.NewModalWindow(宽, 高, 标题, 父窗口HWND, xcc.Window_Style_Modal) 创建模态窗口,
//     模态期间父窗口无法操作; 模态窗口关闭时会自动销毁资源句柄.
//  2. mw.DoModal() 会阻塞, 直到模态窗口结束才返回, 返回值表示结束方式:
//     xcc.MessageBox_Flag_Ok / xcc.MessageBox_Flag_Cancel / xcc.MessageBox_Flag_Other.
//  3. 在按钮事件里调用 mw.EndModal(nResult) 结束模态窗口, nResult 就是 DoModal() 的返回值.
//  4. DoModal() 返回后模态窗口已销毁, 但结果可以通过 Go 闭包变量带回(如 Edit 的内容),
//     显示回主窗口的界面元素上.

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
	// 创建窗口
	w := window.New(0, 0, 430, 300, "模态窗口", 0, xcc.Window_Style_Default)

	// 创建形状文本, 用来显示模态窗口返回的结果
	shape := widget.NewShapeText(30, 100, 370, 24, "结果: (还没有)", w.Handle)

	// 创建按钮_模态窗口
	btn := widget.NewButton(30, 50, 100, 30, "ModalWindow", w.Handle)
	// 给按钮绑定事件
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		// 创建模态窗口, 模态期间父窗口无法操作
		mw := window.NewModalWindow(360, 170, "输入对话框", w.GetHWND(), xcc.Window_Style_Modal)

		// 模态窗口里放一个输入框, 坐标相对于整个窗口, 注意避开标题栏
		input := widget.NewEdit(30, 45, 300, 26, mw.Handle)
		input.SetDefaultText("请输入点什么...")
		input.SetFocus()

		// 结果通过闭包变量带回, 确定时记录输入内容
		var result string
		var confirmed bool

		// 确定按钮: 记录输入内容, 用 EndModal 结束模态, 参数会成为 DoModal 的返回值
		btnOk := widget.NewButton(30, 95, 90, 30, "确定", mw.Handle)
		btnOk.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			result = input.GetText()
			confirmed = true
			mw.EndModal(xcc.MessageBox_Flag_Ok)
			return 0
		})

		// 取消按钮: 直接结束模态
		btnCancel := widget.NewButton(140, 95, 90, 30, "取消", mw.Handle)
		btnCancel.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			mw.EndModal(xcc.MessageBox_Flag_Cancel)
			return 0
		})

		// DoModal 阻塞显示, 直到模态窗口结束, 返回结束方式
		nRet := mw.DoModal()

		// DoModal 返回后模态窗口已销毁, 通过闭包变量拿到结果显示回主窗口
		if confirmed && nRet == xcc.MessageBox_Flag_Ok {
			shape.SetText("结果: " + result)
		} else {
			shape.SetText("结果: (已取消)")
		}
		// 修改元素显示后必须重绘
		shape.Redraw()
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
