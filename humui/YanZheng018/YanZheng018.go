// 未闻花名 - 验证界面.
// https://mall.xcgui.com/1698.html
package main

import (
	_ "embed"
	"fmt"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

var (
	//go:embed res/data.zip
	resData []byte // 打包的炫彩资源文件

	m_code    bool                // false: 登录模式，true: 卡密模式
	m_layouts []*widget.LayoutEle // 布局面板，0.卡密，1.登录，2.注册，3.改密，4.充值
	w         *window.Window
)

func main() {
	// 初始化UI库
	app.InitOrExit()
	a := app.New(true)
	// 启用自适应DPI
	a.EnableAutoDPI(true).EnableDPI(true)
	// 加载资源文件
	a.LoadResourceZipMem(resData, "资源文件\\resource.res", "")

	// 创建窗口
	w = window.NewByLayoutZipMem(resData, "布局文件\\main.xml", "", 0, 0)
	// 禁止拖动边框
	w.EnableDragBorder(false)

	// 原窗口是没有阴影的, 由于要加阴影, 就给窗口宽高加上阴影大小*2
	shawdowSize := int32(14) // 阴影大小
	w.SetSize(500+shawdowSize*2, 670+shawdowSize*2)
	// 设置窗口阴影
	w.SetTransparentType(xcc.Window_Transparent_Shadow)
	w.SetShadowInfo(shawdowSize, 100, 25, false, xcc.COLOR_BLACK)

	// 加载布局面板
	loadLayout(true)

	// 调整布局
	w.AdjustLayout()
	// 显示窗口
	w.Show(true)
	a.Run()
	a.Exit()
}

const (
	layout_code            = iota // 0.卡密
	layout_login                  // 1.登录
	layout_register               // 2.注册
	layout_change_password        // 3.改密
	layout_densification          // 4.充值
)

// 加载布局面板
//   - isCode: 是否为卡密模式
func loadLayout(isCode bool) {
	m_code = isCode
	layoutParent := widget.NewLayoutEleByName("warp") // 获取父级布局元素句柄
	// 加载布局面板到布局元素
	// 0.卡密，1.登录，2.注册，3.改密，4.充值
	layoutNames := []string{"layout_code.xml", "layout_login.xml", "layout_register.xml", "layout_change_password.xml", "layout_densification.xml"}
	for i := 0; i < len(layoutNames); i++ {
		m_layouts = append(m_layouts, widget.NewLayoutEleByLayoutZipMem(resData, "布局文件\\"+layoutNames[i], "", layoutParent.Handle, 0))
		if i == layout_code {
			m_layouts[i].Show(isCode)
		} else if i == layout_login {
			m_layouts[i].Show(!isCode)
		} else {
			m_layouts[i].Show(false)
		}
	}
	regEvent()
	layoutParent.AdjustLayout(0)
	w.Redraw(false)
}

// 注册事件
func regEvent() {
	// 注册返回按钮事件
	widget.NewButtonByName("back").AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if m_code {
			changeLayout(layout_code)
		} else {
			changeLayout(layout_login)
		}
		return 0
	})

	codeEvent()
	loginEvent()
	registerEvent()
	densificationEvent()
	changePasswordEvent()
}

// 切换布局面板
//   - 0.卡密，1.登录，2.注册，3.改密，4.充值
func changeLayout(index int) {
	for i := 0; i < len(m_layouts); i++ {
		if i == index && m_layouts[i].IsShow() {
			return
		}
		m_layouts[i].Show(i == index)
	}
	widget.NewButtonByName("back").Show(index > 1) // 只在0卡密和1登录界面显示返回按钮
	w.AdjustLayout().Redraw(false)
}

// 卡密界面事件
func codeEvent() {
	code_edit_key := widget.NewEditByName("code_edit_key")     // 卡密编辑框
	code_check_rmr := widget.NewButtonByName("code_check_rmr") // 记住卡密
	code_buy := widget.NewButtonByName("code_buy")             // 购买卡密
	code_login := widget.NewButtonByName("code_login")         // 登录卡密
	code_go_login := widget.NewButtonByName("code_go_login")   // 使用账户模式

	// 卡密编辑框, 按键弹起事件
	code_edit_key.AddEvent_KeyUp(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
		if wParam == xcc.VK_Enter { // 登录按钮被点击
			code_login.SendEvent(xcc.XE_BNCLICK, 0, 0)
		}
		return 0
	})
	// 记住卡密, 按钮被选择事件
	code_check_rmr.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		fmt.Printf("记住卡密状态: %v\n", bCheck)
		return 0
	})
	// 购买卡密, 按钮单击事件
	code_buy.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		changeLayout(layout_densification)
		return 0
	})
	// 登录卡密, 按钮单击事件
	code_login.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if Msg(w.Handle, "登录卡密, 按钮单击事件\n这个例子里是没有具体功能的") == xcc.MessageBox_Flag_Ok {
			fmt.Println("提示框确认按钮被点击")
		}
		return 0
	})
	// 使用账户模式, 按钮单击事件
	code_go_login.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		changeLayout(layout_login)
		return 0
	})
}

