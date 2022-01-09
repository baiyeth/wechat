package wechat

import (
	"context"

	wechat2 "github.com/silenceper/wechat/v2"

	"github.com/baiyeth/wechat/internal/conf"
	"github.com/baiyeth/wechat/miniProgram"
	"github.com/baiyeth/wechat/officialAccount"
)

type wechat struct {
	cfgPath string
	ctx     context.Context
	wc      *wechat2.Wechat
	cfg     conf.Configuration
}

func NewWechat(ctx context.Context, cfgPath string) *wechat {
	return &wechat{
		cfgPath: cfgPath,
		ctx:     ctx,
		wc:      wechat2.NewWechat(),
		cfg:     conf.LoadConfig(cfgPath),
	}
}

func (w *wechat) GetMiniProgram(ctx context.Context) miniProgram.MiniProgram {
	return miniProgram.NewMiniProgram(ctx, w.wc)
}

func (w *wechat) GetOfficialAccount(ctx context.Context) officialAccount.OfficialAccount {
	return officialAccount.NewOfficialAccount(ctx, w.wc)
}

func (w *wechat) GetCfg() conf.Configuration {
	return w.cfg
}
