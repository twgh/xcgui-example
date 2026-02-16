# SaveMemory - WebView 内存优化示例

本示例演示如何通过两种方式减少 WebView2 的内存占用。

## 内存优化方法

### 1. 设置低内存目标级别

```go
wv.WebView2_19.SetMemoryUsageTargetLevel(edge.COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_LOW)
```

此方法可将内存占用降低近一半。

### 2. 挂起 WebView

当窗口最小化一段时间后，挂起 WebView 可将内存占用降至个位数：

- 窗口最小化 10 秒后自动挂起
- 窗口恢复时自动从挂起状态恢复

**注意**：正在导航时不应挂起，否则恢复后 WebView 可能不可见。

## 运行要求

- Windows 系统
- 已安装 WebView2 运行时
- Go 1.18+

## 运行方式

```bash
go run SaveMemory.go
```
