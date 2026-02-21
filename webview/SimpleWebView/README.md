# Simple - 简单WebView示例

> 本文由 AI 生成

## 项目简介

Simple 是最基础的 WebView2 示例，展示了如何用最少的代码创建一个功能完整的 WebView 桌面应用。该示例使用了 `NewWebViewWithWindow()` 便捷方法，一步完成窗口和 WebView 的创建，专注于演示 WebView2 的核心功能，去除了复杂的配置和高级特性，是初学者学习 XCGUI + WebView2 开发的最佳入门示例。

## 功能特性

### 核心功能
- **基础WebView创建**: 最简化的WebView2实例创建过程
- **网页浏览**: 可以访问任意网页（示例访问百度首页）
- **开发者工具**: 内置开发者工具支持，便于调试

### 界面特性
- **极简界面**: 纯WebView界面，无额外UI元素
- **填充窗口**: WebView完全填充整个窗口区域
- **炫彩窗口**: 使用炫彩窗口标题栏，提供更好的视觉体验

### 技术特性
- **版本检测**: 包含完整的WebView2版本检测机制
- **错误处理**: 基础但完整的错误处理
- **DPI支持**: 支持高DPI显示器适配

## 项目结构

```
SimpleWebView/
├── SimpleWebView.go      # 主程序文件
└── README.md           # 本文档
```

## 技术实现

### 技术栈
- **Go 语言**: 主要编程语言
- **XCGUI**: 炫彩界面库，提供窗口框架
- **WebView2**: Microsoft Edge 内核的嵌入式浏览器
- **便捷API**: 使用 `NewWebViewWithWindow()` 方法一步创建窗口和 WebView

### 核心代码解析

#### 1. 应用初始化
```go
// 初始化界面库
app.InitOrExit()
a := app.New(true)
a.EnableAutoDPI(true).EnableDPI(true)
```

#### 2. WebView环境创建
```go
// 创建 WebView 环境
edg, err := edge.New(edge.Option{
    UserDataFolder: os.TempDir(), // 实际应用中应使用固定目录
})
```

#### 3. WebView和窗口创建（使用便捷方法）

**注意**: 本示例使用了 `NewWebViewWithWindow()` 便捷方法，该方法一步完成窗口和 WebView 的创建。如果需要在现有窗口中嵌入 WebView，可以使用 `NewWebView()` 方法。
```go
// 创建 WebView 和窗口（一步完成）
w, wv, err := edg.NewWebViewWithWindow(
    edge.WithXmlWindowTitle("简单 WebView 例子"), // 窗口标题
    edge.WithXmlWindowSize(1400, 900),           // 窗口大小
    edge.WithXmlWindowTitleBar(true),            // 使用炫彩窗口标题栏
    edge.WithFillParent(true),                   // WebView 填充窗口
    edge.WithDebug(true),                        // 可打开开发者工具
    edge.WithAutoFocus(true),                    // 自动聚焦
)
```

#### 4. 导航到网页
```go
// 导航到百度首页
wv.Navigate("https://www.baidu.com")
```

#### 5. 显示窗口并运行
```go
// 显示窗口并运行应用
w.Show(true)
a.Run()
a.Exit()
```

### 版本检查机制
```go
func checkWebView2() {
    // 输出本库使用的 WebView2 版本
    fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())

    // 获取本机已安装的 WebView2 运行时版本
    localVersion, err := edge.GetAvailableBrowserVersion()
    if err != nil {
        // 错误处理...
    }
    if localVersion == "" {
        // 提示安装WebView2运行时...
        edge.DownloadWebView2()
        os.Exit(2)
    }
    
    // 检查版本兼容性
    if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
        log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
    }
}
```

## 使用方法

### 快速启动
1. 确保系统已安装 WebView2 运行时

2. 运行程序：
   ```bash
   go run SimpleWebView.go
   ```

### 开发者工具使用
- 在WebView区域右键点击
- 选择"检查"或"Inspect"
- 可以进行元素检查、JavaScript调试等操作

### 自定义导航
修改代码中的导航地址：
```go
// 将百度改为其他网站
wv.Navigate("https://www.github.com")
```

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（Microsoft Edge WebView2）

### Go 依赖
- **github.com/twgh/xcgui**: XCGUI Go语言绑定库
- **Go 1.18+**: 支持泛型和最新特性

### 运行时检测
程序启动时会自动检测：
1. WebView2运行时是否已安装
2. 运行时版本是否兼容
3. 如未安装会自动打开下载页面

## 代码特点

### 简洁性
- 无复杂配置和高级特性
- 专注核心功能演示

### 完整性
- 包含完整的错误处理
- 版本兼容性检查
- 资源清理和退出处理

### 可扩展性
- 代码结构清晰，易于扩展
- 可以在此基础上添加更多功能
- 作为其他复杂示例的基础模板

## 学习价值

### 对于初学者
- **理解基本流程**: 学习WebView2集成的基本步骤
- **掌握核心API**: 了解最重要的几个API函数
- **快速上手**: 最快的方式体验WebView2功能

### 对于开发者
- **参考模板**: 作为新项目的起始模板
- **调试基础**: 简单的代码便于调试和问题定位
- **性能基准**: 作为性能对比的基础版本

## 扩展方向

基于这个简单示例，可以扩展为：

### 1. 浏览器应用
```go
// 添加地址栏、前进后退按钮等
// 实现完整的浏览器功能
```

### 2. Web应用容器
```go
// 导航到特定的Web应用
wv.Navigate("https://your-web-app.com")
// 添加应用特定的功能
```

### 3. 混合桌面应用
```go
// 结合XCGUI控件和WebView
// 实现桌面UI和Web内容的混合应用
// 也可以使用 NewWebView() 在现有窗口中嵌入WebView
```

## 最佳实践

### 1. 用户数据目录
```go
// 生产环境建议使用固定目录
UserDataFolder: filepath.Join(os.Getenv("APPDATA"), "YourAppName")
```

### 2. 错误处理
```go
// 添加更详细的错误处理
if err != nil {
    log.Printf("WebView创建失败: %v", err)
    // 显示用户友好的错误信息
}
```

## 注意事项

1. **WebView2运行时**: 确保目标机器安装了WebView2运行时
2. **网络连接**: 访问外部网页需要网络连接
3. **用户数据目录**: 示例使用临时目录，生产环境应使用固定目录
4. **调试模式**: 生产版本可以关闭调试模式提高安全性
5. **资源管理**: 程序退出时会自动清理资源

这个简单示例是学习WebView2开发的理想起点，提供了最基础但完整的实现参考。