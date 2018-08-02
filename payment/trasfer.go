package payment

import (
	"encoding/xml"
	"errors"
	"strconv"
	"time"

	"github.com/medivhzhan/weapp/util"
)

const (
	transferTimeFormat = "2006-01-02 15:04:05"
	transferAPI        = "/mmpaymkttransfers/promotion/transfers"
)

// Transferer transfer params
type Transferer struct {
	// required
	AppID      string `xml:"mch_appid"`
	MchID      string `xml:"mchid"`            // 商户号
	OutTradeNo string `xml:"partner_trade_no"` // 商户订单号
	ToUser     string `xml:"openid"`
	Amount     int    `xml:"amount"`
	// 企业付款描述信息
	Desc string `xml:"desc"`

	// optional
	IP string `xml:"spbill_create_ip,omitempty"`
	// 校验用户姓名选项
	CheckName bool   `xml:"-"`
	Device    string `xml:"device_info,omitempty"`
	// 收款用户真实姓名
	// 如果check_name设置为FORCE_CHECK，则必填用户真实姓名
	RealName string `xml:"re_user_name,omitempty"`
}

type transferer struct {
	XMLName xml.Name `xml:"xml"`
	Transferer
	// 校验用户姓名选项
	// NO_CHECK:不校验真实姓名
	// FORCE_CHECK:强校验真实姓名
	CheckName string `xml:"check_name"`
	NonceStr  string `xml:"nonce_str"`
	Sign      string `xml:"sign"` // 签名
}

type transferResponse struct {
	response
	AppID         string `xml:"mch_appid"` // 小程序ID
	MchID         string `xml:"mchid"`
	Device        string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	OutTradeNo    string `xml:"partner_trade_no"` // 商户订单号
	TransactionID string `xml:"payment_no"`
	// 微信支付成功时间
	// format: 2015-05-19 15:26:59
	Datetime string `xml:"payment_time"`
}

// TransferResponse 转账返回数据
type TransferResponse struct {
	transferResponse
	Datetime time.Time
}

// 请求前准备
func (t *Transferer) prepare(key string) (transferer, error) {
	tra := transferer{
		Transferer: *t,
		NonceStr:   util.RandomString(32),
	}

	signData := map[string]string{
		"mch_appid":        tra.AppID,
		"mchid":            tra.MchID,
		"nonce_str":        tra.NonceStr,
		"partner_trade_no": tra.OutTradeNo,
		"openid":           tra.ToUser,
		"check_name":       tra.CheckName,
		"desc":             tra.Desc,
		"amount":           strconv.Itoa(tra.Amount),
	}

	if t.CheckName {
		tra.CheckName = "FORCE_CHECK"
		if t.RealName == "" {
			return tra, errors.New("选择校验用户姓名时用户姓名不能为空")
		}
		signData["re_user_name"] = tra.RealName

	} else {
		tra.CheckName = "NO_CHECK"
	}
	signData["check_name"] = tra.CheckName

	if t.IP == "" {
		ip, err := util.FetchIP()
		if err != nil {
			return tra, err
		}

		tra.IP = ip.String()
	}
	signData["spbill_create_ip"] = tra.IP

	if t.Device != "" {
		signData["device_info"] = tra.Device
	}

	sign, err := util.SignByMD5(signData, key)
	if err != nil {
		return tra, err
	}
	tra.Sign = sign

	return tra, nil
}

// Transfer 转账到微信用户零钱
func (t Transferer) Transfer(key string, certPath, keyPath string) (res TransferResponse, err error) {
	reqData, err := t.prepare(key)
	if err != nil {
		return
	}

	resData, err := util.TSLPostXML(baseURL+transferAPI, reqData, certPath, keyPath)
	if err != nil {
		return
	}
	var tres transferResponse
	if err = xml.Unmarshal(resData, &tres); err != nil {
		return
	}

	if err = tres.Check(); err != nil {
		return
	}

	res.transferResponse = tres
	res.Datetime, err = time.Parse(transferTimeFormat, tres.Datetime)

	return
}
