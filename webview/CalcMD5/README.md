# CalcMD5 - 文件 MD5 计算器

> 本文由 AI 生成

## 项目简介

CalcMD5 是一个基于 XCGUI 和 WebView2 技术的文件 MD5 哈希值计算工具。这个示例演示了如何使用 Go 语言创建一个现代化的桌面应用程序，结合 Web 前端技术实现用户界面，展示圆角窗口、文件拖拽、异步计算等高级特性。

## 功能特性

### 核心功能
- **文件 MD5 计算**：快速计算选定文件的 MD5 哈希值
- **文件选择**：通过点击按钮选择文件进行计算
- **拖拽支持**：支持直接拖拽文件到应用窗口进行计算
- **实时显示**：在界面上实时显示计算结果和进度

### 界面特性
- **圆角窗口**：8px 圆角的现代化窗口设计
- **自定义标题栏**：完全自定义的 HTML 标题栏，包含窗口控制按钮
- **紧凑设计**：550x500 的紧凑窗口大小
- **美观界面**：使用 HTML/CSS 实现的现代化用户界面
- **自动焦点**：首次加载完成后自动聚焦输入框

### 技术特性
- **异步处理**：MD5 计算在后台协程中执行，不阻塞界面
- **消息循环等待**：通过消息循环机制等待协程完成，保持界面响应
- **嵌入式资源**：前端资源（HTML、CSS、JS）嵌入到可执行文件中
- **虚拟主机映射**：使用虚拟主机名访问嵌入的文件系统
- **环境选项优化**：禁用跟踪防护、现代化滚动条样式

## 项目结构

```
CalcMD5/
├── CalcMD5.go         # 主程序文件
├── 1.jpg              # 窗口界面展示图片
├── assets/            # 前端资源目录（会被嵌入）
│   ├── CalcMD5.html   # 主界面 HTML
│   ├── CalcMD5.css    # 样式文件
│   └── CalcMD5.js     # JavaScript 逻辑
└── README.md          # 本文档
```

## 技术实现

### 后端技术栈
- **Go 语言**：主要编程语言
- **XCGUI**：GUI 框架，提供窗口和 UI 组件
- **WebView2**：嵌入式浏览器组件
- **embed**：Go 1.16+ 的文件嵌入功能

### 前端技术栈
- **HTML5**：页面结构
- **CSS3**：样式设计
- **JavaScript**：交互逻辑
- **WebView2 API**：Go 与 JavaScript 之间的通信桥梁

### 关键技术点

#### 1. 文件嵌入

使用 Go 的 `embed` 包将前端资源嵌入到可执行文件中：

```go
//go:embed assets/**
embedAssets embed.FS // 嵌入 assets 目录及其所有子目录文件，不包括隐藏文件
```

#### 2. 虚拟主机映射

设置虚拟主机名与嵌入文件系统的映射：

```go
const hostName = "app.example"

// 启用虚拟主机名和嵌入文件系统之间的映射
m.wv.EnableVirtualHostNameToEmbedFSMapping(true)

// 设置虚拟主机名和嵌入文件系统之间的映射
edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
```

#### 3. 圆角窗口

通过两个配置实现 8px 圆角窗口：

```go
// 炫彩 XML 窗口阴影圆角大小，设置后会使窗口变为圆角
edge.WithXmlWindowShadowAngleSize(8)

// WebView 圆角 8px
edge.WithRoundRadius(8)
```

#### 4. 自定义标题栏

使用 HTML 实现自定义标题栏，并通过绑定 Go 函数实现窗口控制：

```html
<div id="title-bar">
    <div>MD5 计算器</div>
    <div id="window-controls">
        <button class="window-btn" onclick="go.minimizeWindow()">−</button>
        <button class="window-btn" onclick="go.toggleMaximize()">□</button>
        <button class="window-btn close" onclick="go.closeWindow()">×</button>
    </div>
</div>
```

