# SharedBuffer 共享缓冲区示例

> 本文由 AI 生成

## 简介

SharedBuffer 示例演示了如何在 Go 应用程序和 WebView2 浏览器控件之间通过共享内存高效传输数据。该技术利用 WebView2 的共享缓冲区（Shared Buffer）功能，避免了传统数据传输中的复制开销，特别适用于需要频繁或大量数据交换的场景。

## 核心功能

1. **高效内存共享**：利用 WebView2 的共享缓冲区功能，在 Go 应用和前端 JavaScript 之间共享内存数据，避免数据复制开销
2. **图片传输**：演示如何通过共享缓冲区传输图片数据到 WebView，使用可复用缓冲区模式
3. **文本传输**：展示如何通过共享缓冲区传输文本数据到 WebView，使用一次性缓冲区模式
4. **权限控制**：支持设置共享缓冲区的读写权限（只读/读写）
5. **资源管理**：正确的资源释放机制，确保共享内存被及时回收
6. **双向通信**：通过 WebMessage 实现前端到 Go 端的消息传递，触发数据发送

## 技术特点

### 核心技术
- 使用 `ICoreWebView2Environment12.CreateSharedBuffer` 创建共享缓冲区
- 通过 `ICoreWebView2_17.PostSharedBufferToScript` 将缓冲区发送到 WebView
- 在 JavaScript 中通过 `chrome.webview.addEventListener("sharedbufferreceived")` 接收共享缓冲区
- 使用嵌入式文件系统（embed.FS）加载前端资源
- 虚拟主机名映射，实现资源本地化加载

### 两种缓冲区模式
1. **可复用缓冲区**（`SharedBufferSender`）：
   - 创建固定大小的缓冲区（本例为 20MB）
   - 可多次发送数据，每次发送前清空缓冲区
   - 适用于频繁发送同一类型数据的场景
   - 前提条件：已知数据大小范围，可预估缓冲区容量

2. **一次性缓冲区**（`sendData` 函数）：
   - 根据数据长度动态创建缓冲区
   - 发送后立即关闭，不重复使用
   - 适用于数据大小不确定或偶尔发送的场景
   - 简化资源管理，避免缓冲区浪费

### 自定义数据协议
- **数据长度字段**：前 4 个字节（uint32，小端序）存储实际数据长度
- **数据字段**：从第 5 个字节开始，存储实际数据内容
- **附加数据**：通过 JSON 格式传递元数据（如数据类型 "type":"img"）
- **前端解析**：JavaScript 端根据数据长度字段正确提取数据

### 资源生命周期管理
- **创建阶段**：通过 WebView2 环境创建共享缓冲区和流对象
- **使用阶段**：通过 IStream 接口写入数据，通过 PostSharedBufferToScript 发送
- **释放阶段**：
  - Go 端调用 `Close()` 方法释放 COM 对象和缓冲区
  - JavaScript 端调用 `chrome.webview.releaseBuffer()` 释放引用
  - 双端释放后，操作系统才会真正释放底层共享内存
- **销毁事件**：通过 WebView 的 Event_Destroy 事件确保缓冲区在窗口销毁时释放

## 项目结构

```
SharedBuffer/
├── SharedBuffer.go      # Go 主程序，包含共享缓冲区实现
├── assets/
│   └── SharedBuffer.html  # 前端 HTML 页面
├── res/
│   └── 1.png              # 示例图片资源
└── README.md            # 本文档
```

## 运行示例

### 环境要求
- Windows 操作系统
- Go 1.18+
- WebView2 运行时（本程序会自动检测并提示安装）

### 运行步骤

1. **安装依赖**
   ```bash
   # 下载本库依赖
   go mod tidy
   ```

2. **运行程序**
   ```bash
   go run SharedBuffer.go
   ```

3. **测试功能**
   程序启动后会自动：
   - 创建 1200x900 的窗口
   - 加载 WebView 控件
   - 发送初始的示例图片和随机文本
   
   用户操作：
   - **"选择图片"** 按钮：打开文件选择器，选择本地图片文件，通过共享缓冲区传输到 WebView 显示
   - **"发送文本"** 按钮：生成 100 字符的随机文本，通过共享缓冲区传输到 WebView 显示

## 核心 API

### Go 端

