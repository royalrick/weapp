package server

import "encoding/xml"

// EncryptedResult 接收的加密数据
type EncryptedResult struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	ToUserName string   `json:"ToUserName" xml:"ToUserName"` // 接收者 为小程序 AppID
	Encrypt    string   `json:"Encrypt" xml:"Encrypt"`       // 加密消息
}

// EncryptedMsgRequest 发送的加密消息格式
type EncryptedMsgRequest struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `json:"Encrypt" xml:"Encrypt"`                               // 加密消息
	TimeStamp    string   `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`       // 时间戳
	Nonce        string   `json:"Nonce,omitempty" xml:"Nonce,omitempty"`               // 随机数
	MsgSignature string   `json:"MsgSignature,omitempty" xml:"MsgSignature,omitempty"` // 签名
}

// CommonServerResult 基础通知数据
type CommonServerResult struct {
	XMLName      xml.Name  `xml:"xml" json:"-"`
	ToUserName   string    `json:"ToUserName" xml:"ToUserName"`     // 小程序的原始ID
	FromUserName string    `json:"FromUserName" xml:"FromUserName"` // 发送者的 openID | 平台推送服务UserName
	CreateTime   uint      `json:"CreateTime" xml:"CreateTime"`     // 消息创建时间(整型）
	MsgType      MsgType   `json:"MsgType" xml:"MsgType"`           // 消息类型
	Event        EventType `json:"Event" xml:"Event"`               // 事件类型
}

// CommonServerReturn 没收到通知后返回的基础数据
type CommonServerReturn struct {
	ToUserName   string `json:"ToUserName" xml:"ToUserName"`     // 是 原样返回请求中的 FromUserName
	FromUserName string `json:"FromUserName" xml:"FromUserName"` // 是 快递公司小程序 UserName
	CreateTime   uint   `json:"CreateTime" xml:"CreateTime"`     // 是 事件时间，Unix时间戳
	MsgType      string `json:"MsgType" xml:"MsgType"`           // 是 消息类型，固定为 event
	Event        string `json:"Event" xml:"Event"`               // 是 事件类型，固定为 transport_add_order，不区分大小写
	ResultCode   int    `json:"resultcode" xml:"resultcode"`     // 是 错误码
	ResultMsg    string `json:"resultmsg" xml:"resultmsg"`       // 是 错误描述
}

// UserTempsessionEnterResult 接收的文本消息
type UserTempsessionEnterResult struct {
	CommonServerResult
	SessionFrom string `json:"SessionFrom" xml:"SessionFrom"` // 开发者在客服会话按钮设置的 session-from 属性
}

// TextMessageResult 接收的文本消息
type TextMessageResult struct {
	CommonServerResult
	MsgID   int    `json:"MsgId" xml:"MsgId"` // 消息 ID
	Content string `json:"Content" xml:"Content"`
}

// ImageMessageResult 接收的图片消息
type ImageMessageResult struct {
	CommonServerResult
	MsgID   int    `json:"MsgId" xml:"MsgId"` // 消息 ID
	PicURL  string `json:"PicUrl" xml:"PicUrl"`
	MediaID string `json:"MediaId" xml:"MediaId"`
}

// CardMessageResult 接收的卡片消息
type CardMessageResult struct {
	CommonServerResult
	MsgID        int    `json:"MsgId" xml:"MsgId"`               // 消息 ID
	Title        string `json:"Title" xml:"Title"`               // 标题
	AppID        string `json:"AppId" xml:"AppId"`               // 小程序 appid
	PagePath     string `json:"PagePath" xml:"PagePath"`         // 小程序页面路径
	ThumbURL     string `json:"ThumbUrl" xml:"ThumbUrl"`         // 封面图片的临时cdn链接
	ThumbMediaID string `json:"ThumbMediaId" xml:"ThumbMediaId"` // 封面图片的临时素材id
}

// MediaCheckAsyncResult 异步校验的图片/音频结果
type MediaCheckAsyncResult struct {
	CommonServerResult
	AppID   string `json:"appid" xml:"appid"`       // 小程序的appid
	TraceID string `json:"trace_id" xml:"trace_id"` // 任务id
	Version string `json:"version" xml:"version"`   // 可用于区分接口版本
	// 综合结果
	Result struct {
		// 建议，有risky、pass、review三种值
		Suggest string `json:"suggest" xml:"suggest"`
		// 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
		Label string `json:"label" xml:"label"`
	} `json:"result" xml:"result"`
	// 详细检测结果
	Detail []struct {
		// 策略类型
		Strategy string `json:"strategy" xml:"strategy"`
		// 错误码，仅当该值为0时，该项结果有效
		Errcode int `json:"errcode" xml:"errcode"`
		// 建议，有risky、pass、review三种值
		Suggest string `json:"suggest" xml:"suggest"`
		// 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
		Label int `json:"label" xml:"label"`
		// 0-100，代表置信度，越高代表越有可能属于当前返回的标签（label）
		Prob int `json:"prob" xml:"prob"`
	} `json:"detail" xml:"detail"`
}

// AddNearbyPoiResult 附近小程序添加地点审核状态通知数据
type AddNearbyPoiResult struct {
	CommonServerResult
	AuditID uint   `xml:"audit_id"` // 审核单id
	Status  uint8  `xml:"status"`   // 审核状态（3：审核通过，2：审核失败）
	Reason  string `xml:"reason"`   // 如果status为2，会返回审核失败的原因
	PoiID   uint   `xml:"poi_id"`
}

// ExpressPathUpdateResult 运单轨迹更新事件需要返回的数据
type ExpressPathUpdateResult struct {
	CommonServerResult
	DeliveryID string `json:"DeliveryID" xml:"DeliveryID"` // 快递公司ID
	OrderID    string `json:"OrderID" xml:"OrderID"`       // 	传入的唯一标识订单的 ID，由商户生成，原样返回
	WayBillID  string `json:"WayBillId" xml:"WayBillId"`   // 运单ID
	Version    uint   `json:"Version" xml:"Version"`       // 轨迹版本号（整型）
	Count      uint   `json:"Count" xml:"Count"`           // 轨迹节点数（整型）
	Actions    []struct {
		ActionTime uint   `json:"ActionTime" xml:"ActionTime"` // 轨迹节点 Unix 时间戳
		ActionType uint   `json:"ActionType" xml:"ActionType"` // 轨迹节点类型
		ActionMsg  string `json:"ActionMsg" xml:"ActionMsg"`   // 轨迹节点详情
	} `json:"Actions" xml:"Actions"` // 轨迹列表
}

