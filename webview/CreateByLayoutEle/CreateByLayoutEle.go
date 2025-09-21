// 在布局元素中创建 WebView
package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed res/main.xml
	xmlStr string
	//go:embed res/title.png
	png_title []byte
	//go:embed res/player.html
	playerHTML string

	w       *window.Window
	editUrl *widget.Edit

	isDubug         = true                  // 是否为调试版
	eventSwitch     = make(map[string]bool) // 事件开关
	eventSwitchFile = "webview/CreateByLayoutEle/event.switch.json"
)

func loadConfig() {
	if !xc.PathExists2(eventSwitchFile) {
		if err := os.WriteFile(eventSwitchFile, []byte("{}"), 0666); err != nil {
			log.Println("创建 event.switch.json 失败:", err)
		}
	} else {
		// 读取 json
		data, err := os.ReadFile(eventSwitchFile)
		if err != nil {
			log.Println("读取 event.switch.json 失败:", err)
		} else {
			if err := json.Unmarshal(data, &eventSwitch); err != nil {
				log.Println("解析 event.switch.json 失败:", err)
			}
		}
	}
}

func saveConfig() {
	data, err := json.MarshalIndent(eventSwitch, "", "  ")
	if err != nil {
		log.Println("序列化 eventSwitch 失败:", err)
	} else {
		if err := os.WriteFile(eventSwitchFile, data, 0666); err != nil {
			log.Println("保存 event.switch.json 失败:", err)
		}
	}
}

func main() {
	CheckWebView2()
	app.InitOrExit()
	loadConfig()

	// 初始化界面库
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 创建窗口
	w = window.NewByLayoutStringW(xmlStr, 0, 0)
	// 设置窗口透明度
	w.SetTransparentAlpha(255)

	// 放置 WebView 的布局元素
	layoutWV := widget.NewLayoutEleByName("布局WV")

	// 设置全局 WebView 错误回调. 调用 Must 系列方法时, 出错会触发该回调. 还有一些不方便直接 return 的地方也会把错误报告到该回调.
	edge.SetErrorCallBack(func(err *edge.WebViewError) {
		if isDubug {
			log.Println(err.ErrorWithFile())
		} else {
			log.Println(err.ErrorWithFullName()) // 不包含报错源码路径, 这是敏感信息
		}
	})

	//  Edge 在整个应用程序的生命周期里应该只创建一次.
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 自己的软件应该在固定位置创建一个自己的目录, 而不是用临时目录
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 环境失败: "+err.Error(), "错误", wapi.MB_IconError)
		os.Exit(3)
	}

	// 创建 WebView
	wv, err := createWebView(edg, layoutWV.Handle)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_IconError)
		os.Exit(3)
	}

	addXcEvent(wv)

	// 窗口_调整布局, 只要是从布局文件创建窗口, 显示前都要调用, 否则布局会错乱
	w.AdjustLayout()
	w.Show(true)
	a.Run()
	a.Exit()
}

