// 美化编辑框(使用背景管理器).
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BeautifyEdit 美化编辑框封装
type BeautifyEdit struct {
	*widget.Edit
	normalColor uint32 // 正常状态边框色
	hoverColor  uint32 // 悬浮状态边框色
	focusColor  uint32 // 焦点状态边框色
	bgColor     uint32 // 背景色
	textColor   uint32 // 文本颜色
}

// NewBeautifyEdit 创建美化编辑框
func NewBeautifyEdit(x, y, cx, cy int32, hParent int) *BeautifyEdit {
	edit := widget.NewEdit(x, y, cx, cy, hParent)

	w := &BeautifyEdit{
		Edit:        edit,
		normalColor: xc.RGBA(200, 200, 200, 255),
		hoverColor:  xc.RGBA(118, 118, 118, 255),
		focusColor:  xc.RGBA(0, 120, 212, 255), // 默认主题色(蓝色)
		bgColor:     xc.RGBA(255, 255, 255, 255),
		textColor:   xc.RGBA(0, 0, 0, 255),
	}

	w.initStyle()
	return w
}

// initStyle 初始化样式
func (b *BeautifyEdit) initStyle() {
	b.SetTextColor(b.textColor)
	b.SetDefaultTextColor(xc.RGBA(150, 150, 150, 255))

	bkMgr := b.GetBkManagerObj()
	if bkMgr == nil {
		return
	}

	bkMgr.Clear()

	// 正常状态
	bkMgr.AddFill(xcc.Element_State_Flag_Nothing, b.bgColor, 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Nothing, b.normalColor, 2, 0)

	// 悬浮状态
	bkMgr.AddFill(xcc.Element_State_Flag_Stay, b.bgColor, 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Stay, b.hoverColor, 2, 0)

	// 获得焦点状态
	bkMgr.AddFill(xcc.Element_State_Flag_Focus, b.bgColor, 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Focus, b.focusColor, 2, 0)

	// 禁用状态
	bkMgr.AddFill(xcc.Element_State_Flag_Disable, xc.RGBA(245, 245, 245, 255), 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Disable, xc.RGBA(220, 220, 220, 255), 2, 0)

	// 设置边框大小, 为了调整文本的位置(包括默认文本)
	b.SetBorderSize(8, 4, 8, 4)

	// 设置插入符颜色
	b.SetCaretColor(b.focusColor)
}

// SetThemeColor 设置主题色
func (b *BeautifyEdit) SetThemeColor(color uint32) *BeautifyEdit {
	b.focusColor = color
	b.initStyle()
	b.Redraw(false)
	return b
}

// GetNormalColor 获取正常状态边框色
func (b *BeautifyEdit) GetNormalColor() uint32 { return b.normalColor }

// SetNormalColor 设置正常状态边框色，设置后自动刷新
func (b *BeautifyEdit) SetNormalColor(color uint32) *BeautifyEdit {
	b.normalColor = color
	b.initStyle()
	b.Redraw(false)
	return b
}

// GetHoverColor 获取悬浮状态边框色
func (b *BeautifyEdit) GetHoverColor() uint32 { return b.hoverColor }

// SetHoverColor 设置悬浮状态边框色，设置后自动刷新
func (b *BeautifyEdit) SetHoverColor(color uint32) *BeautifyEdit {
	b.hoverColor = color
	b.initStyle()
	b.Redraw(false)
	return b
}

// GetFocusColor 获取焦点状态边框色
func (b *BeautifyEdit) GetFocusColor() uint32 { return b.focusColor }

// SetFocusColor 设置焦点状态边框色，设置后自动刷新
func (b *BeautifyEdit) SetFocusColor(color uint32) *BeautifyEdit {
	b.focusColor = color
	b.initStyle()
	b.Redraw(false)
	return b
}

// GetBgColor 获取背景色
func (b *BeautifyEdit) GetBgColor() uint32 { return b.bgColor }

// SetBgColor 设置背景色，设置后自动刷新
func (b *BeautifyEdit) SetBgColor(color uint32) *BeautifyEdit {
	b.bgColor = color
	b.initStyle()
	b.Redraw(false)
	return b
}

// GetTextColor 获取文本颜色
func (b *BeautifyEdit) GetTextColor() uint32 { return b.textColor }

// SetTextColor 设置文本颜色
func (b *BeautifyEdit) SetTextColor(color uint32) *BeautifyEdit {
	b.textColor = color
	b.Edit.SetTextColor(color)
	return b
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 500, 450, "美化编辑框示例", 0, xcc.Window_Style_Default)

	// 主题色 (蓝色)
	themeColor := xc.RGBA(0, 120, 212, 255)

	// 1. 普通编辑框
	edit1 := NewBeautifyEdit(30, 80, 300, 36, w.Handle)
	edit1.SetThemeColor(themeColor).SetDefaultText("请输入用户名")

	// 2. 密码框
	edit2 := NewBeautifyEdit(30, 140, 300, 36, w.Handle)
	edit2.SetThemeColor(themeColor).SetDefaultText("请输入密码")
	edit2.EnablePassword(true)

	// 3. 只读编辑框
	edit3 := NewBeautifyEdit(30, 200, 300, 36, w.Handle)
	edit3.SetThemeColor(themeColor).SetText("这是只读文本")
	edit3.EnableReadOnly(true)

	// 4. 不同主题色的编辑框 (绿色主题)
	edit4 := NewBeautifyEdit(30, 260, 300, 36, w.Handle)
	edit4.SetThemeColor(xc.RGBA(0, 180, 80, 255)).SetDefaultText("绿色主题编辑框")

	// 5. 紫色主题, 浅黄背景, 橙色焦点
	edit5 := NewBeautifyEdit(30, 320, 300, 36, w.Handle)
	edit5.SetDefaultText("自定义颜色编辑框")
	edit5.SetThemeColor(xc.RGBA(120, 50, 180, 255)).
		SetNormalColor(xc.RGBA(180, 130, 220, 255)).
		SetHoverColor(xc.RGBA(140, 90, 190, 255)).
		SetFocusColor(xc.RGBA(255, 120, 50, 255)).
		SetBgColor(xc.RGBA(255, 248, 230, 255)).
		SetTextColor(xc.RGBA(60, 30, 90, 255))

	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}
