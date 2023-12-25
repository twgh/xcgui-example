// 月历卡片
// 要美化的话, 就得自绘, 看这个: http://www.xcgui.com/doc-ui/page_draw__month_cal.html
// 我以后有空会翻译几个好看的: http://mall.xcgui.com/1618.html
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 400, 300, "月历卡片", 0, xcc.Window_Style_Default)

	// 创建MonthCal
	monthCal := widget.NewMonthCal(30, 40, 290, 240, w.Handle)
	// 注册月历元素日期改变事件
	monthCal.Event_MONTHCAL_CHANGE(func(pbHandled *bool) int {
		// 获取被选择的年月日
		var pnYear, pnMonth, pnDay int32
		monthCal.GetSelDate(&pnYear, &pnMonth, &pnDay)

		fmt.Printf("%d年%d月%d日\n", pnYear, pnMonth, pnDay)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
