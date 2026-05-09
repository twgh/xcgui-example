// 动画特效展示
package main

import (
	_ "embed"
	"fmt"
	"math/rand"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/svg"
	"github.com/twgh/xcgui/wapi"

	"github.com/twgh/xcgui/ani"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	w *window.Window

	list_svg       []int
	list_animation []int
	list_xcgui     []int
	list_object    []interface{}

	m_hLayout1 int
	m_hLayout2 int
	m_hLayout3 int

	m_hSvg int
)

var (
	//go:embed svg/公益.svg
	svg1 string
	//go:embed svg/时间戳.svg
	svg2 string
	//go:embed svg/技术服务.svg
	svg3 string
	//go:embed svg/底层架构.svg
	svg4 string
	//go:embed svg/查验.svg
	svg5 string
	//go:embed svg/接口配置.svg
	svg6 string
	//go:embed svg/淘公仔文字.svg
	svg7 string
	//go:embed svg/成功.svg
	svgSuccess string
	//go:embed svg/消息.svg
	svgMessage string
	//go:embed svg/警告.svg
	svgWarning string
	//go:embed svg/错误.svg
	svgError string

	//go:embed image/img-1.jpg
	img1 []byte
	//go:embed image/img-2.jpg
	img2 []byte
	//go:embed image/img-3.jpg
	img3 []byte
	//go:embed image/img-4.jpg
	img4 []byte
	//go:embed image/img-5.jpg
	img5 []byte
	//go:embed image/img-6.jpg
	img6 []byte

	//go:embed image/progressBar.png
	imgProgressBar []byte
	//go:embed image/sliderBar.png
	imgSliderBar []byte

	svg11 = `<svg x="0" y="0" width="25" height="25" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#ee6362" /></svg>`
	svg12 = `<svg x="0" y="0" width="25" height="25" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#2cb0b2" /></svg>`
	svg13 = `<svg x="0" y="0" width="20" height="20" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#f00" /></svg>`
	svg14 = `<svg x="0" y="0" width="15" height="15" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#f00" /></svg>`
	svg15 = `<svg viewBox="0 0 200 200"><circle cx="100" cy="100" r="100" fill="#ff0" /></svg>`
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 设置UI的最小重绘频率
	a.SetPaintFrequency(10)
	// a.ShowLayoutFrame(true).ShowSvgFrame(true)
	// 创建窗口
	w = window.New(0, 0, 1020, 650, "炫彩界面库-动画特效-SVG特效", 0, xcc.Window_Style_Default)

	// 创建按钮, 注册按钮单击事件
	var top int32 = 35
	var left int32 = 10
	CreateButtonRadio(left, &top, "1.下落 缩放 缓动").AddEvent_BnClick(OnBtnClick1)
	CreateButtonRadio(left, &top, "2.下落 呼吸SVG").AddEvent_BnClick(OnBtnClick2)
	CreateButtonRadio(left, &top, "3.呼吸SVG").AddEvent_BnClick(OnBtnClick3)
	CreateButtonRadio(left, &top, "4.不透明度SVG").AddEvent_BnClick(OnBtnClick4)
	CreateButtonRadio(left, &top, "5.移动SVG").AddEvent_BnClick(OnBtnClick5)
	CreateButtonRadio(left, &top, "6.形状文本").AddEvent_BnClick(OnBtnClick6)
	CreateButtonRadio(left, &top, "7.按钮").AddEvent_BnClick(OnBtnClick7)
	CreateButtonRadio(left, &top, "8.布局焦点展开").AddEvent_BnClick(OnBtnClick8)
	CreateButtonRadio(left, &top, "9.图片切换").AddEvent_BnClick(OnBtnClick9)
	CreateButtonRadio(left, &top, "10.图片切换2").AddEvent_BnClick(OnBtnClick10)
	CreateButtonRadio(left, &top, "11.进度 等待").AddEvent_BnClick(OnBtnClick11)
	CreateButtonRadio(left, &top, "12.旋转 移动").AddEvent_BnClick(OnBtnClick12)
	CreateButtonRadio(left, &top, "13.旋转 摇摆").AddEvent_BnClick(OnBtnClick13)
	CreateButtonRadio(left, &top, "14.旋转 移动 缩放").AddEvent_BnClick(OnBtnClick14)
	CreateButtonRadio(left, &top, "15.旋转 开合效果").AddEvent_BnClick(OnBtnClick15)
	CreateButtonRadio(left, &top, "16.颜色渐变").AddEvent_BnClick(OnBtnClick16)
	CreateButtonRadio(left, &top, "17.缩放 位置").AddEvent_BnClick(OnBtnClick17)
	CreateButtonRadio(left, &top, "18.按钮 宽度").AddEvent_BnClick(OnBtnClick18)

	top = 35
	left = 900
	CreateButtonRadio(left, &top, "19.窗口特效").AddEvent_BnClick(OnBtnClick19)
	CreateButtonRadio(left, &top, "20.遮盖弹窗").AddEvent_BnClick(OnBtnClick20)
	CreateButtonRadio(left, &top, "21.通知消息").AddEvent_BnClick(OnBtnClick21)
	CreateButtonRadio(left, &top, "22.进度条").AddEvent_BnClick(OnBtnClick22)
	CreateButtonRadio(left, &top, "23.焦点追踪").AddEvent_BnClick(OnBtnClick23)
	CreateButtonRadio(left, &top, "24.页面切换 滑动").AddEvent_BnClick(OnBtnClick24)
	CreateButtonRadio(left, &top, "25.折叠面板").AddEvent_BnClick(OnBtnClick25)
	// todo: 翻译剩下的动画
	// CreateButtonRadio(left, &top, "26.图片轮播").AddEvent_BnClick(OnBtnClick26)
	// CreateButtonRadio(left, &top, "27.背景管理器").AddEvent_BnClick(OnBtnClick27)

	w.AddEvent_Paint(OnWndDrawWindow)
	w.AddEvent_Destroy(func(hWindow int, pbHandled *bool) int {
		ReleaseAnimation()
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 创建单选按钮
func CreateButtonRadio(left int32, top *int32, name string) *widget.Button {
	btn := widget.NewButton(left, *top, 110, 30, name, w.Handle)
	btn.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter)
	btn.SetTypeEx(xcc.Button_Type_Radio)
	btn.SetGroupID(1)
	*top += 29
	return btn
}

// 创建按钮
func CreateButton(left, top, width, height int32, name string) *widget.Button {
	btn := widget.NewButton(left, top, width, height, name, w.Handle)
	btn.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter)
	btn.SetPadding(10, 0, 0, 0)
	return btn
}

// 释放资源
func ReleaseAnimation() {
	for i := 0; i < len(list_object); i++ {
		switch obj := list_object[i].(type) {
		case *CFocusTraceButton_Line:
			obj.Release()
		case *CFocusTraceEdit_Border:
			obj.Release()
		case *CFocusTraceEdit_Line:
			obj.Release()
		case *CExpandGroup:
			obj.Release()
		}
	}

	for _, v := range list_animation {
		xc.XAnima_Release(v, true)
	}

	for _, v := range list_svg {
		xc.XSvg_Release(v)
	}

	for _, v := range list_xcgui {
		t := xc.XObj_GetTypeBase(v)
		switch t {
		case xcc.XC_ELE:
			xc.XEle_Destroy(v)
		case xcc.XC_SHAPE:
			xc.XShape_Destroy(v)
		case xcc.XC_SVG:
			xc.XSvg_Release(v)
		}
	}

	list_animation = list_animation[:0]
	list_svg = list_svg[:0]
	list_xcgui = list_xcgui[:0]
	list_object = list_object[:0]
}

func ReleaseObject(hObject int) {
	xc_type := xc.XObj_GetTypeBase(hObject)
	switch xc_type {
	case xcc.XC_ELE:
		xc.XEle_Destroy(hObject)
	case xcc.XC_SHAPE:
		xc.XShape_Destroy(hObject)
	case xcc.XC_SVG:
		xc.XSvg_Release(hObject)
	}
}

// 窗口绘制消息.
func OnWndDrawWindow(hWindow int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	w.DrawWindow(hDraw)

	if m_hSvg != 0 {
		xc.XDraw_DrawSvgSrc(hDraw, m_hSvg)
	}

	for _, v := range list_svg {
		xc.XDraw_DrawSvgSrc(hDraw, v)
	}
	return 0
}

// 1.下落 缩放 缓动
func OnBtnClick1(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 130
	var top int32 = 22

	// 加载svg图片
	list_svg = append(list_svg,
		xc.XSvg_LoadString(svg1),
		xc.XSvg_LoadString(svg2),
		xc.XSvg_LoadString(svg3),
		xc.XSvg_LoadString(svg4),
		xc.XSvg_LoadString(svg5),
		xc.XSvg_LoadString(svg6),
	)

	// 创建动画组
	hGroup := xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)
	xc.XAnima_Run(hGroup, w.Handle)

	for i, hSvg := range list_svg {
		// 设置svg图片大小和位置
		xc.XSvg_SetSize(hSvg, 100, 100)
		xc.XSvg_SetPosition(hSvg, left, top)

		// 创建动画序列
		hAnimation := xc.XAnima_Create(hSvg, 0)
		// 将动画序列添加到动画组中
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)

		xc.XAnima_Move(hAnimation, 500, float32(left), 22, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
		xc.XAnima_Delay(hAnimation, 500)

		xc.XAnima_Delay(hAnimation, 100*float32(i))
		xc.XAnima_Alpha(hAnimation, 500, 0, 1, 0, false)

		xc.XAnima_Delay(hAnimation, 500)

		xc.XAnima_Alpha(hAnimation, 500, 255, 1, 0, false)
		xc.XAnima_Delay(hAnimation, 1000)

		xc.XAnima_Move(hAnimation, 2000, float32(left), 500, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
		xc.XAnima_Delay(hAnimation, 1000)
		left += 130
		{
			hAnimation = xc.XAnima_Create(hSvg, 0)
			xc.XAnima_Delay(hAnimation, 6000+float32(i)*200)
			xc.XAnima_Scale(hAnimation, 1200, 2, 2, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, true)

			xc.XAnimaGroup_AddItem(hGroup, hAnimation)
		}
	}
	return 0
}

// 2.下落 呼吸SVG
func OnBtnClick2(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 450
	var top int32 = 22

	// 加载svg图片
	list_svg = append(list_svg, xc.XSvg_LoadString(svg1))
	// 设置svg图片大小和位置
	xc.XSvg_SetSize(list_svg[0], 100, 100)
	xc.XSvg_SetPosition(list_svg[0], left, top)

	// 创建动画组
	group := ani.NewAnimaGroup(0)
	list_animation = append(list_animation, group.Handle)
	group.Run(w.Handle)

	// 下落
	ani1 := ani.NewAnima(list_svg[0], 0)
	group.AddItem(ani1.Handle)
	{
		ani1.Move(2000, float32(left), 500, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
		// 停留
		ani1.Delay(2000)
		// 返回顶部
		ani1.Move(500, float32(left), 22, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
	}

	// 缩放
	ani2 := ani.NewAnima(list_svg[0], 1)
	group.AddItem(ani2.Handle)
	{
		ani2.Delay(2000)
		ani2.Scale(1000, 2, 2, 0, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, true)
	}

	/* 以下是纯函数方式实现
		// 创建动画组
	   	hGroup := xc.XAnimaGroup_Create(0)
	   	list_animation = append(list_animation, hGroup)
	   	xc.XAnima_Run(hGroup, w.Handle)

	   	// 下落
	   	hAnimation := xc.XAnima_Create(list_svg[0], 0)
	   	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	   	xc.XAnima_Move(hAnimation, 2000, float32(left), 500, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)

	   	// 停留
	   	xc.XAnima_Delay(hAnimation, 2000)

	   	// 返回顶部
	   	xc.XAnima_Move(hAnimation, 500, float32(left), 22, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)

	   	// 缩放
	   	hAnimation = xc.XAnima_Create(list_svg[0], 1)
	   	xc.XAnimaGroup_AddItem(hGroup, hAnimation)

	   	xc.XAnima_Delay(hAnimation, 2000)
	   	xc.XAnima_Scale(hAnimation, 1000, 2, 2, 0, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, true)
	*/
	return 0
}

// 3.呼吸SVG
func OnBtnClick3(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 300
	var top int32 = 150

	// 加载svg图片
	list_svg = append(list_svg, xc.XSvg_LoadString(svg1))
	// 设置svg图片大小和位置
	xc.XSvg_SetSize(list_svg[0], 300, 300)
	xc.XSvg_SetPosition(list_svg[0], left, top)

	// 创建动画序列
	ani1 := ani.NewAnima(list_svg[0], 1)
	list_animation = append(list_animation, ani1.Handle)

	// 缩放
	ani1.Scale(1500, 2, 2, 0, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, true)
	ani1.Run(w.Handle)

	/* 以下是纯函数方式实现
		// 创建动画序列
	   	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	   	list_animation = append(list_animation, hAnimation)

	   	// 缩放
	   	xc.XAnima_Scale(hAnimation, 1500, 2, 2, 0, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, true)
	   	xc.XAnima_Run(hAnimation, w.Handle)
	*/
	return 0
}

// 4.不透明度SVG
func OnBtnClick4(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 200
	var top int32 = 30

	// 加载svg图片
	list_svg = append(list_svg,
		xc.XSvg_LoadString(svg1),
		xc.XSvg_LoadString(svg1),
		xc.XSvg_LoadString(svg1),
	)

	// 设置svg图片大小和位置
	for i, hSvg := range list_svg {
		xc.XSvg_SetSize(hSvg, 100, 100)
		xc.XSvg_SetPosition(hSvg, left+int32(i)*100, top)
	}

	// 创建动画序列
	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_AlphaEx(hAnimation, 3000, 0, 255, 1, 0, false)
	xc.XAnima_Run(hAnimation, w.Handle)

	hAnimation = xc.XAnima_Create(list_svg[1], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Alpha(hAnimation, 3000, 0, 1, 0, true)
	xc.XAnima_Run(hAnimation, w.Handle)

	hAnimation = xc.XAnima_Create(list_svg[2], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Alpha(hAnimation, 3000, 0, 0, 0, true)
	xc.XAnima_Run(hAnimation, w.Handle)

	{
		top = 100
		m_hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, m_hSvg)
		xc.XSvg_SetPosition(m_hSvg, left, top)

		hAnimation = xc.XAnima_Create(m_hSvg, 0)
		list_animation = append(list_animation, hAnimation)

		xc.XAnima_Alpha(hAnimation, 3000, 0, 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		m_hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, m_hSvg)
		xc.XSvg_SetPosition(m_hSvg, left, top)

		hAnimation = xc.XAnima_Create(m_hSvg, 0)
		list_animation = append(list_animation, hAnimation)

		xc.XAnima_AlphaEx(hAnimation, 3000, 255, 50, 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		m_hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, m_hSvg)
		xc.XSvg_SetPosition(m_hSvg, left, top)

		hAnimation = xc.XAnima_Create(m_hSvg, 0)
		list_animation = append(list_animation, hAnimation)

		xc.XAnima_AlphaEx(hAnimation, 3000, 50, 255, 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	return 0
}

// 5.移动SVG
func OnBtnClick5(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	var top int32 = 30

	// 加载svg图片
	list_svg = append(list_svg,
		xc.XSvg_LoadString(svg1),
		xc.XSvg_LoadString(svg2),
		xc.XSvg_LoadString(svg3),
	)

	// 设置svg图片大小和位置
	for i, hSvg := range list_svg {
		xc.XSvg_SetSize(hSvg, 100, 100)
		xc.XSvg_SetPosition(hSvg, left, top+int32(i)*100)
	}
	top = 22

	// 循环
	ani1 := ani.NewAnima(list_svg[0], 1)
	list_animation = append(list_animation, ani1.Handle)
	ani1.Run(w.Handle)
	ani1.Move(2000, 750, float32(top), 10, 0, true)
	top += 100

	// 一次, 往返
	ani2 := ani.NewAnima(list_svg[1], 1)
	list_animation = append(list_animation, ani2.Handle)
	ani2.Run(w.Handle)
	ani2.Move(2000, 750, float32(top), 1, 0, true)
	top += 100

	// 一次, 不往返
	ani3 := ani.NewAnima(list_svg[2], 1)
	list_animation = append(list_animation, ani3.Handle)
	ani3.Run(w.Handle)
	ani3.Move(2000, 750, float32(top), 1, 0, false)
	top += 100

	/* 以下是纯函数方式实现
	// 循环
	hAnimation := xc.XAnima_Create(list_svg[0], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 10, 0, true)
	top += 100

	// 一次, 往返
	hAnimation = xc.XAnima_Create(list_svg[1], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 1, 0, true)
	top += 100

	// 一次, 不往返
	hAnimation = xc.XAnima_Create(list_svg[2], 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 2000, 750, float32(top), 1, 0, false)
	*/
	return 0
}

// 6.形状文本
func OnBtnClick6(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 140
	var top int32 = 100

	// 创建形状文本
	hShapeText1 := xc.XShapeText_Create(left, top, 100, 30, "循环滚动", w.Handle)
	top += 50
	hShapeText2 := xc.XShapeText_Create(left, top, 100, 30, "往返滚动", w.Handle)
	top += 50
	hShapeText3 := xc.XShapeText_Create(left, top, 100, 30, "移动到末尾", w.Handle)
	top += 50
	list_xcgui = append(list_xcgui,
		hShapeText1,
		hShapeText2,
		hShapeText3,
	)
	top = 100

	ani1 := ani.NewAnima(hShapeText1, 0)
	list_animation = append(list_animation, ani1.Handle)
	ani1.Run(w.Handle)
	ani1.Move(3000, 750, float32(top), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)

	ani2 := ani.NewAnima(hShapeText2, 1)
	list_animation = append(list_animation, ani2.Handle)
	ani2.Run(w.Handle)
	ani2.Move(3000, 750, float32(top+50), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)

	ani3 := ani.NewAnima(hShapeText3, 1)
	list_animation = append(list_animation, ani3.Handle)
	ani3.Run(w.Handle)
	ani3.Move(1500, 750, float32(top+100), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)

	/* 	以下是纯函数方式实现
	hAnimation := xc.XAnima_Create(hShapeText1, 0)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 3000, 750, float32(top), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)

	hAnimation = xc.XAnima_Create(hShapeText2, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 3000, 750, float32(top+50), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)

	hAnimation = xc.XAnima_Create(hShapeText3, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, w.Handle)
	xc.XAnima_Move(hAnimation, 1500, 750, float32(top+100), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
	*/
	return 0
}

// 7.按钮
func OnBtnClick7(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 125
	var top int32 = 50

	{
		group1 := ani.NewAnimaGroup(0)
		list_animation = append(list_animation, group1.Handle)
		group1.Run(w.Handle)
		for i := 0; i < 13; i++ {
			hButton := xc.XBtn_Create(left, top, 60, 30, "透明度", w.Handle)
			list_xcgui = append(list_xcgui, hButton)

			hAnimation := xc.XAnima_Create(hButton, 0)
			group1.AddItem(hAnimation)

			xc.XAnima_Delay(hAnimation, 500)

			xc.XAnima_Delay(hAnimation, 100*float32(i))
			xc.XAnima_AlphaEx(hAnimation, 1200, 255, 20, 1, 0, true)
			left += 61
		}
	}

	left = 125
	top = 100
	group2 := ani.NewAnimaGroup(0)
	list_animation = append(list_animation, group2.Handle)
	group2.Run(w.Handle)
	for i := 0; i < 7; i++ {
		hButton := xc.XBtn_Create(left, top, 80, 30, "循环滚动", w.Handle)
		list_xcgui = append(list_xcgui, hButton)

		{
			hAnimation := xc.XAnima_Create(hButton, 0)
			group2.AddItem(hAnimation)

			xc.XAnima_Move(hAnimation, 500, float32(left), float32(top), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
			xc.XAnima_Delay(hAnimation, 500)

			xc.XAnima_Delay(hAnimation, 100*float32(i))
			xc.XAnima_AlphaEx(hAnimation, 500, 255, 0, 1, 0, false)

			xc.XAnima_Delay(hAnimation, 500)

			xc.XAnima_AlphaEx(hAnimation, 500, 0, 255, 1, 0, false)
			xc.XAnima_Delay(hAnimation, 1000)

			xc.XAnima_Move(hAnimation, 2000, float32(left), 500, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
			xc.XAnima_Delay(hAnimation, 1000)
		}
		{
			hAnimation := xc.XAnima_Create(hButton, 1)
			xc.XAnimaGroup_AddItem(group2.Handle, hAnimation)
			xc.XAnima_Delay(hAnimation, 6000+float32(i)*200)
			xc.XAnima_Scale(hAnimation, 1200, 1.5, 2, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, true)
		}
		left += 110
	}
	return 0
}

// 8.布局焦点展开
func OnBtnClick8(hEle int, pbHandled *bool) int {
	ReleaseAnimation()

	layBody := widget.NewLayoutEle(140, 100, 750, 100, w.Handle)
	layBody.SetSpace(20)
	list_xcgui = append(list_xcgui, layBody.Handle)

	for i := 0; i < 3; i++ {
		lay := widget.NewLayoutEle(0, 0, 100, 100, layBody.Handle)
		lay.SetPadding(10, 0, 10, 0)

		st := widget.NewShapeText(0, 0, 100, 100, "鼠标放上来查看-炫彩界面库-github.com/twgh/xcgui", lay.Handle)
		st.SetTextColor(xc.RGBA(255, 255, 255, 255))
		st.LayoutItem_SetWidth(xcc.Layout_Size_Fill, 0)

		list_xcgui = append(list_xcgui, lay.Handle)
		lay.EnableMouseThrough(false)
		lay.LayoutItem_SetWidth(xcc.Layout_Size_Weight, 100)

		xc.XBkM_SetInfo(lay.GetBkManager(), "{99:1.9.9;98:16(0);5:2(15)20(1)21(3)26(1)22(-7839744)23(255)9(5,5,5,5);}") // 这种字符串是在设计器里设计好后, 从xml里复制出来的
		lay.AddEvent_MouseStay(OnMouseStay8)
		lay.AddEvent_MouseLeave(OnMouseLeave8)

		switch i {
		case 0:
			m_hLayout1 = lay.Handle
		case 1:
			m_hLayout2 = lay.Handle
		case 2:
			m_hLayout3 = lay.Handle
		}
	}

	w.AdjustLayout().Redraw(false)
	return 0
}

// 鼠标进入事件8
func OnMouseStay8(hLayout int, pbHandled *bool) int {
	if m_hLayout1 != hLayout {
		xc.XEle_SetAlpha(m_hLayout1, 200)
	}

	if m_hLayout2 != hLayout {
		xc.XEle_SetAlpha(m_hLayout2, 200)
	}

	if m_hLayout3 != hLayout {
		xc.XEle_SetAlpha(m_hLayout3, 200)
	}

	hAnimation := xc.XAnima_Create(hLayout, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_LayoutWidth(hAnimation, 300, xcc.Layout_Size_Weight, 200, 1, 0, false)
	xc.XAnima_Run(hAnimation, w.Handle)
	return 0
}

// 鼠标离开事件8
func OnMouseLeave8(hLayout, hEleStay int, pbHandled *bool) int {
	hAnimation := xc.XAnima_Create(hLayout, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_LayoutWidth(hAnimation, 300, xcc.Layout_Size_Weight, 100, 1, 0, false)
	xc.XAnima_Run(hAnimation, w.Handle)

	xc.XEle_SetAlpha(m_hLayout1, 255)
	xc.XEle_SetAlpha(m_hLayout2, 255)
	xc.XEle_SetAlpha(m_hLayout3, 255)
	return 0
}

// 9.图片切换 - 两个基础元素透明度切换
func OnBtnClick9(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	var top int32 = 50

	imgMap := map[int][]byte{
		1: img1,
		2: img2,
		3: img3,
		4: img4,
		5: img5,
		6: img6,
	}

	for i := 0; i < 3; i++ {
		hImage := xc.XImage_LoadMemory(imgMap[i*2+1])
		xc.XImage_SetDrawType(hImage, xcc.Image_Draw_Type_Fixed_Ratio)

		ele1 := widget.NewElement(left, top, 211, 270, w.Handle)
		ele1.AddBkImage(xcc.Element_State_Flag_Leave, hImage)
		list_xcgui = append(list_xcgui, ele1.Handle)

		hImage2 := xc.XImage_LoadMemory(imgMap[i*2+2])
		xc.XImage_SetDrawType(hImage2, xcc.Image_Draw_Type_Fixed_Ratio)

		ele2 := widget.NewElement(left, top, 211, 270, w.Handle)
		ele2.AddBkImage(xcc.Element_State_Flag_Leave, hImage2)
		list_xcgui = append(list_xcgui, ele2.Handle)

		ele1.SetUserData(ele2.Handle)
		ele2.SetUserData(ele1.Handle)
		ele2.Show(false)

		hText := xc.XShapeText_Create(left, top+280, 200, 40, "炫彩界面库-图片切换\r\n$66.66", w.Handle)
		xc.XShapeText_SetTextColor(hText, xc.RGBA(80, 80, 80, 255))
		list_xcgui = append(list_xcgui, hText)

		ele1.AddEvent_MouseStay(OnMouseStay9)
		ele2.AddEvent_MouseLeave(OnMouseLeave9)

		left += 211 + 10
	}
	w.Redraw(false)
	return 0
}

// 鼠标进入事件9
func OnMouseStay9(hEle int, pbHandled *bool) int {
	hEle2 := xc.XEle_GetUserData(hEle)
	// 释放当前对象关联的动画
	for i := len(list_animation) - 1; i >= 0; i-- {
		hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
		if hEle == hObjectUI || hEle2 == hObjectUI {
			xc.XAnima_Release(list_animation[i], false)
			list_animation = append(list_animation[:i], list_animation[i+1:]...)
		}
	}

	hAnimation := xc.XAnima_Create(hEle, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_AlphaEx(hAnimation, 1000, 255, 0, 1, 0, false)
	xc.XAnima_Show(hAnimation, 0, false)

	xc.XEle_SetAlpha(hEle2, 0)
	xc.XWidget_Show(hEle2, true)

	hAnimation = xc.XAnima_Create(hEle2, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle2)
	xc.XAnima_Delay(hAnimation, 500)
	xc.XAnima_AlphaEx(hAnimation, 1000, 0, 255, 1, 0, false)
	return 0
}

// 鼠标离开事件9
func OnMouseLeave9(hEle2, hEleStay int, pbHandled *bool) int {
	hEle := xc.XEle_GetUserData(hEle2)
	// 释放当前对象关联的动画
	for i := len(list_animation) - 1; i >= 0; i-- {
		hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
		if hEle == hObjectUI || hEle2 == hObjectUI {
			xc.XAnima_Release(list_animation[i], false)
			list_animation = append(list_animation[:i], list_animation[i+1:]...)
		}
	}

	hAnimation := xc.XAnima_Create(hEle2, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle2)
	xc.XAnima_AlphaEx(hAnimation, 1000, 255, 0, 1, 0, false)
	xc.XAnima_Show(hAnimation, 0, false)

	xc.XEle_SetAlpha(hEle, 0)
	xc.XWidget_Show(hEle, true)

	hAnimation = xc.XAnima_Create(hEle, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Delay(hAnimation, 500)
	xc.XAnima_AlphaEx(hAnimation, 1000, 0, 255, 1, 0, false)
	return 0
}

// 10.图片切换2 - 滚动切换
func OnBtnClick10(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	var top int32 = 50

	imgMap := map[int][]byte{
		1: img1,
		2: img2,
		3: img3,
		4: img4,
		5: img5,
		6: img6,
	}

	for i := 0; i < 3; i++ {
		ele := widget.NewElement(left, top, 211, 270, w.Handle)
		ele.EnableDrawBorder(false)
		list_xcgui = append(list_xcgui, ele.Handle)

		hImage := xc.XImage_LoadMemory(imgMap[i*2+1])
		xc.XImage_SetDrawType(hImage, xcc.Image_Draw_Type_Fixed_Ratio)

		hImage2 := xc.XImage_LoadMemory(imgMap[i*2+2])
		xc.XImage_SetDrawType(hImage2, xcc.Image_Draw_Type_Fixed_Ratio)

		hShapePic := xc.XShapePic_Create(0, 0, 211, 270, ele.Handle)
		xc.XShapePic_SetImage(hShapePic, hImage)

		hShapePic2 := xc.XShapePic_Create(211+10, 0, 211, 270, ele.Handle)
		xc.XShapePic_SetImage(hShapePic2, hImage2)

		hText := xc.XShapeText_Create(left, top+280, 200, 40, "炫彩界面库-图片切换2\r\n$66.66", w.Handle)
		xc.XShapeText_SetTextColor(hText, xc.RGBA(80, 80, 80, 255))
		list_xcgui = append(list_xcgui, hText)

		ele.AddEvent_MouseStay(OnMouseStay10)
		ele.AddEvent_MouseLeave(OnMouseLeave10)

		left += 211 + 10
	}
	w.Redraw(false)
	return 0
}

// 鼠标进入事件10
func OnMouseStay10(hEle int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := len(list_animation) - 1; i >= 0; i-- {
		hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
		if hEle == hObjectUI {
			xc.XAnima_Release(list_animation[i], false)
			list_animation = append(list_animation[:i], list_animation[i+1:]...)
		}
	}

	hPic := xc.XEle_GetChildByIndex(hEle, 0)

	hAnimation := xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, -(211 + 10), 0, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, false)

	hPic = xc.XEle_GetChildByIndex(hEle, 1)

	hAnimation = xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, 0, 0, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, false)
	return 0
}

// 鼠标离开事件10
func OnMouseLeave10(hEle, hEleStay int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := len(list_animation) - 1; i >= 0; i-- {
		hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
		if hEle == hObjectUI {
			xc.XAnima_Release(list_animation[i], false)
			list_animation = append(list_animation[:i], list_animation[i+1:]...)
		}
	}

	hPic := xc.XEle_GetChildByIndex(hEle, 0)

	hAnimation := xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, 0, 0, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, false)

	hPic = xc.XEle_GetChildByIndex(hEle, 1)

	hAnimation = xc.XAnima_Create(hPic, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Run(hAnimation, hEle)
	xc.XAnima_Move(hAnimation, 500, 211+10, 0, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_In, false)
	return 0
}

// 11.进度 等待
func OnBtnClick11(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 160
	var top int32 = 80
	var hSvg, hGroup, hAnimation int

	// 两个球型交替移动
	{
		hSvg := xc.XSvg_LoadString(svg11)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)

		hGroup := xc.XAnimaGroup_Create(0)
		list_animation = append(list_animation, hGroup)
		xc.XAnima_Run(hGroup, w.Handle)

		hAnimation := xc.XAnima_Create(hSvg, 1)
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)
		xc.XAnima_Move(hAnimation, 1000, float32(left)+50, float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)
		xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)

		hSvg = xc.XSvg_LoadString(svg12)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left+50, top)

		hGroup = xc.XAnimaGroup_Create(0)
		list_animation = append(list_animation, hGroup)
		xc.XAnima_Run(hGroup, w.Handle)

		hAnimation = xc.XAnima_Create(hSvg, 1)
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)
		xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)
		xc.XAnima_Move(hAnimation, 1000, float32(left)+50, float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)
	}

	// 一排小球 缩放
	{
		left = 350
		hGroup = xc.XAnimaGroup_Create(0)
		list_animation = append(list_animation, hGroup)
		xc.XAnima_Run(hGroup, w.Handle)

		for i := 0; i < 10; i++ {
			hSvg = xc.XSvg_LoadString(svg13)
			list_svg = append(list_svg, hSvg)
			xc.XSvg_SetPosition(hSvg, left+int32(i)*50, top)

			hAnimation = xc.XAnima_Create(hSvg, 0)
			xc.XAnimaGroup_AddItem(hGroup, hAnimation)

			xc.XAnima_Delay(hAnimation, float32(i)*200)
			xc.XAnima_Scale(hAnimation, 1000, 2, 2, 1, 0, true)
		}
	}

	// 一排小球 垂直缩放
	{
		top = 150
		hGroup = xc.XAnimaGroup_Create(0)
		list_animation = append(list_animation, hGroup)
		xc.XAnima_Run(hGroup, w.Handle)

		for i := 0; i < 10; i++ {
			hSvg = xc.XSvg_LoadString(svg13)
			list_svg = append(list_svg, hSvg)
			xc.XSvg_SetPosition(hSvg, left+int32(i)*50, top)

			hAnimation = xc.XAnima_Create(hSvg, 0)
			xc.XAnimaGroup_AddItem(hGroup, hAnimation)

			xc.XAnima_Delay(hAnimation, float32(i)*200)
			xc.XAnima_Scale(hAnimation, 1000, 1, 2, 1, 0, true)
		}
	}

	// 一排小球 上下波浪
	{
		left = 150
		top = 200
		for i := 0; i < 10; i++ {
			hSvg = xc.XSvg_LoadString(svg13)
			list_svg = append(list_svg, hSvg)
			x := left + int32(i)*35
			xc.XSvg_SetPosition(hSvg, x, top)

			hAnimation = xc.XAnima_Create(hSvg, 0)
			list_animation = append(list_animation, hAnimation)
			xc.XAnima_Run(hAnimation, w.Handle)

			xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*100), true)
			xc.XAnima_Move(hAnimation, 1200, float32(x), float32(top)+100, 1, 0, true)
		}
	}

	// 一排小球上下波浪
	{
		left = 550
		for i := 0; i < 10; i++ {
			hSvg = xc.XSvg_LoadString(svg13)
			list_svg = append(list_svg, hSvg)
			x := left + int32(i)*35
			xc.XSvg_SetPosition(hSvg, x, top)

			hAnimation = xc.XAnima_Create(hSvg, 0)
			list_animation = append(list_animation, hAnimation)
			xc.XAnima_Run(hAnimation, w.Handle)

			xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*150), true)
			xc.XAnima_Move(hAnimation, 1000, float32(x), float32(top)+50, 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, true)
		}
	}

	// 一排小球 跳动
	{
		left = 150
		top = 350
		for i := 0; i < 10; i++ {
			hSvg = xc.XSvg_LoadString(svg13)
			list_svg = append(list_svg, hSvg)
			x := left + int32(i)*35
			xc.XSvg_SetPosition(hSvg, x, top)

			hAnimation = xc.XAnima_Create(hSvg, 0)
			list_animation = append(list_animation, hAnimation)
			xc.XAnima_Run(hAnimation, w.Handle)

			xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*200), true)
			xc.XAnima_Move(hAnimation, 500, float32(x), float32(top)+50, 1, xcc.Ease_Flag_Quint|xcc.Ease_Flag_Out, true)
			xc.XAnima_Delay(hAnimation, 1700)
		}
	}

	// 一排小球 移动
	{
		left = 220
		top = 600
		for i := 5; i >= 0; i-- {
			hSvg = xc.XSvg_LoadString(svg14)
			list_svg = append(list_svg, hSvg)
			xc.XSvg_SetPosition(hSvg, 100-int32(i)*25, top)
			xc.XSvg_SetAlpha(hSvg, 0)

			{
				hAnimation = xc.XAnima_Create(hSvg, 0)
				xc.XAnima_Run(hAnimation, w.Handle)
				list_animation = append(list_animation, hAnimation)

				xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*100), true)
				xc.XAnima_Move(hAnimation, 2000, 550-float32(i)*25, float32(top), 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
				xc.XAnima_Move(hAnimation, 2000, 900-float32(i)*25, float32(top), 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, false)
				xc.XAnima_Move(hAnimation, 0, 100-float32(i)*25, float32(top), 1, 0, false)
				xc.XAnima_Delay(hAnimation, 500)
			}
			{
				hAnimation = xc.XAnima_Create(hSvg, 0)
				xc.XAnima_Run(hAnimation, w.Handle)
				list_animation = append(list_animation, hAnimation)

				xc.XAnimaItem_EnableCompleteRelease(xc.XAnima_Delay(hAnimation, float32(i)*100), true)
				xc.XAnima_AlphaEx(hAnimation, 2000, 0, 255, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
				xc.XAnima_AlphaEx(hAnimation, 2000, 255, 0, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, false)
				xc.XAnima_Delay(hAnimation, 500)
			}
		}
	}
	w.Redraw(false)
	return 0
}

