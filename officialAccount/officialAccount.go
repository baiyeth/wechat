package officialAccount

import (
	"context"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"

	wechat2 "github.com/baiyeth/wechat"
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

func (oc *OfficialAccount) GetOfficialAccount() *officialaccount.OfficialAccount {
	return oc.oc
}

// getOfficialAccount 获取公众号配置
func getOfficialAccount() *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	cacheIns := cache.NewMemory()
	cfg := &config.Config{
		AppID:          wechat2.Conf.Official.AppId,
		AppSecret:      wechat2.Conf.Official.AppSecret,
		EncodingAESKey: wechat2.Conf.Official.Encoding,
		Cache:          cacheIns,
	}
	return wc.GetOfficialAccount(cfg)
}
