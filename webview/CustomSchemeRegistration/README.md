# CustomSchemeRegistration - 注册自定义协议（Custom Scheme）

> 本文由 AI 生成

## 项目简介

CustomSchemeRegistration 是一个演示如何在 WebView2 中注册和使用自定义协议的高级示例。该示例展示了如何通过自定义协议实现 WebView2 与本地系统的交互，包括文件打开、系统调用等功能，是学习 WebView2 高级特性开发的理想起点。

## 功能特性

### 核心功能
- **自定义协议注册**: 注册 `myapp://` 自定义协议
- **导航事件拦截**: 拦截 WebView2 导航事件，处理自定义协议
- **URL 路由**: 根据不同的路径和参数执行不同的操作
- **文件打开**: 支持通过自定义协议打开本地文件
- **系统交互**: 与 Windows Shell API 交互，执行系统命令
- **JavaScript 通信**: 通过 JavaScript 更新页面内容

### 界面特性
- **简洁界面**: 简单的 HTML 页面作为示例
- **炫彩标题栏**: 保留炫彩窗口的标题栏和窗口控制按钮
- **自适应布局**: WebView 自动填充整个窗口客户区

### 技术特性
- **EnvOptions 使用**: 使用简化的 `EnvOptions` 配置 WebView2 环境
- **版本兼容检查**: 自动检测 WebView2 运行时版本兼容性
- **错误处理**: 完善的错误提示和异常处理机制
- **DPI 适配**: 支持高 DPI 显示器的自适应缩放
- **文件嵌入**: 使用 Go 1.16+ 的 embed 功能嵌入 HTML 资源

## 项目结构

```
CustomSchemeRegistration/
├── CustomSchemeRegistration.go  # 主程序文件
├── index.html               # 内嵌的 HTML 页面
└── README.md                # 本文档
```

## 技术实现

### 技术栈
- **Go 语言**: 主要编程语言
- **XCGUI**: 炫彩界面库，提供窗口和 UI 框架
- **WebView2**: 基于 Microsoft Edge 内核的嵌入式浏览器
- **embed**: Go 1.16+ 的文件嵌入功能
- **便捷 API**: 使用 `NewWebViewWithWindow()` 方法一步创建窗口和 WebView
- **EnvOptions**: 使用简化的环境选项配置

### 自定义协议说明

本示例注册了 `myapp://` 自定义协议，支持以下路径：

#### 1. `/openFile` - 打开文件
**格式**: `myapp://openFile?path=C:\path\to\file`

**功能**:

- 如果 `path` 参数为空，弹出文件选择对话框
- 如果 `path` 参数存在，使用系统默认程序打开指定文件
- 将选择的文件路径通过 JavaScript 更新到页面

**示例**:

```go
// 打开指定文件
myapp://openFile?path=C:\Windows\system.ini

// 弹出文件选择对话框
myapp://openFile
```

#### 2. `/showSettings` - 显示设置
**格式**: `myapp://showSettings`

**功能**: 显示提示对话框，模拟打开设置窗口

**示例**:

```go
myapp://showSettings
```

#### 3. `/navigateTo` - 页面跳转
**格式**: `myapp://navigateTo?page=home`

**功能**:

- 跳转到指定页面
- 如果 `page` 参数为空，默认跳转到 `home`
- 显示提示对话框

**示例**:
```go
myapp://navigateTo?page=home
myapp://navigateTo?page=settings
```

### 关键代码解析

#### 1. 自定义协议注册
```go
// 创建自定义方案
customScheme, err := edge.CreateCustomSchemeRegistration("myapp")
if err != nil {
    wapi.MessageBoxW(0, "创建自定义方案 myapp 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
    os.Exit(1)
}
defer customScheme.Release()

// 创建 WebView 环境时使用 EnvOptions
edg, err := edge.New(edge.Option{
    UserDataFolder: os.TempDir(),
    EnvOptions: &edge.EnvOptions{
        ExclusiveUserDataFolderAccess: true,  // 多进程共享用户数据
        EnableTrackingPrevention:        false,  // 禁用跟踪防护
        ScrollBarStyle:                edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY,
        CustomSchemeRegistrations:       []*edge.ICoreWebView2CustomSchemeRegistration{customScheme},
    },
})
```

**注意**:
- 本示例使用了简化的 `EnvOptions` 配置，更加简洁易读
- `CustomSchemeRegistrations` 字段用于注册自定义协议
- 其他配置项如 `ExclusiveUserDataFolderAccess`、`EnableTrackingPrevention` 等也一并设置

#### 2. WebView 和窗口创建
```go
// 创建 WebView 和窗口（一步完成）
w, wv, err := edg.NewWebViewWithWindow(
    edge.WithXmlWindowTitle("注册自定义方案"),
    edge.WithXmlWindowSize(600, 400),
    edge.WithXmlWindowTitleBar(true),       // 使用炫彩窗口标题栏
    edge.WithFillParent(true),             // WebView 填充窗口
    edge.WithDebug(true),                  // 启用调试模式
    edge.WithAutoFocus(true),              // 自动聚焦
)
```

