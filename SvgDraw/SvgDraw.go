// svg绘制
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/svg"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	w    *window.Window
	svg1 *svg.Svg
)

func main() {
	a := app.New(true)
	w = window.NewWindow(0, 0, 350, 200, "", 0, xcc.Window_Style_Default)

	// SVG_加载从字符串
	svg1 = svg.NewSvg_LoadStringW(svgStr)
	if svg1.Handle == 0 {
		panic("svg1.Handle = 0")
	}

	// 窗口绘制消息
	w.Event_PAINT(OnWndDrawWindow)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

func OnWndDrawWindow(hDraw int, pbHandled *bool) int {
	*pbHandled = true
	// 在自绘事件函数中,用户手动调用绘制窗口, 以便控制绘制顺序
	w.DrawWindow(hDraw)
	// 创建绘制对象
	draw := drawx.NewDrawByHandle(hDraw)

	left := 20
	top := 50
	draw.DrawSvgEx(svg1.Handle, left, top, 100, 100)
	left += 100
	draw.DrawSvgEx(svg1.Handle, left, top+(100-72)/2, 72, 72)
	left += 72
	draw.DrawSvgEx(svg1.Handle, left, top+(100-48)/2, 48, 48)
	left += 48
	draw.DrawSvgEx(svg1.Handle, left, top+(100-32)/2, 32, 32)
	left += 32
	draw.DrawSvgEx(svg1.Handle, left, top+(100-24)/2, 24, 24)
	left += 24
	return 0
}

var svgStr = `<svg t="1637645455197" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="31134" width="22" height="22"><path d="M934.8096 278.2208v-35.8912c0-48.0768-39.0144-87.0912-87.0912-87.0912H160.3072c-48.0768 0-87.0912 39.0144-87.0912 87.0912v35.8912h861.5936z" fill="#C65EDB" p-id="31135"></path><path d="M725.9136 536.6272h-48.0256v39.8336h48.0256c11.008 0 19.9168-8.96 19.9168-19.9168 0-11.008-8.96-19.9168-19.9168-19.9168z" fill="#BD50D3" p-id="31136"></path><path d="M73.216 343.6032v440.9856c0 48.0768 39.0144 87.0912 87.0912 87.0912h687.4624c48.0768 0 87.0912-39.0144 87.0912-87.0912V343.6032H73.216z m347.8528 181.76l-79.6672 176.8448a32.54784 32.54784 0 0 1-29.6448 19.1488h-0.2048a32.58368 32.58368 0 0 1-29.5936-19.5584l-76.9536-176.7936a29.01504 29.01504 0 0 1 15.0016-38.1952c14.6944-6.4 31.7952 0.3072 38.1952 15.0528l53.9648 123.9552L368.128 501.504a29.04064 29.04064 0 0 1 38.4-14.5408 29.10208 29.10208 0 0 1 14.5408 38.4z m115.0464 176.9984c0 16.0256-13.0048 29.0304-29.0304 29.0304s-29.0304-13.0048-29.0304-29.0304V507.5968c0-16.0256 13.0048-29.0304 29.0304-29.0304s29.0304 13.0048 29.0304 29.0304v194.7648z m189.7984-67.84h-48.0256v67.84c0 16.0256-13.0048 29.0304-29.0304 29.0304s-29.0304-13.0048-29.0304-29.0304V507.5968c0-16.0256 13.0048-29.0304 29.0304-29.0304h77.056c43.008 0 77.9776 34.9696 77.9776 77.9776s-35.0208 77.9776-77.9776 77.9776z" fill="#BD50D3" p-id="31137"></path><path d="M934.8096 343.6032H73.216v440.9856c0 48.0768 39.0144 87.0912 87.0912 87.0912h393.9328c182.3232-72.2432 323.84-225.1264 380.6208-414.6688V343.6032z m-513.7408 181.76l-79.6672 176.8448a32.54784 32.54784 0 0 1-29.6448 19.1488h-0.2048a32.58368 32.58368 0 0 1-29.5936-19.5584l-76.9536-176.7936a29.01504 29.01504 0 1 1 53.1968-23.1936l53.9648 123.9552L368.128 501.504a29.04064 29.04064 0 0 1 38.4-14.5408 29.10208 29.10208 0 0 1 14.5408 38.4z m115.0464 176.9984c0 16.0256-13.0048 29.0304-29.0304 29.0304s-29.0304-13.0048-29.0304-29.0304V507.5968c0-16.0256 13.0048-29.0304 29.0304-29.0304s29.0304 13.0048 29.0304 29.0304v194.7648z m189.7984-67.84h-48.0256v67.84c0 16.0256-13.0048 29.0304-29.0304 29.0304s-29.0304-13.0048-29.0304-29.0304V507.5968c0-16.0256 13.0048-29.0304 29.0304-29.0304h77.056c43.008 0 77.9776 34.9696 77.9776 77.9776s-35.0208 77.9776-77.9776 77.9776z" fill="#C65EDB" p-id="31138"></path><path d="M73.216 684.2368c21.9648 2.2528 44.288 3.4304 66.8672 3.4304 44.4928 0 87.9616-4.5056 129.9456-13.1072l-65.1264-149.5552a29.01504 29.01504 0 1 1 53.1968-23.1936l53.9648 123.9552L368.128 501.504a29.04064 29.04064 0 0 1 38.4-14.5408 28.99968 28.99968 0 0 1 14.5408 38.4l-54.7328 121.4976a645.3248 645.3248 0 0 0 111.7184-54.6816V507.5968c0-16.0256 13.0048-29.0304 29.0304-29.0304s29.0304 13.0048 29.0304 29.0304v44.288a648.33536 648.33536 0 0 0 174.7456-208.2816H73.216v340.6336zM775.5264 155.1872H160.3072c-48.0768 0-87.0912 39.0144-87.0912 87.0912v35.8912h667.7504c15.36-39.2192 27.0848-80.384 34.56-122.9824z" fill="#CA6EE0" p-id="31139"></path><path d="M73.216 452.2496c100.5056-14.8992 193.4848-52.992 273.3056-108.6464H73.216v108.6464zM525.5168 155.1872H160.3072c-48.0768 0-87.0912 39.0144-87.0912 87.0912v35.8912h352.768a648.0384 648.0384 0 0 0 99.5328-122.9824z" fill="#D786EA" p-id="31140"></path></svg>`