// 登录界面事件
func loginEvent() {
	// login_edit_username := widget.NewEditByName("login_edit_username") // 用户名编辑框
	login_edit_password := widget.NewEditByName("login_edit_password")   // 密码编辑框
	login_check_rmr := widget.NewButtonByName("login_check_rmr")         // 记住登录状态
	login_check_visible := widget.NewButtonByName("login_check_visible") // 密码显示隐藏
	create_account := widget.NewButtonByName("create_account")           // 创建账户
	login_forget := widget.NewButtonByName("login_forget")               // 忘记密码
	login_login := widget.NewButtonByName("login_login")                 // 登录按钮
	login_go_code := widget.NewButtonByName("login_go_code")             // 使用卡密模式

	// 密码编辑框, 设置密码字符, 按键弹起事件
	login_edit_password.EnablePassword(true).SetPasswordCharacter('●')
	login_edit_password.AddEvent_KeyUp(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
		if wParam == xcc.VK_Enter { // 登录按钮被点击
			login_login.SendEvent(xcc.XE_BNCLICK, 0, 0)
		}
		return 0
	})
	// 记住登录状态, 按钮被选择事件
	login_check_rmr.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		fmt.Printf("记住登录状态: %v\n", bCheck)
		return 0
	})
	// 密码显示隐藏, 按钮被选择事件
	login_check_visible.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		fmt.Printf("密码显示隐藏: %v\n", bCheck)
		login_edit_password.EnablePassword(!bCheck).Redraw(false)
		return 0
	})
	// 创建账户, 按钮单击事件
	create_account.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		changeLayout(layout_register)
		return 0
	})
	// 忘记密码, 按钮单击事件
	login_forget.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		changeLayout(layout_change_password)
		return 0
	})
	// 登录按钮, 按钮单击事件
	login_login.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if Msg(w.Handle, "登录按钮, 按钮单击事件\n这个例子里是没有具体功能的") == xcc.MessageBox_Flag_Ok {
			fmt.Println("提示框确认按钮被点击")
		}
		return 0
	})
	// 使用卡密模式, 按钮单击事件
	login_go_code.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		changeLayout(layout_code)
		return 0
	})
}

// 注册界面事件
func registerEvent() {
	// reg_edit_username := widget.NewEditByName("reg_edit_username") // 注册用户名编辑框
	reg_edit_password := widget.NewEditByName("reg_edit_password")   // 注册密码编辑框
	reg_edit_ID := widget.NewEditByName("reg_edit_ID")               // 推荐人ID编辑框
	reg_check_visible := widget.NewButtonByName("reg_check_visible") // 注册密码显示隐藏
	reg_go_login := widget.NewButtonByName("reg_go_login")           // 去登录
	reg_reg := widget.NewButtonByName("reg_reg")                     // 注册按钮

	// 注册密码编辑框, 设置密码字符
	reg_edit_password.EnablePassword(true).SetPasswordCharacter('●')
	// 注册密码编辑框, 推荐人ID编辑框, 按键弹起事件
	cbKeyUp := func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
		if wParam == xcc.VK_Enter { // 注册按钮被点击
			reg_reg.SendEvent(xcc.XE_BNCLICK, 0, 0)
		}
		return 0
	}
	reg_edit_password.AddEvent_KeyUp(cbKeyUp)
	reg_edit_ID.AddEvent_KeyUp(cbKeyUp)
	// 注册密码显示隐藏, 按钮被选择事件
	reg_check_visible.AddEvent_Button_Check(func(hEle int, bCheck bool, pbHandled *bool) int {
		fmt.Printf("注册密码显示隐藏: %v\n", bCheck)
		reg_edit_password.EnablePassword(!bCheck).Redraw(false)
		return 0
	})
	// 去登录, 按钮单击事件
	reg_go_login.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		changeLayout(layout_login)
		return 0
	})
	// 注册按钮, 按钮单击事件
	reg_reg.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if Msg(w.Handle, "注册按钮, 按钮单击事件\n这个例子里是没有具体功能的") == xcc.MessageBox_Flag_Ok {
			fmt.Println("提示框确认按钮被点击")
		}
		return 0
	})
}

