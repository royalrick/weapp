package express

import (
	"github.com/medivhzhan/weapp"
)

const (
	apiBindAccount    = "/cgi-bin/express/business/account/bind"
	apiGetAllAccount  = "/cgi-bin/express/business/account/getall"
	apiGetPath        = "/cgi-bin/express/business/path/get"
	apiAddOrder       = "/cgi-bin/express/business/order/add"
	apiCancelOrder    = "/cgi-bin/express/business/order/cancel"
	apiGetAllDelivery = "/cgi-bin/express/business/delivery/getall"
	apiGetOrder       = "/cgi-bin/express/business/order/get"
	apiGetPrinter     = "/cgi-bin/express/business/printer/getall"
	apiGetQuota       = "/cgi-bin/express/business/quota/get"
	apiUpdatePrinter  = "/cgi-bin/express/business/printer/update"
)

// Account 物流账号
type Account struct {
	Type          BindType `json:"type"`           // bind表示绑定，unbind表示解除绑定
	BizID         string   `json:"biz_id"`         // 快递公司客户编码
	DeliveryID    string   `json:"delivery_id"`    // 快递公司 ID
	Password      string   `json:"password"`       // 快递公司客户密码
	RemarkContent string   `json:"remark_content"` // 备注内容（提交EMS审核需要）
}

// BindType 绑定动作类型
type BindType = string

// 所有绑定动作类型
const (
	Bind   = "bind"   // 绑定
	Unbind = "unbind" // 解除绑定
)

// Bind 绑定、解绑物流账号
// @accessToken 接口调用凭证
func (ac *Account) Bind(accessToken string) (*weapp.BaseResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiBindAccount, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(weapp.BaseResponse)
	if err := weapp.PostJSON(api, ac, res); err != nil {
		return nil, err
	}

	return res, nil
}

// AccountList 所有绑定的物流账号
type AccountList struct {
	weapp.BaseResponse
	Count uint `json:"count"` // 账号数量
	List  []struct {
		BizID           string     `json:"biz_id"`            // 	快递公司客户编码
		DeliveryID      string     `json:"delivery_id"`       // 	快递公司 ID
		CreateTime      uint       `json:"create_time"`       // 	账号绑定时间
		UpdateTime      uint       `json:"update_time"`       // 	账号更新时间
		StatusCode      BindStatus `json:"status_code"`       // 	绑定状态
		Alias           string     `json:"alias"`             // 	账号别名
		RemarkWrongMsg  string     `json:"remark_wrong_msg"`  // 	账号绑定失败的错误信息（EMS审核结果）
		RemarkContent   string     `json:"remark_content"`    // 	账号绑定时的备注内容（提交EMS审核需要)）
		QuotaNum        uint       `json:"quota_num"`         // 	电子面单余额
		QuotaUpdateTime uint       `json:"quota_update_time"` // 	电子面单余额更新时间

	} `json:"list"` // 账号列表
}

// BindStatus 账号绑定状态
type BindStatus = int8

// 所有账号绑定状态
const (
	BindSuccess = 0  // 成功
	BindFailed  = -1 // 系统失败
)

