// 调用 wapi 打开/保存文件, 浏览文件夹
package main

import (
	"fmt"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"strings"
	"syscall"
	"unsafe"
)

var (
	a *app.App
	w *window.Window

	btn1 *widget.Button
	btn2 *widget.Button
	btn3 *widget.Button
	btn4 *widget.Button
)

func main() {
	a = app.New(true)
	w = window.New(0, 0, 430, 300, "", 0, xcc.Window_Style_Default)

	// 创建按钮
	btn1 = widget.NewButton(20, 40, 100, 30, "浏览文件夹", w.Handle)
	btn2 = widget.NewButton(20, 80, 100, 30, "单选打开文件", w.Handle)
	btn3 = widget.NewButton(130, 80, 100, 30, "多选打开文件", w.Handle)
	btn4 = widget.NewButton(20, 120, 100, 30, "保存文件", w.Handle)

	// 注册按钮事件
	btn1.Event_BnClick1(onBnClick)
	btn2.Event_BnClick1(onBnClick)
	btn3.Event_BnClick1(onBnClick)
	btn4.Event_BnClick1(onBnClick)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func onBnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn1.Handle:
		ExampleSHGetPathFromIDListW()
	case btn2.Handle:
		ExampleGetOpenFileNameW()
	case btn3.Handle:
		ExampleGetOpenFileNameW_2()
	case btn4.Handle:
		ExampleGetSaveFileNameW()
	}
	return 0
}

// 浏览文件夹
func ExampleSHGetPathFromIDListW() {
	buf := make([]uint16, 260)
	bi := wapi.BrowseInfoW{
		HwndOwner:      0,
		PidlRoot:       0,
		PszDisplayName: common.Uint16SliceDataPtr(&buf),
		LpszTitle:      common.StrPtr("显示在对话框中树视图控件上方的文本"),
		UlFlags:        wapi.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}
	var pszPath string
	ret := wapi.SHGetPathFromIDListW(wapi.SHBrowseForFolderW(&bi), &pszPath)
	fmt.Println(ret)
	fmt.Println("pszPath:", pszPath)                           // 用户选择的文件夹完整路径
	fmt.Println("PszDisplayName:", syscall.UTF16ToString(buf)) // 用户选择的文件夹的名称
}

// 打开单个文件
func ExampleGetOpenFileNameW() {
	// 多个过滤器, 打开单个文件.
	c := "\x00"
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, c) + c + c

	lpstrFile := make([]uint16, 260)
	lpstrFileTitle := make([]uint16, 260)

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         w.GetHWND(),
		HInstance:         0,
		LpstrFilter:       common.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    &lpstrFileTitle[0],
		NMaxFileTitle:     260,
		LpstrInitialDir:   common.StrPtr("D:"),
		LpstrTitle:        common.StrPtr("打开文件"),
		Flags:             wapi.OFN_PATHMUTEXIST, // 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := wapi.GetOpenFileNameW(&ofn)
	fmt.Println(ret)
	fmt.Println("lpstrFile:", syscall.UTF16ToString(lpstrFile))
	fmt.Println("lpstrFileTitle:", syscall.UTF16ToString(lpstrFileTitle))
}

// 打开多个文件
func ExampleGetOpenFileNameW_2() {
	// 多个过滤器, 打开多个文件.
	c := "\x00"
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, c) + c + c

	lpstrFile := make([]uint16, 512)

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         w.GetHWND(),
		HInstance:         0,
		LpstrFilter:       common.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      2,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          512,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   common.StrPtr("D:"),
		LpstrTitle:        common.StrPtr("打开文件(可选多个)"),
		Flags:             wapi.OFN_ALLOWMULTISELECT | wapi.OFN_EXPLORER | wapi.OFN_PATHMUTEXIST, // 允许文件多选 | 使用新界面 | 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := wapi.GetOpenFileNameW(&ofn)
	fmt.Println(ret)

	s := common.Uint16SliceToStringSlice(lpstrFile)
	fmt.Println("选择的文件个数:", len(s)-1) // -1是因为切片中第一个元素是文件目录, 不是文件
	fmt.Println("lpstrFile:", s)
}

// 保存文件
func ExampleGetSaveFileNameW() {
	// 多个过滤器, 保存文件.
	c := "\x00"
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, c) + c + c

	lpstrFile := make([]uint16, 260)
	lpstrFileTitle := make([]uint16, 260)

	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         w.GetHWND(),
		HInstance:         0,
		LpstrFilter:       common.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    &lpstrFileTitle[0],
		NMaxFileTitle:     260,
		LpstrInitialDir:   common.StrPtr("D:"),
		LpstrTitle:        common.StrPtr("保存文件"),
		Flags:             wapi.OFN_OVERWRITEPROMPT, // 如果所选文件已存在，则使“另存为”对话框生成一个消息框。用户必须确认是否覆盖文件。
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       common.StrPtr("txt"), // 如果用户没有输入文件扩展名, 则默认使用这个
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := wapi.GetSaveFileNameW(&ofn)
	fmt.Println(ret)
	fmt.Println("lpstrFile:", syscall.UTF16ToString(lpstrFile))
	fmt.Println("lpstrFileTitle:", syscall.UTF16ToString(lpstrFileTitle))
}
