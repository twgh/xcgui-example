// 框架窗口
package main

import (
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
	// 启用自动 DPI 缩放
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建框架窗口
	w := window.NewFrameWindow(0, 0, 1000, 800, "FrameWindow", 0, xcc.Window_Style_Default)

	// 创建窗格
	paneLeft := widget.NewPane("left", 200, 280, w.Handle)
	paneRight := widget.NewPane("right", 200, 280, w.Handle)
	paneBottom := widget.NewPane("bottom", 770, 170, w.Handle)

	// 把窗格句柄存到数组里
	paneList := []int{paneLeft.Handle, paneRight.Handle, paneBottom.Handle}

	// 设置窗格 ID, id 必须大于0
	for i := 0; i < len(paneList); i++ {
		xc.XWidget_SetID(paneList[i], int32(i+1))
	}

	// 添加窗格
	w.AddPane(0, paneLeft.Handle, xcc.Pane_Align_Left)
	w.AddPane(0, paneRight.Handle, xcc.Pane_Align_Right)
	w.AddPane(0, paneBottom.Handle, xcc.Pane_Align_Bottom)

	// 创建按钮, 保存布局信息文件
	btnSaveLayout := widget.NewButton(20, 50, 100, 30, "保存布局", paneLeft.Handle)
	btnSaveLayout.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		btnSaveLayout.Enable(false).Redraw(false)
		defer btnSaveLayout.Enable(true).Redraw(false)

		w.SaveLayoutToFile("FrameWindow/frameWnd_layout.xml")
		w.MessageBox("提示", "布局信息已保存到文件", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
		return 0
	})

	// 窗口关闭事件, 关闭前保存布局信息
	w.AddEvent_Close(func(hWindow int, pbHandled *bool) int {
		w.SaveLayoutToFile("FrameWindow/frameWnd_layout.xml")
		return 0
	})

	// 创建编辑框
	edit := widget.NewEdit(0, 0, 0, 0, w.Handle)
	edit.EnableMultiLine(true)
	// 设置主视图为编辑框
	w.SetView(edit.Handle)

	// 框架窗口_加载布局信息文件
	w.LoadLayoutFile(paneList, int32(len(paneList)), "FrameWindow/frameWnd_layout.xml")

	// 窗口_调整布局
	w.AdjustLayout()

	// 显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
