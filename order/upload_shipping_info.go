package order

import (
	"github.com/medivhzhan/weapp/v3/request"
)

type OrderKey struct {
	OrderNumberType int    `json:"order_number_type"`        //订单单号类型，用于确认需要上传详情的订单。枚举值1，使用下单商户号和商户侧单号；枚举值2，使用微信支付单号。
	TransactionId   string `json:"transaction_id,omitempty"` //原支付交易对应的微信订单号
	Mchid           string `json:"mchid,omitempty"`          //支付下单商户的商户号，由微信支付生成并下发。
	OutTradeNo      string `json:"out_trade_no,omitempty"`   //商户系统内部订单号，只能是数字、大小写字母`_-*`且在同一个商户号下唯一
}

type Payer struct {
	Openid string `json:"openid"`
}

type Contact struct {
	ConsignorContact string `json:"consignor_contact,omitempty"` //寄件人联系方式，寄件人联系方式，采用掩码传输，最后4位数字不能打掩码 示例值: `189****1234, 021-****1234, ****1234, 0**2-***1234, 0**2-******23-10, ****123-8008` 值限制: 0 ≤ value ≤ 1024
	ReceiverContact  string `json:"receiver_contact,omitempty"`  //收件人联系方式，收件人联系方式为，采用掩码传输，最后4位数字不能打掩码 示例值: `189****1234, 021-****1234, ****1234, 0**2-***1234, 0**2-******23-10, ****123-8008` 值限制: 0 ≤ value ≤ 1024
}

type ShippingList struct {
	TrackingNo     string   `json:"tracking_no,omitempty"`     //物流单号，物流快递发货时必填，示例值: 323244567777 字符字节限制: [1, 128]
	ExpressCompany string   `json:"express_company,omitempty"` //物流公司编码，快递公司ID，参见「查询物流公司编码列表」，物流快递发货时必填， 示例值: DHL 字符字节限制: [1, 128]
	ItemDesc       string   `json:"item_desc"`                 //商品信息，例如：微信红包抱枕*1个，限120个字以内
	Contact        *Contact `json:"contact,omitempty"`         //联系方式，当发货的物流公司为顺丰时，联系方式为必填，收件人或寄件人联系方式二选一
}

type UploadShippingInfoRequest struct {
	OrderKey       OrderKey       `json:"order_key"`                  //合单订单，需要上传物流详情的合单订单，根据订单类型二选一
	DeliveryMode   string         `json:"delivery_mode"`              //发货模式，发货模式枚举值：1、UNIFIED_DELIVERY（统一发货）2、SPLIT_DELIVERY（分拆发货） 示例值: UNIFIED_DELIVERY
	LogisticsType  int            `json:"logistics_type"`             //物流模式，发货方式枚举值：1、实体物流配送采用快递公司进行实体物流配送形式 2、同城配送 3、虚拟商品，虚拟商品，例如话费充值，点卡等，无实体配送形式 4、用户自提
	IsAllDelivered bool           `json:"is_all_delivered,omitempty"` //分拆发货模式时必填，用于标识分拆发货模式下是否已全部发货完成，只有全部发货完成的情况下才会向用户推送发货完成通知。示例值: true/false
	ShippingList   []ShippingList `json:"shipping_list"`              //物流信息列表，发货物流单列表，支持统一发货（单个物流单）和分拆发货（多个物流单）两种模式，多重性: [1, 10]
	UploadTime     string         `json:"upload_time"`                //上传时间，用于标识请求的先后顺序 示例值: `2022-12-15T13:29:35.120+08:00` RFC 3339 格式
	Payer          Payer          `json:"payer"`                      //支付者，支付者信息
}

// UploadShippingInfo 发货信息录入接口
func (cli *Order) UploadShippingInfo(req *UploadShippingInfoRequest) (*request.CommonError, error) {

	url, err := cli.combineURI("/wxa/sec/order/upload_shipping_info", nil, true)
	if err != nil {
		return nil, err
	}

	rsp := new(request.CommonError)
	if err := cli.request.Post(url, req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
