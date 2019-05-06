// Package payment 微信支付
package payment

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/medivhzhan/weapp/util"
)

const (
	baseURL = "https://api.mch.weixin.qq.com"

	unifyAPI          = "/pay/unifiedorder"
	queryAPI          = "/pay/orderquery"
	paymentTimeFormat = "20060102150405"
)

// Params 前端调用支付必须的参数
// 注意返回后得大小写格式不能变动
type Params struct {
	Timestamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
	Package   string `json:"package"`
}

// Order 商户统一订单
type Order struct {
	// 必填 ...
	AppID      string `xml:"appid"`        // 小程序ID
	MchID      string `xml:"mch_id"`       // 商户号
	TotalFee   int    `xml:"total_fee"`    // 标价金额
	NotifyURL  string `xml:"notify_url"`   // 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	OpenID     string `xml:"openid"`       // 下单用户ID
	Body       string `xml:"body"`         // 商品描述
	OutTradeNo string `xml:"out_trade_no"` // 商户订单号

	// 选填 ...
	IP        string    `xml:"spbill_create_ip,omitempty"` // 终端IP
	NoCredit  bool      `xml:"-"`                          // 上传此参数 no_credit 可限制用户不能使用信用卡支付
	StartedAt time.Time `xml:"-"`                          // 交易起始时间 格式为yyyyMMddHHmmss
	ExpiredAt time.Time `xml:"-"`                          // 交易结束时间 订单失效时间 格式为yyyyMMddHHmmss
	Tag       string    `xml:"goods_tag,omitempty"`        // 订单优惠标记，使用代金券或立减优惠功能时需要的参数，
	Detail    string    `xml:"detail,omitempty"`           // 商品详情
	Attach    string    `xml:"attach,omitempty"`           // 附加数据
}

// 下单所需所有数据
type order struct {
	XMLName xml.Name `xml:"xml"`
	Order
	Sign      string `xml:"sign"`                // 签名
	NonceStr  string `xml:"nonce_str"`           // 随机字符串
	TradeType string `xml:"trade_type"`          // 小程序取值如下: JSAPI
	SignType  string `xml:"sign_type,omitempty"` // 签名类型: 目前支持HMAC-SHA256和MD5，默认为MD5

	NoCredit  string `xml:"limit_pay,omitempty"`   // 上传此参数 no_credit 可限制用户不能使用信用卡支付
	StartedAt string `xml:"time_start,omitempty"`  // 交易起始时间 格式为yyyyMMddHHmmss
	ExpiredAt string `xml:"time_expire,omitempty"` // 交易结束时间 订单失效时间 格式为yyyyMMddHHmmss
}

// 请求前准备
func (o *Order) prepare(key string) (order, error) {

	od := order{
		Order:     *o,
		TradeType: "JSAPI",
		SignType:  "MD5",
		NonceStr:  util.RandomString(32),
	}

	signData := map[string]string{
		"appid":        od.AppID,
		"body":         od.Body,
		"mch_id":       od.MchID,
		"nonce_str":    od.NonceStr,
		"notify_url":   od.NotifyURL,
		"openid":       od.OpenID,
		"out_trade_no": od.OutTradeNo,
		"total_fee":    strconv.Itoa(od.TotalFee),
		"trade_type":   od.TradeType,
		"sign_type":    od.SignType,
	}

	if o.IP == "" {
		ip, err := util.FetchIP()
		if err != nil {
			return od, err
		}

		od.IP = ip.String()
	}
	signData["spbill_create_ip"] = od.IP

	if !o.StartedAt.IsZero() {
		od.StartedAt = o.StartedAt.Format(paymentTimeFormat)
		signData["time_start"] = od.StartedAt
	}

	if !o.ExpiredAt.IsZero() {
		od.ExpiredAt = o.ExpiredAt.Format(paymentTimeFormat)
		signData["time_expire"] = od.ExpiredAt
	}

	if o.Attach != "" {
		signData["attach"] = od.Attach
	}

	if o.Detail != "" {
		signData["detail"] = od.Detail
	}

	if o.Tag != "" {
		signData["goods_tag"] = od.Tag
	}

	if o.NoCredit {
		od.NoCredit = "no_credit"
		signData["limit_pay"] = od.NoCredit
	}

	sign, err := util.SignByMD5(signData, key)
	if err != nil {
		return od, err
	}
	od.Sign = sign

	return od, nil
}

// response 基础返回数据
type response struct {
	ReturnCode string `xml:"return_code"` // 返回状态码: SUCCESS/FAIL
	ReturnMsg  string `xml:"return_msg"`  // 返回信息: 返回信息，如非空，为错误原因
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
}

