# WebResourceRequestedEvent 示例

> 本文由 AI 生成

## 功能介绍

本示例演示了如何使用 WebView2 的网络资源请求和响应接收事件，实现对网页网络请求的拦截、修改以及响应数据的捕获。

## 核心功能

### 1. 网络资源请求拦截与修改 (WebResourceRequested)
通过 `Event_WebResourceRequested` 事件拦截网络请求，并根据需求修改请求参数：

- **日K线数据查询**：将数据条数从 210 条修改为 570 条
- **个股信息查询**：在请求的 fields 参数中添加额外字段 `f189`

### 2. 网络资源响应接收 (WebResourceResponseReceived)
通过 `Event_WebResourceResponseReceived` 事件捕获网络响应内容，并保存到本地文件：

- 日K线数据响应 → 保存为 `日k线数据响应.json`
- 个股信息查询响应 → 保存为 `个股信息查询响应.json`
- 行业查询响应 → 保存为 `行业查询响应.json`

### 3. 并发请求处理
演示了如何使用多个 `WebViewEventImpl` 实例来处理并发的异步请求，避免回调函数被覆盖：

```go
impls = make([]*edge.WebViewEventImpl, 2)
// 为不同类型的响应分配独立的 WebViewEventImpl
```

### 4. 自动化交互
在页面加载完成后，自动点击日K选项，触发数据请求：

```go
wv.Event_NavigationCompleted(...)
// 自动执行点击操作
```

## 技术要点

### WebView2 事件处理
- `Event_WebResourceRequested`：拦截所有网络请求
- `Event_WebResourceResponseReceived`：接收所有网络响应
- `Event_NavigationCompleted`：页面导航完成时触发

### 请求过滤
支持两种请求过滤方式（根据运行时版本自动选择）：

1. **高级过滤（WebView2_22）**：支持按请求源类型过滤
   ```go
   wv.WebView2_22.AddWebResourceRequestedFilterWithRequestSourceKinds(...)
   ```

2. **基础过滤**：按 URL 模式过滤
   ```go
   wv.AddWebResourceRequestedFilter(...)
   ```

### 异步响应内容获取
使用 `GetContentEx` 异步获取响应内容，确保在并发场景下不会互相干扰。

## 使用场景

本示例适用于以下场景：

1. **数据采集**：拦截网页 API 请求，获取并保存数据
2. **请求修改**：修改网页请求参数，获取自定义数据
3. **调试监控**：监控网页的网络请求和响应
4. **自动化测试**：模拟用户操作并验证数据返回

## 运行环境

- 操作系统：Windows
- 运行时：需要安装 WebView2 Runtime
- Go 版本：建议 1.18+

## 运行示例

```bash
go run WebResourceRequestedEvent.go
```

程序启动后会：
1. 自动检查 WebView2 运行时版本
2. 打开东方财富网股票页面
3. 自动点击日K选项
4. 拦截并保存股票相关数据到当前目录的 JSON 文件

## 注意事项

1. **WebView2 版本要求**：部分高级功能（如 `AddWebResourceRequestedFilterWithRequestSourceKinds`）需要较高版本的 WebView2 运行时
2. **并发处理**：多个异步请求必须使用不同的 `WebViewEventImpl` 实例，否则回调函数会相互覆盖
3. **数据格式**：东方财富网的 API 返回 JSONP 格式数据，代码中会自动去除 `JSONP包装`（删除开头 `(` 和结尾 `);`）
4. **保存路径**：响应数据默认保存在程序所在目录，如需修改请调整 `prefixPath` 常量

## 关键代码说明

### 请求修改示例
```go
// 修改日K线数据条数
newUrl := strings.ReplaceAll(uri, "lmt=210", "lmt=570")
err = req.SetUri(newUrl)
```

### 响应获取示例
```go
// 异步获取响应内容
err = res.GetContentEx(getImpl(resType), func(errorCode syscall.Errno, content []byte) uintptr {
    // 处理响应内容
    contentStr := string(content)
    // ...
    return 0
})
```

### 多实例处理并发
```go
// 根据响应类型返回不同的 WebViewEventImpl
func getImpl(resType string) *edge.WebViewEventImpl {
    switch resType {
    case "日k线数据响应":
        return wv.GetWebViewEventImpl()
    case "个股信息查询响应":
        return impls[0]
    case "行业查询响应":
        return impls[1]
    }
    return nil
}
```

