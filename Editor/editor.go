// 代码编辑框
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	panic("由于代码编辑框的API正在升级, 所以[代码编辑框的部分函数]会用不了, 等待xcgui原作者更新后将会开放大量接口, 比以前更好用")
	a := app.New(true)
	w := window.New(0, 0, 1000, 600, "Editor", 0, xcc.Window_Style_Default)

	// 创建Editor
	Editor := widget.NewEditor(12, 35, 975, 555, w.Handle)
	// 启用接收Tab输入
	Editor.EnableKeyTab(true)
	// 启用自动换行
	Editor.EnableAutoWrap(true)

	// 创建字体
	font1 := font.NewEX("Arial", 12, xcc.FontStyle_Regular)
	// 设置Editor的字体
	Editor.SetFont(font1.Handle)
	// 设置默认颜色
	Editor.SetTextColor(xc.ABGR(100, 100, 100, 255))

	// 添加样式
	iStyle_fun := Editor.AddStyle(0, xc.ABGR(255, 128, 0, 255), true)     // 函数
	iStyle_str := Editor.AddStyle(0, xc.ABGR(206, 145, 120, 255), true)   // 字符串
	iStyle_comment := Editor.AddStyle(0, xc.ABGR(67, 166, 74, 255), true) // 注释
	iStyle_key1 := Editor.AddStyle(0, xc.ABGR(86, 156, 214, 255), true)   // key1
	iStyle_key2 := Editor.AddStyle(0, xc.ABGR(200, 0, 0, 255), true)      // key2

	// 设置样式
	Editor.SetStyleFunction(iStyle_fun)
	Editor.SetStyleString(iStyle_str)
	Editor.SetStyleComment(iStyle_comment)
	// 添加关键字
	Editor.AddKeyword("if", iStyle_key1)
	Editor.AddKeyword("else", iStyle_key1)
	Editor.AddKeyword("switch", iStyle_key1)
	Editor.AddKeyword("case", iStyle_key1)
	Editor.AddKeyword("break", iStyle_key1)
	Editor.AddKeyword("int", iStyle_key1)

	Editor.AddKeyword("function", iStyle_key2)
	Editor.AddKeyword("return", iStyle_key2)

	// 添加自动匹配常量
	Editor.AddConst(`XE_BNCLICK //按钮点击事件`)
	Editor.AddConst(`XE_PAINT //元素绘制事件`)
	// 添加自动匹配函数
	Editor.AddFunction(`function Tmp1(pFileName string, hParent int) //我是 Tmp1`)
	Editor.AddFunction(`function Tmp2(pFileName string, hParent int) //我是 Tmp2`)
	Editor.AddFunction(`function Tmp3(pFileName string, hParent int) //我是 Tmp3`)

	// 设置断点
	Editor.SetBreakpoint(1, true)
	Editor.SetBreakpoint(2, true)
	Editor.SetBreakpoint(3, false)

	/* 	//获取设置的断点
	   	var BreakPoints []int32
	   	Editor.GetBreakpoints(&BreakPoints, Editor.GetBreakpointCount())
	   	for _, v := range BreakPoints {
	   		fmt.Println(v)
	   	} */

	// 设置当前运行行
	Editor.SetRunRow(0)

	code := `// 123456
function foo(a int,b int) int{
	Tmp1("layout.xml",0);
	Tmp2("layout.xml",0);
	Tmp3("layout.xml",0);
	XE_BNCLICK;
	XE_PAINT;
	if(a == 1){

	}else{

	}

	switch(a){
	case 0:
		break;
	case 1:
		break;
	}

	return 0
}`
	Editor.SetText(code)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
