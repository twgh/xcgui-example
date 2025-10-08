# EmbedAssets - 嵌入式资源WebView应用

> 本文由 AI 生成

## 项目简介

EmbedAssets 是一个演示如何将前端资源嵌入到可执行文件中的WebView应用示例。该示例展示了如何实现单文件部署，无需启动HTTP服务器，通过虚拟主机名映射的方式访问嵌入的HTML、CSS、JavaScript等前端资源。

## 功能特性

### 核心功能
- **嵌入式资源**: 将所有前端资源打包到可执行文件中
- **虚拟主机映射**: 通过虚拟主机名访问嵌入的文件系统
- **单文件部署**: 整个应用打包为单个可执行文件

### 界面特性
- **自定义标题栏**: 完全自定义的HTML标题栏，包含窗口控制按钮
- **现代化界面**: 基于HTML/CSS的现代化用户界面
- **透明窗口**: 解决窗口切换时的闪烁问题
- **响应式设计**: 支持窗口拖拽和大小调整

### 技术特性
- **Go与JavaScript交互**: 支持Go函数和JavaScript之间的双向调用
- **环境选项配置**: 详细的WebView2环境配置示例
- **资源热重载**: 开发模式下支持前端资源的实时更新
- **渐进式加载**: 等待页面完全加载后再显示窗口

## 项目结构

```
EmbedAssets/
├── EmbedAssets.go        # 主程序文件
├── 1.jpg                # 窗口界面展示图片
├── assets/              # 前端资源目录（会被嵌入）
│   └── EmbedAssets.html # 主页面文件
├── res/                 # 窗口布局资源
│   └── EmbedAssets.xml  # 窗口布局文件
└── README.md            # 本文档
```

## 技术实现

### 技术栈
- **Go 语言**: 主要编程语言
- **XCGUI**: 炫彩界面库
- **WebView2**: 嵌入式浏览器组件
- **embed**: Go 1.16+ 文件嵌入功能
- **HTML/CSS/JavaScript**: 前端技术栈

### 关键技术点

#### 1. 文件嵌入
```go
//go:embed assets/**
embedAssets embed.FS // 嵌入assets目录及其所有子目录文件
```

#### 2. 虚拟主机映射
```go
const hostName = "app.example"

// 生产模式：映射嵌入文件系统
err = edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)

// 开发模式：映射本地文件夹
folderPath, _ := filepath.Abs("webview/EmbedAssets/assets")
err = wv.SetVirtualHostNameToFolderMapping(hostName, folderPath, 
    edge.COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW)
```

#### 3. 环境配置优化
```go
// 禁用跟踪防护以提升性能
envOpts5.SetEnableTrackingPrevention(false)

// 设置滚动条样式
envOpts8.SetScrollBarStyle(edge.COREWEBVIEW2_SCROLLBAR_STYLE_FLUENT_OVERLAY)
```

#### 4. Go与JavaScript交互
```go
// 绑定Go函数供JavaScript调用
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

### 设计模式

#### MainWindow结构体模式
采用面向对象的设计模式，将相关功能封装在MainWindow结构体中：
```go
type MainWindow struct {
    edg *edge.Edge      // WebView2环境
    w   *window.Window  // 炫彩窗口
    wv  *edge.WebView   // WebView实例
}
```

#### 渐进式显示模式
通过导航完成事件实现无闪烁的窗口显示：
```go
wv.Event_NavigationCompleted(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) uintptr {
    if firstLoad {
        firstLoad = false
        w.Show(true) // 首次加载完成后才显示窗口
    }
    return 0
})
```

## 使用方法

### 开发模式
1. 运行程序：
   ```bash
   go run EmbedAssets.go
   ```

### 生产模式
1. 设置 `isDebug = false`
2. 编译程序
3. 所有资源已嵌入，可直接分发exe文件

### 前端开发
- HTML文件位于 `assets/EmbedAssets.html`
- 可以使用任何前端框架和技术
- 通过虚拟主机名 `app.example` 访问资源

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 配置说明

### 调试开关
```go
var isDebug = false // 控制调试/生产模式
```

### 虚拟主机名
```go
const hostName = "app.example" // 可自定义虚拟主机名
```

### 文件嵌入范围
```go
//go:embed assets/**           // 不包含隐藏文件
//go:embed all:assets/**       // 包含所有文件（包括隐藏文件）
```

## 设计亮点

### 单文件部署
- 所有前端资源打包进可执行文件
- 无需额外的依赖文件
- 简化部署和分发流程

### 用户体验优化
- 渐进式加载避免白屏闪烁
- 自定义标题栏提供一致的体验
- 透明窗口设计解决视觉问题

## 扩展可能性

这个示例可以作为基础，开发：
- **离线Web应用**: 完全离线运行的Web应用
- **混合桌面应用**: 结合桌面和Web技术的应用
- **企业内部工具**: 单文件部署的企业工具
- **游戏配置工具**: 游戏的配置和启动器
- **系统管理工具**: 基于Web界面的系统管理应用

## 最佳实践

1. **资源组织**: 将前端资源合理组织在assets目录中
2. **版本控制**: 开发时使用文件夹映射，生产时使用嵌入模式
3. **错误处理**: 完善的错误处理和用户提示
4. **性能优化**: 合理配置WebView2环境选项
5. **安全考虑**: 生产环境关闭调试功能

这个示例展示了现代桌面应用开发的最佳实践，将Web技术的灵活性与桌面应用的便捷部署完美结合。