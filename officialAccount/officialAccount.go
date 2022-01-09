package officialAccount

import (
	"context"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"

	"github.com/baiyeth/wechat/internal/conf"
)

type OfficialAccount struct {
	oa  *officialaccount.OfficialAccount
	ctx *context.Context
}

func NewOfficialAccount(ctx context.Context, wc *wechat.Wechat) OfficialAccount {
	if ctx == nil {
		ctx = context.Background()
	}
	return OfficialAccount{
		oa:  getOfficialAccount(wc),
		ctx: &ctx,
	}
}

func (oa *OfficialAccount) GetOa() *officialaccount.OfficialAccount {
	return oa.oa
}

func (oa *OfficialAccount) GetTemplateMessage() *TemplateMessage {
	return NewTemplateMessage(oa)
}

// getOfficialAccount 获取公众号配置
func getOfficialAccount(wc *wechat.Wechat) *officialaccount.OfficialAccount {
	cacheIns := cache.NewMemory()
	cfg := &config.Config{
		AppID:          conf.Conf.Official.AppId,
		AppSecret:      conf.Conf.Official.AppSecret,
		EncodingAESKey: conf.Conf.Official.Encoding,
		Cache:          cacheIns,
	}
	return wc.GetOfficialAccount(cfg)
}