Go 端绑定函数：
```go
m.wv.Bind("go.minimizeWindow", func() {
    m.w.ShowWindow(xcc.SW_MINIMIZE)
})

m.wv.Bind("go.toggleMaximize", func() {
    m.w.MaxWindow(!m.w.IsMaxWindow())
})

m.wv.Bind("go.closeWindow", func() {
    m.w.CloseWindow()
})
```

#### 5. 文件选择

使用 `wutil.OpenFile` 实现文件选择功能：

```go
m.wv.Bind("go.openFile", func() string {
    return wutil.OpenFile(m.w.Handle, []string{"All Files(*.*)", "*.*"}, "")
})
```

JavaScript 端调用：
```javascript
async function selectFile() {
    await calculate(await go.openFile());
}
```

#### 6. 拖拽文件处理

通过 WebView2 的 `postMessageWithAdditionalObjects` API 实现拖拽文件功能：

JavaScript 端：
```javascript
dragArea.addEventListener('drop', (e) => {
    e.preventDefault();
    // 通过 WebView2 发送文件作为附加对象
    if (window.chrome && chrome.webview) {
        chrome.webview.postMessageWithAdditionalObjects('drag_files', e.dataTransfer.files);
    }
});
```

Go 端处理：
```go
m.wv.Event_WebMessageReceived(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2WebMessageReceivedEventArgs) uintptr {
    webMessage, _ := args.TryGetWebMessageAsString()
    if webMessage != "drag_files" {
        return 0
    }

    args2, _ := args.GetICoreWebView2WebMessageReceivedEventArgs2()
    defer args2.Release()

    // 获取附加对象集合
    objs, _ := args2.GetAdditionalObjects()
    defer objs.Release()

    objCount, _ := objs.GetCount()
    for i := uint32(0); i < objCount; i++ {
        obj, _ := objs.GetValueAtIndex(i)

        file := new(edge.ICoreWebView2File)
        obj.QueryInterface(wapi.NewGUIDPointer(edge.IID_ICoreWebView2File), unsafe.Pointer(&file))
        defer file.Release()

        // 获取文件路径并计算 MD5
        filePath, _ := file.GetPath()
        m.wv.Eval(`calculate('` + strings.ReplaceAll(filePath, `\`, `\\`) + `');`)
        break // 目前只处理第一个文件
    }
    return 0
})
```

#### 7. 异步 MD5 计算

在 Go 协程中执行耗时的 MD5 计算，通过消息循环等待完成：

```go
m.wv.Bind("go.calculateMD5", func(filePath string) string {
    // 判断文件是否存在
    if !xc.PathExists2(filePath) {
        return "错误: 文件不存在"
    }

    var ret string
    // 耗时操作在协程里执行，不卡界面
    go func() {
        data, err := os.ReadFile(filePath)
        if err != nil {
            ret = "错误: " + err.Error()
            return
        }

        hash := md5.Sum(data)
        md5Str := hex.EncodeToString(hash[:])
        ret = "文件: " + filePath + "\nMD5: " + md5Str
    }()

    // 等待 MD5 计算完成，通过消息循环保持界面响应
    var msg wapi.MSG
    for ret == "" {
        if wapi.GetMessage(&msg, 0, 0, 0) == 0 {
            break
        }
        wapi.TranslateMessage(&msg)
        wapi.DispatchMessage(&msg)
    }
    return ret
})
```

#### 8. 渐进式加载和自动焦点

等待页面完全加载后再显示窗口，避免白屏闪烁：

```go
var firstLoad = true

m.wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
    uri := sender.MustGetSource()

    switch uri {
    case edge.JoinUrlHeader(hostName) + "/CalcMD5.html":
        if firstLoad {
            firstLoad = false
            m.w.Show(true)
            // 使 HTML 中的输入框获取焦点
            m.wv.Eval(`document.getElementById('filePath').focus()`)
        }
    }
    return 0
})
```

#### 9. 环境选项优化

优化 WebView2 性能和体验：

```go
// 禁用跟踪防护以提高运行时性能（仅在呈现已知安全内容时）
envOpts5.SetEnableTrackingPrevention(false)

// 设置现代化滚动条样式
envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY)
```

## 使用方法

### 运行程序

