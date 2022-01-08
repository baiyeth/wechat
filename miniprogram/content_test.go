package wechat_test

import (
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestCheckText(t *testing.T) {
	t.Parallel()
	appid := ""
	appSecret := ""
	text := ""
	err := wechat.CheckText(text, appid, appSecret)
	fmt.Print(err)
}

func TestCheckImage(t *testing.T) {
	t.Parallel()
	appid := ""
	appSecret := ""
	text := ""
	err := wechat.CheckImage(text, appid, appSecret)
	fmt.Print(err)
}
