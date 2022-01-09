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
	wc := wechat.NewWechat(ctx, "../internal/conf/wechat.yml")
	offAcc = wc.GetOfficialAccount(ctx)
	fmt.Println(offAcc)
}
