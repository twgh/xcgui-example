// 列表视图
package main

import (
	_ "embed"
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//go:embed res/1.png
var img1 []byte

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w := window.New(0, 0, 465, 400, "ListView", 0, xcc.Window_Style_Default)

	// 创建ListView
	lv := widget.NewListView(10, 32, 445, 357, w.Handle)
	// 创建数据适配器
	lv.CreateAdapter()

	// 添加分组
	group1 := lv.Group_AddItemText("group1", -1)
	group2 := lv.Group_AddItemText("group2", -1)
	// 图片加载从内存
	img := imagex.NewByMem(img1)

	// 循环把图片加到分组里
	for i := 0; i < 3; i++ {
		// 列表视_项添加图片
		index := lv.Item_AddItemImage(group1, img.Handle, -1)
		lv.Item_SetText(group1, index, 1, fmt.Sprintf("group1-item%d", i))

		index = lv.Item_AddItemImage(group2, img.Handle, -1)
		lv.Item_SetText(group2, index, 1, fmt.Sprintf("group2-item%d", i))
	}

	widget.NewButton(150, 0, 70, 30, "取选中项", w.Handle).Event_BnClick(func(pbHandled *bool) int {
		n := lv.GetSelectItemCount()
		fmt.Println("选中项个数:", n)
		if n == 0 {
			return 0
		}

		// 取选中项id
		var ids []xc.ListView_Item_Id_
		lv.GetSelectAll(&ids, n)
		fmt.Println("选中的列表视-项ID:", ids)

		var groupIndex, itemIndex int32
		lv.GetSelectItem(&groupIndex, &itemIndex)
		fmt.Printf("选中项组索引: %d, 项索引: %d\n", groupIndex, itemIndex)
		return 0
	})

	// 添加列表视元素-项选择事件.
	lv.AddEvent_ListView_Select(func(hEle int, iGroup int32, iItem int32, pbHandled *bool) int {
		fmt.Println("---------- 列表视元素-项选择事件 ----------")
		fmt.Printf("选中项组索引: %d, 项索引: %d\n", iGroup, iItem)
		return 0
	})

	// 添加列表视元素-组展开收缩事件.
	lv.AddEvent_ListView_Expand(func(hEle int, iGroup int32, bExpand bool, pbHandled *bool) int {
		fmt.Println("----------列表视元素-组展开收缩事件 ----------")
		fmt.Printf("选中项组索引: %d, 展开状态: %v\n", iGroup, bExpand)
		return 0
	})

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
