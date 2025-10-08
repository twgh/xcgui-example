# EnvironmentOptions - WebView2环境选项配置示例

## 项目简介

EnvironmentOptions 是一个专门演示如何配置 WebView2 环境选项的示例。该示例详细展示了 WebView2 的各种高级配置选项，包括语言设置、浏览器参数、多进程管理、扩展功能、滚动条样式等，为开发者提供了完整的环境配置参考。

## 功能特性

### 核心功能
- **环境选项配置**: 演示所有主要的WebView2环境配置选项
- **浏览器参数定制**: 自定义浏览器命令行参数
- **多版本接口使用**: 展示如何使用不同版本的环境选项接口
- **配置信息输出**: 实时显示当前配置的各项参数值

### 配置选项演示
- **语言设置**: 设置WebView2界面语言
- **自动播放策略**: 配置媒体自动播放相关参数
- **多进程管理**: 控制用户数据文件夹的访问权限
- **跟踪防护**: 控制隐私跟踪防护功能
- **浏览器扩展**: 启用/禁用浏览器扩展功能
- **滚动条样式**: 设置现代化的滚动条外观

### 技术特性
- **版本兼容检查**: 完整的WebView2版本检测和兼容性检查
- **错误处理**: 详细的错误处理和日志输出
- **配置验证**: 验证和显示所有配置项的当前值

## 项目结构

```
EnvironmentOptions/
├── EnvironmentOptions.go   # 主程序文件
└── README.md               # 本文档
```

## 技术实现

### 技术栈
- **Go 语言**: 主要编程语言  
- **XCGUI**: 炫彩界面库
- **WebView2**: Microsoft Edge 内核的嵌入式浏览器
- **WebView2 Environment Options**: WebView2环境配置接口

### 关键配置项详解

#### 1. 基础环境选项
```go
// 创建环境选项对象
envOpts, err := edge.CreateEnvironmentOptions()

// 设置界面语言
envOpts.SetLanguage("en-us")

// 设置浏览器命令行参数
envOpts.SetAdditionalBrowserArguments(sb.String())
```

#### 2. 自动播放配置
```go
sb := strings.Builder{}
// 允许无需用户交互的自动播放
sb.WriteString("--autoplay-policy=no-user-gesture-required ")
// 禁用媒体参与度检查，绕过自动播放策略
sb.WriteString("--disable-features=PreloadMediaEngagementData,MediaEngagementBypassAutoplayPolicies ")
// 忽略 Web Audio 的自动播放限制
sb.WriteString("--enable-features=AutoplayIgnoreWebAudio")
```

#### 3. 高级选项配置（环境选项2）
```go
envOpts2, err := envOpts.GetICoreWebView2EnvironmentOptions2()
// 设置其他进程可以从使用相同用户数据文件夹创建的WebView2环境创建WebView2
envOpts2.SetExclusiveUserDataFolderAccess(true)
```

#### 4. 隐私和安全配置（环境选项5）
```go
envOpts5, err := envOpts.GetICoreWebView2EnvironmentOptions5()
// 禁用跟踪防护功能以提高运行时性能
// 注意：仅在呈现已知安全内容时才应禁用
envOpts5.SetEnableTrackingPrevention(false)
```

#### 5. 浏览器扩展配置（环境选项6）
```go
envOpts6, err := envOpts.GetICoreWebView2EnvironmentOptions6()
// 启用浏览器扩展功能
envOpts6.SetAreBrowserExtensionsEnabled(true)
```

#### 6. 用户界面配置（环境选项8）
```go
envOpts8, err := envOpts.GetICoreWebView2EnvironmentOptions8()
// 设置滚动条样式为流畅覆盖样式
envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY)
```

### 配置信息输出
程序会实时输出所有配置项的当前值：
```go
fmt.Println("------------------- WebView2 环境选项 -------------------")
fmt.Println("语言:", envOpts.MustGetLanguage())
fmt.Println("命令行参数:", envOpts.MustGetAdditionalBrowserArguments())
fmt.Println("多进程共享用户数据文件夹:", envOpts2.MustGetExclusiveUserDataFolderAccess())
fmt.Println("跟踪防护功能:", envOpts5.MustGetEnableTrackingPrevention())
fmt.Println("浏览器扩展功能:", envOpts6.MustGetAreBrowserExtensionsEnabled())
fmt.Println("滚动条样式:", envOpts8.MustGetScrollBarStyle())
```

## 使用方法

### 运行程序
1. 确保已安装 WebView2 运行时
2. 编译并运行程序：
   ```bash
   go run EnvironmentOptions.go
   ```
3. 查看控制台输出的配置信息
4. 程序会打开一个WebView窗口访问必应搜索

### 配置定制
开发者可以根据需要修改各项配置：

```go
// 修改语言设置
envOpts.SetLanguage("zh-cn") // 改为中文

// 修改滚动条样式
envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY)
// 或者使用默认样式
envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_DEFAULT)

// 启用/禁用跟踪防护
envOpts5.SetEnableTrackingPrevention(true) // 启用隐私保护
```

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测版本）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 配置选项详解

### 语言设置
- **作用**: 设置WebView2界面语言
- **可选值**: "en-us", "zh-cn", "ja-jp" 等标准语言代码
- **默认值**: 系统默认语言

### 自动播放策略
- **no-user-gesture-required**: 允许无用户交互的自动播放
- **disable-features**: 禁用特定功能以绕过播放限制
- **enable-features**: 启用特定功能支持

### 滚动条样式
- **FLUENT_OVERLAY**: 现代化的半透明覆盖样式
- **DEFAULT**: 传统的系统默认样式

### 扩展功能
- **启用**: 支持安装和使用浏览器扩展
- **禁用**: 提高安全性，减少攻击面

## 最佳实践

### 性能优化
```go
// 对于内容可控的应用，可以禁用跟踪防护以提升性能
envOpts5.SetEnableTrackingPrevention(false)

// 使用现代化滚动条样式提升用户体验
envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY)
```

### 安全考虑
```go
// 对于需要隐私保护的应用，应启用跟踪防护
envOpts5.SetEnableTrackingPrevention(true)

// 对于企业应用，可能需要禁用扩展功能
envOpts6.SetAreBrowserExtensionsEnabled(false)
```

### 用户体验
```go
// 设置用户首选语言
envOpts.SetLanguage("zh-cn")

// 配置媒体自动播放以改善用户体验
sb.WriteString("--autoplay-policy=no-user-gesture-required")
```

## 扩展应用

这个配置示例可以应用于：
- **企业内部应用**: 定制化的企业级WebView配置
- **多媒体应用**: 需要特殊媒体播放配置的应用
- **国际化应用**: 需要多语言支持的应用
- **高性能应用**: 需要性能优化的应用
- **安全敏感应用**: 需要特殊安全配置的应用

## 注意事项

1. **版本兼容性**: 不同的环境选项接口需要相应的WebView2版本支持
2. **性能影响**: 某些配置可能影响性能，需要权衡
3. **安全考虑**: 禁用安全功能前要确保内容安全可控
4. **用户数据目录**: 生产环境应使用固定的用户数据目录
5. **命令行参数**: 浏览器参数需要谨慎使用，避免不兼容问题

这个示例为开发者提供了WebView2环境配置的完整参考，帮助创建更加定制化和优化的WebView应用。