package miniProgram_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/baiyeth/wechat"
	"github.com/baiyeth/wechat/miniProgram"
)

var (
	miniapp miniProgram.MiniProgram
)

func Test(t *testing.T) {
	ctx := context.Background()
	var (
		AppId     = "test-AppId"
		AppSecret = "test-AppSecret"
	)
	wc := wechat.NewWechat(ctx, wechat.WithAuth(AppId, AppSecret))
	miniapp = wc.GetMiniProgram()
	fmt.Println(miniapp)
}
