// 侧边导航
package main

import (
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	w *window.Window
)

func main() {
	// 初始化UI库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w = window.New(0, 0, 800, 600, "侧边导航示例", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	w.SetBorderSize(2, 32, 2, 2)
	w.EnableLayout(true)     // 启用布局
	w.EnableHorizon(false)   // 禁用水平布局
	w.SetPadding(6, 6, 6, 6) // 设置窗口内填充

	// 创建树形框作为侧边导航
	createTree()

	// 创建主内容区布局元素
	layoutContent := widget.NewLayoutEle(0, 0, 0, 0, w.Handle)
	layoutContent.SetPadding(4, 4, 4, 4)
	// 宽度占用剩余空间
	layoutContent.LayoutItem_SetWidth(xcc.Layout_Size_Weight, 1)
	// 高度填充父
	layoutContent.LayoutItem_SetHeight(xcc.Layout_Size_Fill, 0)

	// 后续:
	// 1.创建所有页面的布局元素到主内容区布局元素里, 只把首页的show(true), 其他的show(false)
	// 2.点击导航项时, show对应页面的布局元素, 隐藏其他页面的布局元素, 然后主内容区布局元素.调整布局, 主内容区布局元素.刷新

	w.Show(true)
	a.Run()
}

// 创建树形框作为侧边导航
func createTree() {
	tree := widget.NewTree(0, 0, 140, 500, w.Handle)
	tree.LayoutItem_SetHeight(xcc.Layout_Size_Fill, 0) // 高度填充父

	// 创建数据适配器（必须步骤）
	tree.CreateAdapter()

	// 设置树形框样式
	tree.EnableConnectLine(false, false) // 禁用连接线
	tree.SetItemHeightDefault(36, 36)    // 设置项高度

	// 禁用边框和焦点绘制（更简洁的外观）
	tree.EnableDrawBorder(false)
	tree.EnableDrawFocus(false)

	// 添加导航项
	addNavigationItems(tree)

	// 加载图标
	imgTop := app.NewImageBySvgString(svg_top)
	imgBottom := app.NewImageBySvgString(svg_bottom)

	// 创建背景管理器
	bkmTop := app.NewBkManager().AddRef()
	bkmTop.AddImage(xcc.Button_State_Flag_Leave, imgTop.Handle, 0)
	bkmTop.AddImage(xcc.Button_State_Flag_Stay, imgTop.Handle, 0)
	bkmTop.AddImage(xcc.Button_State_Flag_Down, imgTop.Handle, 0)
	bkmBottom := app.NewBkManager().AddRef()
	bkmBottom.AddImage(xcc.Button_State_Flag_Leave, imgBottom.Handle, 0)
	bkmBottom.AddImage(xcc.Button_State_Flag_Stay, imgBottom.Handle, 0)
	bkmBottom.AddImage(xcc.Button_State_Flag_Down, imgBottom.Handle, 0)

	// 窗口销毁时释放
	w.AddEvent_Destroy(func(hWindow int, pbHandled *bool) int {
		bkmTop.Release()
		bkmBottom.Release()
		return 0
	})

	// 列表树元素-项模板创建完成事件
	tree.AddEvent_Tree_Temp_Create_End(func(hEle int, pItem *xc.Tree_Item_, nFlag int32, pbHandled *bool) int {
		if nFlag == 1 { // 1: 新模板实例
			if pItem.NDepth == 0 { // 深度 0 也就是顶级菜单
				// 获取展开按钮句柄
				// 在默认树元素项模板文件中, 展开按钮的 itemID 是 1
				hBtn := tree.GetTemplateObject(pItem.NID, 1)
				btn := widget.NewButtonByHandle(hBtn)

				// 设置展开/收缩图标
				if pItem.BExpand {
					btn.SetBkManager(bkmTop.Handle)
				} else {
					btn.SetBkManager(bkmBottom.Handle)
				}

				btn.Redraw(false)
			}
		}
		return 0
	})

	// 树元素-项展开收缩事件
	tree.AddEvent_Tree_Expand(func(hEle int, id int32, bExpand bool, pbHandled *bool) int {
		// 获取展开按钮句柄
		// 在默认树元素项模板文件中, 展开按钮的 itemID 是 1
		hBtn := tree.GetTemplateObject(id, 1)
		if xc.XC_GetObjectType(hBtn) != xcc.XC_BUTTON {
			return 0
		}
		btn := widget.NewButtonByHandle(hBtn)

		// 设置展开/收缩图标
		if bExpand {
			btn.SetBkManager(bkmTop.Handle)
		} else {
			btn.SetBkManager(bkmBottom.Handle)
		}
		btn.Redraw(false)
		return 0
	})

	// 添加树元素-项选择事件
	tree.AddEvent_Tree_Select(func(hEle int, nItemID int32, pbHandled *bool) int {
		itemText := tree.GetItemText(nItemID, 0) // 获取选中项的文本

		// 根据不同的导航项执行相应操作
		switch itemText {
		case "系统设置":
			fmt.Println("打开系统设置页面")
		case "用户列表":
			fmt.Println("打开用户列表页面")
		case "数据备份":
			fmt.Println("打开数据备份页面")

		// 添加更多case处理其他导航项
		default:
			fmt.Println("选中了:", itemText)
		}
		return 0
	})

	// 树形框鼠标左键弹起事件
	tree.AddEvent_LButtonUp(func(hEle, nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		// 获取选中项ID
		nItemID := tree.GetSelectItem()
		// 获取选中项的文本
		itemText := tree.GetItemText(nItemID, 0)

		switch itemText {
		case "系统管理", "用户管理", "数据管理":
			IsExpand := tree.IsExpand(nItemID)
			fmt.Println(common.Choose(IsExpand, "折叠", "展开"), itemText)

			Expand := !IsExpand
			// 展开或折叠项
			tree.ExpandItem(nItemID, Expand)
			// 触发树元素-项展开收缩事件
			tree.SendEvent(xcc.XE_TREE_EXPAND, uintptr(nItemID), uintptr(common.Choose(Expand, 1, 0)))
			tree.Redraw(false)
		}
		return 0
	})
}

