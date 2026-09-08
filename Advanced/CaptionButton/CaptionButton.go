// 标题栏自定义按钮.
// 窗口启用布局时，在标题栏上（最小化按钮左侧）放置自定义按钮。
// 这是针对启用布局后的窗口写的例子, 不过非布局状态下也是一样的。
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 640, 420, "标题栏自定义按钮示例", 0, xcc.Window_Style_Default)

	// 1) 启用窗口布局（内容区控件将自动排列）
	w.EnableLayout(true)
	w.SetPadding(12, 12, 12, 12) // 内容区内边距，避免贴边
	w.SetSpace(10).SetSpaceRow(10)

	// 2) 内容区放几个普通控件，证明布局是生效的
	widget.NewShapeText(0, 0, 200, 28, "内容区（受布局管理）:", w.Handle)
	widget.NewEdit(0, 0, 320, 30, w.Handle)

	// 3) 在标题栏上、最小化按钮左侧放一个自定义按钮
	//    关键点：EnableLayoutControl(false) 让它脱离窗口布局，使用绝对坐标。
	captionBtn := widget.NewButton(0, 0, 64, 28, "☰ 菜单", w.Handle)
	captionBtn.EnableLayoutControl(false)
	captionBtn.SetZOrder(999) // 设置Z序，防止被覆盖
	captionBtn.SetTextColor(xc.RGBA(60, 60, 60, 255))
	captionBtn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		w.MessageBox("菜单", "你点击了标题栏上的自定义按钮", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		return 0
	})

	var needRepos = true // 是否需要重新定位
	// 定位函数：读取最小化按钮坐标，把自定义按钮放到它的左侧。
	reposCaptionBtn := func() {
		hMin := w.GetButton(xcc.Window_Style_Btn_Min)
		if hMin == 0 {
			return
		}
		minRect := widget.NewButtonByHandle(hMin).GetRectEx()
		if minRect.Right <= minRect.Left { // 坐标尚未就绪
			return
		}
		btnW := int32(64)
		btnH := minRect.Bottom - minRect.Top // 与最小化按钮等高，视觉更协调
		btnX := minRect.Left - btnW - 6      // 6px 间距，位于最小化左侧
		btnY := minRect.Top
		captionBtn.SetRectEx(btnX, btnY, btnW, btnH, true, xcc.AdjustLayout_No)
	}

	// 5) 尺寸变化（含最大化/还原/拖拽缩放）时标记需要重新定位。
	w.AddEvent_Size(func(hWindow int, nFlags uint, pPt *xc.SIZE, pbHandled *bool) int {
		needRepos = true
		return 0
	})

	// 6) 在绘制完成事件里完成定位：此时窗口内置标题栏按钮坐标已是最新值，
	//    自定义按钮即可正确跟随最大化/还原。每个尺寸变化只定位一次。
	w.AddEvent_Paint_End(func(hWindow int, hDraw int, pbHandled *bool) int {
		if needRepos {
			needRepos = false
			reposCaptionBtn()
		}
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