func addXcEvent(wv *edge.WebView) {
	var err error
	hLayout := wv.GetHParent()
	// 按钮_隐藏
	btnHide := widget.NewButtonByName("按钮_隐藏")
	btnHide.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		btnHide.Enable(false).Redraw(false)
		defer btnHide.Enable(true).Redraw(false)
		if wv.CoreWebView == nil {
			w.MessageBox("提示", "webview 不存在!", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}

		isShow := xc.XWidget_IsShow(hLayout)
		if isShow {
			btnHide.SetText("显示")
		} else {
			btnHide.SetText("隐藏")
		}

		xc.XWidget_Show(hLayout, !isShow)
		xc.XEle_Redraw(hLayout, false)
		return 0
	})

	// 按钮_前进
	btnForward := widget.NewButtonByName("按钮_前进")
	btnForward.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		btnForward.Enable(false).Redraw(false)
		defer btnForward.Enable(true).Redraw(false)
		if wv.CoreWebView == nil {
			return 0
		}
		wv.CoreWebView.GoForward()
		return 0
	})

	// 按钮_后退
	btnBack := widget.NewButtonByName("按钮_后退")
	btnBack.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		btnBack.Enable(false).Redraw(false)
		defer btnBack.Enable(true).Redraw(false)
		if wv.CoreWebView == nil {
			return 0
		}
		wv.CoreWebView.GoBack()
		return 0
	})

	// 编辑框_地址栏
	editUrl = widget.NewEditByName("编辑框_地址栏")
	// 按钮_跳转
	btnJump := widget.NewButtonByName("按钮_跳转")
	// 按钮_JS测试
	btnJsTest := widget.NewButtonByName("按钮_JS测试")
	// 按钮_销毁
	btnDestroy := widget.NewButtonByName("按钮_销毁")

	// 跳转网页的函数
	jumpFunc := func() {
		btnJump.Enable(false).Redraw(false)
		defer btnJump.Enable(true).Redraw(false)
		if wv.CoreWebView == nil {
			return
		}
		addr := strings.TrimSpace(editUrl.GetTextEx())
		if addr != "" {
			wv.Navigate(addr)
		}
	}

	// 按钮_跳转事件
	btnJump.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		jumpFunc()
		return 0
	})

	// 编辑框_地址栏事件
	editUrl.AddEvent_KeyDown(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
		if wParam == xcc.VK_Enter { // 判断按下回车键
			jumpFunc()
		}
		return 0
	})

	// 按钮_JS测试事件
	btnJsTest.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			w.MessageBox("提示", "webview 不存在!", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}
		codeWindow := window.New(0, 0, 600, 500, "测试JS代码", w.GetHWND(), xcc.Window_Style_Default)
		codeWindow.EnableLayout(true)
		codeWindow.SetAlignV(xcc.Layout_Align_Top)
		// 代码框
		codeEdit := widget.NewEdit(0, 0, 0, 0, codeWindow.Handle)
		codeEdit.LayoutItem_SetWidth(xcc.Layout_Size_Fill, -1)
		codeEdit.LayoutItem_SetHeight(xcc.Layout_Size_Percent, 60)
		codeEdit.SetText("alert('Hello World')").EnableMultiLine(true).SetDefaultText("请输入代码")
		codeWindow.SetFocusEle(codeEdit.Handle)
		// 选择框
		checkBox := widget.NewButton(0, 0, 0, 0, "获取返回值", codeWindow.Handle)
		checkBox.SetTypeEx(xcc.Button_Type_Check).EnableBkTransparent(true)
		checkBox.LayoutItem_SetWidth(xcc.Layout_Size_Fill, -1)
		checkBox.LayoutItem_SetHeight(xcc.Layout_Size_Percent, 8)
		// 输出框
		resultEdit := widget.NewEdit(0, 0, 0, 0, codeWindow.Handle)
		resultEdit.LayoutItem_SetWidth(xcc.Layout_Size_Fill, -1)
		resultEdit.LayoutItem_SetHeight(xcc.Layout_Size_Percent, 20)
		resultEdit.EnableReadOnly(true).EnableMultiLine(true).SetDefaultText("这里会输出结果").EnableAutoShowScrollBar(true)
		// 日志函数
		elog := func(s string) {
			resultEdit.MoveEnd()
			resultEdit.AddText(time.Now().Format("[15:04:05] ") + s + "\n").ScrollBottom()
			resultEdit.Redraw(false)
		}
		// 执行按钮
		btn := widget.NewButton(0, 0, 0, 0, "执行", codeWindow.Handle)
		btn.LayoutItem_SetWidth(xcc.Layout_Size_Fill, -1)
		btn.LayoutItem_SetHeight(xcc.Layout_Size_Percent, 12)
		btn.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			code := strings.TrimSpace(codeEdit.GetTextEx())
			if code != "" {
				btn.Enable(false).Redraw(false)
				if checkBox.IsCheck() { // 获取返回值
					/*{ // 同步获取
						ret, err := wv.EvalSync(code)
						if err != nil {
							elog("EvalSync, 返回错误: " + err.Error())
							return 0
						}
						elog(fmt.Sprintf("EvalSync, js返回结果: %s", ret))
					}*/

					{ // 异步获取
						wv.EvalAsync(code, func(errorCode syscall.Errno, result string) uintptr {
							if !wapi.IsOK(errorCode) {
								elog("EvalAsync, 返回错误码: " + strconv.Itoa(int(errorCode)))
								return 0
							}
							elog(fmt.Sprintf("EvalAsync, js返回结果: %s", result))
							return 0
						})
					}
				} else {
					wv.Eval(code)
				}
				btn.Enable(true).Redraw(false)
			} else {
				codeWindow.SetFocusEle(codeEdit.Handle)
			}
			return 0
		})
		codeWindow.Show(true)
		return 0
	})

	// 按钮_销毁事件
	btnDestroy.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if btnDestroy.GetText() == "销毁" && wapi.IsWindow(wv.GetHWND()) {
			wv.Close()
			btnDestroy.SetText("创建").Redraw(false)
		} else {
			wv, err = createWebView(wv.Edge, hLayout)
			if err != nil {
				w.MessageBox("提示", err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
				return 0
			}
			xc.XEle_PostEvent(hLayout, xcc.XE_SIZE, 0, 0)
			btnDestroy.SetText("销毁").Redraw(false)
		}
		return 0
	})

	// 按钮_设置搜索关键词
	btnSetSearch := widget.NewButtonByName("按钮_设置搜索关键词")
	btnSetSearch.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Eval("document.querySelectorAll('#chat-textarea')[0].value = 'Microsoft Edge WebView2 简介'")
		return 0
	})

	// 按钮_点击搜索
	btnClickSearch := widget.NewButtonByName("按钮_点击搜索")
	btnClickSearch.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Eval("document.querySelectorAll('#chat-submit-button')[0].click()")
		return 0
	})

	// 按钮_执行Go函数
	btnGoFunc := widget.NewButtonByName("按钮_执行Go函数")
	btnGoFunc.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}

		wv.Eval(`
			goAddStr('Hello World', 666).then(function(result) {
				alert(result);
			});
		`)
		return 0
	})

	// 按钮_刷新
	btnRefresh := widget.NewButtonByName("按钮_刷新")
	btnRefresh.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.CoreWebView.Reload()
		return 0
	})

	// 按钮_打开vea
	btnOpenVea := widget.NewButtonByName("按钮_打开vea")
	btnOpenVea.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Navigate("https://panjiachen.github.io/vue-element-admin/")
		return 0
	})

	// 按钮_打开vben
	btnOpenVben := widget.NewButtonByName("按钮_打开vben")
	btnOpenVben.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Navigate("https://www.vben.pro/")
		return 0
	})

	// 按钮_打开百度
	btnOpenBaidu := widget.NewButtonByName("按钮_打开百度")
	btnOpenBaidu.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Navigate("https://www.baidu.com")
		return 0
	})

	// 按钮_浏览器内核
	btnKernel := widget.NewButtonByName("按钮_浏览器内核")
	btnKernel.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Navigate("https://ie.icoa.cn")
		return 0
	})

	// 按钮_前端发消息到后端
	btnSendMsg := widget.NewButtonByName("按钮_前端发消息到后端")
	btnSendMsg.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Eval("window.chrome.webview.postMessage('我是前端 js 发来的文本消息');" +
			"window.chrome.webview.postMessage(JSON.stringify({type: 'alert', data: '前端来 json 消息了'}));")
		return 0
	})

	// 按钮_新建窗口
	btnNewWindow := widget.NewButtonByName("按钮_新建窗口")
	btnNewWindow.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		*pbHandled = true // 创建新窗口时需要这个拦截
		w2 := window.New(0, 0, 1000, 800, "新建窗口", 0, xcc.Window_Style_Default)

		wv2, err := wv.Edge.NewWebView(w2.Handle, edge.WithFillParent(true))
		if err != nil {
			w2.CloseWindow()
			w.MessageBox("提示", "新建 webview 窗口失败:"+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
			return 0
		}

		// 导航开始事件
		wv2.Event_NavigationStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationStartingEventArgs) uintptr {
			fmt.Println("-------- wv2 导航开始事件 --------")
			uri := args.MustGetUri()
			fmt.Printf("导航开始, 当前uri: %s\n", uri)
			return 0
		})

		wv2.Navigate("https://www.sogou.com")
		w2.Show(true)
		return 0
	})

	// 按钮_是否在新窗口打开新页面
	btnIsNewWindow := widget.NewButtonByName("按钮_在新窗口打开新页面").SetTypeEx(xcc.Button_Type_Check)
	btnIsNewWindow.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		if bCheck {
			w.SetProperty("在新窗口打开新页面", "1")
		} else {
			w.SetProperty("在新窗口打开新页面", "0")
		}
		return 0
	})

	// 按钮_WindowClose
	btnWindowClose := widget.NewButtonByName("按钮_WindowClose")
	btnWindowClose.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Eval("window.close();")
		return 0
	})

	// 按钮_拦截重定向
	btnInterceptRedirection := widget.NewButtonByName("按钮_拦截重定向").SetTypeEx(xcc.Button_Type_Check)
	btnInterceptRedirection.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		if bCheck {
			w.SetProperty("拦截重定向", "1")
		} else {
			w.SetProperty("拦截重定向", "0")
		}
		return 0
	})

	// 按钮_截图
	btnJieTu := widget.NewButtonByName("按钮_截图")
	btnJieTu.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		fileName := time.Now().Format("WebView_2006-01-02_15-04-05") + ".png"
		absPath, _ := filepath.Abs(fileName)
		// 创建流
		stream, err := edge.NewStreamOnFileEx(absPath, wapi.STGM_READWRITE|wapi.STGM_CREATE, wapi.FILE_ATTRIBUTE_NORMAL, true)
		if err != nil {
			w.MessageBox("提示", "创建截图流失败:"+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
			return 0
		}
		err = wv.CapturePreview(edge.COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG, stream, func(errorCode syscall.Errno) uintptr {
			defer stream.Release()
			if !wapi.IsOK(errorCode) {
				w.MessageBox("提示", "截图失败: "+errorCode.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
				return 0
			}
			if w.MessageBox("提示", "截图保存成功: "+absPath+"\n是否打开?", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Cancel, xcc.Window_Style_Default) == xcc.MessageBox_Flag_Ok {
				wapi.ShellExecuteW(0, "open", absPath, "", "", xcc.SW_SHOW)
			}
			return 0
		})
		if err != nil {
			w.MessageBox("提示", "截图失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
			return 0
		}
		return 0
	})

	// 按钮_任务管理器
	btnTaskManager := widget.NewButtonByName("按钮_任务管理器")
	btnTaskManager.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		if wv.WebView2_6 == nil {
			wv.WebView2_6, err = wv.CoreWebView.GetICoreWebView2_6()
			if err != nil {
				w.MessageBox("提示", "打开任务管理器失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
				return 0
			}
		}
		wv.WebView2_6.OpenTaskManagerWindow()
		return 0
	})
	// 按钮_下载文件
	btnDownloadFile := widget.NewButtonByName("按钮_下载文件")
	btnDownloadFile.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		wv.Eval("window.location.href = 'https://msedge.sf.dl.delivery.mp.microsoft.com/filestreamingservice/files/b66a4e63-9084-4d3d-bc7e-1fe47aca92da/MicrosoftEdgeWebView2RuntimeInstallerX64.exe';")
		return 0
	})
	// 按钮_静音
	btnMute := widget.NewButtonByName("按钮_静音")
	btnMute.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		btnMute.Enable(false).Redraw(false)
		defer btnMute.Enable(true).Redraw(false)
		if wv.CoreWebView == nil {
			return 0
		}
		if wv.WebView2_8 == nil {
			wv.WebView2_8, err = wv.CoreWebView.GetICoreWebView2_8()
			if err != nil {
				w.MessageBox("提示", "静音失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
				return 0
			}
		}
		if wv.WebView2_8.MustGetIsMuted() {
			wv.WebView2_8.SetIsMuted(false)
			btnMute.SetText("开启静音")
		} else {
			wv.WebView2_8.SetIsMuted(true)
			btnMute.SetText("取消静音")
		}
		return 0
	})
	// 按钮_事件开关
	btnOutputManager := widget.NewButtonByName("按钮_事件开关")
	btnOutputManager.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		*pbHandled = true // 创建新窗口时需要这个拦截
		eventSwitchWindow := window.New(0, 0, 640, 500, "事件开关", w.GetHWND(), xcc.Window_Style_Default)
		eventSwitchWindow.SetPadding(6, 4, 6, 4)
		eventSwitchWindow.SetBorderSize(1, 30, 1, 1)
		eventSwitchWindow.EnableLayout(true)
		// 全选/取消全选
		btnSelectAll := widget.NewButton(0, 0, 200, 30, "全选", eventSwitchWindow.Handle)
		btnSelectAll.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			btnSelectAll.Enable(false).Redraw(false)
			defer btnSelectAll.Enable(true).Redraw(false)
			var selectAll bool
			if btnSelectAll.GetText() == "全选" {
				selectAll = true
				btnSelectAll.SetText("取消全选")
			} else {
				selectAll = false
				btnSelectAll.SetText("全选")
			}
			// 遍历所有的复选框
			for i := int32(0); i < eventSwitchWindow.GetChildCount(); i++ {
				hEle := eventSwitchWindow.GetChildByIndex(i)
				if xc.XObj_GetTypeEx(hEle) == xcc.Button_Type_Check {
					xc.XBtn_SetCheck(hEle, selectAll)
					xc.XEle_SendEvent(hEle, xcc.XE_BUTTON_CHECK, common.BoolPtr(selectAll), 0)
					xc.XEle_Redraw(hEle, false)
				}
			}
			return 0
		})
		keys := make([]string, 0, len(eventSwitch))
		for k := range eventSwitch {
			keys = append(keys, k)
		}
		sort.Strings(keys) // 按键名排序

		count := 0 // 没有选中的复选框数量
		// 批量创建复选框
		for i, name := range keys {
			b := eventSwitch[name]
			cbb := widget.NewButton(0, 0, 200, 30, name, eventSwitchWindow.Handle)
			cbb.SetTypeEx(xcc.Button_Type_Check).EnableBkTransparent(true)
			cbb.SetCheck(b)
			if !b {
				count++
			}
			if i == 0 { // 强制第1个复选框换行
				cbb.LayoutItem_EnableWrap(true)
			}
			cbb.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
				text := cbb.GetText()
				eventSwitch[text] = bCheck
				return 0
			})
		}
		if count == 0 {
			btnSelectAll.SetText("取消全选")
		} else {
			btnSelectAll.SetText("全选")
		}
		// 窗口销毁时保存配置
		eventSwitchWindow.AddEvent_Destroy(func(hWindow int, pbHandled *bool) int {
			saveConfig()
			return 0
		})
		eventSwitchWindow.Show(true)
		return 0
	})
	// 按钮_播放本地视频
	btnPlayLocalVideo := widget.NewButtonByName("按钮_播放本地视频")
	btnPlayLocalVideo.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		// 由于访问本地文件是有限制的, 所以这里需要设置虚拟主机名映射本地文件夹, 可设置多个主机名映射多个文件夹
		err := wv.SetVirtualHostNameToFolderMapping("local.video", "D:\\twgh\\Videos", edge.COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW)
		if err != nil {
			w.MessageBox("提示", "设置播放视频的虚拟主机名失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
			return 0
		}
		html := strings.Replace(playerHTML, `<source src=""`, `<source src="https://local.video/1.mp4"`, 1)
		// todo 目前想要自动播放的话要注册导航完成事件, 用js来触发播放.
		//  创建环境时传入命令行参数也可以, 但在go里怎么传还有待测试.
		wv.NavigateToString(html)
		return 0
	})
	// 按钮_保存为PDF
	btnSaveAsPDF := widget.NewButtonByName("按钮_保存为PDF")
	btnSaveAsPDF.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		// 获取保存路径
		savePath := wutil.SaveFileEx(wutil.OpenFileOption{
			HwndOwner:   w.GetHWND(),
			Title:       "请选择保存PDF文件路径",
			DefDir:      "%USERPROFILE%\\Desktop",
			DefFileName: "webview2_page.pdf",
			DefExt:      "pdf",
			Filters:     []string{"PDF Files (*.pdf)", "*.pdf"},
		})
		if savePath == "" {
			return 0
		}
		if wv.WebView2_7 == nil {
			wv.WebView2_7, err = wv.CoreWebView.GetICoreWebView2_7()
			if err != nil {
				w.MessageBox("提示", "保存为PDF失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
				return 0
			}
		}
		// 创建 ICoreWebView2PrintSettings 对象并设置打印参数
		env6 := wv.Edge.Environment.MustGetICoreWebView2Environment6()
		printSettings, _ := env6.CreatePrintSettings()
		if printSettings != nil {
			// 设置打印背景颜色和图像
			printSettings.SetShouldPrintBackgrounds(true)
		}

		// 异步将当前页面打印为 PDF 文件
		wv.WebView2_7.PrintToPdfEx(wv.GetWebViewEventImpl(), savePath, printSettings, func(errorCode syscall.Errno, isSuccessful bool) uintptr {
			if !wapi.IsOK(errorCode) || !isSuccessful {
				w.MessageBox("提示", "保存为PDF失败: "+errorCode.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
			} else {
				w.MessageBox("提示", "保存为PDF成功: "+savePath, xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			}
			return 0
		})
		return 0
	})
	// 按钮_保存网页
	btnSavePage := widget.NewButtonByName("按钮_保存网页")
	btnSavePage.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if wv.CoreWebView == nil {
			return 0
		}
		if wv.WebView2_25 == nil {
			wv.WebView2_25, err = wv.CoreWebView.GetICoreWebView2_25()
			if err != nil {
				w.MessageBox("提示", "保存网页失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Default)
				return 0
			}
		}
		// 显示"另存为"UI，允许用户保存当前页面。
		wv.WebView2_25.ShowSaveAsUIEx(wv.GetWebViewEventImpl(), func(errorCode syscall.Errno, result edge.COREWEBVIEW2_SAVE_AS_UI_RESULT) uintptr {
			fmt.Println("ShowSaveAsUIEx 结果:", errorCode, result)
			return 0
		})
		return 0
	})
}

