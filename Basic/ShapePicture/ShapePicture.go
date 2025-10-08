// 形状图片, 用按钮来代替形状图片.
package main

import (
	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed icon1.ico
var imgData []byte

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 400, 300, "ShapePicture", 0, xcc.Window_Style_Default)

	// 加载图片
	img := imagex.NewByMem(imgData)

	// 创建形状图片元素
	spic := widget.NewShapePicture(50, 50, 32, 32, w.Handle)
	// 设置图片
	spic.SetImage(img.Handle)

	// 形状元素都是没有事件的, 但有时候想要鼠标点击图片就打开一个超链接, 这时候应该用其他元素来代替.
	// 炫彩是很灵活的, 不要局限于元素类型, 大胆的去 DIY, 能用就行.
	// 这里采用万能的按钮来代替.
	// 创建按钮
	btnPic := NewImageButton(50, 120, 32, 32, w.Handle)
	// 设置图片
	btnPic.SetImage(img.Handle)

	// 注册按钮单击事件
	btnPic.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		a.Alert("提示", "图片按钮被单击了!")
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}

// ImageButton 图片按钮, 与形状图片不同的是, 它支持按钮的事件.
type ImageButton struct {
	widget.Button
}

func NewImageButton(x int32, y int32, cx int32, cy int32, hParent int) *ImageButton {
	btn := widget.NewButton(x, y, cx, cy, "", hParent)
	btn.EnableBkTransparent(true)
	i := &ImageButton{}
	i.SetHandle(btn.Handle)
	return i
}

func (i *ImageButton) SetImage(hImage int) {
	i.SetIcon(hImage)
}
