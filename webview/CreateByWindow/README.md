# CreateByWindow - 在炫彩窗口中创建WebView

> 本文由 AI 生成

## 项目简介

CreateByWindow 是一个演示如何在炫彩窗口中创建 WebView2 组件的基础示例。该示例使用了 `NewWebViewWithWindow()` 便捷方法，一步完成炫彩窗口和 WebView2 的创建，同时保留炫彩窗口的标题栏，展示了最简洁的 WebView2 集成方式。

## 功能特性

### 核心功能
- **基础WebView创建**: 在炫彩窗口中创建WebView2实例
- **网页浏览**: 可访问任意网页（示例中访问 vben.pro）
- **调试支持**: 启用开发者工具，便于调试和开发

### 界面特性
- **炫彩标题栏**: 保留原生炫彩窗口的标题栏和窗口控制按钮
- **自定义标题栏颜色**: 支持通过 API 设置标题栏背景颜色
- **自适应布局**: WebView自动填充整个窗口客户区

### 技术特性
- **版本兼容检查**: 自动检测WebView2运行时版本兼容性
- **错误处理**: 完善的错误提示和异常处理机制
- **DPI适配**: 支持高DPI显示器的自适应缩放

## 项目结构

```
CreateByWindow/
├── CreateByWindow.go     # 主程序文件
└── README.md            # 本文档
```

**注意**: 本示例使用 `NewWebViewWithWindow()` 便捷方法创建窗口，不需要额外的 XML 布局文件或资源文件。

## 技术实现

### 技术栈
- **Go 语言**: 主要编程语言
- **XCGUI**: 炫彩界面库，提供窗口和UI框架
- **WebView2**: 基于 Microsoft Edge 内核的嵌入式浏览器
- **embed**: Go 1.16+ 的文件嵌入功能
- **便捷API**: 使用 `NewWebViewWithWindow()` 方法一步创建窗口和 WebView

### 关键代码解析

#### 1. WebView2环境创建
```go
// 创建 WebView 环境
edg, err := edge.New(edge.Option{
    UserDataFolder: os.TempDir(), // 实际应用应使用固定目录
})
```

#### 2. WebView和窗口创建（使用便捷方法）
```go
// 创建 WebView 和窗口（一步完成）
w, wv, err := edg.NewWebViewWithWindow(
    edge.WithXmlWindowTitle("在窗口中创建 WebView, 仍然使用炫彩窗口的标题栏"),
    edge.WithXmlWindowTitleBar(true),                    // 使用炫彩窗口标题栏
    edge.WithXmlWindowTitleBarBgColor(xc.RGBA(17, 17, 26, 255)), // 设置标题栏背景颜色
    edge.WithXmlWindowSize(1400, 900),                // 窗口大小
    edge.WithFillParent(true),                          // WebView 填充窗口
    edge.WithDebug(true),                               // 启用调试模式
    edge.WithAutoFocus(true),                           // 自动聚焦
)
```

**注意**:
- 本示例使用 `NewWebViewWithWindow()` 便捷方法，该方法一步完成窗口和 WebView 的创建
- `WithXmlWindowTitleBar(true)` 保留炫彩窗口的标题栏和窗口控制按钮
- `WithXmlWindowTitleBarBgColor()` 可以自定义标题栏的背景颜色
- 如需在现有窗口中嵌入 WebView，可以使用 `NewWebView()` 方法

#### 3. 版本兼容性检查
```go
func checkWebView2() {
    fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())
    localVersion, err := edge.GetAvailableBrowserVersion()
    // 检查版本兼容性...
}
```

### 设计亮点

#### 标题栏颜色自定义
- 使用 `WithXmlWindowTitleBarBgColor()` 设置标题栏背景颜色
- 可根据应用主题调整颜色，提升视觉一致性
- 也可以在运行时通过 `SetXmlWindowTitleBarBgColor()` 动态修改

#### 版本管理机制
- **版本检测**: 自动检测本机WebView2运行时版本
- **兼容性警告**: 当版本不匹配时给出明确提示
- **自动下载**: 未安装时可自动打开下载页面

## 使用方法

### 运行程序
1. 确保系统已安装 WebView2 运行时
2. 运行程序：
   ```bash
   go run CreateByWindow.go
   ```
3. 程序将自动打开窗口并导航到指定网页

### 开发者工具
- 在WebView区域右键可打开开发者工具
- 支持标准的Web调试功能
- 可以检查元素、调试JavaScript等

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测并提示安装）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 配置说明

### 关键配置项
```go
// 用户数据目录
UserDataFolder: os.TempDir() // 建议改为专用目录

// 窗口配置
edge.WithXmlWindowTitle("窗口标题")       // 设置窗口标题
edge.WithXmlWindowSize(1400, 900)       // 设置窗口大小
edge.WithXmlWindowTitleBar(true)          // 启用炫彩窗口标题栏
edge.WithXmlWindowTitleBarBgColor(...)     // 设置标题栏背景颜色

// WebView 配置
edge.WithFillParent(true)                // WebView 填充整个窗口
edge.WithDebug(true)                     // 启用调试模式（生产环境建议设为 false）
edge.WithAutoFocus(true)                 // 自动聚焦
```

## 扩展可能性

这个基础示例可以扩展为：
- **企业内网应用**: 嵌入内部Web系统
- **混合桌面应用**: 结合炫彩UI和Web技术
- **在线文档查看器**: 集成各种在线服务
- **Web应用包装器**: 将Web应用包装为桌面程序

## 注意事项

1. **用户数据目录**: 示例中使用临时目录，实际应用应使用固定目录存储用户数据
2. **调试模式**: 生产环境中应关闭调试模式以提高安全性
3. **错误处理**: 程序包含完善的错误处理，但建议根据实际需求调整
4. **网络连接**: 需要网络连接才能正常访问外部网页

这个示例为初学者提供了使用 `NewWebViewWithWindow()` 便捷方法的 WebView2 集成方案，展示了如何一步完成窗口和 WebView 的创建，同时保留炫彩窗口的标题栏功能，是学习 XCGUI + WebView2 开发的理想起点。