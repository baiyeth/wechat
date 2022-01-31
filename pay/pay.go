package pay

import (
	"context"
	"sync"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/pay"
	payConfig "github.com/silenceper/wechat/v2/pay/config"
	"github.com/silenceper/wechat/v2/pay/notify"
	"github.com/silenceper/wechat/v2/pay/order"
	"github.com/silenceper/wechat/v2/pay/refund"
	"github.com/silenceper/wechat/v2/pay/transfer"

	"github.com/baiyeth/wechat/internal/conf"
)

var (
	orderMap    sync.Map // 订单
	refundMap   sync.Map // 退款
	transferMap sync.Map // 提现
	notifyMap   sync.Map // 通知
)

type Config payConfig.Config

// OrderParams 创建订单传入的参数，用于生成 prepay_id 的必需参数
type OrderParams order.Params

type PaidResult notify.PaidResult

// QueryParams 查询订单传入的参数
type QueryParams order.QueryParams

// RefundParams 退款订单传入的参数
type RefundParams refund.Params

// RefundedResult 退款回调
type RefundedResult notify.RefundedResult

// RefundedReqInfo 退款结果（明文）
type RefundedReqInfo notify.RefundedReqInfo

// TransferParams 提现传入的参数
type TransferParams transfer.Params

// Pay 微信小程序支付
type Pay struct {
	pay *pay.Pay
	ctx context.Context
	cfg conf.PayConfiguration
}

type Option func(*Pay)

func WithPayConfig(appId, mchId, key, notifyUrl string) Option {
	return func(m *Pay) {
		m.cfg = conf.PayConfiguration{
			AppId:     appId,
			MchID:     mchId,
			Key:       key,
			NotifyURL: notifyUrl,
		}
	}
}

func WithMchId(mchId, key string) Option {
	return func(m *Pay) {
		m.cfg.MchID = mchId
		m.cfg.Key = key
	}
}

func WithNotifyUrl(notifyUrl string) Option {
	return func(m *Pay) {
		m.cfg.NotifyURL = notifyUrl
	}
}

func NewPay(ctx context.Context, wc *wechat.Wechat, opts ...Option) Pay {
	if ctx == nil {
		ctx = context.Background()
	}
	p := Pay{
		ctx: ctx,
		cfg: conf.PayConfiguration{},
	}
	for _, o := range opts {
		o(&p)
	}
	p.pay = p.getNewPay(wc)
	return p
}

// GetPay ...
func (p *Pay) GetPay() *pay.Pay {
	return p.pay
}

// CreateOrder  创建订单
func (p *Pay) CreateOrder(param *OrderParams) (string, error) {
	return p.getOrder().PrePayID((*order.Params)(param))
}

// QueryOrder  查询订单状态
func (p *Pay) QueryOrder(param *QueryParams) (notify.PaidResult, error) {
	return p.getOrder().QueryOrder((*order.QueryParams)(param))
}

// Refund 发起退款
func (p *Pay) Refund(param *RefundParams) (refund.Response, error) {
	return p.getRefund().Refund((*refund.Params)(param))
}

// Transfer 提现
func (p *Pay) Transfer(param *TransferParams) (transfer.Response, error) {
	resp, err := p.getTransfer().WalletTransfer((*transfer.Params)(param))
	return *resp, err
}

// OrderNotify 订单通知校验
func (p *Pay) OrderNotify(res *PaidResult) bool {
	status := p.getNotify().PaidVerifySign((notify.PaidResult)(*res))
	return status
}

// RefundNotify 订单退款通知校验
func (p *Pay) RefundNotify(res *RefundedResult) (*RefundedReqInfo, error) {
	reqInfo, err := p.getNotify().DecryptReqInfo((*notify.RefundedResult)(res))
	return (*RefundedReqInfo)(reqInfo), err
}

// getNewPay 获取支付实例
func (p *Pay) getNewPay(wc *wechat.Wechat) *pay.Pay {
	if p.cfg.AppId == "" || p.cfg.MchID == "" {
		panic("invalid AppId or AppSecret")
	}
	wcfg := &payConfig.Config{
		AppID:     p.cfg.AppId,
		MchID:     p.cfg.MchID,
		Key:       p.cfg.Key,
		NotifyURL: p.cfg.NotifyURL,
	}
	return wc.GetPay(wcfg)
}

// getOrder  下单
func (p *Pay) getOrder() *order.Order {
	orderIns, ok := orderMap.Load(p.cfg)
	if !ok {
		orderIns = p.pay.GetOrder()
		orderMap.Store(p.cfg, orderIns)
	}
	return orderIns.(*order.Order)
}

// getRefund  退款
func (p *Pay) getRefund() *refund.Refund {
	refundIns, ok := refundMap.Load(p.cfg)
	if !ok {
		refundIns = p.pay.GetRefund()
		refundMap.Store(p.cfg, refundIns)
	}
	return refundIns.(*refund.Refund)
}

// getTransfer 付款（提现）
func (p *Pay) getTransfer() *transfer.Transfer {
	transferIns, ok := refundMap.Load(p.cfg)
	if !ok {
		transferIns = p.pay.GetTransfer()
		refundMap.Store(p.cfg, transferIns)
	}
	return transferIns.(*transfer.Transfer)
}

// getNotify  通知
func (p *Pay) getNotify() *notify.Notify {
	notifyIns, ok := refundMap.Load(p.cfg)
	if !ok {
		notifyIns = p.pay.GetNotify()
		refundMap.Store(p.cfg, notifyIns)
	}
	return notifyIns.(*notify.Notify)
}
