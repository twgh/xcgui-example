// 音频播放器
package main

import (
	"fmt"
	"log"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/wapi/wutil"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	// 初始化界面库
	app.InitOrExit()
	a := app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)

	NewMainWindow()

	a.Run()
	a.Exit()
}

type MainWindow struct {
	w              *window.Window
	ap             *wutil.AudioPlayer
	btnPlay        *widget.Button
	btnPause       *widget.Button
	btnSelectSound *widget.Button
	editSoundPath  *widget.Edit
	sliderVolume   *widget.SliderBar
	progressBar    *widget.ProgressBar
	textTime       *widget.ShapeText
	volPct         *widget.ShapeText
}

func NewMainWindow() *MainWindow {
	w := window.New(0, 0, 600, 300, "音频播放器", 0, xcc.Window_Style_Default)

	w.EnableLayout(true)                             // 启用布局
	w.SetPadding(10, 15, 10, 10)                     // 设置布局内填充
	w.SetSpace(20).SetSpaceRow(20)                   // 布局盒子设置行间距和列间距
	w.SetAlignBaseline(xcc.Layout_Align_Axis_Center) // 布局盒子_置对齐基线, 垂直居中

	// 创建音频播放器
	ap := wutil.NewAudioPlayer()

	// 标签_音频路径
	widget.NewShapeText(0, 0, 60, 30, "音频路径:", w.Handle)
	// 编辑框_音频路径
	editSoundPath := widget.NewEdit(0, 0, 360, 30, w.Handle)

	// 按钮_选择音频
	btnSelectSound := widget.NewButton(0, 0, 100, 30, "选择音频", w.Handle)
	// 按钮_播放/停止
	btnPlay := widget.NewButton(0, 0, 100, 30, "播放", w.Handle)
	btnPlay.Enable(false)
	// 按钮_暂停/继续
	btnPause := widget.NewButton(0, 0, 100, 30, "暂停", w.Handle)
	btnPause.Enable(false)

	// 水平进度条
	progressBar := widget.NewProgressBar(0, 0, 280, 12, w.Handle)
	progressBar.SetRange(100)
	progressBar.SetPos(0)
	progressBar.EnableShowText(false)
	progressBar.SetBorderSize(1, 1, 1, 1)
	progressBar.SetColorLoad(xc.RGBA(43, 170, 255, 255))
	progressBar.AddBkFill(xcc.Element_State_Flag_Leave, xc.RGBA(221, 221, 223, 255))
	// 窗口组件_布局项_启用换行, 强制换行
	progressBar.LayoutItem_EnableWrap(true)

	// 时间文本
	textTime := widget.NewShapeText(0, 0, 120, 20, "00:00 / 00:00", w.Handle)

	// 音量标签
	widget.NewShapeText(0, 0, 40, 24, "音量:", w.Handle).LayoutItem_EnableWrap(true)
	// 水平滑块条
	sliderVolume := widget.NewSliderBar(0, 0, 200, 24, w.Handle)
	sliderVolume.SetRange(100)
	sliderVolume.SetPos(80) // 默认音量 80%
	sliderVolume.SetButtonWidth(20)
	sliderVolume.SetButtonHeight(20)
	sliderVolume.EnableBkTransparent(true).EnableDrawFocus(false)

	// 音量百分比文本
	volPct := widget.NewShapeText(0, 0, 40, 24, "80%", w.Handle)

	m := &MainWindow{
		w:              w,
		ap:             ap,
		btnPlay:        btnPlay,
		btnPause:       btnPause,
		btnSelectSound: btnSelectSound,
		editSoundPath:  editSoundPath,
		sliderVolume:   sliderVolume,
		progressBar:    progressBar,
		textTime:       textTime,
		volPct:         volPct,
	}
	// 注册事件
	m.regEvents()

	w.Show(true)
	return m
}

func (m *MainWindow) regEvents() {
	// 窗口关闭事件
	m.w.AddEvent_Close(func(hWindow int, pbHandled *bool) int {
		m.closeAudioPlayer() // 关闭音频播放器
		return 0
	})

	// 按钮_选择音频, 单击事件
	m.btnSelectSound.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		filePath := wutil.OpenFile(m.w.Handle, []string{"音频文件(*.mp3;*.wav;*.wma)", "*.mp3;*.wav;*.wma"}, "%USERPROFILE%\\Music")
		if filePath == "" {
			return 0
		}
		m.editSoundPath.SetText(filePath).Redraw(false)

		// 先关闭上一个
		m.closeAudioPlayer()

		// 打开音频文件
		err := m.ap.Open(filePath)
		if err != nil {
			m.w.MessageBox("提示", "打开音频文件失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Modal)
			m.btnPlay.Enable(false).Redraw(false)
			return 0
		}

		m.btnPlay.Enable(true).Redraw(false)
		return 0
	})

	// 按钮_播放/停止, 单击事件
	m.btnPlay.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		switch m.btnPlay.GetText() {
		case "播放":
			m.playAudio()
		case "停止":
			m.stopAudio()
		}
		return 0
	})

	// 按钮_暂停/继续, 单击事件
	m.btnPause.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		switch m.btnPause.GetText() {
		case "暂停":
			m.pauseAudio()
		case "继续":
			m.resumeAudio()
		}
		return 0
	})

	// 音量滑块, 滑块位置改变事件
	m.sliderVolume.AddEvent_SliderBar_Change(func(hEle int, pos int32, pbHandled *bool) int {
		// 更新音量百分比显示
		m.volPct.SetText(fmt.Sprintf("%d%%", pos))
		m.volPct.Redraw()
		// 在播放中才设置音量
		if m.ap.IsPlaying() || m.ap.IsPaused() {
			_ = m.ap.SetVolume(int(pos) * 10)
		}
		return 0
	})
}

