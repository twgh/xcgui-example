# BindTypes — Bind 类型演示

> 本文由 AI 生成

演示 WebView 的 `Bind` 函数所支持的参数和返回值类型。

## 功能

通过可视化的 Web 界面展示了以下类型的绑定：

| 分类 | 函数名 | Go 签名字符串 | 说明 |
|------|--------|--------------|------|
| **基本类型** | `demo.stringParam` | `func(string) string` | string 参数 + 返回值 |
| | `demo.intParam` | `func(int, int) int` | int 参数 + 返回值 |
| | `demo.floatParam` | `func(float64, float64) float64` | float64 参数 + 返回值 |
| | `demo.boolParam` | `func(bool) bool` | bool 参数 + 返回值 |
| | `demo.noop` | `func()` | 无参数、无返回值 |
| | `demo.hello` | `func() string` | 无参数、有返回值 |
| **Slice** | `demo.sliceInt` | `func([]int) []int` | []int 参数 + 返回值 |
| | `demo.sliceString` | `func([]string) string` | []string 参数 + string 返回值 |
| **Map** | `demo.mapParam` | `func(map[string]interface{}) string` | map 参数 + 返回值 |
| **Struct** | `demo.structParam` | `func(Person) Person` | struct 参数 + 返回值 |
| | `demo.companyInfo` | `func(Company) Company` | 嵌套 struct 参数 + 返回值 |
| **error** | `demo.divideOK` | `func(float64, float64) (float64, error)` | 返回值和 error (成功) |
| | `demo.divideFail` | `func(float64, float64) (float64, error)` | 返回值和 error (失败) |
| | `demo.checkPositive` | `func(int) error` | 仅返回 error |
| **混合参数** | `demo.mixedArgs` | `func(string, int, []int, map[string]interface{}) Person` | 混合多种类型 |
| | `demo.calcSum` | `func([]int) int` | Slice 参数 |

## 运行

在项目根目录执行：

```bash
go run .\webview\BindTypes\BindTypes.go
```

## 文件结构

```
webview/BindTypes/
├── BindTypes.go           # Go 主文件
├── README.md              # 本文件
└── assets/
    └── BindTypes.html     # 前端 UI 页面
```

## 关于 Bind

`Bind` 是 `edge.WebView` 绑定的方法，用于将 Go 函数暴露为 JavaScript 全局函数。实现原理是通过 JSON-RPC 经 `window.external.invoke` 进行通信：

- 参数通过 `json.Unmarshal` 从 JSON 反序列化为 Go 类型
- 返回值通过 `json.Marshal` 序列化为 JSON 返回给 JS，以 Promise 形式接收
- 暂不支持的、不适合的类型：`[]byte`、`uintptr` 等
