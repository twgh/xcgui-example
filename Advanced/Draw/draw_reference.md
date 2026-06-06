# Draw 对象使用指南 — xcgui 图形绘制完全参考

> 版本: 基于 xcgui v1.4.0
> 包路径: `github.com/twgh/xcgui/drawx`
> 源码位置: `source/xcgui/drawx/draw.go`

---

## 一、Draw 对象的三种获取方式

### 1. 从窗口创建（主动创建）

```go
draw := drawx.New(hWindow)    // 创建独立 Draw 实例
// 或
draw := app.NewDraw(hWindow)  // 同上, 包级快捷函数
defer draw.Destroy()          // 用完手动释放
```

### 2. 从 Paint 事件回调获取（自绘场景, 最常用）

```go
w.AddEvent_Paint(func(hWindow int, hDraw int, pbHandled *bool) int {
    draw := drawx.NewByHandle(hDraw)
    // ... 绘制代码 ...
    return 0
})
```

> 注意: 此场景下 draw 不能调用 `Destroy()`, 由系统管理。

### 3. 从 GDI DC 创建

```go
draw := drawx.NewGDI(hWindow, hdc)  // 绑定 GDI 设备上下文
// 或
draw := app.NewDrawGDI(hWindow, hdc)
```

---

## 二、颜色系统

| 函数 | 说明 |
|------|------|
| `xc.RGBA(r, g, b, a byte) uint32` | 基础 RGBA 颜色 |
| `xc.HexRGB2RGBA("#1E88E5", 255) uint32` | 十六进制 → RGBA |
| `xc.RGB2RGBA(rgb, a byte) uint32` | RGB → RGBA |
| `xc.ParseRGBA(color) (r,g,b,a byte)` | RGBA → 分量 |

---

## 三、API 分类速查

### 画笔 / 画刷设置

| 方法 | 说明 |
|------|------|
| `SetBrushColor(color)` | 设置画刷颜色 |
| `SetLineWidth(w)` | 设置线宽 |
| `SetLineWidthF(w)` | 设置线宽（float） |
| `SetFont(hFont)` | 设置字体 |
| `SetTextAlign(flags)` | 设置文本对齐 |
| `SetTextVertical(b)` | 设置文本垂直显示 |
| `SetTextRenderingHint(t)` | 设置文本渲染提示 |
| `SetD2dTextRenderingMode(m)` | 设置 D2D 文本渲染模式 |

### 全局开关

| 方法 | 说明 |
|------|------|
| `EnableSmoothingMode(b)` | 启用平滑/抗锯齿（强烈建议开启） |
| `EnableWndTransparent(b)` | 启用窗口/元素背景透明 |

### 形状 — 填充

| 方法 | 说明 |
|------|------|
| `FillRect(pRect)` | 填充矩形 |
| `FillRectColor(pRect, color)` | 填充矩形（指定颜色） |
| `FillRoundRect(pRect, w, h)` | 填充圆角矩形 |
| `FillRoundRectEx(pRect, lt, rt, rb, lb)` | 填充四角独立圆角矩形 |
| `FillEllipse(pRect)` | 填充椭圆/圆形 |
| `FillPolygon(points, count)` | 填充多边形 |

### 形状 — 边框

| 方法 | 说明 |
|------|------|
| `DrawRect(pRect)` | 矩形边框 |
| `DrawRoundRect(pRect, w, h)` | 圆角矩形边框 |
| `DrawRoundRectEx(pRect, lt, rt, rb, lb)` | 四角独立圆角边框 |
| `DrawEllipse(pRect)` | 椭圆边框 |
| `DrawPolygon(points, count)` | 多边形 |
| `DrawLine(x1, y1, x2, y2)` | 直线 |
| `Dottedline(x1, y1, x2, y2)` | 虚线（水平/垂直） |
| `DrawArc(x, y, w, h, start, sweep)` | 圆弧 |
| `DrawCurve(points, count, tens)` | 曲线（D2D 暂不支持） |
| `FocusRect(pRect)` | 焦点虚线框 |

### 渐变

| 方法 | 说明 |
|------|------|
| `GradientFill2(pRect, c1, c2, mode)` | 双色渐变 |
| `GradientFill4(pRect, c1, c2, c3, c4, mode)` | 四色渐变 |

`mode` 取值:
- `xcc.GRADIENT_FILL_RECT_H` — 水平填充
- `xcc.GRADIENT_FILL_RECT_V` — 垂直填充

### 文本

| 方法 | 说明 |
|------|------|
| `TextOutEx(x, y, str)` | 简单文本 |
| `DrawText(str, pRect)` | 矩形内文本（自动换行） |
| `DrawTextUnderline(str, pRect, c)` | 带下划线文本 |
| `TextOut(x, y, str, cbStr)` | GDI 方式 |
| `TextOutA(x, y, str)` | GDI 方式 |

### 图片

| 方法 | 说明 |
|------|------|
| `Image(hImg, x, y)` | 原始大小 |
| `XDraw_ImageEx(hImg, x, y, w, h)` | 指定宽高 |
| `ImageAdaptive(hImg, pRect, border)` | 自适应/九宫格 |
| `ImageSuper(hImg, pRect, clip)` | 增强绘制 |
| `ImageSuperEx(hImg, dst, src)` | 增强绘制（源区域） |
| `ImageTile(hImg, mask, pRect, flag)` | 平铺 |
| `ImageMask(hImg, mask, pRect, x, y)` | 遮盖绘制 |
| `ImageMaskRect(hImg, rect, mask, rd)` | 矩形遮罩 |
| `ImageMaskEllipse(hImg, rect, mask)` | 圆形遮罩 |

