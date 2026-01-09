# EmbedAssets - 嵌入式资源 WebView 应用

> 本文由 AI 生成

## 项目简介

EmbedAssets 是一个演示如何将前端资源嵌入到可执行文件中的 WebView 应用示例。该示例展示了如何实现单文件部署，无需启动 HTTP 服务器，通过虚拟主机名映射的方式访问嵌入的 HTML、CSS、JavaScript 等前端资源。

## 功能特性

### 核心功能
- **嵌入式资源**：将所有前端资源打包到可执行文件中
- **虚拟主机映射**：通过虚拟主机名访问嵌入的文件系统
- **单文件部署**：整个应用打包为单个可执行文件
- **开发/生产双模式**：支持开发时的热重载和生产时的资源嵌入

### 界面特性
- **自定义标题栏**：完全自定义的 HTML 标题栏，包含窗口控制按钮
- **现代化界面**：基于 HTML/CSS 的现代化用户界面（左侧导航栏 + 主内容区）
- **无闪烁加载**：通过延迟显示窗口避免白屏闪烁
- **窗口控制**：支持最小化、最大化/还原、关闭操作

### 技术特性
- **Go 与 JavaScript 交互**：支持 Go 函数和 JavaScript 之间的双向调用
- **环境选项配置**：禁用跟踪防护、设置现代化滚动条样式
- **资源热重载**：开发模式下支持前端资源的实时更新（通过文件夹映射）
- **渐进式加载**：等待页面完全加载后再显示窗口

## 项目结构

```
EmbedAssets/
├── EmbedAssets.go        # 主程序文件
├── 1.jpg                # 窗口界面展示图片
├── assets/              # 前端资源目录（会被嵌入）
│   ├── EmbedAssets.html # 主页面文件
│   ├── css/             # 样式文件目录
│   │   └── EmbedAssets.css
│   └── js/              # JavaScript 文件目录
│       └── EmbedAssets.js
└── README.md            # 本文档
```

## 技术实现

### 技术栈
- **Go 语言**：主要编程语言
- **XCGUI**：炫彩界面库
- **WebView2**：嵌入式浏览器组件
- **embed**：Go 1.16+ 文件嵌入功能
- **HTML/CSS/JavaScript**：前端技术栈（引用 Font Awesome 图标库）

### 关键技术点

#### 1. 文件嵌入

使用 Go 的 embed 包嵌入前端资源：

```go
//go:embed assets/**
embedAssets embed.FS // 嵌入 assets 目录及其所有子目录文件，不包括隐藏文件
```

**注意**：如果需要包含隐藏文件（以 `.` 或 `_` 开头的文件和目录），可以使用：
```go
//go:embed all:assets/**
```

#### 2. 虚拟主机映射

通过虚拟主机名访问嵌入的文件系统：

```go
const hostName = "app.example"
```

**生产模式**（`isDebug = false`）：映射嵌入文件系统
```go
// 启用虚拟主机名和嵌入文件系统之间的映射
err = m.wv.EnableVirtualHostNameToEmbedFSMapping(true)

// 设置虚拟主机名和嵌入文件系统之间的映射
err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
```

**开发模式**（`isDebug = true`）：映射本地文件夹，方便热重载
```go
folderPath, _ := filepath.Abs("webview/EmbedAssets/assets")
err = m.wv.SetVirtualHostNameToFolderMapping(hostName,
    folderPath, edge.COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW)
```

#### 3. 环境配置优化

优化 WebView2 性能和体验：

```go
// 禁用跟踪防护以提高运行时性能（仅在呈现已知安全内容时）
envOpts5.SetEnableTrackingPrevention(false)

// 设置现代化滚动条样式
envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY)
```

#### 4. Go 与 JavaScript 交互

绑定 Go 函数供 JavaScript 调用：

