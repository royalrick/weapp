package weapp

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

type CommonServerResult struct {
	XMLName      xml.Name  `xml:"xml" json:"-"`
	ToUserName   string    `json:"ToUserName" xml:"ToUserName"`     // 小程序的原始ID
	FromUserName string    `json:"FromUserName" xml:"FromUserName"` // 发送者的 openID | 平台推送服务UserName
	CreateTime   uint64    `json:"CreateTime" xml:"CreateTime"`     // 消息创建时间(整型）
	MsgType      MsgType   `json:"MsgType" xml:"MsgType"`           // 消息类型
	Event        EventType `json:"Event" xml:"Event"`               // 事件类型
}

// UserEnterTempsessionResult 接收的文本消息
type UserEnterTempsessionResult struct {
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
	IsRisky       uint8  `json:"isrisky"`         // 检测结果，0：暂未检测到风险，1：风险
	ExtraInfoJSON string `json:"extra_info_json"` // 附加信息，默认为空
	AppID         string `json:"appid"`           // 小程序的appid
	TraceID       string `json:"trace_id"`        // 任务id
	StatusCode    int    `json:"status_code"`     // 默认为：0，4294966288(-1008)为链接无法下载
}

// AddNearbyPoiAuditResult 附近小程序添加地点审核状态通知数据
type AddNearbyPoiAuditResult struct {
	CommonServerResult
	AuditID uint   `xml:"audit_id"` // 审核单id
	Status  uint8  `xml:"status"`   // 审核状态（3：审核通过，2：审核失败）
	Reason  string `xml:"reason"`   // 如果status为2，会返回审核失败的原因
	PoiID   uint   `xml:"poi_id"`
}

// ExpressPathUpdateResult 运单轨迹更新事件返回数据
type ExpressPathUpdateResult struct {
	CommonServerResult
	DeliveryID string `json:"DeliveryID" xml:"DeliveryID"` // 快递公司ID
	WayBillID  string `json:"WayBillId" xml:"WayBillId"`   // 运单ID
	Version    uint   `json:"Version" xml:"Version"`       // 轨迹版本号（整型）
	Count      uint   `json:"Count" xml:"Count"`           // 轨迹节点数（整型）
	Actions    []struct {
		ActionTime uint   `json:"ActionTime" xml:"ActionTime"` // 轨迹节点 Unix 时间戳
		ActionType uint   `json:"ActionType" xml:"ActionType"` // 轨迹节点类型
		ActionMsg  string `json:"ActionMsg" xml:"ActionMsg"`   // 轨迹节点详情
	} `json:"Actions" xml:"Actions"` // 轨迹列表
}

// AddExpressOrderReturn 请求下单事件返回数据
type AddExpressOrderReturn struct {
	ToUserName   string `json:"ToUserName" xml:"ToUserName"`     // 	原样返回请求中的 FromUserName
	FromUserName string `json:"FromUserName" xml:"FromUserName"` // 	快递公司小程序 UserName
	CreateTime   uint   `json:"CreateTime" xml:"CreateTime"`     // 	事件时间，Unix时间戳
	MsgType      string `json:"MsgType" xml:"MsgType"`           // 	消息类型，固定为event
	Event        string `json:"Event" xml:"Event"`               // 	事件类型
	Token        string `json:"Token" xml:"Token"`               // 	传入的 Token，原样返回
	OrderID      string `json:"OrderID" xml:"OrderID"`           // 	传入的唯一标识订单的 ID，由商户生成，原样返回
	BizID        string `json:"BizID" xml:"BizID"`               // 	商户 ID，原样返回
	WayBillID    string `json:"WayBillID" xml:"WayBillID"`       // 	运单 ID
	ResultCode   int    `json:"ResultCode" xml:"ResultCode"`     // 	处理结果错误码
	ResultMsg    string `json:"ResultMsg" xml:"ResultMsg"`       // 	处理结果的详细信息
	WaybillData  string `json:"WaybillData" xml:"WaybillData"`   // 	集包地、三段码、大头笔等信息，用于生成面单信息。详见后文返回值说明
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
		Used  InsureStatus `json:"UseInsured" xml:"UseInsured"`     // 是否保价，0 表示不保价，1 表示保价
		Value uint         `json:"InsuredValue" xml:"InsuredValue"` // 保价金额，单位是分，比如: 10000 表示 100 元
	} `json:"Insured" xml:"Insured"` // 保价信息
	Service struct {
		Type uint8  `json:"ServiceType" xml:"ServiceType"` // 服务类型 ID
		Name string `json:"ServiceName" xml:"ServiceName"` // 服务名称
	} `json:"Service" xml:"Service"` // 服务类型
}

