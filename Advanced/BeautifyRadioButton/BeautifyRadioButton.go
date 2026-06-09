// 美化单选按钮.
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BeautifyRadioButton 美化单选按钮.
//
// 继承 widget.Button, 通过自绘实现单选按钮外观.
// 默认使用蓝色主题, 选中状态: 蓝色外环 + 蓝色实心内圆.
// 未选中状态: 灰色外环, 内圆空心.
type BeautifyRadioButton struct {
	*widget.Button
	circleSize          int32  // 圆形指示器大小
	textOffset          int32  // 文本与圆形间距
	borderColorOn       uint32 // 选中: 外环颜色
	borderColorOff      uint32 // 未选中: 外环颜色
	dotColorOn          uint32 // 选中: 内圆颜色
	dotColorOff         uint32 // 未选中: 内圆颜色(透明)
	hoverBorderColorOn  uint32 // 悬停 + 选中: 外环颜色(高亮)
	hoverBorderColorOff uint32 // 悬停 + 未选中: 外环颜色
	disabledBorderColor uint32 // 禁用: 外环颜色
	disabledDotColor    uint32 // 禁用: 内圆颜色
}

// NewBeautifyRadioButton 创建美化单选按钮.
//
// x, y: 坐标;
//
// w, h: 宽度, 高度;
//
// text: 文本标签;
//
// hParent: 父句柄.
func NewBeautifyRadioButton(x, y, w, h int32, text string, hParent int) *BeautifyRadioButton {
	btn := widget.NewButton(x, y, w, h, text, hParent)
	if btn == nil {
		return nil
	}
	t := &BeautifyRadioButton{
		Button:              btn,
		circleSize:          18,
		textOffset:          8,
		borderColorOn:       xc.RGBA(24, 144, 255, 255),  // 选中蓝环
		borderColorOff:      xc.RGBA(191, 191, 191, 255), // 未选中灰环
		dotColorOn:          xc.RGBA(24, 144, 255, 255),  // 选中实心蓝点
		dotColorOff:         xc.RGBA(0, 0, 0, 0),         // 未选中透明
		hoverBorderColorOn:  xc.RGBA(64, 169, 255, 255),  // 悬停+选中: 亮蓝
		hoverBorderColorOff: xc.RGBA(168, 168, 168, 255), // 悬停+未选中: 加深灰
		disabledBorderColor: xc.RGBA(217, 217, 217, 255), // 禁用灰环
		disabledDotColor:    xc.RGBA(217, 217, 217, 255), // 禁用灰点
	}
	// 设置按钮类型为单选按钮
	t.SetTypeEx(xcc.Button_Type_Radio)
	// 设置文本左对齐, 文本垂直居中
	t.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter | xcc.TextTrimming_EllipsisCharacter)
	// 注册自绘事件
	t.AddEvent_Paint(t.onPaint)
	return t
}

// SetCircleSize 设置圆形指示器大小.
func (t *BeautifyRadioButton) SetCircleSize(size int32) {
	t.circleSize = size
}

// GetCircleSize 获取圆形指示器大小.
func (t *BeautifyRadioButton) GetCircleSize() int32 {
	return t.circleSize
}

// SetTextOffset 设置文本与圆形间距.
func (t *BeautifyRadioButton) SetTextOffset(offset int32) {
	t.textOffset = offset
}

// GetTextOffset 获取文本与圆形间距.
func (t *BeautifyRadioButton) GetTextOffset() int32 {
	return t.textOffset
}

// SetBorderColorOn 设置选中状态外环颜色.
func (t *BeautifyRadioButton) SetBorderColorOn(color uint32) {
	t.borderColorOn = color
}

// GetBorderColorOn 获取选中状态外环颜色.
func (t *BeautifyRadioButton) GetBorderColorOn() uint32 {
	return t.borderColorOn
}

// SetBorderColorOff 设置未选中状态外环颜色.
func (t *BeautifyRadioButton) SetBorderColorOff(color uint32) {
	t.borderColorOff = color
}

// GetBorderColorOff 获取未选中状态外环颜色.
func (t *BeautifyRadioButton) GetBorderColorOff() uint32 {
	return t.borderColorOff
}

