package wechat

import (
	"context"

	wechat2 "github.com/silenceper/wechat/v2"

	"github.com/baiyeth/wechat/internal/conf"
	"github.com/baiyeth/wechat/miniProgram"
	"github.com/baiyeth/wechat/officialAccount"
)

type wechat struct {
	ctx context.Context
	wc  *wechat2.Wechat
	cfg conf.Configuration
}

type Option func(*wechat)

func WithConfiguration(cfg conf.Configuration) Option {
	return func(m *wechat) {
		m.cfg = cfg
	}
}

func WithOfficialTplMessage(Expr, Users, TemplateId, MiniProgramAppId, MiniProgramPagePath string) Option {
	return func(m *wechat) {
		m.cfg.Official.TplMessage.Expr = Expr
		m.cfg.Official.TplMessage.Users = Users
		m.cfg.Official.TplMessage.TemplateId = TemplateId
		m.cfg.Official.TplMessage.MiniProgramAppId = MiniProgramAppId
		m.cfg.Official.TplMessage.MiniProgramPagePath = MiniProgramPagePath
	}
}

func WithAuth(AppId, AppSecret string) Option {
	return func(m *wechat) {
		m.cfg.Official = conf.OfficialConfiguration{
			AppId:     AppId,
			AppSecret: AppSecret,
		}
		m.cfg.MiniProgram = conf.MiniProgramConfiguration{
			AppId:     AppId,
			AppSecret: AppSecret,
		}
	}
}

func NewWechat(ctx context.Context, opts ...Option) wechat {
	w := wechat{
		ctx: ctx,
		wc:  wechat2.NewWechat(),
		cfg: conf.Configuration{},
	}
	for _, o := range opts {
		o(&w)
	}
	return w
}

func NewDefaultWechat(ctx context.Context, AppId, AppSecret string) wechat {
	return NewWechat(ctx, WithAuth(AppId, AppSecret))
}

func (w *wechat) GetMiniProgram() miniProgram.MiniProgram {
	return w.getMiniProgram(w.cfg.MiniProgram.AppId, w.cfg.MiniProgram.AppSecret)
}

func (w *wechat) GetOfficialAccount() officialAccount.OfficialAccount {
	return w.getOfficialAccount(w.cfg.Official.AppId, w.cfg.Official.AppSecret)
}

func (w *wechat) GetCfg() conf.Configuration {
	return w.cfg
}

func (w *wechat) getMiniProgram(AppId, AppSecret string) miniProgram.MiniProgram {
	return miniProgram.NewDefaultMiniProgram(w.ctx, w.wc, AppId, AppSecret)
}

func (w *wechat) getOfficialAccount(AppId, AppSecret string) officialAccount.OfficialAccount {
	return officialAccount.NewOfficialAccount(w.ctx, w.wc, officialAccount.WithOfficialAccountAuth(AppId, AppSecret))
}
