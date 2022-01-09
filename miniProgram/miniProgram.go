package miniProgram

import (
	"context"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	"github.com/baiyeth/wechat/internal/conf"
)

// MiniProgram 微信小程序相关API
type MiniProgram struct {
	mp  *miniprogram.MiniProgram
	ctx context.Context
}

func NewMiniProgram(ctx context.Context, wc *wechat.Wechat) MiniProgram {
	if ctx == nil {
		ctx = context.Background()
	}
	return MiniProgram{
		mp:  getMiniProgram(wc),
		ctx: ctx,
	}
}

func (mp *MiniProgram) GetMp() *miniprogram.MiniProgram {
	return mp.mp
}

func (mp *MiniProgram) GetAnalysis() *Analysis {
	return NewAnalysis(mp)
}

func (mp *MiniProgram) GetAuth() *Auth {
	return NewAuth(mp)
}

func (mp *MiniProgram) GetContent() *Content {
	return NewContent(mp)
}

func (mp *MiniProgram) GetQrcode() *Qrcode {
	return NewQrcode(mp)
}

func (mp *MiniProgram) GetSubscribe() *Subscribe {
	return NewSubscribe(mp)
}

func (mp *MiniProgram) GetWerun() *Werun {
	return NewWerun(mp)
}

// getMiniProgram 获取小程序实例
func getMiniProgram(wc *wechat.Wechat) *miniprogram.MiniProgram {
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     conf.Conf.MiniProgram.AppId,
		AppSecret: conf.Conf.MiniProgram.AppSecret,
		Cache:     memory,
	}
	return wc.GetMiniProgram(cfg)
}
