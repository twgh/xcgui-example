// 美化开关按钮.
// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BeautifyToggleButton 美化开关按钮.
//
// 继承 widget.Button, 通过自绘实现美化开关外观.
// 开: 蓝色轨道(#006CD1) + 右侧白色圆形滑块.
// 关: 灰色轨道(#C4C4C4) + 左侧白色圆形滑块.
type BeautifyToggleButton struct {
	*widget.Button
	trackColorOn  uint32 // 开状态轨道颜色
	trackColorOff uint32 // 关状态轨道颜色
	thumbColor    uint32 // 滑块颜色
	disabledColor uint32 // 禁用状态轨道颜色
}

// NewBeautifyToggleButton 创建美化开关按钮. 默认尺寸 44x22.
//
// x, y: 坐标;
//
// hParent: 父句柄.
func NewBeautifyToggleButton(x, y int32, hParent int) *BeautifyToggleButton {
	btn := widget.NewButton(x, y, 44, 22, "", hParent)
	if btn == nil {
		return nil
	}
	t := &BeautifyToggleButton{
		Button:        btn,
		trackColorOn:  xc.RGBA(0, 108, 209, 255),   // #006CD1
		trackColorOff: xc.RGBA(196, 196, 196, 255), // #C4C4C4
		thumbColor:    xc.RGBA(255, 255, 255, 255), // 白色
		disabledColor: xc.RGBA(220, 220, 220, 255),
	}
	// 设置按钮类型为多选按钮
	t.SetTypeEx(xcc.Button_Type_Check)
	// 启用背景透明, 由自绘接管全部绘制
	t.EnableBkTransparent(true)
	// 注册自绘事件
	t.AddEvent_Paint(t.onPaint)
	return t
}

// SetTrackColorOn 设置开状态轨道颜色.
func (t *BeautifyToggleButton) SetTrackColorOn(color uint32) {
	t.trackColorOn = color
}

// GetTrackColorOn 获取开状态轨道颜色.
func (t *BeautifyToggleButton) GetTrackColorOn() uint32 {
	return t.trackColorOn
}

// SetTrackColorOff 设置关状态轨道颜色.
func (t *BeautifyToggleButton) SetTrackColorOff(color uint32) {
	t.trackColorOff = color
}

// GetTrackColorOff 获取关状态轨道颜色.
func (t *BeautifyToggleButton) GetTrackColorOff() uint32 {
	return t.trackColorOff
}

// SetThumbColor 设置滑块颜色.
func (t *BeautifyToggleButton) SetThumbColor(color uint32) {
	t.thumbColor = color
}

// GetThumbColor 获取滑块颜色.
func (t *BeautifyToggleButton) GetThumbColor() uint32 {
	return t.thumbColor
}

// SetDisabledColor 设置禁用状态轨道颜色.
func (t *BeautifyToggleButton) SetDisabledColor(color uint32) {
	t.disabledColor = color
}

// GetDisabledColor 获取禁用状态轨道颜色.
func (t *BeautifyToggleButton) GetDisabledColor() uint32 {
	return t.disabledColor
}

// onPaint 自绘事件: 绘制轨道背景 + 圆形滑块.
func (t *BeautifyToggleButton) onPaint(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true // 拦截元素原本的绘制
	draw := drawx.NewByHandle(hDraw)
	if draw == nil {
		return 0
	}
	draw.EnableSmoothingMode(true)

	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)

	h := rc.Bottom - rc.Top
	margin := int32(2)        // 滑块距轨道边缘距离
	thumbSize := h - margin*2 // 滑块为正方形

	isCheck := t.IsCheck()

	// 计算滑块 X 坐标
	var thumbX int32
	if isCheck {
		thumbX = rc.Right - thumbSize - margin
	} else {
		thumbX = margin
	}

	// 选择轨道颜色
	trackColor := t.trackColorOff
	if isCheck {
		trackColor = t.trackColorOn
	}
	if !xc.XEle_IsEnable(hEle) {
		trackColor = t.disabledColor
	}

	// 1. 绘制轨道背景(全圆角矩形, 圆角半径 = 高度/2)
	draw.SetBrushColor(trackColor)
	draw.FillRoundRect(&rc, h/2, h/2)

	// 2. 绘制滑块(圆形)
	draw.SetBrushColor(t.thumbColor)
	var thumbRC xc.RECT
	thumbRC.Left = thumbX
	thumbRC.Top = margin
	thumbRC.Right = thumbX + thumbSize
	thumbRC.Bottom = margin + thumbSize
	draw.FillEllipse(&thumbRC)

	// 3. 滑块底部绘制微阴影
	shadowColor := xc.RGBA(0, 0, 0, 20)
	draw.SetBrushColor(shadowColor)
	var shadowRC xc.RECT
	shadowRC.Left = thumbX + 1
	shadowRC.Top = margin + 2
	shadowRC.Right = thumbX + thumbSize - 1
	shadowRC.Bottom = margin + thumbSize
	draw.FillEllipse(&shadowRC)

	return 0
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 400, 250, "美化开关按钮示例", 0, xcc.Window_Style_Default)

	// 1. 默认开关按钮
	tgl1 := NewBeautifyToggleButton(30, 40, w.Handle)
	tgl1.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		w.MessageBox("提示", "开关状态: "+common.Choose(bCheck, "开", "关"), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	// 2. 默认选中
	tgl2 := NewBeautifyToggleButton(30, 80, w.Handle)
	tgl2.SetCheck(true)

	// 3. 禁用状态
	tgl3 := NewBeautifyToggleButton(30, 120, w.Handle)
	tgl3.Enable(false)

	// 4. 自定义颜色(开=绿色, 关=红色)
	tgl4 := NewBeautifyToggleButton(30, 160, w.Handle)
	tgl4.SetTrackColorOn(xc.RGBA(16, 185, 129, 255)) // 开: 绿色
	tgl4.SetTrackColorOff(xc.RGBA(255, 77, 79, 255)) // 关: 红色

	w.Show(true)
	a.Run()
	a.Exit()
}
