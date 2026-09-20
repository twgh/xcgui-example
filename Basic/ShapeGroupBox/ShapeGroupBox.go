// 组框形状对象
package main

// 演示 ShapeGroupBox: 设置标题文字/字体/颜色、边框颜色、圆角、文本偏移, 并在其中放置按钮展示分组效果.
// 形状对象没有事件, 重点在于演示各种外观设置 API.

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 520, 440, "ShapeGroupBox", 0, xcc.Window_Style_Default)
	// 设置窗口边框大小
	w.SetBorderSize(1, 30, 1, 1)

	// 给组框标题用的粗体字体
	titleFont := font.NewEX("微软雅黑", 12, xcc.FontStyle_Bold)

	// ---------------- 组框一: 基础用法 ----------------
	// 创建时通过 name 参数直接指定标题文字
	gb1 := widget.NewShapeGroupBox(20, 45, 230, 170, "基础设置", w.Handle)
	gb1.SetBorderColor(xc.HexRGB2RGBA("#4285F4", 255))
	gb1.SetTextColor(xc.HexRGB2RGBA("#4285F4", 255))
	gb1.SetFontX(titleFont.Handle)

	// 组框是形状对象, 不能作为元素的父对象, 把按钮创建到窗口上,
	// 坐标落在组框范围内即可呈现分组效果.
	widget.NewButton(45, 85, 180, 32, "开机自启动", w.Handle)
	widget.NewButton(45, 125, 180, 32, "关闭主窗口最小化", w.Handle)
	widget.NewButton(45, 165, 180, 32, "退出时清空缓存", w.Handle)

	// ---------------- 组框二: 圆角与文本偏移 ----------------
	gb2 := widget.NewShapeGroupBox(270, 45, 230, 170, "网络设置", w.Handle)
	// 启用圆角并设置圆角大小
	gb2.EnableRoundAngle(true)
	gb2.SetRoundAngle(16, 16)
	gb2.SetBorderColor(xc.HexRGB2RGBA("#34A853", 255))
	gb2.SetTextColor(xc.HexRGB2RGBA("#188038", 255))
	// 设置标题文本偏移量, 让标题向右下移动
	gb2.SetTextOffset(20, 2)

	widget.NewButton(295, 85, 180, 32, "使用系统代理", w.Handle)
	widget.NewButton(295, 125, 180, 32, "自动检测网络", w.Handle)
	widget.NewButton(295, 165, 180, 32, "测试连接", w.Handle)

	// ---------------- 组框三: 空标题横排按钮 ----------------
	// 标题传空字符串就是一个纯边框分组容器
	gb3 := widget.NewShapeGroupBox(20, 235, 480, 90, "", w.Handle)
	gb3.SetBorderColor(xc.HexRGB2RGBA("#DADCE0", 255))

	widget.NewButton(45, 265, 130, 32, "保存", w.Handle)
	widget.NewButton(195, 265, 130, 32, "重置", w.Handle)
	widget.NewButton(345, 265, 130, 32, "取消", w.Handle)

	// ---------------- 组框四: 动态设置标题 ----------------
	gb4 := widget.NewShapeGroupBox(20, 345, 480, 70, "", w.Handle)
	gb4.SetBorderColor(xc.HexRGB2RGBA("#EA4335", 255))
	gb4.SetTextColor(xc.HexRGB2RGBA("#EA4335", 255))
	// 创建后用 SetText 动态设置标题文字
	gb4.SetText("关于")
	// 用 GetText 可以取回标题文字
	println("组框标题: " + gb4.GetText())

	widget.NewShapeText(45, 375, 400, 24, "ShapeGroupBox 示例 — 组框是一个纯外观的形状对象, 没有事件", w.Handle)

	w.Show(true)
	a.Run()
	a.Exit()
}
