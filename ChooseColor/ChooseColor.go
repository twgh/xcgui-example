// 调用 wapi 选择颜色
package main

import (
	"fmt"
	"unsafe"

	"github.com/twgh/xcgui/xc"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	w          *window.Window
	custColors [16]uint32 // 保存自定义颜色的数组
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w = window.New(0, 0, 430, 300, "选择颜色", 0, xcc.Window_Style_Default)

	widget.NewButton(20, 40, 100, 30, "选择颜色", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		// 或调用封装好的函数: wutil.ChooseColor(w.Handle)
		ExampleChooseColor()
		return 0
	})

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func ExampleChooseColor() {
	cc := wapi.ChooseColor{
		LStructSize:    36,
		HwndOwner:      w.GetHWND(),
		HInstance:      0,
		RgbResult:      0,
		LpCustColors:   &custColors[0],
		Flags:          wapi.CC_FULLOPEN, // 默认打开自定义颜色
		LCustData:      0,
		LpfnHook:       0,
		LpTemplateName: 0,
	}
	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	ret := wapi.ChooseColorW(&cc)
	fmt.Println("选择颜色:", ret)
	if !ret {
		return
	}

	rgb := cc.RgbResult
	rgba := xc.RGB2RGBA(int(rgb), 255)
	fmt.Println("rgb颜色:", rgb)
	fmt.Println("炫彩使用的颜色:", rgba)
	fmt.Println("自定义颜色:", custColors) // 如果你添加了自定义颜色, 会保存在这个数组里面, 然后只要这个数组还在, 再次打开选择颜色界面时, 之前添加的自定义颜色还会存在

	// 设置窗口背景颜色
	w.AddBkFill(xcc.Window_State_Flag_Leave, rgba)
	w.Redraw(true)
}
