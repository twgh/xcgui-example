// 自动检测并安装 WebView2 运行时
package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

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

func main() {
	AutomaticInstallWebView2Runtime()
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	// 这里 false 是为了兼容 win7
	a := app.New(false)
	a.EnableAutoDPI(true).EnableDPI(true)

	// 创建窗口
	w := window.New(0, 0, 1400, 900, "自动检测并安装 WebView2 运行时", 0, xcc.Window_Style_Default)

	// 创建 WebView
	wv, err := edg.NewWebView(w.Handle,
		edge.WithFillParent(true), // WebView 填充窗口
		edge.WithDebug(true),      // 可打开开发者工具
		edge.WithAutoFocus(true),  // 在窗口获得焦点时尝试保持 WebView 的焦点
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 导航
	wv.Navigate("https://www.baidu.com")

	// 显示窗口并运行应用
	w.Show(true)
	a.Run()
	a.Exit()
}

//go:embed MicrosoftEdgeWebview2Setup.exe
var WebView2Installer []byte

// 自动检测并安装 WebView2 运行时
func AutomaticInstallWebView2Runtime() {
	// 输出本库使用的 WebView2 版本
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())

	// 获取本机已安装的 WebView2 运行时版本
	localVersion, _ := edge.GetAvailableBrowserVersion()
	if localVersion == "" {
		go func() {
			wapi.MessageBoxW(0, "首次运行请等待安装必要的运行环境...", "提示", wapi.MB_OK)
		}()

		// 运行 WebView2 运行时的小型安装引导程序
		err := RunWebView2Installer(false)
		if err != nil {
			wapi.MessageBoxW(0, "安装 WebView2 运行时失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(1)
		}

		// 等待安装完成, 等使用 edge.GetAvailableBrowserVersion() 获取到版本号就是安装完成了
		for i := 0; i < 300; i++ {
			time.Sleep(time.Second)
			localVersion, _ = edge.GetAvailableBrowserVersion()
			if localVersion != "" {
				break
			}
		}
		if localVersion == "" {
			wapi.MessageBoxW(0, "WebView2 运行时安装超时", "错误", wapi.MB_OK|wapi.MB_IconError)
			os.Exit(1)
		}
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)
}

// RunWebView2Installer 运行 WebView2 运行时的小型安装引导程序
//
// isSilent: 是否静默安装
func RunWebView2Installer(isSilent ...bool) error {
	// 文件路径
	installerPath := filepath.Join(os.TempDir(), "Webview2Installer_"+time.Now().Format("20060102")+".exe")
	// 写出文件
	err := os.WriteFile(installerPath, WebView2Installer, 0777)
	if err != nil {
		return err
	}

	// 运行安装程序, 等待安装完成
	cmd := exec.Command(installerPath)
	// 是否静默安装
	if len(isSilent) > 0 && isSilent[0] {
		cmd.Args = append(cmd.Args, "/silent", "/install")
	}

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("运行安装程序失败: %w", err)
	}

	return nil
}
