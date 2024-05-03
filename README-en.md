# XCGUI Examples

English | [简体中文](./README.md)

[https://github.com/twgh/xcgui](https://github.com/twgh/xcgui) Examples

# Usage

## 1. Download xcgui.dll to the system32 directory (if already downloaded, ignore this step)

#### (1) Download link

| x64  | [download](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest) |
| ---- | ------------------------------------------------------------ |
| x86  | [download](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest) |

#### (2) Command line download

x64

```bash
iwr https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest -OutFile xcgui.dll
```

x86

```
iwr https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest -OutFile xcgui.dll
```

#### (3) Network disk download

[download](https://wwi.lanzoup.com/b0cqd6nkb)

## 2. Clone project to local

#### (1) git clone

```bash
git clone https://github.com/twgh/xcgui-example
```

#### (2) If you don't have git, you can download the source zip and decompress it locally.

[download](https://codeload.github.com/twgh/xcgui-example/zip/refs/heads/main)

## 3. Execute commands in the project directory

```go
go mod tidy
cd SimpleWindow && go run .
```

# Visualization UI Designer

[![uidesigner](https://z3.ax1x.com/2021/09/15/4Vmh9S.png)](https://github.com/twgh/xcgui-example/blob/main/uidesigner/uidesigner.png)


# Simple Window

[![SimpleWindow](https://s1.ax1x.com/2022/05/24/XiEWtg.png)](https://github.com/twgh/xcgui-example/tree/main/SimpleWindow)

[![ButtonImage](https://s1.ax1x.com/2022/05/24/XiuLAx.jpg)](https://github.com/twgh/xcgui-example/tree/main/ButtonImage)
