// 调用UI设计器设计好的布局文件和资源文件
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	// 添加文件搜索路径, 请使用go run运行程序, 如果你使用go build运行, 那么请把这里改成`res`
	a.AddFileSearchPath(`uidesigner\res`)
	// 从zip中加载资源文件
	a.LoadResourceZip(`qqmusic.zip`, "resource.res", "")
	// 从zip中加载布局文件
	hWindow := a.LoadLayoutZip(`qqmusic.zip`, "main.xml", "", 0)
	if hWindow == 0 {
		panic("LoadLayoutZip Error")
	}

	// 创建窗口对象
	w := window.NewWindowByHandle(hWindow)
	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
