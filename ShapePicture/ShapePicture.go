// 形状图片, 用按钮来代替形状图片.
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 400, 300, "ShapePicture", 0, xcc.Window_Style_Default)

	// 加载图片
	hImage := imagex.NewByFile("ShapePicture/icon1.ico").Handle

	// 创建形状图片元素
	sp := widget.NewShapePicture(50, 50, 32, 32, w.Handle)
	// 设置图片
	sp.SetImage(hImage)

	// 形状元素都是没有事件的, 但有时候想要鼠标点击图片就打开一个超链接, 这时候应该用其他元素来代替.
	// 炫彩是很灵活的, 不要局限于元素类型, 大胆的去DIY, 能用就行.
	// 这里采用万能的按钮来代替.
	// 创建按钮
	btnPic := widget.NewButton(50, 120, 32, 32, "", w.Handle)

	// 按钮启用背景透明
	btnPic.EnableBkTransparent(true)
	// 按钮添加背景图片, 两种方式都可以
	btnPic.SetIcon(hImage)
	// btnPic.AddBkImage(xcc.Element_State_Flag_Enable, hImage)

	// 注册按钮单击事件
	btnPic.Event_BnClick(func(pbHandled *bool) int {
		a.Alert("提示", "图片按钮被单击了!")
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
