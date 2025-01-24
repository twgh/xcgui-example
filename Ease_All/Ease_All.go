// 全部缓动类型
package main

import (
	"runtime"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/ease"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	w *window.Window

	m_easeFlag           = xcc.Ease_Type_Out // 缓动方式
	m_easeType   int32   = 11                // 缓动类型
	m_pos        int32   = 0                 // 当前位置
	m_time       int32   = 60                // 缓动点数量
	m_time_pos   int32   = 0                 // 当前点
	m_rect       xc.RECT                     // 窗口客户区坐标
	m_windowType = 2                         // 窗口水平或垂直缓动
)

func main() {
	// 这是必要的, 这将保证main函数中对UI库命令的调用是在一个系统线程中执行的。
	// 如果不在一个系统线程中执行, 那程序大概率卡死.
	// 其他例子里没有加是因为简单的例子确实不需要这两句代码, 总之从初始化到Run需要保证是在一个系统线程中执行.
	// 程序运行就窗口卡死未响应就是go的运行时调度的原因, 还没到Run就切换到其他线程了, 比如你在Run前http访问网页了main中加上这两句就不会有问题, 不加就可能出问题.
	// 因为下面用了time.Sleep(), go的运行时可能会进行调度, 就跳到其他线程了, 所以必须用这个.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	a := app.New(true)
	a.SetPaintFrequency(10)
	w = window.New(0, 0, 700, 450, "炫彩缓动测试", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)

	var left int32 = 30
	var top int32 = 35
	CreateButton(2, 11, left, top, 100, "Linear")
	left += 105
	CreateButton(2, 12, left, top, 100, "Quadratic")
	left += 105
	CreateButton(2, 13, left, top, 100, "Cubic")
	left += 105
	CreateButton(2, 14, left, top, 100, "Quartic")
	left += 105
	CreateButton(2, 15, left, top, 100, "Quintic")
	left += 105

	left = 30
	top += 30
	CreateButton(2, 16, left, top, 100, "Sinusoidal")
	left += 105
	CreateButton(2, 17, left, top, 100, "Exponential")
	left += 105
	CreateButton(2, 18, left, top, 100, "Circular")
	left += 105

	left = 30
	top += 30
	CreateButton(2, 19, left, top, 100, "Elastic")
	left += 105
	CreateButton(2, 20, left, top, 100, "Back")
	left += 105
	CreateButton(2, 21, left, top, 100, "Bounce")
	left += 105

	left = 30
	top += 40
	CreateButton(1, 0, left, top, 100, "easeIn")
	left += 105
	CreateButton(1, 1, left, top, 100, "easeOut")
	left += 105
	CreateButton(1, 2, left, top, 100, "easeInOut")
	left += 105

	btn := widget.NewButton(445, top, 100, 24, "快速", w.Handle)
	btn.SetTypeEx(xcc.Button_Type_Check)
	btn.SetCheck(true)
	btn.Event_BUTTON_CHECK(OnBtnCheckSlow)

	btn = widget.NewButton(445, 65, 100, 24, "从左向右", w.Handle)
	btn.SetTypeEx(xcc.Button_Type_Radio)
	btn.SetGroupID(3)
	btn.Event_BUTTON_CHECK(OnBtnCheck_LeftToRight)

	btn = widget.NewButton(445, 92, 100, 24, "从上向下", w.Handle)
	btn.SetTypeEx(xcc.Button_Type_Radio)
	btn.SetGroupID(3)
	btn.SetCheck(true)
	btn.Event_BUTTON_CHECK(OnBtnCheck_TopToBottom)

	btn = widget.NewButton(550, 65, 110, 50, "Run - 窗口缓动", w.Handle)
	btn.Event_BnClick(OnBtnStartWindow)

	btn = widget.NewButton(550, 120, 110, 50, "Run - 缓动曲线", w.Handle)
	btn.Event_BnClick(OnBtnStart)

	w.AdjustLayout()
	w.ShowWindow(xcc.SW_SHOW)

	// 窗口绘制事件
	w.Event_PAINT(OnDrawWindow)

	// 获取窗口坐标
	rc := w.GetRectDPI()
	// 第一次缓动
	for t := 0; t <= 30; t++ {
		v := ease.Bounce(float32(t)/30.0, xcc.Ease_Type_Out)
		y := int32(v * float32(rc.Top))
		w.SetPosition(rc.Left, y).Redraw(true)
		time.Sleep(time.Millisecond * 10)
	}

	a.Run()
	a.Exit()
}