func createWebView(edg *edge.Edge, hParent int) (*edge.WebView, error) {
	// 创建 webview
	wv, err := edg.NewWebView(hParent,
		edge.WithFillParent(true),
		edge.WithDebug(true),
	)
	if err != nil {
		return nil, fmt.Errorf("创建 webview 失败: %v", err)
	}

	// 获取浏览器设置
	settings, _ := wv.GetSettings()
	// 获取浏览器设置2
	s2, _ := settings.GetICoreWebView2Settings2()
	if s2 != nil {
		ua := s2.MustGetUserAgent()
		fmt.Println("浏览器 UserAgent:", ua)
		s2.Release()
	}
	settings.Release()

	var suspendTimer *time.Timer
	var suspendTime = 10 * time.Minute
	// 最小化一段时间后挂起 webview 以节省内存.
	// 注意: 如果正在导航, 则不应挂起, 否则等导航完成, 恢复时webview将不可见.
	// 这里我没有写判断的代码, 如果要写的话可以结合导航开始/完成事件来判断.
	w.AddEvent_WindProc(func(hWindow int, message uint32, wParam, lParam uintptr, pbHandled *bool) int {
		if wv.Controller == nil { // webview 未创建或已被销毁
			return 0
		}

		if message == wapi.WM_SIZE {
			switch wParam {
			case wapi.SIZE_MINIMIZED: // 窗口最小化
				// 设置定时器
				if suspendTimer != nil {
					suspendTimer.Stop()
				}
				suspendTimer = time.AfterFunc(suspendTime, func() {
					xc.XC_CallUT(func() { // 需在 UI 线程操作 WebView
						// 挂起前要隐藏 webview 才行
						wv.Show(false)
						fmt.Println("开始挂起")
						err := wv.TrySuspend(func(errorCode syscall.Errno, isSuccessful bool) uintptr {
							if !wapi.IsOK(errorCode) || !isSuccessful {
								log.Println("挂起失败, errorCode:", errorCode, ", isSuccessful:", isSuccessful)
							} else {
								fmt.Println("挂起成功")
							}
							return 0
						})
						if err != nil {
							log.Println("执行 TrySuspend 失败, err:", err)
						}
					})
				})
				fmt.Printf("窗口最小化，%v 内恢复可避免挂起\n", suspendTime)
			case wapi.SIZE_RESTORED: // 窗口恢复
				// 取消挂起定时器
				if suspendTimer != nil {
					suspendTimer.Stop()
					suspendTimer = nil
					fmt.Println("已停止挂起定时器")
				}
				if wv.IsSuspended() {
					wv.Resume()
					wv.Show(true)
					fmt.Println("从挂起状态恢复正常")
				}
			}
		}
		return 0
	})

	addWebviewEvent(wv)

	// 绑定Go函数
	if err := wv.Bind("goAddStr", func(str string, num int) string {
		fmt.Println("执行Go函数: goAddStr")
		return "传进Go函数 goAddStr 的参数: " + str + ", " + strconv.Itoa(num)
	}); err != nil {
		log.Println("绑定Go函数 goAddStr 失败:", err.Error())
	}

	// 绑定一个输出函数, 方便在js中调用
	if err := wv.Bind("go.log", func(a interface{}) {
		fmt.Printf("js输出: %v\n", a)
	}); err != nil {
		log.Println("绑定Go函数 go.log 失败:", err.Error())
	}

	// 添加网络资源请求过滤, 这里获取出错是因为运行时版本低
	wv.WebView2_22, err = wv.CoreWebView.GetICoreWebView2_22()
	if err == nil {
		wv.WebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds("*", edge.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL, edge.COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_ALL)
	} else {
		wv.AddWebResourceRequestedFilter("*", edge.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
	}

	// 加载网页
	wv.Navigate("https://www.baidu.com")
	return wv, nil
}

// webview 事件
func addWebviewEvent(wv *edge.WebView) {
	// 网页消息事件
	wv.Event_WebMessageReceived(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2WebMessageReceivedEventArgs) uintptr {
		if !eventSwitch["网页消息事件"] {
			return 0
		}
		fmt.Println("-------- 网页消息事件 --------")
		message := args.MustTryGetWebMessageAsString()
		fmt.Println("消息源:", args.MustGetSource())
		fmt.Println("消息:", message)
		if message == "" {
			return 0
		}

		// 从前端发来的json消息: {"type":"alert","data":"前端来json消息了"} 解析出参数
		if strings.HasPrefix(message, "{\"type\":\"") {
			var data map[string]string
			if err := json.Unmarshal([]byte(message), &data); err != nil {
				log.Println("解析json消息失败:", err.Error())
				return 0
			}
			switch data["type"] {
			case "alert":
				xc.XC_Alert("提示", data["data"])
			}
		}
		return 0
	})

	// 快捷键事件
	wv.Event_AcceleratorKeyPressed(func(sender *edge.ICoreWebView2Controller, args *edge.ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
		if !eventSwitch["快捷键事件"] {
			return 0
		}
		fmt.Println("-------- 快捷键事件 --------")
		eventKind, _ := args.GetKeyEventKind()
		virtualKey, _ := args.GetVirtualKey()
		status, _ := args.GetPhysicalKeyStatus()
		if eventKind == edge.COREWEBVIEW2_KEY_EVENT_KIND_KEY_DOWN ||
			eventKind == edge.COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_DOWN {
			if !status.WasKeyDown {
				fmt.Println("快捷键按下:", virtualKey)
				// 这只是演示快捷键事件怎么用, 实际只需要 wv.MustGetSettings().MustGetICoreWebView2Settings3().SetAreBrowserAcceleratorKeysEnabled(false) 就能禁止浏览器的一些特定快捷键
				switch virtualKey {
				case xcc.VK_F5:
					fmt.Println("拦截 F5 刷新")
					args.SetHandled(true)
				case xcc.VK_R:
					if wutil.IsKeyPressed(xcc.VK_Ctrl) {
						fmt.Println("拦截 Ctrl+R 刷新")
						args.SetHandled(true)
					}
				}
			}
		}
		return 0
	})

	// 网络资源请求事件
	wv.Event_WebResourceRequested(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2WebResourceRequestedEventArgs) uintptr {
		if !eventSwitch["网络资源请求事件"] {
			return 0
		}
		request, _ := args.GetRequest()
		method := request.MustGetMethod()
		if method == http.MethodPost {
			uri := request.MustGetUri()
			headers := request.MustGetHeadersMap()
			content := string(request.MustGetContent())
			fmt.Println("-------- 网络资源请求事件回调 Post --------")
			fmt.Printf("请求方法: %s, 地址:%s\n", method, uri)
			fmt.Println("请求头:", headers)
			fmt.Println("请求内容:", content)
		}
		return 0
	})

	// 导航开始事件
	wv.Event_NavigationStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationStartingEventArgs) uintptr {
		if !eventSwitch["导航开始事件"] {
			return 0
		}
		fmt.Println("-------- 导航开始事件 --------")
		uri := args.MustGetUri()
		navigationId := args.MustGetNavigationId()
		isUserInitiated := args.MustGetIsUserInitiated()
		isRedirect := args.MustGetIsRedirected()
		editUrl.SetText(uri).Redraw(false)
		fmt.Printf("导航开始, 当前uri: %s, 导航id: %d, 用户发起: %v, 重定向: %v\n", uri, navigationId, isUserInitiated, isRedirect)
		if isRedirect && w.GetProperty("拦截重定向") == "1" {
			args.SetCancel(true)
			fmt.Println("-------- 已拦截重定向 --------")
			return 0
		}
		return 0
	})

	// 源改变事件
	wv.Event_SourceChanged(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2SourceChangedEventArgs) uintptr {
		if !eventSwitch["源改变事件"] {
			return 0
		}
		fmt.Println("-------- 源改变事件 --------")
		fmt.Printf("当前源: %s, 导航到的页面是否为一个新文档: %v\n", sender.MustGetSource(), args.MustGetIsNewDocument())
		return 0
	})

	// 网页内容正在加载事件
	wv.Event_ContentLoading(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2ContentLoadingEventArgs) uintptr {
		if !eventSwitch["网页内容正在加载事件"] {
			return 0
		}
		fmt.Println("-------- 网页内容正在加载事件 --------")
		fmt.Printf("当前源: %s, 导航id: %d, 是否为错误页面: %v\n", sender.MustGetSource(), args.MustGetNavigationId(), args.MustGetIsErrorPage())
		return 0
	})

	// 导航完成事件
	wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		if !eventSwitch["导航完成事件"] {
			return 0
		}
		fmt.Println("-------- 导航完成事件 --------")
		source := sender.MustGetSource()
		navigationId := args.MustGetNavigationId()
		webErrorStatus := args.MustGetWebErrorStatus()
		if args.MustGetIsSuccess() {
			fmt.Printf("导航成功, 当前源: %s, 导航id: %d, 错误码: %d\n", source, navigationId, webErrorStatus)
		} else {
			log.Printf("导航失败, 当前源: %s, 导航id: %d, 错误码: %d\n", source, navigationId, webErrorStatus)
		}

		// 获取 cookies
		if err := wv.GetCookies("https://www.baidu.com", func(errorCode syscall.Errno, cookies *edge.ICoreWebView2CookieList) uintptr {
			if !wapi.IsOK(errorCode) {
				log.Printf("获取 cookies 失败: %v\n", errorCode)
				return 0
			}
			fmt.Println("-------- 获取 Cookies 开始--------")
			for i := uint32(0); i < cookies.MustGetCount(); i++ {
				cookie := cookies.MustGetValueAtIndex(i)
				fmt.Println(cookie.MustGetName(), "\t", cookie.MustGetValue())
			}
			fmt.Println("-------- 获取 Cookies 结束 --------")
			return 0
		}); err != nil {
			log.Printf("获取 cookies 失败: %v\n", err)
		}
		return 0
	})

	// 框架导航开始事件
	wv.Event_Frame_NavigationStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationStartingEventArgs) uintptr {
		if !eventSwitch["框架导航开始事件"] {
			return 0
		}
		fmt.Println("-------- 框架导航开始事件 --------")
		uri := args.MustGetUri()
		fmt.Printf("框架导航开始, 当前uri: %s\n", uri)
		return 0
	})

	// 框架导航完成事件
	wv.Event_Frame_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		if !eventSwitch["框架导航完成事件"] {
			return 0
		}
		fmt.Println("-------- 框架导航完成事件 --------")
		source := sender.MustGetSource()
		navigationId := args.MustGetNavigationId()
		webErrorStatus := args.MustGetWebErrorStatus()
		if args.MustGetIsSuccess() {
			fmt.Printf("子框架导航成功, 当前源: %s, 导航id: %d, 错误码: %d\n", source, navigationId, webErrorStatus)
		} else {
			log.Printf("子框架导航失败, 当前源: %s, 导航id: %d, 错误码: %d\n", source, navigationId, webErrorStatus)
		}
		return 0
	})

	// 新窗口请求事件
	wv.Event_NewWindowRequested(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NewWindowRequestedEventArgs) uintptr {
		if !eventSwitch["新窗口请求事件"] {
			return 0
		}
		fmt.Println("-------- 新窗口请求事件 --------")
		if w.GetProperty("在新窗口打开新页面") == "1" {
			return 0
		}
		// 拦截
		args.SetHandled(true)
		// 在当前 webview 中导航
		uri := args.MustGetUri()
		fmt.Println("拦截新窗口请求, 在当前 webview 中导航, uri: " + uri)
		wv.Navigate(uri)
		return 0
	})

	// 窗口关闭请求事件
	wv.Event_WindowCloseRequested(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["窗口关闭请求事件"] {
			return 0
		}
		fmt.Println("-------- 窗口关闭请求事件 --------")
		ret := w.MessageBox("提示", "接收到了 js 传来的关闭请求: window.close(), 是否关闭 WebView?", xcc.MessageBox_Flag_Icon_Qustion|xcc.MessageBox_Flag_Cancel|xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
		if ret == xcc.MessageBox_Flag_Ok {
			wv.Close()
			return 0
		}
		// args 是 nil
		return 0
	})

	// 文档标题改变事件
	wv.Event_DocumentTitleChanged(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["文档标题改变事件"] {
			return 0
		}
		fmt.Println("-------- 文档标题改变事件 --------")
		fmt.Println("文档标题改变, 当前标题: " + sender.MustGetDocumentTitle())
		// args 是 nil
		return 0
	})
	// 全屏元素状态改变事件
	wv.Event_ContainsFullScreenElementChanged(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["全屏元素状态改变事件"] {
			return 0
		}
		fmt.Println("-------- 全屏元素状态改变事件 --------")
		// 视频全屏/取消时会触发
		ContainsFullScreenElement := sender.MustGetContainsFullScreenElement()
		fmt.Println("ContainsFullScreenElement(是否包含全屏元素) 属性:", ContainsFullScreenElement)
		// args 是 nil
		return 0
	})
	// 进程失败事件
	wv.Event_ProcessFailed(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2ProcessFailedEventArgs) uintptr {
		if !eventSwitch["进程失败事件"] {
			return 0
		}
		fmt.Println("-------- 进程失败事件 --------")
		fmt.Println("进程失败, 错误码: ", args.MustGetProcessFailedKind())
		return 0
	})
	// 历史记录改变事件
	wv.Event_HistoryChanged(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["历史记录改变事件"] {
			return 0
		}
		fmt.Println("-------- 历史记录改变事件 --------")
		source := sender.MustGetSource()
		fmt.Printf("历史记录改变, 当前源: %s\n", source)
		// args 是 nil
		return 0
	})

	// 脚本对话框打开事件
	// 必须先禁用 js 默认对话框:
	// wv.CoreWebView.MustGetSettings().SetAreDefaultScriptDialogsEnabled(false)
	wv.Event_ScriptDialogOpening(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2ScriptDialogOpeningEventArgs) uintptr {
		if !eventSwitch["脚本对话框打开事件"] {
			return 0
		}
		fmt.Println("-------- 脚本对话框打开事件 --------")
		dialogType := args.MustGetKind()
		fmt.Println("脚本对话框打开, 类型: ", dialogType)
		return 0
	})

	// DOM内容加载完成事件
	wv.Event_DOMContentLoaded(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2DOMContentLoadedEventArgs) uintptr {
		if !eventSwitch["DOM内容加载完成事件"] {
			return 0
		}
		fmt.Println("-------- DOM内容加载完成事件 --------")
		source := sender.MustGetSource()
		fmt.Printf("DOM 内容加载完成, 当前源: %s\n", source)
		return 0
	})
	// 框架创建完成事件
	wv.Event_FrameCreated(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2FrameCreatedEventArgs) uintptr {
		if !eventSwitch["框架创建完成事件"] {
			return 0
		}
		fmt.Println("-------- 框架创建完成事件 --------")
		frame := args.MustGetFrame()
		fmt.Println("框架名称:", frame.MustGetName())
		return 0
	})
	// 下载开始事件
	wv.Event_DownloadStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2DownloadStartingEventArgs) uintptr {
		if !eventSwitch["下载开始事件"] {
			return 0
		}
		fmt.Println("-------- 下载开始事件 --------")
		fmt.Println("下载开始, 文件路径:", args.MustGetResultFilePath())
		// 获取下载操作对象
		downloadOperation, err := args.GetDownloadOperation()
		if err != nil {
			log.Println("获取下载操作对象失败:", err.Error())
			return 0
		}
		defer downloadOperation.Release()

		fmt.Println("下载地址:", downloadOperation.MustGetUri())
		// 下载字节改变事件
		downloadOperation.Event_BytesReceivedChanged(wv.GetWebViewEventImpl(), func(sender *edge.ICoreWebView2DownloadOperation, args *edge.IUnknown) uintptr {
			state := sender.MustGetState()                             // 获取下载状态
			if state == edge.COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS { // 正在下载
				fileSize := sender.MustGetTotalBytesToReceive()      // 文件总大小
				bytesReceived := sender.MustGetBytesReceived()       // 已下载大小
				estimatedEndTime := sender.MustGetEstimatedEndTime() // 预计结束时间文本
				endTime, _ := time.Parse(time.RFC3339, estimatedEndTime)
				utcNow := time.Now().UTC()
				fmt.Printf("下载进度: %.2f%%, %.2fMB / %.2fMB, 剩余时间: %v\n", float32(bytesReceived)/float32(fileSize)*100, float32(bytesReceived)/1024/1024, float32(fileSize)/1024/1024, endTime.Sub(utcNow))
			}
			return 0
		})
		// 下载状态改变事件
		downloadOperation.Event_StateChanged(wv.GetWebViewEventImpl(), func(sender *edge.ICoreWebView2DownloadOperation, args *edge.IUnknown) uintptr {
			state := sender.MustGetState() // 获取下载状态
			switch state {
			case edge.COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS:
				// 正在下载
			case edge.COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED:
				reason := sender.MustGetInterruptReason() // 获取中断原因
				fmt.Println("下载中断, 原因:", reason)
			case edge.COREWEBVIEW2_DOWNLOAD_STATE_COMPLETED:
				fmt.Println("下载完成, 文件路径:", sender.MustGetResultFilePath())
			}
			return 0
		})
		// 预计结束时间改变事件
		downloadOperation.Event_EstimatedEndTimeChanged(wv.GetWebViewEventImpl(), func(sender *edge.ICoreWebView2DownloadOperation, args *edge.IUnknown) uintptr {
			estimatedEndTime := sender.MustGetEstimatedEndTime() // 预计结束时间文本
			endTime, _ := time.Parse(time.RFC3339, estimatedEndTime)
			utcNow := time.Now().UTC()
			fmt.Println("预计结束时间改变, 预计剩余时间:", endTime.Sub(utcNow))
			return 0
		})
		return 0
	})
	// 静音状态改变事件
	wv.Event_IsMutedChanged(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["静音状态改变事件"] {
			return 0
		}
		fmt.Println("-------- 静音状态改变事件 --------")
		if wv.WebView2_8 == nil {
			var err error
			wv.WebView2_8, err = wv.CoreWebView.GetICoreWebView2_8()
			if err != nil {
				log.Println("获取 ICoreWebView2_8 失败:", err.Error())
				return 0
			}
		}
		fmt.Println("是否静音:", wv.WebView2_8.MustGetIsMuted(), "当前文档是否正在播放音频:", wv.WebView2_8.MustGetIsDocumentPlayingAudio())
		return 0
	})
	// 文档播放音频状态改变事件
	wv.Event_DocumentPlayingAudioChanged(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["文档播放音频状态改变事件"] {
			return 0
		}
		fmt.Println("-------- 文档播放音频状态改变事件 --------")
		if wv.WebView2_8 == nil {
			var err error
			wv.WebView2_8, err = wv.CoreWebView.GetICoreWebView2_8()
			if err != nil {
				log.Println("获取 ICoreWebView2_8 失败:", err.Error())
				return 0
			}
		}
		fmt.Println("当前文档是否正在播放音频:", wv.WebView2_8.MustGetIsDocumentPlayingAudio())
		return 0
	})
	// 创建一个自定义菜单项: 显示文档标题
	var menuItemShowTitle *edge.ICoreWebView2ContextMenuItem
	env9 := wv.Edge.Environment.MustGetICoreWebView2Environment9()
	if env9 != nil {
		defer env9.Release()
		iconStream, err := edge.NewStreamMem(png_title)
		if err != nil {
			log.Println("创建自定义菜单项的图标流失败:", err.Error())
		}
		defer iconStream.Release()
		// 创建自定义菜单项: 显示文档标题
		menuItemShowTitle, err = env9.CreateContextMenuItem("显示文档标题", iconStream, edge.COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND)
		if err != nil {
			log.Println("创建自定义菜单项失败:", err.Error())
		}
		CommandId := menuItemShowTitle.MustGetCommandId()
		// 自定义菜单项选中事件
		menuItemShowTitle.Event_CustomItemSelected(wv.GetWebViewEventImpl(), func(sender *edge.ICoreWebView2ContextMenuItem, args *edge.IUnknown) uintptr {
			fmt.Println("对象是否相等:", sender == menuItemShowTitle) // 理论上判断对象也是可以的
			if sender.MustGetCommandId() == CommandId {
				title := wv.CoreWebView.MustGetDocumentTitle()
				app.Alert("提示", "当前文档标题: "+title)
			}
			return 0
		})
	}
	// 上下文菜单请求事件
	wv.Event_ContextMenuRequested(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2ContextMenuRequestedEventArgs) uintptr {
		if !eventSwitch["上下文菜单请求事件"] {
			return 0
		}
		fmt.Println("-------- 上下文菜单请求事件 --------")
		items := args.MustGetMenuItems()
		defer items.Release()

		// 添加自定义菜单项
		if menuItemShowTitle != nil {
			err := items.InsertValueAtIndex(0, menuItemShowTitle)
			if err != nil {
				log.Println("添加自定义菜单项失败:", err.Error())
			}
		}
		// 遍历上下文菜单项
		for i := uint32(0); i < items.MustGetCount(); i++ {
			item := items.MustGetValueAtIndex(i)
			text := item.MustGetLabel()
			fmt.Println("菜单索引:", i, "文本:", text)
			if text == "打印(&P)" { // 删除菜单项"打印"
				items.RemoveValueAtIndex(i)
			}
			item.Release()
		}
		return 0
	})
	// 进程信息改变事件
	wv.Event_ProcessInfosChanged(func(sender *edge.ICoreWebView2Environment, args *edge.IUnknown) uintptr {
		if !eventSwitch["进程信息改变事件"] {
			return 0
		}
		fmt.Println("-------- 进程信息改变事件 --------")
		e8, err := sender.GetICoreWebView2Environment8()
		if err != nil {
			log.Println("获取 ICoreWebView2Environment8 失败:", err.Error())
			return 0
		}
		defer e8.Release()
		// 获取进程信息列表
		processInfos := e8.MustGetProcessInfos()
		for i := uint32(0); i < processInfos.MustGetCount(); i++ {
			info := processInfos.MustGetValueAtIndex(i)
			if info == nil {
				continue
			}
			fmt.Printf("进程ID: %d, 类型: %d\n", info.MustGetProcessId(), info.MustGetKind())
			info.Release()
		}
		processInfos.Release()
		return 0
	})
	// 网站图标改变事件
	wv.Event_FaviconChanged(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["网站图标改变事件"] {
			return 0
		}
		fmt.Println("-------- 网站图标改变事件 --------")
		if wv.WebView2_15 == nil {
			var err error
			wv.WebView2_15, err = wv.CoreWebView.GetICoreWebView2_15()
			if err != nil {
				log.Println("获取 ICoreWebView2_15 失败:", err.Error())
				return 0
			}
		}
		fmt.Printf("网站图标URI: %s\n", wv.WebView2_15.MustGetFaviconUri())

		// 获取网站图标
		wv.WebView2_15.GetFaviconEx(wv.GetWebViewEventImpl(), edge.COREWEBVIEW2_FAVICON_IMAGE_FORMAT_PNG, func(errorCode syscall.Errno, favicon []byte) uintptr {
			if wapi.IsOK(errorCode) && len(favicon) > 0 {
				img := imagex.NewByMemAdaptive(favicon, 0, 0, 0, 0)
				w.SetIcon(img.Handle)
				w.Redraw(true)
			}
			return 0
		})
		return 0
	})
	// 缩放因子改变事件
	wv.Event_ZoomFactorChanged(func(sender *edge.ICoreWebView2Controller, args *edge.IUnknown) uintptr {
		if !eventSwitch["缩放因子改变事件"] {
			return 0
		}
		fmt.Println("-------- 缩放因子改变事件 --------")
		zoomFactor, _ := sender.GetZoomFactor()
		fmt.Printf("缩放因子改变, 当前缩放因子: %.2f\n", zoomFactor)
		return 0
	})
	// 移动焦点请求事件
	wv.Event_MoveFocusRequested(func(sender *edge.ICoreWebView2Controller, args *edge.ICoreWebView2MoveFocusRequestedEventArgs) uintptr {
		if !eventSwitch["移动焦点请求事件"] {
			return 0
		}
		fmt.Println("-------- 移动焦点请求事件 --------")
		reason, _ := args.GetReason()
		fmt.Println("移动焦点请求, 原因: ", reason)
		return 0
	})
	// 获得焦点事件
	wv.Event_GotFocus(func(sender *edge.ICoreWebView2Controller, args *edge.IUnknown) uintptr {
		if !eventSwitch["获得焦点事件"] {
			return 0
		}
		fmt.Println("-------- 获得焦点事件 --------")
		return 0
	})
	// 失去焦点事件
	wv.Event_LostFocus(func(sender *edge.ICoreWebView2Controller, args *edge.IUnknown) uintptr {
		if !eventSwitch["失去焦点事件"] {
			return 0
		}
		fmt.Println("-------- 失去焦点事件 --------")
		return 0
	})
	// 状态栏文本改变事件
	wv.Event_StatusBarTextChanged(func(sender *edge.ICoreWebView2, args *edge.IUnknown) uintptr {
		if !eventSwitch["状态栏文本改变事件"] {
			return 0
		}
		fmt.Println("-------- 状态栏文本改变事件 --------")
		if wv.WebView2_12 == nil {
			var err error
			wv.WebView2_12, err = wv.CoreWebView.GetICoreWebView2_12()
			if err != nil {
				log.Println("获取 ICoreWebView2_12 失败:", err.Error())
				return 0
			}
		}
		text := wv.WebView2_12.MustGetStatusBarText()
		if text != "" {
			fmt.Println("状态栏文本改变, 当前文本: " + text)
		}
		return 0
	})
	// 检测到服务器证书错误事件
	wv.Event_ServerCertificateErrorDetected(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2ServerCertificateErrorDetectedEventArgs) uintptr {
		if !eventSwitch["检测到服务器证书错误事件"] {
			return 0
		}
		fmt.Println("-------- 检测到服务器证书错误事件 --------")
		fmt.Printf("检测到服务器证书错误, 错误状态: %d, 请求地址: %s\n", args.MustGetErrorStatus(), args.MustGetRequestUri())
		// 忽略证书错误继续导航
		// args.SetAction(edge.COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_ALWAYS_ALLOW)
		return 0
	})
	// 另存为界面显示事件
	wv.Event_SaveAsUIShowing(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2SaveAsUIShowingEventArgs) uintptr {
		if !eventSwitch["另存为界面显示事件"] {
			return 0
		}
		fmt.Println("-------- 另存为界面显示事件 --------")
		return 0
	})
	// 保存文件安全检查开始事件
	wv.Event_SaveFileSecurityCheckStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2SaveFileSecurityCheckStartingEventArgs) uintptr {
		if !eventSwitch["保存文件安全检查开始事件"] {
			return 0
		}
		fmt.Println("-------- 保存文件安全检查开始事件 --------")
		return 0
	})
	// 屏幕截图开始事件
	wv.Event_ScreenCaptureStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2ScreenCaptureStartingEventArgs) uintptr {
		if !eventSwitch["屏幕截图开始事件"] {
			return 0
		}
		fmt.Println("-------- 屏幕截图开始事件 --------")
		return 0
	})
}

func CheckWebView2() {
	// 高版本运行时是兼容低版本的。低版本不用高版本的方法也没有问题。
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())
	// 获取本机已安装的 WebView2 运行时版本号
	localVersion, err := edge.GetAvailableBrowserVersion()
	if err != nil {
		wapi.MessageBoxW(0, "获取 WebView2 运行时版本号失败: "+err.Error(), "提示", wapi.MB_IconError)
		os.Exit(1)
	}
	if localVersion == "" {
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序! 下载完后请使用管理员权限运行安装包!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		// 使用默认浏览器打开 WebView2 运行时小型引导程序下载地址.
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)

	// 判断本机 WebView2 运行时版本是否低于本库使用的 WebView2 运行时版本
	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}
