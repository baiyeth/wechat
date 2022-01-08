package wechat_test

import (
	"context"
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

var (
	page = "/page/index/index"
	path = "/page/index/index"
	w    = 144
	qr   = wechat.NewQrcode(context.Background())
)

func TestCreateWXAQRCode(t *testing.T) {
	t.Parallel()
	_, err := qr.CreateWXAQRCode(page, path, w)
	fmt.Print(err)
}

func TestGetWXACode(t *testing.T) {
	t.Parallel()
	_, err := qr.GetWXACode(page, path, w)
	fmt.Print(err)
}

func TestGetWXACodeUnlimit(t *testing.T) {
	t.Parallel()
	_, err := qr.GetWXACodeUnlimit(page, path, w)
	fmt.Print(err)
}
