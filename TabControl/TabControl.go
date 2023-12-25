// 纯代码模拟选择夹(选项卡)切换页面
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	a *app.App
	w *window.Window

	layoutTab *widget.LayoutEle
	tabBtn1   *widget.Button
	tabBtn2   *widget.Button
	tabBtn3   *widget.Button

	layoutBody  *widget.LayoutEle
	layoutPage1 *widget.LayoutEle
	layoutPage2 *widget.LayoutEle
	layoutPage3 *widget.LayoutEle
)

func main() {
	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 600, 400, "选择夹切换页面", 0, xcc.Window_Style_Default)

	// 我是喜欢创建一个水平布局元素, tab按钮都放在里面
	layoutTab = widget.NewLayoutEle(14, 35, 500, 30, w.Handle)
	// 放在水平布局元素中的组件, x, y绝对坐标是无效的
	tabBtn1 = widget.NewButton(0, 0, 100, 30, "页面1", layoutTab.Handle)
	tabBtn2 = widget.NewButton(0, 0, 100, 30, "页面2", layoutTab.Handle)
	tabBtn3 = widget.NewButton(0, 0, 100, 30, "页面3", layoutTab.Handle)

	// 设为单选类型按钮
	tabBtn1.SetTypeEx(xcc.Button_Type_Radio)
	tabBtn2.SetTypeEx(xcc.Button_Type_Radio)
	tabBtn3.SetTypeEx(xcc.Button_Type_Radio)
	// 设置为默认按钮样式, 就不会是单选按钮样式了
	tabBtn1.SetStyle(xcc.Button_Style_Default)
	tabBtn2.SetStyle(xcc.Button_Style_Default)
	tabBtn3.SetStyle(xcc.Button_Style_Default)
	// 默认选中第一个
	tabBtn1.SetCheck(true)

	// 主体部分, 页面都放进这里面, 我是喜欢这样设计, 不是必须
	layoutBody = widget.NewLayoutEle(14, 65, 500, 350, w.Handle)
	// 第一页
	layoutPage1 = widget.NewLayoutEle(0, 0, 500, 350, layoutBody.Handle)
	// 第二页
	layoutPage2 = widget.NewLayoutEle(0, 0, 500, 350, layoutBody.Handle)
	// 第三页
	layoutPage3 = widget.NewLayoutEle(0, 0, 500, 350, layoutBody.Handle)
	// 只让第一页显示, 其他页都设为不显示
	layoutPage2.Show(false)
	layoutPage3.Show(false)

	// 给按钮绑定页面, 绑定后切换页面的原理就是: 你点哪个按钮就显示哪个页面
	tabBtn1.SetBindEle(layoutPage1.Handle)
	tabBtn2.SetBindEle(layoutPage2.Handle)
	tabBtn3.SetBindEle(layoutPage3.Handle)

	// 页面中放置内容
	widget.NewShapeText(0, 0, 100, 30, "我是第一页", layoutPage1.Handle)
	widget.NewShapeText(0, 0, 100, 30, "我是第二页", layoutPage2.Handle)
	widget.NewShapeText(0, 0, 100, 30, "我是第三页", layoutPage3.Handle)

	w.Show(true)
	a.Run()
	a.Exit()
}
