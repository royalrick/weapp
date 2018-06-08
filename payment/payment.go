// Package payment 微信支付
package payment

import (
	"encoding/xml"
	"errors"
	"net"
	"strconv"

	"github.com/medivhzhan/weapp/util"
)

const (
	baseURI = "https://api.mch.weixin.qq.com"

	unifyAPI = "/pay/unifiedorder"
)

// Params 前端调用支付必须的参数
// 注意返回后得大小写格式不能变动
type Params struct {
	Timestamp int64  `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
	Package   string `json:"package"`
}

// Payment 支付基础表单
type Payment struct {
	XMLName  xml.Name `xml:"xml" json:"-"`
	AppID    string   `xml:"appid"`               // 小程序ID
	MchID    string   `xml:"mch_id"`              // 商户号
	NonceStr string   `xml:"nonce_str"`           // 随机字符串
	Sign     string   `xml:"sign"`                // 签名
	SignType string   `xml:"sign_type,omitempty"` // 签名类型: 目前支持HMAC-SHA256和MD5，默认为MD5
	TotalFee float64  `xml:"total_fee"`           // 标价金额
}

// BaseResponse 基础返回数据
type BaseResponse struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	ReturnCode string   `xml:"return_code"`          // 返回状态码: SUCCESS/FAIL
	ReturnMsg  string   `xml:"return_msg,omitempty"` // 返回信息: 返回信息，如非空，为错误原因
}

// Response 请求返回结果
type Response struct {
	BaseResponse
	ResultCode string `xml:"result_code,omitempty"`  // 业务结果: SUCCESS/FAIL SUCCESS退款申请接收成功，结果通过退款查询接口查询 FAIL 提交业务失败
	ErrCode    string `xml:"err_code,omitempty"`     // 错误码
	ErrCodeDes string `xml:"err_code_des,omitempty"` // 错误代码描述
}

// Check 检测返回信息是否包含错误
func (res BaseResponse) Check() error {
	switch res.ReturnCode {
	case "SUCCESS":
		return nil
	case "FAIL":
		return errors.New(res.ReturnMsg)
	default:
		return errors.New("未知微信返回状态码: " + res.ReturnCode)
	}
}

// Check 检测返回信息是否包含错误
func (res Response) Check() error {

	if err := res.BaseResponse.Check(); err != nil {
		return err
	}

	switch res.ResultCode {
	case "SUCCESS":
		return nil
	case "FAIL":
		return errors.New(res.ErrCodeDes)
	default:
		return errors.New("未知微信返回业务结果代码: " + res.ResultCode)
	}
}

// Order 商户统一订单
type Order struct {
	Payment

	NotifyURL      string `xml:"notify_url"`            // 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	OpenID         string `xml:"openid"`                // 下单用户ID
	DeviceInfo     string `xml:"device_info,omitempty"` // 设备号， 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	Body           string `xml:"body"`                  // 商品描述
	Detail         string `xml:"detail,omitempty"`      // 商品详情
	Attach         string `xml:"attach,omitempty"`      // 附加数据
	OutTradeNo     string `xml:"out_trade_no"`          // 商户订单号
	FeeType        string `xml:"fee_type,omitempty"`    // 标价币种
	SPBillCreateIP net.IP `xml:"spbill_create_ip"`      // 终端IP
	TimeStart      string `xml:"time_start,omitempty"`  // 交易起始时间 格式为yyyyMMddHHmmss
	TimeExpire     string `xml:"time_expire,omitempty"` // 交易结束时间 订单失效时间 格式为yyyyMMddHHmmss
	GoodsTag       string `xml:"goods_tag,omitempty"`   // 订单优惠标记，使用代金券或立减优惠功能时需要的参数，
	TradeType      string `xml:"trade_type"`            // 小程序取值如下：JSAPI
	LimitPay       string `xml:"limit_pay,omitempty"`   // 上传此参数 no_credit 可限制用户不能使用信用卡支付
}

// GetParams 获取支付参数
// @ appID
// @ key
// @ nonceStr
// @ prepayID
// @ timestamp
func GetParams(appID, key, nonceStr, prepayID string, timestamp int64) (p Params, err error) {

	sign, err := util.SignByMD5(map[string]string{
		"key":       key,
		"appId":     appID,
		"signType":  "MD5",
		"nonceStr":  nonceStr,
		"package":   "prepay_id" + "=" + prepayID,
		"timeStamp": strconv.FormatInt(timestamp, 10),
	})
	if err != nil {
		return
	}

	p = Params{
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		SignType:  "MD5",
		PaySign:   sign,
	}

	return
}

// Unify 统一下单
func (o *Order) Unify() {
	if o.TradeType == "" {
		o.TradeType = "JSAPI"
	}
}
