package miniProgram

import (
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

// Subscribe 微信小程序相关API
type Subscribe struct {
	*MiniProgram
}

func NewSubscribe(mp *MiniProgram) *Subscribe {
	return &Subscribe{mp}
}

func (sb *Subscribe) Send(msg *subscribe.Message) (err error) {
	return sb.GetMp().GetSubscribe().Send(msg)
}