1. 确保已安装 WebView2 运行时
2. 运行程序：
   ```bash
   go run CalcMD5.go
   ```

### 计算文件 MD5

**方法一**：通过按钮选择文件
1. 点击"选择文件"按钮
2. 在弹出的文件对话框中选择文件
3. 等待计算完成，查看结果

**方法二**：拖拽文件
1. 将文件拖拽到应用程序窗口中
2. 程序会自动计算并显示文件的 MD5 哈希值

### 调试模式

设置 `isDebug = true` 可以启用：
- 开发者工具
- 右键菜单
- 浏览器快捷键

```go
var isDebug = false
```

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测并提示下载）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 设计亮点

### 用户体验
- **无闪烁启动**：等待页面完全加载后再显示窗口，避免白屏闪烁
- **实时反馈**：计算过程中显示"计算中..."状态，完成后显示结果
- **自动焦点**：启动后自动聚焦到输入框，方便直接拖拽文件
- **紧凑设计**：550x500 的窗口大小，适合工具类应用

### 技术创新
- **圆角窗口**：通过 `WithXmlWindowShadowAngleSize` 和 `WithRoundRadius` 实现 8px 圆角
- **拖拽文件**：使用 WebView2 的 `postMessageWithAdditionalObjects` API 实现文件拖拽
- **异步计算**：协程 + 消息循环机制，既不阻塞界面又能等待计算完成
- **资源一体化**：将所有前端资源嵌入可执行文件，简化部署
- **双圆角一致**：WebView2 和外层窗口都设置 8px 圆角，保持视觉一致性

## 扩展可能性

这个示例可以作为基础，扩展为：
- **多种哈希算法支持**：SHA1、SHA256、SHA512 等
- **批量文件处理**：一次选择多个文件批量计算
- **哈希值比较**：比较两个文件的哈希值是否相同
- **文件完整性验证**：验证文件是否被修改
- **更复杂的桌面应用程序**：各种需要文件处理的工具

## 最佳实践

### 异步处理
对于耗时操作（如 MD5 计算），应该：
- 在协程中执行，避免阻塞 UI 线程
- 使用消息循环等待协程完成
- 在界面上显示加载状态

### 文件拖拽
使用 `postMessageWithAdditionalObjects` API：
- 支持从文件管理器拖拽文件到窗口
- 可以获取文件的完整路径
- 支持多个文件拖拽（本示例只处理第一个）

### 窗口设计
- 使用渐进式加载避免白屏闪烁
- 自动焦点提升用户体验
- 圆角设计增加现代感
- 自定义标题栏实现统一风格

## 注意事项

1. **调试模式**：代码中的 `isDebug` 变量控制调试功能的开启，生产环境应设置为 `false`
2. **用户数据目录**：示例中使用临时目录，实际应用应使用固定的应用数据目录
3. **WebView2 版本**：程序会检查本地 WebView2 版本兼容性，某些功能需要较新版本
4. **文件大小**：对于大文件的 MD5 计算可能需要较长时间
5. **协程等待**：使用消息循环等待协程完成时，要确保消息循环正常退出
6. **路径转义**：将文件路径传给 JavaScript 时，需要转义反斜杠字符

## 常见问题

### Q: 为什么计算大文件时界面不卡？
A: 因为 MD5 计算在后台协程中执行，通过消息循环等待完成，不会阻塞 UI 线程。

### Q: 如何支持更多哈希算法？
A: 可以在 Go 端添加对应的哈希函数（如 `crypto/sha256`），然后在 JavaScript 中调用。

### Q: 为什么窗口是圆角的？
A: 通过 `WithXmlWindowShadowAngleSize(8)` 和 `WithRoundRadius(8)` 两个配置实现 8px 圆角。

### Q: 如何禁用调试功能？
A: 将 `isDebug` 变量设置为 `false` 即可。

### Q: 拖拽多个文件如何处理？
A: 修改 Go 端代码，不要在获取第一个文件后 `break`，而是遍历所有文件。

这个示例完美展示了 XCGUI 与 WebView2 结合开发现代桌面应用的优势，为开发者提供了一个实用的参考模板。