// 充值卡密界面事件
func densificationEvent() {
	// den_edit_username := widget.NewEditByName("den_edit_username") // 充值用户名编辑框
	den_edit_password := widget.NewEditByName("den_edit_password") // 充值卡密编辑框
	den_service := widget.NewButtonByName("den_service")           // 充值联系客服
	den_den := widget.NewButtonByName("den_den")                   // 充值按钮

	// 充值卡密编辑框, 按键弹起事件
	den_edit_password.AddEvent_KeyUp(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
		if wParam == xcc.VK_Enter { // 充值按钮被点击
			den_den.SendEvent(xcc.XE_BNCLICK, 0, 0)
		}
		return 0
	})
	// 充值联系客服, 按钮单击事件
	den_service.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		fmt.Println("充值联系客服按钮被点击")
		return 0
	})
	// 充值按钮, 按钮单击事件
	den_den.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if Msg(w.Handle, "充值按钮, 按钮单击事件\n这个例子里是没有具体功能的") == xcc.MessageBox_Flag_Ok {
			fmt.Println("提示框确认按钮被点击")
		}
		return 0
	})
}

// 改密界面事件
func changePasswordEvent() {
	forget_edit_username := widget.NewEditByName("forget_edit_username") // 改密用户名编辑框
	forget_service := widget.NewButtonByName("forget_service")           // 改密联系客服
	forget_forget := widget.NewButtonByName("forget_forget")             // 改密按钮

	// 改密用户名编辑框, 按键弹起事件
	forget_edit_username.AddEvent_KeyUp(func(hEle int, wParam, lParam uintptr, pbHandled *bool) int {
		if wParam == xcc.VK_Enter { // 改密按钮被点击
			forget_forget.SendEvent(xcc.XE_BNCLICK, 0, 0)
		}
		return 0
	})
	// 改密联系客服, 按钮单击事件
	forget_service.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		fmt.Println("改密联系客服按钮被点击")
		return 0
	})
	// 改密按钮, 按钮单击事件
	forget_forget.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		if Msg(w.Handle, "改密按钮, 按钮单击事件\n这个例子里是没有具体功能的") == xcc.MessageBox_Flag_Ok {
			fmt.Println("提示框确认按钮被点击")
		}
		return 0
	})
}

// 提示框选项
type MsgOptins struct {
	Title      string // 标题, 为空则是"系统提示"
	ButtonType int    // 按钮类型, 0: 确认+取消, 1: 确认
}

// 提示框
//   - hWindow: 父窗口句柄
//   - text: 内容
//   - opt: 可选选项
func Msg(hWindow int, text string, opt ...MsgOptins) xcc.MessageBox_Flag_ {
	o := MsgOptins{}
	if len(opt) > 0 {
		o = opt[0]
	}
	// 创建提示框窗口
	mw := window.NewModalWindowByLayoutZipMem(resData, "布局文件\\msg.xml", "", hWindow, 0)
	mw.EnableDragBorder(false) // 禁止拖动边框
	if o.Title != "" {
		widget.NewButtonByName("msg_title").SetText(o.Title)
	}
	// 确认按钮, 按钮单击事件
	widget.NewButtonByName("msg_btn1").AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
		mw.EndModal(xcc.MessageBox_Flag_Ok)
		return 0
	})
	// 取消按钮, 按钮单击事件
	msg_btn2 := widget.NewButtonByName("msg_btn2")
	if o.ButtonType == int(xcc.MessageBox_Flag_Ok) { // 只显示确认按钮
		msg_btn2.Show(false)
	} else {
		msg_btn2.AddEvent_BnClick(func(hEle int, pbHandled *bool) int {
			mw.EndModal(xcc.MessageBox_Flag_Cancel)
			return 0
		})
	}
	// 内容
	widget.NewShapeTextByName("msg_content").SetText(text)
	// 从xml加载的显示前要调整布局
	mw.AdjustLayout()
	return mw.DoModal()
}