// GetAllAccount 获取所有绑定的物流账号
// @accessToken 接口调用凭证
func GetAllAccount(accessToken string) (*AccountList, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetAllAccount, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(AccountList)
	if err := weapp.PostJSON(api, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PathGetter 查询运单轨迹所需参数
type PathGetter struct {
	OrderID    string `json:"order_id"`         //	订单 ID，需保证全局唯一
	OpenID     string `json:"openid,omitempty"` //	用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID string `json:"delivery_id"`      //	快递公司ID，参见getAllDelivery
	WaybillID  string `json:"waybill_id"`       //	运单ID
}

// Path 运单轨迹
type Path struct {
	weapp.BaseResponse
	OpenID       string     `json:"openid"`         // 用户openid
	DeliveryID   string     `json:"delivery_id"`    // 快递公司 ID
	WaybillID    string     `json:"waybill_id"`     // 运单 ID
	PathItemNum  uint       `json:"path_item_num"`  // 轨迹节点数量
	PathItemList []PathNode `json:"path_item_list"` // 轨迹节点列表
}

// PathNode 运单轨迹节点
type PathNode struct {
	ActionTime uint   `json:"action_time"` // 轨迹节点 Unix 时间戳
	ActionType uint   `json:"action_type"` // 轨迹节点类型
	ActionMsg  string `json:"action_msg"`  // 轨迹节点详情
}

// Get 查询运单轨迹
// @accessToken 接口调用凭证
func (pg *PathGetter) Get(accessToken string) (*Path, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetPath, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(Path)
	if err := weapp.PostJSON(api, pg, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Order 订单
type Order struct {
	OrderID      string        `json:"order_id"`                // 订单ID，须保证全局唯一，不超过512字节
	OpenID       string        `json:"openid,omitempty"`        //	用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID   string        `json:"delivery_id"`             // 快递公司ID，参见getAllDelivery
	BizID        string        `json:"biz_id"`                  // 快递客户编码或者现付编码
	CustomRemark string        `json:"custom_remark,omitempty"` //	快递备注信息，比如"易碎物品"，不超过1024字节
	Sender       OrderSender   `json:"sender"`                  // 发件人信息
	Receiver     OrderReceiver `json:"receiver"`                // 收件人信息
	Cargo        OrderCargo    `json:"cargo"`                   // 包裹信息，将传递给快递公司
	Shop         OrderShop     `json:"shop,omitempty"`          //	商家信息，会展示到物流服务通知中，当add_source=2时无需填写（不发送物流服务通知）
	Insured      OrderInsure   `json:"insured"`                 // 保价信息
	Service      OrderService  `json:"service"`                 // 服务类型
	ExpectTime   uint          `json:"expect_time,omitempty"`   //	顺丰必须填写此字段。预期的上门揽件时间，0表示已事先约定取件时间；否则请传预期揽件时间戳，需大于当前时间，收件员会在预期时间附近上门。例如expect_time为“1557989929”，表示希望收件员将在2019年05月16日14:58:49-15:58:49内上门取货。

}

// OrderSource 订单来源
type OrderSource = int8

// 所有订单来源
const (
	FromWeapp   OrderSource = iota // 小程序订单
	FromAppOrH5                    // APP或H5订单
)

// OrderCreator 订单创建器
type OrderCreator struct {
	Order
	AddSource OrderSource `json:"add_source"` // 订单来源，0为小程序订单，2为App或H5订单，填2则不发送物流服务通知
	WXAppID   string      `json:"wx_appid"`   // App或H5的appid，add_source=2时必填，需和开通了物流助手的小程序绑定同一open帐号
}

// OrderSender 发件人信息
type OrderSender struct {
	Name     string `json:"name"`                // 发件人姓名，不超过64字节
	Tel      string `json:"tel,omitempty"`       // 发件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Mobile   string `json:"mobile,omitempty"`    // 发件人手机号码，若不填写则必须填写 tel，不超过32字节
	Company  string `json:"company,omitempty"`   // 发件人公司名称，不超过64字节
	PostCode string `json:"post_code,omitempty"` // 发件人邮编，不超过10字节
	Country  string `json:"country,omitempty"`   // 发件人国家，不超过64字节
	Province string `json:"province"`            // 发件人省份，比如："广东省"，不超过64字节
	City     string `json:"city"`                // 发件人市/地区，比如："广州市"，不超过64字节
	Area     string `json:"area"`                // 发件人区/县，比如："海珠区"，不超过64字节
	Address  string `json:"address"`             // 发件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
}

// OrderReceiver 收件人信息
type OrderReceiver struct {
	Name     string `json:"name"`                // 收件人姓名，不超过64字节
	Tel      string `json:"tel,omitempty"`       // 收件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Mobile   string `json:"mobile,omitempty"`    // 收件人手机号码，若不填写则必须填写 tel，不超过32字节
	Company  string `json:"company,omitempty"`   // 收件人公司名，不超过64字节
	PostCode string `json:"post_code,omitempty"` // 收件人邮编，不超过10字节
	Country  string `json:"country,omitempty"`   // 收件人所在国家，不超过64字节
	Province string `json:"province"`            // 收件人省份，比如："广东省"，不超过64字节
	City     string `json:"city"`                // 收件人地区/市，比如："广州市"，不超过64字节
	Area     string `json:"area"`                // 收件人区/县，比如："天河区"，不超过64字节
	Address  string `json:"address"`             // 收件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
}

// OrderCargo 包裹信息
type OrderCargo struct {
	Count      uint          `json:"count"`       // 包裹数量
	Height     uint          `json:"weight"`      // 包裹总重量，单位是千克(kg)
	SpaceX     uint          `json:"space_x"`     // 包裹长度，单位厘米(cm)
	SpaceY     uint          `json:"space_y"`     // 包裹宽度，单位厘米(cm)
	SpaceZ     uint          `json:"space_z"`     // 包裹高度，单位厘米(cm)
	DetailList []CargoDetail `json:"detail_list"` // 包裹中商品详情列表
}

// OrderShop 商家信息
type OrderShop struct {
	WXAPath    string `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
	IMGUrl     string `json:"img_url"`     // 商品缩略图 url
	GoodsName  string `json:"goods_name"`  // 商品名称
	GoodsCount uint   `json:"goods_count"` // 商品数量
}

// CargoDetail 包裹详情
type CargoDetail struct {
	Name  string `json:"name"`  // 商品名，不超过128字节
	Count uint   `json:"count"` // 商品数量
}

// OrderInsure 订单保价
type OrderInsure struct {
	UseInsured   InsureStatus // 是否保价，0 表示不保价，1 表示保价
	InsuredValue uint         // 保价金额，单位是分，比如: 10000 表示 100 元
}

// OrderService 服务类型
type OrderService struct {
	Type int8   `json:"service_type"` // 服务类型ID
	Name string `json:"service_name"` // 服务名称
}

// InsureStatus 保价状态
type InsureStatus = int8

// 所有保价状态
const (
	Uninsured = 0 // 不保价
	Insured   = 1 // 保价
)

// AddOrderResponse 创建订单返回数据
type AddOrderResponse struct {
	weapp.BaseResponse
	OrderID     string `json:"order_id"`   //	订单ID，下单成功时返回
	WaybillID   string `json:"waybill_id"` //	运单ID，下单成功时返回
	WaybillData []struct {
		Key   string `json:"key"`   // 运单信息 key
		Value string `json:"value"` // 运单信息 value
	} `json:"waybill_data"` //	运单信息，下单成功时返回
	DeliveryResultcode int    `json:"delivery_resultcode"` //	快递侧错误码，下单失败时返回
	DeliveryResultmsg  string `json:"delivery_resultmsg"`  //	快递侧错误信息，下单失败时返回
}

// Add 生成运单
// @accessToken 接口调用凭证
func (oc *OrderCreator) Add(accessToken string) (*AddOrderResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetPath, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(AddOrderResponse)
	if err := weapp.PostJSON(api, oc, res); err != nil {
		return nil, err
	}

	return res, nil
}

// OrderCanceler din
type OrderCanceler struct {
	OrderID    string `json:"order_id"`         // 订单 ID，需保证全局唯一
	OpenID     string `json:"openid,omitempty"` // 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID string `json:"delivery_id"`      // 快递公司ID，参见getAllDelivery
	WaybillID  string `json:"waybill_id"`       // 运单ID

}

// CancelOrderResponse 取消订单返回数据
type CancelOrderResponse struct {
	weapp.BaseResponse
	Count uint `json:"count"` //快递公司数量
	Data  []struct {
		DeliveryID   string `json:"delivery_id"`   // 快递公司 ID
		DeliveryName string `json:"delivery_name"` // 快递公司名称

	} `json:"data"` //快递公司信息列表
}

// Cancel 取消运单
// @accessToken 接口调用凭证
func (oc *OrderCanceler) Cancel(accessToken string) (*weapp.BaseResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiCancelOrder, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(weapp.BaseResponse)
	if err := weapp.PostJSON(api, oc, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeliveryList 支持的快递公司列表
type DeliveryList struct {
	weapp.BaseResponse
	Count uint `json:"count"` // 快递公司数量
	Data  []struct {
		ID   string `json:"delivery_id"`   // 快递公司 ID
		Name string `json:"delivery_name"` // 快递公司名称
	} `json:"data"` // 快递公司信息列表
}

// GetAllDelivery 获取支持的快递公司列表
// @accessToken 接口调用凭证
func GetAllDelivery(accessToken string) (*DeliveryList, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetAllDelivery, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(DeliveryList)
	if err := weapp.PostJSON(api, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// OrderGetter 订单获取器
type OrderGetter struct {
	OrderID    string `json:"order_id"`         // 订单 ID，需保证全局唯一
	OpenID     string `json:"openid,omitempty"` // 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID string `json:"delivery_id"`      // 快递公司ID，参见getAllDelivery
	WaybillID  string `json:"waybill_id"`       // 运单ID

}

// GetOrderResponse 获取运单返回数据
type GetOrderResponse struct {
	weapp.BaseResponse
	PrintHTML   string `json:"print_html"` // 运单 html 的 BASE64 结果
	WaybillData []struct {
		Key   string `json:"key"`   // 运单信息 key
		Value string `json:"value"` // 运单信息 value

	} `json:"waybill_data"` // 运单信息
}

// Get 获取运单数据
// @accessToken 接口调用凭证
func (og *OrderGetter) Get(accessToken string) (*GetOrderResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetOrder, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(GetOrderResponse)
	if err := weapp.PostJSON(api, og, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPrinterResponse 获取打印员返回数据
type GetPrinterResponse struct {
	weapp.BaseResponse
	Count  uint     `json:"count"`  // 已经绑定的打印员数量
	OpenID []string `json:"openid"` // 打印员 openid 列表

}

// GetPrinter 获取打印员。若需要使用微信打单 PC 软件，才需要调用。
// @accessToken 接口调用凭证
func GetPrinter(accessToken string) (*GetPrinterResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetPrinter, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(GetPrinterResponse)
	if err := weapp.PostJSON(api, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// QuotaGetter 电子面单余额获取器
type QuotaGetter struct {
	DeliveryID string `json:"delivery_id"` // 快递公司ID，参见getAllDelivery
	BizID      string `json:"biz_id"`      // 快递公司客户编码

}

// Quota 电子面单余额
type Quota struct {
	weapp.BaseResponse
	Number uint // 电子面单余额
}

// GetQuota 获取电子面单余额。仅在使用加盟类快递公司时，才可以调用。
func (qg *QuotaGetter) GetQuota(accessToken string) (*Quota, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetQuota, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(Quota)
	if err := weapp.PostJSON(api, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PrinterUpdater 打印员更新器
type PrinterUpdater struct {
	OpenID string   `json:"openid"`      // 打印员 openid
	Type   BindType `json:"update_type"` // 更新类型
}

// Update 更新打印员。若需要使用微信打单 PC 软件，才需要调用。
func (pu *PrinterUpdater) Update(accessToken string) (*weapp.BaseResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiUpdatePrinter, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(weapp.BaseResponse)
	if err := weapp.PostJSON(api, pu, res); err != nil {
		return nil, err
	}

	return res, nil
}