// AddExpressOrderReturn 请求下单事件需要返回的数据
type AddExpressOrderReturn struct {
	CommonServerReturn
	Token       string `json:"Token" xml:"Token"`             // 	传入的 Token，原样返回
	OrderID     string `json:"OrderID" xml:"OrderID"`         // 	传入的唯一标识订单的 ID，由商户生成，原样返回
	BizID       string `json:"BizID" xml:"BizID"`             // 	商户 ID，原样返回
	WayBillID   string `json:"WayBillID" xml:"WayBillID"`     // 	运单 ID
	WaybillData string `json:"WaybillData" xml:"WaybillData"` // 	集包地、三段码、大头笔等信息，用于生成面单信息。详见后文返回值说明
}

// TransferCustomerMessage 需要转发的客服消息
type TransferCustomerMessage struct {
	XMLName xml.Name `xml:"xml"`
	// 接收方帐号（收到的OpenID）
	ToUserName string `json:"ToUserName" xml:"ToUserName"`
	// 开发者微信号
	FromUserName string `json:"FromUserName" xml:"FromUserName"`
	// 消息创建时间 （整型）
	CreateTime uint `json:"CreateTime" xml:"CreateTime"`
	// 转发消息类型
	MsgType MsgType `json:"MsgType" xml:"MsgType"`
}

// AddExpressOrderResult 请求下单事件参数
type AddExpressOrderResult struct {
	CommonServerResult
	Token     string `json:"Token" xml:"Token"`         // 订单 Token。请保存该 Token，调用logistics.updatePath时需要传入
	OrderID   string `json:"OrderID" xml:"OrderID"`     // 唯一标识订单的 ID，由商户生成。快递需要保证相同的 OrderID 生成相同的运单ID。
	BizID     string `json:"BizID" xml:"BizID"`         // 商户 ID，即商户在快递注册的客户编码或月结账户名
	BizPwd    string `json:"BizPwd" xml:"BizPwd"`       // BizID 对应的密码
	ShopAppID string `json:"ShopAppID" xml:"ShopAppID"` // 商户的小程序 AppID
	WayBillID string `json:"WayBillID" xml:"WayBillID"` // 运单 ID，从微信号段中生成。若为 0，则表示需要快递来生成运单 ID。
	Remark    string `json:"Remark" xml:"Remark"`       // 快递备注，会打印到面单上，比如"易碎物品"
	Sender    struct {
		Name     string `json:"Name" xml:"Name"`         // 收件人/发件人姓名，不超过64字节
		Tel      string `json:"Tel" xml:"Tel"`           // 收件人/发件人座机号码，若不填写则必须填写 mobile，不超过32字节
		Mobile   string `json:"Mobile" xml:"Mobile"`     // 收件人/发件人手机号码，若不填写则必须填写 tel，不超过32字节
		Company  string `json:"Company" xml:"Company"`   // 收件人/发件人公司名称，不超过64字节
		PostCode string `json:"PostCode" xml:"PostCode"` // 收件人/发件人邮编，不超过10字节
		Country  string `json:"Country" xml:"Country"`   // 收件人/发件人国家，不超过64字节
		Province string `json:"Province" xml:"Province"` // 收件人/发件人省份，比如："广东省"，不超过64字节
		City     string `json:"City" xml:"City"`         // 收件人/发件人市/地区，比如："广州市"，不超过64字节
		Area     string `json:"Area" xml:"Area"`         // 收件人/发件人区/县，比如："海珠区"，不超过64字节
		Address  string `json:"Address" xml:"Address"`   // 收件人/发件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
	} `json:"Sender" xml:"Sender"` // 发件人信息
	Receiver struct {
		Name     string `json:"Name" xml:"Name"`         // 收件人/发件人姓名，不超过64字节
		Tel      string `json:"Tel" xml:"Tel"`           // 收件人/发件人座机号码，若不填写则必须填写 mobile，不超过32字节
		Mobile   string `json:"Mobile" xml:"Mobile"`     // 收件人/发件人手机号码，若不填写则必须填写 tel，不超过32字节
		Company  string `json:"Company" xml:"Company"`   // 收件人/发件人公司名称，不超过64字节
		PostCode string `json:"PostCode" xml:"PostCode"` // 收件人/发件人邮编，不超过10字节
		Country  string `json:"Country" xml:"Country"`   // 收件人/发件人国家，不超过64字节
		Province string `json:"Province" xml:"Province"` // 收件人/发件人省份，比如："广东省"，不超过64字节
		City     string `json:"City" xml:"City"`         // 收件人/发件人市/地区，比如："广州市"，不超过64字节
		Area     string `json:"Area" xml:"Area"`         // 收件人/发件人区/县，比如："海珠区"，不超过64字节
		Address  string `json:"Address" xml:"Address"`   // 收件人/发件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
	} `json:"Receiver" xml:"Receiver"` // 收件人信息
	Cargo struct {
		Weight float64 `json:"Weight" xml:"Weight"`   // 包裹总重量，单位是千克(kg)
		SpaceX float64 `json:"Space_X" xml:"Space_X"` // 包裹长度，单位厘米(cm)
		SpaceY float64 `json:"Space_Y" xml:"Space_Y"` // 包裹宽度，单位厘米(cm)
		SpaceZ float64 `json:"Space_Z" xml:"Space_Z"` // 包裹高度，单位厘米(cm)
		Count  uint    `json:"Count" xml:"Count"`     // 包裹数量
	} `json:"Cargo" xml:"Cargo"` // 包裹信息
	Insured struct {
		Used  uint8 `json:"UseInsured" xml:"UseInsured"`     // 是否保价，0 表示不保价，1 表示保价
		Value uint  `json:"InsuredValue" xml:"InsuredValue"` // 保价金额，单位是分，比如: 10000 表示 100 元
	} `json:"Insured" xml:"Insured"` // 保价信息
	Service struct {
		Type uint8  `json:"ServiceType" xml:"ServiceType"` // 服务类型 ID
		Name string `json:"ServiceName" xml:"ServiceName"` // 服务名称
	} `json:"Service" xml:"Service"` // 服务类型
}

// GetExpressQuotaReturn 查询商户余额事件需要返回的数据
type GetExpressQuotaReturn struct {
	CommonServerReturn
	BizID string  `json:"BizID" xml:"BizID"` // 	商户ID
	Quota float64 `json:"Quota" xml:"Quota"` // 	商户可用余额，0 表示无可用余额
}

// GetExpressQuotaResult 查询商户余额事件参数
type GetExpressQuotaResult struct {
	CommonServerResult
	BizID     string `json:"BizID" xml:"BizID"`         // 商户ID，即商户在快递注册的客户编码或月结账户名
	BizPwd    string `json:"BizPwd" xml:"BizPwd"`       // BizID 对应的密码
	ShopAppID string `json:"ShopAppID" xml:"ShopAppID"` // 商户小程序的 AppID
}

