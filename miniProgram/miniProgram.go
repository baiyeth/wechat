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
	cfg conf.MiniProgramConfiguration
}

type Option func(*MiniProgram)

func WithMiniProgramAuth(AppId, AppSecret string) Option {
	return func(m *MiniProgram) {
		m.cfg = conf.MiniProgramConfiguration{
			AppId:     AppId,
			AppSecret: AppSecret,
		}
	}
}

func NewMiniProgram(ctx context.Context, wc *wechat.Wechat, opts ...Option) MiniProgram {
	if ctx == nil {
		ctx = context.Background()
	}
	miniApp := MiniProgram{
		ctx: ctx,
		cfg: conf.MiniProgramConfiguration{},
	}
	for _, o := range opts {
		o(&miniApp)
	}
	miniApp.mp = miniApp.getMiniProgram(wc)
	return miniApp
}

func NewDefaultMiniProgram(ctx context.Context, wc *wechat.Wechat, AppId, AppSecret string) MiniProgram {
	return NewMiniProgram(ctx, wc, WithMiniProgramAuth(AppId, AppSecret))
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
func (mp *MiniProgram) getMiniProgram(wc *wechat.Wechat) *miniprogram.MiniProgram {
	if mp.cfg.AppId == "" || mp.cfg.AppSecret == "" {
		panic("invalid AppId or AppSecret")
	}
	memory := cache.NewMemory()
	wcfg := &miniConfig.Config{
		AppID:     mp.cfg.AppId,
		AppSecret: mp.cfg.AppSecret,
		Cache:     memory,
	}
	return wc.GetMiniProgram(wcfg)
}
