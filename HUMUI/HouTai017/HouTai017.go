// 未闻花名 - 后台界面017
// https://mall.xcgui.com/1562.html
package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"time"

	"github.com/twgh/xcgui/ani"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/bkmanager"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
	"github.com/twgh/xcgui/tmpl"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed res/data.zip
	resData []byte // 打包的炫彩资源文件

	w *window.Window // 主窗口
)

func main() {
	// 初始化UI库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 加载资源文件
	a.LoadResourceZipMem(resData, "资源文件\\resource.res", "")
	// 设置默认字体
	a.SetDefaultFont(font.NewEX("微软雅黑", 10, xcc.FontStyle_Regular).Handle)

	// 创建窗口从布局文件
	w = window.NewByLayoutZipMem(resData, "布局文件\\main.xml", "", 0, 0)

	// 窗口事件
	setWindowEvent()
	// 左侧导航, 导航展开收缩按钮属性设置
	setNavAdjAttr()
	// 美化日期时间元素
	setDateTimeUI()
	// 加载列表
	loadList()
	// 流量明细UI
	setDetailUI()

	// 调整布局
	w.AdjustLayout()
	w.PostMessage(wapi.WM_SIZE, 0, 0) // 为了调整窗口左下角按钮 navAdj 的位置
	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}

// 窗口事件
func setWindowEvent() {
	// 左侧导航-展开收缩按钮
	btnNavAdj := widget.NewButtonByName("navAdj")
	// 窗口消息处理
	w.AddEvent_WindProc(func(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
		switch message {
		case wapi.WM_CLOSE: // 窗口关闭时，释放图片
			destroyDTIcons()
		case wapi.WM_SIZE:
			rc := w.GetRectEx()
			btnNavAdj.SetPosition(16, rc.Bottom-rc.Top-44, false, 0, 0)
		}
		return 0
	})
	// 双击窗口标题栏切换最大化状态
	w.AddEvent_LButtonDBClick(func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int {
		if pPt.Y <= 35 {
			w.MaxWindow(!w.IsMaxWindow())
		}
		return 0
	})
}

// 左侧导航, 导航展开收缩按钮属性设置
func setNavAdjAttr() {
	svg_hide := svg.NewByZipMem(resData, "资源文件\\nav_hide.svg", "").SetSize(20, 20)
	svg_show := svg.NewByZipMem(resData, "资源文件\\nav_show.svg", "").SetSize(20, 20)
	img_hide := imagex.NewBySvg(svg_hide.Handle).AddRef()
	img_show := imagex.NewBySvg(svg_show.Handle).AddRef()
	// 获取左侧导航布局元素
	navLayout := widget.NewLayoutEleByName("nav")
	// 获取左侧导航-展开收缩按钮, 设置图标
	btnNavAdj := widget.NewButtonByName("navAdj").SetIcon(img_hide.Handle)
	btnNavAdj.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		app.SetPaintFrequency(10)
		defer app.SetPaintFrequency(30)

		width2 := float32(300)
		// 获取左侧导航宽度
		width := navLayout.GetWidth()
		if width == 0 {
			btnNavAdj.SetIcon(img_hide.Handle).SetToolTip("收起").Redraw(false)
		} else {
			btnNavAdj.SetIcon(img_show.Handle).SetToolTip("展开").Redraw(false)
			width2 = 0
		}
		// 展开收缩动画
		ani1 := ani.NewAnima(navLayout.Handle, 1)
		ani1.LayoutWidth(300, xcc.Layout_Size_Fixed, width2, 1, xcc.Ease_Flag_Expo, false)
		ani1.Run(w.Handle)
		return 0
	})
}

var m_DTIcons []int // 日期时间图标句柄

// 美化日期时间元素
func setDateTimeUI() {
	// 加载日期时间图标
	m_DTIcons = append(m_DTIcons,
		imagex.NewByZipMem(resData, "资源文件\\ll.png", "").Handle,
		imagex.NewByZipMem(resData, "资源文件\\rr.png", "").Handle,
		imagex.NewByZipMem(resData, "资源文件\\l.png", "").Handle,
		imagex.NewByZipMem(resData, "资源文件\\r.png", "").Handle,
	)
	// 禁止自动销毁日期时间图标
	for i := 0; i < len(m_DTIcons); i++ {
		xc.XImage_EnableAutoDestroy(m_DTIcons[i], false)
	}

	for i := 1; i <= 4; i++ {
		// 获取日期时间元素
		dt := widget.NewDateTimeByName(fmt.Sprintf("date%d", i))
		// 普通按钮替代日期按钮
		widget.NewButtonByName(fmt.Sprintf("date_btn%d", i)).AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			dt.Popup()
			return 0
		})
		// 设置日期时间元素的属性
		setDateTimeAttr(dt)
	}
}