// CancelExpressOrderResult 取消订单事件参数
type CancelExpressOrderResult struct {
	CommonServerResult
	OrderID   string `json:"OrderID" xml:"OrderID"`     // 唯一标识订单的 ID，由商户生成
	BizID     string `json:"BizID" xml:"BizID"`         // 商户 ID
	BizPwd    string `json:"BizPwd" xml:"BizPwd"`       // 商户密码
	ShopAppID string `json:"ShopAppID" xml:"ShopAppID"` // 商户的小程序 AppID
	WayBillID string `json:"WayBillID" xml:"WayBillID"` // 运单 ID，从微信号段中生成
}

// CancelExpressOrderReturn 取消订单事件需要返回的数据
type CancelExpressOrderReturn struct {
	CommonServerReturn
	BizID     string `json:"BizID" xml:"BizID"`         // 商户ID，请原样返回
	OrderID   string `json:"OrderID" xml:"OrderID"`     // 唯一标识订单的ID，由商户生成。请原样返回
	WayBillID string `json:"WayBillID" xml:"WayBillID"` // 运单ID，请原样返回
}

// CheckExpressBusinessResult 审核商户事件参数
type CheckExpressBusinessResult struct {
	CommonServerResult
	BizID         string `json:"BizID" xml:"BizID"`                 // 商户ID，即商户在快递注册的客户编码或月结账户名
	BizPwd        string `json:"BizPwd" xml:"BizPwd"`               // BizID 对应的密码
	ShopAppID     string `json:"ShopAppID" xml:"ShopAppID"`         // 商户的小程序 AppID
	ShopName      string `json:"ShopName" xml:"ShopName"`           // 商户名称，即小程序昵称（仅EMS可用）
	ShopTelphone  string `json:"ShopTelphone" xml:"ShopTelphone"`   // 商户联系电话（仅EMS可用）
	ShopContact   string `json:"ShopContact" xml:"ShopContact"`     // 商户联系人姓名（仅EMS可用）
	ServiceName   string `json:"ServiceName" xml:"ServiceName"`     // 预开通的服务类型名称（仅EMS可用）
	SenderAddress string `json:"SenderAddress" xml:"SenderAddress"` // 商户发货地址（仅EMS可用）
}

// CheckExpressBusinessReturn 审核商户事件需要需要返回的数据
type CheckExpressBusinessReturn struct {
	CommonServerReturn
	BizID string  `json:"BizID" xml:"BizID"` //	商户ID
	Quota float64 `json:"Quota" xml:"Quota"` //	商户可用余额，0 表示无可用余额
}

// DeliveryOrderStatusUpdateResult 服务器携带的参数
type DeliveryOrderStatusUpdateResult struct {
	CommonServerResult
	ShopID      string `json:"shopid" xml:"shopid"`               // 商家id， 由配送公司分配的appkey
	ShopOrderID string `json:"shop_order_id" xml:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo      string `json:"shop_no" xml:"shop_no"`             // 商家门店编号， 在配送公司侧登记
	WaybillID   string `json:"waybill_id" xml:"waybill_id"`       // 配送单id
	ActionTime  uint   `json:"action_time" xml:"action_time"`     // Unix时间戳
	OrderStatus int    `json:"order_status" xml:"order_status"`   // 配送状态，枚举值
	ActionMsg   string `json:"action_msg" xml:"action_msg"`       // 附加信息
	Agent       struct {
		Name  string `json:"name" xml:"name"`   // 骑手姓名
		Phone string `json:"phone" xml:"phone"` // 骑手电话
	} `json:"agent" xml:"agent"` // 骑手信息
}

// DeliveryOrderStatusUpdateReturn 需要返回的数据
type DeliveryOrderStatusUpdateReturn CommonServerReturn

// AgentPosQueryReturn 需要返回的数据
type AgentPosQueryReturn struct {
	CommonServerReturn
	Lng       float64 `json:"lng" xml:"lng"`               // 必填 经度，火星坐标，精确到小数点后6位
	Lat       float64 `json:"lat" xml:"lat"`               // 必填 纬度，火星坐标，精确到小数点后6位
	Distance  float64 `json:"distance" xml:"distance"`     // 必填 和目的地距离，已取货配送中需返回，单位米
	ReachTime uint    `json:"reach_time" xml:"reach_time"` // 必填 预计还剩多久送达时间, 单位秒， 已取货配送中需返回，比如5分钟后送达，填300
}

// AgentPosQueryResult 服务器携带的参数
type AgentPosQueryResult struct {
	CommonServerResult
	ShopID      string `json:"shopid" xml:"shopid"`               // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID string `json:"shop_order_id" xml:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo      string `json:"shop_no" xml:"shop_no"`             // 商家门店编号， 在配送公司侧登记
	WaybillID   string `json:"waybill_id" xml:"waybill_id"`       // 配送单id
}

// AuthInfoGetReturn 需要返回的数据
type AuthInfoGetReturn struct {
	CommonServerReturn
	AppKey      string `json:"appkey" xml:"appkey"`             // 必填 配送公司分配的appkey，对应shopid
	Account     string `json:"account" xml:"account"`           // 必填 帐号名称
	AccountType uint   `json:"account_type" xml:"account_type"` // 必填 帐号类型：0.不确定，1.预充值，2，月结，3，其它
}

// AuthInfoGetResult 服务器携带的参数
type AuthInfoGetResult struct {
	CommonServerResult
	WxAppID string `json:"wx_appid" xml:"wx_appid"` // 	发起授权的商户小程序appid
	Code    string `json:"code" xml:"code"`         // 	授权码
}

// CancelAuthReturn 需要返回的数据
type CancelAuthReturn CommonServerReturn

// CancelAuthResult 服务器携带的参数
type CancelAuthResult struct {
	CommonServerResult
	ShopID  string `json:"shopid" xml:"shopid"`     // 	商家id， 配送公司唯一标识
	WxAppID string `json:"wx_appid" xml:"wx_appid"` // 	发起授权的商户小程序appid
}

// DeliveryOrderAddReturn 需要返回的数据
type DeliveryOrderAddReturn struct {
	CommonServerReturn
	Event            string  `json:"Event" xml:"Event"`                                             // 是 事件类型，固定为 transport_add_order，不区分大小写
	Fee              uint    `json:"fee" xml:"fee"`                                                 // 是 实际运费(单位：元)，运费减去优惠券费用
	Deliverfee       uint    `json:"deliverfee" xml:"deliverfee"`                                   // 是 运费(单位：元)
	Couponfee        uint    `json:"couponfee" xml:"couponfee"`                                     // 是 优惠券费用(单位：元)
	Tips             uint    `json:"tips" xml:"tips"`                                               // 是 小费(单位：元)
	Insurancefee     uint    `json:"insurancefee" xml:"insurancefee"`                               // 是 保价费(单位：元)
	Distance         float64 `json:"distance,omitempty" xml:"distance,omitempty"`                   // 否 配送距离(单位：米)
	WaybillID        string  `json:"waybill_id,omitempty" xml:"waybill_id,omitempty"`               // 否 配送单号, 可以在API1更新配送单状态异步返回
	OrderStatus      int     `json:"order_status" xml:"order_status"`                               // 是 配送单状态
	FinishCode       uint    `json:"finish_code,omitempty" xml:"finish_code,omitempty"`             // 否 收货码
	PickupCode       uint    `json:"pickup_code,omitempty" xml:"pickup_code,omitempty"`             // 否 取货码
	DispatchDuration uint    `json:"dispatch_duration,omitempty" xml:"dispatch_duration,omitempty"` // 否 预计骑手接单时间，单位秒，比如5分钟，就填300, 无法预计填0
	SenderLng        float64 `json:"sender_lng,omitempty" xml:"sender_lng,omitempty"`               // 否 发货方经度，火星坐标，精确到小数点后6位， 用于消息通知，如果下单请求里有发货人信息则不需要
	SenderLat        float64 `json:"sender_lat,omitempty" xml:"sender_lat,omitempty"`               // 否 发货方纬度，火星坐标，精确到小数点后6位， 用于消息通知，如果下单请求里有发货人信息则不需要
}

