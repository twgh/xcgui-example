# CalcMD5 - 文件MD5计算器

> 本文由 AI 生成

## 项目简介

CalcMD5 是一个基于 XCGUI 和 WebView2 技术的文件MD5哈希值计算工具。这个示例演示了如何使用 Go 语言创建一个现代化的桌面应用程序，结合 Web 前端技术实现用户界面。

## 功能特性

### 核心功能
- **文件MD5计算**: 快速计算选定文件的MD5哈希值
- **文件选择**: 通过点击按钮选择文件进行计算
- **拖拽支持**: 支持直接拖拽文件到应用窗口进行计算
- **实时显示**: 在界面上实时显示计算结果

### 界面特性
- **圆角窗口**: 8px圆角的现代化窗口设计
- **自定义标题栏**: 自定义的窗口控制按钮（最小化、最大化、关闭）
- **响应式设计**: 支持窗口拖拽和大小调整
- **美观界面**: 使用 HTML/CSS 实现的现代化用户界面

### 技术特性
- **异步处理**: MD5计算在后台线程执行，不阻塞界面
- **嵌入式资源**: 前端资源（HTML、CSS、JS）嵌入到可执行文件中
- **虚拟主机**: 使用虚拟主机名映射嵌入的文件系统

## 项目结构

```
CalcMD5/
├── CalcMD5.go         # 主程序文件
├── 1.jpg              # 窗口界面展示图片
├── assets/            # 前端资源目录
│   ├── CalcMD5.html   # 主界面HTML
│   ├── CalcMD5.css    # 样式文件
│   └── CalcMD5.js     # JavaScript逻辑
├── res/               # 窗口布局资源
│   └── CalcMD5.xml    # 窗口布局文件
└── README.md          # 本文档
```

## 技术实现

### 后端技术栈
- **Go 语言**: 主要编程语言
- **XCGUI**: GUI框架，提供窗口和UI组件
- **WebView2**: 嵌入式浏览器组件
- **embed**: Go 1.16+ 的文件嵌入功能

### 前端技术栈
- **HTML5**: 页面结构
- **CSS3**: 样式设计
- **JavaScript**: 交互逻辑
- **WebView2 API**: Go与JavaScript之间的通信桥梁

### 关键技术点

#### 1. 文件嵌入
使用 Go 的 `embed` 包将前端资源嵌入到可执行文件中：
```go
//go:embed assets/**
embedAssets embed.FS
```

#### 2. 虚拟主机映射
设置虚拟主机名与嵌入文件系统的映射：
```go
edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)
```

#### 3. Go与JavaScript通信
- **Go调用JS**: 使用 `webview.Eval()` 执行JavaScript代码
- **JS调用Go**: 使用 `webview.Bind()` 绑定Go函数供JavaScript调用

#### 4. 拖拽文件处理
通过 WebView2 的 `postMessageWithAdditionalObjects` API 实现拖拽文件功能：
```javascript
chrome.webview.postMessageWithAdditionalObjects('drag_files', e.dataTransfer.files);
```

#### 5. 异步MD5计算
在Go协程中执行耗时的MD5计算，避免阻塞UI线程：
```go
go func() {
    // MD5计算逻辑
    data, _ := os.ReadFile(filePath)
    hash := md5.Sum(data)
    // ...
}()
```

## 使用方法

### 运行程序
1. 确保已安装 WebView2 运行时
2. 运行程序：
   ```bash
   go run CalcMD5.go
   ```

### 计算文件MD5
1. **方法一**: 点击"选择文件"按钮，在弹出的文件对话框中选择文件
2. **方法二**: 直接将文件拖拽到应用程序窗口中
3. 程序会自动计算并显示文件的MD5哈希值

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测并提示下载）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

## 设计亮点

### 用户体验
- **无闪烁启动**: 等待页面完全加载后再显示窗口，避免白屏闪烁
- **实时反馈**: 计算过程中显示"计算中..."状态

### 技术创新
- **透明窗口设计**: 使用透明的XCGUI窗口作为WebView2的容器，避免窗口切换时的视觉闪烁
- **圆角一致性**: WebView2和外层窗口都设置8px圆角，保持视觉一致性
- **资源一体化**: 将所有前端资源嵌入可执行文件，简化部署

## 扩展可能性

这个示例可以作为基础，扩展为：
- 多种哈希算法支持（SHA1、SHA256等）
- 批量文件处理
- 哈希值比较功能
- 文件完整性验证工具
- 更复杂的桌面应用程序

## 注意事项

1. **调试模式**: 代码中的 `isDebug` 变量控制调试功能的开启
2. **用户数据目录**: 示例中使用临时目录，实际应用应使用固定的应用数据目录
3. **WebView2版本**: 程序会检查本地WebView2版本兼容性
4. **文件大小**: 对于大文件的MD5计算可能需要较长时间

这个示例完美展示了 XCGUI 与 WebView2 结合开发现代桌面应用的优势，为开发者提供了一个实用的参考模板。