// 月历卡片
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 460, 330, "月历卡片", 0, xcc.Window_Style_Default)

	// 创建MonthCal
	monthCal := widget.NewMonthCal(30, 40, 290, 240, w.Handle)

	// 用形状文本在窗口上显示当前选中的日期
	stDate := widget.NewShapeText(30, 285, 400, 24, "当前选中:", w.Handle)

	// 读取月历选中日期并刷新到界面.
	updateText := func() {
		var pnYear, pnMonth, pnDay int32
		monthCal.GetSelDate(&pnYear, &pnMonth, &pnDay)
		stDate.SetText(fmt.Sprintf("当前选中: %d年%d月%d日", pnYear, pnMonth, pnDay))
		// 修改元素显示内容后必须重绘, 否则界面不刷新
		stDate.Redraw()
	}

	// 注册月历元素日期改变事件, 使用 MonthCal 的 SetToday 方法也会触发.
	monthCal.AddEvent_MonthCal_Change(func(hEle int, pbHandled *bool) int {
		updateText()
		return 0
	})

	// 获取月历里的今天按钮.
	btnToday := widget.NewButtonByHandle(monthCal.GetButton(xcc.MonthCal_Button_Type_Today))

	// 按钮: 用 SetToday(年, 月, 日) 配合点击 btnToday 跳转到今天.
	btnJumpToday := widget.NewButton(340, 40, 100, 30, "跳到今天", w.Handle)
	btnJumpToday.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		n := time.Now()
		// SetToday 设置月历选中的年月日, 并没有使月历中显示的那个选中日期状态发生改变
		monthCal.SetToday(int32(n.Year()), int32(n.Month()), int32(n.Day()))
		// 发送点击事件, 会把月历中显示的选中日期改变了, 就是让月份跳转到上面设置的日期那里
		btnToday.PostEvent(xcc.XE_BNCLICK, 0, 0)
		monthCal.Redraw(false)
		return 0
	})

	// 按钮: 随机跳转到一个 2000~2030 年的日期.
	btnRandom := widget.NewButton(340, 80, 100, 30, "随机日期", w.Handle)
	btnRandom.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		year := int32(2000 + rand.Intn(31))
		month := int32(1 + rand.Intn(12))
		day := int32(1 + rand.Intn(28))
		monthCal.SetToday(year, month, day)
		btnToday.PostEvent(xcc.XE_BNCLICK, 0, 0)
		monthCal.Redraw(false)
		return 0
	})

	// 初始显示一次当前选中日期.
	updateText()

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
