# Chart - 多图表展示示例

> 本文由 AI 生成

## 项目简介

Chart 是一个基于 XCGUI 和 WebView2 技术的多图表展示应用示例。该示例演示了如何在单个窗口中创建多个 WebView 实例，分别显示不同类型的图表，并实现 Go 后端与前端 JavaScript 的数据交互。

## 功能特性

### 核心功能
- **多图表展示**: 在一个窗口中同时显示三种不同类型的图表
  - **静态折线图**: 展示预设的月度数据趋势
  - **静态饼图**: 显示数据比例分布  
  - **动态柱状图**: 支持实时数据更新的交互式图表
- **实时数据更新**: 通过按钮触发随机数据生成，动态更新第三个图表
- **数据传输**: Go 后端向前端发送 JSON 格式的图表数据

### 界面特性
- **多WebView布局**: 使用炫彩UI布局管理器排列三个WebView控件
- **响应式设计**: 图表自适应容器大小变化
- **现代化图表**: 基于 Chart.js 4.5.0 的专业图表组件
- **流畅动画**: Chart.js 内置的图表更新动画效果

### 技术特性
- **多实例管理**: 同时管理三个独立的WebView实例
- **虚拟主机映射**: 统一的资源访问方案
- **事件驱动架构**: 基于事件的数据更新机制
- **模块化设计**: 清晰的代码结构和功能分离

## 项目结构

```
Chart/
├── chart.go              # 主程序文件
├── 1.jpg                 # 窗口界面展示图片
├── assets/               # 前端资源目录
│   ├── css/             # 样式文件目录
│   │   └── Chart.css    # 通用图表样式
│   ├── js/              # JavaScript库目录
│   │   └── chart.js     # Chart.js图表库
│   ├── chart1.html      # 静态折线图页面
│   ├── chart2.html      # 静态饼图页面
│   └── chart3.html      # 动态柱状图页面
├── res/                 # 窗口布局资源
│   └── Chart.xml        # 窗口布局文件
└── README.md            # 本文档
```

## 技术实现

### 后端技术栈
- **Go 语言**: 主要编程语言
- **XCGUI**: GUI框架，提供窗口和布局管理
- **WebView2**: 嵌入式浏览器组件
- **embed**: Go 1.16+ 的文件嵌入功能

### 前端技术栈
- **HTML5**: 页面结构
- **CSS3**: 响应式样式设计  
- **Chart.js 4.5.0**: 专业图表库
- **WebView2 API**: JavaScript与Go的通信桥梁

### 图表类型实现

#### 1. 静态折线图 (chart1.html)
```javascript
// 展示月度数据趋势
type: 'line',
data: {
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
    datasets: [{
        label: 'Dataset 1',
        data: [65, 59, 80, 81, 56, 55, 40],
        // 样式配置...
    }]
}
```

#### 2. 静态饼图 (chart2.html)
```javascript
// 展示数据比例分布
type: 'pie',
data: {
    labels: ['Red', 'Blue', 'Yellow'],
    datasets: [{
        data: [12, 19, 3],
        // 颜色配置...
    }]
}
```

#### 3. 动态柱状图 (chart3.html)
```javascript
// 支持数据动态更新
type: 'bar',
// 监听来自Go的消息
window.chrome.webview.addEventListener('message', arg => {
    const labels = arg.data.map(item => item.label);
    const values = arg.data.map(item => item.value);
    // 更新图表数据...
});
```

### 关键技术点

#### 1. 多WebView管理
为每个图表创建独立的WebView实例：
```go
func createWebView1(edg *edge.Edge, wvOption []edge.WebViewOption) {
    layout_chart := widget.NewLayoutEleByName("layout_chart1")
    wv, err := edg.NewWebView(layout_chart.Handle, wvOption...)
    // 配置和事件处理...
}
```

#### 2. 数据结构定义
```go
type DataPoint struct {
    Label string  `json:"label"`
    Value float64 `json:"value"`
}
```

#### 3. 实时数据传输
Go向JavaScript发送JSON数据：
```go
data := []DataPoint{
    {"A", rand.Float64()*100 + 10},
    {"B", rand.Float64()*100 + 10},
    // ...
}
bs, _ := json.Marshal(data)
wv.PostWebMessageAsJSON(string(bs))
```

#### 4. 响应式CSS设计
```css
.chart-container {
    width: 100%;
    height: 100%;
    max-width: 90vw;
    max-height: 90vh;
    position: relative;
}
```

## 使用方法

### 运行程序
1. 确保已安装 WebView2 运行时
2. 编译并运行程序：
   ```bash
   go run chart.go
   ```

### 操作说明
1. **查看图表**: 程序启动后会显示三个不同的图表
2. **更新数据**: 点击"发送数据"按钮更新第三个柱状图的数据
3. **观察效果**: 柱状图会以动画形式更新为新的随机数据

## 依赖要求

### 系统要求
- Windows 10 1809 或更高版本
- WebView2 运行时（程序会自动检测并提示下载）

### Go 依赖
- github.com/twgh/xcgui
- Go 1.18+

### 前端依赖
- Chart.js 4.5.0（已嵌入到项目中）

## 设计亮点

### 架构设计
- **模块化WebView**: 每个图表独立管理，互不干扰
- **统一资源管理**: 通过虚拟主机名统一访问嵌入资源
- **事件驱动更新**: 基于WebView2消息机制的数据同步

### 用户体验
- **渐进式加载**: 等待每个图表完全加载后再显示
- **流畅动画**: Chart.js提供的专业图表动画效果
- **响应式布局**: 图表自动适应窗口大小变化

### 技术创新
- **多实例协同**: 展示了在单应用中管理多个WebView的最佳实践
- **实时数据流**: 演示了Go与JavaScript之间的高效数据传输
- **专业图表库**: 集成业界标准的Chart.js图表库

## 扩展可能性

这个示例可以作为基础，扩展为：
- **数据看板应用**: 实时监控系统的数据可视化
- **商业智能工具**: 多维度数据分析和展示
- **实时监控系统**: 服务器性能、网络状态等监控
- **金融数据分析**: 股票、期货等金融数据图表
- **IoT数据展示**: 物联网设备数据的实时可视化

## 学习价值

### 对于初学者
- 学习如何在桌面应用中集成Web技术
- 理解WebView2的基本使用方法
- 掌握Go与JavaScript的数据交互

### 对于进阶开发者
- 多WebView实例的管理和协调
- 复杂数据结构的序列化和传输
- 响应式图表界面的设计实现

## 注意事项

1. **调试模式**: 代码中的 `isDebug = true` 会启用开发者工具，生产环境建议设为 false

这个示例完美展示了 XCGUI + WebView2 + Chart.js 的强大组合，为构建现代化的数据可视化桌面应用提供了优秀的参考模板。