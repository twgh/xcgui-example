# EnvironmentOptions - WebView2 环境选项配置示例

> 本文由 AI 生成

## 项目简介

本示例演示如何使用 WebView2 环境选项（Environment Options）创建定制化的 WebView2 环境。展示了两种配置方式，包括语言设置、浏览器参数、多进程管理、扩展功能、滚动条样式等多种配置选项。

## 功能特性

### 核心功能
- **两种配置方式**：提供自动配置和手动配置两种创建环境的方式
- **环境选项配置**：演示所有主要的 WebView2 环境配置选项
- **浏览器参数定制**：自定义浏览器命令行参数
- **配置信息输出**：实时显示当前配置的各项参数值

### 配置选项演示
- **语言设置**：设置 WebView2 界面语言
- **自动播放策略**：配置媒体自动播放相关参数
- **多进程管理**：控制用户数据文件夹的访问权限
- **跟踪防护**：控制隐私跟踪防护功能
- **浏览器扩展**：启用/禁用浏览器扩展功能
- **频道搜索**：选择 WebView2 频道搜索类型
- **滚动条样式**：设置现代化的滚动条外观

## 技术实现

### 技术栈
- **Go 语言**：主要编程语言
- **XCGUI**：炫彩界面库
- **WebView2**：Microsoft Edge 内核的嵌入式浏览器
- **WebView2 Environment Options**：WebView2 环境配置接口

### 两种创建方式

#### 方式一：自动配置（createEdge1）

使用 `edge.EnvOptions` 结构体自动配置环境选项：

```go
edg, err := edge.New(edge.Option{
    UserDataFolder: os.TempDir(),
    EnvOptions: &edge.EnvOptions{
        Language: "en-us",
        AdditionalBrowserArguments: []string{
            "--autoplay-policy=no-user-gesture-required",
            "--disable-features=PreloadMediaEngagementData,MediaEngagementBypassAutoplayPolicies",
            "--enable-features=AutoplayIgnoreWebAudio",
        },
        ExclusiveUserDataFolderAccess: true,
        DisableTrackingPrevention:     true,
        AreBrowserExtensionsEnabled:   true,
        ChannelSearchKind:             edge.COREWEBVIEW2_CHANNEL_SEARCH_KIND_MOST_STABLE,
        ScrollBarStyle:                edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY,
        ReleaseChannels: &edge.ReleaseChannels{
            ReleaseChannels: edge.COREWEBVIEW2_RELEASE_CHANNELS_NONE,
        },
    },
})
```

**特点**：
- 简洁易用，一次性配置所有选项
- 不需要手动释放资源
- 如果某个选项设置失败，会导致整个环境创建失败
- 适用于配置稳定、不需要特殊错误处理的场景

#### 方式二：手动配置（createEdge2）

使用 `CreateEnvironmentOptions` 手动创建并配置环境选项：

```go
envOpts, err := edge.CreateEnvironmentOptions()
if err != nil {
    // 处理错误
}

// 设置基础选项
envOpts.SetLanguage("en-us")
envOpts.SetAdditionalBrowserArguments(args)

// 获取并设置高级选项（支持版本检测）
if envOpts2, err := envOpts.GetICoreWebView2EnvironmentOptions2(); err == nil {
    envOpts2.SetExclusiveUserDataFolderAccess(true)
    envOpts2.Release()
}

// ... 设置其他版本的选项

edg, err := edge.New(edge.Option{
    UserDataFolder:     os.TempDir(),
    EnvironmentOptions: envOpts,
})
```

**特点**：
- 可以精确控制每个选项的设置
- 支持版本检测，不同版本的选项可以分别处理
- 设置失败不会导致整个环境创建失败
- 需要手动释放资源（调用 Release 方法）
- 适用于需要灵活处理不同版本和错误的场景

### 关键配置项说明

#### 1. 语言设置
```go
envOpts.SetLanguage("en-us") // 或 "zh-cn", "ja-jp" 等
```

#### 2. 自动播放配置
```go
AdditionalBrowserArguments: []string{
    "--autoplay-policy=no-user-gesture-required",         // 允许无需用户交互的自动播放
    "--disable-features=PreloadMediaEngagementData,MediaEngagementBypassAutoplayPolicies", // 禁用媒体参与度检查
    "--enable-features=AutoplayIgnoreWebAudio",          // 忽略 Web Audio 的自动播放限制
}
```

#### 3. 多进程共享用户数据文件夹
```go
ExclusiveUserDataFolderAccess: true
// 其他进程可以从使用相同用户数据文件夹创建的 WebView2 环境创建 WebView2
// 从而共享同一个 WebView 浏览器进程实例
```

#### 4. 跟踪防护
```go
DisableTrackingPrevention: true
// 禁用跟踪防护功能以提高运行时性能
// 仅在呈现已知安全内容时才应禁用
```

#### 5. 浏览器扩展
```go
AreBrowserExtensionsEnabled: true // 启用浏览器扩展功能
```

