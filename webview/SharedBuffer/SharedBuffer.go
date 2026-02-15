// 共享缓冲区.
// SharedBuffer 的核心优势在于它提供了一种高效的方式来共享内存数据，避免了数据复制的开销。这使得它在需要频繁或大量数据交换的场景中非常有用.
// 可用于高效的传输大量数据（如大型 JSON 对象、二进制数据块）, 还可以跨进程通信.
package main

import (
	"embed"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/edge"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed assets/**
	embedAssets embed.FS // 嵌入 assets 目录以及子目录下的文件, 不包括隐藏文件
	//go:embed res/1.png
	img1 []byte
	// 图片共享缓冲区发送者
	imgSender *SharedBufferSender
)

const hostName = "app.example"

type MainWindow struct {
	edg          *edge.Edge
	w            *window.Window
	wv           *edge.WebView
	btnSelectImg *widget.Button
	btnSendText  *widget.Button
}

func NewMainWindow(edg *edge.Edge) *MainWindow {
	m := &MainWindow{edg: edg}

	// 创建窗口
	m.w = window.New(0, 0, 1200, 900, "共享缓冲区例子", 0, xcc.Window_Style_Default)
	m.w.EnableLayout(true)         // 窗口启用布局
	m.w.SetBorderSize(2, 30, 2, 2) // 设置窗口边框大小

	// 创建放置顶部按钮的布局
	layoutTop := widget.NewLayoutEle(0, 0, 0, 34, m.w.Handle)
	layoutTop.LayoutItem_SetWidth(xcc.Layout_Size_Fill, 0) // 宽度填充父
	layoutTop.SetSpace(2)                                  // 子项间距
	layoutTop.SetPadding(2, 1, 2, 1)                       // 内填充
	layoutTop.AddBkBorder(xcc.Element_State_Flag_Leave, xc.RGBA(224, 224, 224, 255), 1)

	// 按钮_选择图片
	m.btnSelectImg = widget.NewButton(0, 0, 100, 30, "选择图片", layoutTop.Handle)
	// 按钮_发送文本
	m.btnSendText = widget.NewButton(0, 0, 100, 30, "发送文本", layoutTop.Handle)

	// 创建放置 WebView 的布局
	layoutWV := widget.NewLayoutEle(0, 0, 0, 0, m.w.Handle)
	layoutWV.LayoutItem_SetWidth(xcc.Layout_Size_Fill, 0)    // 宽度填充父
	layoutWV.LayoutItem_SetHeight(xcc.Layout_Size_Weight, 1) // 高度占据剩余空间

	// 创建 WebView
	m.createWebView(layoutWV.Handle)

	// 注册炫彩事件
	m.regXcEvents()

	// 显示窗口
	m.w.Show(true)
	return m
}

