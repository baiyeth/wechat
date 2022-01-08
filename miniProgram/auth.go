package miniProgram

import (
	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

// Auth 微信小程序相关API
type Auth struct {
	*MiniProgram
}

func NewAuth(mp *MiniProgram) *Auth {
	return &Auth{mp}
}

func (au *Auth) Code2Session(jsCode string) (result auth.ResCode2Session, err error) {
	return au.GetMp().GetAuth().Code2Session(jsCode)
}
