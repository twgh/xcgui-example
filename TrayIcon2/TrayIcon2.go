// 托盘图标2. 使用 xc.XTrayIcon_ 系列函数来操作.
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wnd"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"math/rand"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	defer a.Exit()
	a.EnableDPI(true)
	a.EnableAutoDPI(true)

	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "TrayIcon2", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// 窗口设置边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// 窗口启用布局
	w.EnableLayout(true)
	// 窗口设置左右两边的内填充大小
	w.SetPadding(10, 0, 10, 0)
	// 窗口设置布局盒子垂直居中
	w.SetAlignV(xcc.Layout_Align_Center)
	// 窗口设置内部元素间距
	w.SetSpace(8)
	// 窗口设置内部元素行间距
	w.SetSpaceRow(20)

	// 创建并显示
	btnAdd := widget.NewButton(0, 0, 80, 30, "创建并显示", w.Handle)
	// 显示或隐藏
	btnShow := widget.NewButton(0, 0, 80, 30, "隐藏", w.Handle)
	btnShow.Enable(false)
	// 重置
	btnReset := widget.NewButton(0, 0, 80, 30, "重置", w.Handle)
	btnReset.Enable(false)
	// 置焦点
	btnFocus := widget.NewButton(0, 0, 80, 30, "置焦点", w.Handle)
	btnFocus.Enable(false)

	// 修改图标和提示信息
	btnMod := widget.NewButton(0, 0, 160, 30, "修改图标和提示信息", w.Handle)
	btnMod.Enable(false)
	// 设置弹出气泡消息
	btnMsg := widget.NewButton(0, 0, 160, 30, "设置弹出气泡消息", w.Handle)
	btnMsg.Enable(false)

	// 加载图标
	hIcon1 := wutil.HIcon("TrayIcon2/icon1.ico")
	fmt.Println("hIcon1:", hIcon1, "LastErr:", syscall.GetLastError())
	hIcon2 := wutil.HIcon("TrayIcon2/icon2.ico")
	fmt.Println("hIcon2:", hIcon2, "LastErr:", syscall.GetLastError())

	// 创建并显示托盘图标
	btnAdd.Event_BnClick(func(pbHandled *bool) int {
		btnAdd.Enable(false).Redraw(false)

		// 托盘图标_置图标
		xc.XTrayIcon_SetIcon(hIcon1)
		// 托盘图标_置提示文本
		xc.XTrayIcon_SetTips("托盘提示信息")
		// 添加
		xc.XTrayIcon_Add(w.Handle, 111) // 自定义的托盘图标唯一标识符会传到托盘事件的第一个参数里

		w.SetProperty("记录当前托盘图标", "1")
		btnMod.Enable(true).Redraw(false)
		btnShow.Enable(true).Redraw(false)
		btnMsg.Enable(true).Redraw(false)
		btnFocus.Enable(true).Redraw(false)
		return 0
	})

	// 设置弹出气泡消息
	btnMsg.Event_BnClick(func(pbHandled *bool) int {
		btnMsg.Enable(false).Redraw(false)
		defer btnMsg.Enable(true).Redraw(false)

		xc.XTrayIcon_SetPopupBalloon("弹出气泡标题", "弹出气泡内容: "+strconv.Itoa(rand.Int()), 0, xcc.TrayIcon_Flag_Icon_Info)
		xc.XTrayIcon_Modify()
		return 0
	})

	// 修改
	btnMod.Event_BnClick(func(pbHandled *bool) int {
		btnMod.Enable(false).Redraw(false)
		defer btnMod.Enable(true).Redraw(false)

		// 修改为新的图标
		if w.GetProperty("记录当前托盘图标") == "1" {
			xc.XTrayIcon_SetIcon(hIcon2)
			w.SetProperty("记录当前托盘图标", "2")
		} else {
			xc.XTrayIcon_SetIcon(hIcon1)
			w.SetProperty("记录当前托盘图标", "1")
		}

		// 修改托盘提示信息
		rand.Seed(time.Now().Unix())
		xc.XTrayIcon_SetTips("修改了图标和托盘提示信息: " + strconv.Itoa(rand.Int()))

		// 应用修改
		xc.XTrayIcon_Modify()
		return 0
	})

	// 显示或隐藏
	btnShow.Event_BnClick(func(pbHandled *bool) int {
		btnShow.Enable(false).Redraw(false)

		if btnShow.GetText() == "显示" {
			btnShow.SetText("隐藏")
			xc.XTrayIcon_Add(w.Handle, 111)
		} else {
			btnShow.SetText("显示")
			xc.XTrayIcon_Del()
		}

		btnReset.Enable(!btnReset.IsEnable()).Redraw(false)
		btnShow.Enable(true).Redraw(false)
		return 0
	})

	// 重置会清空已经设置的图标(图标句柄并没有释放)、文字提示等信息, 只能在托盘图标不在系统托盘显示的时候使用
	btnReset.Event_BnClick(func(pbHandled *bool) int {
		btnReset.Enable(false).Redraw(false)
		defer btnReset.Enable(true).Redraw(false)

		xc.XTrayIcon_Reset()
		return 0
	})

	// 托盘图标置焦点
	btnFocus.Event_BnClick(func(pbHandled *bool) int {
		btnFocus.Enable(false).Redraw(false)
		defer btnFocus.Enable(true).Redraw(false)

		xc.XTrayIcon_SetFocus()
		return 0
	})

	// 注册托盘图标事件
	w.Event_TRAYICON(func(wParam, lParam uintptr, pbHandled *bool) int {
		fmt.Println(wParam, lParam)
		if wParam != 111 { // 不是自定义的托盘图标唯一标识符.
			return 0
		}

		switch xcc.WM_(lParam) {
		case xcc.WM_LBUTTONDOWN:
			w.ShowWindow(xcc.SW_SHOWNORMAL)
		case xcc.WM_RBUTTONDOWN:
			// 创建菜单
			menu := widget.NewMenu()
			// 一级菜单
			menu.AddItem(10001, "窗口置顶", 0, xcc.Menu_Item_Flag_Select)
			// 获取自己 SetProperty 的值, 这不是读写元素的属性, 可理解为对元素里内置的一个map进行读写
			// 这样可以不用另外声明变量, 能用到很多地方记录一些东西
			if a.GetProperty(w.Handle, "记录窗口置顶状态") == "1" {
				menu.SetItemCheck(10001, true)
			} else {
				menu.SetItemCheck(10001, false)
			}

			menu.AddItem(99999, "退出", 0, xcc.Menu_Item_Flag_Normal)

			// 获取鼠标光标的屏幕坐标
			var pt wapi.POINT
			wapi.GetCursorPos(&pt)
			// 弹出菜单
			menu.Popup(w.GetHWND(), pt.X+10, pt.Y-30, 0, xcc.Menu_Popup_Position_Left_Top)
		}
		return 0
	})

	// 菜单被选择事件
	w.Event_MENU_SELECT(func(nID int32, pbHandled *bool) int {
		fmt.Println("托盘菜单被选择:", nID)
		switch nID {
		case 10001:
			if a.GetProperty(w.Handle, "记录窗口置顶状态") == "1" {
				a.SetProperty(w.Handle, "记录窗口置顶状态", "0")
				wnd.SetTop(w.GetHWND(), false)
				fmt.Println("窗口已取消置顶")
			} else {
				a.SetProperty(w.Handle, "记录窗口置顶状态", "1")
				wnd.SetTop(w.GetHWND(), true)
				fmt.Println("窗口已被置顶")
			}
		case 99999:
			w.CloseWindow()
			a.PostQuitMessage(0)
		}
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	a.Run()
}