// DeliveryOrderAddResult 服务器携带的参数
type DeliveryOrderAddResult struct {
	CommonServerResult
	WxToken       string `json:"wx_token" xml:"wx_token"`             // 	微信订单 Token。请保存该Token，调用更新配送单状态接口（updateOrder）时需要传入
	DeliveryToken string `json:"delivery_token" xml:"delivery_token"` // 	配送公司侧在预下单时候返回的token，用于保证运费不变
	ShopID        string `json:"shopid" xml:"shopid"`                 // 	商家id， 由配送公司分配的appkey
	ShopNo        string `json:"shop_no" xml:"shop_no"`               // 	商家门店编号， 在配送公司侧登记
	ShopOrderID   string `json:"shop_order_id" xml:"shop_order_id"`   // 	唯一标识订单的 ID，由商户生成
	DeliverySign  string `json:"delivery_sign" xml:"delivery_sign"`   // 	用配送公司侧提供的appSecret加密的校验串
	Sender        struct {
		Name           string  `json:"name" xml:"name"`                       // 姓名，最长不超过256个字符
		City           string  `json:"city" xml:"city"`                       // 城市名称，如广州市
		Address        string  `json:"address" xml:"address"`                 // 地址(街道、小区、大厦等，用于定位)
		AddressDetail  string  `json:"address_detail" xml:"address_detail"`   // 地址详情(楼号、单元号、层号)
		Phone          string  `json:"phone" xml:"phone"`                     // 电话/手机号，最长不超过64个字符
		Lng            float64 `json:"lng" xml:"lng"`                         // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
		Lat            float64 `json:"lat" xml:"lat"`                         // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
		CoordinateType uint8   `json:"coordinate_type" xml:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
	} `json:"sender" xml:"sender"` // 发件人信息，如果配送公司能从shopid+shop_no对应到门店地址，则不需要填写，否则需填写
	Receiver struct {
		Name           string  `json:"name" xml:"name"`                       // 姓名，最长不超过256个字符
		City           string  `json:"city" xml:"city"`                       // 城市名称，如广州市
		Address        string  `json:"address" xml:"address"`                 // 地址(街道、小区、大厦等，用于定位)
		AddressDetail  string  `json:"address_detail" xml:"address_detail"`   // 地址详情(楼号、单元号、层号)
		Phone          string  `json:"phone" xml:"phone"`                     // 电话/手机号，最长不超过64个字符
		Lng            float64 `json:"lng" xml:"lng"`                         // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
		Lat            float64 `json:"lat" xml:"lat"`                         // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
		CoordinateType uint8   `json:"coordinate_type" xml:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
	} `json:"receiver" xml:"receiver"` // 收件人信息
	Cargo struct {
		GoodsValue  float64 `json:"goods_value" xml:"goods_value"`   // 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-5000]
		GoodsHeight float64 `json:"goods_height" xml:"goods_height"` // 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-45]
		GoodsLength float64 `json:"goods_length" xml:"goods_length"` // 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-65]
		GoodsWidth  float64 `json:"goods_width" xml:"goods_width"`   // 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
		GoodsWeight float64 `json:"goods_weight" xml:"goods_weight"` // 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
		GoodsDetail struct {
			Goods []struct {
				Count uint    `json:"good_count" xml:"good_count"` // 货物数量
				Name  string  `json:"good_name" xml:"good_name"`   // 货品名称
				Price float64 `json:"good_price" xml:"good_price"` // 货品单价，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）
				Unit  string  `json:"good_unit" xml:"good_unit"`   // 货品单位，最长不超过20个字符
			} `json:"goods" xml:"goods"` // 货物列表
		} `json:"goods_detail" xml:"goods_detail"` // 货物详情，最长不超过10240个字符
		GoodsPickupInfo   string `json:"goods_pickup_info" xml:"goods_pickup_info"`     // 货物取货信息，用于骑手到店取货，最长不超过100个字符
		GoodsDeliveryInfo string `json:"goods_delivery_info" xml:"goods_delivery_info"` // 货物交付信息，最长不超过100个字符
		CargoFirstClass   string `json:"cargo_first_class" xml:"cargo_first_class"`     // 品类一级类目
		CargoSecondClass  string `json:"cargo_second_class" xml:"cargo_second_class"`   // 品类二级类目
	} `json:"cargo" xml:"cargo"` // 货物信息
	OrderInfo struct {
		DeliveryServiceCode  string  `json:"delivery_service_code" xml:"delivery_service_code"`   // 配送服务代码 不同配送公司自定义,微信侧不理解
		OrderType            uint8   `json:"order_type" xml:"order_type"`                         // 订单类型, 0: 即时单 1 预约单，如预约单，需要设置expected_delivery_time或expected_finish_time或expected_pick_time
		ExpectedDeliveryTime uint    `json:"expected_delivery_time" xml:"expected_delivery_time"` // 期望派单时间(达达支持，表示达达系统调度时间)，unix-timestamp
		ExpectedFinishTime   uint    `json:"expected_finish_time" xml:"expected_finish_time"`     // 期望送达时间(美团、顺丰同城急送支持），unix-timestamp)
		ExpectedPickTime     uint    `json:"expected_pick_time" xml:"expected_pick_time"`         // 期望取件时间（闪送、顺丰同城急送支持，顺丰同城急送只需传expected_finish_time或expected_pick_time其中之一即可，同时都传则以expected_finish_time为准），unix-timestamp
		PoiSeq               string  `json:"poi_seq" xml:"poi_seq"`                               // 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
		Note                 string  `json:"note" xml:"note"`                                     // 备注，最长不超过200个字符
		OrderTime            uint    `json:"order_time" xml:"order_time"`                         // 用户下单付款时间
		IsInsured            uint8   `json:"is_insured" xml:"is_insured"`                         // 是否保价，0，非保价，1.保价
		DeclaredValue        float64 `json:"declared_value" xml:"declared_value"`                 // 保价金额，单位为元，精确到分
		Tips                 float64 `json:"tips" xml:"tips"`                                     // 小费，单位为元, 下单一般不加小费
		IsDirectDelivery     float64 `json:"is_direct_delivery" xml:"is_direct_delivery"`         // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑手只能配送此订单至完成，配送费用也相应高一些，闪送必须选1，达达可选0或1，其余配送公司不支持直拿直送）
		CashOnDelivery       float64 `json:"cash_on_delivery" xml:"cash_on_delivery"`             // 骑手应付金额，单位为元，精确到分
		CashOnPickup         float64 `json:"cash_on_pickup" xml:"cash_on_pickup"`                 // 骑手应收金额，单位为元，精确到分
		RiderPickMethod      uint8   `json:"rider_pick_method" xml:"rider_pick_method"`           // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
		IsFinishCodeNeeded   uint8   `json:"is_finish_code_needed" xml:"is_finish_code_needed"`   // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
		IsPickupCodeNeeded   uint8   `json:"is_pickup_code_needed" xml:"is_pickup_code_needed"`   // 取货码（0：不需要；1：需要。取货码的作用是：骑手必须输入取货码才能从商家取货）
	} `json:"order_info" xml:"order_info"` // 订单信息
}

