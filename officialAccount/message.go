package wechat

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

// SendTplMessage 发送微信公众号模板消息
func SendTplMessage(msg *message.TemplateMessage) error {
	o := GetOfficialAccount()
	_, err := o.GetTemplate().Send(msg)
	return err
}
