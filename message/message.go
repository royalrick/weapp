// Package message 消息
package message

// MsgType 消息类型
type MsgType string

const (
	// Text 文本消息类型
	Text MsgType = "text"

	// Img 图片消息类型
	Img MsgType = "image"

	// Link 图文链接消息类型
	Link MsgType = "link"

	// Card 小程序卡片消息类型
	Card MsgType = "miniprogrampage"

	// Event 事件类型
	Event MsgType = "event"
)

// EventType 事件类型
type EventType string

const (
	// Enter 用户进入临时会话状态
	Enter EventType = "user_enter_tempsession"
)
