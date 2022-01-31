### 微信相关模块api

#### wechat实例

```go
func Test() {
    ctx := context.Background()
    var (
        AppId     = "test-Appid"
        AppSecret = "test-AppSecret"
    )
    wc := wechat.NewWechat(ctx, wechat.WithAuth(AppId, AppSecret))
    fmt.Println(wc)
}
```



#### 小程序

```go
func TestMiniProgram() {
	miniapp = wc.GetMiniProgram()
    fmt.Println(miniapp)
    
    //auth
    au := miniapp.GetAuth()
	res, err := au.Code2Session("test-code")
    fmt.Println(au, res, err)
    
    //其他还有
    //GetAnalysis
    //GetContent
    //GetQrcode
    //GetSubscribe
    //GetWerun
}
```



#### 支付

```go
func TestPay {
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
            }
        ]
    }
]`,
		Attach:    "百业特惠",
		GoodsTag:  "baiyetag",
		NotifyURL: "https://baiyeth-sandbox.young-ging.cn",
	})
    fmt.Println(orderId, err)
    // wait notify callback
    status := payer.OrderNotify(res)
    if status {
        // 发货...
    }
}
```



#### 公众号

- TODO