# XCGUI例子

[English](./README-en.md) | 简体中文

[https://github.com/twgh/xcgui](https://github.com/twgh/xcgui) 的例子

# 用法
## 一、下载xcgui.dll到系统system32目录(如已下载则忽略这步)

#### （1）文件直链

| 64位 | [点击下载](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest) |
| ---- | ------------------------------------------------------------ |
| 32位 | [点击下载](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest) |

#### （2）命令行下载

64位

```bash
curl -fL "https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest" -o xcgui.dll
```

32位

```bash
curl -fL "https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest" -o xcgui.dll
```

#### （3）使用getxcgui工具下载

> 请确保 `%GOPATH%\bin` 在环境变量`path`中

```bash
go install github.com/twgh/getxcgui@latest
getxcgui  
```

如果要把dll直接下载到`C:\Windows\System32`目录里，请使用如下命令：

```bash
getxcgui -o %windir%\system32\xcgui.dll
```

此工具的源码在[这里](https://github.com/twgh/getxcgui)，更多flags可以点[进去](https://github.com/twgh/getxcgui#flags)查看

#### （4）网盘下载

网盘内还包含`界面设计器`和`chm帮助文档`

| 网盘         | 下载地址                                                     |
| ------------ | ------------------------------------------------------------ |
| 百度网盘     | [下载](https://pan.baidu.com/s/1rC3unQGaxnRUCMm8z8qzvA?pwd=1111) |
| 蓝奏云     | [下载](https://wwi.lanzoup.com/b0cqd6nkb) |

## 二、克隆项目到本地

#### （1）git克隆

```bash
git clone https://github.com/twgh/xcgui-example
```

#### （2）没有git的可以下载源码zip到本地后解压

[点击下载](https://codeload.github.com/twgh/xcgui-example/zip/refs/heads/main)

## 三、在项目目录里执行命令

```bash
go mod tidy
cd SimpleWindow && go run .
```

# 可视化UI设计器

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/tree/main/uidesigner)

# 简单窗口

[![SimpleWindow](https://s1.ax1x.com/2022/05/24/XiEWtg.png)](https://github.com/twgh/xcgui-example/tree/main/SimpleWindow)

[![ButtonImage](https://s1.ax1x.com/2022/05/24/XiuLAx.jpg)](https://github.com/twgh/xcgui-example/tree/main/ButtonImage)

