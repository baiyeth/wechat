package miniProgram

import (
	"context"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	wechat2 "github.com/baiyeth/wechat"
)

// MiniProgram 微信小程序相关API
type MiniProgram struct {
	mp  *miniprogram.MiniProgram
	ctx context.Context
	cfg wechat2.Configuration
}

func NewMiniProgram(ctx context.Context, cfgFile string) MiniProgram {
	if ctx == nil {
		ctx = context.Background()
	}
	if cfgFile == "" {
		panic("invalid cfg file path")
	}
	conf := wechat2.InitConfig(cfgFile)
	return MiniProgram{
		mp:  getMiniProgram(cfgFile),
		ctx: ctx,
		cfg: conf,
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
func getMiniProgram(cfgFile string) *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     wechat2.Conf.MiniProgram.AppId,
		AppSecret: wechat2.Conf.MiniProgram.AppSecret,
		Cache:     memory,
	}
	return wc.GetMiniProgram(cfg)
}
