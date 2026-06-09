// 美化复选框(多选按钮).
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

// BeautifyCheckBox 美化复选框.
//
// 继承 widget.Button, 通过自绘实现复选框外观.
// 未选中: 灰色方框 + 白色底; 选中: 蓝色填充方框 + 白色勾号.
type BeautifyCheckBox struct {
	*widget.Button
	boxSize             int32  // 方框大小
	textOffset          int32  // 文本与方框间距
	borderColorOn       uint32 // 选中: 边框颜色(与填充色相同)
	borderColorOff      uint32 // 未选中: 边框颜色
	fillColorOn         uint32 // 选中: 填充颜色
	fillColorOff        uint32 // 未选中: 填充颜色(一般为白色)
	checkColor          uint32 // 勾号颜色(一般为白色)
	hoverBorderColorOn  uint32 // 悬停+选中: 边框颜色(高亮)
	hoverBorderColorOff uint32 // 悬停+未选中: 边框颜色(高亮)
	hoverFillColorOn    uint32 // 悬停+选中: 填充颜色(高亮)
	disabledBorderColor uint32 // 禁用: 边框颜色
	disabledFillColor   uint32 // 禁用: 填充颜色
	disabledCheckColor  uint32 // 禁用: 勾号颜色
}

// NewBeautifyCheckBox 创建美化复选框.
//
// x, y: 坐标; cx, cy: 宽高;
//
// text: 文本标签;
//
// hParent: 父句柄.
func NewBeautifyCheckBox(x, y, cx, cy int32, text string, hParent int) *BeautifyCheckBox {
	btn := widget.NewButton(x, y, cx, cy, text, hParent)
	if btn == nil {
		return nil
	}
	t := &BeautifyCheckBox{
		Button:              btn,
		boxSize:             18,
		textOffset:          8,
		borderColorOn:       xc.RGBA(24, 144, 255, 255),  // 选中: 蓝色边框
		borderColorOff:      xc.RGBA(191, 191, 191, 255), // 未选中: 灰色边框
		fillColorOn:         xc.RGBA(24, 144, 255, 255),  // 选中: 蓝色填充
		fillColorOff:        xc.RGBA(255, 255, 255, 255), // 未选中: 白色填充
		checkColor:          xc.RGBA(255, 255, 255, 255), // 勾号: 白色
		hoverBorderColorOn:  xc.RGBA(64, 169, 255, 255),  // 悬停+选中: 亮蓝边框
		hoverBorderColorOff: xc.RGBA(168, 168, 168, 255), // 悬停+未选中: 加深灰边框
		hoverFillColorOn:    xc.RGBA(64, 169, 255, 255),  // 悬停+选中: 亮蓝填充
		disabledBorderColor: xc.RGBA(217, 217, 217, 255), // 禁用: 浅灰边框
		disabledFillColor:   xc.RGBA(245, 245, 245, 255), // 禁用: 浅灰填充
		disabledCheckColor:  xc.RGBA(200, 200, 200, 255), // 禁用: 浅灰勾号
	}
	// 设置按钮类型为多选按钮
	t.SetTypeEx(xcc.Button_Type_Check)
	// 设置文本左对齐, 文本垂直居中
	t.SetTextAlign(xcc.TextAlignFlag_Left | xcc.TextAlignFlag_Vcenter | xcc.TextTrimming_EllipsisCharacter)
	// 注册自绘事件
	t.AddEvent_Paint(t.onPaint)
	return t
}

// SetBoxSize 设置方框大小.
func (t *BeautifyCheckBox) SetBoxSize(size int32) {
	t.boxSize = size
}

// GetBoxSize 获取方框大小.
func (t *BeautifyCheckBox) GetBoxSize() int32 {
	return t.boxSize
}

// SetTextOffset 设置文本与方框间距.
func (t *BeautifyCheckBox) SetTextOffset(offset int32) {
	t.textOffset = offset
}

// GetTextOffset 获取文本与方框间距.
func (t *BeautifyCheckBox) GetTextOffset() int32 {
	return t.textOffset
}

// SetBorderColorOn 设置选中状态边框颜色.
func (t *BeautifyCheckBox) SetBorderColorOn(color uint32) {
	t.borderColorOn = color
}

// GetBorderColorOn 获取选中状态边框颜色.
func (t *BeautifyCheckBox) GetBorderColorOn() uint32 {
	return t.borderColorOn
}

// SetBorderColorOff 设置未选中状态边框颜色.
func (t *BeautifyCheckBox) SetBorderColorOff(color uint32) {
	t.borderColorOff = color
}

// GetBorderColorOff 获取未选中状态边框颜色.
func (t *BeautifyCheckBox) GetBorderColorOff() uint32 {
	return t.borderColorOff
}