// 12.旋转 移动
func OnBtnClick12(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 120
	var top int32 = 100
	var hSvg, hAnimation int

	// 移动 360度旋转
	{
		hSvg := xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, 0)

		hAnimation := xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Rotate(hAnimation, 1700, 360, 1, 0, false)
		xc.XAnima_Run(hAnimation, w.Handle)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Move(hAnimation, 3000, float32(left)+500, float32(top), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 移动 往返旋转
	{
		top = 350
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, -45)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 0, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Rotate(hAnimation, 1500, 45, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, false)
		xc.XAnima_Rotate(hAnimation, 1500, -45, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, false)
		xc.XAnima_Run(hAnimation, w.Handle)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Move(hAnimation, 3000, float32(left)+500, float32(top), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	return 0
}

// 13.旋转 摇摆
func OnBtnClick13(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 130
	var top int32 = 80
	var hSvg, hAnimation, hRotate int

	// 自身 摇摆 往返
	{
		hSvg := xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, -45)

		hAnimation := xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Rotate(hAnimation, 1000, 45, 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 自身 旋转
	{
		left = 500
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Rotate(hAnimation, 1000, 360, 1, xcc.Ease_Flag_Expo|xcc.Ease_Flag_In, false)
		xc.XAnima_Rotate(hAnimation, 0, 0, 1, xcc.Ease_Flag_Linear, false)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 两个叠加 悬挂摆动
	{
		left = 300
		top = 250
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, 45)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hRotate := xc.XAnima_Rotate(hAnimation, 3000, 100, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_InOut, true)
		xc.XAnimaRotate_SetCenter(hRotate, float32(left)+10, float32(top)+50, false)
		xc.XAnima_Run(hAnimation, w.Handle)

		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, 45)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 0, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hRotate = xc.XAnima_Rotate(hAnimation, 3000, 100, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_InOut, true)
		xc.XAnimaRotate_SetCenter(hRotate, float32(left)+10, float32(top)+50, false)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 砍东西效果
	{
		left = 500
		top = 400
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, -45)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 0, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hRotate = xc.XAnima_Rotate(hAnimation, 1500, 0, 1, xcc.Ease_Flag_Expo|xcc.Ease_Flag_In, true)
		xc.XAnimaRotate_SetCenter(hRotate, float32(left), float32(top), false)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	return 0
}

// 14.旋转 移动 缩放
func OnBtnClick14(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 130
	var top int32 = 50

	// 加载svg, 设置大小和填充颜色
	hSvg := xc.XSvg_LoadString(svg7)
	list_svg = append(list_svg, hSvg)
	xc.XSvg_SetSize(hSvg, 50, 50)
	xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 0, 0, 255), true)

	// 移动 360度旋转
	xc.XSvg_SetPosition(hSvg, left, top)
	xc.XSvg_SetRotateAngle(hSvg, 0)

	// 创建动画组
	hGroup := xc.XAnimaGroup_Create(0)
	list_animation = append(list_animation, hGroup)

	// 旋转
	hAnimation := xc.XAnima_Create(hSvg, 0)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Rotate(hAnimation, 600, 360, 4, 0, false)

	// 缩放
	hAnimation = xc.XAnima_Create(hSvg, 0)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Scale(hAnimation, 2400, 7, 7, 1, 0, false)
	xc.XAnima_Delay(hAnimation, 1000)
	xc.XAnima_Scale(hAnimation, 1000, 1.0/7.0, 1.0/7.0, 1, 0, false)

	// 移动
	hAnimation = xc.XAnima_Create(hSvg, 0)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)
	xc.XAnima_Move(hAnimation, 2400, float32(left)+500, float32(top)+300, 1, 0, false)
	xc.XAnima_Delay(hAnimation, 1000)
	xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, 0, false)
	xc.XAnima_Run(hGroup, w.Handle)
	return 0
}