// 注册炫彩事件
func (m *MainWindow) regXcEvents() {
	m.btnSelectImg.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if m.wv.CoreWebView == nil {
			return 0
		}
		if imgSender == nil {
			m.w.MessageBox("提示", "图片缓冲区发送者没有创建成功", xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}

		// 选择文件
		filePath := wutil.OpenFile(m.w.Handle, []string{"图片文件 (*.png;*.jpg;*.jpeg;*.bmp;*.gif)", "*.png;*.jpg;*.jpeg;*.bmp;*.gif"}, "")
		if filePath == "" {
			return 0
		}
		// 读取文件
		data, err := os.ReadFile(filePath)
		if err != nil {
			m.w.MessageBox("提示", "读取图片失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}
		// 发送数据给 WebView
		err = imgSender.Send(data)
		if err != nil {
			m.w.MessageBox("提示", "发送图片给 WebView 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}
		return 0

		// 发送数据给 WebView
		/* err = sendData(wv, data, "img")
		if err != nil {
			w.MessageBox("提示", "发送图片给 WebView 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}
		return 0 */
	})

	m.btnSendText.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if m.wv.CoreWebView == nil {
			return 0
		}
		// 生成随机字符串
		data := []byte(generateRandomString(100))

		// 发送数据给 WebView
		err := sendData(m.wv, data, "text")
		if err != nil {
			m.w.MessageBox("提示", "发送文本给 WebView 失败: "+err.Error(), xcc.MessageBox_Flag_Ok, xcc.Window_Style_Default)
			return 0
		}
		return 0
	})
}

// 创建 WebView
func (m *MainWindow) createWebView(hParent int) {
	var err error
	// 设置虚拟主机名和嵌入文件系统之间的映射
	edge.SetVirtualHostNameToEmbedFSMapping(hostName, embedAssets)

	// 创建 WebView
	m.wv, err = m.edg.NewWebView(hParent,
		edge.WithFillParent(true),
		edge.WithDebug(true),
	)
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView 失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(2)
	}

	// 创建图片的共享缓冲区发送者, 这个不是一次性的缓冲区
	imgSender, err = NewSharedBufferSender(m.wv, 20*1024*1024) // 20MB
	if err != nil {
		log.Println("创建图片的共享缓冲区发送者失败: " + err.Error())
	} else {
		imgSender.SetAdditionalDataAsJson(`{"type":"img"}`)
	}

	// 在宿主原生窗口销毁时释放图片缓冲区.
	// 因为 WebView2 并没有销毁事件, 所以用宿主原生窗口的销毁事件来释放图片缓冲区.
	m.wv.Event_Destroy(func(wv *edge.WebView) {
		if imgSender != nil {
			imgSender.Close()
		}
	})

	// 注册 WebView 事件
	m.regWebViewEvents()

	// 启用虚拟主机名和嵌入文件系统之间的映射
	m.wv.EnableVirtualHostNameToEmbedFSMapping(true)

	// 导航
	m.wv.Navigate(edge.JoinUrlHeader(hostName) + "/SharedBuffer.html")
}

// 注册 WebView 事件
func (m *MainWindow) regWebViewEvents() {
	// 网页消息事件
	m.wv.Event_WebMessageReceived(func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2WebMessageReceivedEventArgs) uintptr {
		msg, err := args.TryGetWebMessageAsString()
		if err != nil {
			log.Println("Event_WebMessageReceived, TryGetWebMessageAsString: " + err.Error())
			return 0
		}

		switch msg {
		case "RequestImg": // 收到前端的图片请求, 通过共享缓冲区发送图片给 WebView
			if imgSender != nil {
				imgSender.Send(img1)
			}
			/* err = sendData(wv, img1, "img")
			if err != nil {
				log.Println("发送图片给 WebView 失败: " + err.Error())
			} */
		case "RequestText": // 收到前端的文本请求, 通过共享缓冲区发送文本给 WebView
			err = sendData(m.wv, []byte(generateRandomString(100)), "text")
			if err != nil {
				log.Println("发送文本给 WebView 失败: " + err.Error())
			}
		}
		return 0
	})
}

// SharedBufferSender 共享缓冲区发送者.
type SharedBufferSender struct {
	wv2_17               *edge.ICoreWebView2_17
	buf                  *edge.ICoreWebView2SharedBuffer
	stream               *edge.IStream
	additionalDataAsJson string
	bufSize              uint64
	access               edge.COREWEBVIEW2_SHARED_BUFFER_ACCESS
}