// 创建按钮
func CreateButton(nGroup, id, x, y, cx int32, title string) {
	btn := widget.NewButton(x, y, cx, 22, title, w.Handle)
	// 设置为单选按钮
	btn.SetTypeEx(xcc.Button_Type_Radio)
	// 设置按钮组id
	btn.SetGroupID(nGroup)
	// 设置元素ID
	btn.SetID(id)

	if id == 1 || id == 21 {
		btn.SetCheck(true)
	}
	// 注册按钮选中事件
	btn.Event_BUTTON_CHECK1(OnButtonCheck)
}

// 按钮选中事件
func OnButtonCheck(hEle int, bCheck bool, pbHandled *bool) int {
	if !bCheck {
		return 0
	}
	id := xc.XWidget_GetID(hEle)

	if id <= 2 {
		m_easeFlag = xcc.Ease_Type_(id)
	} else {
		m_easeType = id - 10
	}

	w.Redraw(true)
	return 0
}

// 快速
func OnBtnCheckSlow(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_time = 60
	} else {
		m_time = 120
	}
	return 0
}

// 从左向右
func OnBtnCheck_LeftToRight(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_windowType = 1
	}
	return 0
}

// 从上向下
func OnBtnCheck_TopToBottom(bCheck bool, pbHandled *bool) int {
	if bCheck {
		m_windowType = 2
	}
	return 0
}

// 窗口缓动
func OnBtnStartWindow(pbHandled *bool) int {
	rect := w.GetRectDPI()

	time2 := m_time / 2
	for t := int32(0); t <= time2; t++ {
		var v float32
		switch m_easeType {
		case 1:
			v = ease.Linear(float32(t) / float32(time2))
		case 2:
			v = ease.Quad(float32(t)/float32(time2), m_easeFlag)
		case 3:
			v = ease.Cubic(float32(t)/float32(time2), m_easeFlag)
		case 4:
			v = ease.Quart(float32(t)/float32(time2), m_easeFlag)
		case 5:
			v = ease.Quint(float32(t)/float32(time2), m_easeFlag)
		case 6:
			v = ease.Sine(float32(t)/float32(time2), m_easeFlag)
		case 7:
			v = ease.Expo(float32(t)/float32(time2), m_easeFlag)
		case 8:
			v = ease.Circ(float32(t)/float32(time2), m_easeFlag)
		case 9:
			v = ease.Elastic(float32(t)/float32(time2), m_easeFlag)
		case 10:
			v = ease.Back(float32(t)/float32(time2), m_easeFlag)
		case 11:
			v = ease.Bounce(float32(t)/float32(time2), m_easeFlag)
		}

		if m_windowType == 1 {
			x := int32(v * float32(rect.Left))
			w.SetPosition(x, rect.Top)
		} else {
			y := int32(v * float32(rect.Top))
			w.SetPosition(rect.Left, y)
		}
		w.Redraw(true)
		time.Sleep(10 * time.Millisecond)
	}
	return 0
}

// 缓动曲线
func OnBtnStart(pbHandled *bool) int {
	var width float32 = 400.0
	for t := int32(0); t <= m_time; t++ {
		var v float32
		switch m_easeType {
		case 1:
			v = ease.Linear(float32(t) / float32(m_time))
		case 2:
			v = ease.Quad(float32(t)/float32(m_time), m_easeFlag)
		case 3:
			v = ease.Cubic(float32(t)/float32(m_time), m_easeFlag)
		case 4:
			v = ease.Quart(float32(t)/float32(m_time), m_easeFlag)
		case 5:
			v = ease.Quint(float32(t)/float32(m_time), m_easeFlag)
		case 6:
			v = ease.Sine(float32(t)/float32(m_time), m_easeFlag)
		case 7:
			v = ease.Expo(float32(t)/float32(m_time), m_easeFlag)
		case 8:
			v = ease.Circ(float32(t)/float32(m_time), m_easeFlag)
		case 9:
			v = ease.Elastic(float32(t)/float32(m_time), m_easeFlag)
		case 10:
			v = ease.Back(float32(t)/float32(m_time), m_easeFlag)
		case 11:
			v = ease.Bounce(float32(t)/float32(m_time), m_easeFlag)
		}

		m_pos = int32(v * width)
		m_time_pos = t
		time.Sleep(10 * time.Millisecond)

		rc := m_rect
		rc.Top = 170
		w.RedrawRect(&rc, true)
	}
	return 0
}

