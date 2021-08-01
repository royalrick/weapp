package server

// MsgType 消息类型
type MsgType = string

// 所有消息类型
const (
	MsgText  MsgType = "text"                      // 文本消息类型
	MsgImg   MsgType = "image"                     // 图片消息类型
	MsgCard  MsgType = "miniprogrampage"           // 小程序卡片消息类型
	MsgEvent MsgType = "event"                     // 事件类型
	MsgTrans MsgType = "transfer_customer_service" // 转发客服消息
)
