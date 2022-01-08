package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"

	"github.com/baiyeth/wechat/conf"
)

// GetMiniProgram 获取小程序实例
func GetMiniProgram(appid, appSecret string) *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	if appid == "" {
		appid = conf.WechatConf.MiniProgram.AppId
	}
	if appSecret == "" {
		appSecret = conf.WechatConf.MiniProgram.AppSecret
	}
	cfg := &miniConfig.Config{
		AppID:     appid,
		AppSecret: appSecret,
		Cache:     memory,
	}
	return wc.GetMiniProgram(cfg)
}
