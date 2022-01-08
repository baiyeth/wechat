package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram/werun"
)

func GetWeRunData(sessionKey, encryptedData, iv string, appid, appSecret string) (*werun.Data, error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetWeRun().GetWeRunData(sessionKey, encryptedData, iv)
}
