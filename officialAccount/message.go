package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/officialaccount/message"
	
	"github.com/baiyeth/wechat"
)

type TemplateMessage struct {
	wechat.OfficialAccount
}

func NewTemplateMessage(ctx context.Context) *TemplateMessage {
	return &TemplateMessage{
		wechat.NewOfficialAccount(ctx),
	}
}

// SendTplMessage 发送微信公众号模板消息
func (tm *TemplateMessage) SendTplMessage(msg *message.TemplateMessage) error {
	_, err := tm.GetOfficialAccount().GetTemplate().Send(msg)
	return err
}
