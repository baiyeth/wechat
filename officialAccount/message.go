package officialAccount

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type TemplateMessage struct {
	*OfficialAccount
}

func NewTemplateMessage(oa *OfficialAccount) *TemplateMessage {
	return &TemplateMessage{oa}
}

// SendTplMessage 发送微信公众号模板消息
func (tm *TemplateMessage) SendTplMessage(msg *message.TemplateMessage) error {
	_, err := tm.GetOa().GetTemplate().Send(msg)
	return err
}
