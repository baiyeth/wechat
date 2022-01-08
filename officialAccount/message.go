package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type TemplateMessage struct {
	OfficialAccount
}

func NewTemplateMessage(ctx context.Context) *TemplateMessage {
	return &TemplateMessage{
		NewOfficialAccount(ctx),
	}
}

// SendTplMessage 发送微信公众号模板消息
func (tm *TemplateMessage) SendTplMessage(msg *message.TemplateMessage) error {
	_, err := tm.oc.GetTemplate().Send(msg)
	return err
}
