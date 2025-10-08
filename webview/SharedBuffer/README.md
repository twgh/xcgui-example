# SharedBuffer 共享缓冲区示例

> 本文由 AI 生成

## 简介

SharedBuffer 示例演示了如何在 Go 应用程序和 WebView2 浏览器控件之间通过共享内存高效传输数据。该技术避免了传统数据传输中的复制开销，特别适用于需要频繁或大量数据交换的场景。

## 核心功能

1. **高效内存共享**：利用 WebView2 的共享缓冲区功能，在 Go 应用和前端 JavaScript 之间共享内存数据，避免数据复制开销
2. **图片传输**：演示如何通过共享缓冲区传输图片数据到 WebView
3. **文本传输**：展示如何通过共享缓冲区传输文本数据到 WebView
4. **权限控制**：支持设置共享缓冲区的读写权限
5. **资源管理**：正确的资源释放机制，确保共享内存被及时回收

## 技术特点

- 使用 `ICoreWebView2Environment12.CreateSharedBuffer` 创建共享缓冲区
- 通过 `ICoreWebView2_17.PostSharedBufferToScript` 将缓冲区发送到 WebView
- 在 JavaScript 中通过 `chrome.webview.addEventListener("sharedbufferreceived")` 接收共享缓冲区
- 封装了一次性缓冲区和可复用缓冲区两种使用模式
- 实现了完整的资源管理机制，确保内存正确释放
- **自定义数据协议**：在 `SharedBufferSender` 创建的共享缓冲区中，前4个字节用于存储数据长度，后面跟随实际数据，便于前端正确解析数据

## 运行示例

1. 确保已安装 WebView2 运行时
2. 进入SharedBuffer目录并运行程序：
   ```bash
   go run SharedBuffer.go
   ```
3. 点击界面中的按钮测试功能：
   - "选择图片"：选择本地图片文件并通过共享缓冲区传输到 WebView 显示
   - "发送文本"：生成随机文本并通过共享缓冲区传输到 WebView 显示

## 核心 API

### Go 端

- `ICoreWebView2Environment12.CreateSharedBuffer`：创建共享缓冲区
- `ICoreWebView2SharedBuffer.GetBuffer`：获取共享缓冲区的内存地址
- `ICoreWebView2SharedBuffer.OpenStream`：获取访问共享缓冲区的 IStream 对象
- `ICoreWebView2_17.PostSharedBufferToScript`：将共享缓冲区发送到 WebView
- `SharedBufferSender.Send`：发送数据到共享缓冲区，前4个字节为数据长度，后面为实际数据

### JavaScript 端

- `chrome.webview.addEventListener("sharedbufferreceived")`：监听共享缓冲区接收事件
- `event.getBuffer()`：获取共享缓冲区的 ArrayBuffer 对象
- `chrome.webview.releaseBuffer(buffer)`：释放共享缓冲区

  **自定义协议**:
- `GetDataLength(buffer)`：从缓冲区前4个字节获取数据长度
- `GetImgUrl(buffer)`：根据数据长度从缓冲区获取图片数据并创建URL

## 使用场景

1. **大文件传输**：传输大型 JSON 对象、二进制数据块等
2. **实时数据更新**：频繁更新的数据，如实时图表、监控数据等
3. **多媒体处理**：图片、音频、视频等多媒体数据的传输
4. **跨进程通信**：在不同进程间高效共享数据

## 注意事项

1. 共享缓冲区使用完毕后，需要在 Go 端调用 `Close()` 方法，在 JavaScript 端调用 `chrome.webview.releaseBuffer()` 方法
2. 只有当两端都释放了缓冲区，操作系统才会真正释放底层的共享内存
3. 可以设置缓冲区的访问权限（只读/读写），写入只读缓冲区会导致渲染器进程崩溃
4. 目前共享缓冲区大小限制为 2GB 以内
5. 数据在 `SharedBufferSender` 创建的共享缓冲区中的存储格式为：前4个字节为数据长度，后面为实际数据，前端需要按照此协议解析数据