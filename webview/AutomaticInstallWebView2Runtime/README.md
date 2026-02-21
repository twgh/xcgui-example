# 自动检测并安装 WebView2 运行时

> 本文由 AI 生成

本示例演示如何在程序启动时自动检测 WebView2 运行时是否已安装，如果未安装则自动运行嵌入的安装程序进行安装。

## 功能说明

- **自动检测**：启动时检测本机是否已安装 WebView2 运行时
- **自动安装**：如果未检测到 WebView2 运行时，自动运行嵌入的安装程序
- **等待安装完成**：程序会等待安装完成后才继续创建 WebView2 窗口
- **WebView2 浏览**：安装完成后自动打开窗口并加载百度首页

## 前提条件

在构建本示例之前，需要将 WebView2 运行时安装程序嵌入到可执行文件中：

1. 下载 **WebView2 运行时小型安装引导程序**：
   - 访问 [WebView2 运行时下载页面](https://developer.microsoft.com/zh-cn/microsoft-edge/webview2/)
   - 下载 **"Evergreen Standalone Installer"** 或 **"Evergreen Bootstrapper"**
   - 建议下载 `MicrosoftEdgeWebview2Setup.exe`（小型安装引导程序，约 2MB）

2. 将下载的安装程序重命名为 `MicrosoftEdgeWebview2Setup.exe` 并放置在本目录下

## 构建与运行

### 构建

```bash
go build -ldflags "-s -w -H widnowsgui" -trimpath -o AutomaticInstallWebView2Runtime.exe
```

### 运行

```bash
.\AutomaticInstallWebView2Runtime.exe
```

运行后程序会自动检测 WebView2 运行时状态：
- 如果已安装，直接创建窗口并加载网页
- 如果未安装，会弹出提示框，然后自动运行安装程序，等待安装完成后才继续

## 代码说明

### 核心函数

| 函数 | 说明 |
|------|------|
| `AutomaticInstallWebView2Runtime()` | 检测并安装 WebView2 运行时 |
| `RunWebView2Installer()` | 运行嵌入的安装程序，支持静默安装 |
| `createEdge()` | 创建 WebView2 环境 |

### 关键代码片段

**检测本地 WebView2 版本：**

```go
localVersion, _ := edge.GetAvailableBrowserVersion()
if localVersion == "" {
    // 未安装，执行安装逻辑
}
```

**运行安装程序：**
```go
// 将嵌入的安装程序写入临时目录并执行
installerPath := filepath.Join(os.TempDir(), "Webview2Installer.exe")
os.WriteFile(installerPath, WebView2Installer, 0777)
exec.Command(installerPath).Run()
```

**等待安装完成：**
```go
for i := 0; i < 300; i++ {
    time.Sleep(time.Second)
    localVersion, _ = edge.GetAvailableBrowserVersion()
    if localVersion != "" {
        break
    }
}
```

### 静默安装

`RunWebView2Installer` 函数支持静默安装参数：

```go
// 静默安装（不显示安装界面）
RunWebView2Installer(true)

// 普通安装（显示安装界面）
RunWebView2Installer(false)
```

## 注意事项

1. **安装包大小**：使用小型安装引导程序（约 2MB）时，实际运行时还需要联网下载完整的 WebView2 运行时（约 130MB），安装时间取决于网络速度

2. **安装时间**：首次安装可能需要几分钟时间，程序会等待安装完成（最多等待 5 分钟）

3. **临时目录**：安装程序会被释放到系统临时目录，安装完成后不会自动删除

4. **UserDataFolder**：本示例使用 `os.TempDir()` 作为用户数据目录，实际应用中应使用固定的应用数据目录

5. **权限**：安装 WebView2 运行时可能需要管理员权限，建议在程序清单中配置或提示用户以管理员身份运行
