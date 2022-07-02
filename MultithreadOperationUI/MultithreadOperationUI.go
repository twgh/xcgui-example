// 多线程操作UI, 方式1.
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

	rwm sync.RWMutex
	wg  sync.WaitGroup
	t   time.Time
)

func main() {
	// 置随机数种子
	rand.Seed(time.Now().UnixNano())

	a = app.New(true)
	w = window.New(0, 0, 550, 300, "MultithreadOperationUI", 0, xcc.Window_Style_Default)
	// 创建按钮
	btn = widget.NewButton(15, 33, 70, 24, "click", w.Handle)
	btn.Event_BnClick(onBnClick)
	// 创建列表
	ls = widget.NewList(10, 60, 530, 230, w.Handle)
	ls.CreateAdapterHeader() // 创建表头数据适配器
	ls.CreateAdapter(5)      // 创建数据适配器, 5列
	// 添加列
	ls.AddColumnText(100, "name1", "column1")
	ls.AddColumnText(100, "name2", "column2")
	ls.AddColumnText(100, "name3", "column3")
	ls.AddColumnText(100, "name4", "column4")
	ls.AddColumnText(100, "name5", "column5")
	// 置列表数据
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

// 按钮单击事件
func onBnClick(pbHandled *bool) int {
	if !btn.IsEnable() {
		return 0
	}
	btn.Enable(false)
	btn.Redraw(true)

	go func() {
		t = time.Now() // 记录开始的时间

		// 多线程操作列表框数据
		for i := 0; i < 2022; i++ {
			wg.Add(1)

			// go setText(0) // 像这样直接另起线程操作UI次数多了程序必将崩溃, 你可以测试一下

			go func() {
				xc.XC_CallUiThreadEx(setText, 0) // 这样是在界面线程进行UI操作, 就不会崩溃了
				wg.Done()
			}()
		}
		wg.Wait()

		xc.XC_CallUiThreadEx(func(data int) int {
			ls.RefreshData() // 刷新列表项数据
			ls.Redraw(false) // 列表重绘
			btn.Enable(true)
			btn.Redraw(true)
			w.MessageBox("提示", fmt.Sprintf("全部执行完毕, 耗时: %v", time.Since(t)), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}, 0)
	}()
	return 0
}

func setText(data int) int {
	item := rand.Intn(ls.GetCount_AD())
	col := rand.Intn(ls.GetColumnCount())
	text := strconv.Itoa(rand.Intn(1000) + 1000)

	rwm.RLock()
	ls.SetItemText(item, col, text)
	rwm.RUnlock()
	return 0
}
