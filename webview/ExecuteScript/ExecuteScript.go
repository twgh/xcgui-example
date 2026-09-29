// 执行JS
package main

// 演示 Go 侧执行页面 JavaScript 并取回返回值,
// 以及通过 AddScriptToExecuteOnDocumentCreated 在每次页面加载前注入脚本(可移除).

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//go:embed assets/**
var embedAssets embed.FS // 嵌入 assets 目录以及子目录下的文件

const hostName = "app.example"

// injectScript 注入到每个新文档的脚本: 每次加载页面(包括导航到新页面)都会执行,
// 它给 window 加个标记并在页面顶部插入标记条, 方便直观确认注入是否生效.
const injectScript = `window.__xcguiInjected = '由 Go 侧注入的脚本';
window.addEventListener('DOMContentLoaded', function () {
	var bar = document.createElement('div');
	bar.style.cssText = 'padding:8px 16px;background:#e8f4ff;border-bottom:1px solid #b8dcff;color:#1a5fb4;font-size:13px;';
	bar.textContent = '[注入脚本生效] ' + window.__xcguiInjected + ' (AddScriptToExecuteOnDocumentCreated)';
	document.body.insertBefore(bar, document.body.firstChild);
});`

// computeJS 执行一段 JS 计算(日期/随机数/Math), 把结果序列化为 JSON 字符串返回给 Go 侧.
const computeJS = `JSON.stringify({time: new Date().toLocaleString('zh-CN'), random: Math.floor(Math.random() * 9000) + 1000, pow2_10: Math.pow(2, 10), sqrt2: Math.sqrt(2).toFixed(6)})`

var (
	w          *window.Window
	wv         *edge.WebView
	resultEdit *widget.Edit // 炫彩侧编辑框, 显示最近一次从页面取回的值
	scriptID   string       // 注入脚本的 id, 用于 RemoveScriptToExecuteOnDocumentCreated
)

func main() {
	checkWebView2()
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)

	newWebView(edg)       // 创建窗口和 WebView 并导航到首页
	addInjectScript()     // 在每次页面加载前注入脚本
	createGoSideWidgets() // 创建炫彩侧的编辑框和按钮

	// 显示窗口并运行应用
	w.Show(true)
	a.Run()
	a.Exit()
}