// SetDotColorOn 设置选中状态内圆颜色.
func (t *BeautifyRadioButton) SetDotColorOn(color uint32) {
	t.dotColorOn = color
}

// GetDotColorOn 获取选中状态内圆颜色.
func (t *BeautifyRadioButton) GetDotColorOn() uint32 {
	return t.dotColorOn
}

// SetDotColorOff 设置未选中状态内圆颜色(通常为透明).
func (t *BeautifyRadioButton) SetDotColorOff(color uint32) {
	t.dotColorOff = color
}

// GetDotColorOff 获取未选中状态内圆颜色.
func (t *BeautifyRadioButton) GetDotColorOff() uint32 {
	return t.dotColorOff
}

// SetHoverBorderColorOn 设置悬停+选中状态外环颜色.
func (t *BeautifyRadioButton) SetHoverBorderColorOn(color uint32) {
	t.hoverBorderColorOn = color
}

// GetHoverBorderColorOn 获取悬停+选中状态外环颜色.
func (t *BeautifyRadioButton) GetHoverBorderColorOn() uint32 {
	return t.hoverBorderColorOn
}

// SetHoverBorderColorOff 设置悬停+未选中状态外环颜色.
func (t *BeautifyRadioButton) SetHoverBorderColorOff(color uint32) {
	t.hoverBorderColorOff = color
}

// GetHoverBorderColorOff 获取悬停+未选中状态外环颜色.
func (t *BeautifyRadioButton) GetHoverBorderColorOff() uint32 {
	return t.hoverBorderColorOff
}

// SetDisabledBorderColor 设置禁用状态外环颜色.
func (t *BeautifyRadioButton) SetDisabledBorderColor(color uint32) {
	t.disabledBorderColor = color
}

// GetDisabledBorderColor 获取禁用状态外环颜色.
func (t *BeautifyRadioButton) GetDisabledBorderColor() uint32 {
	return t.disabledBorderColor
}

// SetDisabledDotColor 设置禁用状态内圆颜色.
func (t *BeautifyRadioButton) SetDisabledDotColor(color uint32) {
	t.disabledDotColor = color
}

// GetDisabledDotColor 获取禁用状态内圆颜色.
func (t *BeautifyRadioButton) GetDisabledDotColor() uint32 {
	return t.disabledDotColor
}

