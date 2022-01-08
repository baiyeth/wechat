package wechat_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/silenceper/wechat/v2/officialaccount/message"

	"github.com/baiyeth/wechat/conf"
	wechat "github.com/baiyeth/wechat/officialAccount"
)

func TestSendTplMessage(t *testing.T) {
	t.Parallel()
	msg := message.TemplateMessage{
		ToUser:     "xxx",
		TemplateID: conf.WechatConf.Official.TplMessageCronTask.TemplateId,
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
	msg.MiniProgram.AppID = conf.WechatConf.Official.TplMessageCronTask.MiniProgramAppId
	msg.MiniProgram.PagePath = conf.WechatConf.Official.TplMessageCronTask.MiniProgramPagePath
	tm := wechat.NewTemplateMessage(context.Background())
	err := tm.SendTplMessage(&msg)
	fmt.Println(err)
}
