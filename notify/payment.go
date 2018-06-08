package notify

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/medivhzhan/weapp/payment"
	"github.com/medivhzhan/weapp/util"
)

// PaidNotify 支付结果返回数据
type PaidNotify struct {
	payment.Response
	payment.Payment

	Device        string  `xml:"device_info,omitempty"`
	OpenID        string  `xml:"openid"`
	TradeType     string  `xml:"trade_type"`                     // 交易类型 JSAPI
	Bank          string  `xml:"bank_type"`                      // 银行类型，采用字符串类型的银行标识
	Settlement    float64 `xml:"settlement_total_fee,omitempty"` // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType       string  `xml:"fee_type,omitempty"`             // 货币种类: 符合ISO4217标准的三位字母代码，默认人民币：CNY
	CashFee       float64 `xml:"cash_fee"`                       // 现金支付金额订单的现金支付金额
	CashFeeType   string  `xml:"cash_fee_type,omitempty"`        // 现金支付货币类型: 符合ISO4217标准的三位字母代码，默认人民币：CNY
	CouponFee     float64 `xml:"coupon_fee,omitempty"`           // 总代金券金额: 代金券金额<=订单金额，订单金额-代金券金额=现金支付金额
	CouponCount   int     `xml:"coupon_count,omitempty"`         // 代金券使用数量
	TransactionID string  `xml:"transaction_id"`                 // 微信支付订单号
	Attach        string  `xml:"attach,omitempty"`               // 商家数据包，原样返回

	// 商户系统内部订单号: 要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	OutTradeNo string `xml:"out_trade_no"`

	// 支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010
	Timeend string `xml:"time_end"`
}

// RefundedResponse 退款结果通知
type RefundedResponse struct {
	payment.BaseResponse

	AppID    string `xml:"appid,omitempty"` // 小程序ID
	MchID    string `xml:"mch_id"`          // 商户号
	NonceStr string `xml:"nonce_str"`       // 随机字符串
	ReqInfo  []byte `xml:"req_info"`        // 加密信息
}

// RefundedNotify 解密后的退款通知消息体
type RefundedNotify struct {
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
	// SUCCESS 退款成功
	// CHANGE 退款异常
	// REFUNDCLOSE 退款关闭
	RefundStatus string `xml:"refund_status"`

	// 退款成功时间
	// 资金退款至用户帐号的时间，格式2017-12-15 09:46:01
	SuccessTime string `xml:"success_time,omitempty"`

	// 退款入账账户:取当前退款单的退款入账方
	// 1）退回银行卡： {银行名称}{卡类型}{卡尾号}
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

// paymentReturn 收到退款和支付通知后返回给微信服务器的消息
type paymentReturn struct {
	Code string `xml:"return_code"`          // 返回状态码: SUCCESS/FAIL
	Msg  string `xml:"return_msg,omitempty"` // 返回信息: 返回信息，如非空，为错误原因
}

// 根据结果创建返回数据
// ok 是否处理成功
// msg 处理不成功原因
func newPaymentReturn(ok bool, msg string) paymentReturn {

	ret := paymentReturn{Msg: msg}

	if ok {
		ret.Code = "SUCCESS"
	} else {
		ret.Code = "FAIL"
	}

	return ret
}

// HandlePaidNotify 处理支付结果通知
func HandlePaidNotify(res http.ResponseWriter, req *http.Request, fuck func(PaidNotify) (bool, string)) error {
	// dev: 是否需要限制请求方法
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var ntf PaidNotify
	if err := xml.Unmarshal(body, &ntf); err != nil {
		return err
	}

	if err := ntf.Check(); err != nil {
		return err
	}

	pr := newPaymentReturn(fuck(ntf))

	b, err := xml.Marshal(pr)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)
	_, err = res.Write(b)

	return err
}

// HandleRefundedNotify 处理退款结果通知
// key: 微信支付 KEY
func HandleRefundedNotify(res http.ResponseWriter, req *http.Request, key string, fuck func(RefundedNotify) (bool, string)) error {
	// dev: 是否需要限制请求方法
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var ref RefundedResponse

	if err := xml.Unmarshal(body, &ref); err != nil {
		return err
	}

	if err := ref.Check(); err != nil {
		return err
	}

	bts, err := util.AesECBDecrypt(ref.ReqInfo, key)
	if err != nil {
		return err
	}

	var ntf RefundedNotify
	if err := xml.Unmarshal(bts, &ntf); err != nil {
		return err
	}

	pr := newPaymentReturn(fuck(ntf))

	b, err := xml.Marshal(pr)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)
	_, err = res.Write(b)

	return err
}
