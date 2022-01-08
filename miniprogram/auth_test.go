package wechat_test

import (
	"context"
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestCode2Session(t *testing.T) {
	t.Parallel()
	au := wechat.NewAuth(context.Background())
	res, err := au.Code2Session("")
	fmt.Print(res, err)
}
