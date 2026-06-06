// 美化按钮
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

// BeautifyButton 美化按钮
type BeautifyButton struct {
	*widget.Button
	themeColor          uint32 // 主题色
	bgColor             uint32 // 正常背景色
	hoverBgColor        uint32 // 悬停背景色
	pressedBgColor      uint32 // 按下背景色
	disabledBgColor     uint32 // 禁用背景色
	borderColor         uint32 // 正常边框色
	hoverBorderColor    uint32 // 悬停边框色
	pressedBorderColor  uint32 // 按下边框色
	disabledBorderColor uint32 // 禁用边框色
	textColor           uint32 // 文字颜色
	disabledTextColor   uint32 // 禁用状态文字颜色
	round               int32  // 圆角大小
	style               int    // 样式: 0=标准, 1=强调(填充)
}

// NewBeautifyButton 创建美化按钮
func NewBeautifyButton(x, y, cx, cy int32, text string, hParent int) *BeautifyButton {
	btn := widget.NewButton(x, y, cx, cy, text, hParent)

	w := &BeautifyButton{
		Button:              btn,
		themeColor:          xc.RGBA(0, 120, 212, 255),   // 默认主题色(蓝色)
		bgColor:             xc.RGBA(240, 243, 249, 255), // 正常: 浅灰 #F0F3F9
		hoverBgColor:        xc.RGBA(229, 233, 242, 255), // 悬停: 稍深 #E5E9F2
		pressedBgColor:      xc.RGBA(218, 223, 234, 255), // 按下: 更深 #DADEEA
		disabledBgColor:     xc.RGBA(244, 244, 244, 255), // 禁用: 灰色 #F4F4F4
		borderColor:         xc.RGBA(180, 188, 200, 255), // 正常边框
		hoverBorderColor:    xc.RGBA(150, 158, 170, 255), // 悬停边框
		pressedBorderColor:  xc.RGBA(120, 128, 140, 255), // 按下边框
		disabledBorderColor: xc.RGBA(210, 210, 210, 255), // 禁用边框
		textColor:           xc.RGBA(0, 0, 0, 255),       // 黑色文字
		disabledTextColor:   xc.RGBA(160, 160, 160, 255), // 禁用: 灰色文字
		round:               4,                           // 圆角
		style:               0,                           // 默认标准样式
	}

	w.initStyle()
	return w
}

// initStyle 初始化样式
func (b *BeautifyButton) initStyle() {
	// 启用背景透明（使用自绘）
	b.EnableBkTransparent(true)

	// 设置文字颜色和对齐
	b.SetTextColor(b.textColor)
	b.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)

	// 注册绘制事件
	b.AddEvent_Paint(b.onPaint)
}

// onPaint 绘制按钮
func (b *BeautifyButton) onPaint(hEle int, hDraw int, pbHandled *bool) int {
	draw := drawx.NewByHandle(hDraw)
	draw.EnableSmoothingMode(true)

	// 获取按钮状态
	nState := xc.XBtn_GetStateEx(hEle)

	// 根据状态选择背景色
	var bgColor uint32
	switch nState {
	case xcc.Button_State_Stay:
		// 悬停状态
		bgColor = b.hoverBgColor
	case xcc.Button_State_Down:
		// 按下状态
		bgColor = b.pressedBgColor
	case xcc.Button_State_Disable:
		// 禁用状态
		bgColor = b.disabledBgColor
	default:
		// 正常状态
		bgColor = b.bgColor
	}

	// 绘制圆角矩形背景
	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)
	draw.SetBrushColor(bgColor)
	draw.FillRoundRect(&rc, b.round, b.round)

	// 标准样式绘制边框，让按钮与背景区分开
	if b.style == 0 {
		var borderColor uint32
		switch nState {
		case xcc.Button_State_Stay:
			borderColor = b.hoverBorderColor
		case xcc.Button_State_Down:
			borderColor = b.pressedBorderColor
		case xcc.Button_State_Disable:
			borderColor = b.disabledBorderColor
		default:
			borderColor = b.borderColor
		}
		draw.SetLineWidth(1)
		draw.SetBrushColor(borderColor)
		draw.DrawRoundRect(&rc, b.round, b.round)
	}

	return 0
}

