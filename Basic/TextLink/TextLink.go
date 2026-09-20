// 文本链接按钮
package main

// 文本链接(TextLink)是一个静态文本链接按钮, 外观像网页里的超链接文本,
// 与普通形状文本(ShapeText)的区别是: 它是可以点击的, 能响应鼠标事件,
// 并且支持鼠标停留(悬停)/离开两种状态下的下划线和文本颜色设置,
// 常用于"忘记密码"、"查看协议"、"打开官网"这类界面链接.
//
// 本例演示: 创建多个不同样式的文本链接, 注册点击事件, 点击后用 MessageBox 提示,
// 并与普通 ShapeText 对比.

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 480, 340, "TextLink - 文本链接", 0, xcc.Window_Style_Default)

	// 用形状文本显示点击结果
	stResult := widget.NewShapeText(20, 240, 440, 48, "点击结果: 还没有点击任何链接", w.Handle)
	stResult.SetTextAlign(xcc.TextAlignFlag_Left)

	// 记录结果到界面并弹窗提示
	showResult := func(name string) {
		text := fmt.Sprintf("点击了%s", name)
		stResult.SetText("点击结果: " + text)
		// 修改形状文本显示内容后必须重绘
		stResult.Redraw()
		w.MessageBox("提示", text, xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
	}

	// 说明: 普通形状文本只能显示, 不能点击, 也没有事件
	widget.NewShapeText(20, 40, 440, 24, "下面是普通形状文本(ShapeText), 只能显示, 点击无任何反应:", w.Handle)
	stPlain := widget.NewShapeText(40, 70, 300, 24, "我只是一段普通文本, 点我无效", w.Handle)
	stPlain.SetTextColor(xc.RGBA(120, 120, 120, 255))

	// 说明: TextLink 可以点击并响应事件
	widget.NewShapeText(20, 110, 440, 24, "下面是文本链接(TextLink), 可以点击并响应事件:", w.Handle)

	// 链接1: 蓝色字体, 离开和悬停状态都显示下划线, 悬停时变红色
	link1 := widget.NewTextLink(40, 145, 140, 24, "炫彩界面库官网", w.Handle)
	link1.SetTextColor(xc.RGBA(0, 102, 204, 255))           // 离开状态文本颜色
	link1.EnableUnderlineLeave(true)                        // 鼠标离开状态显示下划线
	link1.SetUnderlineColorLeave(xc.RGBA(0, 102, 204, 255)) // 离开状态的下划线颜色
	link1.EnableUnderlineStay(true)                         // 鼠标悬停状态显示下划线
	link1.SetTextColorStay(xc.RGBA(220, 60, 60, 255))       // 悬停状态文本颜色
	link1.SetUnderlineColorStay(xc.RGBA(220, 60, 60, 255))  // 悬停状态的下划线颜色
	link1.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		showResult("炫彩界面库官网链接")
		return 0
	})

	// 链接2: 绿色字体, 只有悬停时才显示下划线
	link2 := widget.NewTextLink(40, 180, 140, 24, "查看使用协议", w.Handle)
	link2.SetTextColor(xc.RGBA(0, 150, 90, 255))
	link2.EnableUnderlineLeave(false)                 // 鼠标离开状态不显示下划线
	link2.EnableUnderlineStay(true)                   // 鼠标悬停时才显示下划线
	link2.SetTextColorStay(xc.RGBA(0, 190, 120, 255)) // 悬停时变亮绿色
	link2.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		showResult("查看使用协议链接")
		return 0
	})

	// 链接3: 使用默认样式, 不做任何设置
	link3 := widget.NewTextLink(40, 215, 140, 24, "忘记密码?", w.Handle)
	link3.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		showResult("忘记密码链接")
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
