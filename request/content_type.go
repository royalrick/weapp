package request

// 消息返回数据类型
type ContentType uint

const (
	Plaintext ContentType = iota
	ContentTypeXML
	ContentTypeJSON
)

func (ctp ContentType) String() string {
	switch ctp {
	case ContentTypeXML:
		return "application/xml"
	case ContentTypeJSON:
		return "application/json"
	default:
		return "plain/text"
	}
}