// 15.旋转 开合效果
func OnBtnClick15(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	var top int32 = 200
	var height, width int32
	var hSvg, hAnimation, hRotate int

	// 砍东西效果
	{
		hSvg := xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		height = xc.XSvg_GetHeight(hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, -45)

		hAnimation := xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hRotate := xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)
		xc.XAnimaRotate_SetCenter(hRotate, float32(left), float32(top+height/2.0), false)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 砍东西效果
	{
		top = 300
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		height = xc.XSvg_GetHeight(hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, 45)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hRotate = xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)
		xc.XAnimaRotate_SetCenter(hRotate, float32(left), float32(top+height/2.0), false)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 砍东西效果
	{
		left = 500
		top = 200
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		width = xc.XSvg_GetWidth(hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, 45)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 0, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hRotate = xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)
		xc.XAnimaRotate_SetCenter(hRotate, float32(left+width), float32(top+height/2.0), false)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 砍东西效果
	{
		top = 300
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		width = xc.XSvg_GetWidth(hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetRotateAngle(hSvg, -45)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 0, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hRotate = xc.XAnima_Rotate(hAnimation, 2000, 0, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, true)
		xc.XAnimaRotate_SetCenter(hRotate, float32(left+width), float32(top+height/2.0), false)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	return 0
}

// 16.颜色渐变
func OnBtnClick16(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	var top int32 = 50
	var hSvg, hAnimation int

	{
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 0, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Color(hAnimation, 1500, xc.RGBA(0, 0, 255, 255), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top = 225
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(0, 255, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Color(hAnimation, 1500, xc.RGBA(255, 0, 0, 255), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top = 400
		hSvg = xc.XSvg_LoadString(svg7)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 255, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Color(hAnimation, 1500, xc.RGBA(0, 0, 255, 255), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		hSvg = xc.XSvg_LoadString(svg15)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, 500, 300)
		xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 255, 0, 255), true)

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Color(hAnimation, 1500, xc.RGBA(0, 255, 255, 255), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		hFontx := xc.XFont_CreateEx("微软雅黑", 36, xcc.FontStyle_Bold)
		hShapeText := xc.XShapeText_Create(500, 100, 400, 50, "炫彩界面库", w.Handle)
		xc.XWidget_LayoutItem_SetWidth(hShapeText, xcc.Layout_Size_Auto, -1) // 自动宽度
		list_xcgui = append(list_xcgui, hShapeText)
		xc.XShapeText_SetFont(hShapeText, hFontx)
		xc.XShapeText_SetTextColor(hShapeText, xc.RGBA(255, 0, 0, 255))

		hAnimation = xc.XAnima_Create(hShapeText, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Color(hAnimation, 1500, xc.RGBA(0, 0, 255, 255), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		hShapeText := xc.XShapeText_Create(500, 200, 100, 20, "炫彩界面库", w.Handle)
		xc.XWidget_LayoutItem_SetWidth(hShapeText, xcc.Layout_Size_Auto, -1) // 自动宽度
		list_xcgui = append(list_xcgui, hShapeText)

		hAnimation = xc.XAnima_Create(hShapeText, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnima_Color(hAnimation, 1500, xc.RGBA(0, 255, 0, 255), 1, 0, true)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	return 0
}

// 17.缩放 位置
func OnBtnClick17(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	var top int32 = 50
	var hSvg, hAnimation, hScale int

	{
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_leftTop", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_LeftTop)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_left", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_Left)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_leftBottom", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_LeftBottom)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top = 50
		left += 150
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_top", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_Top)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_center", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_Center)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_bottom", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_Bottom)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		left += 150
		top = 50
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_rightTop", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_RightTop)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_right", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_Right)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	{
		top += 150
		hSvg = xc.XSvg_LoadString(svg5)
		list_svg = append(list_svg, hSvg)
		xc.XSvg_SetPosition(hSvg, left, top)
		list_xcgui = append(list_xcgui, xc.XShapeText_Create(left, top+65, 150, 20, "position_flag_rightBottom", w.Handle))

		hAnimation = xc.XAnima_Create(hSvg, 0)
		list_animation = append(list_animation, hAnimation)
		hScale = xc.XAnima_Scale(hAnimation, 3000, 0.5, 0.5, 1, 0, true)
		xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_RightBottom)
		xc.XAnima_Run(hAnimation, w.Handle)
	}
	return 0
}

// 18.按钮 宽度
func OnBtnClick18(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var left int32 = 150
	var top int32 = 50
	var hFont = font.New(10).Handle

	for i := 0; i < 5; i++ {
		btn := widget.NewButton(left, top, 100, 50, "鼠标 停留 离开", w.Handle)
		list_xcgui = append(list_xcgui, btn.Handle)
		btn.SetFont(hFont)
		btn.SetTextColor(xc.RGBA(255, 255, 255, 255))
		xc.XBkM_SetInfo(btn.GetBkManager(), "{99:1.9.9;98:16(0)32(1)64(2);5:2(15)20(1)21(3)26(1)22(-25024)23(255)9(4,4,4,4);5:2(15)20(1)21(3)26(1)22(-20122)23(255)9(4,4,4,4);5:2(15)20(1)21(3)26(1)22(-1667526)23(255)9(4,4,4,4);}") // 这种字符串是在设计器里设计好后, 从xml里复制出来的

		btn.AddEvent_MouseStay(OnMouseStay18)
		btn.AddEvent_MouseLeave(OnMouseLeave18)
		top += 60
	}
	w.Redraw(false)
	return 0
}

// 鼠标进入事件18
func OnMouseStay18(hButton int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := len(list_animation) - 1; i >= 0; i-- {
		hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
		if hButton == hObjectUI {
			xc.XAnima_Release(list_animation[i], false)
			list_animation = append(list_animation[:i], list_animation[i+1:]...)
		}
	}

	hAnimation := xc.XAnima_Create(hButton, 1)
	list_animation = append(list_animation, hAnimation)
	hScale := xc.XAnima_ScaleSize(hAnimation, 400, 250, 40, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
	xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_Left)
	xc.XAnima_Run(hAnimation, w.Handle)
	return 0
}

// 鼠标离开事件18
func OnMouseLeave18(hButton, hEleStay int, pbHandled *bool) int {
	// 释放当前对象关联的动画
	for i := len(list_animation) - 1; i >= 0; i-- {
		hObjectUI := xc.XAnima_GetObjectUI(list_animation[i])
		if hButton == hObjectUI {
			xc.XAnima_Release(list_animation[i], false)
			list_animation = append(list_animation[:i], list_animation[i+1:]...)
		}
	}

	hAnimation := xc.XAnima_Create(hButton, 1)
	list_animation = append(list_animation, hAnimation)
	hScale := xc.XAnima_ScaleSize(hAnimation, 400, 150, 40, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, false)
	xc.XAnimaScale_SetPosition(hScale, xcc.Position_Flag_Left)
	xc.XAnima_Run(hAnimation, w.Handle)
	return 0
}

// 19.窗口特效
func OnBtnClick19(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var top int32 = 200
	var left int32 = 140
	var width int32 = 120
	var height_btn int32 = 35
	var height int32 = 34

	btn := CreateButton(left, top, width, height_btn, "窗口 从上往下")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick19_1)

	btn = CreateButton(left, top, width, height_btn, "窗口 从左往右")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick19_2)

	btn = CreateButton(left, top, width, height_btn, "窗口 缩放")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick19_3)

	btn = CreateButton(left, top, width, height_btn, "窗口 缩放2")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick19_4)

	btn = CreateButton(left, top, width, height_btn, "窗口 透明")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick19_5)

	w.Redraw(false)
	return 0
}

