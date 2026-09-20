// 形状对象GIF
package main

// 演示 ShapeGif 的使用: 创建形状GIF播放动画, 并演示透明度, 显示/隐藏, 移动位置等操作.
//
// ShapeGif 与 GifPlayer 的区别:
//   - ShapeGif 是形状对象, 跟随父元素/窗口一起绘制, 没有独立的事件, 不能响应鼠标键盘等交互,
//     由界面库内部自动循环播放 GIF 帧, 适合做界面装饰动画(加载提示, 背景动效等).
//   - GifPlayer 是 Go 层实现的播放器, 需要挂到一个元素上, 由协程逐帧驱动,
//     提供播放/暂停/继续/停止/帧回调等控制, 适合需要精确控制播放逻辑的场景.
//   - ShapeGif 自身只有创建, 置图/取图等少量方法, 没有暂停/播放速度控制;
//     显示/隐藏, 移动, 透明度等通过继承的 Shape 方法实现.

import (
	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//go:embed bg1.gif
var bg1 []byte

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 560, 400, "ShapeGif", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// 从内存加载 GIF 图片
	img := imagex.NewByMem(bg1)
	if img == nil {
		panic("加载 GIF 失败")
	}

	// 创建形状GIF_1: 正常播放, GIF 动画会被缩放到形状对象的大小
	sg1 := widget.NewShapeGif(70, 50, 180, 180, w.Handle)
	sg1.SetImage(img.Handle)

	// 创建形状GIF_2: 半透明播放
	sg2 := widget.NewShapeGif(310, 50, 180, 180, w.Handle)
	sg2.SetImage(img.Handle)
	// 设置透明度, 255 为不透明
	sg2.SetAlpha(128)

	// 说明文字, 形状文本也是形状对象, 同样跟随父元素绘制
	txt1 := widget.NewShapeText(70, 236, 180, 22, "正常播放", w.Handle)
	txt1.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)
	txt1.SetTextColor(xc.RGBA(128, 128, 128, 255))
	txt2 := widget.NewShapeText(310, 236, 180, 22, "半透明 (Alpha=128)", w.Handle)
	txt2.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)
	txt2.SetTextColor(xc.RGBA(128, 128, 128, 255))

	// 按钮: 显示/隐藏 右侧的 ShapeGif
	btnShow := widget.NewButton(70, 290, 110, 30, "隐藏右侧", w.Handle)
	btnShow.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if sg2.IsShow() {
			sg2.Show(false)
			btnShow.SetText("显示右侧")
		} else {
			sg2.Show(true)
			btnShow.SetText("隐藏右侧")
		}
		btnShow.Redraw(false)
		return 0
	})

	// 按钮: 切换右侧 ShapeGif 的透明度
	btnAlpha := widget.NewButton(230, 290, 110, 30, "不透明", w.Handle)
	btnAlpha.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if sg2.GetAlpha() == 128 {
			sg2.SetAlpha(255)
			btnAlpha.SetText("半透明")
		} else {
			sg2.SetAlpha(128)
			btnAlpha.SetText("不透明")
		}
		sg2.Redraw()
		btnAlpha.Redraw(false)
		return 0
	})

	// 按钮: 移动左侧 ShapeGif 的位置
	btnMove := widget.NewButton(390, 290, 110, 30, "移动左侧", w.Handle)
	moved := false
	btnMove.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if moved {
			sg1.SetPosition(70, 50)
		} else {
			sg1.SetPosition(190, 50)
		}
		moved = !moved
		w.Redraw() // 移动后要刷新窗口, 否则 ShapeGif 会在原位置有残留
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
