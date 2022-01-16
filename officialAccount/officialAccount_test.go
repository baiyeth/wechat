package officialAccount_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/baiyeth/wechat"
	"github.com/baiyeth/wechat/officialAccount"
)

var (
	offAcc officialAccount.OfficialAccount
)

func Test(t *testing.T) {
	ctx := context.Background()
	var (
		AppId     = "test-AppId"
		AppSecret = "test-AppSecret"
	)
	wc := wechat.NewDefaultWechat(ctx, AppId, AppSecret)
	offAcc = wc.GetOfficialAccount()
	fmt.Println(offAcc)
}