// 设置日期时间元素的属性
func setDateTimeAttr(dt *widget.DateTime) {
	// 设置日期分割栏为-
	dt.EnableSplitSlash(false)
	// 设置文本颜色
	dt.SetTextColor(xc.RGBA(43, 43, 46, 255))
	// 设置选择日期背景颜色
	dt.SetSelBkColor(xc.RGBA(200, 227, 255, 255))
	// 日期时间元素弹出月历卡片事件
	dt.AddEvent_DateTime_Popup_MonthCal(func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int {
		// 月历卡片窗口
		mcw := window.NewByHandle(hMonthCalWnd)
		mcw.SetTransparentType(xcc.Window_Transparent_Shadow)
		mcw.SetTransparentAlpha(255)
		rc := mcw.GetRectEx()
		rc.Right += 40
		rc.Bottom += 30
		mcw.SetRect(&rc)
		// 月历卡片元素
		mc := widget.NewMonthCalByHandle(hMonthCal)
		mc.SetSize(rc.Right-rc.Left-40, rc.Bottom-rc.Top, false, xcc.AdjustLayout_All, 0)
		mcw.SetShadowInfo(10, 50, 5, false, xcc.COLOR_BLACK)
		mc.SetPosition(20, 20, true, xcc.AdjustLayout_Self, 0)
		mc.EnableBkTransparent(true)
		mc.EnableDrawBorder(false)
		mc.EnableFocus(false)
		xc.XEle_SetTextColor(hMonthCal, xc.RGBA(96, 98, 102, 255))
		mc.SetHeight(234)
		mc.SetFont(font.New(10).Handle)
		mc.SetTextColor(1, xc.RGBA(96, 98, 102, 255))
		mc.SetTextColor(2, xc.RGBA(96, 98, 102, 255))
		// 设置处理颜色
		bgColor := xcc.COLOR_WHITE
		shapeRect := widget.NewShapeRect(0, 0, rc.Right-rc.Left, rc.Bottom-rc.Top, hMonthCalWnd)
		shapeRect.SetRoundAngle(10, 10)
		shapeRect.SetFillColor(bgColor)
		bkm := bkmanager.NewByHandle(mc.GetBkManager())
		bkm.AddFill(xcc.Element_State_Flag_Leave, bgColor, 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Last_Month|xcc.MonthCal_State_Flag_Item_Select_No|xcc.MonthCal_State_Flag_Item_Leave, xc.RGBA(81, 167, 255, 20), 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Next_Month|xcc.MonthCal_State_Flag_Item_Select_No|xcc.MonthCal_State_Flag_Item_Leave, xc.RGBA(81, 167, 255, 20), 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Cur_Month|xcc.MonthCal_State_Flag_Item_Select_No|xcc.MonthCal_State_Flag_Item_Leave, xc.RGBA(81, 167, 255, 5), 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Select|xcc.MonthCal_State_Flag_Leave|xcc.MonthCal_State_Flag_Item_Leave, xc.RGBA(96, 191, 255, 255), 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Select|xcc.MonthCal_State_Flag_Item_Stay, xc.RGBA(50, 149, 225, 255), 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Stay|xcc.MonthCal_State_Flag_Item_Select_No, xc.RGBA(81, 167, 255, 80), 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Down|xcc.MonthCal_State_Flag_Item_Select_No, xc.RGBA(81, 167, 255, 200), 0)
		bkm.AddFill(xcc.MonthCal_State_Flag_Item_Select|xcc.MonthCal_State_Flag_Item_Down, xc.RGBA(81, 167, 255, 200), 0)
		// 设置按钮样式
		bgColor = xc.RGBA(192, 196, 204, 255)
		xc.XEle_Destroy(mc.GetButton(xcc.MonthCal_Button_Type_Today)) // 销毁今天按钮
		mcButtonTypes := []xcc.MonthCal_Button_Type_{xcc.MonthCal_Button_Type_Last_Year, xcc.MonthCal_Button_Type_Next_Year, xcc.MonthCal_Button_Type_Last_Month, xcc.MonthCal_Button_Type_Next_Month}
		mcButtonToolTips := []string{"上一年", "下一年", "上一月", "下一月"}
		posXs := []int32{0, rc.Right - rc.Left - 60, 20, rc.Right - rc.Left - 80}
		for i := 0; i < 4; i++ {
			btn := widget.NewButtonByHandle(mc.GetButton(mcButtonTypes[i]))
			btn.SetToolTip(mcButtonToolTips[i])
			btn.SetSize(20, 20, false, xcc.AdjustLayout_Self, 0)
			btn.SetPosition(posXs[i], 0, false, xcc.AdjustLayout_Self, 0)
			btn.EnableDrawBorder(false)
			btn.SetText("")
			btn.SetTextColor(bgColor)
			btn.SetIcon(m_DTIcons[i])
			btn.SetOffsetIcon(0, -3)
		}
		return 0
	})
}