// SetThemeColor 设置主题色
func (w *BeautifyButton) SetThemeColor(color uint32) *BeautifyButton {
	w.themeColor = color
	r, g, b, a := xc.ParseRGBA(color)

	if w.style == 1 { // 强调样式: 填充背景色
		w.bgColor = color
		w.hoverBgColor = xc.RGBA(
			uint8(minInt(int(r)+30, 255)),
			uint8(minInt(int(g)+30, 255)),
			uint8(minInt(int(b)+30, 255)),
			a)
		w.pressedBgColor = xc.RGBA(
			uint8(minInt(int(r)-30, 0)),
			uint8(minInt(int(g)-30, 0)),
			uint8(minInt(int(b)-30, 0)),
			a)
		w.disabledBgColor = xc.RGBA(
			uint8((int(r)+240)/2),
			uint8((int(g)+240)/2),
			uint8((int(b)+240)/2),
			a)
		w.disabledTextColor = xc.RGBA(160, 160, 160, 255)
		w.SetTextColor(xc.RGBA(255, 255, 255, 255)) // 白色文字
	} else { // 标准样式: 边框体现主题色
		w.borderColor = color
		w.hoverBorderColor = xc.RGBA(
			uint8(minInt(int(r)+30, 255)),
			uint8(minInt(int(g)+30, 255)),
			uint8(minInt(int(b)+30, 255)),
			a)
		w.pressedBorderColor = xc.RGBA(
			uint8(maxInt(int(r)-30, 0)),
			uint8(maxInt(int(g)-30, 0)),
			uint8(maxInt(int(b)-30, 0)),
			a)
	}
	w.Redraw(false)
	return w
}

// SetStyle 设置按钮样式
// style: 0=标准按钮(浅灰背景), 1=强调按钮(主题色背景)
func (w *BeautifyButton) SetStyle(style int) *BeautifyButton {
	w.style = style
	if style == 1 {
		// 强调按钮: 使用主题色
		w.bgColor = w.themeColor
		r, g, b, a := xc.ParseRGBA(w.themeColor)
		w.hoverBgColor = xc.RGBA(
			uint8(minInt(int(r)+30, 255)),
			uint8(minInt(int(g)+30, 255)),
			uint8(minInt(int(b)+30, 255)),
			a)
		w.pressedBgColor = xc.RGBA(
			uint8(maxInt(int(r)-30, 0)),
			uint8(maxInt(int(g)-30, 0)),
			uint8(maxInt(int(b)-30, 0)),
			a)
		w.disabledBgColor = xc.RGBA(
			uint8((int(r)+240)/2),
			uint8((int(g)+240)/2),
			uint8((int(b)+240)/2),
			a)
		w.disabledTextColor = xc.RGBA(160, 160, 160, 255)
		w.textColor = xc.RGBA(255, 255, 255, 255) // 白色文字
		w.SetTextColor(w.textColor)
	}
	w.Redraw(false)
	return w
}

// GetThemeColor 获取主题色
func (b *BeautifyButton) GetThemeColor() uint32 { return b.themeColor }

// GetBgColor 获取正常背景色
func (b *BeautifyButton) GetBgColor() uint32 { return b.bgColor }

// SetBgColor 设置正常背景色
func (b *BeautifyButton) SetBgColor(color uint32) *BeautifyButton {
	b.bgColor = color
	b.Redraw(false)
	return b
}

// GetHoverBgColor 获取悬停背景色
func (b *BeautifyButton) GetHoverBgColor() uint32 { return b.hoverBgColor }

// SetHoverBgColor 设置悬停背景色
func (b *BeautifyButton) SetHoverBgColor(color uint32) *BeautifyButton {
	b.hoverBgColor = color
	b.Redraw(false)
	return b
}

// GetPressedBgColor 获取按下背景色
func (b *BeautifyButton) GetPressedBgColor() uint32 { return b.pressedBgColor }

// SetPressedBgColor 设置按下背景色
func (b *BeautifyButton) SetPressedBgColor(color uint32) *BeautifyButton {
	b.pressedBgColor = color
	b.Redraw(false)
	return b
}

// GetDisabledBgColor 获取禁用背景色
func (b *BeautifyButton) GetDisabledBgColor() uint32 { return b.disabledBgColor }

// SetDisabledBgColor 设置禁用背景色
func (b *BeautifyButton) SetDisabledBgColor(color uint32) *BeautifyButton {
	b.disabledBgColor = color
	b.Redraw(false)
	return b
}

// GetBorderColor 获取正常边框色
func (b *BeautifyButton) GetBorderColor() uint32 { return b.borderColor }

// SetBorderColor 设置正常边框色
func (b *BeautifyButton) SetBorderColor(color uint32) *BeautifyButton {
	b.borderColor = color
	b.Redraw(false)
	return b
}

// GetHoverBorderColor 获取悬停边框色
func (b *BeautifyButton) GetHoverBorderColor() uint32 { return b.hoverBorderColor }

// SetHoverBorderColor 设置悬停边框色
func (b *BeautifyButton) SetHoverBorderColor(color uint32) *BeautifyButton {
	b.hoverBorderColor = color
	b.Redraw(false)
	return b
}

