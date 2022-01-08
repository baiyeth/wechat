package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	
	"github.com/baiyeth/wechat"
)

// Qrcode 微信小程序相关API
type Qrcode struct {
	wechat.MiniProgram
}

func NewQrcode(ctx context.Context) *Qrcode {
	return &Qrcode{
		wechat.NewMiniProgram(ctx),
	}
}

func (qr *Qrcode) CreateWXAQRCode(page, path string, w int) (response []byte, err error) {
	qrCfg := qrcode.QRCoder{
		// page 必须是已经发布的小程序存在的页面,根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
		Page: page,
		// path 扫码进入的小程序页面路径
		Path: path,
		// width 图片宽度
		Width: w,
	}
	return qr.GetMp().GetQRCode().CreateWXAQRCode(qrCfg)
}

func (qr *Qrcode) GetWXACode(page, path string, w int) (response []byte, err error) {
	qrCfg := qrcode.QRCoder{
		// page 必须是已经发布的小程序存在的页面,根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
		Page: page,
		// path 扫码进入的小程序页面路径
		Path: path,
		// width 图片宽度
		Width: w,
	}
	return qr.GetMp().GetQRCode().GetWXACode(qrCfg)
}

func (qr *Qrcode) GetWXACodeUnlimit(page, path string, w int) (response []byte, err error) {
	qrCfg := qrcode.QRCoder{
		// page 必须是已经发布的小程序存在的页面,根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
		Page: page,
		// path 扫码进入的小程序页面路径
		Path: path,
		// width 图片宽度
		Width: w,
	}
	return qr.GetMp().GetQRCode().GetWXACodeUnlimit(qrCfg)
}
