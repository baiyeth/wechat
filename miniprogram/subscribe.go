package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

func Send(msg *subscribe.Message, appid, appSecret string) (err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetSubscribe().Send(msg)
}
