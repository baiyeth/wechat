package wechat

import (
	"context"

	wechat2 "github.com/silenceper/wechat/v2"

	"github.com/baiyeth/wechat/internal/conf"
	"github.com/baiyeth/wechat/miniProgram"
	"github.com/baiyeth/wechat/officialAccount"
	"github.com/baiyeth/wechat/pay"
)

type Wechat struct {
	ctx context.Context
	wc  *wechat2.Wechat
	cfg conf.Configuration
}

type Option func(*Wechat)

func WithConfiguration(cfg conf.Configuration) Option {
	return func(m *Wechat) {
		m.cfg = cfg
	}
}

func WithOfficialTplMessage(Expr, Users, TemplateId, MiniProgramAppId, MiniProgramPagePath string) Option {
	return func(m *Wechat) {
		m.cfg.Official.TplMessage.Expr = Expr
		m.cfg.Official.TplMessage.Users = Users
		m.cfg.Official.TplMessage.TemplateId = TemplateId
		m.cfg.Official.TplMessage.MiniProgramAppId = MiniProgramAppId
		m.cfg.Official.TplMessage.MiniProgramPagePath = MiniProgramPagePath
	}
}

func WithAuth(AppId, AppSecret string) Option {
	return func(m *Wechat) {
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

func WithPayConfig(appId, appSecret, mchId, key, notifyUrl string) Option {
	return func(m *Wechat) {
		m.cfg.Pay = conf.PayConfiguration{
			AppId:     appId,
			AppSecret: appSecret,
			MchID:     mchId,
			Key:       key,
			NotifyURL: notifyUrl,
		}
		m.cfg.Official = conf.OfficialConfiguration{
			AppId:     appId,
			AppSecret: appSecret,
		}
		m.cfg.MiniProgram = conf.MiniProgramConfiguration{
			AppId:     appId,
			AppSecret: appSecret,
		}
	}
}

func NewWechat(ctx context.Context, opts ...Option) *Wechat {
	w := &Wechat{
		ctx: ctx,
		wc:  wechat2.NewWechat(),
		cfg: conf.Configuration{},
	}
	for _, o := range opts {
		o(w)
	}
	return w
}

func NewDefaultWechat(ctx context.Context, AppId, AppSecret string) *Wechat {
	return NewWechat(ctx, WithAuth(AppId, AppSecret))
}

func (w *Wechat) GetMiniProgram() miniProgram.MiniProgram {
	return w.getMiniProgram(w.cfg.MiniProgram.AppId, w.cfg.MiniProgram.AppSecret)
}

func (w *Wechat) GetOfficialAccount() officialAccount.OfficialAccount {
	return w.getOfficialAccount(w.cfg.Official.AppId, w.cfg.Official.AppSecret)
}

func (w *Wechat) GetPay() pay.Pay {
	return w.getPay(pay.Config{
		AppID:     w.cfg.Pay.AppId,
		MchID:     w.cfg.Pay.MchID,
		Key:       w.cfg.Pay.Key,
		NotifyURL: w.cfg.Pay.NotifyURL,
	})
}

func (w *Wechat) GetCfg() conf.Configuration {
	return w.cfg
}

func (w *Wechat) getMiniProgram(AppId, AppSecret string) miniProgram.MiniProgram {
	return miniProgram.NewDefaultMiniProgram(w.ctx, w.wc, AppId, AppSecret)
}

func (w *Wechat) getOfficialAccount(AppId, AppSecret string) officialAccount.OfficialAccount {
	return officialAccount.NewOfficialAccount(w.ctx, w.wc, officialAccount.WithOfficialAccountAuth(AppId, AppSecret))
}

func (w *Wechat) getPay(cfg pay.Config) pay.Pay {
	return pay.NewPay(w.ctx, w.wc, pay.WithPayConfig(cfg.AppID, cfg.MchID, cfg.Key, cfg.NotifyURL))
}