```go
// 绑定最小化窗口函数
m.wv.Bind("go.minimizeWindow", func() {
    m.w.ShowWindow(xcc.SW_MINIMIZE)
})

// 绑定切换最大化窗口函数
m.wv.Bind("go.toggleMaximize", func() {
    m.w.MaxWindow(!m.w.IsMaxWindow())
})

// 绑定关闭窗口函数
m.wv.Bind("go.closeWindow", func() {
    m.w.CloseWindow()
})
```

在 HTML 中调用：
```html
<button onclick="go.minimizeWindow()">最小化</button>
<button onclick="go.toggleMaximize()">最大化</button>
<button onclick="go.closeWindow()">关闭</button>
```

#### 5. 渐进式加载（避免白屏闪烁）

在导航完成事件中判断首次加载完毕时才显示窗口：

```go
var firstLoad = true

m.wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
    uri := sender.MustGetSource()

    switch uri {
    case edge.JoinUrlHeader(hostName) + "/EmbedAssets.html":
        if firstLoad {
            firstLoad = false
            m.w.Show(true) // 首次加载完成后才显示窗口
        }
    }
    return 0
})
```

**原因**：采用嵌入文件系统的方式时，网页还没加载出来时会显示 WebView 白色的背景，然后才会加载出网页，表现出来的就是有一瞬间的闪烁。通过延迟显示窗口可以解决这个问题。

### 设计模式

#### MainWindow 结构体模式

采用面向对象的设计模式，将相关功能封装在 MainWindow 结构体中：

```go
type MainWindow struct {
    edg *edge.Edge      // WebView2 环境
    w   *window.Window  // 炫彩窗口
    wv  *edge.WebView   // WebView 实例
}
```

创建窗口时使用 `NewWebViewWithWindow` 方法：

```go
m.w, m.wv, err = m.edg.NewWebViewWithWindow(
    edge.WithXmlWindowTitle("我的应用"),
    edge.WithXmlWindowClassName("EmbedAssets"),
    edge.WithXmlWindowSize(1300, 900),
    edge.WithFillParent(true),
    edge.WithAppDrag(true),          // 启用应用拖拽
    edge.WithDebug(isDebug),         // 调试模式
    edge.WithDefaultContextMenus(isDebug), // 控制右键菜单
    edge.WithBrowserAcceleratorKeys(isDebug), // 控制浏览器快捷键
    edge.WithStatusBar(false),       // 禁用状态栏
    edge.WithZoomControl(false),     // 禁用缩放控制
    edge.WithAutoFocus(true),
)
```

## 使用方法

### 开发模式（isDebug = true）

1. 设置 `var isDebug = true`
2. 运行程序：
   ```bash
   go run EmbedAssets.go
   ```
3. 修改 `assets/` 目录下的文件，支持热重载

**特点**：
- 使用本地文件夹映射，支持实时修改和刷新
- 启用调试菜单和快捷键
- 方便开发和调试

### 生产模式（isDebug = false）

1. 设置 `var isDebug = false`
2. 编译程序：
   ```bash
   go build -o myapp.exe EmbedAssets.go
   ```
3. 所有资源已嵌入，可直接分发 exe 文件

**特点**：
- 使用嵌入文件系统，无需外部资源文件
- 禁用调试功能，提升用户体验
- 单文件部署，简化分发流程

### 前端开发

- HTML 文件位于 `assets/EmbedAssets.html`
- CSS 文件位于 `assets/css/EmbedAssets.css`
- JavaScript 文件位于 `assets/js/EmbedAssets.js`
- 通过虚拟主机名 `app.example` 访问资源
- 可以使用任何前端框架和技术栈

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测版本）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 配置说明

### 调试开关

```go
var isDebug = false // 控制调试/生产模式
```

- `true`：开发模式，使用文件夹映射，启用调试功能
- `false`：生产模式，使用嵌入文件系统，禁用调试功能

### 虚拟主机名

```go
const hostName = "app.example" // 可自定义虚拟主机名
```

访问方式：`https://app.example/EmbedAssets.html`

### 文件嵌入范围