#### 3. 导航事件拦截
```go
// 注册导航开始事件
wv.Event_NavigationStarting(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationStartingEventArgs) uintptr {
    // 获取当前导航 URI
    uri, err := args.GetUri()
    if err != nil {
        log.Println("NavigationStarting GetUri 失败:", err)
        return 0
    }
    fmt.Println("uri:", uri)

    // 检查是否是自定义协议
    if strings.HasPrefix(uri, "myapp://") {
        // 取消导航，防止 WebView 尝试加载该 URL
        err := args.SetCancel(true)
        if err != nil {
            log.Println("NavigationStarting SetCancel 失败:", err)
            return 0
        }

        // 解析 URL 并处理
        u, err := url.Parse(uri)
        if err != nil {
            log.Printf("URL 解析失败: %v\n", err)
            return 0
        }

        path := "/" + u.Host  // 路径，如 "/openFile"
        query := u.Query()       // 查询参数，如 map[string][]string

        // 根据路径执行不同操作
        switch path {
        case "/openFile":
            handleOpenFile(query)
        case "/showSettings":
            handleShowSettings()
        case "/navigateTo":
            handleNavigateTo(query)
        default:
            log.Printf("未知命令: %s\n", path)
        }
    }
    return 0
})
```

**关键点**:
- 使用 `strings.HasPrefix()` 检查是否是自定义协议
- 调用 `args.SetCancel(true)` 取消导航，防止错误
- 使用 `url.Parse()` 解析 URL，获取路径和查询参数
- 使用 `switch` 语句根据路径执行不同操作

#### 4. 文件打开处理
```go
case "/openFile":
    filePath := query.Get("path")
    if filePath != "" {
        // 使用系统默认程序打开指定文件
        wapi.ShellExecuteW(0, "open", filePath, "", "", xcc.SW_SHOW)
        return 0
    }

    // 弹出文件选择对话框
    filePath = wutil.OpenFile(w.Handle, []string{"All Files(*.*)", "*.*"}, "")

    fmt.Println("打开的文件路径:", filePath)
    filePath = strings.ReplaceAll(filePath, `\`, `\\`)  // 转义反斜杠

    // 通过 JavaScript 更新页面
    wv.Eval(`document.getElementById('filePath').value = '` + filePath + `'`)
```

**功能说明**:
- `wutil.OpenFile()` 弹出文件选择对话框
- `wapi.ShellExecuteW()` 使用系统默认程序打开文件
- `wv.Eval()` 执行 JavaScript 更新页面内容
- 需要转义路径中的反斜杠，避免 JavaScript 语法错误

#### 5. HTML 页面
```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>自定义协议示例</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            padding: 20px;
        }
        input {
            width: 100%;
            padding: 8px;
            margin: 10px 0;
            border: 1px solid #ccc;
            border-radius: 4px;
        }
        button {
            padding: 10px 20px;
            margin: 5px;
            cursor: pointer;
            background: #0078D4;
            color: white;
            border: none;
            border-radius: 4px;
        }
    </style>
</head>
<body>
    <h2>自定义协议示例</h2>
    <input type="text" id="filePath" readonly placeholder="选择的文件路径">
    <br>
    <button onclick="location.href='myapp://openFile'">选择文件</button>
    <button onclick="location.href='myapp://openFile?path=C:\\Windows\\system.ini'">打开system.ini</button>
    <button onclick="location.href='myapp://showSettings'">显示设置</button>
    <button onclick="location.href='myapp://navigateTo?page=profile'">跳转到个人资料页</button>
    <button onclick="location.href='myapp://navigateTo?page=order&id=123'">查看订单 #123</button>
</body>
</html>
```

**特点**:
- 简洁的 HTML 页面
- 使用 JavaScript 导航到自定义协议
- 展示了三种不同的自定义协议用法

## 使用方法

### 运行程序
1. 确保系统已安装 WebView2 运行时

2. 运行程序：
   ```bash
   go run CustomSchemeRegistration.go
   ```

3. 程序将打开窗口，显示 HTML 页面

### 使用自定义协议

1. **选择文件**: 点击"选择文件"按钮，弹出文件选择对话框，选择文件后会显示文件路径

2. **显示设置**: 点击"显示设置"按钮，弹出提示对话框

3. **跳转页面**: 点击"跳转页面"按钮，弹出提示对话框显示要跳转的页面名

4. **手动测试**: 在开发者工具控制台中手动输入：
   ```javascript
   location.href = 'myapp://openFile?path=C:\\Windows\\system.ini'
   location.href = 'myapp://navigateTo?page=settings'
   ```

### 开发者工具
- 在 WebView 区域右键可打开开发者工具
- 可以在 Console 中测试自定义协议
- 查看导航事件日志输出