// 19.1 窗口缓动 从上往下
func OnBtnClick19_1(hEle int, pbHandled *bool) int {
	m := window.NewModalWindow(400, 300, "窗口缓动", w.GetHWND(), xcc.Window_Style_Modal|xcc.Window_Style_Drag_Window)

	rcWindow := w.GetRectDPI()
	rcModal := m.GetRectDPI()
	left := float32(rcWindow.Left + (rcWindow.Right-rcWindow.Left-(rcModal.Right-rcModal.Left))/2)
	top := float32(rcWindow.Top + (rcWindow.Bottom-rcWindow.Top-(rcModal.Bottom-rcModal.Top))/2)

	hAnimation := xc.XAnima_Create(m.Handle, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_MoveEx(hAnimation, 1000, left, 20, left, top, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
	xc.XAnima_Run(hAnimation, m.Handle)

	m.DoModal()
	return 0
}

// 19.2 窗口缓动 从左往右
func OnBtnClick19_2(hEle int, pbHandled *bool) int {
	m := window.NewModalWindow(400, 300, "窗口缓动", w.GetHWND(), xcc.Window_Style_Modal|xcc.Window_Style_Drag_Window)

	rcWindow := w.GetRectDPI()
	rcModal := m.GetRectDPI()
	left := float32(rcWindow.Left + (rcWindow.Right-rcWindow.Left-(rcModal.Right-rcModal.Left))/2)
	top := float32(rcWindow.Top + (rcWindow.Bottom-rcWindow.Top-(rcModal.Bottom-rcModal.Top))/2)

	hAnimation := xc.XAnima_Create(m.Handle, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_MoveEx(hAnimation, 1000, 20, top, left, top, 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
	xc.XAnima_Run(hAnimation, m.Handle)

	m.DoModal()
	return 0
}

// 19.3 窗口缩放
func OnBtnClick19_3(hEle int, pbHandled *bool) int {
	m := window.NewModalWindow(400, 300, "窗口缩放", w.GetHWND(), xcc.Window_Style_Modal|xcc.Window_Style_Drag_Window)

	rcModal := m.GetRectEx()
	fmt.Println(rcModal)
	width := rcModal.Right - rcModal.Left
	height := rcModal.Bottom - rcModal.Top

	hAnimation := xc.XAnima_Create(m.Handle, 1)
	list_animation = append(list_animation, hAnimation)

	// TODO: 这里有个BUG, 导致窗口位置被改变了, 不应该改变才对, 是开启AutoDPI后出现的BUG
	xc.XAnima_ScaleSize(hAnimation, 1000, float32(width)*1.5, float32(height)*1.5, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, true)
	xc.XAnima_Run(hAnimation, m.Handle)

	rcModal = m.GetRectEx()
	fmt.Println(rcModal)
	m.DoModal()
	return 0
}

// 19.4 窗口缩放2
func OnBtnClick19_4(hEle int, pbHandled *bool) int {
	m := window.NewModalWindow(400, 300, "窗口缩放2", w.GetHWND(), xcc.Window_Style_Modal|xcc.Window_Style_Drag_Window)

	rcModal := m.GetRectEx()
	width := rcModal.Right - rcModal.Left
	height := rcModal.Bottom - rcModal.Top

	hAnimation := xc.XAnima_Create(m.Handle, 1)
	list_animation = append(list_animation, hAnimation)

	// TODO: 这里有个BUG, 导致窗口位置被改变了, 不应该改变才对, 是开启AutoDPI后出现的BUG
	xc.XAnima_ScaleSize(hAnimation, 1000, float32(width)*2, float32(height)*2, 1, xcc.Ease_Flag_Back|xcc.Ease_Flag_Out, false)

	xc.XAnima_Run(hAnimation, m.Handle)
	m.DoModal()
	return 0
}

// 19.5 窗口透明
func OnBtnClick19_5(hEle int, pbHandled *bool) int {
	hModal := xc.XModalWnd_Create(400, 300, "窗口透明", w.GetHWND(), xcc.Window_Style_Modal|xcc.Window_Style_Drag_Window)
	xc.XWnd_SetTransparentType(hModal, xcc.Window_Transparent_Shadow)
	xc.XWnd_SetTransparentAlpha(hModal, 1)

	hAnimation := xc.XAnima_Create(hModal, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Delay(hAnimation, 100)
	xc.XAnima_Alpha(hAnimation, 1000, 255, 1, 0, false)
	xc.XAnima_Run(hAnimation, hModal)

	xc.XModalWnd_DoModal(hModal)
	return 0
}

// 20. 遮盖弹窗
func OnBtnClick20(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var top int32 = 200
	var left int32 = 140
	var width int32 = 150
	var height_btn int32 = 35
	var height int32 = 34

	btn := CreateButton(left, top, width, height_btn, "遮盖层-内嵌子弹窗")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick20_1)

	btn = CreateButton(left, top, width, height_btn, "遮盖层-内嵌消息框")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick20_2)

	btn = CreateButton(left, top, width, height_btn, "遮盖层-消息框")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick20_3)

	btn = CreateButton(left, top, width, height_btn, "遮盖层-等待")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick20_4)

	btn = CreateButton(left, top, width, height_btn, "遮盖层-基础元素弹窗")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick20_5)

	w.Redraw(false)
	return 0
}