// onPaint 自绘事件: 绘制圆形指示器 + 文本标签.
func (t *BeautifyRadioButton) onPaint(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true // 拦截元素原本的绘制
	draw := drawx.NewByHandle(hDraw)
	if draw == nil {
		return 0
	}
	draw.EnableSmoothingMode(true)

	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)

	h := rc.Bottom - rc.Top
	isCheck := t.IsCheck()
	isHover := t.GetState() != xcc.Common_State3_Leave
	isEnabled := xc.XEle_IsEnable(hEle)

	// 圆形指示器居中于左侧
	margin := int32(2)
	circleSize := t.circleSize
	circleTop := (h - circleSize) / 2
	circleBottom := circleTop + circleSize
	circleLeft := margin
	circleRight := circleLeft + circleSize

	var borderColor, dotColor uint32

	if !isEnabled {
		borderColor = t.disabledBorderColor
		if isCheck {
			dotColor = t.disabledDotColor
		} else {
			dotColor = t.dotColorOff
		}
	} else if isHover && isCheck {
		borderColor = t.hoverBorderColorOn
		dotColor = t.dotColorOn
	} else if isHover && !isCheck {
		borderColor = t.hoverBorderColorOff
		dotColor = t.dotColorOff
	} else if isCheck {
		borderColor = t.borderColorOn
		dotColor = t.dotColorOn
	} else {
		borderColor = t.borderColorOff
		dotColor = t.dotColorOff
	}

	// 1. 绘制外环(先填充白色背景, 再绘制彩色边框)
	var circleRC xc.RECT
	circleRC.Left = circleLeft
	circleRC.Top = circleTop
	circleRC.Right = circleRight
	circleRC.Bottom = circleBottom

	// 填充白色背景(确保在透明背景下边缘干净)
	draw.SetBrushColor(xc.RGBA(255, 255, 255, 255))
	draw.FillEllipse(&circleRC)

	// 绘制外环(圆形边框)
	borderThickness := int32(2)
	draw.SetLineWidth(borderThickness)
	draw.SetBrushColor(borderColor)
	draw.DrawEllipse(&circleRC)

	// 2. 绘制内圆(选中时的实心圆点)
	if isCheck {
		dotSize := circleSize/2 - 1 // 取偶数, 方便居中计算
		if !isEnabled {
			dotSize = circleSize / 3
		}
		// 以外环中心为基准
		cx := circleLeft + circleSize/2
		cy := circleTop + circleSize/2
		dotLeft := cx - dotSize/2
		dotTop := cy - dotSize/2
		var dotRC xc.RECT
		dotRC.Left = dotLeft
		dotRC.Top = dotTop
		dotRC.Right = dotLeft + dotSize
		dotRC.Bottom = dotTop + dotSize
		draw.SetBrushColor(dotColor)
		draw.FillEllipse(&dotRC)
	}

	// 3. 绘制文本标签
	text := t.GetText()
	if text != "" {
		textLeft := circleRight + t.textOffset
		var textRC xc.RECT
		textRC.Left = textLeft
		textRC.Top = rc.Top
		textRC.Right = rc.Right - margin
		textRC.Bottom = rc.Bottom

		// 文本颜色: 根据状态选择
		textColor := xc.RGBA(51, 51, 51, 255) // 正常: 深灰
		if !isEnabled {
			textColor = xc.RGBA(191, 191, 191, 255) // 禁用: 浅灰
		} else if isHover {
			textColor = xc.RGBA(24, 144, 255, 255) // 悬停: 蓝色
		}
		draw.SetBrushColor(textColor)
		draw.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter | xcc.TextTrimming_EllipsisCharacter)
		draw.DrawText(text, &textRC)
	}

	return 0
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 400, 290, "美化单选按钮示例", 0, xcc.Window_Style_Default)
	w.SetBorderSize(2, 30, 2, 2)

	// 第一组: 默认颜色主题
	widget.NewShapeText(30, 32, 200, 20, "默认蓝色主题", w.Handle)

	// 使用 SetGroupID 分组, 同一组内多选一
	rb1 := NewBeautifyRadioButton(30, 57, 150, 26, "选项 A", w.Handle)
	rb1.SetGroupID(1)
	rb1.SetCheck(true) // 默认选中

	rb2 := NewBeautifyRadioButton(30, 87, 150, 26, "选项 B", w.Handle)
	rb2.SetGroupID(1)

	rb3 := NewBeautifyRadioButton(30, 117, 150, 26, "选项 C", w.Handle)
	rb3.SetGroupID(1)

	// 选中事件演示
	rb1.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		if bCheck {
			w.MessageBox("提示", "选中: 选项 A", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		}
		return 0
	})
	rb3.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		if bCheck {
			w.MessageBox("提示", "选中: 选项 C", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		}
		return 0
	})

	// 第二组: 绿色主题 + 禁用状态
	widget.NewShapeText(30, 152, 200, 20, "自定义绿色主题", w.Handle)

	rb4 := NewBeautifyRadioButton(30, 177, 180, 26, "选项 X (自定义颜色)", w.Handle)
	rb4.SetGroupID(2)
	rb4.SetBorderColorOn(xc.RGBA(82, 196, 26, 255))       // 选中: 绿色环
	rb4.SetDotColorOn(xc.RGBA(82, 196, 26, 255))          // 选中: 绿点
	rb4.SetHoverBorderColorOn(xc.RGBA(115, 209, 61, 255)) // 悬停+选中: 亮绿
	rb4.SetCheck(true)

	rb5 := NewBeautifyRadioButton(30, 207, 180, 26, "选项 Y (禁用)", w.Handle)
	rb5.SetGroupID(2)
	rb5.Enable(false)
	rb5.SetCheck(true) // 禁用且选中

	rb6 := NewBeautifyRadioButton(30, 237, 180, 26, "选项 Z (禁用未选中)", w.Handle)
	rb6.SetGroupID(2)
	rb6.Enable(false)

	w.Show(true)
	a.Run()
	a.Exit()
}