## 应用场景

自定义协议在以下场景中非常有用：

### 1. 桌面应用包装器
将 Web 应用包装为桌面程序，通过自定义协议访问本地功能：
- 打开本地文件
- 显示系统对话框
- 调用本地 API
- 与系统集成

### 2. 混合应用
结合 Web UI 和本地功能：
- Web 页面提供用户界面
- 自定义协议桥接本地功能
- 保持 Web 技术栈的优势

### 3. 安全通信
通过自定义协议实现安全的本地通信：
- 避免直接暴露本地 API
- 受控的命令执行
- 参数验证和处理

### 4. 跨页面通信
在不同页面之间通过自定义协议通信：
- `myapp://open/settings`
- `myapp://share/data?param=value`
- `myapp://sync/local`

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测并提示安装）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 高级主题

### URL 参数解析

自定义协议支持查询参数，使用 `url.Parse()` 解析：

```go
u, _ := url.Parse("myapp://openFile?path=C:\\test.txt&mode=readonly")
query := u.Query()

path := query.Get("path")     // "C:\\test.txt"
mode := query.Get("mode")    // "readonly"
all := query["param"]        // []string{"value1", "value2"}
```

### JavaScript 与 Go 交互

```go
// 从 Go 向 JavaScript 传递数据
wv.Eval(`document.getElementById('result').textContent = '` + message + `'`)

// 执行 JavaScript 函数
wv.Eval(`updateContent('Hello from Go')`)
```

### 多命令扩展

可以通过 switch 语句轻松扩展更多命令：

```go
switch path {
case "/openFile":
    handleOpenFile(query)
case "/showSettings":
    handleShowSettings()
case "/navigateTo":
    handleNavigateTo(query)
case "/shareContent":
    handleShareContent(query)    // 新增
case "/syncData":
    handleSyncData(query)        // 新增
default:
    log.Printf("未知命令: %s\n", path)
}
```

### 错误处理

完善的错误处理是自定义协议的关键：

```go
// 检查必需参数
if path == "" {
    app.Alert("错误", "缺少 path 参数")
    return 0
}

// 验证参数
if !strings.HasPrefix(path, "C:\\") {
    app.Alert("错误", "路径必须以 C:\\ 开头")
    return 0
}

// 处理错误
err := wapi.ShellExecuteW(0, "open", filePath, "", "", xcc.SW_SHOW)
if err != nil {
    log.Printf("打开文件失败: %v\n", err)
    app.Alert("错误", "无法打开文件")
}
```

## 最佳实践

### 1. 协议命名
- 使用唯一的前缀，避免与系统协议冲突
- 使用小写字母和数字，如 `myapp://`、`mycompany://`
- 保持简短但有意义

### 2. 路径设计
- 使用清晰的路径结构：`/category/action`
- 支持层级路径：`/settings/theme/dark`
- 使用 RESTful 风格：`/api/v1/data`

### 3. 参数处理
- 始终验证参数的完整性和有效性
- 提供默认值
- 支持可选参数

### 4. 安全考虑
- 不要在自定义协议中执行任意代码
- 验证所有输入参数
- 限制可访问的资源

### 5. 用户反馈
- 为每个操作提供视觉反馈
- 使用对话框显示错误信息
- 记录操作日志

## 扩展可能性

基于这个示例可以扩展为：

### 1. 文件管理器
```go
case "/listFiles":
    listFiles(query.Get("path"))
case "/deleteFile":
    deleteFile(query.Get("path"))
case "/createFolder":
    createFolder(query.Get("name"))
```

### 2. 系统信息获取
```go
case "/getSystemInfo":
    info := getSystemInfo()
    wv.Eval(`updateSystemInfo('` + info + `')`)
```

### 3. 本地存储
```go
case "/saveData":
    saveData(query.Get("key"), query.Get("value"))
case "/loadData":
    data := loadData(query.Get("key"))
    wv.Eval(`updateData('` + data + `')`)
```

## 注意事项

1. **协议唯一性**: 确保自定义协议名不会与现有协议冲突
2. **参数验证**: 严格验证所有从 JavaScript 传递的参数
3. **错误处理**: 提供清晰的错误信息给用户
4. **调试模式**: 生产环境建议关闭调试模式
5. **用户数据目录**: 示例使用临时目录，实际应用应使用固定目录
6. **路径转义**: 注意路径中的特殊字符需要转义，避免 JavaScript 错误
7. **导航取消**: 记得在自定义协议处理中调用 `args.SetCancel(true)`

## 总结

本示例展示了如何使用 WebView2 的自定义协议功能，实现 Web 页面与本地系统的深度集成。通过自定义协议，可以将 Web 应用包装为功能完整的桌面应用，同时保持 Web 技术栈的优势。配合简化 EnvOptions 配置和 NewWebViewWithWindow() 便捷方法，代码更加简洁易读，是学习 WebView2 高级特性的理想起点。
