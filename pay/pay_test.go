package pay_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/baiyeth/wechat"
	"github.com/baiyeth/wechat/pay"
)

func TestOrder(t *testing.T) {
	ctx := context.Background()
	var (
		AppId     = "test-AppId"
		AppSecret = "test-AppSecret"
		MchId     = "test-MchId"
		Key       = "test-Key"
		NotifyUrl = "test-notifyUrl"
	)
	wc := wechat.NewWechat(ctx, wechat.WithPayConfig(AppId, AppSecret, MchId, Key, NotifyUrl))
	payer := wc.GetPay()
	orderId, err := payer.CreateOrder(&pay.OrderParams{
		TotalFee:   "88",
		CreateIP:   "127.0.0.1",
		Body:       "百业特惠-优惠券-1张",
		OutTradeNo: "202201311502123",
		TimeExpire: "202201311512123", // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。
		OpenID:     "test-openid",
		TradeType:  "JSAPI",
		SignType:   "MD5",
		Detail: `[{
        "goods_detail": [
            {
                "goods_id": "iphone6s_16G",
                "wxpay_goods_id": "1001",
                "goods_name": "iPhone6s 16G",
                "quantity": 1,
                "price": 528800,
                "goods_category": "123456",
                "body": "百业特惠优惠券"
            },
            {
                "goods_id": "iphone6s_32G",
                "wxpay_goods_id": "1002",
                "goods_name": "iPhone6s 32G",
                "quantity": 1,
                "price": 608800,
                "goods_category": "123789",
                "body": "百业特惠优惠券"
            }
        ]
    }
]`,
		Attach:    "百业特惠",
		GoodsTag:  "baiyetag",
		NotifyURL: "https://baiyeth-sandbox.young-ging.cn",
	})
	fmt.Println(orderId, err)
}
