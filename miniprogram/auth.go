package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

// Auth 微信小程序相关API
type Auth struct {
	MiniProgram
}

func NewAuth(ctx context.Context) *Auth {
	return &Auth{
		NewMiniProgram(ctx),
	}
}

func (au *Auth) Code2Session(jsCode string) (result auth.ResCode2Session, err error) {
	return au.GetMp().GetAuth().Code2Session(jsCode)
}