#### 6. 频道搜索类型
```go
ChannelSearchKind: edge.COREWEBVIEW2_CHANNEL_SEARCH_KIND_MOST_STABLE
// 选择最稳定的频道版本
```

#### 7. 滚动条样式
```go
ScrollBarStyle: edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY
// 使用现代化的流畅覆盖样式
```

#### 8. 发布频道
```go
ReleaseChannels: &edge.ReleaseChannels{
    ReleaseChannels: edge.COREWEBVIEW2_RELEASE_CHANNELS_NONE,
}
```

### 配置信息输出

程序会在创建环境后输出当前配置信息：

```go
fmt.Println("------------------- WebView2 环境选项 -------------------")
fmt.Println("语言:", envOpts.MustGetLanguage())
fmt.Println("命令行参数:", envOpts.MustGetAdditionalBrowserArguments())
fmt.Println("多进程共享用户数据文件夹:", envOpts2.MustGetExclusiveUserDataFolderAccess())
fmt.Println("跟踪防护功能:", envOpts5.MustGetEnableTrackingPrevention())
fmt.Println("浏览器扩展功能:", envOpts6.MustGetAreBrowserExtensionsEnabled())
fmt.Println("频道搜索类型:", envOpts7.MustGetChannelSearchKind())
fmt.Println("发布频道:", envOpts7.MustGetReleaseChannels())
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
4. 程序会打开一个 WebView 窗口访问 Bing 搜索

### 切换配置方式

在 `main()` 函数中切换注释即可选择不同的创建方式：

```go
edg := createEdge1() // 使用自动配置方式
// edg := createEdge2() // 使用手动配置方式
```

### 配置定制

#### 使用自动配置方式
```go
EnvOptions: &edge.EnvOptions{
    Language: "zh-cn",  // 修改为中文
    // ... 其他配置
}
```

#### 使用手动配置方式
```go
envOpts.SetLanguage("zh-cn")  // 修改为中文
// ... 其他配置
```

## 两种方式的选择建议

### 选择自动配置方式（createEdge1）的场景：
- 配置相对简单，不需要复杂的版本兼容性处理
- 希望代码简洁，不需要手动管理资源释放
- 对错误处理要求不高，任何配置失败都直接退出
- 环境版本已知且统一，不存在版本差异

### 选择手动配置方式（createEdge2）的场景：
- 需要支持多个版本的 WebView2 运行时
- 希望在某个配置失败时继续运行（降级处理）
- 需要精细控制错误处理逻辑
- 需要在运行时动态调整配置项
- 需要查看或验证配置项的值

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测版本）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 配置选项详解

### 语言设置
- **作用**：设置 WebView2 界面语言
- **可选值**："en-us", "zh-cn", "ja-jp" 等标准语言代码
- **默认值**：系统默认语言

### 自动播放策略
- **no-user-gesture-required**：允许无用户交互的自动播放
- **disable-features**：禁用特定功能以绕过播放限制
- **enable-features**：启用特定功能支持

### 滚动条样式
- **FLUENT_OVERLAY**：现代化的半透明覆盖样式
- **DEFAULT**：传统的系统默认样式

### 频道搜索类型
- **MOST_STABLE**：选择最稳定的频道版本
- **其他选项**：根据 WebView2 版本支持情况而定

### 扩展功能
- **启用**：支持安装和使用浏览器扩展
- **禁用**：提高安全性，减少攻击面

## 最佳实践

### 性能优化
```go
// 对于内容可控的应用，可以禁用跟踪防护以提升性能
DisableTrackingPrevention: true

// 使用现代化滚动条样式提升用户体验
ScrollBarStyle: edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY
```

### 安全考虑
```go
// 对于需要隐私保护的应用，应启用跟踪防护
DisableTrackingPrevention: false

// 对于企业应用，可能需要禁用扩展功能
AreBrowserExtensionsEnabled: false
```

### 用户体验
```go
// 设置用户首选语言
Language: "zh-cn"

// 配置媒体自动播放以改善用户体验
AdditionalBrowserArguments: []string{
    "--autoplay-policy=no-user-gesture-required",
}
```

## 注意事项

1. **版本兼容性**：不同的环境选项接口需要相应的 WebView2 版本支持
2. **方式选择**：根据实际需求选择合适的配置方式
3. **资源释放**：使用手动配置方式时，务必调用 Release 方法释放资源
4. **性能影响**：某些配置可能影响性能，需要权衡
5. **安全考虑**：禁用安全功能前要确保内容安全可控
6. **用户数据目录**：生产环境应使用固定的用户数据目录，而不是临时目录
7. **命令行参数**：浏览器参数需要谨慎使用，避免不兼容问题

## 扩展应用

这个配置示例可以应用于：
- **企业内部应用**：定制化的企业级 WebView 配置
- **多媒体应用**：需要特殊媒体播放配置的应用
- **国际化应用**：需要多语言支持的应用
- **高性能应用**：需要性能优化的应用
- **安全敏感应用**：需要特殊安全配置的应用
