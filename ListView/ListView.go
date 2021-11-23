// 列表视图
package main

import (
	_ "embed"
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed res/1.png
var img1 []byte

func main() {
	a := app.New(true)
	w := window.NewWindow(0, 0, 465, 400, "ListView", 0, xcc.Window_Style_Default)

	// 创建ListView
	lv := widget.NewListView(10, 32, 445, 357, w.Handle)
	// 创建数据适配器
	lv.CreateAdapter()

	// 添加分组
	group1 := lv.Group_AddItemText("group1", -1)
	group2 := lv.Group_AddItemText("group2", -1)
	// 图片加载从内存
	img := imagex.NewImage_LoadMemory(&img1, len(img1))

	// 循环把图片加到分组里
	var index int
	for i := 0; i < 3; i++ {
		index = lv.Item_AddItemImage(group1, img.Handle, -1)
		lv.Item_SetText(group1, index, 1, fmt.Sprintf("group1-item%d", i))

		index = lv.Item_AddItemImage(group2, img.Handle, -1)
		lv.Item_SetText(group2, index, 1, fmt.Sprintf("group2-item%d", i))
	}

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
