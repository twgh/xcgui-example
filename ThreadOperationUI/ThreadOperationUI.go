// 炫彩_调用界面线程, 在UI线程操作UI, 不在UI线程操作UI, 必将崩溃.
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/twgh/xcgui/wapi"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	w   *window.Window
	btn *widget.Button
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	fmt.Println("UI线程id:", xc.ThreadId)

	// 启用自适应 DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w = window.New(0, 0, 550, 300, "ThreadOperationUI", 0, xcc.Window_Style_Default)

	// 创建按钮
	btn = widget.NewButton(20, 35, 100, 30, "click", w.Handle)

	// 单选按钮
	radioBtn1 := widget.NewButton(20, 90, 70, 30, "方式1", w.Handle)
	radioBtn2 := widget.NewButton(100, 90, 70, 30, "方式2", w.Handle)
	radioBtn3 := widget.NewButton(180, 90, 70, 30, "方式3", w.Handle)
	// 设置按钮类型为单选按钮
	radioBtn1.SetTypeEx(xcc.Button_Type_Radio)
	radioBtn2.SetTypeEx(xcc.Button_Type_Radio)
	radioBtn3.SetTypeEx(xcc.Button_Type_Radio)
	// 默认选中radioBtn1
	radioBtn1.SetCheck(true)

	// 添加按钮点击事件.
	btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		// 事件回调函数是在UI线程内执行的.
		fmt.Printf("事件回调函数所在线程id: %d, 是否UI线程: %v\n", wapi.GetCurrentThreadId(), xc.IsUiThread())

		// 禁用按钮
		btn.Enable(false).Redraw(false)

		// 另起协程/线程, 就不是在UI线程内了.
		// 因为是在事件回调函数里执行耗时操作, 所以必须另起协程/线程的,
		// 不然界面会卡住, 等到操作完才会恢复. 如果你知道窗口消息循环机制的话,
		// 就会懂的, 事件回调函数是在消息循环里执行的.
		if radioBtn1.IsCheck() {
			go updateBtn1() // 第一种方式: app.CallUiThreadEx, 可传1个整数
		} else if radioBtn2.IsCheck() {
			go updateBtn2() // 第二种方式: app.CallUiThreader, 可传更多的参数
		} else if radioBtn3.IsCheck() {
			go updateBtn3() // 第三种方式: app.CallUTAny, 可传任意类型参数, 且数量可变.
		}
		return 0
	})

	a.ShowAndRun(w.Handle)
	a.Exit()
}

// 第一种方式
func updateBtn1() {
	fmt.Println("使用方式1: CallUiThreadEx")
	fmt.Printf("协程所在线程id: %d, 是否UI线程: %v\n", wapi.GetCurrentThreadId(), xc.IsUiThread())
	for i := 0; i < 2010; i += 10 {
		// 模拟处理数据
		btnText := strconv.Itoa(i)
		width := i / 5
		time.Sleep(time.Millisecond * 3) // 加延迟模拟处理数据的耗时

		// 如果直接在非UI线程内操作UI, 次数多了程序必将崩溃, 而且你不会知道它在什么时候崩溃.
		// 使用 app.CallUiThreadEx() 这样是在UI线程进行UI操作, 就不会崩溃了.
		app.CallUiThreadEx(func(data int) int {
			btn.SetText(btnText)
			btn.SetWidth(int32(data))
			w.Redraw(false)
			return 0
		}, width) // 把 width 传进回调函数了. 其实直接闭包就行, 这里是测试传参.
	}

	// 如果不需要传参数进回调函数, 也不需要返回值时可以调用 app.CallUT(), 回调函数写法能简单些.
	app.CallUT(func() {
		btn.Enable(true).Redraw(false) // 解禁按钮.
	})
	fmt.Println("----------------------------------------------------")
}

// 第2种方式, 这种方式明显可以传更多的参数, 完成更复杂的操作
type updateButton struct {
	HEle         int
	Text         string
	Width        int32
	RedrawWindow bool
}

func (u *updateButton) UiThreadCallBack(data int) int {
	xc.XBtn_SetText(u.HEle, u.Text)
	xc.XEle_SetWidth(u.HEle, u.Width)
	w.Redraw(u.RedrawWindow)
	return 0
}

func updateBtn2() {
	fmt.Println("使用方式2: CallUiThreader")
	fmt.Printf("协程所在线程id: %d, 是否UI线程: %v\n", wapi.GetCurrentThreadId(), xc.IsUiThread())
	u := &updateButton{
		HEle:         btn.Handle,
		RedrawWindow: false,
	}

	for i := int32(0); i < 2010; i += 10 {
		// 模拟处理数据
		u.Text = xc.Itoa(i)
		u.Width = i / 5
		time.Sleep(time.Millisecond * 3) // 加延迟模拟处理数据的耗时

		// 如果直接在非界面线程内操作UI, 次数多了程序必将崩溃, 而且你不会知道它在什么时候崩溃.
		// 使用 app.CallUiThreader 这样是在界面线程进行UI操作, 就不会崩溃了.
		app.CallUiThreader(u, 0)
	}

	// 如果不需要传参数进回调函数, 也不需要返回值时可以调用 app.CallUT(), 回调函数写法能简单些.
	app.CallUT(func() {
		btn.Enable(true).Redraw(false) // 解禁按钮.
	})
	fmt.Println("----------------------------------------------------")
}

// 第3种方式, 可传任意类型参数, 且数量可变
func updateBtn3() {
	fmt.Println("使用方式3: CallUTAny")
	fmt.Printf("协程所在线程id: %d, 是否UI线程: %v\n", wapi.GetCurrentThreadId(), xc.IsUiThread())
	for i := 0; i < 2010; i += 10 {
		// 模拟处理数据
		btnText := strconv.Itoa(i)
		width := i / 5
		time.Sleep(time.Millisecond * 3) // 加延迟模拟处理数据的耗时

		// 如果直接在非UI线程内操作UI, 次数多了程序必将崩溃, 而且你不会知道它在什么时候崩溃.
		// 使用 app.CallUTAny() 这样是在UI线程进行UI操作, 就不会崩溃了.
		app.CallUTAny(func(data ...interface{}) int {
			btn.SetText(data[0].(string))
			btn.SetWidth(int32(data[1].(int)))
			w.Redraw(false)
			return 0
		}, btnText, width) // 把 btnText, width 传进回调函数了. 其实直接闭包就行, 这里是测试传参.
	}

	// 如果不需要传参数进回调函数, 也不需要返回值时可以调用 app.CallUT(), 回调函数写法能简单些.
	app.CallUT(func() {
		btn.Enable(true).Redraw(false) // 解禁按钮.
	})
	fmt.Println("----------------------------------------------------")
}

// 第4种方式, 使用 xc.Auto 系列函数, 会自动调用系统 API getCurrentThreadId 来判断是否在 UI 线程, 如果不在, 则自动调用界面线程来执行回调函数, 由于多了一步调用系统 API, 肯定会增加耗时的.
