// 形状文本自动根据内容改变大小, 设置字体、字体大小、字体样式
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	w := window.New(0, 0, 500, 200, "", 0, xcc.Window_Style_Default)

	st := widget.NewShapeText(15, 35, 100, 30, "测试字体大小", w.Handle)
	// 自动根据内容改变大小
	st.LayoutItem_SetWidth(xcc.Layout_Size_Auto, -1)
	st.LayoutItem_SetHeight(xcc.Layout_Size_Auto, -1)

	// 设置字体大小
	st.SetFont(font.New(50).Handle)

	// 设置个新字体， 粗体
	// st.SetFont(font.NewEX("幼圆", 50, xcc.FontStyle_Bold).Handle)

	w.Show(true)
	a.Run()
	a.Exit()
}
