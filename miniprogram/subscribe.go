package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

// Subscribe 微信小程序相关API
type Subscribe struct {
	MiniProgram
}

func NewSubscribe(ctx context.Context) *Subscribe {
	return &Subscribe{
		NewMiniProgram(ctx),
	}
}

func (sb *Subscribe) Send(msg *subscribe.Message) (err error) {
	return sb.GetMp().GetSubscribe().Send(msg)
}
