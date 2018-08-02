package payment

import (
	"encoding/xml"
	"time"

	"github.com/medivhzhan/weapp/util"
)

const transferInfoAPI = "/mmpaymkttransfers/gettransferinfo"

// TransferInfo params to get transfer info
type TransferInfo struct {
	AppID      string `xml:"appid"`
	MchID      string `xml:"mch_id"`           // 商户号
	OutTradeNo string `xml:"partner_trade_no"` // 商户订单号
}

type transferInfo struct {
	XMLName xml.Name `xml:"xml"`
	TransferInfo
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"` // 签名
}

type transferInfoResponse struct {
	response
	OutTradeNo    string `xml:"partner_trade_no"` // 商户订单号
	MchID         string `xml:"mch_id"`
	TransactionID string `xml:"detail_id"` // TODO: 确认是这个破玩意儿
	// 转账状态
	// SUCCESS:转账成功
	// FAILED:转账失败
	// PROCESSING:处理中
	Status   string `xml:"status"` // 如果失败则有失败原因
	Reason   string `xml:"reason"`
	ToUser   string `xml:"openid"`         // 收款用户openid
	RealName string `xml:"re_user_name"`   // 收款用户姓名
	Amount   int    `xml:"payment_amount"` // 付款金额单位分）
	Desc     string `xml:"desc"`           // 付款时候的描述
	// 发起转账的时间
	// format: 2015-04-21 20:00:00
	TransferTime string `xml:"payment_time"`
}

// TransferInfoResponse 转账返回数据
type TransferInfoResponse struct {
	transferInfoResponse
	TransferTime time.Time
}

// 请求前准备
func (t *TransferInfo) prepare(key string) (transferInfo, error) {
	info := transferInfo{
		TransferInfo: *t,
		NonceStr:     util.RandomString(32),
	}

	signData := map[string]string{
		"appid":            info.AppID,
		"mch_id":           info.MchID,
		"nonce_str":        info.NonceStr,
		"partner_trade_no": info.OutTradeNo,
	}

	sign, err := util.SignByMD5(signData, key)
	if err != nil {
		return info, err
	}
	info.Sign = sign

	return info, nil
}

// GetInfo 转账信息
func (t TransferInfo) GetInfo(key string, certPath, keyPath string) (res TransferInfoResponse, err error) {
	reqData, err := t.prepare(key)
	if err != nil {
		return
	}

	resData, err := util.TSLPostXML(baseURL+transferInfoAPI, reqData, certPath, keyPath)
	if err != nil {
		return
	}
	var tres transferInfoResponse
	if err = xml.Unmarshal(resData, &tres); err != nil {
		return
	}

	if err = tres.Check(); err != nil {
		return
	}

	res.transferInfoResponse = tres
	res.TransferTime, err = time.Parse(transferTimeFormat, tres.TransferTime)

	return
}
