package weapp

// ExpressOrder 物流订单
type ExpressOrder struct {
	OrderID      string           `json:"order_id"`                // 订单ID，须保证全局唯一，不超过512字节
	OpenID       string           `json:"openid,omitempty"`        //	用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID   string           `json:"delivery_id"`             // 快递公司ID，参见getAllDelivery
	BizID        string           `json:"biz_id"`                  // 快递客户编码或者现付编码
	CustomRemark string           `json:"custom_remark,omitempty"` //	快递备注信息，比如"易碎物品"，不超过1024字节
	Sender       ExpreseeUserInfo `json:"sender"`                  // 发件人信息
	Receiver     ExpreseeUserInfo `json:"receiver"`                // 收件人信息
	Cargo        ExpressCargo     `json:"cargo"`                   // 包裹信息，将传递给快递公司
	Shop         ExpressShop      `json:"shop,omitempty"`          //	商家信息，会展示到物流服务通知中，当add_source=2时无需填写（不发送物流服务通知）
	Insured      ExpressInsure    `json:"insured"`                 // 保价信息
	Service      ExpressService   `json:"service"`                 // 服务类型
}

// ExpreseeUserInfo 收件人/发件人信息
type ExpreseeUserInfo struct {
	Name     string `json:"name"`                // 收件人/发件人姓名，不超过64字节
	Tel      string `json:"tel,omitempty"`       // 收件人/发件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Mobile   string `json:"mobile,omitempty"`    // 收件人/发件人手机号码，若不填写则必须填写 tel，不超过32字节
	Company  string `json:"company,omitempty"`   // 收件人/发件人公司名称，不超过64字节
	PostCode string `json:"post_code,omitempty"` // 收件人/发件人邮编，不超过10字节
	Country  string `json:"country,omitempty"`   // 收件人/发件人国家，不超过64字节
	Province string `json:"province"`            // 收件人/发件人省份，比如："广东省"，不超过64字节
	City     string `json:"city"`                // 收件人/发件人市/地区，比如："广州市"，不超过64字节
	Area     string `json:"area"`                // 收件人/发件人区/县，比如："海珠区"，不超过64字节
	Address  string `json:"address"`             // 收件人/发件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
}

// ExpressCargo 包裹信息
type ExpressCargo struct {
	Count      uint          `json:"count"`       // 包裹数量
	Weight     float64       `json:"weight"`      // 包裹总重量，单位是千克(kg)
	SpaceX     float64       `json:"space_x"`     // 包裹长度，单位厘米(cm)
	SpaceY     float64       `json:"space_y"`     // 包裹宽度，单位厘米(cm)
	SpaceZ     float64       `json:"space_z"`     // 包裹高度，单位厘米(cm)
	DetailList []CargoDetail `json:"detail_list"` // 包裹中商品详情列表
}

// CargoDetail 包裹详情
type CargoDetail struct {
	Name  string `json:"name"`  // 商品名，不超过128字节
	Count uint   `json:"count"` // 商品数量
}

// ExpressShop 商家信息
type ExpressShop struct {
	WXAPath    string `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
	IMGUrl     string `json:"img_url"`     // 商品缩略图 url
	GoodsName  string `json:"goods_name"`  // 商品名称
	GoodsCount uint   `json:"goods_count"` // 商品数量
}

// ExpressInsure 订单保价
type ExpressInsure struct {
	Used  InsureStatus `json:"use_insured"`   // 是否保价，0 表示不保价，1 表示保价
	Value uint         `json:"insured_value"` // 保价金额，单位是分，比如: 10000 表示 100 元
}

// ExpressService 服务类型
type ExpressService struct {
	Type uint8  `json:"service_type"` // 服务类型ID
	Name string `json:"service_name"` // 服务名称
}
