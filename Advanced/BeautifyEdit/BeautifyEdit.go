// 美化编辑框(自绘方式)
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

// BeautifyEdit 美化编辑框
type BeautifyEdit struct {
	*widget.Edit
	themeColor          uint32 // 主题色
	bgColor             uint32 // 正常背景色
	hoverBgColor        uint32 // 悬停背景色
	focusBgColor        uint32 // 焦点背景色
	disabledBgColor     uint32 // 禁用背景色
	borderColor         uint32 // 正常边框色
	hoverBorderColor    uint32 // 悬停边框色
	focusBorderColor    uint32 // 焦点边框色
	disabledBorderColor uint32 // 禁用边框色
	textColor           uint32 // 文本颜色
	disabledTextColor   uint32 // 禁用文字颜色
	round               int32  // 圆角大小
}

// NewBeautifyEdit 创建美化编辑框
func NewBeautifyEdit(x, y, cx, cy int32, hParent int) *BeautifyEdit {
	edit := widget.NewEdit(x, y, cx, cy, hParent)

	w := &BeautifyEdit{
		Edit:                edit,
		themeColor:          xc.RGBA(0, 120, 212, 255),   // 默认主题色(蓝色)
		bgColor:             xc.RGBA(255, 255, 255, 255), // 白色背景
		hoverBgColor:        xc.RGBA(248, 250, 252, 255), // 悬停: 微灰
		focusBgColor:        xc.RGBA(255, 255, 255, 255), // 焦点: 白色
		disabledBgColor:     xc.RGBA(245, 245, 245, 255), // 禁用: 浅灰
		borderColor:         xc.RGBA(200, 200, 200, 255), // 正常边框
		hoverBorderColor:    xc.RGBA(140, 140, 140, 255), // 悬停边框
		focusBorderColor:    xc.RGBA(0, 120, 212, 255),   // 焦点边框(主题色)
		disabledBorderColor: xc.RGBA(220, 220, 220, 255), // 禁用边框
		textColor:           xc.RGBA(0, 0, 0, 255),       // 黑色文字
		disabledTextColor:   xc.RGBA(160, 160, 160, 255), // 禁用文字
		round:               4,                           // 圆角
	}

	w.initStyle()
	return w
}

// initStyle 初始化样式
func (b *BeautifyEdit) initStyle() {
	// 启用背景透明（使用自绘）
	b.EnableBkTransparent(true)
	// 禁用绘制默认边框和焦点边框
	b.EnableDrawBorder(false).EnableDrawFocus(false)

	// 设置文字颜色和提示文字颜色
	b.SetTextColor(b.textColor)
	b.SetDefaultTextColor(xc.RGBA(150, 150, 150, 255))

	// 设置边框内边距, 让文字不贴在圆角上
	b.SetBorderSize(10, 4, 10, 4)

	// 设置插入符颜色
	b.SetCaretColor(b.themeColor)

	// 注册绘制事件
	b.AddEvent_Paint(b.onPaint)
}

// onPaint 绘制编辑框
func (b *BeautifyEdit) onPaint(hEle int, hDraw int, pbHandled *bool) int {
	draw := drawx.NewByHandle(hDraw)
	draw.EnableSmoothingMode(true)

	// 获取元素状态
	state := xc.XEle_GetStateFlags(hEle)

	// 根据状态选择背景色和边框色
	var bgColor, borderColor uint32

	if state&xcc.Element_State_Flag_Disable != 0 {
		// 禁用状态
		bgColor = b.disabledBgColor
		borderColor = b.disabledBorderColor
	} else if state&xcc.Element_State_Flag_Focus != 0 {
		// 焦点状态
		bgColor = b.focusBgColor
		borderColor = b.focusBorderColor
	} else if state&xcc.Element_State_Flag_Stay != 0 {
		// 悬停状态
		bgColor = b.hoverBgColor
		borderColor = b.hoverBorderColor
	} else {
		// 正常状态
		bgColor = b.bgColor
		borderColor = b.borderColor
	}

	// 获取客户区
	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)

	// 绘制圆角矩形背景
	draw.SetBrushColor(bgColor)
	draw.FillRoundRect(&rc, b.round, b.round)

	// 绘制 1px 圆角矩形边框
	draw.SetLineWidth(1)
	draw.SetBrushColor(borderColor)
	draw.DrawRoundRect(&rc, b.round, b.round)

	return 0
}

// SetThemeColor 设置主题色
func (b *BeautifyEdit) SetThemeColor(color uint32) *BeautifyEdit {
	b.themeColor = color
	b.focusBorderColor = color
	b.Redraw(false)
	return b
}

// GetThemeColor 获取主题色
func (b *BeautifyEdit) GetThemeColor() uint32 { return b.themeColor }

// GetBgColor 获取正常背景色
func (b *BeautifyEdit) GetBgColor() uint32 { return b.bgColor }

// SetBgColor 设置正常背景色
func (b *BeautifyEdit) SetBgColor(color uint32) *BeautifyEdit {
	b.bgColor = color
	b.Redraw(false)
	return b
}

// GetHoverBgColor 获取悬停背景色
func (b *BeautifyEdit) GetHoverBgColor() uint32 { return b.hoverBgColor }

// SetHoverBgColor 设置悬停背景色
func (b *BeautifyEdit) SetHoverBgColor(color uint32) *BeautifyEdit {
	b.hoverBgColor = color
	b.Redraw(false)
	return b
}

// GetFocusBgColor 获取焦点背景色
func (b *BeautifyEdit) GetFocusBgColor() uint32 { return b.focusBgColor }