// DeliveryOrderAddTipsReturn 需要返回的数据
type DeliveryOrderAddTipsReturn CommonServerReturn

// DeliveryOrderAddTipsResult 服务器携带的参数
type DeliveryOrderAddTipsResult struct {
	CommonServerResult
	ShopID       string  `json:"shopid" xml:"shopid"`               // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID  string  `json:"shop_order_id" xml:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo       string  `json:"shop_no" xml:"shop_no"`             // 商家门店编号， 在配送公司侧登记
	WaybillID    string  `json:"waybill_id" xml:"waybill_id"`       // 配送单id
	DeliverySign string  `json:"delivery_sign" xml:"delivery_sign"` // 用配送公司侧提供的appSecret加密的校验串
	Tips         float64 `json:"tips" xml:"tips"`                   // 小费金额(单位：元)
	Remark       string  `json:"remark" xml:"remark"`               // 备注
}

// DeliveryOrderCancelReturn 需要返回的数据
type DeliveryOrderCancelReturn struct {
	CommonServerReturn
	DeductFee uint   `json:"deduct_fee" xml:"deduct_fee"` // 是	预计扣除的违约金(单位：元)，可能没有
	Desc      string `json:"desc" xml:"desc"`             // 是	扣费说明
}

// DeliveryOrderCancelResult 服务器携带的参数
type DeliveryOrderCancelResult struct {
	CommonServerResult
	ShopID         string `json:"shopid" xml:"shopid"`                     // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID    string `json:"shop_order_id" xml:"shop_order_id"`       // 唯一标识订单的 ID，由商户生成
	ShopNo         string `json:"shop_no" xml:"shop_no"`                   // 商家门店编号， 在配送公司侧登记
	WaybillID      string `json:"waybill_id" xml:"waybill_id"`             // 配送单id
	DeliverySign   string `json:"delivery_sign" xml:"delivery_sign"`       // 用配送公司侧提供的appSecret加密的校验串
	CancelReasonID uint   `json:"cancel_reason_id" xml:"cancel_reason_id"` // 取消原因id
	CancelReason   string `json:"cancel_reason" xml:"cancel_reason"`       // 取消原因
}

// DeliveryOrderReturnConfirmReturn 需要返回的数据
type DeliveryOrderReturnConfirmReturn CommonServerReturn

// DeliveryOrderReturnConfirmResult 服务器携带的参数
type DeliveryOrderReturnConfirmResult struct {
	CommonServerResult
	ShopID       string `json:"shopid" xml:"shopid"`               // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID  string `json:"shop_order_id" xml:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo       string `json:"shop_no" xml:"shop_no"`             // 商家门店编号， 在配送公司侧登记
	WaybillID    string `json:"waybill_id" xml:"waybill_id"`       // 配送单id
	DeliverySign string `json:"delivery_sign" xml:"delivery_sign"` // 用配送公司侧提供的appSecret加密的校验串
}

// DeliveryOrderPreAddReturn 需要返回的数据
type DeliveryOrderPreAddReturn struct {
	CommonServerReturn
	Fee              uint    `json:"fee" xml:"fee"`                             // 是	实际运费(单位：元)，运费减去优惠券费用
	Deliverfee       uint    `json:"deliverfee" xml:"deliverfee"`               // 是	运费(单位：元)
	Couponfee        uint    `json:"couponfee" xml:"couponfee"`                 // 是	优惠券费用(单位：元)
	Tips             float64 `json:"tips" xml:"tips"`                           // 是	小费(单位：元)
	Insurancefee     uint    `json:"insurancefee" xml:"insurancefee"`           // 是	保价费(单位：元)
	Distance         uint    `json:"distance" xml:"distance"`                   // 否	配送距离(单位：米)
	DispatchDuration uint    `json:"dispatch_duration" xml:"dispatch_duration"` // 否	预计骑手接单时间，单位秒，比如5分钟，就填300, 无法预计填0
	DeliveryToken    string  `json:"delivery_token" xml:"delivery_token"`       // 否	配送公司可以返回此字段，当用户下单时候带上这个字段，配送公司可保证在一段时间内运费不变
}