#### WebView2 核心 API
- `ICoreWebView2Environment12.CreateSharedBuffer(size uint64)`：创建指定大小的共享缓冲区
- `ICoreWebView2SharedBuffer.GetBuffer()`：获取共享缓冲区的内存地址（指针）
- `ICoreWebView2SharedBuffer.OpenStream()`：获取访问共享缓冲区的 IStream 对象，用于读写数据
- `ICoreWebView2SharedBuffer.Close()`：关闭共享缓冲区
- `ICoreWebView2_17.PostSharedBufferToScript(buffer, access, additionalData)`：将共享缓冲区发送到 WebView 脚本端

#### 封装类和方法

**SharedBufferSender**（可复用缓冲区发送者）
```go
// 创建共享缓冲区发送者，默认只读权限
func NewSharedBufferSender(wv *edge.WebView, size uint64) (*SharedBufferSender, error)

// 发送数据，前 4 个字节为数据长度
func (s *SharedBufferSender) Send(data []byte, additionalDataAsJson ...string) error

// 设置缓冲区访问权限
func (s *SharedBufferSender) SetAccess(access edge.COREWEBVIEW2_SHARED_BUFFER_ACCESS) *SharedBufferSender

// 设置附加 JSON 数据（如数据类型标识）
func (s *SharedBufferSender) SetAdditionalDataAsJson(json string) *SharedBufferSender

// 获取缓冲区大小
func (s *SharedBufferSender) GetBufferSize() uint64

// 关闭缓冲区，释放资源
func (s *SharedBufferSender) Close()
```

**sendData**（一次性缓冲区发送函数）
```go
// 创建一次性共享缓冲区并发送数据
func sendData(wv *edge.WebView, data []byte, typeStr string) error
```

#### 辅助函数
- `Uint32ToBytes(i uint32)`：将 uint32 转换为字节数组（小端序）
- `generateRandomString(length int)`：生成指定长度的随机字符串

### JavaScript 端

#### WebView2 API
- `chrome.webview.addEventListener("sharedbufferreceived", callback)`：监听共享缓冲区接收事件
- `event.getBuffer()`：获取共享缓冲区的 ArrayBuffer 对象
- `event.additionalData`：获取附加的 JSON 数据（已解析为对象）
- `chrome.webview.releaseBuffer(buffer)`：释放共享缓冲区引用
- `chrome.webview.postMessage(message)`：向 Go 端发送消息

#### 自定义函数
```javascript
// 从缓冲区前 4 个字节获取数据长度（小端序）
function GetDataLength(buffer) {
    return new DataView(buffer).getUint32(0, true);
}

// 根据数据长度从缓冲区获取图片数据并创建 Blob URL
function GetImgUrl(buffer) {
    const len = GetDataLength(buffer);
    const imgData = new Uint8Array(buffer, 4, len);
    return URL.createObjectURL(new Blob([imgData]));
}

// 将 ArrayBuffer 转换为字符串
function BufferToString(buffer) {
    return new TextDecoder().decode(new Uint8Array(buffer));
}

// 释放缓冲区
function ReleaseBuffer(buffer) {
    window.chrome.webview.releaseBuffer(buffer);
}

// 发送消息到 Go 端
function PostMessage(message) {
    window.chrome.webview.postMessage(message);
}
```

## 数据流程

### 图片传输流程（可复用缓冲区）
```
用户点击"选择图片" 
  ↓
选择本地图片文件
  ↓
读取文件到内存
  ↓
SharedBufferSender.Send(data) 
  ↓
写入数据长度（4字节）到缓冲区
  ↓
写入实际数据到缓冲区
  ↓
PostSharedBufferToScript 发送到 WebView
  ↓
JavaScript 接收 sharedbufferreceived 事件
  ↓
根据 additionalData.type="img" 判断为图片
  ↓
GetImgUrl() 解析缓冲区创建 Blob URL
  ↓
显示图片到 <img> 标签
  ↓
调用 releaseBuffer() 释放缓冲区引用
```

### 文本传输流程（一次性缓冲区）
```
用户点击"发送文本"
  ↓
生成随机字符串
  ↓
sendData(wv, data, "text")
  ↓
创建与数据等大的共享缓冲区
  ↓
写入数据到缓冲区
  ↓
PostSharedBufferToScript 发送到 WebView（附加 type="text"）
  ↓
JavaScript 接收事件，解析为文本
  ↓
显示文本到 <p> 标签
  ↓
调用 releaseBuffer() 释放
  ↓
Go 端自动关闭缓冲区
```