// SetFillColorOn 设置选中状态填充颜色.
func (t *BeautifyCheckBox) SetFillColorOn(color uint32) {
	t.fillColorOn = color
}

// GetFillColorOn 获取选中状态填充颜色.
func (t *BeautifyCheckBox) GetFillColorOn() uint32 {
	return t.fillColorOn
}

// SetFillColorOff 设置未选中状态填充颜色.
func (t *BeautifyCheckBox) SetFillColorOff(color uint32) {
	t.fillColorOff = color
}

// GetFillColorOff 获取未选中状态填充颜色.
func (t *BeautifyCheckBox) GetFillColorOff() uint32 {
	return t.fillColorOff
}

// SetCheckColor 设置勾号颜色.
func (t *BeautifyCheckBox) SetCheckColor(color uint32) {
	t.checkColor = color
}

// GetCheckColor 获取勾号颜色.
func (t *BeautifyCheckBox) GetCheckColor() uint32 {
	return t.checkColor
}

// SetHoverBorderColorOn 设置悬停+选中状态边框颜色.
func (t *BeautifyCheckBox) SetHoverBorderColorOn(color uint32) {
	t.hoverBorderColorOn = color
}

// GetHoverBorderColorOn 获取悬停+选中状态边框颜色.
func (t *BeautifyCheckBox) GetHoverBorderColorOn() uint32 {
	return t.hoverBorderColorOn
}

// SetHoverBorderColorOff 设置悬停+未选中状态边框颜色.
func (t *BeautifyCheckBox) SetHoverBorderColorOff(color uint32) {
	t.hoverBorderColorOff = color
}

// GetHoverBorderColorOff 获取悬停+未选中状态边框颜色.
func (t *BeautifyCheckBox) GetHoverBorderColorOff() uint32 {
	return t.hoverBorderColorOff
}

// SetHoverFillColorOn 设置悬停+选中状态填充颜色.
func (t *BeautifyCheckBox) SetHoverFillColorOn(color uint32) {
	t.hoverFillColorOn = color
}

// GetHoverFillColorOn 获取悬停+选中状态填充颜色.
func (t *BeautifyCheckBox) GetHoverFillColorOn() uint32 {
	return t.hoverFillColorOn
}

// SetDisabledBorderColor 设置禁用状态边框颜色.
func (t *BeautifyCheckBox) SetDisabledBorderColor(color uint32) {
	t.disabledBorderColor = color
}

// GetDisabledBorderColor 获取禁用状态边框颜色.
func (t *BeautifyCheckBox) GetDisabledBorderColor() uint32 {
	return t.disabledBorderColor
}

// SetDisabledFillColor 设置禁用状态填充颜色.
func (t *BeautifyCheckBox) SetDisabledFillColor(color uint32) {
	t.disabledFillColor = color
}

// GetDisabledFillColor 获取禁用状态填充颜色.
func (t *BeautifyCheckBox) GetDisabledFillColor() uint32 {
	return t.disabledFillColor
}

// SetDisabledCheckColor 设置禁用状态勾号颜色.
func (t *BeautifyCheckBox) SetDisabledCheckColor(color uint32) {
	t.disabledCheckColor = color
}

// GetDisabledCheckColor 获取禁用状态勾号颜色.
func (t *BeautifyCheckBox) GetDisabledCheckColor() uint32 {
	return t.disabledCheckColor
}

// drawCheckmark 绘制勾号(✓)路径.
func (t *BeautifyCheckBox) drawCheckmark(draw *drawx.Draw, left, top, size int32, color uint32) {
	// 勾号路径: 三根折线, 比例参考标准复选框勾号
	strokeWidth := int32(2)
	if size > 24 {
		strokeWidth = int32(float32(size) / 12.0)
	}

	draw.SetLineWidth(strokeWidth)
	draw.SetBrushColor(color)

	// 使用三点折线绘制勾号
	// 比例: 起笔在左下约(0.18,0.55), 转折点在(0.42,0.78), 终点在(0.82,0.22)
	margin := int32(float32(size)*0.15 + 0.5) // 内边距
	p1x := left + margin
	p1y := top + int32(float32(size)*0.55+0.5)
	p2x := left + int32(float32(size)*0.42+0.5)
	p2y := top + int32(float32(size)*0.78+0.5)
	p3x := left + size - margin
	p3y := top + int32(float32(size)*0.22+0.5)

	// 绘制两条线段模拟勾号
	draw.DrawLine(p1x, p1y, p2x, p2y)
	draw.DrawLine(p2x, p2y, p3x, p3y)
}

