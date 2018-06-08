package payment

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
)

const (
	refundAPI = "/secapi/pay/refund"
)

// Refunder 退款表单数据
type Refunder struct {
	Payment

	TransactionID string  `xml:"transaction_id,emitempty"`  // 微信订单号: 微信生成的订单号，在支付通知中有返回。和商户订单号二选一
	OutTradeNo    string  `xml:"out_trade_no,emitempty"`    // 商户订单号: 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 和微信订单号二选一
	OutRefundNo   string  `xml:"out_refund_no"`             // 商户退款单号: 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundFee     float64 `xml:"refund_fee"`                // 退款金额: 退款总金额，订单总金额，单位为分，只能为整数
	RefundFeeType string  `xml:"refund_fee_type,emitempty"` // 货币种类: 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY
	RefundDesc    string  `xml:"refund_desc,emitempty"`     // 退款原因: 若商户传入，会在下发给用户的退款消息中体现退款原因

	// 退款资金来源: 仅针对老资金流商户使用
	// REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款（默认使用未结算资金退款）
	// REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款
	RefundAccount string `xml:"refund_account,emitempty"`

	// 退款结果通知url: 异步接收微信支付退款结果通知的回调地址
	// 通知 URL 必须为外网可访问且不允许带参数
	// 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效。
	NotifyURL string `xml:"notify_url,omitempty"`
}

// Refund 发起退款请求
func (r *Refunder) Refund() error {

	if r.TransactionID == "" && r.OutRefundNo == "" {
		return errors.New("out_trade_no 和 out_trade_no 必须填写一个")
	}

	if r.TransactionID != "" && r.OutRefundNo != "" {
		return errors.New("out_trade_no 和 out_trade_no 只能填写一个")
	}

	r.SignType = "MD5"
	sign, err := util.SignByMD5(map[string]string{
	// dev: 需要MD5加密签名的数据
	})
	if err != nil {
		return err
	}
	r.Sign = sign

	body, err := xml.Marshal(r)
	if err != nil {
		return err
	}
	res, err := http.Post(baseURI+refundAPI, "application/xml", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(weapp.WeChatServerError)
	}

	var resp Response
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	return resp.Check()
}