### 页面初始化流程
```
页面加载完成
  ↓
发送 "RequestImg" 消息到 Go 端
  ↓
发送 "RequestText" 消息到 Go 端
  ↓
Go 端接收 WebMessageReceived 事件
  ↓
根据消息类型分别发送数据
  ↓
前端接收并显示
```

## 使用场景

1. **大文件传输**：传输大型 JSON 对象、二进制数据块等，避免复制开销
2. **实时数据更新**：频繁更新的数据，如实时图表、监控数据等
3. **多媒体处理**：图片、音频、视频等多媒体数据的高效传输
4. **跨进程通信**：在不同进程间高效共享数据，零拷贝传输
5. **游戏应用**：游戏资源的快速加载和更新
6. **桌面应用**：现代化桌面应用中的高性能数据交换

## 注意事项

### 资源管理
1. **双端释放原则**：
   - Go 端：使用完毕后调用 `Close()` 方法释放 SharedBufferSender
   - JavaScript 端：使用完毕后调用 `chrome.webview.releaseBuffer()` 释放引用
   - 只有两端都释放，操作系统才会真正释放底层共享内存

2. **生命周期控制**：
   - 可复用缓冲区在程序关闭或不再需要时调用 Close()
   - 一次性缓冲区在发送后立即释放（defer 模式）
   - 建议在 WebView 销毁事件中主动释放缓冲区

### 权限控制
3. **访问权限**：
   - 默认为只读权限（`COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY`）
   - 可通过 `SetAccess()` 设置为读写权限
   - **警告**：写入只读缓冲区会导致渲染器进程崩溃

### 性能考虑
4. **缓冲区大小**：
   - 目前限制为 2GB 以内
   - 可复用缓冲区需要预估最大数据量
   - 过大的缓冲区会浪费内存，过小会导致发送失败

5. **数据协议**：
   - 前 4 字节为数据长度字段（uint32，小端序）
   - JavaScript 端必须按照此协议解析数据
   - 可根据需要扩展协议（如添加校验和、压缩标志等）

### 兼容性
6. **WebView2 版本**：
   - 需要 WebView2 运行时支持 Shared Buffer 功能
   - 本程序会自动检测版本并提示更新
   - 建议使用最新版本的 WebView2 运行时

### 调试建议
7. **调试技巧**：
   - 使用 `edge.WithDebug(true)` 启用开发者工具
   - 在 JavaScript 控制台查看缓冲区接收情况
   - Go 端使用 log.Println 输出调试信息

## 扩展建议

### 性能优化
- 对于超大数据，考虑分片传输
- 实现数据压缩以减少传输量
- 使用内存池管理缓冲区分配

### 功能扩展
- 添加数据校验机制（CRC32、MD5）
- 支持双向共享缓冲区（JavaScript 写入，Go 读取）
- 实现多缓冲区并发传输
- 添加传输进度反馈

### 安全增强
- 实现数据加密传输
- 添加缓冲区访问权限验证
- 限制单次传输数据大小

## 常见问题

**Q: 为什么需要自定义数据协议（前4字节存储长度）？**
A: WebView2 的 Shared Buffer 只传递原始内存，没有内置的数据边界标识。自定义协议可以让 JavaScript 端正确提取实际数据，避免解析错误。

**Q: 何时使用可复用缓冲区 vs 一次性缓冲区？**
A: 频繁发送数据且数据大小相对固定时，使用可复用缓冲区提高性能；偶尔发送或数据大小不确定时，使用一次性缓冲区简化管理。

**Q: 如果数据超过缓冲区大小会怎样？**
A: `SharedBufferSender.Send()` 会返回错误 "数据长度超过缓冲区大小"，需要处理此错误或重新创建更大的缓冲区。

**Q: 共享缓冲区是否线程安全？**
A: WebView2 的 Shared Buffer 本身不保证线程安全。在 Go 端应确保同一时间只有一个线程写入缓冲区，或在应用层实现同步机制。

**Q: 能否在多个 WebView 实例间共享缓冲区？**
A: 不可以。Shared Buffer 只能在单个 WebView2 实例的宿主和渲染器进程间共享。如需跨 WebView 共享数据，应使用其他 IPC 机制。