// Check 检测返回信息是否包含错误
func (res response) Check() error {
	if res.ReturnCode != "SUCCESS" {
		return errors.New("交易失败: " + res.ReturnMsg)
	}

	if res.ResultCode != "SUCCESS" {
		return errors.New("发生错误: " + res.ErrCodeDes)
	}

	return nil
}

// PaidResponse 支付返回面向用户的集合
type PaidResponse struct {
	AppID    string `xml:"appid"` // 小程序ID
	MchID    string `xml:"mch_id"`
	PrePayID string `xml:"prepay_id"`
	Sign     string `xml:"sign"`
	NonceStr string `xml:"nonce_str"`
}

// paidResponse 支付返回集合
type paidResponse struct {
	response
	PaidResponse
}

// GetParams 获取支付参数
//
// @appID 小程序 APPID
// @key 微信支付密钥
// @nonceStr 统一下单得到的 nonceStr
// @prepayID 统一下单得到的 prepayID
func GetParams(appID, key, nonceStr, prepayID string) (p Params, err error) {

	if len(nonceStr) > 32 {
		err = errors.New("随机字符串长度为32个字符以下")
		return
	}

	p.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	p.SignType = "MD5"
	p.NonceStr = nonceStr
	p.Package = "prepay_id=" + prepayID

	p.PaySign, err = util.SignByMD5(map[string]string{
		"appId":     appID,
		"signType":  p.SignType,
		"nonceStr":  nonceStr,
		"package":   p.Package,
		"timeStamp": p.Timestamp,
	}, key)

	return
}

// Unify 统一下单
//
// @key payment secret key
func (o Order) Unify(key string) (pres PaidResponse, err error) {

	reqData, err := o.prepare(key)
	if err != nil {
		return
	}

	data, err := util.PostXML(baseURL+unifyAPI, reqData)
	if err != nil {
		return
	}

	var res paidResponse
	if err = xml.Unmarshal(data, &res); err != nil {
		return
	}

	if err = res.Check(); err != nil {
		return
	}

	pres = res.PaidResponse
	return
}

// PaidNotify 支付结果返回数据
type PaidNotify struct {
	AppID         string  `xml:"appid"`               // 小程序ID
	MchID         string  `xml:"mch_id"`              // 商户号
	TotalFee      int     `xml:"total_fee"`           // 标价金额
	NonceStr      string  `xml:"nonce_str"`           // 随机字符串
	Sign          string  `xml:"sign"`                // 签名
	SignType      string  `xml:"sign_type,omitempty"` // 签名类型: 目前支持HMAC-SHA256和MD5，默认为MD5
	OpenID        string  `xml:"openid"`
	TradeType     string  `xml:"trade_type"`                     // 交易类型 JSAPI
	Bank          string  `xml:"bank_type"`                      // 银行类型，采用字符串类型的银行标识
	Settlement    float64 `xml:"settlement_total_fee,omitempty"` // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType       string  `xml:"fee_type,omitempty"`             // 货币种类: 符合ISO4217标准的三位字母代码，默认人民币: CNY
	CashFee       float64 `xml:"cash_fee"`                       // 现金支付金额订单的现金支付金额
	CashFeeType   string  `xml:"cash_fee_type,omitempty"`        // 现金支付货币类型: 符合ISO4217标准的三位字母代码，默认人民币: CNY
	CouponFee     float64 `xml:"coupon_fee,omitempty"`           // 总代金券金额: 代金券金额<=订单金额，订单金额-代金券金额=现金支付金额
	CouponCount   int     `xml:"coupon_count,omitempty"`         // 代金券使用数量
	TransactionID string  `xml:"transaction_id"`                 // 微信支付订单号
	Attach        string  `xml:"attach,omitempty"`               // 商家数据包，原样返回
	// 商户系统内部订单号: 要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	OutTradeNo string `xml:"out_trade_no"`
	// 支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010
	Timeend string `xml:"time_end"`
}

type paidNotify struct {
	response
	PaidNotify
}

// 收到退款和支付通知后返回给微信服务器的消息
type replay struct {
	Code string `xml:"return_code"` // 返回状态码: SUCCESS/FAIL
	Msg  string `xml:"return_msg"`  // 返回信息: 返回信息，如非空，为错误原因
}

// 根据结果创建返回数据
//
// ok 是否处理成功
// msg 处理不成功原因
func newReplay(ok bool, msg string) replay {

	ret := replay{Msg: msg}

	if ok {
		ret.Code = "SUCCESS"
	} else {
		ret.Code = "FAIL"
	}

	return ret
}

