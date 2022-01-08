package wechat

import (
	"context"
	
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	
	byWechat "github.com/baiyeth/wechat"
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
		AppID:          byWechat.Conf.Official.AppId,
		AppSecret:      byWechat.Conf.Official.AppSecret,
		EncodingAESKey: byWechat.Conf.Official.Encoding,
		Cache:          cacheIns,
	}
	return wc.GetOfficialAccount(cfg)
}
