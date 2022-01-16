package wechat_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/baiyeth/wechat"
)

func TestNewWechat(t *testing.T) {
	ctx := context.Background()
	var (
		AppId     = "test-Appid"
		AppSecret = "test-AppSecret"
	)
	wc := wechat.NewWechat(ctx, wechat.WithAuth(AppId, AppSecret))
	mp := wc.GetMiniProgram()
	oc := wc.GetOfficialAccount()
	fmt.Printf("wechat:%v\n miniapp: %v\n officialAccount:%v", wc, mp, oc)
}
