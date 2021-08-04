package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiAbnormalConfirm         = "/cgi-bin/express/local/business/order/confirm_return"
	apiAddDeliveryOrder        = "/cgi-bin/express/local/business/order/add"
	apiAddDeliveryTip          = "/cgi-bin/express/local/business/order/addtips"
	apiCancelDeliveryOrder     = "/cgi-bin/express/local/business/order/cancel"
	apiGetAllImmediateDelivery = "/cgi-bin/express/local/business/delivery/getall"
	apiGetDeliveryBindAccount  = "/cgi-bin/express/local/business/shop/get"
	apiGetDeliveryOrder        = "/cgi-bin/express/local/business/order/get"
	apiPreAddDeliveryOrder     = "/cgi-bin/express/local/business/order/pre_add"
	apiPreCancelDeliveryOrder  = "/cgi-bin/express/local/business/order/precancel"
	apiReAddDeliveryOrder      = "/cgi-bin/express/local/business/order/readd"
	apiMockUpdateDeliveryOrder = "/cgi-bin/express/local/business/test_update_order"
	apiUpdateDeliveryOrder     = "/cgi-bin/express/local/delivery/update_order"
)

// AbnormalConfirmer 异常件退回商家商家确认器
type AbnormalConfirmer struct {
	ShopID       string `json:"shopid"`        // 商家id， 由配送公司分配的appkey
	ShopOrderID  string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo       string `json:"shop_no"`       // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
	DeliverySign string `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串
	WaybillID    string `json:"waybill_id"`    // 配送单id
	Remark       string `json:"remark"`        // 备注
}

// Confirm 异常件退回商家商家确认收货
func (cli *Client) AbnormalImmediateDeliveryConfirm(confirmer *AbnormalConfirmer) (*request.CommonResult, error) {
	api := baseURL + apiAbnormalConfirm

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.abnormalImmediateDeliveryConfirm(api, token, confirmer)
}

func (cli *Client) abnormalImmediateDeliveryConfirm(api, token string, confirmer *AbnormalConfirmer) (*request.CommonResult, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonResult)
	if err := cli.request.Post(url, confirmer, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeliveryOrderCreator 下配送单参数
type DeliveryOrderCreator struct {
	DeliveryToken string            `json:"delivery_token,omitempty"` // 预下单接口返回的参数，配送公司可保证在一段时间内运费不变
	ShopID        string            `json:"shopid"`                   // 商家id， 由配送公司分配的appkey
	ShopOrderID   string            `json:"shop_order_id"`            // 唯一标识订单的 ID，由商户生成
	ShopNo        string            `json:"shop_no"`                  // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
	DeliverySign  string            `json:"delivery_sign"`            // 用配送公司提供的appSecret加密的校验串
	DeliveryID    string            `json:"delivery_id"`              // 配送公司ID
	OpenID        string            `json:"openid"`                   // 下单用户的openid
	Sender        DeliveryUser      `json:"sender"`                   // 发件人信息，闪送、顺丰同城急送必须填写，美团配送、达达，若传了shop_no的值可不填该字段
	Receiver      DeliveryUser      `json:"receiver"`                 // 收件人信息
	Cargo         DeliveryCargo     `json:"cargo"`                    // 货物信息
	OrderInfo     DeliveryOrderInfo `json:"order_info"`               // 订单信息
	Shop          DeliveryShop      `json:"shop"`                     // 商品信息，会展示到物流通知消息中
	SubBizID      string            `json:"sub_biz_id"`               // 子商户id，区分小程序内部多个子商户
}

// DeliveryUser 发件人信息，闪送、顺丰同城急送必须填写，美团配送、达达，若传了shop_no的值可不填该字段
type DeliveryUser struct {
	Name           string  `json:"name"`            // 姓名，最长不超过256个字符
	City           string  `json:"city"`            // 城市名称，如广州市
	Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
	AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
	Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
	Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
	Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
	CoordinateType uint8   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
}

// DeliveryCargo 货物信息
type DeliveryCargo struct {
	GoodsValue        float64             `json:"goods_value"`         // 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-5000]
	GoodsHeight       float64             `json:"goods_height"`        // 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-45]
	GoodsLength       float64             `json:"goods_length"`        // 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-65]
	GoodsWidth        float64             `json:"goods_width"`         // 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
	GoodsWeight       float64             `json:"goods_weight"`        // 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
	GoodsDetail       DeliveryGoodsDetail `json:"goods_detail"`        // 货物详情，最长不超过10240个字符
	GoodsPickupInfo   string              `json:"goods_pickup_info"`   // 货物取货信息，用于骑手到店取货，最长不超过100个字符
	GoodsDeliveryInfo string              `json:"goods_delivery_info"` // 货物交付信息，最长不超过100个字符
	CargoFirstClass   string              `json:"cargo_first_class"`   // 品类一级类目
	CargoSecondClass  string              `json:"cargo_second_class"`  // 品类二级类目
}

// DeliveryGoodsDetail 货物详情
type DeliveryGoodsDetail struct {
	Goods []DeliveryGoods `json:"goods"` // 货物交付信息，最长不超过100个字符
}

// DeliveryGoods 货物
type DeliveryGoods struct {
	Count uint    `json:"good_count"` // 货物数量
	Name  string  `json:"good_name"`  // 货品名称
	Price float32 `json:"good_price"` // 货品单价，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）
	Unit  string  `json:"good_unit"`  // 货品单位，最长不超过20个字符
}

// DeliveryOrderInfo 订单信息
type DeliveryOrderInfo struct {
	DeliveryServiceCode  string  `json:"delivery_service_code"`  // 配送服务代码 不同配送公司自定义,微信侧不理解
	OrderType            uint8   `json:"order_type"`             // 订单类型, 0: 即时单 1 预约单，如预约单，需要设置expected_delivery_time或expected_finish_time或expected_pick_time
	ExpectedDeliveryTime uint    `json:"expected_delivery_time"` // 期望派单时间(顺丰同城急送、达达、支持)，unix-timestamp
	ExpectedFinishTime   uint    `json:"expected_finish_time"`   // 期望送达时间(顺丰同城急送、美团配送支持），unix-timestamp
	ExpectedPickTime     uint    `json:"expected_pick_time"`     // 期望取件时间（闪送支持），unix-timestamp
	PoiSeq               string  `json:"poi_seq"`                // 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
	Note                 string  `json:"note"`                   // 备注，最长不超过200个字符
	OrderTime            uint    `json:"order_time"`             // 用户下单付款时间
	IsInsured            uint8   `json:"is_insured"`             // 是否保价，0，非保价，1.保价
	DeclaredValue        float64 `json:"declared_value"`         // 保价金额，单位为元，精确到分
	Tips                 float64 `json:"tips"`                   // 小费，单位为元, 下单一般不加小费
	IsDirectDelivery     uint    `json:"is_direct_delivery"`     // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑手只能配送此订单至完成，配送费用也相应高一些，闪送必须选1，达达可选0或1，其余配送公司不支持直拿直送）
	CashOnDelivery       uint    `json:"cash_on_delivery"`       // 骑手应付金额，单位为元，精确到分
	CashOnPickup         uint    `json:"cash_on_pickup"`         // 骑手应收金额，单位为元，精确到分
	RiderPickMethod      uint8   `json:"rider_pick_method"`      // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
	IsFinishCodeNeeded   uint8   `json:"is_finish_code_needed"`  // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
	IsPickupCodeNeeded   uint8   `json:"is_pickup_code_needed"`  // 取货码（0：不需要；1：需要。取货码的作用是：骑手必须输入取货码才能从商家取货）
}

// DeliveryShop 商品信息，会展示到物流通知消息中
type DeliveryShop struct {
	WxaPath    string `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
	ImgURL     string `json:"img_url"`     // 商品缩略图 url
	GoodsName  string `json:"goods_name"`  // 商品名称
	GoodsCount uint   `json:"goods_count"` // 商品数量
}