// onPaint 自绘事件: 绘制方框指示器 + 勾号 + 文本标签.
func (t *BeautifyCheckBox) onPaint(hEle int, hDraw int, pbHandled *bool) int {
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

	// 方框指示器居中于左侧
	margin := int32(2)
	boxSize := t.boxSize
	boxTop := (h - boxSize) / 2
	boxBottom := boxTop + boxSize
	boxLeft := margin
	boxRight := boxLeft + boxSize

	var borderColor, fillColor, checkColor uint32

	if !isEnabled {
		borderColor = t.disabledBorderColor
		fillColor = t.disabledFillColor
		checkColor = t.disabledCheckColor
	} else if isHover && isCheck {
		borderColor = t.hoverBorderColorOn
		fillColor = t.hoverFillColorOn
		checkColor = t.checkColor
	} else if isHover && !isCheck {
		borderColor = t.hoverBorderColorOff
		fillColor = t.fillColorOff
		checkColor = t.checkColor
	} else if isCheck {
		borderColor = t.borderColorOn
		fillColor = t.fillColorOn
		checkColor = t.checkColor
	} else {
		borderColor = t.borderColorOff
		fillColor = t.fillColorOff
		checkColor = t.checkColor
	}

	// 1. 绘制方框背景(填充)
	roundRadius := boxSize / 4 // 小圆角
	var boxRC xc.RECT
	boxRC.Left = boxLeft
	boxRC.Top = boxTop
	boxRC.Right = boxRight
	boxRC.Bottom = boxBottom
	draw.SetBrushColor(fillColor)
	draw.FillRoundRect(&boxRC, roundRadius, roundRadius)

	// 2. 绘制方框边框
	borderThickness := int32(2)
	draw.SetLineWidth(borderThickness)
	draw.SetBrushColor(borderColor)
	draw.DrawRoundRect(&boxRC, roundRadius, roundRadius)

	// 3. 绘制勾号(选中状态)
	if isCheck {
		t.drawCheckmark(draw, boxLeft, boxTop, boxSize, checkColor)
	}

	// 4. 绘制文本标签
	text := t.GetText()
	if text != "" {
		textLeft := boxRight + t.textOffset
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

	w := window.New(0, 0, 500, 320, "美化复选框示例", 0, xcc.Window_Style_Default)
	w.SetBorderSize(2, 30, 2, 2)

	// 第一组: 默认蓝色主题
	widget.NewShapeText(30, 35, 200, 20, "默认蓝色主题", w.Handle)

	cb1 := NewBeautifyCheckBox(30, 60, 130, 26, "选项 1", w.Handle)
	cb1.SetCheck(true) // 默认选中

	cb2 := NewBeautifyCheckBox(30, 90, 130, 26, "选项 2", w.Handle)

	cb3 := NewBeautifyCheckBox(30, 120, 130, 26, "选项 3", w.Handle)

	// 选中事件演示
	cb1.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		status := "已选中"
		if !bCheck {
			status = "未选中"
		}
		w.MessageBox("提示", "选项 1: 状态: "+status, xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	// 选项 2 选中事件
	cb2.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		status := "已选中"
		if !bCheck {
			status = "未选中"
		}
		w.MessageBox("提示", "选项 2: 状态: "+status, xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	// 选项 3 选中事件
	cb3.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		status := "已选中"
		if !bCheck {
			status = "未选中"
		}
		w.MessageBox("提示", "选项 3: 状态: "+status, xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	// 第二组: 绿色主题 + 禁用状态
	widget.NewShapeText(30, 160, 200, 20, "自定义绿色主题", w.Handle)

	cb4 := NewBeautifyCheckBox(30, 185, 220, 26, "选项 A (自定义颜色)", w.Handle)
	cb4.SetBorderColorOn(xc.RGBA(82, 196, 26, 255))
	cb4.SetFillColorOn(xc.RGBA(82, 196, 26, 255))
	cb4.SetHoverBorderColorOn(xc.RGBA(115, 209, 61, 255))
	cb4.SetHoverFillColorOn(xc.RGBA(115, 209, 61, 255))
	cb4.SetCheck(true)

	cb5 := NewBeautifyCheckBox(30, 215, 180, 26, "选项 B (禁用选中)", w.Handle)
	cb5.Enable(false)
	cb5.SetCheck(true)

	cb6 := NewBeautifyCheckBox(30, 245, 180, 26, "选项 C (禁用未选中)", w.Handle)
	cb6.Enable(false)
	cb6.SetCheck(false)

	// 第二个禁用状态展示: 不同颜色主题
	cb7 := NewBeautifyCheckBox(280, 215, 180, 26, "选项 D (禁用选中2)", w.Handle)
	cb7.Enable(false)
	cb7.SetCheck(true)

	w.Show(true)
	a.Run()
	a.Exit()
}
