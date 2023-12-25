// 复选按钮
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
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "复选按钮", 0, xcc.Window_Style_Default)

	// 创建按钮
	Check1 := widget.NewButton(10, 35, 70, 30, "Check1", w.Handle)
	Check2 := widget.NewButton(10, 75, 70, 30, "Check2", w.Handle)
	Check3 := widget.NewButton(10, 115, 70, 30, "Check3", w.Handle)
	// 设置按钮类型
	Check1.SetTypeEx(xcc.Button_Type_Check)
	Check2.SetTypeEx(xcc.Button_Type_Check)
	Check3.SetTypeEx(xcc.Button_Type_Check)

	// 设置选中
	Check1.SetCheck(true)

	// 注册事件_按钮被选中
	Check1.Event_BUTTON_CHECK1(btn_check)
	Check2.Event_BUTTON_CHECK1(btn_check)
	Check3.Event_BUTTON_CHECK1(btn_check)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 事件_按钮被选中
func btn_check(hEle int, bCheck bool, pbHandled *bool) int {
	if bCheck {
		fmt.Println(xc.XBtn_GetText(hEle), "Selected")
	} else {
		fmt.Println(xc.XBtn_GetText(hEle), "Unselected")
	}
	return 0
}