// 销毁日期时间图标
func destroyDTIcons() {
	for i := 0; i < len(m_DTIcons); i++ {
		xc.XImage_Destroy(m_DTIcons[i])
	}
}

// 加载列表
func loadList() {
	hParent := app.GetObjectByName("home_listBox")
	widget.NewLayoutEleByLayoutZipMemEx(resData, "布局文件\\list.xml", "", "home", hParent, 0, 0)

	ls := widget.NewListByUID(10001)
	// 设置两遍, 一遍列表头, 一遍列表项
	hTmp := tmpl.NewByZipMem(xcc.ListItemTemp_Type_List_Head, resData, "布局文件\\listTemp.xml", "").Handle
	ls.SetItemTemplate(hTmp) // 列表_置项模板, 列表头
	hTmp = tmpl.NewByZipMem(xcc.ListItemTemp_Type_List_Item, resData, "布局文件\\listTemp.xml", "").Handle
	ls.SetItemTemplate(hTmp) // 列表_置项模板, 列表项

	setListAttr(ls) // 设置列表属性
	addListData(ls) // 列表加载默认数据
}

// 设置列表属性
func setListAttr(ls *widget.List) {
	ls.SetHeaderHeight(40) // 列表_置列表头高度
	ls.EnableMultiSel(true)
	ls.SetSplitLineColor(xc.RGBA(239, 240, 241, 255))
	ls.SetRowSpace(0)
	ls.EnableDragChangeColumnWidth(false)
	ls.SetRowHeightDefault(76, 76)
	ls.SetDrawRowBkFlags(xcc.List_DrawItemBk_Flag_Leave | xcc.List_DrawItemBk_Flag_Stay | xcc.List_DrawItemBk_Flag_Line)
	ls.CreateAdapterHeader()
	ls.CreateAdapter(0)
	// 列表自适应窗口宽度
	w.AddEvent_WindProc(func(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
		if message == 0x0047 { // WM_WINDOWPOSCHANGED 窗口位置改变
			hParentObj := ls.GetParentEle()
			if hParentObj < 1 {
				return 0
			}

			var rc xc.RECT
			if app.IsHWINDOW(hParentObj) {
				xc.XWnd_GetClientRect(hParentObj, &rc)
			} else {
				xc.XWnd_AdjustLayout(xc.XWidget_GetHWINDOW(hParentObj))
				xc.XEle_GetClientRect(hParentObj, &rc)
			}

			var count = ls.GetColumnCount()
			var isAverage = false // 是否均分宽度
			var width int32
			if isAverage { // 每列均分宽度
				width = (rc.Right - rc.Left) / count
			} else { // 有三列固定宽度和为270, 剩余的三列均分剩余宽度
				count = 3                                  // 剩余的列数为3
				width = (rc.Right - rc.Left - 270) / count // (总宽度 - 固定列宽度的和270) / 剩余的列数3
			}
			indexs := []int32{1, 4, 5} // 除了三列固定, 剩余的三列的索引

			isSet := false // 是否有设置列宽
			for i := int32(0); i < count; i++ {
				index := i
				if !isAverage {
					index = indexs[i]
				}

				if width == ls.GetColumnWidth(index) { // 如果已经设置，直接返回
					return 0
				}
				ls.SetColumnWidth(index, width)
				isSet = true
			}
			if isSet {
				if app.IsHWINDOW(hParentObj) {
					xc.XWnd_AdjustLayout(hParentObj)
				} else {
					xc.XEle_AdjustLayout(hParentObj, 0)
				}
				ls.Redraw(false)
			}
		}
		return 0
	}, true) // 这里要特别注意, 因为该窗口的windproc事件上面已经添加过了, AddEvent开头的函数你想给一个元素的一个类型事件添加多个回调函数, 最后这个参数必须为true, 不然会覆盖掉之前添加的那个

	// 列表头按钮被选择事件
	onBtnCheckListListHeader := func(hEle int, bCheck bool, pbHandled *bool) int {
		for i := int32(0); i < ls.GetCount_AD(); i++ {
			hBtn := ls.GetTemplateObject(i, 0, 1)
			if app.IsHXCGUI(hBtn, xcc.XC_BUTTON) { // 未显示的无法获取，因为没有创建出来
				xc.XBtn_SetCheck(hBtn, bCheck)
			}
			ls.SetItemData(i, 0, common.BoolToInt(bCheck)) // 设置当前项数据，下次创建模板时，复原选中状态
		}
		ls.Redraw(false)
		return 0
	}

	// 列表头项模板创建完成事件
	ls.AddEvent_List_Header_Temp_Create_End(func(hEle int, pItem *xc.List_Header_Item_, pbHandled *bool) int {
		// 获取列表头按钮
		btn := widget.NewButtonByHandle(ls.GetHeaderTemplateObject(0, 1))
		if app.IsHXCGUI(btn.Handle, xcc.XC_BUTTON) {
			// 列表头按钮被选择事件
			btn.AddEvent_Button_Check(onBtnCheckListListHeader)
		}
		return 0
	})
	// 列表项绘制事件
	ls.AddEvent_List_DrawItem(func(hEle int, hDraw int, pItem *xc.List_Item_, pbHandled *bool) int {
		*pbHandled = true
		// 隔行换色
		/*if pItem.Index%2 == 1 {
			xc.XDraw_SetBrushColor(hDraw, xc.RGBA(249, 249, 251, 255))
			xc.XDraw_FillRect(hDraw, &pItem.RcItem)
		}*/

		switch xcc.List_DrawItemBk_Flag_(pItem.NState) {
		case xcc.List_DrawItemBk_Flag_Stay:
			xc.XDraw_SetBrushColor(hDraw, xc.RGBA(226, 249, 254, 255))
			xc.XDraw_FillRect(hDraw, &pItem.RcItem)
		case xcc.List_DrawItemBk_Flag_Leave:
			xc.XDraw_SetBrushColor(hDraw, xc.RGBA(240, 251, 255, 255))
			xc.XDraw_FillRect(hDraw, &pItem.RcItem)
		}
		return 0
	})
	// 列表鼠标左键弹起事件
	ls.AddEvent_LButtonUp(func(hEle int, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		iItem := ls.GetSelectRow()
		if iItem < 0 {
			return 0
		}
		hBtn := ls.GetTemplateObject(iItem, 0, 1)
		if app.IsHXCGUI(hBtn, xcc.XC_BUTTON) {
			isCheck := xc.XBtn_IsCheck(hBtn)
			xc.XBtn_SetCheck(hBtn, !isCheck)
			xc.XEle_Redraw(hBtn, false)
			ls.SetItemData(iItem, 0, common.BoolToInt(xc.XBtn_IsCheck(hBtn))) // 设置当前项数据，下次创建模板时，复原选中状态
		}
		return 0
	})

	// 运行中, 暂停中, 已停止, 未运行
	listStateColors := []uint32{xc.RGBA(75, 198, 121, 255), xc.RGBA(252, 189, 80, 255), xc.RGBA(253, 138, 138, 255), xc.RGBA(121, 121, 123, 255)}

	// 列表项模板创建完成事件
	ls.AddEvent_List_Temp_Create_End(func(hEle int, pItem *xc.List_Item_, nFlag int32, pbHandled *bool) int {
		// 设置首页列表 状态项 对应文本颜色
		for i := int32(0); i < ls.GetCount_AD(); i++ {
			hShapeText := xc.XList_GetTemplateObject(hEle, i, 3, 1)
			if xc.XC_IsHXCGUI(hShapeText, xcc.XC_SHAPE_TEXT) {
				usData := xc.XList_GetItemData(hEle, i, 3)
				xc.XShapeText_SetTextColor(hShapeText, listStateColors[usData])
			}
		}
		// 取模板设置按钮
		btn := widget.NewButtonByHandle(ls.GetTemplateObject(pItem.Index, 5, 1))
		if app.IsHXCGUI(btn.Handle, xcc.XC_BUTTON) {
			// 列表项按钮点击事件
			btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
				itemIndex := ls.GetRowIndexFromHXCGUI(btn.Handle)
				fmt.Println("列表设置按钮被点击 OnBtnListItem() 当前所在项:", itemIndex+1)
				return 0
			})

			btn2 := widget.NewButtonByHandle(ls.GetTemplateObject(pItem.Index, 0, 1))
			if app.IsHXCGUI(btn2.Handle, xcc.XC_BUTTON) {
				// 按钮被选择事件
				btn2.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
					itemIndex := ls.GetRowIndexFromHXCGUI(btn2.Handle)
					ls.SetItemData(itemIndex, 0, common.BoolToInt(bCheck)) // 设置当前项数据，下次创建模板时，复原选中状态
					fmt.Println("列表按钮被点击 OnBtnCheckList() 当前所在项:", itemIndex+1)
					return 0
				})
				if pItem.NUserData == 1 { // OnListSelect
					btn2.SetCheck(true)
				}
			}
		}
		return 0
	})
}

