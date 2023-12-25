// 多协程操作UI, 方式2.
// 必须在UI线程操作UI, 否则随机发生崩溃.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	a   *app.App
	w   *window.Window
	btn *widget.Button
	ls  *widget.List
)

func main() {
	// 置随机数种子
	rand.Seed(time.Now().UnixNano())

	a = app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 550, 300, "MultithreadOperationUI2", 0, xcc.Window_Style_Default)

	btn = widget.NewButton(15, 33, 70, 24, "click", w.Handle)
	btn.Event_BnClick(onBnClick)

	ls = widget.NewList(10, 60, 530, 230, w.Handle)
	ls.CreateAdapterHeader() // 创建表头数据适配器
	ls.CreateAdapter(5)      // 创建数据适配器, 5列

	ls.AddColumnText(100, "name1", "column1")
	ls.AddColumnText(100, "name2", "column2")
	ls.AddColumnText(100, "name3", "column3")
	ls.AddColumnText(100, "name4", "column4")
	ls.AddColumnText(100, "name5", "column5")

	for i := 1; i < 10; i++ {
		id := strconv.Itoa(i)
		index := ls.AddItemText("item" + id + "-column" + id)
		ls.SetItemText(index, 1, "item"+id+"-column"+id)
		ls.SetItemText(index, 2, "item"+id+"-column"+id)
		ls.SetItemText(index, 3, "item"+id+"-column"+id)
		ls.SetItemText(index, 4, "item"+id+"-column"+id)
	}

	a.ShowAndRun(w.Handle)
	a.Exit()
}

type updateList struct {
	Item int    // 项索引
	Col  int    // 列索引
	Text string // 项文本

	rwm sync.RWMutex // 保证同时只有1个在给List置入数据
	wg  sync.WaitGroup
}

// 在这里面写操作UI的代码, 是在ui线程操作ui
func (l *updateList) UiThreadCallBack(data int) int {
	xc.XList_SetItemText(data, l.Item, l.Col, l.Text)
	return 0
}

// 按钮单击事件
func onBnClick(pbHandled *bool) int {
	if !btn.IsEnable() {
		return 0
	}
	btn.Enable(false)
	btn.Redraw(false)

	go func() {
		u := new(updateList)
		t := time.Now() // 记录开始的时间

		// 多协程操作列表框数据
		for i := 0; i < 2022; i++ {
			u.wg.Add(1)

			go func() {
				u.rwm.RLock()
				u.Item = rand.Intn(ls.GetCount_AD())
				u.Col = rand.Intn(ls.GetColumnCount())
				u.Text = strconv.Itoa(rand.Intn(1000) + 1000)
				// 这种方式能够传递更多的数据进回调函数
				a.CallUiThreader(u, ls.Handle) // 这样是在UI线程进行UI操作, 就不会崩溃了
				u.rwm.RUnlock()

				u.wg.Done()
			}()
		}
		u.wg.Wait()

		// 如果不需要传参数进回调函数, 也不需要返回值时可以调用a.CallUT(), 回调函数写法能简单些.
		a.CallUT(func() {
			ls.RefreshData() // 刷新列表项数据
			ls.Redraw(false) // 列表重绘
			btn.Enable(true)
			btn.Redraw(true)
			w.MessageBox("提示", fmt.Sprintf("全部执行完毕, 耗时: %v", time.Since(t)), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
		})
	}()
	return 0
}
