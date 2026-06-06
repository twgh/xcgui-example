// Draw 实战 — 自绘圆角按钮/窗口底栏.
// 其它自绘控件可以看 Advanced 目录下的 Beautify* 系列例子.
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// ============================================================
// BeautifyButton — 自绘圆角按钮封装
// 继承 widget.Button, 通过自绘实现圆角外观, 支持鼠标悬停/按下颜色变化.
// ============================================================

// BeautifyButton 美化按钮
type BeautifyButton struct {
	*widget.Button
	text      string // 按钮文字
	font      int    // 字体句柄
	radius    int32  // 圆角半径
	colorNorm uint32 // 正常颜色
	colorStay uint32 // 悬停颜色
	colorDown uint32 // 按下颜色
	textColor uint32 // 文字颜色
}

// NewBeautifyButton 创建美化按钮.
//
// x, y: 坐标; cx, cy: 宽高; text: 文字; font: 字体句柄; hParent: 父句柄.
//
// 默认主题: 蓝色系(正常#4285F4, 悬停#286EE6, 按下#1E5AC8, 文字白色).
// 使用 SetColorNorm/SetColorStay/SetColorDown/SetTextColor 自定义颜色.
func NewBeautifyButton(x, y, cx, cy int32, text string, font int, hParent int) *BeautifyButton {
	btn := widget.NewButton(x, y, cx, cy, "", hParent)
	if btn == nil {
		return nil
	}
	b := &BeautifyButton{
		Button:    btn,
		text:      text,
		font:      font,
		radius:    8,
		colorNorm: xc.RGBA(66, 133, 244, 255),
		colorStay: xc.RGBA(40, 110, 230, 255),
		colorDown: xc.RGBA(30, 90, 200, 255),
		textColor: xc.RGBA(255, 255, 255, 255),
	}
	b.AddEvent_Paint(b.onPaint)
	return b
}

// SetText 设置按钮文字.
func (b *BeautifyButton) SetText(text string) {
	b.text = text
	xc.XEle_Redraw(b.Handle, false)
}

// GetText 获取按钮文字.
func (b *BeautifyButton) GetText() string {
	return b.text
}

// SetRadius 设置圆角半径.
func (b *BeautifyButton) SetRadius(radius int32) {
	b.radius = radius
	xc.XEle_Redraw(b.Handle, false)
}

// GetRadius 获取圆角半径.
func (b *BeautifyButton) GetRadius() int32 {
	return b.radius
}

// SetColorNorm 设置正常颜色.
func (b *BeautifyButton) SetColorNorm(color uint32) {
	b.colorNorm = color
	xc.XEle_Redraw(b.Handle, false)
}

// GetColorNorm 获取正常颜色.
func (b *BeautifyButton) GetColorNorm() uint32 {
	return b.colorNorm
}

// SetColorStay 设置悬停颜色.
func (b *BeautifyButton) SetColorStay(color uint32) {
	b.colorStay = color
	xc.XEle_Redraw(b.Handle, false)
}

// GetColorStay 获取悬停颜色.
func (b *BeautifyButton) GetColorStay() uint32 {
	return b.colorStay
}

// SetColorDown 设置按下颜色.
func (b *BeautifyButton) SetColorDown(color uint32) {
	b.colorDown = color
	xc.XEle_Redraw(b.Handle, false)
}

// GetColorDown 获取按下颜色.
func (b *BeautifyButton) GetColorDown() uint32 {
	return b.colorDown
}

// SetTextColor 设置文字颜色.
func (b *BeautifyButton) SetTextColor(color uint32) {
	b.textColor = color
	xc.XEle_Redraw(b.Handle, false)
}

// GetTextColor 获取文字颜色.
func (b *BeautifyButton) GetTextColor() uint32 {
	return b.textColor
}

// onPaint 自绘事件: 绘制圆角矩形背景 + 居中文字.
func (b *BeautifyButton) onPaint(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true // 拦截默认绘制
	draw := drawx.NewByHandle(hDraw)
	if draw == nil {
		return 0
	}
	draw.EnableSmoothingMode(true)

	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)

	bg := b.colorNorm
	switch xc.XBtn_GetStateEx(hEle) {
	case xcc.Button_State_Stay:
		bg = b.colorStay
	case xcc.Button_State_Down:
		bg = b.colorDown
	}

	draw.SetBrushColor(bg).FillRoundRect(&rc, b.radius, b.radius)

	draw.SetFont(b.font)
	draw.SetBrushColor(b.textColor)
	draw.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)
	draw.DrawText(b.text, &rc)
	return 0
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 400, 250, "Draw 自定义控件实战", 0, xcc.Window_Style_Default)
	w.SetBorderSize(1, 30, 1, 1)

	fontBtn := font.NewEX("微软雅黑", 14, xcc.FontStyle_Bold)
	fontSmall := font.NewEX("微软雅黑", 11, xcc.FontStyle_Regular)

	border := w.GetBorderSizeEx()
	offY := border.Top

	// ====== 1. 自绘圆角按钮 (蓝色) ======
	btn := NewBeautifyButton(20, offY+30, 120, 36, "确认", fontBtn.Handle, w.Handle)
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		w.MessageBox("提示", "点击了确认按钮", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	// ====== 2. 自绘危险按钮 (红色) ======
	btnDanger := NewBeautifyButton(160, offY+30, 120, 36, "删除", fontBtn.Handle, w.Handle)
	btnDanger.SetColorNorm(xc.RGBA(234, 67, 53, 255))
	btnDanger.SetColorStay(xc.RGBA(210, 50, 40, 255))
	btnDanger.SetColorDown(xc.RGBA(180, 30, 20, 255))
	btnDanger.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		w.MessageBox("警告", "危险操作不可逆！", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	// ====== 3. 窗口底栏 ======
	w.AddEvent_Paint(func(hWindow int, hDraw int, pbHandled *bool) int {
		*pbHandled = true
		xc.XWnd_DrawWindow(hWindow, hDraw)
		draw := drawx.NewByHandle(hDraw)
		draw.EnableSmoothingMode(true)

		windowW := w.GetWidth()
		barH := int32(50)
		barY := w.GetHeight() - barH - border.Bottom
		rc := xc.RECT{Left: border.Left, Top: barY, Right: windowW - border.Right, Bottom: barY + barH}
		draw.GradientFill2(
			&rc,
			xc.RGBA(245, 245, 245, 255),
			xc.RGBA(225, 225, 225, 255),
			xcc.GRADIENT_FILL_RECT_V,
		)

		draw.SetBrushColor(xc.RGBA(200, 200, 200, 255))
		draw.DrawLine(border.Left, barY, windowW, barY)

		draw.SetFont(fontSmall.Handle)
		draw.SetBrushColor(xc.RGBA(120, 120, 120, 255))
		draw.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)
		draw.DrawText("基于 Draw 自绘控件 — 圆角按钮 / 窗口底栏",
			&rc)
		return 0
	})

	w.Show()
	a.Run()
	a.Exit()
}
