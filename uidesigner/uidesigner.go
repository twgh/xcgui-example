// 调用UI设计器设计好的布局文件和资源文件
package main

import (
	_ "embed"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
)

//go:embed res/qqmusic.zip
var qqmusic []byte

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 从内存zip中加载资源文件
	a.LoadResourceZipMem(qqmusic, "resource.res", "")
	// 从内存zip中加载布局文件, 创建窗口对象
	w := window.NewByLayoutZipMem(qqmusic, "main.xml", "", 0, 0)

	// songTitle是在main.xml中给歌曲名(shapeText组件)设置的name属性的值.
	// 通过 GetObjectByName 可以获取布局文件中设置了name属性的组件的句柄.
	// 可简化为: widget.NewShapeTextByName("songTitle").
	song := widget.NewShapeTextByHandle(a.GetObjectByName("songTitle"))
	println(song.GetText()) // 输出: 两只老虎爱跳舞

	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}
