package payment

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/medivhzhan/weapp/util"
)

const (
	refundAPI = "/secapi/pay/refund"
)

// Refunder 退款表单数据
type Refunder struct {
	// 必填 ...
	AppID         string `xml:"appid"`  // 小程序ID
	MchID         string `xml:"mch_id"` // 商户号
	TotalFee      int    `xml:"total_fee"`
	RefundFee     int    `xml:"refund_fee"`               // 退款金额: 退款总金额，订单总金额，单位为分，只能为整数
	TransactionID string `xml:"transaction_id,omitempty"` // 微信订单号: 微信生成的订单号，在支付通知中有返回。和商户订单号二选一
	OutTradeNo    string `xml:"out_trade_no,omitempty"`   // 商户订单号: 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 和微信订单号二选一
	OutRefundNo   string `xml:"out_refund_no"`            // 商户退款单号: 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。

	// 选填 ...
	// RefundFeeType string `xml:"refund_fee_type,omitempty"` // 货币种类: 货币类型，符合ISO 4217标准的三位字母代码，默认人民币: CNY
	RefundDesc string `xml:"refund_desc,omitempty"` // 退款原因: 若商户传入，会在下发给用户的退款消息中体现退款原因

	// 退款结果通知url: 异步接收微信支付退款结果通知的回调地址
	// 通知 URL 必须为外网可访问且不允许带参数
	// 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效。
	NotifyURL string `xml:"notify_url,omitempty"`

	// 退款资金来源: 仅针对老资金流商户使用
	// REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款（默认使用未结算资金退款）
	// REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款
	// RefundAccount string `xml:"refund_account,omitempty"`
}

type refunder struct {
	XMLName xml.Name `xml:"xml"`
	Refunder
	Sign     string `xml:"sign"`                // 签名
	NonceStr string `xml:"nonce_str"`           // 随机字符串
	SignType string `xml:"sign_type,omitempty"` // 签名类型: 目前支持HMAC-SHA256和MD5，默认为MD5
}

// 请求前准备
func (r Refunder) prepare(key string) (refunder, error) {
	ref := refunder{
		Refunder: r,
		SignType: "MD5",
		NonceStr: util.RandomString(32),
	}

	signData := map[string]string{
		"appid":         ref.AppID,
		"mch_id":        ref.MchID,
		"nonce_str":     ref.NonceStr,
		"out_refund_no": ref.OutRefundNo,
		"total_fee":     strconv.Itoa(ref.TotalFee),
		"refund_fee":    strconv.Itoa(ref.RefundFee),
		"sign_type":     ref.SignType,
	}

	switch {
	case r.TransactionID == "" && r.OutRefundNo == "":
		return ref, errors.New("out_trade_no 和 transition_id 必须填写一个")
	case r.TransactionID != "" && r.OutRefundNo != "":
		return ref, errors.New("out_trade_no 和 transition_id 只能填写一个")
	case r.TransactionID != "":
		signData["transaction_id"] = r.TransactionID
	case r.OutRefundNo != "":
		signData["out_trade_no"] = r.OutTradeNo
	}

	if r.RefundDesc != "" {
		signData["refund_desc"] = r.RefundDesc
	}

	if r.NotifyURL != "" {
		signData["notify_url"] = r.NotifyURL
	}

	sign, err := util.SignByMD5(signData, key)
	ref.Sign = sign

	return ref, err
}

// RefundedResponse 请求退款返回数据
type RefundedResponse struct {
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	TransactionID string `xml:"transaction_id"` // 微信订单号: 微信生成的订单号，在支付通知中有返回。和商户订单号二选一
	OutTradeNo    string `xml:"out_trade_no"`   // 商户订单号: 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 和微信订单号二选一
	OutRefundNo   string `xml:"out_refund_no"`  // 商户退款单号: 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	// 微信退款单号
	RefundID string `xml:"refund_id"`
	// 退款总金额,单位为分,可以做部分退款
	RefundFee int `xml:"refund_fee"`
	// 应结退款金额
	// 去掉非充值代金券退款金额后的退款金额，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	SettlementRefundFee int `xml:"settlement_refund_fee"`
	// 标价金额
	// 订单总金额，单位为分，只能为整数
	TotalFee int `xml:"total_fee"`
	// 应结订单金额
	// 去掉非充值代金券金额后的订单总金额，应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	SettlementTotalFee int `xml:"settlement_total_fee"`
	// 标价币种
	// FeeType            int `xml:"fee_type"`
	// 现金支付金额
	CashFee       int    `xml:"cash_fee"`
	CashRefundFee int    `xml:"cash_refund_fee"`
	Sign          string `xml:"sign"`
	NonceStr      string `xml:"nonce_str"`

	// TODO: ...
	// coupon_type_$n
	// coupon_refund_fee
	// coupon_refund_fee_$n
	// coupon_refund_count
	// coupon_refund_id_$n
}

