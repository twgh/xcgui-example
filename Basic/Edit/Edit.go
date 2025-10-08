// 编辑框, 多行编辑框, 密码框, 数字框
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 340, 280, "", 0, xcc.Window_Style_Default)

	// 1.普通编辑框
	edit := widget.NewEdit(12, 35, 150, 30, w.Handle)
	// 设置文本颜色
	edit.SetTextColor(xc.RGBA(236, 64, 122, 255))
	edit.SetText("hello, 普通编辑框")

	// 2.密码输入框
	editPwd := widget.NewEdit(12, 75, 150, 30, w.Handle)
	editPwd.EnablePassword(true)
	editPwd.SetText("pwd")
	// 这个可以改变密码字符, 不改的话默认是*
	editPwd.SetPasswordCharacter('#')
	// 创建眼睛按钮
	eyeBtn := widget.NewButton(165, 75, 30, 30, "", w.Handle)
	eyeBtn.EnableBkTransparent(true)
	// 加载图片, 禁止自动销毁
	imgShow := imagex.NewBySvgStringW(eye_show).EnableAutoDestroy(false)
	imgHide := imagex.NewBySvgStringW(eye_hide).EnableAutoDestroy(false)
	eyeBtn.SetIcon(imgHide.Handle)
	// 按钮事件, 控制是否显示密码
	eyeBtn.Event_BnClick(func(pbHandled *bool) int {
		// 这样记录不用新创建一个变量, 是自定义的, 你就当成在读写元素内置的一个map, 就是用来让你方便记录东西的
		if eyeBtn.GetProperty("isPwd") != "0" {
			eyeBtn.SetProperty("isPwd", "0")
			eyeBtn.SetIcon(imgShow.Handle)
			editPwd.EnablePassword(false)
		} else {
			eyeBtn.SetProperty("isPwd", "1")
			eyeBtn.SetIcon(imgHide.Handle)
			editPwd.EnablePassword(true)
		}
		editPwd.Redraw(true)
		eyeBtn.Redraw(true)
		return 0
	})

	// 3.多行编辑框
	editMultiLine := widget.NewEdit(12, 115, 300, 100, w.Handle)
	editMultiLine.EnableAutoShowScrollBar(true)
	editMultiLine.EnableMultiLine(true)
	editMultiLine.EnableAutoWrap(true)
	editMultiLine.AddText("你好, 世界, 我是普通文本")
	// 添加样式, 需要注意的是Arial字体不支持中文
	style1 := editMultiLine.AddStyleEx("Arial", 14, xcc.FontStyle_Bold, xc.RGBA(0, 191, 165, 255), true)
	// 添加带样式的文本
	editMultiLine.AddTextEx("\nhello world, I am styled text!", style1)
	// 获取编辑框文本
	fmt.Println("获取编辑框文本:", editMultiLine.GetTextEx())

	// 设置定时器, 循环获取多行编辑框鼠标选中的文本
	editMultiLine.SetXCTimer(111, 1000)
	editMultiLine.Event_XC_TIMER(func(nTimerID int, pbHandled *bool) int {
		if nTimerID == 111 {
			text := editMultiLine.GetSelectTextEx()
			if text != "" {
				fmt.Println("选中的文本:", text)
			}
		}
		return 0
	})

	// 4.只能输入数字的编辑框
	editOnlyNumber := widget.NewEdit(12, 225, 150, 30, w.Handle)
	editOnlyNumber.SetDefaultText("只能输入数字的编辑框")
	editOnlyNumber.Event_CHAR(func(wParam, lParam uintptr, pbHandled *bool) int {
		fmt.Println(wParam)
		if wParam < 58 && wParam > 47 { // 0-9
			return 0
		}

		switch wParam { // 放行删除,复制,撤销,剪切,全选. 有其它需要的自己添加.
		case 8, 3, 26, 24, 1:
			return 0
		}

		*pbHandled = true // 其它的都拦截
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}

const (
	eye_hide = `<svg t="1720166247153" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4295" width="30" height="30"><path d="M332.8 729.6l34.133333-34.133333c42.666667 12.8 93.866667 21.333333 145.066667 21.333333 162.133333 0 285.866667-68.266667 375.466667-213.333333-46.933333-72.533333-102.4-128-166.4-162.133334l29.866666-29.866666c72.533333 42.666667 132.266667 106.666667 183.466667 192-98.133333 170.666667-243.2 256-426.666667 256-59.733333 4.266667-119.466667-8.533333-174.933333-29.866667z m-115.2-64c-51.2-38.4-93.866667-93.866667-132.266667-157.866667 98.133333-170.666667 243.2-256 426.666667-256 38.4 0 76.8 4.266667 110.933333 12.8l-34.133333 34.133334c-25.6-4.266667-46.933333-4.266667-76.8-4.266667-162.133333 0-285.866667 68.266667-375.466667 213.333333 34.133333 51.2 72.533333 93.866667 115.2 128l-34.133333 29.866667z m230.4-46.933333l29.866667-29.866667c8.533333 4.266667 21.333333 4.266667 29.866666 4.266667 46.933333 0 85.333333-38.4 85.333334-85.333334 0-12.8 0-21.333333-4.266667-29.866666l29.866667-29.866667c12.8 17.066667 17.066667 38.4 17.066666 64 0 72.533333-55.466667 128-128 128-17.066667-4.266667-38.4-12.8-59.733333-21.333333zM384 499.2c4.266667-68.266667 55.466667-119.466667 123.733333-123.733333 0 4.266667-123.733333 123.733333-123.733333 123.733333zM733.866667 213.333333l29.866666 29.866667-512 512-34.133333-29.866667L733.866667 213.333333z" fill="#444444" p-id="4296"></path></svg>`
	eye_show = `<svg t="1720166310093" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4449" width="30" height="30"><path d="M512 298.666667c-162.133333 0-285.866667 68.266667-375.466667 213.333333 89.6 145.066667 213.333333 213.333333 375.466667 213.333333s285.866667-68.266667 375.466667-213.333333c-89.6-145.066667-213.333333-213.333333-375.466667-213.333333z m0 469.333333c-183.466667 0-328.533333-85.333333-426.666667-256 98.133333-170.666667 243.2-256 426.666667-256s328.533333 85.333333 426.666667 256c-98.133333 170.666667-243.2 256-426.666667 256z m0-170.666667c46.933333 0 85.333333-38.4 85.333333-85.333333s-38.4-85.333333-85.333333-85.333333-85.333333 38.4-85.333333 85.333333 38.4 85.333333 85.333333 85.333333z m0 42.666667c-72.533333 0-128-55.466667-128-128s55.466667-128 128-128 128 55.466667 128 128-55.466667 128-128 128z" fill="#444444" p-id="4450"></path></svg>`
)
