// 调用UI设计器设计好的布局文件和资源文件
package main

import (
	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed res/qqmusic.zip
var qqmusic []byte

func main() {
	a := app.New(true)
	// 从内存zip中加载资源文件
	a.LoadResourceZipMem(qqmusic, "resource.res", "")
	// 从内存zip中加载布局文件
	hWindow := a.LoadLayoutZipMem(qqmusic, "main.xml", "", 0, 0)
	// 创建窗口对象
	w := window.NewWindowByHandle(hWindow)
	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
