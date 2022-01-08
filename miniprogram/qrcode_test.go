package wechat_test

import (
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestCreateWXAQRCode(t *testing.T) {
	t.Parallel()
	appid := ""
	appSecret := ""
	res, err := wechat.CreateWXAQRCode(appid, appSecret)
	fmt.Print(res, err)
}

func TestGetWXACode(t *testing.T) {
	t.Parallel()
	appid := ""
	appSecret := ""
	res, err := wechat.GetWXACode(appid, appSecret)
	fmt.Print(res, err)
}

func TestGetWXACodeUnlimit(t *testing.T) {
	t.Parallel()
	appid := ""
	appSecret := ""
	res, err := wechat.GetWXACodeUnlimit(appid, appSecret)
	fmt.Print(res, err)
}