// 列表加载默认数据
func addListData(ls *widget.List) {
	rand.Seed(time.Now().UnixNano())
	texts := []string{"运行中", "暂停中", "已停止", "-"}
	s := []string{"本界面由 未闻花名UI 设计", "炫彩界面库 XCGUI"}

	ls.AddColumnText(90, "name1", "编号")
	ls.AddColumnText(280, "name2", "项目名称")
	ls.AddColumnText(100, "name3", "软件版本")
	ls.AddColumnText(120, "name4", "状态")
	ls.AddColumnText(300, "name5", "说明")
	ls.AddColumnText(200, "name6", "操作")

	for i := int32(0); i < 20; i++ {
		ls.AddRowText(xc.Itoa(i + 1))
		ls.SetItemText(i, 1, "未闻花名UI "+xc.Itoa(i+1))
		ls.SetItemText(i, 2, "2.0.1248")
		index := rand.Intn(4)
		ls.SetItemData(i, 3, index)
		ls.SetItemText(i, 3, texts[index])
		ls.SetItemText(i, 4, s[rand.Intn(2)])
	}
}

// 流量明细UI
func setDetailUI() {
	names := []string{"剩余流量", "会员剩余流量"}
	for _, name := range names {
		setDetailAttr(widget.NewElementByName(name), true, 60)
	}
}

// 流量明细 - 圆盘 属性
//   - ele: 圆盘元素
//   - bCCW: 是否逆时针绘制
//   - pos: 绘制进度，0-100
func setDetailAttr(ele *widget.Element, bCCW bool, pos float32) {
	// 元素绘制事件
	ele.AddEvent_Paint(func(hEle int, hDraw int, pbHandled *bool) int {
		*pbHandled = true
		d := drawx.NewByHandle(hDraw)
		d.EnableSmoothingMode(true)
		d.EnableWndTransparent(true)
		d.SetBrushColor(xc.RGBA(114, 123, 134, 100))
		width := ele.GetWidth() - 16
		height := ele.GetHeight() - 16
		d.SetLineWidth(8)
		d.DrawArc(8, 8, width, height, 0, 360)
		d.SetBrushColor(xcc.COLOR_WHITE)

		angle := pos * 360.0 / 100.0
		if bCCW {
			d.DrawArc(8, 8, width, height, -90, angle)
		} else {
			d.DrawArc(8, 8, width, height, -90+360-angle, angle)
		}

		rc := ele.GetRectEx()
		d.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)
		d.DrawText(xc.Ftoa(pos)+"%", &rc)
		return 0
	})
}
