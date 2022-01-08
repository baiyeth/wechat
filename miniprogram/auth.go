package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram/auth"
)

func Code2Session(jsCode string, appid, appSecret string) (result auth.ResCode2Session, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAuth().Code2Session(jsCode)
}