var hEle_mask int // 遮罩

// 20.1 遮盖层 内嵌子弹窗
func OnBtnClick20_1(hEle int, pbHandled *bool) int {
	var rect xc.RECT
	w.GetBodyRect(&rect)

	hEle_mask = xc.XEle_Create(rect.Left, rect.Top, rect.Right-rect.Left, rect.Bottom-rect.Top, w.Handle)
	xc.XEle_AddBkFill(hEle_mask, xcc.CombinedState_Window_State_Flag_Leave, xc.RGBA(0, 0, 0, 200))
	xc.XEle_Redraw(hEle_mask, true)

	wd := window.NewEx(0, xcc.WS_CHILD, "", 0, 0, 300, 200, "内嵌子弹窗", w.GetHWND(), xcc.Window_Style_Default)
	wd.Show(true)
	wd.AddEvent_Destroy(OnWndDestroy20)
	return 0
}

func OnWndDestroy20(hWindow int, pbHandled *bool) int {
	if hEle_mask != 0 {
		xc.XEle_Destroy(hEle_mask)
		hEle_mask = 0
		w.Redraw(false)
	}
	return 0
}

// 20.2 遮盖层 内嵌消息框
func OnBtnClick20_2(hEle int, pbHandled *bool) int {
	var rect xc.RECT
	w.GetBodyRect(&rect)

	hEle_mask = xc.XEle_Create(rect.Left, rect.Top, rect.Right-rect.Left, rect.Bottom-rect.Top, w.Handle)
	xc.XEle_AddBkFill(hEle_mask, xcc.CombinedState_Window_State_Flag_Leave, xc.RGBA(0, 0, 0, 200))
	xc.XEle_Redraw(hEle_mask, true)

	wd := w.Msg_CreateEx(0, xcc.WS_CHILD, "", "标题", "内嵌消息框", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Default)
	wd.Show(true)
	wd.AddEvent_Destroy(OnWndDestroy20)
	return 0
}

// 20.3 遮盖层 消息框
func OnBtnClick20_3(hEle int, pbHandled *bool) int {
	var rect xc.RECT
	w.GetBodyRect(&rect)

	hEle_mask = xc.XEle_Create(rect.Left, rect.Top, rect.Right-rect.Left, rect.Bottom-rect.Top, w.Handle)
	xc.XEle_AddBkFill(hEle_mask, xcc.CombinedState_Window_State_Flag_Leave, xc.RGBA(0, 0, 0, 200))
	xc.XEle_Redraw(hEle_mask, true)

	wd := w.Msg_Create("标题", "消息框", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Default)
	wd.Show(true)
	wd.AddEvent_Destroy(OnWndDestroy20)
	return 0
}

var hSvg1, hSvg2 int

// 20.4 遮盖层 等待
func OnBtnClick20_4(hEle int, pbHandled *bool) int {
	const str = `<svg x="0" y="0" width="25" height="25" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#ee6362" /></svg>`
	const str2 = `<svg x="0" y="0" width="25" height="25" viewBox="0 0 100 100"><circle cx="50" cy="50" r="50" fill="#2cb0b2" /></svg>`

	hSvg1 = xc.XSvg_LoadString(str)
	hSvg2 = xc.XSvg_LoadString(str2)

	var rect xc.RECT
	w.GetBodyRect(&rect)

	eleMask := widget.NewElement(rect.Left, rect.Top, rect.Right-rect.Left, rect.Bottom-rect.Top, w.Handle)
	hEle_mask = eleMask.Handle

	eleMask.AddBkFill(xcc.CombinedState_Window_State_Flag_Leave, xc.RGBA(0, 0, 0, 200))
	eleMask.Redraw(true)

	eleMask.AddEvent_Paint(OnDraw20_4)
	eleMask.AddEvent_LButtonDown(OnLButtonDown20_4and5)

	left := rect.Left + (rect.Right-rect.Left-100)/2
	top := (rect.Bottom-rect.Top)/2 - 50
	xc.XShapeText_SetTextColor(xc.XShapeText_Create(left, top-25, 100, 20, "正在加载...", hEle_mask), xc.RGBA(255, 255, 255, 255))

	// 两个球型交替移动
	{
		xc.XSvg_SetPosition(hSvg1, left, top)

		hGroup := xc.XAnimaGroup_Create(0)
		list_animation = append(list_animation, hGroup)
		xc.XAnima_Run(hGroup, w.Handle)

		hAnimation := xc.XAnima_Create(hSvg1, 1)
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)
		xc.XAnima_Move(hAnimation, 1000, float32(left+50), float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)
		xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)

		xc.XSvg_SetPosition(hSvg2, left+50, top)

		hGroup = xc.XAnimaGroup_Create(0)
		list_animation = append(list_animation, hGroup)
		xc.XAnima_Run(hGroup, w.Handle)

		hAnimation = xc.XAnima_Create(hSvg2, 1)
		xc.XAnimaGroup_AddItem(hGroup, hAnimation)
		xc.XAnima_Move(hAnimation, 1000, float32(left), float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)
		xc.XAnima_Move(hAnimation, 1000, float32(left+50), float32(top), 1, xcc.Ease_Flag_Sine|xcc.Ease_Flag_InOut, false)
	}
	return 0
}

func OnDraw20_4(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	xc.XEle_DrawEle(hEle, hDraw)

	xc.XDraw_DrawSvgSrc(hDraw, hSvg1)
	xc.XDraw_DrawSvgSrc(hDraw, hSvg2)
	return 0
}

func OnLButtonDown20_4and5(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
	*pbHandled = true
	xc.XEle_Destroy(hEle)

	if hSvg1 != 0 {
		xc.XSvg_Destroy(hSvg1)
		hSvg1 = 0
	}
	if hSvg2 != 0 {
		xc.XSvg_Destroy(hSvg2)
		hSvg2 = 0
	}
	w.Redraw(false)
	return 0
}

// 20.5 遮盖层 基础元素弹窗
func OnBtnClick20_5(hEle int, pbHandled *bool) int {
	var rect xc.RECT
	w.GetBodyRect(&rect)

	eleMask := widget.NewElement(rect.Left, rect.Top, rect.Right-rect.Left, rect.Bottom-rect.Top, w.Handle)
	hEle_mask = eleMask.Handle

	eleMask.AddBkFill(xcc.CombinedState_Window_State_Flag_Leave, xc.RGBA(0, 0, 0, 200))
	eleMask.AddEvent_LButtonDown(OnLButtonDown20_4and5)

	var width int32 = 350
	var height int32 = 170
	left := (rect.Right - rect.Left - width) / 2
	top := (rect.Bottom - rect.Top - height) / 2

	hEleDlg := xc.XEle_Create(left, 10, width, height, eleMask.Handle)
	xc.XWidget_Show(hEleDlg, false)
	xc.XEle_EnableBkTransparent(hEleDlg, true)
	xc.XBkM_SetInfo(xc.XEle_GetBkManager(hEleDlg), "{99:1.9.9;98:1(0);5:2(15)20(1)21(3)26(1)22(-1)23(255)9(10,10,10,10);}") // 这种字符串是在设计器里设计好后, 从xml里复制出来的
	xc.XShapeText_SetTextColor(xc.XShapeText_Create(50, 5, 220, 20, "炫彩界面库-仅作功能演示,没有美化处理", hEleDlg), xc.RGBA(80, 80, 80, 255))

	btnClose := widget.NewButton(width-40, 2, 30, 22, "", hEleDlg)
	xc.XBkM_SetInfo(btnClose.GetBkManager(), "{99:1.9.9;98:16(0,1)32(0,1)64(0,1);5:2(48)8(45.00)3(2,10,2,10)20(1)21(3)26(0)22(-8355712)23(255);5:2(48)8(45.00)3(10,2,100,100)20(1)21(3)26(0)22(-8355712)23(255);}")

	xc.XShapeText_SetTextColor(xc.XShapeText_Create(20, 60, 200, 20, "请输入内容(这是一个演示)", hEleDlg), xc.RGBA(80, 80, 80, 255))

	strBkm := "{99:1.9.9;98:16(0)32(1)64(1);5:2(15)20(1)21(3)26(0)22(-1)23(255)10(1)7(1)11(3)16(0)12(-3618616)13(255)9(5,5,5,5);5:2(15)20(1)21(3)26(0)22(-1)23(255)10(1)7(1)11(3)16(0)12(-17897)13(255)9(5,5,5,5);}"
	hEdit := xc.XEdit_Create(20, 82, width-40, 26, hEleDlg)
	xc.XEdit_SetDefaultText(hEdit, "请输入内容...")
	xc.XEle_SetBorderSize(hEdit, 10, 0, 10, 0)
	xc.XBkM_SetInfo(xc.XEle_GetBkManager(hEdit), strBkm)

	var left_ int32 = 190
	var top_ = height - 35
	btnOK := widget.NewButton(left_, top_, 60, 22, "确定", hEleDlg)

	left_ += 80
	btnCancel := widget.NewButton(left_, top_, 60, 22, "取消", hEleDlg)
	xc.XBkM_SetInfo(btnOK.GetBkManager(), strBkm)
	xc.XBkM_SetInfo(btnCancel.GetBkManager(), strBkm)

	btnOK.SetUserData(eleMask.Handle)
	btnClose.SetUserData(eleMask.Handle)
	btnCancel.SetUserData(eleMask.Handle)

	btnOK.AddEvent_BnClick(OnBtnClick20_5_close)
	btnClose.AddEvent_BnClick(OnBtnClick20_5_close)
	btnCancel.AddEvent_BnClick(OnBtnClick20_5_close)

	hAnimation := xc.XAnima_Create(eleMask.Handle, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_AlphaEx(hAnimation, 500, 0, 255, 1, 0, false)
	xc.XAnima_Run(hAnimation, eleMask.Handle)

	hAnimation = xc.XAnima_Create(hEleDlg, 1)
	list_animation = append(list_animation, hAnimation)
	xc.XAnima_Show(hAnimation, 500, true)
	xc.XAnima_Move(hAnimation, 500, float32(left), float32(top), 1, xcc.Ease_Flag_Bounce|xcc.Ease_Flag_Out, false)
	xc.XAnima_Run(hAnimation, eleMask.Handle)

	w.Redraw(false)
	return 0
}

func OnBtnClick20_5_close(hEle int, pbHandled *bool) int {
	*pbHandled = true
	xc.XEle_Destroy(xc.XEle_GetUserData(hEle))
	w.Redraw(false)
	return 0
}

// 21.消息通知
func OnBtnClick21(hEle int, pbHandled *bool) int {
	ReleaseAnimation()
	var top int32 = 200
	var left int32 = 140
	var width int32 = 150
	var height_btn int32 = 35
	var height int32 = 34

	// -----top------------
	btn := CreateButton(left, top, width, height_btn, "top-成功")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_1)

	btn = CreateButton(left, top, width, height_btn, "top-警告消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_2)

	btn = CreateButton(left, top, width, height_btn, "top-消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_3)

	btn = CreateButton(left, top, width, height_btn, "top-错误消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_4)

	btn = CreateButton(left, top, width, height_btn, "top-没有关闭按钮")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_5)

	btn = CreateButton(left, top, width, height_btn, "top-手动关闭消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_6)

	btn = CreateButton(left, top, width, height_btn, "top-不带标题")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_7)

	btn = CreateButton(left, top, width, height_btn, "top-自定义大小")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_8)

	// -----right------------
	left += 160
	top = 200
	btn = CreateButton(left, top, width, height_btn, "right-成功")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_right_1)

	btn = CreateButton(left, top, width, height_btn, "right-警告消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_right_2)

	btn = CreateButton(left, top, width, height_btn, "right-消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_right_3)

	btn = CreateButton(left, top, width, height_btn, "right-错误消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_right_4)

	btn = CreateButton(left, top, width, height_btn, "right-没有关闭按钮")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_right_5)

	btn = CreateButton(left, top, width, height_btn, "right-手动关闭消息")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_right_6)

	btn = CreateButton(left, top, width, height_btn, "right-不带标题")
	top += height
	list_xcgui = append(list_xcgui, btn.Handle)
	btn.AddEvent_BnClick(OnBtnClick21_right_7)

	w.Redraw(false)
	return 0
}

func OnBtnClick21_1(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(16, 16)

	w.NotifyMsg_WindowPopup(xcc.Position_Flag_Top, "成功", "这是一条成功的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Success)
	return 0
}

func OnBtnClick21_2(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgWarning).SetSize(16, 16)

	w.NotifyMsg_WindowPopup(xcc.Position_Flag_Top, "警告", "这是一条警告的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Warning)
	return 0
}

func OnBtnClick21_3(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgMessage).SetSize(16, 16)

	w.NotifyMsg_WindowPopup(xcc.Position_Flag_Top, "消息", "这是一条消息的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Message)
	return 0
}

func OnBtnClick21_4(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgError).SetSize(16, 16)

	w.NotifyMsg_WindowPopup(xcc.Position_Flag_Top, "错误", "这是一条错误的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Error)
	return 0
}

func OnBtnClick21_5(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(16, 16)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Top, "成功", "这是一条成功的提示消息, 没有关闭按钮", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Success, false, true, -1, -1)
	return 0
}

func OnBtnClick21_6(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(16, 16)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Top, "成功", "这是一条成功的提示消息, 手动关闭, 这是一个自动换行文本", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_No, true, false, -1, -1)
	return 0
}

