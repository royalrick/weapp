package order

import "github.com/medivhzhan/weapp/v3/request"

type NotifyConfirmReceiveRequest struct {
	TransactionId   string `json:"transaction_id,omitempty"`    //原支付交易对应的微信订单号
	MerchantId      string `json:"merchant_id,omitempty"`       //支付下单商户的商户号，由微信支付生成并下发
	SubMerchantId   string `json:"sub_merchant_id,omitempty"`   //二级商户号
	MerchantTradeNo string `json:"merchant_trade_no,omitempty"` //商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	ReceivedTime    int    `json:"received_time"`               //快递签收时间，时间戳形式。

}

// NotifyConfirmReceive 确认收货提醒接口
func (cli *Order) NotifyConfirmReceive(req *IsTradeManagedRequest) (*request.CommonError, error) {

	url, err := cli.combineURI("/wxa/sec/order/notify_confirm_receive", nil, true)
	if err != nil {
		return nil, err
	}

	rsp := new(request.CommonError)
	if err := cli.request.Post(url, req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
