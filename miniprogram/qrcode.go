package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
)

func CreateWXAQRCode(appid, appSecret string) (response []byte, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	var qrcfg qrcode.QRCoder
	return miniprogram.GetQRCode().CreateWXAQRCode(qrcfg)
}

func GetWXACode(appid, appSecret string) (response []byte, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	var qrcfg qrcode.QRCoder
	return miniprogram.GetQRCode().GetWXACode(qrcfg)
}

func GetWXACodeUnlimit(appid, appSecret string) (response []byte, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	var qrcfg qrcode.QRCoder
	return miniprogram.GetQRCode().GetWXACodeUnlimit(qrcfg)
}
