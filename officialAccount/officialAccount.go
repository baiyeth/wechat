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
	ctx context.Context
	cfg conf.OfficialConfiguration
}

type Option func(*OfficialAccount)

func WithOfficialAccountAuth(AppId, AppSecret string) Option {
	return func(o* OfficialAccount) {
		o.cfg = conf.OfficialConfiguration{
			AppId:     AppId,
			AppSecret: AppSecret,
		}
	}
}

func WithOfficialTplMessage(Expr, Users, TemplateId, MiniProgramAppId, MiniProgramPagePath string) Option {
	return func(o* OfficialAccount) {
		o.cfg.TplMessage.Expr = Expr
		o.cfg.TplMessage.Users = Users
		o.cfg.TplMessage.TemplateId = TemplateId
		o.cfg.TplMessage.MiniProgramAppId = MiniProgramAppId
		o.cfg.TplMessage.MiniProgramPagePath = MiniProgramPagePath
	}
}

func NewOfficialAccount(ctx context.Context, wc *wechat.Wechat, opts ...Option) OfficialAccount {
	if ctx == nil {
		ctx = context.Background()
	}
	officialAccount := OfficialAccount{
		ctx: ctx,
		cfg: conf.OfficialConfiguration{},
	}
	for _, o := range opts {
		o(&officialAccount)
	}
	officialAccount.oa = officialAccount.getOfficialAccount(wc)
	return officialAccount
}

func NewDefaultOfficialAccount(ctx context.Context, wc *wechat.Wechat, AppId, AppSecret string) OfficialAccount {
	return NewOfficialAccount(ctx, wc, WithOfficialAccountAuth(AppId, AppSecret))
}

func (oa *OfficialAccount) GetOa() *officialaccount.OfficialAccount {
	return oa.oa
}

func (oa *OfficialAccount) GetTemplateMessage() *TemplateMessage {
	return NewTemplateMessage(oa)
}

// getOfficialAccount 获取公众号配置
func (oa *OfficialAccount) getOfficialAccount(wc *wechat.Wechat) *officialaccount.OfficialAccount {
	cacheIns := cache.NewMemory()
	wcfg := &config.Config{
		AppID:          oa.cfg.AppId,
		AppSecret:      oa.cfg.AppSecret,
		EncodingAESKey: oa.cfg.Encoding,
		Cache:          cacheIns,
	}
	return wc.GetOfficialAccount(wcfg)
}