// DeliveryOrderPreAddResult 服务器携带的参数
type DeliveryOrderPreAddResult struct {
	CommonServerResult
	ShopID       string `json:"shopid" xml:"shopid"`               // 商家id， 由配送公司分配的appkey
	ShopNo       string `json:"shop_no" xml:"shop_no"`             // 商家门店编号， 在配送公司侧登记
	ShopOrderID  string `json:"shop_order_id" xml:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	DeliverySign string `json:"delivery_sign" xml:"delivery_sign"` // 用配送公司侧提供的appSecret加密的校验串
	Sender       struct {
		Name           string  `json:"name" xml:"name"`                       // 姓名，最长不超过256个字符
		City           string  `json:"city" xml:"city"`                       // 城市名称，如广州市
		Address        string  `json:"address" xml:"address"`                 // 地址(街道、小区、大厦等，用于定位)
		AddressDetail  string  `json:"address_detail" xml:"address_detail"`   // 地址详情(楼号、单元号、层号)
		Phone          string  `json:"phone" xml:"phone"`                     // 电话/手机号，最长不超过64个字符
		Lng            float64 `json:"lng" xml:"lng"`                         // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
		Lat            float64 `json:"lat" xml:"lat"`                         // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
		CoordinateType uint8   `json:"coordinate_type" xml:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
	} `json:"sender" xml:"sender"` // 发件人信息，如果配送公司能从shopid+shop_no对应到门店地址，则不需要填写，否则需填写
	Receiver struct {
		Name           string  `json:"name" xml:"name"`                       // 姓名，最长不超过256个字符
		City           string  `json:"city" xml:"city"`                       // 城市名称，如广州市
		Address        string  `json:"address" xml:"address"`                 // 地址(街道、小区、大厦等，用于定位)
		AddressDetail  string  `json:"address_detail" xml:"address_detail"`   // 地址详情(楼号、单元号、层号)
		Phone          string  `json:"phone" xml:"phone"`                     // 电话/手机号，最长不超过64个字符
		Lng            float64 `json:"lng" xml:"lng"`                         // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
		Lat            float64 `json:"lat" xml:"lat"`                         // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
		CoordinateType uint8   `json:"coordinate_type" xml:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
	} `json:"receiver" xml:"receiver"` // 收件人信息
	Cargo struct {
		GoodsValue  float64 `json:"goods_value" xml:"goods_value"`   // 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-5000]
		GoodsHeight float64 `json:"goods_height" xml:"goods_height"` // 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-45]
		GoodsLength float64 `json:"goods_length" xml:"goods_length"` // 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-65]
		GoodsWidth  float64 `json:"goods_width" xml:"goods_width"`   // 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
		GoodsWeight float64 `json:"goods_weight" xml:"goods_weight"` // 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
		GoodsDetail struct {
			Goods []struct {
				Count uint    `json:"good_count" xml:"good_count"` // 货物数量
				Name  string  `json:"good_name" xml:"good_name"`   // 货品名称
				Price float64 `json:"good_price" xml:"good_price"` // 货品单价，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）
				Unit  string  `json:"good_unit" xml:"good_unit"`   // 货品单位，最长不超过20个字符
			} `json:"goods" xml:"goods"` // 货物列表
		} `json:"goods_detail" xml:"goods_detail"` // 货物详情，最长不超过10240个字符
		GoodsPickupInfo   string `json:"goods_pickup_info" xml:"goods_pickup_info"`     // 货物取货信息，用于骑手到店取货，最长不超过100个字符
		GoodsDeliveryInfo string `json:"goods_delivery_info" xml:"goods_delivery_info"` // 货物交付信息，最长不超过100个字符
		CargoFirstClass   string `json:"cargo_first_class" xml:"cargo_first_class"`     // 品类一级类目
		CargoSecondClass  string `json:"cargo_second_class" xml:"cargo_second_class"`   // 品类二级类目
	} `json:"cargo" xml:"cargo"` // 货物信息
	OrderInfo struct {
		DeliveryServiceCode  string  `json:"delivery_service_code" xml:"delivery_service_code"`   // 配送服务代码 不同配送公司自定义,微信侧不理解
		OrderType            uint8   `json:"order_type" xml:"order_type"`                         // 订单类型, 0: 即时单 1 预约单，如预约单，需要设置expected_delivery_time或expected_finish_time或expected_pick_time
		ExpectedDeliveryTime uint    `json:"expected_delivery_time" xml:"expected_delivery_time"` // 期望派单时间(达达支持，表示达达系统调度时间)，unix-timestamp
		ExpectedFinishTime   uint    `json:"expected_finish_time" xml:"expected_finish_time"`     // 期望送达时间(美团、顺丰同城急送支持），unix-timestamp)
		ExpectedPickTime     uint    `json:"expected_pick_time" xml:"expected_pick_time"`         // 期望取件时间（闪送、顺丰同城急送支持，顺丰同城急送只需传expected_finish_time或expected_pick_time其中之一即可，同时都传则以expected_finish_time为准），unix-timestamp
		PoiSeq               string  `json:"poi_seq" xml:"poi_seq"`                               // 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
		Note                 string  `json:"note" xml:"note"`                                     // 备注，最长不超过200个字符
		OrderTime            uint    `json:"order_time" xml:"order_time"`                         // 用户下单付款时间
		IsInsured            uint8   `json:"is_insured" xml:"is_insured"`                         // 是否保价，0，非保价，1.保价
		DeclaredValue        float64 `json:"declared_value" xml:"declared_value"`                 // 保价金额，单位为元，精确到分
		Tips                 float64 `json:"tips" xml:"tips"`                                     // 小费，单位为元, 下单一般不加小费
		IsDirectDelivery     float64 `json:"is_direct_delivery" xml:"is_direct_delivery"`         // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑手只能配送此订单至完成，配送费用也相应高一些，闪送必须选1，达达可选0或1，其余配送公司不支持直拿直送）
		CashOnDelivery       float64 `json:"cash_on_delivery" xml:"cash_on_delivery"`             // 骑手应付金额，单位为元，精确到分
		CashOnPickup         float64 `json:"cash_on_pickup" xml:"cash_on_pickup"`                 // 骑手应收金额，单位为元，精确到分
		RiderPickMethod      uint8   `json:"rider_pick_method" xml:"rider_pick_method"`           // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
		IsFinishCodeNeeded   uint8   `json:"is_finish_code_needed" xml:"is_finish_code_needed"`   // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
		IsPickupCodeNeeded   uint8   `json:"is_pickup_code_needed" xml:"is_pickup_code_needed"`   // 取货码（0：不需要；1：需要。取货码的作用是：骑手必须输入取货码才能从商家取货）
	} `json:"order_info" xml:"order_info"` // 订单信息
}

// DeliveryOrderPreCancelReturn 需要返回的数据
type DeliveryOrderPreCancelReturn struct {
	CommonServerReturn
	DeductFee uint   `json:"deduct_fee" xml:"deduct_fee"` // 是	预计扣除的违约金(单位：元)，可能没有
	Desc      string `json:"desc" xml:"desc"`             // 是	扣费说明
}

// DeliveryOrderPreCancelResult 服务器携带的参数
type DeliveryOrderPreCancelResult struct {
	CommonServerResult
	ShopID         string `json:"shopid" xml:"shopid"`                     // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID    string `json:"shop_order_id" xml:"shop_order_id"`       // 唯一标识订单的 ID，由商户生成
	ShopNo         string `json:"shop_no" xml:"shop_no"`                   // 商家门店编号， 在配送公司侧登记
	WaybillID      string `json:"waybill_id" xml:"waybill_id"`             // 配送单id
	DeliverySign   string `json:"delivery_sign" xml:"delivery_sign"`       // 用配送公司侧提供的appSecret加密的校验串
	CancelReasonID uint   `json:"cancel_reason_id" xml:"cancel_reason_id"` // 取消原因id
	CancelReason   string `json:"cancel_reason" xml:"cancel_reason"`       // 取消原因
}

// DeliveryOrderQueryReturn 需要返回的数据
type DeliveryOrderQueryReturn struct {
	CommonServerReturn
	OrderStatus float64 `json:"order_status" xml:"order_status"` // 是	当前订单状态，枚举值
	ActionMsg   string  `json:"action_msg" xml:"action_msg"`     // 否	附加信息
	WaybillID   string  `json:"waybill_id" xml:"waybill_id"`     // 是	配送单id
}

// DeliveryOrderQueryResult 服务器携带的参数
type DeliveryOrderQueryResult struct {
	CommonServerResult
	ShopID       string `json:"shopid" xml:"shopid"`               // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID  string `json:"shop_order_id" xml:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo       string `json:"shop_no" xml:"shop_no"`             // 商家门店编号， 在配送公司侧登记
	WaybillID    string `json:"waybill_id" xml:"waybill_id"`       // 配送单id
	DeliverySign string `json:"delivery_sign" xml:"delivery_sign"` // 用配送公司侧提供的appSecret加密的校验串
}

