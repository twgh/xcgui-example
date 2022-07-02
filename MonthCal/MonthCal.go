// 月历卡片
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
	w := window.New(0, 0, 400, 300, "", 0, xcc.Window_Style_Default)

	// 创建MonthCal
	monthCal := widget.NewMonthCal(30, 40, 290, 240, w.Handle)
	// 注册月历元素日期改变事件
	monthCal.Event_MONTHCAL_CHANGE(func(pbHandled *bool) int {
		// 获取被选择的年月日
		var pnYear, pnMonth, pnDay int
		monthCal.GetSelDate(&pnYear, &pnMonth, &pnDay)

		fmt.Printf("%d年%d月%d日\n", pnYear, pnMonth, pnDay)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