```go
//go:embed assets/**           // 不包含隐藏文件
//go:embed all:assets/**       // 包含所有文件（包括隐藏文件）
```

### 窗口配置

通过 `NewWebViewWithWindow` 的选项参数配置：

```go
edge.WithXmlWindowTitle("我的应用")    // 窗口标题
edge.WithXmlWindowSize(1300, 900)     // 窗口大小
edge.WithFillParent(true)             // WebView 填充窗口
edge.WithAppDrag(true)                // 启用应用拖拽
edge.WithDebug(isDebug)               // 是否启用调试
edge.WithDefaultContextMenus(isDebug) // 是否显示右键菜单
edge.WithBrowserAcceleratorKeys(isDebug) // 是否启用浏览器快捷键
edge.WithStatusBar(false)             // 是否显示状态栏
edge.WithZoomControl(false)           // 是否显示缩放控制
```

## 设计亮点

### 单文件部署
- 所有前端资源打包进可执行文件
- 无需额外的依赖文件
- 简化部署和分发流程
- 便携性极佳

### 用户体验优化
- 渐进式加载避免白屏闪烁
- 自定义标题栏提供一致的体验
- 现代化的滚动条样式
- 开发时支持热重载

### 灵活开发流程
- 开发模式：实时修改和调试
- 生产模式：一次编译，随处运行
- 统一的虚拟主机名访问方式
- 支持任意前端技术栈

## 扩展可能性

这个示例可以作为基础，开发：

- **离线 Web 应用**：完全离线运行的 Web 应用
- **混合桌面应用**：结合桌面和 Web 技术的应用
- **企业内部工具**：单文件部署的企业工具
- **游戏配置工具**：游戏的配置和启动器
- **系统管理工具**：基于 Web 界面的系统管理应用
- **数据可视化应用**：基于现代 Web 技术的数据展示工具

## 最佳实践

### 资源组织
- 将前端资源合理组织在 `assets/` 目录中
- 使用规范的目录结构（html、css、js 分离）
- 考虑使用前端框架（React、Vue 等）时的构建流程

### 开发流程
- 开发时使用文件夹映射（`isDebug = true`）
- 生产时使用嵌入模式（`isDebug = false`）
- 定期测试两种模式的兼容性

### 性能优化
- 合理配置 WebView2 环境选项
- 禁用跟踪防护可以提升性能（仅在内容可控时）
- 使用现代化滚动条样式改善用户体验

### 安全考虑
- 生产环境关闭调试功能
- 避免在嵌入的资源中包含敏感信息
- 注意跨域安全策略

### 错误处理
- 完善的错误处理和用户提示
- WebView2 版本检测和兼容性处理
- 资源加载失败时的降级方案

## 注意事项

1. **WebView2 版本要求**：某些高级功能需要较新版本的 WebView2 运行时
3. **更新机制**：嵌入的资源无法像外部文件那样独立更新
4. **内存占用**：嵌入的文件系统会占用一定内存
5. **用户数据目录**：生产环境应使用固定的用户数据目录，而非临时目录
6. **开发路径**：开发模式下的文件夹路径需要注意相对路径和绝对路径的转换

## 常见问题

### Q: 为什么生产模式要等待加载完成才显示窗口？
A: 因为使用嵌入文件系统时，页面加载会先显示白屏，延迟显示可以避免这种闪烁现象。

### Q: 如何支持更多前端框架？
A: 可以在 `assets/` 目录中使用构建工具（如 webpack、vite 等），将构建结果嵌入。

### Q: 如何处理资源更新？
A: 由于资源被嵌入，更新需要重新编译程序。如需动态更新，可以考虑混合模式（核心资源嵌入，动态资源从网络加载）。

### Q: 开发模式和生产模式可以同时使用吗？
A: 不可以，需要通过修改 `isDebug` 变量来切换。建议开发时用 `true`，发布时用 `false`。

这个示例展示了现代桌面应用开发的最佳实践，将 Web 技术的灵活性与桌面应用的便捷部署完美结合。