// refundedResponse 支付返回集合
type refundedResponse struct {
	response
	RefundedResponse
}

// Refund 发起退款请求
func (r Refunder) Refund(key, certPath, keyPath string) (rres RefundedResponse, err error) {
	data, err := r.prepare(key)
	if err != nil {
		return
	}

	resData, err := util.TSLPostXML(baseURL+refundAPI, data, certPath, keyPath)
	if err != nil {
		return
	}

	var res refundedResponse
	if err = xml.Unmarshal(resData, &res); err != nil {
		return
	}
	err = res.Check()
	if err != nil {
		return
	}

	rres = res.RefundedResponse
	return
}

//  退款结果通知
type refundNotify struct {
	AppID      string `xml:"appid"`       // 小程序 APPID
	MchID      string `xml:"mch_id"`      // 商户号
	NonceStr   string `xml:"nonce_str"`   // 随机字符串
	Ciphertext string `xml:"req_info"`    // 加密信息
	ReturnCode string `xml:"return_code"` // 返回状态码: SUCCESS/FAIL
	ReturnMsg  string `xml:"return_msg"`  // 返回信息: 返回信息，如非空，为错误原因
}

// 检测返回信息是否包含错误
func (res refundNotify) Check() error {

	if res.ReturnCode != "SUCCESS" {
		return errors.New("交易失败: " + res.ReturnMsg)
	}

	return nil
}

// RefundedNotify 解密后的退款通知消息体
type RefundedNotify struct {
	AppID         string // 小程序ID
	MchID         string // 商户号
	NonceStr      string // 随机字符串
	TransactionID string `xml:"transaction_id"` // 微信支付订单号
	// 商户系统内部订单号: 要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	OutTradeNo  string  `xml:"out_trade_no"`
	RefundID    string  `xml:"refund_id"`     // 微信退款单号
	OutRefundNo string  `xml:"out_refund_no"` // 商户退款单号
	TotalFee    float64 `xml:"total_fee"`     // 标价金额
	// 当该订单有使用非充值券时，返回此字段。
	// 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	Settlement float64 `xml:"settlement_total_fee,omitempty"`
	RefundFee  float64 `xml:"refund_fee"` // 退款总金额,单位为分
	// 退款金额
	// 退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	SettlementRefund float64 `xml:"settlement_refund_fee"`
	// 退款状态
	// SUCCESS 退款成功 | CHANGE 退款异常 | REFUNDCLOSE 退款关闭
	RefundStatus string `xml:"refund_status"`
	// 退款成功时间
	// 资金退款至用户帐号的时间，格式2017-12-15 09:46:01
	SuccessTime string `xml:"success_time,omitempty"`
	// 退款入账账户:取当前退款单的退款入账方
	// 1）退回银行卡:  {银行名称}{卡类型}{卡尾号}
	// 2）退回支付用户零钱: 支付用户零钱
	// 3）退还商户: 商户基本账户 商户结算银行账户
	// 4）退回支付用户零钱通: 支付用户零钱通
	ReceiveAccount string `xml:"refund_recv_accout"`
	// 退款资金来源
	// REFUND_SOURCE_RECHARGE_FUNDS 可用余额退款/基本账户
	// REFUND_SOURCE_UNSETTLED_FUNDS 未结算资金退款
	RefundAccount string `xml:"refund_account"`
	// 退款发起来源
	// API接口
	// VENDOR_PLATFORM商户平台
	Source string `xml:"refund_request_source"`
}

// HandleRefundedNotify 处理退款结果通知
// key: 微信支付 KEY
func HandleRefundedNotify(res http.ResponseWriter, req *http.Request, key string, fuck func(RefundedNotify) (bool, string)) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var ref refundNotify

	if err := xml.Unmarshal(body, &ref); err != nil {
		return err
	}

	if err := ref.Check(); err != nil {
		return err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(ref.Ciphertext)
	if err != nil {
		return err
	}
	key, err = util.MD5(key)
	if err != nil {
		return err
	}
	key = strings.ToLower(key)

	bts, err := util.AesECBDecrypt(ciphertext, []byte(key))
	if err != nil {
		return err
	}

	ntf := RefundedNotify{
		AppID:    ref.AppID,
		NonceStr: ref.NonceStr,
		MchID:    ref.MchID,
	}

	if err := xml.Unmarshal(bts, &ntf); err != nil {
		return err
	}

	pr := newReplay(fuck(ntf))

	b, err := xml.Marshal(pr)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)
	_, err = res.Write(b)

	return err
}
