// 美化滑块条
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BeautifySliderBar 美化滑块条.
//
// 继承 widget.SliderBar, 通过自绘实现:
//   - 轨道: 圆角矩形, 左侧填充蓝、右侧灰色
//   - 滑块: 圆形, 带阴影效果
type BeautifySliderBar struct {
	*widget.SliderBar
	trackHeight        int32  // 轨道高度
	trackColorOff      uint32 // 未填充轨道颜色(右侧)
	trackColorOn       uint32 // 已填充轨道颜色(左侧)
	thumbSize          int32  // 滑块直径
	thumbColor         uint32 // 滑块默认颜色
	thumbBorderColor   uint32 // 滑块边框颜色
	thumbHoverColor    uint32 // 悬停滑块颜色
	thumbDownColor     uint32 // 按下滑块颜色
	thumbShadowColor   uint32 // 滑块阴影颜色
	disabledTrackOff   uint32 // 禁用: 未填充轨道
	disabledTrackOn    uint32 // 禁用: 已填充轨道
	disabledThumbColor uint32 // 禁用: 滑块颜色
}

// NewBeautifySliderBar 创建美化滑块条. 默认水平方向.
//
// x, y: 坐标; cx, cy: 宽度/高度; hParent: 父句柄.
func NewBeautifySliderBar(x, y, cx, cy int32, hParent int) *BeautifySliderBar {
	sb := widget.NewSliderBar(x, y, cx, cy, hParent)
	if sb == nil {
		return nil
	}
	t := &BeautifySliderBar{
		SliderBar:          sb,
		trackHeight:        6,
		trackColorOff:      xc.RGBA(227, 227, 227, 255), // 灰
		trackColorOn:       xc.RGBA(24, 144, 255, 255),  // 蓝
		thumbSize:          20,
		thumbColor:         xc.RGBA(255, 255, 255, 255), // 白色
		thumbBorderColor:   xc.RGBA(24, 144, 255, 255),  // 蓝色边框
		thumbHoverColor:    xc.RGBA(230, 244, 255, 255), // 悬停浅蓝
		thumbDownColor:     xc.RGBA(24, 144, 255, 255),  // 按下蓝色实心
		thumbShadowColor:   xc.RGBA(0, 0, 0, 25),        // 半透明黑阴影
		disabledTrackOff:   xc.RGBA(240, 240, 240, 255),
		disabledTrackOn:    xc.RGBA(217, 217, 217, 255),
		disabledThumbColor: xc.RGBA(245, 245, 245, 255),
	}
	t.EnableBkTransparent(true)

	// 设置滑块的宽高
	t.SetButtonWidth(t.thumbSize)
	t.SetButtonHeight(t.thumbSize)
	t.SetRange(100)

	// 注册轨道自绘事件
	t.AddEvent_Paint(t.onTrackPaint)

	// 获取滑块按钮, 注册滑块自绘事件
	thumbBtn := t.GetButtonObj()
	if thumbBtn != nil {
		thumbBtn.AddEvent_Paint(t.onThumbPaint)
	}
	return t
}

// SetRange 设置滑动范围.
func (t *BeautifySliderBar) SetRange(range_ int32) *BeautifySliderBar {
	t.SliderBar.SetRange(range_)
	return t
}

// SetTrackHeight 设置轨道高度.
func (t *BeautifySliderBar) SetTrackHeight(height int32) {
	t.trackHeight = height
}

// SetTrackColorOff 设置未填充轨道颜色.
func (t *BeautifySliderBar) SetTrackColorOff(color uint32) {
	t.trackColorOff = color
}

// SetTrackColorOn 设置已填充轨道颜色.
func (t *BeautifySliderBar) SetTrackColorOn(color uint32) {
	t.trackColorOn = color
}

// SetThumbSize 设置滑块大小.
func (t *BeautifySliderBar) SetThumbSize(size int32) {
	t.thumbSize = size
	t.SetButtonWidth(size)
	t.SetButtonHeight(size)
}

// SetThumbColor 设置滑块默认颜色.
func (t *BeautifySliderBar) SetThumbColor(color uint32) {
	t.thumbColor = color
}

// SetThumbBorderColor 设置滑块边框颜色.
func (t *BeautifySliderBar) SetThumbBorderColor(color uint32) {
	t.thumbBorderColor = color
}