// 绘制
func OnDrawWindow(hDraw int, pbHandled *bool) int {
	*pbHandled = true

	var rect xc.RECT
	w.GetClientRect(&rect)

	draw := drawx.NewByHandle(hDraw)
	draw.SetBrushColor(xc.RGBA(230, 230, 230, 255))
	draw.FillRect(&rect)
	m_rect = rect

	draw.SetBrushColor(xc.RGBA(200, 200, 200, 255))
	draw.DrawRect(&rect)

	draw.SetBrushColor(xc.RGBA(0, 0, 200, 255))
	draw.TextOutEx(260, 10, "炫彩界面库(XCGUI) - 缓动测试")

	var rc xc.RECT
	rc.Left = 150
	rc.Top = 190
	rc.Right = rc.Left + 400 + 30
	rc.Bottom = rc.Top + 50

	rcBorder := rc
	rcBorder.Left -= 2
	rcBorder.Top -= 2
	rcBorder.Right += 2
	rcBorder.Bottom += 2
	draw.SetBrushColor(xc.RGBA(0, 0, 200, 255))
	draw.DrawRect(&rcBorder)

	rcFill := rc
	rcFill.Left = rcFill.Left + m_pos
	rcFill.Right = rcFill.Left + 30
	draw.SetBrushColor(xc.RGBA(128, 0, 0, 255))
	draw.FillRect(&rcFill)

	var rcBorder_Line xc.RECT
	rcBorder_Line.Left = 150
	rcBorder_Line.Right = 150 + 400
	rcBorder_Line.Top = 255
	rcBorder_Line.Bottom = 255 + 180

	rcBorder = rcBorder_Line
	rcBorder.Right++
	rcBorder.Bottom++
	draw.SetBrushColor(xc.RGBA(180, 180, 180, 255))
	draw.DrawRect(&rcBorder)

	pts := make([]xc.POINTF, 121)
	for t := int32(0); t <= m_time; t++ {
		var v float32
		switch m_easeType {
		case 1:
			v = ease.Linear(float32(t) / float32(m_time))
		case 2:
			v = ease.Quad(float32(t)/float32(m_time), m_easeFlag)
		case 3:
			v = ease.Cubic(float32(t)/float32(m_time), m_easeFlag)
		case 4:
			v = ease.Quart(float32(t)/float32(m_time), m_easeFlag)
		case 5:
			v = ease.Quint(float32(t)/float32(m_time), m_easeFlag)
		case 6:
			v = ease.Sine(float32(t)/float32(m_time), m_easeFlag)
		case 7:
			v = ease.Expo(float32(t)/float32(m_time), m_easeFlag)
		case 8:
			v = ease.Circ(float32(t)/float32(m_time), m_easeFlag)
		case 9:
			v = ease.Elastic(float32(t)/float32(m_time), m_easeFlag)
		case 10:
			v = ease.Back(float32(t)/float32(m_time), m_easeFlag)
		case 11:
			v = ease.Bounce(float32(t)/float32(m_time), m_easeFlag)
		}

		pts[t].X = float32(rc.Left) + float32(t)*400.0/float32(m_time)
		pts[t].Y = float32(rcBorder_Line.Bottom) - v*180.0
	}

	draw.EnableSmoothingMode(true)
	draw.SetBrushColor(xc.RGBA(128, 0, 0, 255))

	left := rc.Left + int32(float32(m_time_pos)*400.0/float32(m_time))

	draw.DrawLine(left, rcBorder_Line.Top, left, rcBorder_Line.Bottom)
	draw.DrawCurveF(pts, m_time+1, 0.5)
	return 0
}