// GetQuotaReturn 查询商户余额事件返回数据
type GetQuotaReturn struct {
	ToUserName   string  `json:"ToUserName" xml:"ToUserName"`     // 	原样返回请求中的 FromUserName
	FromUserName string  `json:"FromUserName" xml:"FromUserName"` // 	快递公司小程序 UserName
	CreateTime   uint    `json:"CreateTime" xml:"CreateTime"`     // 	事件时间，Unix时间戳
	MsgType      string  `json:"MsgType" xml:"MsgType"`           // 	消息类型，固定为event
	Event        string  `json:"Event" xml:"Event"`               // 	事件类型
	BizID        string  `json:"BizID" xml:"BizID"`               // 	商户ID
	ResultCode   int     `json:"ResultCode" xml:"ResultCode"`     // 	处理结果错误码
	ResultMsg    string  `json:"ResultMsg" xml:"ResultMsg"`       // 	处理结果详情
	Quota        float64 `json:"Quota" xml:"Quota"`               // 	商户可用余额，0 表示无可用余额
}

// GetQuotaResult 查询商户余额事件参数
type GetQuotaResult struct {
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

// CancelExpressOrderReturn 取消订单事件返回数据
type CancelExpressOrderReturn struct {
	ToUserName   string `json:"ToUserName" xml:"ToUserName"`     // 原样返回请求中的 FromUserName
	FromUserName string `json:"FromUserName" xml:"FromUserName"` // 快递公司小程序 UserName
	CreateTime   uint   `json:"CreateTime" xml:"CreateTime"`     // 事件时间，Unix 时间戳
	MsgType      string `json:"MsgType" xml:"MsgType"`           // 消息类型，固定为 event
	Event        string `json:"Event" xml:"Event"`               // 事件类型，固定为 cancel_waybill，不区分大小写
	BizID        string `json:"BizID" xml:"BizID"`               // 商户ID，请原样返回
	OrderID      string `json:"OrderID" xml:"OrderID"`           // 唯一标识订单的ID，由商户生成。请原样返回
	WayBillID    string `json:"WayBillID" xml:"WayBillID"`       // 运单ID，请原样返回
	ResultCode   int    `json:"ResultCode" xml:"ResultCode"`     // 处理结果错误码
	ResultMsg    string `json:"ResultMsg" xml:"ResultMsg"`       // 处理结果详情
}

// CheckBusinessResult 审核商户事件参数
type CheckBusinessResult struct {
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

// CheckBusinessReturn 审核商户事件返回数据
type CheckBusinessReturn struct {
	ToUserName   string  `json:"ToUserName" xml:"ToUserName"`     //	原样返回请求中的 FromUserName
	FromUserName string  `json:"FromUserName" xml:"FromUserName"` //	快递公司小程序 UserName
	CreateTime   uint    `json:"CreateTime" xml:"CreateTime"`     //	事件时间，Unix时间戳
	MsgType      string  `json:"MsgType" xml:"MsgType"`           //	消息类型，固定为event
	Event        string  `json:"Event" xml:"Event"`               //	事件类型，固定为check_biz，不区分大小写
	BizID        string  `json:"BizID" xml:"BizID"`               //	商户ID
	ResultCode   int     `json:"ResultCode" xml:"ResultCode"`     //	处理结果错误码
	ResultMsg    string  `json:"ResultMsg" xml:"ResultMsg"`       //	处理结果详情
	Quota        float64 `json:"Quota" xml:"Quota"`               //	商户可用余额，0 表示无可用余额
}
