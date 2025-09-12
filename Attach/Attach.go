// 附加窗口, 可以附加到其他窗口上
package main

import (
	"syscall"
	"unsafe"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 用 wapi 创建了一个原生窗口
	hwnd := createWindow()
	// 附加到原生窗口, 返回炫彩窗口对象
	w := window.Attach(hwnd, xcc.Window_Style_Default)

	w.ShowWindow(xcc.SW_RESTORE)
	a.Run()
	a.Exit()
}

func createWindow() uintptr {
	hInstance := wapi.GetModuleHandleEx(0, "")
	icow := wapi.GetSystemMetrics(wapi.SM_CXICON)
	icoh := wapi.GetSystemMetrics(wapi.SM_CYICON)
	icon := wapi.LoadImageW(hInstance, uintptr(32512), wapi.IMAGE_ICON, icow, icoh, wapi.LR_DEFAULTCOLOR)

	className := "windowclass"
	wc := wapi.WNDCLASSEX{
		Style:         wapi.CS_HREDRAW | wapi.CS_VREDRAW | wapi.CS_PARENTDC | wapi.CS_DBLCLKS,
		CbSize:        uint32(unsafe.Sizeof(wapi.WNDCLASSEX{})),
		HInstance:     hInstance,
		LpszClassName: common.StrPtr(className),
		HIcon:         icon,
		HIconSm:       icon,
		LpfnWndProc:   syscall.NewCallback(wndproc),
	}
	wapi.RegisterClassEx(&wc)

	width := int32(800)
	height := int32(600)
	// 居中
	x := (wapi.GetSystemMetrics(wapi.SM_CXSCREEN) - width) / 2
	y := (wapi.GetSystemMetrics(wapi.SM_CYSCREEN) - height) / 2

	hwnd := wapi.CreateWindowEx(0, className, "原生窗口", xcc.WS_MINIMIZE, x, y, width, height, 0, 0, hInstance, 0)

	wapi.ShowWindow(hwnd, xcc.SW_HIDE)
	wapi.UpdateWindow(hwnd)
	wapi.SetFocus(hwnd)
	return hwnd
}

func wndproc(hwnd uintptr, msg uint32, wp, lp uintptr) uintptr {
	return wapi.DefWindowProc(hwnd, msg, wp, lp)
}
