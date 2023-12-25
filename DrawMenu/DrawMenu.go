// 自己美化菜单.
// 这属于是自绘, 如果直接用自带的菜单会简单许多.
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	a   *app.App
	w   *window.Window
	btn *widget.Button

	item_selected = true                 // 控制item_select是否选中
	svgMap        = make(map[int]int, 7) // 存放图片句柄
)

// 菜单项ID

const (
	menuid_item1 = iota + 10000
	menuid_subitem1
	menuid_subitem2

	menuid_item2
	menuid_item_select
	menuid_item3
	menuid_item4
	menuid_item_disable
)

func main() {
	// 1.初始化UI库
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w = window.New(0, 0, 500, 350, "DrawMenu", 0, xcc.Window_Style_Default)
	w.SetBorderSize(1, 30, 1, 1)

	// 加载所有的svg图片
	loadAllSvg()

	// 创建一个按钮
	btn = widget.NewButton(50, 50, 100, 30, "ShowMenu", w.Handle)
	// 注册按钮被单击事件
	btn.Event_BnClick(onBnClick)

	// 注册菜单背景绘制事件
	w.Event_MENU_DRAW_BACKGROUND(onMenuDrawBackground)
	// 注册菜单项绘制事件
	w.Event_MENU_DRAWITEM(onMenuDrawItem)
	// 注册菜单被选择事件
	w.Event_MENU_SELECT(onMenuSelect)

	// 3.显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	a.Exit()
}

// 加载所有的svg图片
func loadAllSvg() {
	svgMap[menuid_item1] = loadSvg(svg1)
	svgMap[menuid_item2] = loadSvg(svg2)
	svgMap[menuid_item3] = loadSvg(svg3)
	svgMap[menuid_item4] = loadSvg(svg4)
	svgMap[menuid_item_select] = loadSvg(svg_select)
	svgMap[menuid_subitem1] = loadSvg(svg_subitem1)
	svgMap[menuid_subitem2] = loadSvg(svg_subitem2)
	svgMap[menuid_item_disable] = loadSvg(svg_disable)
}

// 加载svg图片, 并禁止自动销毁
func loadSvg(svgText string) int {
	hSvg := xc.XImage_LoadSvgStringW(svgText)
	xc.XImage_EnableAutoDestroy(hSvg, false)
	return hSvg
}

// 按钮被单击事件
func onBnClick(pbHandled *bool) int {
	// 创建菜单
	menu := widget.NewMenu()
	menu.SetItemHeight(30)          // 设置菜单项高度
	menu.EnableDrawBackground(true) // 菜单_启用用户绘制背景
	menu.EnableDrawItem(true)       // 菜单_启用用户绘制项

	// 一级菜单
	menu.AddItemIcon(menuid_item1, "item1测试", 0, svgMap[menuid_item1], xcc.Menu_Item_Flag_Normal)
	menu.SetItemWidth(menuid_item1, 100) // 设置菜单宽度
	menu.AddItemIcon(menuid_item2, "item2中文", 0, svgMap[menuid_item2], xcc.Menu_Item_Flag_Normal)
	if item_selected {
		menu.AddItemIcon(menuid_item_select, "item_select", 0, svgMap[menuid_item_select], xcc.Menu_Item_Flag_Check)
	} else {
		menu.AddItem(menuid_item_select, "item_select", 0, xcc.Menu_Item_Flag_Normal)
	}
	menu.AddItemIcon(menuid_item3, "item3", 0, svgMap[menuid_item3], xcc.Menu_Item_Flag_Normal)
	menu.AddItem(-1, "", 0, xcc.Menu_Item_Flag_Separator) // 分隔栏
	menu.AddItemIcon(menuid_item4, "item4", 0, svgMap[menuid_item4], xcc.Menu_Item_Flag_Normal)
	menu.AddItemIcon(menuid_item_disable, "item_disable", 0, svgMap[menuid_item_disable], xcc.Menu_Item_Flag_Disable)

	// item1的二级菜单
	menu.AddItemIcon(menuid_subitem1, "subitem1", menuid_item1, svgMap[menuid_subitem1], xcc.Menu_Item_Flag_Normal)
	menu.AddItemIcon(menuid_subitem2, "subitem2", menuid_item1, svgMap[menuid_subitem2], xcc.Menu_Item_Flag_Normal)

	// 获取按钮坐标
	var rc xc.RECT
	btn.GetWndClientRectDPI(&rc)
	// 转换到屏幕坐标
	pt := wapi.POINT{X: rc.Left, Y: rc.Bottom}
	wapi.ClientToScreen(w.GetHWND(), &pt)
	// 弹出菜单
	menu.Popup(w.GetHWND(), pt.X, pt.Y, 0, xcc.Menu_Popup_Position_Left_Top)
	return 0
}

