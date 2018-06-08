package notify

// EventType 事件类型
type EventType string

const (
	// UserEnterEvent 用户进入临时会话状态
	UserEnterEvent EventType = "user_enter_tempsession"
)
