// 加载布局文件从string
package main

import (
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

func main() {
	a := app.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 加载资源文件从字符串
	a.LoadResourceFromStringW(resStr, "resource.res")
	// 加载窗口布局文件从字符串
	w := window.NewByLayoutStringW(layoutStr, 0, 0)
	// 加载布局文件后, 调整布局是必须的, 因为窗口布局改变了
	w.AdjustLayout()

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}

// 可以到设计器里创建一个xml, 把文本覆盖进去, 关掉xml重新打开, 可看到效果, 然后可自己修改加以使用
const layoutStr = `<?xml version="1.0" encoding="UTF-8"?>
<!--炫彩界面库-窗口布局文件-->
<head />
<windowUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(0)22(-10927566)23(255);98:1(0);}" center="true" content="炫彩界面库 - 我的窗口名" dragWindow="true" enableLayout="true" overlayBorder="true" rect="20,20,350,600" showT="true" transparentAlpha="255" transparentFlag="shadow" windowStyle="0">
	<layoutEleUI expandT="true" layout.alignH="center" layout.height="fill" layout.horizon="false" layout.space="0" layout.width="fill" padding="0,0,0,0" rect="137,368,100,100" showT="true">
		<layoutEleUI expandT="false" layout.alignH="center" layout.alignV="center" layout.height="180" layout.horizon="false" layout.space="0" layout.width="fill" padding="0,0,0,0" rect="231,116,100,100" showT="true">
			<layoutEleUI expandT="false" layout.alignH="right" layout.height="30" layout.width="fill" rect="245,21,100,100" showT="true">
				<buttonUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(0)22(-10927566)23(255);5:2(48)8(135.00)3(10,2,10,2)20(1)21(3)26(0)22(-3163991)23(255);5:2(48)8(45.00)3(10,2,10,2)20(1)21(3)26(0)22(-3163991)23(255);5:2(48)8(135.00)3(10,2,10,2)20(1)21(3)26(0)22(-1)23(255);5:2(48)8(45.00)3(10,2,10,2)20(1)21(3)26(0)22(-1)23(255);98:16(0,1,2)32(0,3,4)64(0,4,3);}" buttonType="close" expandT="true" rect="265,17,30,30" showT="true" />
			</layoutEleUI>
			<layoutEleUI expandT="false" layout.alignH="center" layout.alignV="center" layout.height=":1" layout.width="fill" rect="267,136,100,100" showT="true">
				<buttonUI bindEle="@ID_1" bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(1)22(-1)23(255)9(20,0,0,20);5:2(15)20(1)21(3)26(1)22(-13031139)23(255)9(20,0,0,20);5:2(15)20(1)21(3)26(1)22(-1644826)23(255)9(20,0,0,20);98:272(0)288(2)320(2)128(1);}" buttonType="radio" check="true" content="登陆" expandT="true" font="@ID_FONT_12" name="登陆" rect="200,132,90,40" showT="true" textColor="#FFFFFFFF" transparent="true" />
				<buttonUI bindEle="@ID_2" bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(1)22(-1)23(255)9(0,20,20,0);5:2(15)20(1)21(3)26(1)22(-13031139)23(255)9(0,20,20,0);5:2(15)20(1)21(3)26(1)22(-1644826)23(255)9(0,20,20,0);98:272(0)288(2)320(2)128(1);}" buttonType="radio" content="注册" expandT="true" font="@ID_FONT_12" name="注册" rect="201,93,90,40" showT="true" textColor="#FF000000" transparent="true" />
			</layoutEleUI>
		</layoutEleUI>
		<layoutEleUI expandT="true" layout.height=":1" layout.width="fill" name="content" rect="229,330,100,100" showT="true">
			<layoutEleUI expandT="false" id="@ID_1" layout.height="fill" layout.horizon="false" layout.space="10" layout.width="fill" padding="25,0,25,0" rect="150,282,100,100" showT="true">
				<shapeText content="用户名" expandT="true" font="@ID_FONT_2" layout.height="20" layout.width="auto" rect="167,97,100,20" showT="true" textColor="#FF9CA9BC" />
				<editUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(0)22(-10927566)23(255);5:2(13)3(0,2,0,0)20(1)21(3)26(0)22(-9215146)23(255);98:16(0,1);}" caretColor="#FFFFFFFF" contentTips="您的用户名" contentTipsColor="#FF6A768C" expandT="true" font="@ID_FONT_2" layout.width="fill" rect="211,154,100,25" showT="true" textColor="#FFFFFFFF" />
				<layoutEleUI expandT="true" layout.width="fill" rect="186,219,100,10" showT="true" />
				<shapeText content="密码" expandT="true" font="@ID_FONT_2" layout.height="20" layout.width="auto" rect="187,117,100,20" showT="true" textColor="#FF9CA9BC" />
				<editUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(0)22(-10927566)23(255);5:2(13)3(0,2,0,0)20(1)21(3)26(0)22(-9215146)23(255);98:16(0,1);}" caretColor="#FFFFFFFF" contentTips="您的密码" contentTipsColor="#FF6A768C" expandT="true" font="@ID_FONT_2" layout.width="fill" rect="231,174,100,25" showT="true" textColor="#FFFFFFFF" />
				<layoutEleUI expandT="true" layout.width="fill" rect="206,239,100,10" showT="true" />
				<shapeText content="忘记密码?" expandT="true" layout.height="20" layout.width="auto" rect="87,328,100,20" showT="true" textColor="#FF9CA9BC" />
				<layoutEleUI expandT="true" layout.width="fill" rect="226,259,100,10" showT="true" />
				<layoutEleUI expandT="true" layout.alignH="center" layout.alignV="center" layout.horizon="true" layout.width="fill" rect="246,279,100,50" showT="true">
					<buttonUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(1)22(-1)23(255)9(20,20,20,20);5:2(15)20(1)21(3)26(1)22(-1644826)23(255)9(20,20,20,20);5:2(15)20(1)21(3)26(1)22(-2302756)23(255)9(20,20,20,20);98:272(0)288(1)320(2);}" content="确定登陆" expandT="true" font="@ID_FONT_12" rect="220,152,130,40" showT="true" transparent="true" />
				</layoutEleUI>
			</layoutEleUI>
			<layoutEleUI expandT="false" id="@ID_2" layout.height="fill" layout.horizon="false" layout.space="10" layout.width="fill" padding="25,0,25,0" rect="150,282,100,100" showT="false">
				<shapeText content="用户名" expandT="true" font="@ID_FONT_2" layout.height="20" layout.width="auto" rect="167,97,100,20" showT="true" textColor="#FF9CA9BC" />
				<editUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(0)22(-10927566)23(255);5:2(13)3(0,2,0,0)20(1)21(3)26(0)22(-9215146)23(255);98:16(0,1);}" caretColor="#FFFFFFFF" contentTips="请设置您的用户名" contentTipsColor="#FF6A768C" expandT="true" font="@ID_FONT_2" layout.width="fill" rect="211,154,100,25" showT="true" textColor="#FFFFFFFF" />
				<layoutEleUI expandT="true" layout.width="fill" rect="186,219,100,10" showT="true" />
				<shapeText content="密码" expandT="true" font="@ID_FONT_2" layout.height="20" layout.width="auto" rect="187,117,100,20" showT="true" textColor="#FF9CA9BC" />
				<editUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(0)22(-10927566)23(255);5:2(13)3(0,2,0,0)20(1)21(3)26(0)22(-9215146)23(255);98:16(0,1);}" caretColor="#FFFFFFFF" contentTips="请设置您的密码" contentTipsColor="#FF6A768C" expandT="true" font="@ID_FONT_2" layout.width="fill" rect="231,174,100,25" showT="true" textColor="#FFFFFFFF" />
				<layoutEleUI expandT="true" layout.width="fill" rect="206,239,100,10" showT="true" />
				<shapeText content="E-mail" expandT="true" font="@ID_FONT_2" layout.height="20" layout.width="auto" rect="207,137,100,20" showT="true" textColor="#FF9CA9BC" />
				<editUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(0)22(-10927566)23(255);5:2(13)3(0,2,0,0)20(1)21(3)26(0)22(-9215146)23(255);98:16(0,1);}" caretColor="#FFFFFFFF" contentTips="请设置您的E-mail" contentTipsColor="#FF6A768C" expandT="true" font="@ID_FONT_2" layout.width="fill" rect="251,194,100,25" showT="true" textColor="#FFFFFFFF" />
				<layoutEleUI expandT="true" layout.alignV="center" layout.horizon="true" layout.space="5" layout.width="fill" rect="226,259,100,30" showT="true">
					<elementUI bkInfoM="{99:1.9.9;6:2(15)20(1)21(3)26(1)22(-1)23(255);5:2(3)8(30.00)3(10,5,2,12)20(1)21(3)26(1)22(-12110809)23(255);5:2(3)8(120.00)3(5,10,2,6)20(1)21(3)26(1)22(-12110809)23(255);98:16(0,1,2);}" expandT="true" rect="37,16,20,20" showT="true" transparent="true" />
					<shapeText content="我已阅读并同意账号的使用协议" expandT="true" font="@ID_FONT_2" layout.height="20" layout.width="auto" rect="117,20,100,20" showT="true" textColor="#FF9CA9BC" />
				</layoutEleUI>
				<layoutEleUI expandT="true" layout.width="fill" rect="226,259,100,10" showT="true" />
				<layoutEleUI expandT="true" layout.alignH="center" layout.alignV="center" layout.horizon="true" layout.width="fill" rect="246,279,100,50" showT="true">
					<buttonUI bkInfoM="{99:1.9.9;5:2(15)20(1)21(3)26(1)22(-1)23(255)9(20,20,20,20);5:2(15)20(1)21(3)26(1)22(-1644826)23(255)9(20,20,20,20);5:2(15)20(1)21(3)26(1)22(-2302756)23(255)9(20,20,20,20);98:272(0)288(1)320(2);}" content="确定注册" expandT="true" font="@ID_FONT_12" rect="220,152,130,40" showT="true" transparent="true" />
				</layoutEleUI>
			</layoutEleUI>
		</layoutEleUI>
	</layoutEleUI>
</windowUI>
`

const resStr = `<?xml version="1.0" encoding="UTF-8"?>
<!--炫彩界面库-资源文件-->
<resources>
	<type_image />
	<type_id>
		<item name="ID_1" value="1" />
		<item name="ID_2" value="2" />
	</type_id>
	<type_font>
		<font name="ID_FONT_1" size="16" fontName="微软雅黑" />
		<font name="ID_FONT_12" size="12" fontName="微软雅黑" />
		<font name="ID_FONT_2" size="9" fontName="微软雅黑" />
	</type_font>
	<type_string />
	<type_color />
	<type_bkInfo />
	<type_nameTable />
</resources>
`
