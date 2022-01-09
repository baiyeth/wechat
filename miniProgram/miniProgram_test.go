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
	wc := wechat.NewWechat(ctx, "../internal/conf/wechat.yml")
	miniapp = wc.GetMiniProgram(ctx)
	fmt.Println(miniapp)
}