// PreDeliveryOrderResponse 返回数据
type PreDeliveryOrderResponse struct {
	Fee              float64 `json:"fee"`               // 实际运费(单位：元)，运费减去优惠券费用
	Deliverfee       float64 `json:"deliverfee"`        // 运费(单位：元)
	Couponfee        float64 `json:"couponfee"`         // 优惠券费用(单位：元)
	Tips             float64 `json:"tips"`              // 小费(单位：元)
	Insurancefee     float64 `json:"insurancefee"`      // 保价费(单位：元)
	Distance         float64 `json:"distance"`          // 配送距离(单位：米)
	DispatchDuration uint    `json:"dispatch_duration"` // 预计骑手接单时间，单位秒，比如5分钟，就填300, 无法预计填0
	DeliveryToken    string  `json:"delivery_token"`    // 配送公司可以返回此字段，当用户下单时候带上这个字段，保证在一段时间内运费不变
}

// Prepare 预下配送单接口
func (cli *Client) PreAddImmediateDeliveryOrder(creator *DeliveryOrderCreator) (*PreDeliveryOrderResponse, error) {
	api := baseURL + apiPreCancelDeliveryOrder

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.preAddImmediateDeliveryOrder(api, token, creator)
}

func (cli *Client) preAddImmediateDeliveryOrder(api, token string, creator *DeliveryOrderCreator) (*PreDeliveryOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(PreDeliveryOrderResponse)
	if err := cli.request.Post(url, creator, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateDeliveryOrderResponse 返回数据
type CreateDeliveryOrderResponse struct {
	*request.CommonResult
	Fee              uint    `json:"fee"`               //实际运费(单位：元)，运费减去优惠券费用
	Deliverfee       uint    `json:"deliverfee"`        //运费(单位：元)
	Couponfee        uint    `json:"couponfee"`         //优惠券费用(单位：元)
	Tips             uint    `json:"tips"`              //小费(单位：元)
	Insurancefee     uint    `json:"insurancefee"`      //保价费(单位：元)
	Distance         float64 `json:"distance"`          //	配送距离(单位：米)
	WaybillID        string  `json:"waybill_id"`        //配送单号
	OrderStatus      int     `json:"order_status"`      //配送状态
	FinishCode       uint    `json:"finish_code"`       //	收货码
	PickupCode       uint    `json:"pickup_code"`       //取货码
	DispatchDuration uint    `json:"dispatch_duration"` //	预计骑手接单时间，单位秒，比如5分钟，就填300, 无法预计填0
}

// Create 下配送单
func (cli *Client) AddImmediateDeliveryOrder(creator *DeliveryOrderCreator) (*CreateDeliveryOrderResponse, error) {
	api := baseURL + apiAddDeliveryOrder

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.addImmediateDeliveryOrder(api, token, creator)
}

func (cli *Client) addImmediateDeliveryOrder(api, token string, creator *DeliveryOrderCreator) (*CreateDeliveryOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CreateDeliveryOrderResponse)
	if err := cli.request.Post(url, creator, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Recreate 重新下单
func (cli *Client) ReImmediateDeliveryOrder(creator *DeliveryOrderCreator) (*CreateDeliveryOrderResponse, error) {
	api := baseURL + apiReAddDeliveryOrder

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.reImmediateDeliveryOrder(api, token, creator)
}

func (cli *Client) reImmediateDeliveryOrder(api, token string, creator *DeliveryOrderCreator) (*CreateDeliveryOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CreateDeliveryOrderResponse)
	if err := cli.request.Post(url, creator, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeliveryTipAdder 增加小费参数
type DeliveryTipAdder struct {
	ShopID       string  `json:"shopid"`        // 商家id， 由配送公司分配的appkey
	ShopOrderID  string  `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo       string  `json:"shop_no"`       // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
	DeliverySign string  `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串
	WaybillID    string  `json:"waybill_id"`    // 配送单id
	OpenID       string  `json:"openid"`        // 下单用户的openid
	Tips         float64 `json:"tips"`          // 小费金额(单位：元) 各家配送公司最大值不同
	Remark       string  `json:"Remark"`        // 备注
}

// Add 对待接单状态的订单增加小费。需要注意：订单的小费，以最新一次加小费动作的金额为准，故下一次增加小费额必须大于上一次小费额
func (cli *Client) AddImmediateDeliveryTip(adder *DeliveryTipAdder) (*request.CommonResult, error) {
	api := baseURL + apiAddDeliveryTip
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.addImmediateDeliveryTip(api, token, adder)
}

func (cli *Client) addImmediateDeliveryTip(api, token string, adder *DeliveryTipAdder) (*request.CommonResult, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonResult)
	if err := cli.request.Post(url, adder, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeliveryOrderCanceler 取消配送单参数
type DeliveryOrderCanceler struct {
	ShopID       string `json:"shopid"`           // 商家id， 由配送公司分配的appkey
	ShopOrderID  string `json:"shop_order_id"`    // 唯一标识订单的 ID，由商户生成
	ShopNo       string `json:"shop_no"`          // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
	DeliverySign string `json:"delivery_sign"`    // 用配送公司提供的appSecret加密的校验串
	DeliveryID   string `json:"delivery_id"`      // 快递公司ID
	WaybillID    string `json:"waybill_id"`       // 配送单id
	ReasonID     uint8  `json:"cancel_reason_id"` // 取消原因Id
	Reason       string `json:"cancel_reason"`    // 取消原因
}

// CancelDeliveryOrderResponse 取消配送单返回数据
type CancelDeliveryOrderResponse struct {
	*request.CommonResult
	DeductFee float64 `json:"deduct_fee"` // 	预计扣除的违约金(单位：元)，精确到分
	Desc      string  `json:"desc"`       //说明
}

// Prepare 预取消配送单
func (cli *Client) PreCancelImmediateDeliveryOrder(canceler *DeliveryOrderCanceler) (*CancelDeliveryOrderResponse, error) {
	api := baseURL + apiCancelDeliveryOrder
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}
	return cli.preCancelImmediateDeliveryOrder(api, token, canceler)
}

func (cli *Client) preCancelImmediateDeliveryOrder(api, token string, canceler *DeliveryOrderCanceler) (*CancelDeliveryOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CancelDeliveryOrderResponse)
	if err := cli.request.Post(url, canceler, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Cancel 取消配送单
func (cli *Client) CancelImmediateDeliveryOrder(canceler *DeliveryOrderCanceler) (*CancelDeliveryOrderResponse, error) {
	api := baseURL + apiCancelDeliveryOrder

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.cancelImmediateDeliveryOrder(api, token, canceler)
}

func (cli *Client) cancelImmediateDeliveryOrder(api, token string, canceler *DeliveryOrderCanceler) (*CancelDeliveryOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CancelDeliveryOrderResponse)
	if err := cli.request.Post(url, canceler, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAllImmediateDeliveryResponse 获取已支持的配送公司列表接口返回数据
type GetAllImmediateDeliveryResponse struct {
	*request.CommonResult
	List []struct {
		ID   string `json:"delivery_id"`   //配送公司Id
		Name string `json:"delivery_name"` //	配送公司名称
	} `json:"list"` // 配送公司列表
}

// GetAllImmediateDelivery 获取已支持的配送公司列表接口
func (cli *Client) GetAllImmediateDelivery() (*GetAllImmediateDeliveryResponse, error) {
	api := baseURL + apiGetAllImmediateDelivery

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getAllImmediateDelivery(api, token)
}

func (cli *Client) getAllImmediateDelivery(api, token string) (*GetAllImmediateDeliveryResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetAllImmediateDeliveryResponse)
	if err := cli.request.Post(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetBindAccountResponse 返回数据
type GetBindAccountResponse struct {
	*request.CommonResult
	ShopList []struct {
		DeliveryID  string `json:"delivery_id"`  // 配送公司Id
		ShopID      string `json:"shopid"`       // 商家id
		AuditResult uint8  `json:"audit_result"` // 审核状态
	} `json:"shop_list"` // 配送公司列表
}

// GetBindAccount 拉取已绑定账号
func (cli *Client) GetImmediateDeliveryBindAccount() (*GetBindAccountResponse, error) {
	api := baseURL + apiGetDeliveryBindAccount

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getImmediateDeliveryBindAccount(api, token)
}

func (cli *Client) getImmediateDeliveryBindAccount(api, token string) (*GetBindAccountResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetBindAccountResponse)
	if err := cli.request.Post(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeliveryOrderGetter 请求参数
type DeliveryOrderGetter struct {
	ShopID       string `json:"shopid"`        // 商家id， 由配送公司分配的appkey
	ShopOrderID  string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo       string `json:"shop_no"`       // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
	DeliverySign string `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串说明
}

// GetDeliveryOrderResponse 返回数据
type GetDeliveryOrderResponse struct {
	*request.CommonResult
	OrderStatus int     `json:"order_status"` // 	配送状态，枚举值
	WaybillID   string  `json:"waybill_id"`   // 配送单号
	RiderName   string  `json:"rider_name"`   // 骑手姓名
	RiderPhone  string  `json:"rider_phone"`  // 骑手电话
	RiderLng    float64 `json:"rider_lng"`    // 骑手位置经度, 配送中时返回
	RiderLat    float64 `json:"rider_lat"`    // 骑手位置纬度, 配送中时返回
}

// Get 下配送单
func (cli *Client) GetImmediateDeliveryOrder(getter *DeliveryOrderGetter) (*GetDeliveryOrderResponse, error) {
	api := baseURL + apiGetDeliveryOrder

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getImmediateDeliveryOrder(api, token, getter)
}

func (cli *Client) getImmediateDeliveryOrder(api, token string, getter *DeliveryOrderGetter) (*GetDeliveryOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetDeliveryOrderResponse)
	if err := cli.request.Post(url, getter, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateDeliveryOrderMocker 请求参数
type UpdateDeliveryOrderMocker struct {
	ShopID      string `json:"shopid"`        // 商家id, 必须是 "test_shop_id"
	ShopOrderID string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ActionTime  uint   `json:"action_time"`   // 状态变更时间点，Unix秒级时间戳
	OrderStatus int    `json:"order_status"`  // 配送状态，枚举值
	ActionMsg   string `json:"action_msg"`    // 附加信息
}

// Mock 模拟配送公司更新配送单状态
func (cli *Client) MockUpdateImmediateDeliveryOrder(mocker *UpdateDeliveryOrderMocker) (*request.CommonResult, error) {
	api := baseURL + apiMockUpdateDeliveryOrder

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.mockUpdateImmediateDeliveryOrder(api, token, mocker)
}

func (cli *Client) mockUpdateImmediateDeliveryOrder(api, token string, mocker *UpdateDeliveryOrderMocker) (*request.CommonResult, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonResult)
	if err := cli.request.Post(url, mocker, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeliveryOrderUpdater 请求参数
type DeliveryOrderUpdater struct {
	WXToken              string        `json:"wx_token"`                         // 下单事件中推送的wx_token字段
	ShopID               string        `json:"shopid"`                           // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID          string        `json:"shop_order_id"`                    // 唯一标识订单的 ID，由商户生成
	ShopNo               string        `json:"shop_no,omitempty"`                // 商家门店编号， 在配送公司侧登记
	WaybillID            string        `json:"waybill_id"`                       // 配送单id
	ActionTime           uint          `json:"action_time"`                      // 状态变更时间点，Unix秒级时间戳
	OrderStatus          int           `json:"order_status"`                     // 订单状态，枚举值，下附枚举值列表及说明
	ActionMsg            string        `json:"action_msg,omitempty"`             // 附加信息
	WxaPath              string        `json:"wxa_path"`                         // 配送公司小程序跳转路径，用于用户收到消息会间接跳转到这个页面
	Agent                DeliveryAgent `json:"agent,omitempty"`                  // 骑手信息, 骑手接单时需返回
	ExpectedDeliveryTime uint          `json:"expected_delivery_time,omitempty"` // 预计送达时间戳， 骑手接单时需返回
}

// DeliveryAgent 骑手信息
type DeliveryAgent struct {
	Name      string `json:"name"`                         // 骑手姓名
	Phone     string `json:"phone"`                        // 骑手电话
	Encrypted uint8  `json:"is_phone_encrypted,omitempty"` // 电话是否加密
}

// Update 模拟配送公司更新配送单状态
func (cli *Client) UpdateImmediateDeliveryOrder(updater *DeliveryOrderUpdater) (*request.CommonResult, error) {
	api := baseURL + apiUpdateDeliveryOrder

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.updateImmediateDeliveryOrder(api, token, updater)
}

func (cli *Client) updateImmediateDeliveryOrder(api, token string, updater *DeliveryOrderUpdater) (*request.CommonResult, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonResult)
	if err := cli.request.Post(url, updater, res); err != nil {
		return nil, err
	}

	return res, nil
}