// newWebView 创建窗口和 WebView, 设置虚拟主机映射到嵌入文件系统后导航到首页.
// 注: 本库没有同步执行 JS 取结果的方式, ICoreWebView2_21.ExecuteScriptWithResult 也是异步的.
func newWebView(edg *edge.Edge) {
	var err error
	w, wv, err = edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("ExecuteScript 例子 - Go 侧执行 JS 并取回结果"),
		edge.WithXmlWindowClassName("ExecuteScript"),
		edge.WithXmlWindowSize(1100, 780),         // 窗口大小
		edge.WithXmlWindowTitleBar(true),          // 使用炫彩窗口标题栏
		edge.WithWebViewSize(0, 0, 1100-8*2, 648), // WebView 固定位置: 标题栏下方, 底部留出炫彩控件区域, 宽度减去左右阴影
		edge.WithDebug(true),                      // 可打开开发者工具
		edge.WithAutoFocus(true),                  // 在窗口获得焦点时尝试保持 WebView 的焦点
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 设置窗口背景色.
	// NewWebViewWithWindow 创建的是一个透明窗口, 它的窗口背景是一个透明度为 1(几乎全透明)的填充,
	// 所以 WebView 和炫彩控件没盖住的地方会直接透出桌面.
	// 改成不透明需要注意两点:
	//   1. 必须先用 bkm.Clear() 清空窗口已有的背景, 因为 bkm.AddFill() 是追加式的,
	//      不清空的话 xml 里那个透明填充还在, 新加的填充看不到效果(下面用的就是窗口自己的背景管理器).
	//   2. 再给窗口主体添加不透明的填充.
	bkm := w.GetBkManagerObj()
	bkm.Clear()
	bkm.AddFill(xcc.CombinedState_Window_State_Flag_Body_Leave, xc.RGBA(0xf6, 0xf8, 0xfa, 255), 0) // 窗口主体背景色, 和网页背景色一致

	// 设置并启用虚拟主机名和嵌入文件系统之间的映射
	if err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets); err != nil {
		wapi.MessageBoxW(0, "SetVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(3)
	}
	if err = wv.EnableVirtualHostNameToEmbedFSMapping(true); err != nil {
		wapi.MessageBoxW(0, "EnableVirtualHostNameToEmbedFSMapping 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(4)
	}
	// 导航到首页
	_ = wv.Navigate(edge.JoinUrlHeader(hostName) + "/index.html")
}

// addInjectScript 通过 AddScriptToExecuteOnDocumentCreated 注入脚本,
// 之后每次创建新文档(刷新或导航到新页面)时都会先执行这段脚本.
func addInjectScript() {
	err := wv.AddScriptToExecuteOnDocumentCreated(injectScript, func(errorCode syscall.Errno, id string) uintptr {
		if wapi.IsOK(errorCode) {
			scriptID = id // 保存 id, 之后可用 RemoveScriptToExecuteOnDocumentCreated 移除
			fmt.Println("注入脚本已添加, id:", id)
		}
		return 0
	})
	logErr("AddScriptToExecuteOnDocumentCreated", err)
}

// createGoSideWidgets 在窗口底部创建炫彩编辑框和按钮, 按钮点击事件在 UI 线程触发, 可直接调用 WebView API.
func createGoSideWidgets() {
	// 炫彩侧编辑框: 显示最近一次从页面取回的值
	resultEdit = widget.NewEdit(10, 692, 1080, 32, w.Handle)
	resultEdit.SetText("这里显示最近一次从页面取回的值")

	// 按钮1: ExecuteScript 读取页面输入框的值(异步执行, JS 返回值通过回调取回)
	btnRead := widget.NewButton(10, 732, 300, 36, "读取页面输入框的值", w.Handle)
	btnRead.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		err := wv.ExecuteScript("document.getElementById('userInput').value", func(errorCode syscall.Errno, result string) uintptr {
			if wapi.IsOK(errorCode) {
				// 返回值是 JSON 编码的, 字符串结果带引号, 先用 unwrapJSON 还原
				text := "页面输入框的值: " + unwrapJSON(result)
				writeToPage(text)   // 写回页面显示
				setGoSideText(text) // 更新炫彩侧显示
			}
			return 0
		})
		logErr("ExecuteScript", err)
		return 0
	})

	// 按钮2: ExecuteScript 执行 JS 计算并取回 JSON 结果(对象类型的返回值本身就是 JSON 文本)
	btnCalc := widget.NewButton(330, 732, 320, 36, "执行 JS 计算并取回 JSON", w.Handle)
	btnCalc.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		err := wv.ExecuteScript(computeJS, func(errorCode syscall.Errno, result string) uintptr {
			if wapi.IsOK(errorCode) {
				text := "JS 计算结果: " + result
				writeToPage(text)
				setGoSideText(text)
			}
			return 0
		})
		logErr("ExecuteScript", err)
		return 0
	})

	// 按钮3: 移除注入的脚本, 然后导航到新页面, 新页面顶部将不再出现注入脚本生成的标记条
	btnRemove := widget.NewButton(670, 732, 410, 36, "移除注入脚本并跳转新页面", w.Handle)
	btnRemove.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if scriptID == "" {
			w.MessageBox("提示", "注入脚本尚未添加或已被移除.", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Modal)
			return 0
		}
		if err := wv.RemoveScriptToExecuteOnDocumentCreated(scriptID); err != nil {
			log.Println("RemoveScriptToExecuteOnDocumentCreated 失败:", err)
			return 0
		}
		scriptID = ""
		fmt.Println("注入脚本已移除, 跳转到新页面验证效果")
		_ = wv.Navigate(edge.JoinUrlHeader(hostName) + "/page2.html")
		return 0
	})
}

// writeToPage 通过 ExecuteScript 把一条结果追加显示到页面的 goResult 区域.
// 用 json.Marshal 把 Go 字符串转成合法的 JS 字符串字面量, 避免拼接时引号转义出错.
func writeToPage(text string) {
	b, err := json.Marshal(text)
	if err != nil {
		return
	}
	script := "(function(){var box=document.getElementById('goResult');if(!box){return;}" +
		"var line=document.createElement('div');" +
		"line.textContent='[Go] '+new Date().toLocaleTimeString('zh-CN')+'  '+(" + string(b) + ");" +
		"box.appendChild(line);})();"
	_ = wv.ExecuteScript(script, nil) // 不需要取回结果, 回调传 nil
}

// setGoSideText 更新炫彩侧编辑框文本, 用 xc.UI 包裹以规范地在 UI 线程操作炫彩 UI.
func setGoSideText(text string) {
	xc.UI(func() {
		resultEdit.SetText(text)
	})
}

// unwrapJSON 把 ExecuteScript 返回的 JSON 编码结果还原为字符串:
// 字符串结果带双引号, 对象等结果本身就是 JSON 文本, 原样返回.
func unwrapJSON(result string) string {
	var s string
	if err := json.Unmarshal([]byte(result), &s); err == nil {
		return s
	}
	return result
}

func logErr(op string, err error) {
	if err != nil {
		log.Println(op+" 调用失败:", err)
	}
}

// 创建 WebView 环境
func createEdge() *edge.Edge {
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 实际应用中应使用自己创建的固定目录
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}
	return edg
}

func checkWebView2() {
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())
	localVersion, err := edge.GetAvailableBrowserVersion()
	if err != nil {
		wapi.MessageBoxW(0, "获取 WebView2 运行时版本号失败: "+err.Error(), "提示", wapi.MB_IconError)
		os.Exit(1)
	}
	if localVersion == "" {
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序! 下载完后请使用管理员权限运行安装包!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		// 打开 WebView2 运行时小型引导程序下载地址
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)
	// 判断本机 WebView2 运行时版本是否低于本库使用的版本
	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}
