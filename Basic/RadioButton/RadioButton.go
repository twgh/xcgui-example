// 单选按钮
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 360, 300, "单选按钮", 0, xcc.Window_Style_Default)

	// 创建形状文本, 显示当前选中的选项
	stCur := widget.NewShapeText(110, 40, 200, 24, "当前选中: Radio1", w.Handle)

	// 创建按钮
	Radio1 := widget.NewButton(10, 35, 70, 30, "Radio1", w.Handle)
	Radio2 := widget.NewButton(10, 75, 70, 30, "Radio2", w.Handle)
	Radio3 := widget.NewButton(10, 115, 70, 30, "Radio3", w.Handle)

	// 设置按钮类型
	Radio1.SetTypeEx(xcc.Button_Type_Radio)
	Radio2.SetTypeEx(xcc.Button_Type_Radio)
	Radio3.SetTypeEx(xcc.Button_Type_Radio)

	// 设置背景透明
	Radio1.EnableBkTransparent(true)
	Radio2.EnableBkTransparent(true)
	Radio3.EnableBkTransparent(true)

	// 设置分组id
	Radio1.SetGroupID(1)
	Radio2.SetGroupID(1)
	Radio3.SetGroupID(1)

	// 保存到一个切片, 方便遍历
	radios := []*widget.Button{Radio1, Radio2, Radio3}

	// 更新当前选中显示
	updateSelText := func(name string) {
		stCur.SetText(fmt.Sprintf("当前选中: %s", name))
		stCur.Redraw() // 修改显示后必须重绘
	}

	// 设置选中, 会触发选中事件
	Radio1.SetCheck(true)

	// 注册事件_按钮被选中
	for _, r := range radios {
		r.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
			if bCheck {
				// 只有选中时才更新显示
				updateSelText(xc.XBtn_GetText(hEle))
			}
			return 0
		})
	}

	// 创建提交按钮
	btnSubmit := widget.NewButton(10, 165, 70, 30, "提交", w.Handle)
	btnSubmit.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		// 遍历分组内的单选按钮, 用 IsCheck() 判断哪个处于选中状态
		for _, r := range radios {
			if r.IsCheck() {
				w.MessageBox("提交结果", fmt.Sprintf("你选择了: %s", r.GetText()), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
				return 0
			}
		}
		w.MessageBox("提交结果", "你还没有选择任何项", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