// 添加导航项
func addNavigationItems(tree *widget.Tree) {
	// 首页
	tree.InsertItemText("首页", xcc.XC_ID_ROOT, xcc.XC_ID_LAST)

	// 添加一级菜单项
	systemIndex := tree.InsertItemText("系统管理", xcc.XC_ID_ROOT, xcc.XC_ID_LAST)
	userIndex := tree.InsertItemText("用户管理", xcc.XC_ID_ROOT, xcc.XC_ID_LAST)
	dataIndex := tree.InsertItemText("数据管理", xcc.XC_ID_ROOT, xcc.XC_ID_LAST)

	// 添加子菜单项
	// 系统管理子项
	tree.InsertItemText("系统设置", systemIndex, xcc.XC_ID_LAST)
	tree.InsertItemText("日志管理", systemIndex, xcc.XC_ID_LAST)

	// 用户管理子项
	tree.InsertItemText("用户列表", userIndex, xcc.XC_ID_LAST)
	tree.InsertItemText("角色管理", userIndex, xcc.XC_ID_LAST)
	tree.InsertItemText("权限设置", userIndex, xcc.XC_ID_LAST)

	// 数据管理子项
	tree.InsertItemText("数据备份", dataIndex, xcc.XC_ID_LAST)
	tree.InsertItemText("数据恢复", dataIndex, xcc.XC_ID_LAST)
}

const (
	svg_top    = `<svg t="1760687396783" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4603" width="20" height="20"><path d="M257.28 616.192l46.336 46.336L512 454.093l208.384 208.435 46.336-46.336-208.384-208.384L512 361.472l-46.336 46.336z" fill="#000000" p-id="4604"></path></svg>`
	svg_bottom = `<svg t="1760687462031" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5197" width="20" height="20"><path d="M766.72 407.808l-46.336-46.336L512 569.907 303.616 361.472l-46.336 46.336 208.384 208.384L512 662.528l46.336-46.336z" fill="#000000" p-id="5198"></path></svg>`
)
