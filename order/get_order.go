package order

import "github.com/medivhzhan/weapp/v3/request"

type GetOrderRequest struct {
	TransactionId   string `json:"transaction_id"`    //原支付交易对应的微信订单号。
	MerchantId      string `json:"merchant_id"`       //支付下单商户的商户号，由微信支付生成并下发。
	SubMerchantId   string `json:"sub_merchant_id"`   //二级商户号。
	MerchantTradeNo string `json:"merchant_trade_no"` //商户系统内部订单号，只能是数字、大小写字母`_-*`且在同一个商户号下唯一。
}

type Shipping struct {
	DeliveryMode        int    `json:"delivery_mode"`         //发货模式，发货模式枚举值：1、UNIFIED_DELIVERY（统一发货）2、SPLIT_DELIVERY（分拆发货） 示例值: UNIFIED_DELIVERY
	LogisticsType       int    `json:"logistics_type"`        //物流模式，发货方式枚举值：1、实体物流配送采用快递公司进行实体物流配送形式 2、同城配送 3、虚拟商品，虚拟商品，例如话费充值，点卡等，无实体配送形式 4、用户自提
	FinishShipping      bool   `json:"finish_shipping"`       //是否已完成全部发货。
	FinishShippingCount int    `json:"finish_shipping_count"` //已完成全部发货的次数，未完成时为 0，完成时为 1，重新发货并完成后为 2。
	GoodsDesc           string `json:"goods_desc"`            //在小程序后台发货信息录入页录入的商品描述。
	ShippingList        []struct {
		TrackingNo     string `json:"tracking_no"`     //物流单号，示例值: "323244567777"。
		ExpressCompany string `json:"express_company"` //同城配送公司名或物流公司编码，快递公司ID，参见「查询物流公司编码列表」 示例值: "DHL"。
		GoodsDesc      string `json:"goods_desc"`      //使用上传物流信息 API 录入的该物流信息的商品描述。
		UploadTime     int    `json:"upload_time"`     //该物流信息的上传时间，时间戳形式。
		Contact        struct {
			ConsignorContact string `json:"consignor_contact"` //寄件人联系方式。
			ReceiverContact  string `json:"receiver_contact"`  //收件人联系方式。
		}
	} `json:"shipping_list"` //物流信息列表，发货物流单列表，支持统一发货（单个物流单）和分拆发货（多个物流单）两种模式。
}

type OrderStruct struct {
	TransactionId   string   `json:"transaction_id"`    //原支付交易对应的微信订单号。
	MerchantId      string   `json:"merchant_id"`       //支付下单商户的商户号，由微信支付生成并下发。
	SubMerchantId   string   `json:"sub_merchant_id"`   //二级商户号。
	MerchantTradeNo string   `json:"merchant_trade_no"` //商户系统内部订单号，只能是数字、大小写字母`_-*`且在同一个商户号下唯一。
	Description     string   `json:"description"`       //以分号连接的该支付单的所有商品描述，当超过120字时自动截断并以 “...” 结尾。
	PaidAmount      int      `json:"paid_amount"`       //支付单实际支付金额，整型，单位：分钱。
	Openid          string   `json:"openid"`            //支付者openid。
	TradeCreateTime int      `json:"trade_create_time"` //交易创建时间，时间戳形式。
	PayTime         int      `json:"pay_time"`          //支付时间，时间戳形式。
	InComplaint     bool     `json:"in_complaint"`      //是否处在交易纠纷中。
	OrderState      int      `json:"order_state"`       //订单状态枚举：(1) 待发货；(2) 已发货；(3) 确认收货；(4) 交易完成；(5) 已退款。
	Shipping        Shipping `json:"shipping"`          //发货信息。
}

type GetOrderResponse struct {
	request.CommonError
	Order *OrderStruct `json:"order"`
}

// GetOrder 查询订单发货状态
func (cli *Order) GetOrder(req *GetOrderRequest) (*GetOrderResponse, error) {

	url, err := cli.combineURI("/wxa/sec/order/get_order", nil, true)
	if err != nil {
		return nil, err
	}

	rsp := new(GetOrderResponse)
	if err := cli.request.Post(url, req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
