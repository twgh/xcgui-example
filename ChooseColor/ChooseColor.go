// 调用 wapi 选择颜色
package main

import (
	"fmt"
	"unsafe"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var w *window.Window
var custColors [16]uint32 // 保存自定义颜色的数组

func main() {
	a := app.New(true)
	w = window.New(0, 0, 430, 300, "选择颜色", 0, xcc.Window_Style_Default)

	widget.NewButton(20, 40, 100, 30, "选择颜色", w.Handle).Event_BnClick(onBnClick)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func onBnClick(pbHandled *bool) int {
	ExampleChooseColorW()
	return 0
}

func ExampleChooseColorW() {
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
	fmt.Println(ret)
	fmt.Println(cc.RgbResult) // rgb颜色
	fmt.Println(custColors)   // 如果你添加了自定义颜色, 会保存在这个数组里面, 然后只要这个数组还在, 再次打开选择颜色界面时, 之前添加的自定义颜色还会存在

	// 设置窗口标题颜色
	w.SetTitleColor(xc.ABGR2(int(cc.RgbResult), 255))
	w.Redraw(true)
}
