// 托盘图标.
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wnd"
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
	w := window.New(0, 0, 430, 300, "TrayIcon", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	w.SetBorderSize(1, 30, 1, 1)

	// 加载icon
	hIcon := wapi.LoadImageW(0, "TrayIcon/icon.ico", wapi.IMAGE_ICON, 0, 0, wapi.LR_LOADFROMFILE|wapi.LR_DEFAULTSIZE|wapi.LR_SHARED)
	fmt.Println("hIcon:", hIcon)
	fmt.Println("LastErr:", syscall.GetLastError())

	// 托盘图标_置图标
	xc.XTrayIcon_SetIcon(hIcon)
	// 托盘图标_置提示文本
	xc.XTrayIcon_SetTips("托盘提示信息")
	// 置弹出气泡
	// xc.XTrayIcon_SetPopupBalloon("弹出气泡", "弹出气泡内容测试", 0, xcc.TrayIcon_Flag_Icon_Info)

	// 添加
	btn1 := widget.NewButton(50, 135, 80, 30, "添加", w.Handle)
	btn1.Event_BnClick(func(pbHandled *bool) int {
		xc.XTrayIcon_Add(w.Handle, 111) // 自定义的id会传到托盘事件里
		return 0
	})

	// 修改
	btn2 := widget.NewButton(150, 135, 80, 30, "修改", w.Handle)
	btn2.Event_BnClick(func(pbHandled *bool) int {
		rand.Seed(time.Now().Unix())
		xc.XTrayIcon_SetTips("修改了托盘提示信息: " + strconv.Itoa(rand.Int()))
		xc.XTrayIcon_Modify()
		return 0
	})

	// 删除
	btn3 := widget.NewButton(250, 135, 80, 30, "删除", w.Handle)
	btn3.Event_BnClick(func(pbHandled *bool) int {
		xc.XTrayIcon_Del()
		return 0
	})

	// 注册托盘图标事件
	w.Event_TRAYICON(func(wParam, lParam uint, pbHandled *bool) int {
		switch xcc.WM_(lParam) {
		case xcc.WM_LBUTTONDOWN:
			w.ShowWindow(xcc.SW_SHOWNORMAL)
		case xcc.WM_RBUTTONDOWN:
			// 创建菜单
			menu := widget.NewMenu()
			// 一级菜单
			menu.AddItem(10001, "窗口置顶", 0, xcc.Menu_Item_Flag_Select)
			// 获取自己 SetProperty 的值, 这不是读写元素的属性, 只是对元素里内置的一个map进行读写
			// 这样可以不用另外声明变量, 能用到很多地方记录一些东西
			if a.GetProperty(w.Handle, "窗口置顶") == "1" {
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
			if a.GetProperty(w.Handle, "窗口置顶") == "1" {
				a.SetProperty(w.Handle, "窗口置顶", "0")
				wnd.SetTop(w.GetHWND(), false)
				fmt.Println("窗口已取消置顶")
			} else {
				a.SetProperty(w.Handle, "窗口置顶", "1")
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
