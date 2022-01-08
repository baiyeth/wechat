package miniProgram_test

import (
	"fmt"
	"testing"

	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
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
	sb := miniapp.GetSubscribe()
	err := sb.Send(msg)
	fmt.Print(err)
}
