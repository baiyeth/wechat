package wechat_test

import (
	"context"
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

var (
	content = wechat.NewContent(context.Background())
)

func TestCheckText(t *testing.T) {
	t.Parallel()
	text := ""
	err := content.CheckText(text)
	fmt.Print(err)
}

func TestCheckImage(t *testing.T) {
	t.Parallel()
	media := ""
	err := content.CheckImage(media)
	fmt.Print(err)
}
