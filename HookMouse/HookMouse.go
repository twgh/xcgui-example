// 创建全局鼠标钩子, 监听鼠标消息, 可当热键使用或其他用途.
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.New(true)
	// 启用自适应DPI
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "全局鼠标钩子", 0, xcc.Window_Style_Default)

	widget.NewShapeText(40, 40, 300, 30, "在任何窗口操作鼠标都能够监听到", w.Handle)
	checkBtn := widget.NewButton(40, 80, 300, 30, "拦截鼠标右键按下消息", w.Handle).SetTypeEx(xcc.Button_Type_Check)
	checkBtn.EnableBkTransparent(true)

	// 注册事件_窗口鼠标右键按下, 用来检测是否真的拦截了鼠标右键按下消息
	w.Event_RBUTTONDOWN(func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
		xc.XC_Alert("提示", fmt.Sprintf("响应了炫彩窗口鼠标右键被按下消息, 证明没有被拦截, nFlags: %d, pPt: %v", nFlags, pPt))
		return 0
	})

	msHook := wutil.NewHookMouse(func(nCode int32, wParam xcc.WM_, lParam *wapi.MSLLHOOKSTRUCT) uintptr {
		if nCode < 0 { // nCode小于0时不应继续处理
			return wutil.CallNextHookEx_Mouse(nCode, wParam, lParam)
		}

		switch wParam {
		case xcc.WM_LBUTTONDOWN: // 鼠标左键按下
			fmt.Println("鼠标左键按下, 坐标:", lParam.PT)
		case xcc.WM_RBUTTONDOWN: // 鼠标右键按下
			if checkBtn.GetStateEx() == xcc.Button_State_Check {
				fmt.Println("拦截了鼠标右键按下, 是不会真实响应鼠标右键消息的, 你在任务栏上右键已经没用了, 有些程序窗口拦截不了自行研究, 坐标:", lParam.PT)
				return 1 // 返回1可拦截, 这时按下鼠标右键是不会有响应的, 部分软件窗口拦截不了有多方面原因比如该程序做了特殊处理, 自行研究
			}
		case xcc.WM_MBUTTONDOWN: // 鼠标中键按下
			fmt.Println("鼠标中键按下, 坐标:", lParam.PT)
		case xcc.WM_XBUTTONDOWN: // 鼠标侧键按下
			value := wutil.GetHigh16Bits(lParam.MouseData)
			if value == 1 {
				fmt.Println("鼠标侧键1按下, 坐标:", lParam.PT)
			} else if value == 2 {
				fmt.Println("鼠标侧键2按下, 坐标:", lParam.PT)
			}
		case xcc.WM_MOUSEWHEEL: // 鼠标滚轮滚动
			value := wutil.GetHigh16Bits(lParam.MouseData)
			if lParam.MouseData > 0 {
				fmt.Printf("鼠标滚轮向上滚动, 滚轮增量: %d, 坐标:%v, lParam.MouseData: %v\n", value, lParam.PT, lParam.MouseData)
			} else if lParam.MouseData < 0 {
				fmt.Printf("鼠标滚轮向下滚动, 滚轮增量: %d, 坐标:%v, lParam.MouseData: %v\n", value, lParam.PT, lParam.MouseData)
			}
		}
		return wutil.CallNextHookEx_Mouse(nCode, wParam, lParam)
	})

	w.Event_CLOSE(func(pbHandled *bool) int {
		msHook.Unhook()
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
