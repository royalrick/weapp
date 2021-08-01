package request

// 消息返回数据类型
type ContentType uint

const (
	ContentTypePlain ContentType = iota
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
		return "text/plain"
	}
}