// HandlePaidNotify 处理支付结果通知
func HandlePaidNotify(res http.ResponseWriter, req *http.Request, fuck func(PaidNotify) (bool, string)) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var ntf paidNotify
	if err := xml.Unmarshal(body, &ntf); err != nil {
		return err
	}

	if err := ntf.Check(); err != nil {
		return err
	}

	replay := newReplay(fuck(ntf.PaidNotify))

	b, err := xml.Marshal(replay)
	if err != nil {
		return err
	}

	res.WriteHeader(http.StatusOK)
	_, err = res.Write(b)

	return err
}

type queryResponse struct {
	response
	OrderResponse
}

const (
	TradeStateSuccess    = "SUCCESS"    //支付成功
	TradeStateRefund     = "REFUND"     //转入退款
	TradeStateNotpay     = "NOTPAY"     //未支付
	TradeStateClosed     = "CLOSED"     //已关闭
	TradeStateRevoked    = "REVOKED"    //已撤销
	TradeStateUserpaying = "USERPAYING" //用户支付中
	TradeStatePayerror   = "PAYERROR"   //支付失败(其他原因，如银行返回失败)
)

type OrderResponse struct {
	DeviceInfo         string `xml:"device_info"`          // 设备号	微信支付分配的终端设备号，
	OpenId             string `xml:"openid"`               //用户标识	用户在商户appid下的唯一标识
	IsSubscribe        string `xml:"is_subscribe"`         //是否关注公众账号	用户是否关注公众账号，Y-关注，N-未关注
	TradeType          string `xml:"trade_type"`           //交易类型 	JSAPI	调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	TradeState         string `xml:"trade_state"`          //交易状态 SUCCESS—支付成功 REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭  REVOKED—已撤销（刷卡支付） USERPAYING--用户支付中 PAYERROR--支付失败(其他原因，如银行返回失败)
	BankType           string `xml:"bank_type"`            //付款银行 CMC	银行类型，采用字符串类型的银行标识
	TotalFee           int    `xml:"total_fee"`            //标价金额 订单总金额，单位为分
	SettlementTotalFee int    `xml:"settlement_total_fee"` //应结订单金额 当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType            string `xml:"fee_type"`             //标价币种	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TransactionId      string `xml:"transaction_id"`       //微信支付订单号
	OutTradeNo         string `xml:"out_trade_no"`         //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	Attach             string `xml:"attach"`               //附加数据，原样返回
	TimeEnd            string `xml:"time_end"`             //支付完成时间是 yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	CashFee            int    `xml:"cash_fee"`             //现金支付金额订单现金支付金额
	TradeStateDesc     string `xml:"trade_state_desc"`     //交易状态描述	支付失败，请重新下单支付	对当前查询订单状态的描述和下一步操作的指引
}

type Query struct {
	AppID         string `xml:"appid"`                    // 小程序ID
	MchID         string `xml:"mch_id"`                   // 商户号
	OutTradeNo    string `xml:"out_trade_no,omitempty"`   // 商户订单号
	TransactionID string `xml:"transaction_id,omitempty"` // 微信支付订单号
	SignType      string `xml:"sign_type,omitempty"`      // 签名类型: 目前支持HMAC-SHA256和MD5，默认为MD5
	NonceStr      string `xml:"nonce_str"`                // 随机字符串
	Sign          string `xml:"sign"`                     // 签名
}

func (q Query) Query(key string) (resp OrderResponse, err error) {
	q.SignType = "MD5"
	q.NonceStr = util.RandomString(32)
	signData := map[string]string{
		"appid":     q.AppID,
		"mch_id":    q.MchID,
		"nonce_str": q.NonceStr,
		"sign_type": q.SignType,
	}
	if q.OutTradeNo == "" && q.TransactionID == "" {
		err = errors.New("参数错误: 至少out_trade_no或者transaction_id")
		return
	}
	if q.OutTradeNo != "" {
		signData["out_trade_no"] = q.OutTradeNo
	}
	if q.TransactionID != "" {
		signData["transaction_id"] = q.TransactionID
	}
	sign, err := util.SignByMD5(signData, key)
	if err != nil {
		return
	}
	q.Sign = sign
	data, err := util.PostXML(baseURL+queryAPI, q)
	if err != nil {
		return
	}
	var res queryResponse
	if err = xml.Unmarshal(data, &res); err != nil {
		return
	}
	if err = res.Check(); err != nil {
		return
	}
	resp = res.OrderResponse
	return
}