func OnBtnClick21_7(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(16, 16)

	w.NotifyMsg_WindowPopup(xcc.Position_Flag_Top, "", "这是一条成功的提示消息, 没有标题", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Success)
	return 0
}

func OnBtnClick21_8(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(16, 16)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Top, "成功", "这是一条成功的提示消息,\n自定义大小", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Success, true, true, 300, 200)
	return 0
}

func OnBtnClick21_right_1(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(20, 20)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "成功", "这是一条成功的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Success, true, true, -1, -1)
	return 0
}

func OnBtnClick21_right_2(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgWarning).SetSize(20, 20)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "警告", "这是一条警告的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Warning, true, true, -1, -1)
	return 0
}

func OnBtnClick21_right_3(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgMessage).SetSize(20, 20)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "消息", "这是一条消息的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Message, true, true, -1, -1)
	return 0
}

func OnBtnClick21_right_4(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgError).SetSize(20, 20)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "错误", "这是一条错误的提示消息", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Error, true, true, -1, -1)
	return 0
}

func OnBtnClick21_right_5(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(20, 20)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "成功", "这是一条成功的提示消息, 没有关闭按钮", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Success, false, true, -1, -1)
	return 0
}

func OnBtnClick21_right_6(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(20, 20)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "成功", "这是一条成功的提示消息, 手动关闭, 这是一个自动换行文本", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_No, true, false, -1, -1)
	return 0
}

func OnBtnClick21_right_7(hEle int, pbHandled *bool) int {
	_svg := svg.NewByString(svgSuccess).SetSize(20, 20)

	w.NotifyMsg_WindowPopupEx(xcc.Position_Flag_Right, "成功", "这是一条成功的提示消息, 没有标题", xc.XImage_LoadSvg(_svg.Handle), xcc.NotifyMsg_Skin_Success, true, true, -1, -1)
	return 0
}

var (
	hSliderBar1 int
	hSliderBar2 int
	hSliderBar3 int
	hProgBar1   int
	hProgBar2   int
	hProgBar3   int
)

// 22.进度条
func OnBtnClick22(hEle int, pbHandled *bool) int {
	ReleaseAnimation()

	var left int32 = 150
	var top int32 = 100
	var width int32 = 500

	// 滑动条
	{
		strBackground := "{99:1.9.9;98:16(0);5:2(37)3(0,8,0,0)20(1)21(3)26(1)22(-3618616)23(255)9(3,3,3,3);}"
		hSliderBar1 = xc.XSliderBar_Create(left, top, width, 20, w.Handle)
		top += 50
		hSliderBar2 = xc.XSliderBar_Create(left, top, width, 20, w.Handle)
		top += 50
		hSliderBar3 = xc.XSliderBar_Create(left, top, width, 20, w.Handle)
		top += 50
		list_xcgui = append(list_xcgui, hSliderBar1)
		list_xcgui = append(list_xcgui, hSliderBar2)
		list_xcgui = append(list_xcgui, hSliderBar3)
		xc.XEle_SetBkInfo(hSliderBar1, strBackground)
		xc.XEle_SetBkInfo(hSliderBar2, strBackground)
		xc.XEle_SetBkInfo(hSliderBar3, strBackground)
		xc.XSliderBar_SetButtonWidth(hSliderBar1, 20)
		xc.XSliderBar_SetButtonWidth(hSliderBar2, 20)
		xc.XSliderBar_SetButtonWidth(hSliderBar3, 20)
		xc.XEle_SetBkInfo(xc.XSliderBar_GetButton(hSliderBar1), "{99:1.9.9;98:16(0)32(0)64(0);6:2(15)20(1)21(3)26(1)22(-25024)23(255);}")
		xc.XEle_SetBkInfo(xc.XSliderBar_GetButton(hSliderBar2), "{99:1.9.9;98:16(0)32(0)64(0);6:2(15)20(1)21(3)26(1)22(-25024)23(255);}")
		xc.XEle_SetBkInfo(xc.XSliderBar_GetButton(hSliderBar3), "{99:1.9.9;98:16(0)32(0)64(0);6:2(15)20(1)21(3)26(0)22(-1)23(255)10(1)7(2)11(3)16(1)12(-25024)13(255);}")

		hImage := xc.XImage_LoadMemory(imgSliderBar)
		xc.XImage_SetDrawTypeAdaptive(hImage, 5, 5, 5, 5)
		xc.XSliderBar_SetImageLoad(hSliderBar1, hImage)
		xc.XSliderBar_SetImageLoad(hSliderBar2, hImage)

		hAnimation := xc.XAnima_Create(0, 0)
		list_animation = append(list_animation, hAnimation)
		xc.XAnimaItem_SetCallback(xc.XAnima_DelayEx(hAnimation, 2000, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, true), OnAnimationItem_22)
		xc.XAnima_Run(hAnimation, w.Handle)
	}

	// 进度条
	{
		hProgBar1 = xc.XProgBar_Create(left, top, width, 20, w.Handle)
		top += 50
		hProgBar2 = xc.XProgBar_Create(left, top, width, 20, w.Handle)
		top += 50
		hProgBar3 = xc.XProgBar_Create(left, top, width, 20, w.Handle)
		top += 50
		list_xcgui = append(list_xcgui, hProgBar1)
		list_xcgui = append(list_xcgui, hProgBar2)
		list_xcgui = append(list_xcgui, hProgBar3)
		xc.XEle_SetBkInfo(hProgBar1, "{99:1.9.9;98:16(0);5:2(15)20(1)21(3)26(1)22(-3618616)23(255)9(10,10,10,10);}")
		xc.XEle_SetBkInfo(hProgBar2, "{99:1.9.9;98:16(0);5:2(15)20(1)21(3)26(1)22(-3618616)23(255)9(10,10,10,10);}")

		hImage := xc.XImage_LoadMemory(imgProgressBar)
		xc.XImage_SetDrawTypeAdaptive(hImage, 10, 0, 10, 0)
		xc.XProgBar_SetImageLoad(hProgBar1, hImage)
		xc.XProgBar_SetImageLoad(hProgBar2, hImage)

		xc.XEle_SetTextColor(hProgBar1, xc.RGBA(255, 255, 255, 255))
		xc.XEle_SetTextColor(hProgBar2, xc.RGBA(255, 255, 255, 255))
	}
	w.Redraw(false)
	return 0
}

func OnAnimationItem_22(hAnimation int, pos float32) int {
	if xc.XC_IsHELE(hSliderBar1) {
		xc.XSliderBar_SetPos(hSliderBar1, int32(100.0*pos+0.5))
	}
	if xc.XC_IsHELE(hSliderBar2) {
		xc.XSliderBar_SetPos(hSliderBar2, int32(80.0*pos+0.5))
	}
	if xc.XC_IsHELE(hSliderBar3) {
		xc.XSliderBar_SetPos(hSliderBar3, int32(50.0*pos+0.5))
	}

	if hProgBar1 > 0 {
		xc.XProgBar_SetPos(hProgBar1, int32(100.0*pos+0.5))
	}
	if hProgBar2 > 0 {
		xc.XProgBar_SetPos(hProgBar2, int32(80.0*pos+0.5))
	}
	if hProgBar3 > 0 {
		xc.XProgBar_SetPos(hProgBar3, int32(50.0*pos+0.5))
	}
	return 0
}

// 23.焦点追踪
func OnBtnClick23(hEle int, pbHandled *bool) int {
	ReleaseAnimation()

	var left int32 = 150
	var top int32 = 80
	var width int32 = 100
	var height int32 = 60

	// 图标按钮 切换
	{
		pFocus := NewCFocusTraceButton_Line()
		list_object = append(list_object, pFocus)
		pFocus.FocusOffset = 20
		pFocus.ChangeColor = true
		pFocus.CreatePane2(left, top-20, 600, height+40, w.Handle)
		left += 50
		pFocus.CreateFocusEle(left, top+height-2, width, 2, xc.RGBA(0, 162, 232, 255), w.Handle)

		pFocus.CreateButton2(left, top, width, height, "Button", xc.XSvg_LoadString(svg1), xc.RGBA(171, 72, 188, 255), xc.RGBA(247, 238, 248, 255), w.Handle)
		left += width
		pFocus.CreateButton2(left, top, width, height, "Button", xc.XSvg_LoadString(svg2), xc.RGBA(254, 167, 38, 255), xc.RGBA(253, 244, 232, 255), w.Handle)
		left += width
		pFocus.CreateButton2(left, top, width, height, "Button", xc.XSvg_LoadString(svg3), xc.RGBA(38, 166, 154, 255), xc.RGBA(236, 246, 245, 255), w.Handle)
		left += width
	}

	left = 150
	top += 100
	width = 90
	height = 40
	// 图标按钮 切换2 导航条
	{
		pFocus := NewCFocusTraceButton_Line()
		list_object = append(list_object, pFocus)
		pFocus.CreatePane(left, top, 600, height, w.Handle)
		left += 50
		pFocus.CreateFocusEle(left, top+height-3, width, 3, xc.RGBA(0, 162, 232, 255), w.Handle)

		pFocus.CreateButton(left, top, width, height, "Button", xc.XSvg_LoadString(svg1), w.Handle)
		left += width
		pFocus.CreateButton(left, top, width, height, "Button", xc.XSvg_LoadString(svg2), w.Handle)
		left += width
		pFocus.CreateButton(left, top, width, height, "Button", xc.XSvg_LoadString(svg3), w.Handle)
		left += width
	}

	left = 150
	top += 80
	width = 150
	height = 30
	// 编辑框 焦点边
	{
		pFocus := NewCFocusTraceEdit_Line()
		list_object = append(list_object, pFocus)
		pFocus.CreatePane(left, top, 600, height+40, w.Handle)
		left += 50
		top += 20
		pFocus.CreateEdit(left, top, width, height, w.Handle)
		left += width + 20
		pFocus.CreateEdit(left, top, width, height, w.Handle)
		left += width + 20
		pFocus.CreateEdit(left, top, width, height, w.Handle)
		left += width + 20
	}

	left = 150
	top += 90
	width = 150
	height = 30
	// 编辑框 焦点矩形
	{
		pFocus := NewCFocusTraceEdit_Border()
		list_object = append(list_object, pFocus)
		pFocus.CreatePane(left, top, 600, height+40, w.Handle)
		left += 50
		top += 20
		pFocus.CreateFocusEle(left, top, width, height, w.Handle)

		pFocus.CreateEdit(left, top, width, height, w.Handle)
		left += width + 20
		pFocus.CreateEdit(left, top, width, height, w.Handle)
		left += width + 20
		pFocus.CreateEdit(left, top, width, height, w.Handle)
		left += width + 20
	}

	w.Redraw(false)
	return 0
}

// CFocusTraceButton_Line 焦点追踪 - 线
type CFocusTraceButton_Line struct {
	HShapeRect  int
	HFocusEle   int
	ChangeColor bool
	FocusOffset int32
	List        []int
}

// NewCFocusTraceButton_Line 构造函数
func NewCFocusTraceButton_Line() *CFocusTraceButton_Line {
	return &CFocusTraceButton_Line{
		FocusOffset: -1,
		List:        make([]int, 0),
	}
}

// Release 释放资源
func (c *CFocusTraceButton_Line) Release() {
	for _, hxcgui := range c.List {
		ReleaseObject(hxcgui)
	}
}

func (c *CFocusTraceButton_Line) CreatePane(x, y, width, height int32, hParent int) {
	c.HShapeRect = xc.XShapeRect_Create(x, y, width, height, hParent)
	c.List = append(c.List, c.HShapeRect)
	xc.XShapeRect_SetFillColor(c.HShapeRect, xc.RGBA(156, 39, 176, 255))
	xc.XShapeRect_EnableBorder(c.HShapeRect, false)
}

func (c *CFocusTraceButton_Line) CreatePane2(x, y, width, height int32, hParent int) {
	c.HShapeRect = xc.XShapeRect_Create(x, y, width, height, hParent)
	c.List = append(c.List, c.HShapeRect)
	xc.XShapeRect_SetBorderColor(c.HShapeRect, xc.RGBA(200, 200, 200, 255))
	xc.XShapeRect_EnableFill(c.HShapeRect, false)
}

func (c *CFocusTraceButton_Line) CreateFocusEle(x, y, width, height int32, color uint32, hParent int) {
	c.HFocusEle = xc.XEle_Create(x, y, width, height, hParent)
	c.List = append(c.List, c.HFocusEle)
	xc.XEle_EnableTopmost(c.HFocusEle, true)
	xc.XEle_AddBkFill(c.HFocusEle, xcc.Element_State_Flag_Leave, color)
}