// DeliveryOrderReaddReturn 需要返回的数据
type DeliveryOrderReaddReturn struct {
	CommonServerReturn
	Fee              uint    `json:"fee" xml:"fee"`                             // 是	实际运费(单位：元)，运费减去优惠券费用
	Deliverfee       uint    `json:"deliverfee" xml:"deliverfee"`               // 是	运费(单位：元)
	Couponfee        uint    `json:"couponfee" xml:"couponfee"`                 // 是	优惠券费用(单位：元)
	Tips             float64 `json:"tips" xml:"tips"`                           // 是	小费(单位：元)
	Insurancefee     uint    `json:"insurancefee" xml:"insurancefee"`           // 是	保价费(单位：元)
	Distance         uint    `json:"distance" xml:"distance"`                   // 否	配送距离(单位：米)
	WaybillID        string  `json:"waybill_id" xml:"waybill_id"`               // 否	配送单号, 可以在API1更新配送单状态异步返回
	OrderStatus      float64 `json:"order_status" xml:"order_status"`           // 是	配送单状态
	FinishCode       uint    `json:"finish_code" xml:"finish_code"`             // 否	收货码
	PickupCode       uint    `json:"pickup_code" xml:"pickup_code"`             // 否	取货码
	DispatchDuration uint    `json:"dispatch_duration" xml:"dispatch_duration"` // 否	预计骑手接单时间，单位秒，比如5分钟，就填300, 无法预计填0
	SenderLng        float64 `json:"sender_lng" xml:"sender_lng"`               // 否	发货方经度，火星坐标，精确到小数点后6位， 用于消息通知，如果下单请求里有发货人信息则不需要
	SenderLat        float64 `json:"sender_lat" xml:"sender_lat"`               // 否	发货方纬度，火星坐标，精确到小数点后6位， 用于消息通知，如果下单请求里有发货人信息则不需要
}

// DeliveryOrderReaddResult 服务器携带的参数
type DeliveryOrderReaddResult struct {
	CommonServerResult
	WxToken       string `json:"wx_token" xml:"wx_token"`             // 微信订单 Token。请保存该Token，调用更新配送单状态接口（updateOrder）时需要传入
	DeliveryToken string `json:"delivery_token" xml:"delivery_token"` // 配送公司侧在预下单时候返回的token，用于保证运费不变
	ShopID        string `json:"shopid" xml:"shopid"`                 // 商家id， 由配送公司分配的appkey
	ShopNo        string `json:"shop_no" xml:"shop_no"`               // 商家门店编号， 在配送公司侧登记
	ShopOrderID   string `json:"shop_order_id" xml:"shop_order_id"`   // 唯一标识订单的 ID，由商户生成
	DeliverySign  string `json:"delivery_sign" xml:"delivery_sign"`   // 用配送公司侧提供的appSecret加密的校验串
	Sender        struct {
		Name           string  `json:"name" xml:"name"`                       // 姓名，最长不超过256个字符
		City           string  `json:"city" xml:"city"`                       // 城市名称，如广州市
		Address        string  `json:"address" xml:"address"`                 // 地址(街道、小区、大厦等，用于定位)
		AddressDetail  string  `json:"address_detail" xml:"address_detail"`   // 地址详情(楼号、单元号、层号)
		Phone          string  `json:"phone" xml:"phone"`                     // 电话/手机号，最长不超过64个字符
		Lng            float64 `json:"lng" xml:"lng"`                         // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
		Lat            float64 `json:"lat" xml:"lat"`                         // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
		CoordinateType uint8   `json:"coordinate_type" xml:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
	} `json:"sender" xml:"sender"` // 发件人信息，如果配送公司能从shopid+shop_no对应到门店地址，则不需要填写，否则需填写
	Receiver struct {
		Name           string  `json:"name" xml:"name"`                       // 姓名，最长不超过256个字符
		City           string  `json:"city" xml:"city"`                       // 城市名称，如广州市
		Address        string  `json:"address" xml:"address"`                 // 地址(街道、小区、大厦等，用于定位)
		AddressDetail  string  `json:"address_detail" xml:"address_detail"`   // 地址详情(楼号、单元号、层号)
		Phone          string  `json:"phone" xml:"phone"`                     // 电话/手机号，最长不超过64个字符
		Lng            float64 `json:"lng" xml:"lng"`                         // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
		Lat            float64 `json:"lat" xml:"lat"`                         // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
		CoordinateType uint8   `json:"coordinate_type" xml:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
	} `json:"receiver" xml:"receiver"` // 收件人信息
	Cargo struct {
		GoodsValue  float64 `json:"goods_value" xml:"goods_value"`   // 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-5000]
		GoodsHeight float64 `json:"goods_height" xml:"goods_height"` // 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-45]
		GoodsLength float64 `json:"goods_length" xml:"goods_length"` // 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-65]
		GoodsWidth  float64 `json:"goods_width" xml:"goods_width"`   // 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
		GoodsWeight float64 `json:"goods_weight" xml:"goods_weight"` // 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
		GoodsDetail struct {
			Goods []struct {
				Count uint    `json:"good_count" xml:"good_count"` // 货物数量
				Name  string  `json:"good_name" xml:"good_name"`   // 货品名称
				Price float64 `json:"good_price" xml:"good_price"` // 货品单价，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）
				Unit  string  `json:"good_unit" xml:"good_unit"`   // 货品单位，最长不超过20个字符
			} `json:"goods" xml:"goods"` // 货物列表
		} `json:"goods_detail" xml:"goods_detail"` // 货物详情，最长不超过10240个字符
		GoodsPickupInfo   string `json:"goods_pickup_info" xml:"goods_pickup_info"`     // 货物取货信息，用于骑手到店取货，最长不超过100个字符
		GoodsDeliveryInfo string `json:"goods_delivery_info" xml:"goods_delivery_info"` // 货物交付信息，最长不超过100个字符
		CargoFirstClass   string `json:"cargo_first_class" xml:"cargo_first_class"`     // 品类一级类目
		CargoSecondClass  string `json:"cargo_second_class" xml:"cargo_second_class"`   // 品类二级类目
	} `json:"cargo" xml:"cargo"` // 货物信息
	OrderInfo struct {
		DeliveryServiceCode  string  `json:"delivery_service_code" xml:"delivery_service_code"`   // 配送服务代码 不同配送公司自定义,微信侧不理解
		OrderType            uint8   `json:"order_type" xml:"order_type"`                         // 订单类型, 0: 即时单 1 预约单，如预约单，需要设置expected_delivery_time或expected_finish_time或expected_pick_time
		ExpectedDeliveryTime uint    `json:"expected_delivery_time" xml:"expected_delivery_time"` // 期望派单时间(达达支持，表示达达系统调度时间)，unix-timestamp
		ExpectedFinishTime   uint    `json:"expected_finish_time" xml:"expected_finish_time"`     // 期望送达时间(美团、顺丰同城急送支持），unix-timestamp)
		ExpectedPickTime     uint    `json:"expected_pick_time" xml:"expected_pick_time"`         // 期望取件时间（闪送、顺丰同城急送支持，顺丰同城急送只需传expected_finish_time或expected_pick_time其中之一即可，同时都传则以expected_finish_time为准），unix-timestamp
		PoiSeq               string  `json:"poi_seq" xml:"poi_seq"`                               // 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
		Note                 string  `json:"note" xml:"note"`                                     // 备注，最长不超过200个字符
		OrderTime            uint    `json:"order_time" xml:"order_time"`                         // 用户下单付款时间
		IsInsured            uint8   `json:"is_insured" xml:"is_insured"`                         // 是否保价，0，非保价，1.保价
		DeclaredValue        float64 `json:"declared_value" xml:"declared_value"`                 // 保价金额，单位为元，精确到分
		Tips                 float64 `json:"tips" xml:"tips"`                                     // 小费，单位为元, 下单一般不加小费
		IsDirectDelivery     float64 `json:"is_direct_delivery" xml:"is_direct_delivery"`         // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑手只能配送此订单至完成，配送费用也相应高一些，闪送必须选1，达达可选0或1，其余配送公司不支持直拿直送）
		CashOnDelivery       float64 `json:"cash_on_delivery" xml:"cash_on_delivery"`             // 骑手应付金额，单位为元，精确到分
		CashOnPickup         float64 `json:"cash_on_pickup" xml:"cash_on_pickup"`                 // 骑手应收金额，单位为元，精确到分
		RiderPickMethod      uint8   `json:"rider_pick_method" xml:"rider_pick_method"`           // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
		IsFinishCodeNeeded   uint8   `json:"is_finish_code_needed" xml:"is_finish_code_needed"`   // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
		IsPickupCodeNeeded   uint8   `json:"is_pickup_code_needed" xml:"is_pickup_code_needed"`   // 取货码（0：不需要；1：需要。取货码的作用是：骑手必须输入取货码才能从商家取货）
	} `json:"order_info" xml:"order_info"` // 订单信息
}

