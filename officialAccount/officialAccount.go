package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"

	"github.com/baiyeth/wechat/conf"
)

type OfficialAccount struct {
	oc  *officialaccount.OfficialAccount
	ctx *context.Context
}

func NewOfficialAccount(ctx context.Context) OfficialAccount {
	if ctx == nil {
		ctx = context.Background()
	}
	return OfficialAccount{
		oc:  getOfficialAccount(),
		ctx: &ctx,
	}
}

// getOfficialAccount 获取公众号配置
func getOfficialAccount() *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	cacheIns := cache.NewMemory()
	cfg := &config.Config{
		AppID:          conf.WechatConf.Official.AppId,
		AppSecret:      conf.WechatConf.Official.AppSecret,
		EncodingAESKey: conf.WechatConf.Official.Encoding,
		Cache:          cacheIns,
	}
	return wc.GetOfficialAccount(cfg)
}
