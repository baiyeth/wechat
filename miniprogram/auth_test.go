package wechat_test

import (
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestCode2Session(t *testing.T) {
	t.Parallel()
	appid := ""
	appSecret := ""
	res, err := wechat.Code2Session("", appid, appSecret)
	fmt.Print(res, err)
}