// GetPressedBorderColor 获取按下边框色
func (b *BeautifyButton) GetPressedBorderColor() uint32 { return b.pressedBorderColor }

// SetPressedBorderColor 设置按下边框色
func (b *BeautifyButton) SetPressedBorderColor(color uint32) *BeautifyButton {
	b.pressedBorderColor = color
	b.Redraw(false)
	return b
}

// GetDisabledBorderColor 获取禁用边框色
func (b *BeautifyButton) GetDisabledBorderColor() uint32 { return b.disabledBorderColor }

// SetDisabledBorderColor 设置禁用边框色
func (b *BeautifyButton) SetDisabledBorderColor(color uint32) *BeautifyButton {
	b.disabledBorderColor = color
	b.Redraw(false)
	return b
}

// GetTextColor 获取文字颜色
func (b *BeautifyButton) GetTextColor() uint32 { return b.textColor }

// SetTextColor 设置文字颜色
func (b *BeautifyButton) SetTextColor(color uint32) *BeautifyButton {
	b.textColor = color
	b.Button.SetTextColor(color)
	b.Redraw(false)
	return b
}

// GetDisabledTextColor 获取禁用状态文字颜色
func (b *BeautifyButton) GetDisabledTextColor() uint32 { return b.disabledTextColor }

// SetDisabledTextColor 设置禁用状态文字颜色
func (b *BeautifyButton) SetDisabledTextColor(color uint32) *BeautifyButton {
	b.disabledTextColor = color
	return b
}

// GetRound 获取圆角大小
func (b *BeautifyButton) GetRound() int32 { return b.round }

// SetRound 设置圆角大小
func (b *BeautifyButton) SetRound(round int32) *BeautifyButton {
	b.round = round
	b.Redraw(false)
	return b
}

// GetStyle 获取按钮样式 (0=标准, 1=强调)
func (b *BeautifyButton) GetStyle() int { return b.style }

// minInt 辅助函数
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// maxInt 辅助函数
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 500, 480, "美化按钮示例", 0, xcc.Window_Style_Default)

	// 主题色 (蓝色)
	themeColor := xc.RGBA(0, 120, 212, 255) // #0078D4

	// 1. 标准按钮（浅灰背景，灰色边框）
	btn1 := NewBeautifyButton(50, 50, 120, 36, "标准按钮", w.Handle)
	btn1.SetThemeColor(themeColor)

	// 2. 强调按钮（主题色背景 + 白色文字）
	btn2 := NewBeautifyButton(200, 50, 120, 36, "强调按钮", w.Handle)
	btn2.SetThemeColor(themeColor).SetStyle(1)

	// 3. 绿色主题标准按钮
	btn3 := NewBeautifyButton(50, 120, 120, 36, "绿色", w.Handle)
	btn3.SetThemeColor(xc.RGBA(0, 180, 80, 255))

	// 4. 绿色强调按钮
	btn4 := NewBeautifyButton(200, 120, 120, 36, "绿色强调", w.Handle)
	btn4.SetThemeColor(xc.RGBA(0, 180, 80, 255)).SetStyle(1)

	// 5. 红色警告按钮
	btn5 := NewBeautifyButton(50, 190, 120, 36, "删除", w.Handle)
	btn5.SetThemeColor(xc.RGBA(220, 50, 50, 255)).SetStyle(1)

	// 6. 大圆角标准按钮
	btn6 := NewBeautifyButton(200, 190, 120, 36, "大圆角", w.Handle)
	btn6.SetThemeColor(themeColor).SetRound(18)

	// 7. 禁用状态标准按钮
	btn7 := NewBeautifyButton(50, 260, 120, 36, "禁用", w.Handle)
	btn7.SetThemeColor(themeColor)
	btn7.Enable(false)

	// 8. 禁用状态强调按钮
	btn8 := NewBeautifyButton(200, 260, 120, 36, "禁用强调", w.Handle)
	btn8.SetThemeColor(themeColor).SetStyle(1)
	btn8.Enable(false)

	// 9. 自定义背景色 + 边框色(标准样式)
	btn9 := NewBeautifyButton(50, 330, 250, 36, "自定义: 深蓝底+红边框", w.Handle)
	btn9.SetBgColor(xc.RGBA(220, 230, 250, 255)).
		SetHoverBgColor(xc.RGBA(200, 210, 240, 255)).
		SetPressedBgColor(xc.RGBA(180, 190, 230, 255)).
		SetBorderColor(xc.RGBA(220, 50, 50, 255)).
		SetHoverBorderColor(xc.RGBA(200, 30, 30, 255))

	w.Show(true)
	a.Run()
	a.Exit()
}
