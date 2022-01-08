package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/miniprogram/werun"
	
	"github.com/baiyeth/wechat"
)

// Werun 微信小程序相关API
type Werun struct {
	wechat.MiniProgram
}

func NewWerun(ctx context.Context) *Werun {
	return &Werun{
		wechat.NewMiniProgram(ctx),
	}
}

func (wr *Werun) GetWeRunData(sessionKey, encryptedData, iv string) (*werun.Data, error) {
	return wr.GetMp().GetWeRun().GetWeRunData(sessionKey, encryptedData, iv)
}