// NewSharedBufferSender 创建共享缓冲区发送者.
//   - 默认是只读权限, 在 js 中修改缓冲区数据会报错.
//   - 不再使用时需要调用 Close 方法关闭.
//
// wv: edge.WebView 对象.
//
// size: 缓冲区大小, 单位是字节.
func NewSharedBufferSender(wv *edge.WebView, size uint64) (*SharedBufferSender, error) {
	// 获取 WebView2 环境12
	env12, err := wv.Edge.Environment.GetICoreWebView2Environment12()
	if err != nil {
		return nil, errors.New("GetICoreWebView2Environment12 失败: " + err.Error())
	}
	defer env12.Release()

	// 创建共享缓冲区
	buf, err := env12.CreateSharedBuffer(size)
	if err != nil {
		return nil, errors.New("创建缓冲区失败: " + err.Error())
	}

	// 创建缓冲区流
	stream, err := buf.OpenStream()
	if err != nil {
		buf.Release()
		return nil, errors.New("创建缓冲区流失败: " + err.Error())
	}

	// 获取 WebView2_17
	wv2_17, err := wv.CoreWebView.GetICoreWebView2_17()
	if err != nil {
		stream.Release()
		buf.Release()
		return nil, errors.New("GetICoreWebView2_17 失败: " + err.Error())
	}

	p := &SharedBufferSender{
		buf:     buf,
		stream:  stream,
		wv2_17:  wv2_17,
		bufSize: size,
		access:  edge.COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY,
	}
	return p, nil
}

// GetBufferSize 获取缓冲区大小.
func (s *SharedBufferSender) GetBufferSize() uint64 {
	return s.bufSize
}

// Send 发送数据给 WebView.
//   - 缓冲区的前4个字节固定是数据长度, 后面是数据, js 读取数据的时候要注意这一点. 这个是自己定义的规则, 或者说协议, 因为这样在 js 里才知道读取多少数据.
//
// data: 数据.
//
// additionalDataAsJson: 附加 json 数据, 可不填.
//   - 不填时使用 SetAdditionalDataAsJson 方法设置的 json 文本, 没设置过的话是空字符串.
//   - 填了就相当于本次临时使用的 json 文本, 不会改变使用 SetAdditionalDataAsJson 方法设置的 json 文本.
func (s *SharedBufferSender) Send(data []byte, additionalDataAsJson ...string) error {
	jsonStr := s.additionalDataAsJson
	if len(additionalDataAsJson) > 0 {
		jsonStr = additionalDataAsJson[0]
	}

	if data == nil {
		return s.wv2_17.PostSharedBufferToScript(s.buf, s.access, jsonStr)
	}

	dataLength := len(data)
	// 数据长度+4个字节不能超过缓冲区大小
	if dataLength+4 > int(s.bufSize) {
		return errors.New("数据长度超过缓冲区大小")
	}
	// 清空流数据
	err := s.stream.Clear()
	if err != nil {
		return errors.New("清空缓冲区数据失败: " + err.Error())
	}
	// 将数据长度写到前4个字节
	_, err = s.stream.Write(Uint32ToBytes(uint32(dataLength)))
	if err != nil {
		return errors.New("缓冲区写入数据长度失败: " + err.Error())
	}
	// 将数据写入缓冲区
	_, err = s.stream.Write(data)
	if err != nil {
		return errors.New("缓冲区写入数据失败: " + err.Error())
	}
	return s.wv2_17.PostSharedBufferToScript(s.buf, s.access, jsonStr)
}

// SetAccess 设置缓冲区访问权限.
func (s *SharedBufferSender) SetAccess(access edge.COREWEBVIEW2_SHARED_BUFFER_ACCESS) *SharedBufferSender {
	s.access = access
	return s
}

// GetAccess 获取缓冲区访问权限.
func (s *SharedBufferSender) GetAccess() edge.COREWEBVIEW2_SHARED_BUFFER_ACCESS {
	return s.access
}

// SetAdditionalDataAsJson 设置附加 JSON 数据.
func (s *SharedBufferSender) SetAdditionalDataAsJson(json string) *SharedBufferSender {
	s.additionalDataAsJson = json
	return s
}

// GetAdditionalDataAsJson 获取附加 JSON 数据.
func (s *SharedBufferSender) GetAdditionalDataAsJson() string {
	return s.additionalDataAsJson
}

// Close 关闭缓冲区, 无法再使用.
func (s *SharedBufferSender) Close() {
	if s.wv2_17 != nil {
		// 发送释放缓冲区的消息给 js
		s.Send(nil, `{"type":"close"}`)
		s.wv2_17.Release()
		s.wv2_17 = nil
	}
	if s.stream != nil {
		s.stream.Release()
		s.stream = nil
	}
	if s.buf != nil {
		s.buf.Close()
		s.buf.Release()
		s.buf = nil
	}
}