func (c *CFocusTraceButton_Line) CreateButton(x, y, width, height int32, pName string, hSvg, hParent int) int {
	hButton := xc.XBtn_Create(x, y, width, height, pName, hParent)
	c.List = append(c.List, hButton)

	str := "{99:1.9.9;98:16(0)32(1)64(2);5:2(15)20(1)21(3)26(1)22(-5232740)23(255);5:2(15)20(1)21(3)26(1)22(-4702042)23(255);5:2(15)20(1)21(3)26(1)22(-4303953)23(255);}"
	xc.XEle_SetBkInfo(hButton, str)
	xc.XEle_SetTextColor(hButton, xc.RGBA(255, 255, 255, 255))
	xc.XSvg_SetSize(hSvg, 24, 24)
	xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(255, 255, 255, 255), true)

	hImage := xc.XImage_LoadSvg(hSvg)
	xc.XBtn_SetIcon(hButton, hImage)

	btn := widget.NewButtonByHandle(hButton)
	btn.AddEvent_BnClick(c.OnBtnClick2)
	return hButton
}

func (c *CFocusTraceButton_Line) CreateButton2(x, y, width, height int32, pName string, hSvg int, color, colorBack uint32, hParent int) {
	hButton := xc.XBtn_Create(x, y, width, height, pName, hParent)
	c.List = append(c.List, hButton)

	xc.XBtn_SetIconAlign(hButton, xcc.Button_Icon_Align_Top)
	xc.XSvg_SetUserFillColor(hSvg, color, true)
	xc.XEle_SetTextColor(hButton, color)
	xc.XEle_AddBkFill(hButton, xcc.Element_State_Flag_Stay, colorBack)
	xc.XEle_AddBkFill(hButton, xcc.Element_State_Flag_Down, colorBack)

	xc.XSvg_SetSize(hSvg, 32, 32)
	hImage := xc.XImage_LoadSvg(hSvg)
	xc.XBtn_SetIcon(hButton, hImage)

	xc.XEle_AddBkFill(hButton, xcc.Element_State_Flag_Leave, xc.RGBA(255, 255, 255, 255))

	btn := widget.NewButtonByHandle(hButton)
	btn.AddEvent_BnClick(c.OnBtnClick2)
}

func (c *CFocusTraceButton_Line) OnBtnClick(hButton int, pbHandled *bool) int {
	var x, y int32
	xc.XEle_GetPosition(hButton, &x, &y)

	hAnimation := xc.XAnima_Create(c.HFocusEle, 1)
	xc.XAnimaMove_SetFlag(xc.XAnima_Move(hAnimation, 400, float32(x), float32(y), 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false), xcc.Animation_Move_X)
	xc.XAnima_Run(hAnimation, c.HShapeRect)
	return 0
}

func (c *CFocusTraceButton_Line) OnBtnClick2(hButton int, pbHandled *bool) int {
	if c.ChangeColor {
		xc.XEle_ClearBkInfo(c.HFocusEle)
		xc.XEle_AddBkFill(c.HFocusEle, xcc.Element_State_Flag_Leave, xc.XEle_GetTextColor(hButton))
	}

	var rect xc.RECT
	xc.XEle_GetRect(hButton, &rect)

	hAnimation := xc.XAnima_Create(c.HFocusEle, 1)
	if c.FocusOffset == -1 {
		xc.XEle_SetWidth(c.HFocusEle, rect.Right-rect.Left)
		xc.XAnimaMove_SetFlag(xc.XAnima_Move(hAnimation, 400, float32(rect.Left), 0, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false), xcc.Animation_Move_X)
	} else {
		xc.XEle_SetWidth(c.HFocusEle, (rect.Right-rect.Left)-c.FocusOffset-c.FocusOffset)
		xc.XAnimaMove_SetFlag(xc.XAnima_Move(hAnimation, 400, float32(rect.Left+c.FocusOffset), 0, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false), xcc.Animation_Move_X)
	}
	xc.XAnima_Run(hAnimation, c.HShapeRect)

	hImage := xc.XBtn_GetIcon(hButton, 0)
	hSvg := xc.XImage_GetSvg(hImage)

	hAnimation = xc.XAnima_Create(hSvg, 1)
	index := rand.Intn(3)

	switch index {
	case 0:
		xc.XAnima_Scale(hAnimation, 600, 1.5, 1.5, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, true)
	case 1:
		xc.XAnima_Rotate(hAnimation, 600, 360, 1, 0, false)
	case 2:
		xc.XAnima_Rotate(hAnimation, 200, -45, 1, 0, false)
		xc.XAnima_Rotate(hAnimation, 300, 45, 1, 0, true)
		xc.XAnima_Rotate(hAnimation, 200, 0, 1, 0, false)
	}
	xc.XAnima_Run(hAnimation, hButton)
	return 0
}

// CFocusTraceEdit_Line 焦点追踪 - 线
type CFocusTraceEdit_Line struct {
	hShapeRect int
	list       []int
}

func NewCFocusTraceEdit_Line() *CFocusTraceEdit_Line {
	return &CFocusTraceEdit_Line{
		list: make([]int, 0),
	}
}

func (c *CFocusTraceEdit_Line) Release() {
	for _, hxcgui := range c.list {
		ReleaseObject(hxcgui)
	}
}

func (c *CFocusTraceEdit_Line) CreatePane(x, y, width, height int32, hParent int) {
	c.hShapeRect = xc.XShapeRect_Create(x, y, width, height, hParent)
	c.list = append(c.list, c.hShapeRect)
	xc.XShapeRect_SetBorderColor(c.hShapeRect, xc.RGBA(200, 200, 200, 255))
	xc.XShapeRect_EnableFill(c.hShapeRect, false)
}

func (c *CFocusTraceEdit_Line) CreateEdit(x, y, width, height int32, hParent int) {
	hEdit := xc.XEdit_Create(x, y, width, height, hParent)
	c.list = append(c.list, hEdit)

	xc.XEle_EnableDrawFocus(hEdit, false)

	xc.XEle_AddBkFill(hEdit, xcc.Element_State_Flag_Leave, xc.RGBA(235, 235, 235, 255))
	xc.XEle_AddBkFill(hEdit, xcc.Element_State_Flag_Stay, xc.RGBA(235, 235, 235, 255))

	edit := widget.NewEditByHandle(hEdit)
	edit.AddEvent_SetFocus(c.OnSetFocus)
	edit.AddEvent_KillFocus(c.OnKillFocus)
}

