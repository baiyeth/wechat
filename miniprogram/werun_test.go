package wechat_test

import (
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestGetWeRunData(t *testing.T) {
	t.Parallel()
	appid := ""
	appSecret := ""
	sessionKey := ""
	encryptedData := ""
	iv := ""
	data, err := wechat.GetWeRunData(sessionKey, encryptedData, iv, appid, appSecret)
	fmt.Print(data, err)
}
