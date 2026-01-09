// 复选按钮(多选按钮/选择框)
package main

import (
	"fmt"
	"sync"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 430, 300, "复选按钮(多选按钮/选择框)", 0, xcc.Window_Style_Default)

	// 创建按钮
	Check1 := widget.NewButton(20, 35, 70, 30, "Check1", w.Handle)
	Check2 := widget.NewButton(20, 75, 70, 30, "Check2", w.Handle)

	// 设置按钮类型为多选按钮
	Check1.SetTypeEx(xcc.Button_Type_Check)
	Check2.SetTypeEx(xcc.Button_Type_Check)

	// 启用背景透明
	Check1.EnableBkTransparent(true)
	Check2.EnableBkTransparent(true)

	// 设置Check1选中
	Check1.SetCheck(true)
	// 注册事件_按钮被选中
	Check1.AddEvent_Button_Check(onBtnCheck)
	Check2.AddEvent_Button_Check(onBtnCheck)

	// 美化Check2, 涉及背景管理器和背景对象的使用,
	// 感觉理解起来难的话, 可打开设计器给多选按钮通过背景编辑器添加背景图片试试, 结合可视化界面来理解代码.
	// 反正没必要抵触设计器, 可视化设计熟练后很省事, xml也方便复用.
	{
		// 获取按钮的背景管理器
		bkm := Check2.GetBkManagerObj()
		// 背景_清空
		bkm.Clear()

		// 图片宽高
		var imgWidth int32 = 18
		// 加载svg图片
		imgUnSelect := imagex.NewBySvg(svg.NewByString(svg_unselect).SetSize(imgWidth, imgWidth).Handle)
		imgSelect := imagex.NewBySvg(svg.NewByString(svg_select).SetSize(imgWidth, imgWidth).Handle)

		// 给背景管理器按钮不同状态添加图片背景对象
		bkm.AddImage(xcc.Button_State_Flag_Check_No, imgUnSelect.Handle, 1)
		bkm.AddImage(xcc.Button_State_Flag_Check, imgSelect.Handle, 2)

		for i := int32(1); i <= 2; i++ {
			// 获取上面添加的图片背景对象根据设置的id
			_bkobj := bkm.GetObjectObj(i)
			// 设置对齐方式为左对齐
			_bkobj.SetAlign(xcc.BkObject_Align_Flag_Left)
			// 设置外边距
			_bkobj.SetMargin(0, 0, imgWidth, 0)
		}
	}

	// 创建开关按钮
	SwitchButton := NewSwitchButton(20, 115, 40, 40, w.Handle)
	SwitchButton.AddEvent_Button_Check(onBtnCheck)

	SwitchButton2 := NewSwitchButton(20, 155, 60, 60, w.Handle)
	SwitchButton2.SetCheck(true)
	SwitchButton2.AddEvent_Button_Check(onBtnCheck)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

var (
	imgOff *imagex.Image
	imgOn  *imagex.Image
	once   sync.Once
)

// NewSwitchButton 创建开关按钮. cx 与 cy 相同时效果最好.
func NewSwitchButton(x, y, cx, cy int32, hParent int) *widget.Button {
	btn := widget.NewButton(x, y, cx, cy, "", hParent)
	// 设置按钮类型为多选按钮
	btn.SetTypeEx(xcc.Button_Type_Check)
	// 启用背景透明
	btn.EnableBkTransparent(true)
	// 加载svg图片, 只加载一次
	once.Do(func() {
		imgOff = imagex.NewBySvgString(svg_off)
		imgOn = imagex.NewBySvgString(svg_on)
		// 设置图片绘制类型为拉伸
		imgOff.SetDrawType(xcc.Image_Draw_Type_Stretch)
		imgOn.SetDrawType(xcc.Image_Draw_Type_Stretch)
	})
	// 设置按钮不同状态的背景图片
	btn.AddBkImage(xcc.Button_State_Flag_Check_No, imgOff.Handle)
	btn.AddBkImage(xcc.Button_State_Flag_Check, imgOn.Handle)
	return btn
}

const (
	svg_off = `<svg t="1745761805273" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1075" width="32" height="32"><path d="M2 512c0-141 114-255 255-255h510c141 0 255 114 255 255S908 767 767 767H257C116 767 2 653 2 512z m255 208.8c115.5 0 208.8-93.2 208.8-208.8S372.5 303.2 257 303.2c-115.5 0-208.8 93.2-208.8 208.8S141.5 720.8 257 720.8z" fill="#DCDFE6" p-id="1076"></path></svg>`
	svg_on  = `<svg t="1745761841891" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1445" width="32" height="32"><path d="M2 512c0-141 114-255 255-255h510c141 0 255 114 255 255S908 767 767 767H257C116 767 2 653 2 512z m765 208.8c115.5 0 208.8-93.2 208.8-208.8S882.5 303.2 767 303.2c-115.5 0-208.8 93.2-208.8 208.8S651.5 720.8 767 720.8z" fill="#1890FF" p-id="1446"></path></svg>`
)

// 事件_按钮被选中
func onBtnCheck(hEle int, bCheck bool, pbHandled *bool) int {
	if bCheck {
		fmt.Println(xc.XBtn_GetText(hEle), "selected")
	} else {
		fmt.Println(xc.XBtn_GetText(hEle), "Unselected")
	}
	return 0
}

const (
	svg_unselect = `<svg t="1745750357126" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4423" width="32" height="32"><path d="M896 64H128c-35.296 0-64 28.704-64 64v768c0 35.296 28.704 64 64 64h768c35.296 0 64-28.704 64-64V128c0-35.296-28.704-64-64-64zM128 896V128h768l0.064 768H128z" fill="#1296db" p-id="4424"></path></svg>`
	svg_select   = `<svg t="1745750405994" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4997" width="32" height="32"><path d="M896 64H128c-35.296 0-64 28.704-64 64v768c0 35.296 28.704 64 64 64h768c35.296 0 64-28.704 64-64V128c0-35.296-28.704-64-64-64zM128 896V128h768l0.064 768H128z" p-id="4998" fill="#1296db"></path><path d="M744.64 308.032l-310.368 331.296-157.696-132.48a32 32 0 0 0-41.184 49.024l180.896 152a32 32 0 0 0 43.936-2.624L791.36 351.776a32 32 0 1 0-46.72-43.744z" p-id="4999" fill="#1296db"></path></svg>`
)
