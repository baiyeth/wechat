package miniProgram_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/baiyeth/wechat/miniProgram"
)

var (
	miniapp miniProgram.MiniProgram
)

func Test(t *testing.T) {
	miniapp = miniProgram.NewMiniProgram(context.Background(), "../conf/wechat.yml")
	fmt.Println(miniapp)
}