// SetThumbHoverColor 设置悬停滑块颜色.
func (t *BeautifySliderBar) SetThumbHoverColor(color uint32) {
	t.thumbHoverColor = color
}

// SetThumbDownColor 设置按下滑块颜色.
func (t *BeautifySliderBar) SetThumbDownColor(color uint32) {
	t.thumbDownColor = color
}

// onTrackPaint 自绘轨道: 绘制圆角矩形轨道背景(已填充+未填充).
func (t *BeautifySliderBar) onTrackPaint(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	draw := drawx.NewByHandle(hDraw)
	if draw == nil {
		return 0
	}
	draw.EnableSmoothingMode(true)

	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)
	rcH := rc.Bottom - rc.Top

	// 轨道垂直居中
	trackH := t.trackHeight
	trackTop := (rcH - trackH) / 2
	trackBottom := trackTop + trackH
	thumbSize := t.thumbSize
	// 轨道水平留出滑块半径的边距
	padding := thumbSize / 2
	trackLeft := rc.Left + padding
	trackRight := rc.Right - padding
	roundRadius := trackH / 2

	// 计算当前进度比例
	range_ := t.GetRange()
	pos := t.GetPos()
	ratio := float32(0)
	if range_ > 0 {
		ratio = float32(pos) / float32(range_)
	}
	splitX := trackLeft + int32(float32(trackRight-trackLeft)*ratio+0.5)

	isEnabled := xc.XEle_IsEnable(hEle)

	// 填充颜色
	colorOn := t.trackColorOn
	colorOff := t.trackColorOff
	if !isEnabled {
		colorOn = t.disabledTrackOn
		colorOff = t.disabledTrackOff
	}

	// 1. 绘制已填充部分(左侧)
	if splitX > trackLeft {
		var fillRC xc.RECT
		fillRC.Left = trackLeft
		fillRC.Top = trackTop
		fillRC.Right = splitX
		fillRC.Bottom = trackBottom
		draw.SetBrushColor(colorOn)
		// 如果填充区域小于轨道高度, 用矩形直接填充; 否则用圆角矩形
		if splitX-trackLeft < trackH {
			draw.FillRect(&fillRC)
		} else {
			draw.FillRoundRect(&fillRC, roundRadius, roundRadius)
		}
	}

	// 2. 绘制未填充部分(右侧)
	if splitX < trackRight {
		var unfillRC xc.RECT
		unfillRC.Left = splitX
		unfillRC.Top = trackTop
		unfillRC.Right = trackRight
		unfillRC.Bottom = trackBottom
		draw.SetBrushColor(colorOff)
		if trackRight-splitX < trackH {
			draw.FillRect(&unfillRC)
		} else {
			draw.FillRoundRect(&unfillRC, roundRadius, roundRadius)
		}
	}

	return 0
}

