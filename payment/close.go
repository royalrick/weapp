package payment

import (
	"encoding/xml"

	"github.com/medivhzhan/weapp/util"
)

const (
	closeAPI = "/pay/closeorder"
)

// Closer 关闭订单
type Closer struct {
	// 必填 ...
	AppID      string `xml:"appid"`        // 小程序ID
	MchID      string `xml:"mch_id"`       // 商户号
	OutTradeNo string `xml:"out_trade_no"` // 商户订单号
}

type closer struct {
	XMLName xml.Name `xml:"xml"`
	Closer
	Sign     string `xml:"sign"`                // 签名
	NonceStr string `xml:"nonce_str"`           // 随机字符串
	SignType string `xml:"sign_type,omitempty"` // 签名类型: 目前支持HMAC-SHA256和MD5，默认为MD5
}

// CloseResponse 请求关闭订单时的返回
type CloseResponse struct {
	AppID    string `xml:"appid"`
	MchID    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`
}

type closeResponse struct {
	response
	CloseResponse
}

func (c *Closer) prepare(key string) (closer, error) {
	clo := closer{
		Closer:   *c,
		SignType: "MD5",
		NonceStr: util.RandomString(32),
	}

	signData := map[string]string{
		"appid":        clo.AppID,
		"mch_id":       clo.MchID,
		"nonce_str":    clo.NonceStr,
		"out_trade_no": clo.OutTradeNo,
		"sign_type":    clo.SignType,
	}

	sign, err := util.SignByMD5(signData, key)
	if err != nil {
		return clo, err
	}
	clo.Sign = sign

	return clo, nil
}

// Close 发起关闭支付请求
func (c Closer) Close(key string) (cres CloseResponse, err error) {
	data, err := c.prepare(key)
	if err != nil {
		return
	}

	resData, err := util.PostXML(baseURL+closeAPI, data)
	if err != nil {
		return
	}

	var res closeResponse
	if err = xml.Unmarshal(resData, &res); err != nil {
		return
	}
	err = res.Check()
	if err != nil {
		return
	}

	cres = res.CloseResponse
	return
}