// 菜单背景绘制事件
func onMenuDrawBackground(hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int {
	*pbHandled = true // 作用是拦截菜单原本的绘制
	var rc xc.RECT
	xc.XWnd_GetClientRect(pInfo.HWindow, &rc)

	// 绘制菜单背景
	xc.XDraw_SetBrushColor(hDraw, xc.ABGR(255, 255, 255, 255))
	xc.XDraw_FillRect(hDraw, &rc)

	// 绘制菜单边框
	xc.XDraw_SetBrushColor(hDraw, xc.ABGR(218, 220, 224, 255))
	xc.XDraw_DrawRect(hDraw, &rc)
	return 0
}

// 菜单项绘制事件
func onMenuDrawItem(hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int {
	*pbHandled = true // 作用是拦截菜单原本的绘制

	// 绘制分割栏
	if pInfo.NState&xcc.Menu_Item_Flag_Separator > 0 {
		xc.XDraw_SetBrushColor(hDraw, xc.ABGR(218, 220, 224, 255))
		xc.XDraw_DrawLine(hDraw, int(pInfo.RcItem.Left+3), int(pInfo.RcItem.Top+1), int(pInfo.RcItem.Right-3), int(pInfo.RcItem.Top+1))
		return 0
	}

	// 绘制鼠标停留时菜单项的背景
	if pInfo.NState&xcc.Menu_Item_Flag_Select > 0 {
		// 左右把项填满
		rc := xc.RECT{
			Left:   pInfo.RcItem.Left - 2,
			Top:    pInfo.RcItem.Top,
			Right:  pInfo.RcItem.Right + 2,
			Bottom: pInfo.RcItem.Bottom,
		}
		xc.XDraw_SetBrushColor(hDraw, xc.ABGR(230, 230, 230, 255))
		xc.XDraw_FillRect(hDraw, &rc)
	} else {
		// 如果存在下一个兄弟项, 绘制菜单项之间的分割线
		/* 		if xc.XMenu_GetNextSiblingItem(pInfo.HMenu, int(pInfo.NID)) != xcc.XC_ID_ERROR {
			xc.XDraw_SetBrushColor(hDraw, xc.ABGR(82, 88, 94, 255))
			xc.XDraw_DrawLine(hDraw, int(pInfo.RcItem.Left)+3, int(pInfo.RcItem.Bottom)-1, int(pInfo.RcItem.Right)-3, int(pInfo.RcItem.Bottom)-1)
		} */
	}

	// 绘制右三角
	if pInfo.NState&xcc.Menu_Item_Flag_Popup > 0 {
		var pt [3]xc.POINT
		pt[0].X = pInfo.RcItem.Right - 12
		pt[0].Y = pInfo.RcItem.Top + 10

		pt[1].X = pInfo.RcItem.Right - 12
		pt[1].Y = pInfo.RcItem.Top + 20

		pt[2].X = pInfo.RcItem.Right - 7
		pt[2].Y = pInfo.RcItem.Top + 15
		xc.XDraw_SetBrushColor(hDraw, xc.ABGR(130, 130, 130, 255))
		xc.XDraw_FillPolygon(hDraw, pt[:], 3)
	}

	// 取菜单左侧区域宽度
	leftWidth := xc.XMenu_GetLeftWidth(pInfo.HMenu)
	rc := pInfo.RcItem
	rc.Left = leftWidth + 5

	if pInfo.NState&xcc.Menu_Item_Flag_Disable > 0 {
		// 设置被禁用的菜单项文本颜色
		xc.XDraw_SetBrushColor(hDraw, xc.ABGR(160, 160, 160, 255))
	} else {
		// 设置未禁用的菜单项文本颜色
		xc.XDraw_SetBrushColor(hDraw, xc.ABGR(77, 77, 77, 255))
	}
	// 获取菜单项文本
	text := common.UintPtrToString(pInfo.PText)
	// 绘制菜单项文本
	xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
	xc.XDraw_DrawText(hDraw, text, &rc)

	// 绘制菜单项图标
	if pInfo.HIcon > 0 {
		iconWidth := xc.XImage_GetHeight(pInfo.HIcon)
		iconHeight := xc.XImage_GetWidth(pInfo.HIcon)
		height := pInfo.RcItem.Bottom - pInfo.RcItem.Top
		if height >= 2 && iconWidth >= 2 && iconHeight >= 2 {
			top := (height - iconHeight) / 2
			left := (leftWidth - iconWidth) / 2
			if top < 0 {
				top = 0
			}
			if left < 0 {
				left = 0
			}
			xc.XDraw_Image(hDraw, pInfo.HIcon, left, pInfo.RcItem.Top+top)
		}
	}
	return 0
}

// 菜单被选择事件
func onMenuSelect(nID int32, pbHandled *bool) int {
	if nID == menuid_item_select {
		item_selected = !item_selected
	}
	return 0
}

const (
	svg1 = `<svg t="1644889730063" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5072" width="20" height="20"><path d="M808.192 413.747c-14.797 0-24.576 9.83-24.576 24.576v343.808c0 29.491-19.507 49.101-49.1 49.101H243.353c-29.594 0-49.101-19.61-49.101-49.1V241.868c0-29.491 19.507-49.101 49.1-49.101h368.385c14.796 0 24.576-9.78 24.576-24.576 0-14.746-9.728-24.576-24.576-24.576h-392.96c-44.39 0-73.677 39.066-73.677 83.302v579.738c0 44.237 29.286 73.677 73.677 73.677H759.04c44.39 0 73.677-29.44 73.677-73.677V438.323c0-14.745-9.728-24.576-24.525-24.576z m-430.131 233.78c11.878 11.775 31.13 11.775 43.059 0l448.87-449.23a30.208 30.208 0 0 0 0-42.854 30.566 30.566 0 0 0-43.059 0l-448.87 449.229a30.106 30.106 0 0 0 0 42.854z" p-id="5073" fill="#4d4d4d"></path></svg>`

	svg2 = `<svg t="1644825215895" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="17579" width="26" height="26"><path d="M752.535797 273.701662l-2.230808-2.227738c-51.299363-51.300386-135.227868-51.300386-186.526207 0l-118.846782 118.883621c-51.299363 51.281967-51.299363 135.222751 0 186.544627l2.192945 2.156106c4.27742 4.267187 8.833179 8.116865 13.485129 11.728112l43.49563-43.544749c-5.086855-2.956332-9.885138-6.591115-14.223956-10.92891l-2.192945-2.156106c-27.855418-27.866674-27.855418-73.180719 0-101.048417L606.609263 314.26859c27.782763-27.855418 73.084529-27.855418 100.951203 0l2.218528 2.180666c27.855418 27.867698 27.855418 73.194022 0 101.012601l-53.736878 53.765531c9.304923 23.067368 13.740956 47.64002 13.304004 72.114434l83.152838-83.117023C803.83516 408.918273 803.83516 325.004095 752.535797 273.701662L752.535797 273.701662 752.535797 273.701662 752.535797 273.701662zM576.877101 444.959118c-4.266164-4.264117-8.820899-8.118911-13.521968-11.680017l-43.472094 43.496653c5.088902 3.00545 9.888208 6.615675 14.249539 10.952446l2.215458 2.227738c27.855418 27.820626 27.855418 73.135694 0 101.002368L417.465438 709.790762c-27.854395 27.821649-73.15616 27.821649-101.010555 0l-2.229784-2.204202c-27.854395-27.864628-27.854395-73.204256 0-100.999298l53.771671-53.745065c-9.317203-23.070438-13.763468-47.665603-13.340843-72.140017l-83.176374 83.068927c-51.312666 51.349505-51.312666 135.288243 0 186.563046l2.216481 2.228761c51.299363 51.251268 135.227868 51.251268 186.526207 0l118.835526-118.883621c51.250244-51.299363 51.250244-135.263683 0-186.513928L576.877101 444.959118 576.877101 444.959118 576.877101 444.959118 576.877101 444.959118zM576.877101 444.959118" p-id="17580" fill="#4d4d4d"></path></svg>`

	svg3 = `<svg t="1644825317661" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="19032" width="18" height="18"><path d="M485.376 931.84c-231.424 0-419.84-188.416-419.84-419.84s188.416-419.84 419.84-419.84c197.632 0 370.688 140.288 410.624 333.824 2.048 11.264 4.096 23.552 6.144 34.816 2.048 16.384-10.24 31.744-26.624 33.792-16.384 2.048-31.744-10.24-33.792-26.624-1.024-10.24-3.072-20.48-5.12-29.696-34.816-165.888-182.272-286.72-352.256-286.72-197.632 1.024-358.4 161.792-358.4 360.448S286.72 871.424 485.376 871.424c177.152 0 329.728-132.096 355.328-306.176 2.048-16.384 17.408-27.648 33.792-25.6 16.384 2.048 27.648 17.408 25.6 33.792C870.4 778.24 691.2 931.84 485.376 931.84z" p-id="19033" fill="#4d4d4d"></path><path d="M95.232 758.784c-45.056 0-70.656-12.288-80.896-36.864-19.456-45.056 34.816-94.208 77.824-126.976 13.312-10.24 31.744-8.192 41.984 5.12 10.24 13.312 8.192 31.744-5.12 41.984-36.864 28.672-51.2 46.08-56.32 55.296 39.936 10.24 202.752-16.384 464.896-130.048 262.144-113.664 393.216-215.04 412.672-250.88-12.288-3.072-47.104-5.12-131.072 14.336-16.384 4.096-32.768-6.144-35.84-22.528-4.096-16.384 6.144-32.768 22.528-35.84 122.88-28.672 185.344-22.528 203.776 18.432 17.408 40.96-19.456 90.112-121.856 158.72-82.944 56.32-198.656 117.76-325.632 173.056-126.976 55.296-250.88 98.304-349.184 119.808-48.128 11.264-87.04 16.384-117.76 16.384z" p-id="19034" fill="#4d4d4d"></path></svg>`

	svg4 = `<svg t="1644820967538" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4979" width="20" height="20"><path d="M508.5 889.1c-51.1 0-100.7-10-147.4-29.8-45.1-19.1-85.6-46.4-120.3-81.1-34.8-34.8-62.1-75.3-81.1-120.3-19.7-46.7-29.8-96.3-29.8-147.4 0-54.2 11.2-106.5 33.3-155.6l36.5 16.5c-19.8 43.8-29.8 90.6-29.8 139.1 0 186.7 151.9 338.6 338.6 338.6 69.7 0 136.7-21 193.7-60.8l22.9 32.8c-63.7 44.4-138.6 68-216.6 68zM858.1 656.1l-36.9-15.4c17.2-41.3 25.9-85.1 25.9-130.2 0-186.7-151.9-338.6-338.6-338.6-69.3 0-135.9 20.8-192.6 60.1l-22.8-33c63.4-44 137.9-67.2 215.4-67.2 51.1 0 100.7 10 147.4 29.8 45.1 19.1 85.6 46.4 120.3 81.1C811 277.5 838.3 318 857.3 363c19.7 46.7 29.8 96.3 29.8 147.4 0 50.5-9.7 99.5-29 145.7z" p-id="4980" fill="#4d4d4d"></path><path d="M271 454.9L166.6 318.1 57.5 453.9zM965.5 554.6L861 691.4 752 555.5z" p-id="4981" fill="#4d4d4d"></path></svg>`

	svg_select = `<svg t="1644821064310" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9602" width="20" height="20"><path d="M480 736c-8 0-16-3.2-22.4-9.6l-288-288c-12.8-12.8-12.8-32 0-44.8 12.8-12.8 32-12.8 44.8 0L480 659.2l425.6-425.6c12.8-12.8 32-12.8 44.8 0 12.8 12.8 12.8 32 0 44.8l-448 448C496 732.8 488 736 480 736z" p-id="9603" fill="#4d4d4d"></path></svg>`

	svg_disable = `<svg t="1644892973051" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5040" width="20" height="20"><path d="M561.1 708.454V413.798c0-47.155 49.102-47.155 49.102 0v294.708c0.05 34.099-49.101 38.656-49.101-0.052z m-147.302 0V413.798c0-47.155 49.101-47.155 49.101 0v294.708c0 38.656-49.1 34.099-49.1-0.052z m442.01-442.01H708.454v-49.151c0-71.731-22.988-98.253-98.252-98.253H413.798c-74.035 0-98.252 24.166-98.252 98.253v49.152H168.192c-53.094 0-53.094 49.1 0 49.1h687.616c53.094 0 53.094-49.1 0-49.1z m-491.162-49.151c0-47.667 2.97-49.101 49.101-49.101h196.455c46.08 0 49.1 1.126 49.1 49.1v49.153H364.646v-49.152z m343.91 687.616h-393.01c-70.964 0-98.253-27.239-98.253-98.253V413.798c0-49.612 49.1-49.612 49.1 0v392.91c0 47.718-0.102 49.151 49.101 49.151h392.91c47.718 0 49.1 0.154 49.1-49.152V413.798c0-48.486 49.1-48.486 49.1 0v392.91c0.103 69.426-23.96 98.2-98.047 98.2z" p-id="5041" fill="#A0A0A0"></path></svg>`

	svg_subitem1 = `<svg t="1644821662815" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9848" width="20" height="20"><path d="M836.48 867.2l22.08-122.752c20.224-112.832 33.504-188.672 35.776-204.704 6.144-44.352-2.048-52.576-48.736-51.584-16.736 0.32-57.28 0.608-114.752 0.736-26.784 0.064-55.424 0.128-84.064 0.128h-38.208c-24.48 0-39.968-26.464-28.16-48.064 0.8-1.408 2.592-5.184 5.12-11.328 4.512-10.944 9.088-24.32 13.376-40 12.576-45.888 20.16-101.44 20.16-166.592 0-56.384-26.752-92.544-54.944-94.144-27.2-1.536-51.392 32-51.392 103.872 0 96.384-42.72 169.6-112.768 220.864a345.824 345.824 0 0 1-86.272 46.496v394.976h517.728l5.056-27.872v-0.032z m-558.144 92.64H100.8c-20.32 0-36.8-14.56-36.8-32.512V468.864c0-17.952 16.48-32.512 36.832-32.512h197.44l2.336-0.96a288.32 288.32 0 0 0 61.536-34.176c54.432-39.904 86.336-94.464 86.336-168.416 0-104.96 48.768-172.672 119.232-168.64 65.92 3.776 115.648 70.912 115.648 158.88 0 71.072-8.352 132.384-22.496 183.808-1.664 6.176-3.328 11.936-5.056 17.312 25.568 0 50.976-0.064 74.848-0.128v0.032c56.96-0.128 97.44-0.32 113.6-0.704 84.544-1.792 125.664 39.52 113.696 125.44-2.4 17.344-15.584 92.416-36.16 207.232a72131.418 72131.418 0 0 1-29.248 162.496l-2.016 10.976-0.512 2.88C887.04 948.8 873.696 960 858.208 960h-576.64a32.32 32.32 0 0 1-3.232-0.16z m-28.928-64.96V501.312H137.6v393.504h111.776z" p-id="9849" fill="#4d4d4d"></path></svg>`

	svg_subitem2 = `<svg t="1644821694360" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10142" width="20" height="20"><path d="M465.728 707.2a48.96 48.96 0 1 0 97.92 0 48.96 48.96 0 1 0-97.92 0zM398.08 309.952c-32.32 35.136-67.328 138.24 7.936 140.032 41.984 1.024 41.6-25.088 43.712-31.296s6.336-35.648 19.008-47.936 18.88-18.368 38.592-18.368c19.328 0 34.624 5.12 46.144 15.168 11.584 10.112 17.28 22.336 17.28 36.672a54.848 54.848 0 0 1-7.04 28.416c-4.736 7.936-18.304 22.464-40.896 43.584-23.424 22.016-40.64 45.568-51.008 70.528-22.656 54.656 61.056 71.36 77.888 47.104 16.512-23.872 6.912-28.48 46.528-67.968C616 507.2 628.416 493.696 633.6 485.568c8.96-13.44 15.296-26.688 19.392-39.616a139.072 139.072 0 0 0-13.056-110.464c-38.912-68.48-160.96-113.472-241.856-25.536z" p-id="10143" fill="#4d4d4d"></path><path d="M508.544 66.88A448 448 0 0 0 60.992 514.432a448 448 0 0 0 447.552 447.552 448 448 0 0 0 447.552-447.552A448 448 0 0 0 508.544 66.88z m0 831.104a384 384 0 0 1-383.552-383.552A384 384 0 0 1 508.544 130.88a384 384 0 0 1 383.552 383.552 384 384 0 0 1-383.552 383.552z" p-id="10144" fill="#4d4d4d"></path></svg>`
)
