# Vue + Vite 桌面应用

> 本文由 AI 生成

基于 Vue 3 + Vite + WebView2 的现代桌面应用，支持开发时热重载和发布时嵌入资源成单文件。

## 项目特性

- 🎨 **现代 UI 设计**: 左侧导航栏 + 右侧内容区的经典布局
- ⚡ **Vue 3 + Vite**: 使用最新的 Vue 3 组合式 API 和 Vite 构建工具
- 🔄 **极速热重载**: 开发模式使用 Vite HMR，修改代码立即生效，无需手动刷新
- 📦 **单文件部署**: 发布时自动嵌入资源，无需外部文件依赖
- 🖥️ **WebView2**: 基于 Microsoft WebView2 的现代浏览器引擎
- 🚀 **高性能**: 快速启动，流畅交互
- 🛠️ **简单开发**: 无需额外的监听工具，Vite 自带热重载

## 系统要求

- Windows 10 1809 或更高版本
- Go 1.18+
- Node.js 18+ (用于前端开发)
- WebView2 运行时 (自动检测，缺失时会提示安装)

## 项目结构

```
VueAndVite/
├── VueAndVite.go              # Go 主程序
├── package.json         # Node.js 依赖配置
├── vite.config.js       # Vite 配置文件
├── index.html           # 入口 HTML
├── src/                 # Vue 源代码目录
│   ├── main.js          # 应用入口
│   ├── App.vue          # 主应用组件
│   └── assets/          # 资源文件
│       └── main.css      # 全局样式
├── dist/                # 构建输出目录 (会被嵌入到 Go 程序)
├── dev.bat              # 开发模式启动脚本
├── build.bat            # 生产模式构建脚本
├── .gitignore           # Git 忽略配置
└── README.md            # 本文档
```

## 快速开始

> **最快的方式是直接启动 dev.bat 即可, 下面是手动启动的步骤.**

### 1. 安装前端依赖

```bash
cd webview/VueAndVite
npm install
```

### 2. 开发模式

**第一步**: 设置 `VueAndVite.go` 中的调试模式为 `true` (默认已设置)

```go
var isDebug = true // 开发模式
```

**第二步**: 启动 Vite 开发服务器 (在新的终端窗口)

```bash
npm run dev
```

这将在 `http://localhost:5173` 启动开发服务器，支持热重载。

**第三步**: 运行 Go 程序 (在另一个终端窗口)

在项目根目录 (xcgui-example) 下运行:

```bash
go run webview/VueAndVite/VueAndVite.go
```

**热重载**: 在开发模式下，Vite 会自动监听文件变化，修改 Vue 代码后会自动刷新，无需手动构建或刷新！

这就是 Vite 的强大之处，开发体验非常流畅。

如果更改了 go 代码, 那需要重新 go run.

### 3. 生产模式 (单文件部署)

> 可直接运行 build.bat, 下面是手动构建的步骤.

**第一步**: 构建前端资源

```bash
npm run build
```

这会将构建结果输出到 `assets/` 目录。

**第二步**: 设置 `VueAndVite.go` 中的调试模式为 `false`

```go
var isDebug = false // 生产模式
```

**第三步**: 编译 Go 程序

```bash
go build -ldflags="-s -w -H windowsgui" -trimpath -o VueAndVite.exe webview/VueAndVite/VueAndVite.go
```

生成的 `VueAndVite.exe` 就是完整的单文件应用，可以直接分发运行。

## 技术栈

### 前端
- **Vue 3**: 渐进式 JavaScript 框架
- **Vite**: 下一代前端构建工具，支持极速热重载
- **CSS3**: 现代样式，支持渐变、动画等效果

### 后端
- **Go**: 主要编程语言
- **XCGUI**: 炫彩界面库
- **WebView2**: Microsoft Edge WebView2 控件

### 开发特性
- **Vite HMR**: 开发模式下支持极速热模块替换，无需手动刷新

## 依赖获取

### 获取 Go 依赖

运行:

```bash
go mod tidy
```

### 获取 Node.js 依赖

```bash
npm install
```

主要依赖:
- `vue`: ^3.5.7
- `vite`: ^6.2.3
- `@vitejs/plugin-vue`: ^5.0.3

## 配置说明

### Vite 配置

`vite.config.js` 主要配置:

```javascript
{
  base: './',                 // 使用相对路径
  build: {
    outDir: 'dist',           // 输出到 dist 目录
    emptyOutDir: true,        // 清空输出目录
  },
  server: {
    port: 5173,              // 开发服务器端口
    strictPort: true,         // 端口被占用时报错
    cors: true               // 启用 CORS
  }
}
```

### Go 程序配置

`VueAndVite.go` 主要配置:

```go
var isDebug = true              // 调试/生产模式切换
const hostName = "app.example" // 虚拟主机名 (仅生产模式使用)
```

窗口配置:
```go
edge.WithXmlWindowTitle("Vue Desktop App"), // 窗口标题
edge.WithXmlWindowSize(1300, 900),          // 窗口大小
edge.WithAppDrag(true),                     // 启用拖拽
```

## 功能模块

### 导航菜单

- **仪表盘**: 显示应用概览和统计数据
- **设置**: 应用设置选项
- **关于**: 应用信息和版本号

### 窗口控制

支持窗口最小化、最大化/还原、关闭操作，通过 Go 与 JavaScript 双向绑定实现。

## 扩展开发

### 添加新页面

1. 在 `App.vue` 中添加新的菜单项:

```vue
const menuItems = [
  { id: 'newpage', name: '新页面', icon: '📄' },
  // ...
]
```

2. 添加页面内容:

```vue
<div v-else-if="activeMenu === 'newpage'" class="newpage">
  <h2>新页面</h2>
  <!-- 页面内容 -->
</div>
```

### 添加 Go 函数

在 `bindBasicFuncs()` 中绑定新函数:

```go
m.wv.Bind("go.myFunction", func() {
    // 你的代码
})
```

在 Vue 中调用:

```javascript
if (window.go && window.go.myFunction) {
  window.go.myFunction()
}
```

### 常见问题

### Q: 如何启用热重载？

A: 开发模式下自动启用！
1. 运行 `npm run dev` 启动 Vite 开发服务器
2. 运行 `go run VueAndVite.go` 启动 Go 应用
3. 修改 Vue 代码，Vite 会自动热重载，无需手动操作

### Q: 开发模式和生产模式有什么区别？

A:
- **开发模式**: 连接 Vite 开发服务器 (`http://localhost:5173`)，支持热重载，文件未嵌入
- **生产模式**: 使用嵌入的文件系统，所有资源打包到 exe 中，不支持热重载

### Q: 如何修改窗口大小？

A: 修改 `VueAndVite.go` 中的 `WithXmlWindowSize(1300, 900)` 参数。

### Q: 支持自定义主题吗？

A: 支持。修改 `src/assets/main.css` 或 `App.vue` 中的样式即可。

## 性能优化

- 使用 Vite 的快速构建和热模块替换 (HMR)
- WebView2 环境优化:
  - 禁用跟踪防护
  - 使用现代化滚动条样式
- 代码分割: Vite 自动将 Vue 打包为单独的 chunk
- 资源压缩: Vite 生产构建自动压缩代码

## 构建优化

如需优化最终 exe 文件大小:

使用 UPX 压缩:

```bash
upx --best --lzma myapp.exe
```

## 相关资源

- [Vue 3 官方文档](https://vuejs.org/)
- [Vite 官方文档](https://vitejs.dev/)
- [XCGUI 文档](https://github.com/twgh/xcgui)
- [WebView2 文档](https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/)

