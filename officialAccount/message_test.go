package officialAccount_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func TestSendTplMessage(t *testing.T) {
	t.Parallel()
	msg := message.TemplateMessage{
		ToUser:     "xxx",
		TemplateID: "tmp-1",
		Data: map[string]*message.TemplateDataItem{
			"first": {
				Value: "日常事项定时提醒",
			},
			"keyword1": {
				Value: "每日购买",
			},
			"keyword2": {
				Value: "请到商城下单支付一单(杨博士店有一分钱的单)",
			},
			"keyword3": {
				Value: time.Now().String(),
			},
			"remark": {
				Value: "下单完成记得将截图发到群里哦~",
			},
		},
	}
	msg.MiniProgram.AppID = "test-appid"
	msg.MiniProgram.PagePath = "test-pagePath"
	tm := offAcc.GetTemplateMessage()
	err := tm.SendTplMessage(&msg)
	fmt.Println(err)
}
