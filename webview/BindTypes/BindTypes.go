// 演示 WebView 的 Bind 函数支持的参数和返回值类型.
// 包括: 基本类型 (string, int, float, bool), struct, slice, map, 以及 error 返回值.
package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed assets/**
	embedAssets embed.FS
)

const hostName = "app.example"

// Person 演示 struct 类型.
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

// Company 演示嵌套 struct.
type Company struct {
	Name    string   `json:"name"`
	Staff   []Person `json:"staff"`
	Revenue float64  `json:"revenue"`
}

type Bridge struct {
	wv *edge.WebView
	w  *window.Window
}

func (b *Bridge) bindAll() {
	// ==================== 基本类型 ====================

	// string 参数 + string 返回值
	b.wv.Bind("demo.stringParam", func(s string) string {
		return "Hello, " + s
	})

	// int 参数 + int 返回值
	b.wv.Bind("demo.intParam", func(a, b int) int {
		return a + b
	})

	// float64 参数 + float64 返回值
	b.wv.Bind("demo.floatParam", func(a, b float64) float64 {
		return a * b
	})

	// bool 参数 + bool 返回值
	b.wv.Bind("demo.boolParam", func(flag bool) bool {
		return !flag
	})

	// 无参数, 无返回值
	b.wv.Bind("demo.noop", func() {
		fmt.Println("demo.noop called")
	})

	// 无参数, string 返回值
	b.wv.Bind("demo.hello", func() string {
		return "你好, 世界!"
	})

	// ==================== Slice ====================

	// []int 参数 + []int 返回值 (对每个元素 +1)
	b.wv.Bind("demo.sliceInt", func(nums []int) []int {
		ret := make([]int, len(nums))
		for i, v := range nums {
			ret[i] = v + 1
		}
		return ret
	})

	// []string 参数 + string 返回值
	b.wv.Bind("demo.sliceString", func(strs []string) string {
		var ret string
		for i, s := range strs {
			if i > 0 {
				ret += ", "
			}
			ret += s
		}
		return ret
	})

	// ==================== Map ====================

	// map[string]interface{} 参数 + string 返回值
	b.wv.Bind("demo.mapParam", func(m map[string]interface{}) string {
		var ret string
		for k, v := range m {
			ret += fmt.Sprintf("%s: %v (%T)\n", k, v, v)
		}
		return ret
	})

	// ==================== Struct ====================

	// struct 参数 + struct 返回值
	b.wv.Bind("demo.structParam", func(p Person) Person {
		p.Age += 1 // 过一次生日
		return p
	})

	// 嵌套 struct
	b.wv.Bind("demo.companyInfo", func(c Company) Company {
		c.Revenue *= 1.2 // 营收增加 20%
		for i := range c.Staff {
			c.Staff[i].Age += 1
		}
		return c
	})

	// ==================== error 返回值 ====================

	// 返回 值 + error (成功)
	b.wv.Bind("demo.divideOK", func(a, b float64) (float64, error) {
		return a / b, nil
	})

	// 返回 值 + error (失败)
	b.wv.Bind("demo.divideFail", func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("除数不能为零")
		}
		return a / b, nil
	})

	// 只返回 error
	b.wv.Bind("demo.checkPositive", func(n int) error {
		if n <= 0 {
			return fmt.Errorf("%d 不是正数", n)
		}
		return nil
	})

	// ==================== 混合参数 ====================

	// 混合多种类型参数 + struct 返回值
	b.wv.Bind("demo.mixedArgs", func(name string, age int, scores []int, info map[string]interface{}) Person {
		p := Person{
			Name: name,
			Age:  age,
			City: fmt.Sprintf("总分: %d", sum(scores)),
		}
		if v, ok := info["city"]; ok {
			if s, ok2 := v.(string); ok2 {
				p.City = s
			}
		}
		return p
	})

	// 函数可以接受 []int 等 slice 参数
	b.wv.Bind("demo.calcSum", func(nums []int) int {
		return sum(nums)
	})

	// ==================== 窗口控制 ====================

	b.wv.Bind("wnd.minimize", func() {
		b.w.ShowWindow(xcc.SW_MINIMIZE)
	})
	b.wv.Bind("wnd.toggleMaximize", func() {
		b.w.MaxWindow(!b.w.IsMaxWindow())
	})
	b.wv.Bind("wnd.close", func() {
		b.w.CloseWindow()
	})
}

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func main() {
	checkWebView2()
	edg := createEdge()

	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	win, wv, err := edg.NewWebViewWithWindow(
		edge.WithXmlWindowTitle("Bind 类型演示"),
		edge.WithXmlWindowClassName("BindTypes"),
		edge.WithXmlWindowSize(800, 680),
		edge.WithFillParent(true),
		edge.WithAppDrag(true),
		edge.WithStatusBar(false),
		edge.WithZoomControl(false),
		edge.WithDebug(),
		edge.WithAutoFocus(true),
		edge.WithDefaultBackgroundColor(edge.NewColor(0xFF, 0xFF, 0xFF, 0xFF)),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}

	// 设置虚拟主机名和嵌入文件系统之间的映射
	edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
	wv.EnableVirtualHostNameToEmbedFSMapping(true)

	// 绑定所有演示函数
	b := &Bridge{wv: wv, w: win}
	b.bindAll()

	// 导航完成后再显示窗口, 避免白屏闪烁
	wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
		win.Show(true)
		return 0
	})

	// 访问 HTML
	wv.Navigate(edge.JoinUrlHeader(hostName) + "/BindTypes.html")

	a.Run()
	a.Exit()
}

func createEdge() *edge.Edge {
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(),
		EnvOptions: &edge.EnvOptions{
			DisableTrackingPrevention: true,
			ScrollBarStyle:            edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY,
		},
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 webview 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
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
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)
	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		fmt.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}
