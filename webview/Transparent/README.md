# WebView 局部透明示例

> 本文由 AI 生成

演示如何利用 **xcgui**（炫彩界面库）+ **Edge WebView2** 实现**窗口透明 + WebView 局部不透明**的效果。

## 效果

- 窗口的 `html/body` 背景是**完全透明**的，桌面背景可以直接透出来。
- 页面中的卡片（`.card`）、底部信息栏（`.info-bar`）以及按钮等元素是**不透明**的。
- 透明区域鼠标穿透（点击穿透到桌面），不透明区域可以正常点击。

## 核心原理

1. **窗口设置为透明**：使用 `xcc.Window_Transparent_Shaped` 将窗口设为透明窗口。
2. **WebView 背景透明**：通过 `edge.WithDefaultBackgroundColor(edge.NewColor(0, 0, 0, 0))` 将 WebView 默认背景设为透明。
3. **HTML/body 透明**：CSS 中设置 `background: transparent !important;`。
4. **Go 端对齐不透明区域**：页面加载完成后，通过 JS 获取 `.card` 和 `.info-bar` 元素到布局的边距，在 xcgui 的布局背景管理器中添加对应的不透明填充矩形，并设置相同的圆角。这样炫彩窗口本身上就有了与 HTML 不透明区域位置完全吻合的不透明区域，WebView不透明区域自然就不会鼠标穿透了, 就可以被点击了。

## 文件结构

```
webview/Transparent/
├── Transparent.go        # Go 主程序
├── README.md             # 本文件
├── 1.jpg                 # 效果截图
└── assets/
    └── Transparent.html  # 前端页面（嵌入到 Go 二进制中）
```

## 运行

```bash
go run .
```

> **注意**：运行前请确保本机已安装 [WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)（Evergreen 版本即可）。程序启动时会自动检测，若未安装会弹出提示并打开下载链接。

## 调试模式

将 `Transparent.go` 第 22 行的 `isDebug` 改为 `true`：

```go
var isDebug = true
```

此时程序会采用文件夹映射方式加载 assets 目录，修改 HTML/CSS/JS 后手动刷新页面即可生效，无需重新编译。

## 关键代码说明

| 功能 | 代码位置 |
|------|----------|
| 设置透明窗口 | `NewMainWindow()` 中 `m.w.SetTransparentType(...)` |
| 设置 WebView 透明背景 | `edge.WithDefaultBackgroundColor(edge.NewColor(0, 0, 0, 0))` |
| 加载完成后对齐不透明区域 | `Event_NavigationCompleted` 回调中获取 JS 返回的边距信息 |
| 添加炫彩填充矩形 | 通过 `layContent.GetBkManagerObj().AddFill(...)` 添加与 HTML 元素位置对应的填充矩形 |
| 窗口控制函数绑定 | `bindBasicFuncs()` 中绑定最小化、最大化、关闭 |
| 嵌入文件系统 | 正式版使用 `//go:embed assets/**` 嵌入资源，配合 `EnableVirtualHostNameToEmbedFSMapping` |