### SVG

| 方法 | 说明 |
|------|------|
| `DrawSvgSrc(hSvg)` | 原始位置绘制 |
| `DrawSvg(hSvg, x, y)` | 指定坐标 |
| `DrawSvgEx(hSvg, x, y, w, h)` | 指定坐标+大小 |
| `DrawSvgSize(hSvg, w, h)` | 指定大小 |

### 坐标 / 裁剪 / 偏移

| 方法 | 说明 |
|------|------|
| `SetClipRect(pRect)` | 设置裁剪区域 |
| `ClearClip()` | 清除裁剪 |
| `SetOffset(x, y)` | 设置绘制偏移 |
| `GetOffset(&x, &y)` | 获取偏移 |
| `ConvRect(pRect)` | DPI 坐标转换 |
| `ConvXY(x, y)` | DPI 坐标转换 |

### D2D 专用

| 方法 | 说明 |
|------|------|
| `D2D_Clear(color)` | 清理画布 |
| `GetD2dBitmap()` | 获取 D2D 位图指针 |
| `XDraw_GetD2dRenderTarget()` | 获取 D2D 渲染目标 |

### GDI 专用

| 方法 | 说明 |
|------|------|
| `GetHDC()` | 获取 HDC |
| `GDI_SetBkMode(transparent)` | 设置背景模式 |
| `GDI_CreateSolidBrush(crColor)` | 创建实心画刷 |
| `GDI_CreatePen(style, width, color)` | 创建画笔 |
| `GDI_CreateRectRgn(l, t, r, b)` | 创建矩形区域 |
| `GDI_CreateRoundRectRgn(...)` | 创建圆角区域 |
| `GDI_SelectClipRgn(hRgn)` | 选择裁剪区 |
| `GDI_FillRgn(hrgn, hbr)` | 填充区域 |
| `GDI_FrameRgn(hrgn, hbr, w, h)` | 区域边框 |
| `GDI_Rectangle(l, t, r, b)` | GDI 矩形 |
| `GDI_Ellipse(pRect)` | GDI 椭圆 |
| `GDI_MoveToEx(x, y, &pt)` | GDI 移动起点 |
| `GDI_LineTo(x, y)` | GDI 画线 |
| `GDI_Polyline(pts, count)` | GDI 折线 |
| `GDI_SetPixel(x, y, color)` | GDI 像素 |
| `GDI_BitBlt(...)` | 位图复制 |
| `GDI_AlphaBlend(...)` | 透明混合复制 |
| `GDI_DrawIconEx(...)` | 图标绘制 |
| `GDI_RestoreGDIOBJ()` | 还原 GDI 对象 |

### 其他

| 方法 | 说明 |
|------|------|
| `GetFont()` | 获取当前字体句柄 |
| `GetFontObj()` | 获取当前字体对象 |
| `Destroy()` | 销毁（仅主动创建的 Draw 需要） |

---

## 四、使用场景与事件

### 场景1: 窗口级自绘

```go
w.AddEvent_Paint(func(hWindow int, hDraw int, pbHandled *bool) int {
    *pbHandled = true // 拦截原本的绘制
    xc.XWnd_DrawWindow(hWindow, hDraw)
    draw := drawx.NewByHandle(hDraw)
    // 自定义窗口背景、叠加图案等
    return 0
})
```

### 场景2: 元素级自绘

任何元素（Button, Edit, List...）都支持 Paint 事件，配合 `EnableBkTransparent(true)` 完全自定义外观(或者使用`*pbHandled = true` 拦截默认的绘制)：

```go
btn.AddEvent_Paint(func(hEle int, hDraw int, pbHandled *bool) int {
    draw := drawx.NewByHandle(hDraw)
    // 自绘元素外观
    return 0
})
```

### 场景3: 菜单自绘

窗口事件:
- `Event_MENU_DRAW_BACKGROUND(func(hWindow, hDraw, *Menu_DrawBackground_, *bool) int)`
- `Event_MENU_DRAWITEM(func(hWindow, hDraw, *Menu_DrawItem_, *bool) int)`

元素事件:
- `Event_MENU_DRAW_BACKGROUND(...)`
- `Event_MENU_DRAWITEM(...)`

---

## 五、重要注意事项

1. **链式调用**: 大多数方法返回 `*Draw` 以支持链式调用
   ```go
   draw.SetBrushColor(...).SetLineWidth(2).FillRoundRect(...)
   ```
   例外: `Image()`, `Destroy()`, `GetFontObj()` 等不返回 `*Draw`

2. **抗锯齿**: `EnableSmoothingMode(true)` 在 D2D 模式下很重要，否则图形边缘会有锯齿

3. **画刷颜色持久性**: 设置新画刷颜色后，后续所有填充都使用该颜色，直到再次 `SetBrushColor`；不想改变画刷颜色请用 `FillRectColor`

4. **窗口 Paint 事件顺序**: 使用`*pbHandled = true` 拦截默认的绘制, 后调用 `xc.XWnd_DrawWindow(hWindow, hDraw)` 绘制窗口默认内容，再叠加自定义绘制

5. **重绘**: 修改界面元素后记得调用 `Redraw(false)` 触发重绘

6. **D2D vs GDI 差异**:
   - 曲线绘制 (`DrawCurve`) D2D 暂不支持
   - 图片遮盖 (`ImageSuperMask` / `ImageMask`) 在 D2D 下留空
   - `GradientFill2` 的 `GRADIENT_FILL_TRIANGLE` 模式仅在 GDI 有效