// PreAuthCodeGetReturn 需要返回的数据
type PreAuthCodeGetReturn struct {
	CommonServerReturn
	PreAuthCode string `json:"pre_auth_code" xml:"pre_auth_code"` // 是	预授权码
}

// PreAuthCodeGetResult 服务器携带的参数
type PreAuthCodeGetResult struct {
	CommonServerResult
	WxAppID string `json:"wx_appid" xml:"wx_appid"` // 发起授权的商户小程序appid
}

// RiderScoreSetReturn 需要返回的数据
type RiderScoreSetReturn CommonServerReturn

// RiderScoreSetResult 服务器携带的参数
type RiderScoreSetResult struct {
	CommonServerResult
	ShopID              string `json:"shopid" xml:"shopid"`                               // 商家id， 由配送公司分配，可以是dev_id或者appkey
	ShopOrderID         string `json:"shop_order_id" xml:"shop_order_id"`                 // 唯一标识订单的 ID，由商户生成
	ShopNo              string `json:"shop_no" xml:"shop_no"`                             // 商家门店编号， 在配送公司侧登记
	WaybillID           string `json:"waybill_id" xml:"waybill_id"`                       // 配送单id
	DeliveryOntimeScore uint   `json:"delivery_ontime_score" xml:"delivery_ontime_score"` // 配送准时分数，范围 1 - 5
	CargoIntactScore    uint   `json:"cargo_intact_score" xml:"cargo_intact_score"`       // 货物完整分数，范围1-5
	AttitudeScore       uint   `json:"attitude_score" xml:"attitude_score"`               // 服务态度分数 范围1-5
}

// 订阅结果
type SubscribeResult = string

const (
	SubscribeResultAccept SubscribeResult = "accept"
	SubscribeResultReject SubscribeResult = "reject"
)

// 用户触发订阅消息弹框事件内容
type SubscribeMsgPopupEvent struct {
	CommonServerResult
	// https://developers.weixin.qq.com/community/develop/doc/000e0c47cb85b070d1bc00fcf51c00?fromCreate=0
	SubscribeMsgPopupEvent []*UserSubscribedMsg `json:"List" xml:"SubscribeMsgPopupEvent"`
}

// 订阅消息发送结果通知事件内容
type SubscribeMsgSentEvent struct {
	CommonServerResult
	SubscribeMsgSentEvent struct {
		List struct {
			// 模板id（一次订阅可能有多个id）
			TemplateId string `json:"TemplateId" xml:"TemplateId"`
			// 消息id（调用接口时也会返回）
			MsgID int `json:"MsgId" xml:"MsgId"`
			// 推送结果状态码（0表示成功）
			ErrorCode int `json:"ErrorCode" xml:"ErrorCode"`
			// 推送结果状态码对应的含义
			ErrorStatus int `json:"ErrorStatus" xml:"ErrorStatus"`
		} `json:"List" xml:"List"`
	} `json:"SubscribeMsgSentEvent" xml:"SubscribeMsgSentEvent"`
}

// 订阅的模板
type UserSubscribedMsg struct {
	// 模板id（一次订阅可能有多个id）
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
	// 订阅结果（accept接收；reject拒收）
	SubscribeStatusString string `json:"SubscribeStatusString" xml:"SubscribeStatusString"`
	// 弹框场景，0代表在小程序页面内
	PopupScene string `json:"PopupScene" xml:"PopupScene"`
}

// 用户改变订阅消息事件内容
type SubscribeMsgChangeEvent struct {
	CommonServerResult
	// https://developers.weixin.qq.com/community/develop/doc/000e0c47cb85b070d1bc00fcf51c00?fromCreate=0
	SubscribeMsgChangeEvent []*UserChangesSubscribeMsg `json:"List" xml:"SubscribeMsgChangeEvent"`
}

// 订阅的模板
type UserChangesSubscribeMsg struct {
	// 模板id（一次订阅可能有多个id）
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
	// 订阅结果（accept接收；reject拒收）
	SubscribeStatusString string `json:"SubscribeStatusString" xml:"SubscribeStatusString"`
}