// 关闭音频播放器
func (m *MainWindow) closeAudioPlayer() {
	if m.ap.Alias == "" {
		return
	}

	// 如果没有停止, 最好先停止
	if !m.ap.IsStopped() {
		err := m.ap.Stop()
		if err != nil {
			log.Println("停止播放音频时报错: " + err.Error())
		} else {
			log.Println("停止播放音频成功")
		}
	}

	// 关闭音频设备
	err := m.ap.Close()
	if err != nil {
		log.Println("关闭音频设备时报错: " + err.Error())
	} else {
		log.Println("关闭音频设备成功")
	}
}

// 播放音频
func (m *MainWindow) playAudio() {
	m.btnPlay.SetText("停止").Redraw(false)

	vol := int(m.sliderVolume.GetPos()) * 10 // 滑块 0-100 → 音量 0-1000
	err := m.ap.Play(wutil.PlayOptions{
		Volume:      &vol,
		Wait:        false,
		Repeat:      false,
		SeekToStart: true,
	})
	if err != nil {
		m.w.MessageBox("提示", "播放音频失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Modal)
		m.btnPlay.SetText("播放").Redraw(false)
		return
	}
	m.btnPause.SetText("暂停").Enable(true).Redraw(false)

	// 获取音量
	vol, err = m.ap.GetVolume()
	fmt.Printf("音量: %d, %v\n", vol, err)

	// 获取音频长度
	length, err := m.ap.GetLength()
	fmt.Printf("音频长度: %d毫秒, %v\n", length, err)

	// 获取播放进度, 等待播放完毕
	go func() {
		isStopped := false

		for !isStopped {
			wapi.Sleep(100)

			xc.UI(func() {
				isStopped = m.ap.IsStopped()
				if isStopped {
					m.btnPlay.SetText("播放").Redraw(false)
					m.btnPause.SetText("暂停").Enable(false).Redraw(false)
					m.progressBar.SetPos(0).Redraw(false)
					m.textTime.SetText("00:00 / 00:00")
					m.textTime.Redraw()
					return
				}

				// 获取播放进度
				pos, err := m.ap.GetPosition()
				if err != nil {
					fmt.Println("获取播放进度失败: " + err.Error())
					return
				}
				// 更新进度条 (百分比)
				pct := int32(pos * 100 / length)
				m.progressBar.SetPos(pct).Redraw(false)
				// 更新时间文本
				m.textTime.SetText(fmt.Sprintf("%s / %s",
					formatDuration(pos), formatDuration(length)))
				m.textTime.Redraw()
			})
		}
		fmt.Println("播放完毕!")
	}()
}

// 停止播放音频
func (m *MainWindow) stopAudio() {
	if m.ap.IsStopped() {
		m.btnPlay.SetText("播放").Redraw(false)
		return
	}

	err := m.ap.Stop()
	if err != nil {
		m.w.MessageBox("提示", "停止播放音频失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Modal)
	}
	m.btnPlay.SetText("播放").Redraw(false)
}

// 暂停播放音频
func (m *MainWindow) pauseAudio() {
	if !m.ap.IsPlaying() {
		return
	}

	err := m.ap.Pause()
	if err != nil {
		m.w.MessageBox("提示", "暂停播放音频失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Modal)
	}
	m.btnPause.SetText("继续").Redraw(false)
}

// 继续播放音频
func (m *MainWindow) resumeAudio() {
	if !m.ap.IsPaused() {
		return
	}

	err := m.ap.Resume()
	if err != nil {
		m.w.MessageBox("提示", "继续播放音频失败: "+err.Error(), xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Error, xcc.Window_Style_Modal)
	}
	m.btnPause.SetText("暂停").Redraw(false)
}

// formatDuration 将毫秒数格式化为 mm:ss
func formatDuration(ms int) string {
	totalSec := ms / 1000
	min := totalSec / 60
	sec := totalSec % 60
	return fmt.Sprintf("%02d:%02d", min, sec)
}
