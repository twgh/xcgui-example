// 美化进度条
package main

// 本例子是由 AI 调用 go-xcgui-dev 技能生成的代码.
import (
	"fmt"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/drawx"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// BeautifyProgressBar 美化进度条.
//
// 继承 widget.ProgressBar, 通过自绘实现圆角进度条外观,
// 支持设置颜色/圆角/文字等, 值范围为 0.0 ~ 1.0.
type BeautifyProgressBar struct {
	*widget.ProgressBar
	font       int     // 字体句柄
	radius     int32   // 圆角半径
	value      float32 // 进度值 0.0~1.0
	trackColor uint32  // 轨道背景颜色
	fillColor  uint32  // 进度填充颜色
	textColor  uint32  // 进度百分比文字颜色
}

// NewBeautifyProgressBar 创建美化进度条.
//
// x, y: 坐标.
//
// cx, cy: 宽高.
//
// font: 字体句柄.
//
// hParent: 父句柄.
func NewBeautifyProgressBar(x, y, cx, cy int32, font int, hParent int) *BeautifyProgressBar {
	bar := widget.NewProgressBar(x, y, cx, cy, hParent)
	if bar == nil {
		return nil
	}
	p := &BeautifyProgressBar{
		ProgressBar: bar,
		font:        font,
		value:       0,
		radius:      12,
		trackColor:  xc.RGBA(230, 230, 230, 255),
		fillColor:   xc.RGBA(66, 133, 244, 255),
		textColor:   xc.RGBA(80, 80, 80, 255),
	}
	// 隐藏原生文字, 使用自绘
	p.EnableShowText(false)
	// 注册自绘事件
	p.AddEvent_Paint(p.onPaint)
	return p
}

// SetValue 设置进度值 (0.0~1.0).
func (p *BeautifyProgressBar) SetValue(v float32) {
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	p.value = v
	p.ProgressBar.SetPos(int32(v * 100))
	xc.XEle_Redraw(p.Handle, false)
}

// GetValue 获取进度值 (0.0~1.0).
func (p *BeautifyProgressBar) GetValue() float32 {
	return p.value
}

// SetRadius 设置圆角半径.
func (p *BeautifyProgressBar) SetRadius(radius int32) {
	p.radius = radius
	xc.XEle_Redraw(p.Handle, false)
}

// GetRadius 获取圆角半径.
func (p *BeautifyProgressBar) GetRadius() int32 {
	return p.radius
}

// SetTrackColor 设置轨道背景颜色.
func (p *BeautifyProgressBar) SetTrackColor(color uint32) {
	p.trackColor = color
	xc.XEle_Redraw(p.Handle, false)
}

// GetTrackColor 获取轨道背景颜色.
func (p *BeautifyProgressBar) GetTrackColor() uint32 {
	return p.trackColor
}

// SetFillColor 设置进度填充颜色.
func (p *BeautifyProgressBar) SetFillColor(color uint32) {
	p.fillColor = color
	xc.XEle_Redraw(p.Handle, false)
}

// GetFillColor 获取进度填充颜色.
func (p *BeautifyProgressBar) GetFillColor() uint32 {
	return p.fillColor
}

// SetTextColor 设置百分比文字颜色.
func (p *BeautifyProgressBar) SetTextColor(color uint32) {
	p.textColor = color
	xc.XEle_Redraw(p.Handle, false)
}

// GetTextColor 获取百分比文字颜色.
func (p *BeautifyProgressBar) GetTextColor() uint32 {
	return p.textColor
}

// onPaint 自绘事件: 绘制圆角轨道 + 填充进度 + 百分比文字.
func (p *BeautifyProgressBar) onPaint(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true // 拦截默认绘制
	draw := drawx.NewByHandle(hDraw)
	if draw == nil {
		return 0
	}
	draw.EnableSmoothingMode(true) // 启用平滑模式

	var rc xc.RECT
	xc.XEle_GetClientRect(hEle, &rc)
	rw := rc.Right - rc.Left

	// 1. 轨道背景
	draw.SetBrushColor(p.trackColor)
	draw.FillRoundRect(&rc, p.radius, p.radius)

	// 2. 填充进度
	fillW := int32(float32(rw-8) * p.value)
	if fillW > p.radius {
		// 填充矩形右边缘是否到达右侧圆角区域, 到达则右边也画圆角
		rightRound := int32(0)
		if rc.Left+4+fillW >= rc.Right-p.radius {
			rightRound = p.radius - 4
		}
		draw.SetBrushColor(p.fillColor)
		draw.FillRoundRectEx(&xc.RECT{
			Left: rc.Left + 4, Top: rc.Top + 4,
			Right: rc.Left + 4 + fillW, Bottom: rc.Bottom - 4,
		}, p.radius-4, rightRound, rightRound, p.radius-4)
	}

	// 3. 百分比文字
	draw.SetFont(p.font)
	draw.SetBrushColor(p.textColor)
	draw.SetTextAlign(xcc.TextAlignFlag_Center | xcc.TextAlignFlag_Vcenter)
	draw.DrawText(fmt.Sprintf("%d%%", int(p.value*100+0.5)), &rc)
	return 0
}

func main() {
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	w := window.New(0, 0, 500, 320, "美化进度条示例", 0, xcc.Window_Style_Default)
	w.SetBorderSize(1, 30, 1, 1)

	fontPct := font.NewEX("微软雅黑", 12, xcc.FontStyle_Bold)

	// ====== 1. 默认蓝色进度条 (30%) ======
	bar1 := NewBeautifyProgressBar(30, 40, 440, 24, fontPct.Handle, w.Handle)
	bar1.SetValue(0.3)

	// ====== 2. 自定义颜色: 绿色填充 ======
	bar2 := NewBeautifyProgressBar(30, 80, 440, 24, fontPct.Handle, w.Handle)
	bar2.SetFillColor(xc.RGBA(16, 185, 129, 255))
	bar2.SetTrackColor(xc.RGBA(220, 245, 235, 255))
	bar2.SetTextColor(xc.RGBA(16, 185, 129, 255))
	bar2.SetValue(0.65)

	// ====== 3. 自定义颜色: 橙色进度 ======
	bar3 := NewBeautifyProgressBar(30, 120, 440, 20, fontPct.Handle, w.Handle)
	bar3.SetFillColor(xc.RGBA(255, 159, 67, 255))
	bar3.SetTrackColor(xc.RGBA(255, 240, 220, 255))
	bar3.SetTextColor(xc.RGBA(200, 100, 20, 255))
	bar3.SetRadius(10)
	bar3.SetValue(0.85)

	// ====== 4. 模拟下载进度 (动态增长) ======
	bar4 := NewBeautifyProgressBar(30, 170, 440, 24, fontPct.Handle, w.Handle)
	bar4.SetFillColor(xc.RGBA(99, 102, 241, 255))
	bar4.SetValue(0)

	// 注册进度改变事件
	bar4.AddEvent_ProgressBar_Change(func(hEle int, pos int32, pbHandled *bool) int {
		if pos >= 100 {
			w.MessageBox("提示", "模拟下载完成！", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Modal)
		}
		return 0
	})

	// 启动 goroutine 模拟动态进度
	go func() {
		time.Sleep(500 * time.Millisecond)
		v := float32(0)
		for v < 1 {
			time.Sleep(80 * time.Millisecond)
			v += 0.02
			xc.UI(func() {
				if !app.IsHELE(bar4.Handle) { // 防止还没执行完就关闭窗口导致元素被销毁了弹窗报错
					return
				}
				bar4.SetValue(v)
			})
		}
	}()

	w.Show(true)
	a.Run()
	a.Exit()
}
