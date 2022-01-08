package wechat_test

import (
	"fmt"
	"testing"

	"github.com/silenceper/wechat/v2/miniprogram/subscribe"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestSend(t *testing.T) {
	t.Parallel()
	msg := &subscribe.Message{
		ToUser:           "",
		TemplateID:       "",
		Page:             "",
		Data:             map[string]*subscribe.DataItem{},
		MiniprogramState: "",
		Lang:             "",
	}
	appid := ""
	appSecret := ""
	err := wechat.Send(msg, appid, appSecret)
	fmt.Print(err)
}
