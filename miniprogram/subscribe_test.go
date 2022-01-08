package wechat_test

import (
	"context"
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
	sb := wechat.NewSubscribe(context.Background())
	err := sb.Send(msg)
	fmt.Print(err)
}
