# XCGUI例子

[English](./README-en.md) | 简体中文

[https://github.com/twgh/xcgui](https://github.com/twgh/xcgui) 的例子

# 获取

```bash
git clone https://github.com/twgh/xcgui-example
cd xcgui-example && go mod tidy
go install -ldflags="-s -w" github.com/twgh/getxcgui@latest
getxcgui -o %windir%\system32\xcgui.dll
cd SimpleWindow && go run .
```

# 可视化UI设计器

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/tree/main/uidesigner)

# 简单窗口

[![SimpleWindow](https://s1.ax1x.com/2022/05/24/XiEWtg.png)](https://github.com/twgh/xcgui-example/tree/main/SimpleWindow)
[![ButtonImage](https://s1.ax1x.com/2022/05/24/XiuLAx.jpg)](https://github.com/twgh/xcgui-example/tree/main/ButtonImage)