// onThumbPaint 自绘滑块按钮: 绘制圆形滑块 + 阴影.
func (t *BeautifySliderBar) onThumbPaint(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	draw := drawx.NewByHandle(hDraw)
	if draw == nil {
		return 0
	}
	draw.EnableSmoothingMode(true)

	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)
	rcW := rc.Right - rc.Left
	rcH := rc.Bottom - rc.Top

	isEnabled := xc.XEle_IsEnable(hEle)

	// 获取滑块按钮状态
	state := xc.XBtn_GetState(hEle) // Common_State3_

	// 选择颜色
	var thumbFill, thumbBorder uint32
	if !isEnabled {
		thumbFill = t.disabledThumbColor
		thumbBorder = t.disabledTrackOn
	} else if state == xcc.Common_State3_Down {
		thumbFill = t.thumbDownColor
		thumbBorder = t.thumbDownColor
	} else if state == xcc.Common_State3_Stay {
		thumbFill = t.thumbHoverColor
		thumbBorder = t.thumbBorderColor
	} else {
		thumbFill = t.thumbColor
		thumbBorder = t.thumbBorderColor
	}

	// 居中绘制圆形, 留2px内边距防止边缘被裁剪
	circleSize := rcW
	if rcH < rcW {
		circleSize = rcH
	}
	padding := int32(2)
	circleSize -= padding * 2
	margin := (rcW - circleSize) / 2
	if margin < 0 {
		margin = 0
	}

	var circleRC xc.RECT
	circleRC.Left = rc.Left + margin
	circleRC.Top = rc.Top + margin
	circleRC.Right = circleRC.Left + circleSize
	circleRC.Bottom = circleRC.Top + circleSize

	// 1. 阴影
	if isEnabled {
		draw.SetBrushColor(t.thumbShadowColor)
		var shadowRC xc.RECT
		shadowRC.Left = circleRC.Left
		shadowRC.Top = circleRC.Top + 1
		shadowRC.Right = circleRC.Right
		shadowRC.Bottom = circleRC.Bottom + 1
		draw.FillEllipse(&shadowRC)
	}

	// 2. 填充圆
	draw.SetBrushColor(thumbFill)
	draw.FillEllipse(&circleRC)

	// 3. 边框(仅在非按下状态绘制)
	if state != xcc.Common_State3_Down || !isEnabled {
		borderThickness := int32(2)
		draw.SetLineWidth(borderThickness)
		draw.SetBrushColor(thumbBorder)
		draw.DrawEllipse(&circleRC)
	}

	return 0
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	yOff := int32(15)
	w := window.New(0, 0, 500, 330+yOff*3, "美化滑块条示例", 0, xcc.Window_Style_Default)
	w.SetBorderSize(2, 30, 2, 2)

	// ---- 1. 默认蓝色主题 ----
	widget.NewShapeText(20, 20+yOff, 200, 20, "默认蓝色主题", w.Handle)

	sb1 := NewBeautifySliderBar(20, 45+yOff, 400, 40, w.Handle)
	sb1.SetRange(100)
	sb1.SetPos(30)
	// 显示当前值
	valueText1 := widget.NewShapeText(430, 45+yOff, 50, 40, "30", w.Handle)
	valueText1.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter)
	sb1.AddEvent_SliderBar_Change(func(hEle int, pos int32, pbHandled *bool) int {
		valueText1.SetText(fmt.Sprintf("%d", pos)).Redraw()
		return 0
	})

	// ---- 2. 绿色主题 ----
	widget.NewShapeText(20, 100+yOff, 200, 20, "自定义绿色主题", w.Handle)

	sb2 := NewBeautifySliderBar(20, 125+yOff, 400, 40, w.Handle)
	sb2.SetRange(100)
	sb2.SetPos(60)
	sb2.SetTrackColorOn(xc.RGBA(82, 196, 26, 255))
	sb2.SetThumbBorderColor(xc.RGBA(82, 196, 26, 255))
	sb2.SetThumbDownColor(xc.RGBA(82, 196, 26, 255))
	sb2.SetThumbHoverColor(xc.RGBA(240, 255, 230, 255))
	valueText2 := widget.NewShapeText(430, 125+yOff, 50, 40, "60", w.Handle)
	valueText2.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter)
	sb2.AddEvent_SliderBar_Change(func(hEle int, pos int32, pbHandled *bool) int {
		valueText2.SetText(fmt.Sprintf("%d", pos)).Redraw()
		return 0
	})

	// ---- 3. 禁用状态 ----
	widget.NewShapeText(20, 180+yOff, 200, 20, "禁用状态", w.Handle)

	sb3 := NewBeautifySliderBar(20, 205+yOff, 400, 40, w.Handle)
	sb3.SetRange(100)
	sb3.SetPos(70)
	sb3.Enable(false)

	// ---- 4. 大滑块示例 ----
	widget.NewShapeText(20, 260+yOff, 200, 20, "大滑块 + 无边框按下", w.Handle)

	sb4 := NewBeautifySliderBar(20, 285+yOff, 400, 50, w.Handle)
	sb4.SetThumbSize(30)
	sb4.SetRange(100)
	sb4.SetPos(50)
	sb4.SetTrackColorOn(xc.RGBA(114, 46, 209, 255))
	sb4.SetThumbBorderColor(xc.RGBA(114, 46, 209, 255))
	sb4.SetThumbDownColor(xc.RGBA(114, 46, 209, 255))
	sb4.SetThumbHoverColor(xc.RGBA(245, 235, 255, 255))
	valueText4 := widget.NewShapeText(430, 285+yOff, 50, 50, "50", w.Handle)
	valueText4.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter)
	sb4.AddEvent_SliderBar_Change(func(hEle int, pos int32, pbHandled *bool) int {
		valueText4.SetText(fmt.Sprintf("%d", pos)).Redraw()
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
