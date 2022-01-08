package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	
	"github.com/baiyeth/wechat"
)

// Subscribe 微信小程序相关API
type Subscribe struct {
	wechat.MiniProgram
}

func NewSubscribe(ctx context.Context) *Subscribe {
	return &Subscribe{
		wechat.NewMiniProgram(ctx),
	}
}

func (sb *Subscribe) Send(msg *subscribe.Message) (err error) {
	return sb.GetMp().GetSubscribe().Send(msg)
}