// 获得焦点
func (c *CFocusTraceEdit_Line) OnSetFocus(hEdit int, pbHandled *bool) int {
	var rect xc.RECT
	xc.XEle_GetRect(hEdit, &rect)

	hParent := xc.XWidget_GetParent(hEdit)
	// 创建焦点指示元素，位置在编辑框底部中间
	hFocusEle := xc.XEle_Create(rect.Left+(rect.Right-rect.Left)/2, rect.Bottom-2, 1, 2, hParent)

	// 随机选择颜色
	color := xc.RGBA(200, 0, 0, 255)
	switch rand.Intn(4) {
	case 0:
		color = xc.RGBA(171, 72, 188, 255)
	case 1:
		color = xc.RGBA(254, 167, 38, 255)
	case 2:
		color = xc.RGBA(38, 166, 154, 255)
	}

	xc.XEle_AddBkFill(hFocusEle, xcc.Element_State_Flag_Leave, color)

	// 将 hFocusEle 存储在 Edit 的 UserData 中
	xc.XEle_SetUserData(hEdit, hFocusEle)

	// 创建动画：从中间向两边扩展
	hAnimation := xc.XAnima_Create(hFocusEle, 1)
	xc.XAnima_ScaleSize(hAnimation, 400, float32(rect.Right-rect.Left), 2, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
	xc.XAnima_Run(hAnimation, c.hShapeRect)
	return 0
}

// 失去焦点
func (c *CFocusTraceEdit_Line) OnKillFocus(hEdit int, pbHandled *bool) int {
	// 从 UserData 中取出之前保存的焦点元素句柄
	hFocusEle := xc.XEle_GetUserData(hEdit)

	if hFocusEle != 0 {
		hAnimation := xc.XAnima_Create(hFocusEle, 1)
		// 动画：缩小至宽度0
		xc.XAnima_ScaleSize(hAnimation, 400, 0, 2, 1, xcc.Ease_Flag_Cubic|xcc.Ease_Flag_Out, false)
		xc.XAnima_DestroyObjectUI(hAnimation, 0)
		xc.XAnima_Run(hAnimation, c.hShapeRect)
	}
	return 0
}

// CFocusTraceEdit_Border 焦点追踪- 边
type CFocusTraceEdit_Border struct {
	hShapeRect int
	hFocusEle  int
	list       []int
}

// NewCFocusTraceEdit_Border 构造函数
func NewCFocusTraceEdit_Border() *CFocusTraceEdit_Border {
	return &CFocusTraceEdit_Border{
		list: make([]int, 0),
	}
}

// Release 释放资源
func (c *CFocusTraceEdit_Border) Release() {
	for _, hxcgui := range c.list {
		ReleaseObject(hxcgui)
	}
}

func (c *CFocusTraceEdit_Border) CreatePane(x, y, width, height int32, hParent int) {
	c.hShapeRect = xc.XShapeRect_Create(x, y, width, height, hParent)
	c.list = append(c.list, c.hShapeRect)
	xc.XShapeRect_SetBorderColor(c.hShapeRect, xc.RGBA(200, 200, 200, 255))
	xc.XShapeRect_EnableFill(c.hShapeRect, false)
}

func (c *CFocusTraceEdit_Border) CreateFocusEle(x, y, width, height int32, hParent int) {
	c.hFocusEle = xc.XEle_Create(x, y, width, height, hParent)
	c.list = append(c.list, c.hFocusEle)

	xc.XEle_EnableTopmost(c.hFocusEle, true)
	xc.XEle_EnableMouseThrough(c.hFocusEle, true)
	xc.XEle_EnableBkTransparent(c.hFocusEle, true)
	xc.XEle_AddBkBorder(c.hFocusEle, xcc.Element_State_Flag_Leave, xc.RGBA(254, 167, 38, 255), 2)
}

func (c *CFocusTraceEdit_Border) CreateEdit(x, y, width, height int32, hParent int) {
	hEdit := xc.XEdit_Create(x, y, width, height, hParent)
	c.list = append(c.list, hEdit)

	xc.XEle_EnableDrawFocus(hEdit, false)

	xc.XEle_AddBkFill(hEdit, xcc.Element_State_Flag_Leave, xc.RGBA(245, 245, 245, 255))
	xc.XEle_AddBkFill(hEdit, xcc.Element_State_Flag_Stay, xc.RGBA(245, 245, 245, 255))
	xc.XEle_AddBkBorder(hEdit, xcc.Element_State_Flag_Leave, xc.RGBA(220, 220, 220, 255), 1)
	xc.XEle_AddBkBorder(hEdit, xcc.Element_State_Flag_Stay, xc.RGBA(200, 200, 200, 255), 1)

	edit := widget.NewEditByHandle(hEdit)
	edit.AddEvent_SetFocus(c.OnSetFocus)
}

// 获得焦点
func (c *CFocusTraceEdit_Border) OnSetFocus(hEdit int, pbHandled *bool) int {
	var rect xc.RECT
	xc.XEle_GetRect(hEdit, &rect)

	// 随机选择颜色
	color := xc.RGBA(200, 0, 0, 255)
	switch rand.Intn(4) {
	case 0:
		color = xc.RGBA(171, 72, 188, 255)
	case 1:
		color = xc.RGBA(254, 167, 38, 255)
	case 2:
		color = xc.RGBA(38, 166, 154, 255)
	}

	// 更新焦点元素样式
	xc.XEle_ClearBkInfo(c.hFocusEle)
	xc.XEle_AddBkBorder(c.hFocusEle, xcc.Element_State_Flag_Leave, color, 2)

	// 创建动画：移动焦点框到编辑框位置
	hAnimation := xc.XAnima_Create(c.hFocusEle, 1)
	xc.XAnima_Move(hAnimation, 400, float32(rect.Left), float32(rect.Top), 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
	xc.XAnima_Run(hAnimation, c.hShapeRect)
	return 0
}

var (
	hTabPage_Background int // 页面背景,裁剪
	hTabPage_cur        int // 当前页面
)

// 24.页面切换 滑动
func OnBtnClick24(hEle int, pbHandled *bool) int {
	ReleaseAnimation()

	var left int32 = 150
	var top int32 = 80
	var width int32 = 90
	var height int32 = 40

	pFocus := NewCFocusTraceButton_Line()
	list_object = append(list_object, pFocus)
	pFocus.CreatePane(left, top, 600, height, w.Handle)
	left += 50
	pFocus.CreateFocusEle(left, top+height-3, width, 3, xc.RGBA(255, 255, 255, 255), w.Handle)

	hButton1 := pFocus.CreateButton(left, top, width, height, "Button1", xc.XSvg_LoadString(svg1), w.Handle)
	left += width
	hButton2 := pFocus.CreateButton(left, top, width, height, "Button2", xc.XSvg_LoadString(svg2), w.Handle)
	left += width
	hButton3 := pFocus.CreateButton(left, top, width, height, "Button3", xc.XSvg_LoadString(svg3), w.Handle)
	left += width

	left = 150
	top += 40
	// 作为背景, 对动画区域裁剪
	hTabPage_Background = xc.XEle_Create(left, top, 600, 300, w.Handle)
	list_xcgui = append(list_xcgui, hTabPage_Background)
	xc.XEle_SetUserData(hButton1, 1)
	xc.XEle_SetUserData(hButton2, 2)
	xc.XEle_SetUserData(hButton3, 3)

	widget.NewButtonByHandle(hButton1).AddEvent_BnClick(OnBtnClick24_1, true)
	widget.NewButtonByHandle(hButton2).AddEvent_BnClick(OnBtnClick24_1, true)
	widget.NewButtonByHandle(hButton3).AddEvent_BnClick(OnBtnClick24_1, true)

	OnBtnClick24_1(hButton1, nil)

	w.Redraw(false)
	return 0
}

func OnBtnClick24_1(hButton int, pbHandled *bool) int {
	var bMoveLeft = false

	var id_old = 0
	var id_new = xc.XEle_GetUserData(hButton)

	// 检查当前页面是否存在且是元素类型
	if hTabPage_cur > 0 && xc.XC_IsHELE(hTabPage_cur) {
		id_old = xc.XEle_GetUserData(hTabPage_cur)
		if id_new == id_old {
			return 0
		}
	}

	if id_old < id_new {
		bMoveLeft = true
	}

	var width int32 = 600
	var height int32 = 300

	// 处理旧页面的退出动画
	if hTabPage_cur > 0 {
		if xc.XC_IsHELE(hTabPage_cur) {
			hAnimation := xc.XAnima_Create(hTabPage_cur, 1)
			var moveX int32
			if bMoveLeft {
				moveX = 1 - width
			} else {
				moveX = width
			}
			xc.XAnima_Move(hAnimation, 500, float32(moveX), 0, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
			xc.XAnima_DestroyObjectUI(hAnimation, 0)
			xc.XAnima_Run(hAnimation, hTabPage_Background)
		}
		hTabPage_cur = 0
	}

	// 计算新页面进入的起始位置
	var left int32
	if bMoveLeft {
		left = width + 10
	} else {
		left = 1 - width - 10
	}

	// 创建新页面
	hTabPage_cur = xc.XEle_Create(left, 0, width, height, hTabPage_Background)

	// 根据ID设置不同的内容和背景色
	switch id_new {
	case 1:
		xc.XEle_AddBkFill(hTabPage_cur, xcc.Element_State_Flag_Leave, xc.RGBA(34, 177, 76, 255))
		xc.XBtn_Create(100, 100, 100, 30, "我是页面 1", hTabPage_cur)
	case 2:
		xc.XEle_AddBkFill(hTabPage_cur, xcc.Element_State_Flag_Leave, xc.RGBA(254, 167, 38, 255))
		xc.XBtn_Create(100, 100, 100, 30, "我是页面 2", hTabPage_cur)
	case 3:
		xc.XEle_AddBkFill(hTabPage_cur, xcc.Element_State_Flag_Leave, xc.RGBA(38, 166, 154, 255))
		xc.XBtn_Create(100, 100, 100, 30, "我是页面 3", hTabPage_cur)
	}

	// 新页面入场动画
	if hTabPage_cur > 0 {
		xc.XEle_SetUserData(hTabPage_cur, id_new)

		hAnimation := xc.XAnima_Create(hTabPage_cur, 1)
		xc.XAnima_Move(hAnimation, 500, 0, 0, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
		xc.XAnima_Run(hAnimation, hTabPage_Background)
	}

	xc.XEle_Redraw(hTabPage_Background, false)
	return 0
}

// 25.折叠面板
func OnBtnClick25(hEle int, pbHandled *bool) int {
	ReleaseAnimation()

	var left int32 = 150
	var top int32 = 50
	var width int32 = 500
	var height int32 = 500

	pExpandGroup := NewCExpandGroup()
	list_object = append(list_object, pExpandGroup)
	pExpandGroup.CreateLayout(left, top, width, height, w.Handle)

	hButton := pExpandGroup.CreatePanel("折叠1", xc.XSvg_LoadString(svg1))
	pExpandGroup.CreatePanel("折叠2", xc.XSvg_LoadString(svg2))
	pExpandGroup.CreatePanel("Button3", xc.XSvg_LoadString(svg3))

	xc.XEle_AdjustLayout(pExpandGroup.HLayout, 0)
	pExpandGroup.OnBtnClick(hButton, nil)
	xc.XBtn_SetCheck(hButton, true)

	w.Redraw()
	return 0
}

type panel_info_ struct {
	hPanel     int
	hButton    int
	height     int32
	list_temp  []int
	list_temp2 []int
	list_temp3 []int
	list_temp4 []int
}

// CExpandGroup 展开收缩面板
type CExpandGroup struct {
	m_list  []*panel_info_
	HLayout int
}

// NewCExpandGroup 构造函数
func NewCExpandGroup() *CExpandGroup {
	return &CExpandGroup{
		m_list: make([]*panel_info_, 0),
	}
}

// Release 释放资源
func (c *CExpandGroup) Release() {
	c.m_list = nil

	// 销毁布局元素
	if c.HLayout != 0 && xc.XC_IsHELE(c.HLayout) {
		xc.XEle_Destroy(c.HLayout)
	}
	c.HLayout = 0
}

func (c *CExpandGroup) CreateLayout(x, y, width, height int32, hParent int) int {
	c.HLayout = xc.XLayout_Create(x, y, width, height, hParent)
	xc.XLayoutBox_EnableHorizon(c.HLayout, false)
	xc.XEle_AddBkBorder(c.HLayout, xcc.Element_State_Flag_Leave, xc.RGBA(200, 200, 200, 255), 1)
	xc.XEle_SetPadding(c.HLayout, 1, 0, 1, 1)
	xc.XWidget_LayoutItem_SetHeight(c.HLayout, xcc.Layout_Size_Auto, 0)
	return c.HLayout
}

func (c *CExpandGroup) CreatePanel(pName string, hSvg int) int {
	hButton := xc.XBtn_Create(0, 0, 100, 40, pName, c.HLayout)
	xc.XObj_SetTypeEx(hButton, xcc.Button_Type_Check)
	xc.XWidget_LayoutItem_SetWidth(hButton, xcc.Layout_Size_Fill, 0)
	xc.XEle_EnableDrawBorder(hButton, false)
	xc.XEle_EnableDrawFocus(hButton, false)
	xc.XBtn_SetTextAlign(hButton, xcc.TextAlignFlag_Left|xcc.TextAlignFlag_Vcenter)
	xc.XEle_SetPadding(hButton, 20, 0, 0, 0)
	xc.XEle_SetBkInfo(hButton, "{99:1.9.9;98:272(0)288(0)320(0)128(0);5:2(7)3(0,0,0,1)20(1)21(3)26(0)22(-3618616)23(255);}")
	xc.XEle_EnableBkTransparent(hButton, true)

	xc.XSvg_SetSize(hSvg, 24, 24)
	xc.XBtn_SetIcon(hButton, xc.XImage_LoadSvg(hSvg))

	hCur := wapi.LoadImageW(0, wapi.IDC_HAND, wapi.IMAGE_CURSOR, 0, 0, wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
	if hCur != 0 {
		xc.XEle_SetCursor(hButton, hCur)
	}

	hPanel := xc.XEle_Create(0, 0, 100, 0, c.HLayout)
	xc.XWidget_LayoutItem_SetWidth(hPanel, xcc.Layout_Size_Fill, 0)
	xc.XEle_EnableCanvas(hPanel, false)
	xc.XEle_EnableDrawBorder(hPanel, false)

	// 创建 panel_info_
	pInfo := &panel_info_{
		hPanel:     hPanel,
		hButton:    hButton,
		height:     200,
		list_temp:  make([]int, 0),
		list_temp2: make([]int, 0),
		list_temp3: make([]int, 0),
		list_temp4: make([]int, 0),
	}
	c.m_list = append(c.m_list, pInfo)

	var left int32 = 20
	var top int32 = 10
	var width int32 = 500

	// 第一组文本
	hText := xc.XShapeText_Create(left+width, top, 200, 20, "炫彩界面库 3.3.0", pInfo.hPanel)
	pInfo.list_temp = append(pInfo.list_temp, hText)
	xc.XShapeText_SetTextColor(hText, xc.RGBA(80, 80, 80, 255))

	top += 25
	// 第一组元素
	for i := 0; i < 3; i++ {
		hEle := xc.XEle_Create(left+width, top, 100, 50, pInfo.hPanel)
		pInfo.list_temp2 = append(pInfo.list_temp2, hEle)
		switch i {
		case 0:
			xc.XEle_AddBkFill(hEle, xcc.Element_State_Flag_Leave, xc.RGBA(213, 162, 221, 255))
		case 1:
			xc.XEle_AddBkFill(hEle, xcc.Element_State_Flag_Leave, xc.RGBA(255, 221, 170, 255))
		case 2:
			xc.XEle_AddBkFill(hEle, xcc.Element_State_Flag_Leave, xc.RGBA(151, 232, 223, 255))
		}
		left += 130
	}

	left = 20
	top += 70

	// 第二组文本
	hText = xc.XShapeText_Create(left+width, top, 200, 20, "炫彩界面库 3.3.1", pInfo.hPanel)
	pInfo.list_temp3 = append(pInfo.list_temp3, hText)
	xc.XShapeText_SetTextColor(hText, xc.RGBA(80, 80, 80, 255))

	top += 25
	// 第二组元素
	for i := 0; i < 3; i++ {
		hEle := xc.XEle_Create(left+width, top, 100, 50, pInfo.hPanel)
		pInfo.list_temp4 = append(pInfo.list_temp4, hEle)
		switch i {
		case 0:
			xc.XEle_AddBkFill(hEle, xcc.Element_State_Flag_Leave, xc.RGBA(213, 162, 221, 255))
		case 1:
			xc.XEle_AddBkFill(hEle, xcc.Element_State_Flag_Leave, xc.RGBA(255, 221, 170, 255))
		case 2:
			xc.XEle_AddBkFill(hEle, xcc.Element_State_Flag_Leave, xc.RGBA(151, 232, 223, 255))
		}
		left += 130
	}

	// 注册按钮点击事件
	xc.XEle_RegEventC1(hButton, xcc.XE_BNCLICK, func(hEle int, pbHandled *bool) int {
		return c.OnBtnClick(hEle, pbHandled)
	})

	return hButton
}

func (c *CExpandGroup) OnBtnClick(hButton int, pbHandled *bool) int {
	var pInfo *panel_info_ = nil
	// 查找对应的 panel_info
	for _, var_ := range c.m_list {
		if var_.hButton == hButton {
			pInfo = var_
			break
		}
	}
	if pInfo == nil {
		return 0
	}

	hGroup := xc.XAnimaGroup_Create(1)
	hAnimation := xc.XAnima_Create(pInfo.hPanel, 1)
	xc.XAnimaGroup_AddItem(hGroup, hAnimation)

	width := xc.XEle_GetWidth(pInfo.hPanel)
	bExpand := !xc.XBtn_IsCheck(hButton)

	if len(pInfo.list_temp) > 0 {
		hText := pInfo.list_temp[0]
		c.Animation25_Move(hGroup, hText, common.Choose(bExpand, 0, 600), common.Choose(bExpand, -width, width), true)
	}

	count := len(pInfo.list_temp2)
	for i := 0; i < count; i++ {
		hEle := pInfo.list_temp2[i]
		c.Animation25_Move(hGroup, hEle, common.Choose(bExpand, 200, 400), common.Choose(bExpand, -width, width), true)
	}

	if len(pInfo.list_temp3) > 0 {
		hText := pInfo.list_temp3[0]
		c.Animation25_Move(hGroup, hText, common.Choose(bExpand, 400, 200), common.Choose(bExpand, -width, width), true)
	}

	count = len(pInfo.list_temp4)
	for i := 0; i < count; i++ {
		hEle := pInfo.list_temp4[i]
		c.Animation25_Move(hGroup, hEle, common.Choose(bExpand, 600, 0), common.Choose(bExpand, -width, width), true)
	}

	if bExpand {
		xc.XAnima_LayoutHeight(hAnimation, 400, xcc.Layout_Size_Fixed, float32(pInfo.height), 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, false)
	} else {
		xc.XAnima_Delay(hAnimation, 400)
		xc.XAnima_LayoutHeight(hAnimation, 400, xcc.Layout_Size_Fixed, 0, 1, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In, false)
	}

	xc.XAnima_Run(hGroup, w.Handle)
	return 0
}

func (c *CExpandGroup) Animation25_Move(hGroup int, hObjectUI int, delay int, offsetx int32, bDestroy bool) {
	var left, top int32 = 0, 0
	if xc.XC_IsHELE(hObjectUI) {
		xc.XEle_GetPosition(hObjectUI, &left, &top)
	} else {
		xc.XShape_GetPosition(hObjectUI, &left, &top)
	}

	hAnimationMove := xc.XAnima_Create(hObjectUI, 1)
	xc.XAnimaGroup_AddItem(hGroup, hAnimationMove)
	if delay > 0 {
		xc.XAnima_Delay(hAnimationMove, float32(delay))
	}

	easeFlag := common.Choose(offsetx < 0, xcc.Ease_Flag_Quad|xcc.Ease_Flag_Out, xcc.Ease_Flag_Quad|xcc.Ease_Flag_In)
	xc.XAnima_Move(hAnimationMove, 300, float32(left+offsetx), float32(top), 1, easeFlag, false)
}
