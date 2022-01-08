package wechat_test

import (
	"context"
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestGetWeRunData(t *testing.T) {
	t.Parallel()
	sessionKey := ""
	encryptedData := ""
	iv := ""
	wr := wechat.NewWerun(context.Background())
	data, err := wr.GetWeRunData(sessionKey, encryptedData, iv)
	fmt.Print(data, err)
}
