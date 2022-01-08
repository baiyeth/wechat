package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	"github.com/baiyeth/wechat/conf"
)

// MiniProgram 微信小程序相关API
type MiniProgram struct {
	mp  *miniprogram.MiniProgram
	ctx *context.Context
}

func NewMiniProgram(ctx context.Context) MiniProgram {
	if ctx == nil {
		ctx = context.Background()
	}
	return MiniProgram{
		mp:  getMiniProgram(),
		ctx: &ctx,
	}
}

func (mp *MiniProgram) GetMp() *miniprogram.MiniProgram {
	return mp.mp
}

// getMiniProgram 获取小程序实例
func getMiniProgram() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     conf.WechatConf.MiniProgram.AppId,
		AppSecret: conf.WechatConf.MiniProgram.AppSecret,
		Cache:     memory,
	}
	return wc.GetMiniProgram(cfg)
}
