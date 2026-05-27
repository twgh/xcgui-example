package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// WinUI3Edit WinUI 3 风格编辑框封装
type WinUI3Edit struct {
	*widget.Edit
	themeColor   uint32 // 主题色
	normalColor  uint32 // 正常状态边框色
	hoverColor   uint32 // 悬浮状态边框色
	focusColor   uint32 // 焦点状态边框色
	bgColor      uint32 // 背景色
	textColor    uint32 // 文本颜色
	isPassword   bool   // 是否密码框
	passwordChar uint16 // 密码字符
	hParent      int
}

// NewWinUI3Edit 创建 WinUI 3 风格编辑框
func NewWinUI3Edit(x, y, cx, cy int32, hParent int, themeColor uint32) *WinUI3Edit {
	edit := widget.NewEdit(x, y, cx, cy, hParent)

	w := &WinUI3Edit{
		Edit:         edit,
		themeColor:   themeColor,
		normalColor:  xc.RGBA(200, 200, 200, 255),
		hoverColor:   xc.RGBA(118, 118, 118, 255),
		focusColor:   themeColor,
		bgColor:      xc.RGBA(255, 255, 255, 255),
		textColor:    xc.RGBA(0, 0, 0, 255),
		passwordChar: '*',
		hParent:      hParent,
	}

	w.initStyle()
	return w
}

// initStyle 初始化 WinUI 3 样式
func (w *WinUI3Edit) initStyle() {
	w.SetTextColor(w.textColor)
	w.SetDefaultTextColor(xc.RGBA(150, 150, 150, 255))

	bkMgr := w.GetBkManagerObj()
	if bkMgr == nil {
		return
	}

	bkMgr.Clear()

	// 正常状态
	bkMgr.AddFill(xcc.Element_State_Flag_Nothing, w.bgColor, 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Nothing, w.normalColor, 2, 0)

	// 悬浮状态
	bkMgr.AddFill(xcc.Element_State_Flag_Stay, w.bgColor, 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Stay, w.hoverColor, 2, 0)

	// 获得焦点状态
	bkMgr.AddFill(xcc.Element_State_Flag_Focus, w.bgColor, 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Focus, w.focusColor, 2, 0)

	// 禁用状态
	bkMgr.AddFill(xcc.Element_State_Flag_Disable, xc.RGBA(245, 245, 245, 255), 0)
	bkMgr.AddBorder(xcc.Element_State_Flag_Disable, xc.RGBA(220, 220, 220, 255), 2, 0)

	// 设置内边距
	xc.XEle_SetPadding(w.Handle, 8, 4, 8, 4)

	// 设置插入符颜色
	w.SetCaretColor(w.focusColor)
}

// SetThemeColor 设置主题色
func (w *WinUI3Edit) SetThemeColor(color uint32) *WinUI3Edit {
	w.themeColor = color
	w.focusColor = color
	w.initStyle()
	w.Redraw(true)
	return w
}

// EnableWinUI3Password 启用 WinUI 3 风格密码框
func (w *WinUI3Edit) EnableWinUI3Password(bEnable ...bool) *WinUI3Edit {
	enable := true
	if len(bEnable) > 0 {
		enable = bEnable[0]
	}
	w.isPassword = enable
	w.EnablePassword(enable)
	if enable {
		w.SetPasswordCharacter(w.passwordChar)
	}
	return w
}

// ============================================================
// 示例主函数
// ============================================================
func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 500, 450, "WinUI 3 风格编辑框示例", 0, xcc.Window_Style_Default)

	// WinUI 3 主题色 (蓝色)
	themeColor := xc.RGBA(0, 120, 212, 255)

	// 1. 普通编辑框
	edit1 := NewWinUI3Edit(30, 80, 300, 36, w.Handle, themeColor)
	edit1.SetDefaultText("请输入用户名")

	// 2. 密码框
	edit2 := NewWinUI3Edit(30, 140, 300, 36, w.Handle, themeColor)
	edit2.SetDefaultText("请输入密码")
	edit2.EnableWinUI3Password(true)

	// 3. 只读编辑框
	edit3 := NewWinUI3Edit(30, 200, 300, 36, w.Handle, themeColor)
	edit3.SetText("这是只读文本")
	edit3.EnableReadOnly(true)

	// 4. 不同主题色的编辑框 (绿色主题)
	edit4 := NewWinUI3Edit(30, 260, 300, 36, w.Handle, xc.RGBA(0, 180, 80, 255))
	edit4.SetDefaultText("绿色主题编辑框")

	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}