// SetFocusBgColor 设置焦点背景色
func (b *BeautifyEdit) SetFocusBgColor(color uint32) *BeautifyEdit {
	b.focusBgColor = color
	b.Redraw(false)
	return b
}

// GetDisabledBgColor 获取禁用背景色
func (b *BeautifyEdit) GetDisabledBgColor() uint32 { return b.disabledBgColor }

// SetDisabledBgColor 设置禁用背景色
func (b *BeautifyEdit) SetDisabledBgColor(color uint32) *BeautifyEdit {
	b.disabledBgColor = color
	b.Redraw(false)
	return b
}

// GetBorderColor 获取正常边框色
func (b *BeautifyEdit) GetBorderColor() uint32 { return b.borderColor }

// SetBorderColor 设置正常边框色
func (b *BeautifyEdit) SetBorderColor(color uint32) *BeautifyEdit {
	b.borderColor = color
	b.Redraw(false)
	return b
}

// GetHoverBorderColor 获取悬停边框色
func (b *BeautifyEdit) GetHoverBorderColor() uint32 { return b.hoverBorderColor }

// SetHoverBorderColor 设置悬停边框色
func (b *BeautifyEdit) SetHoverBorderColor(color uint32) *BeautifyEdit {
	b.hoverBorderColor = color
	b.Redraw(false)
	return b
}

// GetFocusBorderColor 获取焦点边框色
func (b *BeautifyEdit) GetFocusBorderColor() uint32 { return b.focusBorderColor }

// SetFocusBorderColor 设置焦点边框色
func (b *BeautifyEdit) SetFocusBorderColor(color uint32) *BeautifyEdit {
	b.focusBorderColor = color
	b.Redraw(false)
	return b
}

// GetDisabledBorderColor 获取禁用边框色
func (b *BeautifyEdit) GetDisabledBorderColor() uint32 { return b.disabledBorderColor }

// SetDisabledBorderColor 设置禁用边框色
func (b *BeautifyEdit) SetDisabledBorderColor(color uint32) *BeautifyEdit {
	b.disabledBorderColor = color
	b.Redraw(false)
	return b
}

// GetTextColor 获取文字颜色
func (b *BeautifyEdit) GetTextColor() uint32 { return b.textColor }

// SetTextColor 设置文字颜色
func (b *BeautifyEdit) SetTextColor(color uint32) *BeautifyEdit {
	b.textColor = color
	b.Edit.SetTextColor(color)
	return b
}

// GetDisabledTextColor 获取禁用文字颜色
func (b *BeautifyEdit) GetDisabledTextColor() uint32 { return b.disabledTextColor }

// SetDisabledTextColor 设置禁用文字颜色
func (b *BeautifyEdit) SetDisabledTextColor(color uint32) *BeautifyEdit {
	b.disabledTextColor = color
	return b
}

// GetRound 获取圆角大小
func (b *BeautifyEdit) GetRound() int32 { return b.round }

// SetRound 设置圆角大小
func (b *BeautifyEdit) SetRound(round int32) *BeautifyEdit {
	b.round = round
	b.Redraw(false)
	return b
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 500, 580, "美化编辑框(自绘)", 0, xcc.Window_Style_Default)

	// 主题色
	themeColor := xc.RGBA(0, 120, 212, 255)

	// 1. 普通编辑框
	edit1 := NewBeautifyEdit(30, 60, 300, 40, w.Handle)
	edit1.SetThemeColor(themeColor).SetDefaultText("请输入用户名")
	edit1.SetRound(0)

	// 2. 密码框
	edit2 := NewBeautifyEdit(30, 120, 300, 40, w.Handle)
	edit2.SetThemeColor(themeColor).SetDefaultText("请输入密码")
	edit2.EnablePassword(true)

	// 3. 只读编辑框
	edit3 := NewBeautifyEdit(30, 180, 300, 40, w.Handle)
	edit3.SetThemeColor(themeColor).SetText("这是只读文本")
	edit3.EnableReadOnly(true)

	// 4. 绿色主题
	edit4 := NewBeautifyEdit(30, 240, 300, 40, w.Handle)
	edit4.SetThemeColor(xc.RGBA(0, 180, 80, 255)).SetDefaultText("绿色主题")

	// 5. 大圆角 + 自定义颜色
	edit5 := NewBeautifyEdit(30, 300, 300, 40, w.Handle)
	edit5.SetDefaultText("橙色大圆角编辑框")
	edit5.SetRound(16).
		SetThemeColor(xc.RGBA(255, 120, 0, 255)).
		SetBgColor(xc.RGBA(255, 248, 240, 255)).
		SetHoverBgColor(xc.RGBA(255, 242, 230, 255)).
		SetFocusBgColor(xc.RGBA(255, 248, 240, 255)).
		SetBorderColor(xc.RGBA(220, 180, 140, 255)).
		SetHoverBorderColor(xc.RGBA(200, 150, 100, 255)).
		SetFocusBorderColor(xc.RGBA(255, 120, 0, 255)).
		SetTextColor(xc.RGBA(80, 40, 10, 255))

	// 6. 禁用状态
	edit6 := NewBeautifyEdit(30, 360, 300, 40, w.Handle)
	edit6.SetThemeColor(themeColor).SetText("禁用状态")
	edit6.Enable(false)

	// 7. 紫色主题
	edit7 := NewBeautifyEdit(30, 420, 300, 40, w.Handle)
	edit7.SetDefaultText("紫色主题编辑框")
	edit7.SetThemeColor(xc.RGBA(160, 50, 200, 255)).
		SetFocusBorderColor(xc.RGBA(160, 50, 200, 255))

	w.Show(true)
	a.Run()
	a.Exit()
}
