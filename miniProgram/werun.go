package miniProgram

import (
	"github.com/silenceper/wechat/v2/miniprogram/werun"
)

// Werun 微信小程序相关API
type Werun struct {
	*MiniProgram
}

func NewWerun(mp *MiniProgram) *Werun {
	return &Werun{mp}
}

func (wr *Werun) GetWeRunData(sessionKey, encryptedData, iv string) (*werun.Data, error) {
	return wr.GetMp().GetWeRun().GetWeRunData(sessionKey, encryptedData, iv)
}