// 通过共享缓冲区发送数据给 WebView, 缓冲区大小是 data 的长度, 一次性的共享缓冲区, 发完就关闭了.
func sendData(wv *edge.WebView, data []byte, typeStr string) error {
	// 获取 WebView2 环境12
	env12, err := wv.Edge.Environment.GetICoreWebView2Environment12()
	if err != nil {
		return errors.New("GetICoreWebView2Environment12 失败: " + err.Error())
	}
	defer env12.Release()

	// 创建共享缓冲区
	buf, err := env12.CreateSharedBuffer(uint64(len(data)))
	if err != nil {
		return errors.New("创建缓冲区失败: " + err.Error())
	}
	defer func() {
		buf.Close() // js里面也要调用释放缓冲区, 只有两边都释放了, 操作系统才会释放底层的共享内存。
		buf.Release()
	}()

	// 创建缓冲区流
	stream, err := buf.OpenStream()
	if err != nil {
		return errors.New("创建缓冲区流失败: " + err.Error())
	}
	defer stream.Release()

	// 将数据写入缓冲区
	_, err = stream.Write(data)
	if err != nil {
		return errors.New("缓冲区写入数据失败: " + err.Error())
	}

	// 获取 WebView2_17
	if wv.WebView2_17 == nil {
		wv.WebView2_17, err = wv.CoreWebView.GetICoreWebView2_17()
		if err != nil {
			return errors.New("GetICoreWebView2_17 失败: " + err.Error())
		}
	}

	// 发送缓冲区给 WebView2, 只读权限
	err = wv.WebView2_17.PostSharedBufferToScript(buf, edge.COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY, `{"type":"`+typeStr+`"}`)
	if err != nil {
		return errors.New("发送缓冲区给 WebView2 失败: " + err.Error())
	}
	return nil
}

func main() {
	checkWebView2()
	// 创建 WebView2 环境
	edg := createEdge()

	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	NewMainWindow(edg)

	a.Run()
	a.Exit()
}

// 创建 WebView2 环境
func createEdge() *edge.Edge {
	// 创建 WebView2 环境
	edg, err := edge.New(edge.Option{
		UserDataFolder: os.TempDir(), // 实际应用中应使用自己创建的固定目录
		EnvOptions: &edge.EnvOptions{
			DisableTrackingPrevention: true,
		},
	})
	if err != nil {
		wapi.MessageBoxW(0, "创建 WebView2 环境失败: "+err.Error(), "错误", wapi.MB_OK|wapi.MB_IconError)
		os.Exit(1)
	}
	return edg
}

func checkWebView2() {
	// 输出本库使用的 WebView2 版本
	fmt.Println("本库使用的 WebView2 运行时版本号:", edge.GetVersion())

	// 获取本机已安装的 WebView2 运行时版本
	localVersion, err := edge.GetAvailableBrowserVersion()
	if err != nil {
		wapi.MessageBoxW(0, "获取 WebView2 运行时版本号失败: "+err.Error(), "提示", wapi.MB_IconError)
		os.Exit(1)
	}
	if localVersion == "" {
		wapi.MessageBoxW(0, "请安装 WebView2 运行时后再打开程序! 下载完后请使用管理员权限运行安装包!", "提示", wapi.MB_IconWarning|wapi.MB_OK)
		// 打开 WebView2 运行时下载页面
		edge.DownloadWebView2()
		os.Exit(2)
	}
	fmt.Println("本机安装的 WebView2 运行时版本号:", localVersion)

	// 检查本机版本是否低于库版本
	if ret, _ := edge.CompareBrowserVersions(localVersion, edge.GetVersion()); ret == -1 {
		log.Println("本机 WebView2 运行时版本低于本库使用的 WebView2 运行时版本!")
	}
}

// 生成随机字符串
func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Uint32ToBytes 将 uint32 转换为字节数组.
func Uint32ToBytes(i uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	return buf
}
