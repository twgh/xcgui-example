// 托盘图标. 使用对象来操作.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"syscall"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wnd"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

const prePath = "Advanced/TrayIcon/"

func main() {
	rand.Seed(time.Now().Unix())

	// 1.初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "TrayIcon", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
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

	// 加载图标, 想把 icon 文件内置在程序里可以看 SetWindowIcon 的例子
	hIcon1 := wutil.HIcon(prePath + "icon1.ico")
	fmt.Println("hIcon1:", hIcon1, "LastErr:", syscall.GetLastError())
	hIcon2 := wutil.HIcon(prePath + "icon2.ico")
	fmt.Println("hIcon2:", hIcon2, "LastErr:", syscall.GetLastError())

	var tray *window.TrayIcon
	// 创建并显示托盘图标
	btnAdd.Event_BnClick(func(pbHandled *bool) int {
		btnAdd.Enable(false).Redraw(false)

		// 创建托盘图标对象
		tray = w.CreateTrayIcon(hIcon1, "托盘提示信息")
		// 显示托盘图标
		tray.Show(true)

		btnMod.Enable(true).Redraw(false)
		btnShow.Enable(true).Redraw(false)
		btnMsg.Enable(true).Redraw(false)
		btnFocus.Enable(true).Redraw(false)
		return 0
	})

	// 修改图标和提示信息
	btnMod.Event_BnClick(func(pbHandled *bool) int {
		btnMod.Enable(false).Redraw(false)
		defer btnMod.Enable(true).Redraw(false)

		// 修改为新的图标
		if tray.HIcon == hIcon1 {
			tray.SetIcon(hIcon2)
		} else {
			tray.SetIcon(hIcon1)
		}

		// 修改托盘提示信息
		tray.SetTips("修改了图标和托盘提示信息: " + strconv.Itoa(rand.Int()))

		// 应用修改
		tray.Modify()
		return 0
	})

	// 设置弹出气泡消息
	btnMsg.Event_BnClick(func(pbHandled *bool) int {
		tray.SetPopupBalloon("弹出气泡标题", "弹出气泡内容: "+strconv.Itoa(rand.Int()), 0, xcc.TrayIcon_Flag_Icon_Info)
		// 应用修改
		tray.Modify()
		return 0
	})

	// 显示或隐藏, 原理实际是添加或删除.
	btnShow.Event_BnClick(func(pbHandled *bool) int {
		btnShow.Enable(false).Redraw(false)

		if btnShow.GetText() == "显示" {
			btnShow.SetText("隐藏")
			tray.Show(true)
		} else {
			btnShow.SetText("显示")
			tray.Show(false)
		}

		/*
			// 当你手动把托盘图标从隐藏区域拖拽到任务栏固定后, 你肯定想让他一直固定在任务栏, 不再回隐藏区域.
			// 但这里有个坑: 如果你在托盘图标创建并显示后, 有过修改图标/提示信息的操作, 那你隐藏图标再显示时它会回到隐藏区域.
			// 如果你在托盘图标创建后, 没有过修改图标/提示信息的操作, 那你随便控制显示和隐藏都没问题.
			// 如果你不需要切换隐藏和显示状态, 那你随便修改都无所谓.
			// 如果你想修改, 还想控制显示和隐藏, 那么你应该在隐藏后调用一下重置函数, 显示后再重新对托盘图标进行设置.
			//
			// 因为我不知道xcgui.dll里这几个函数的内部实现究竟是什么, 所以这只是我测试得出的解决办法, 是否有更好的我不知道.
			// 不过控制隐藏和显示本来就是个小众需求, 用到的人应该很少. 大部分人的托盘图标都只用来呼出托盘菜单罢了.

			if btnShow.GetText() == "显示" {
				btnShow.SetText("隐藏")
				tray.Show(true)

				// 因为隐藏的时候重置了, 这里就得重新设置, 细节是show之后再设置托盘图标信息, 而不是show之前.
				tray.SetIcon(tray.HIcon).SetTips(tray.Tips).Modify()
			} else {
				btnShow.SetText("显示")
				tray.Show(false)

				// 隐藏后多调用了个重置
				tray.Reset()
			}*/

		btnReset.Enable(!btnReset.IsEnable()).Redraw(false)
		btnShow.Enable(true).Redraw(false)
		return 0
	})

	// 重置会清空已经设置的图标(图标句柄并没有释放)、文字提示等信息, 只能在托盘图标不在系统托盘显示的时候使用
	btnReset.Event_BnClick(func(pbHandled *bool) int {
		btnReset.Enable(false).Redraw(false)
		defer btnReset.Enable(true).Redraw(false)

		tray.Reset()
		return 0
	})

	// 托盘图标置焦点
	btnFocus.Event_BnClick(func(pbHandled *bool) int {
		btnFocus.Enable(false).Redraw(false)
		defer btnFocus.Enable(true).Redraw(false)

		tray.SetFocus()
		return 0
	})

	// 注册托盘图标事件
	w.Event_TRAYICON(func(wParam, lParam uintptr, pbHandled *bool) int {
		fmt.Println(wParam, lParam)
		if int32(wParam) != tray.Id { // 不是自定义的托盘图标唯一标识符.
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
			if w.GetProperty("记录窗口置顶状态") == "1" {
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
			if w.GetProperty("记录窗口置顶状态") == "1" {
				w.SetProperty("记录窗口置顶状态", "0")
				wnd.SetTop(w.GetHWND(), false)
				fmt.Println("窗口已取消置顶")
			} else {
				w.SetProperty("记录窗口置顶状态", "1")
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
	a.Exit()
}
