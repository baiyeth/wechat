package wechat

func CheckText(text string, appid, appSecret string) (err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetContentSecurity().CheckText(text)
}

func CheckImage(media string, appid, appSecret string) (err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	context := miniprogram.GetContentSecurity()
	return context.CheckImage(media)
}
