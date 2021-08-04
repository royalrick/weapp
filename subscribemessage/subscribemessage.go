package subscribemessage

import "github.com/medivhzhan/weapp/v3/request"

type SubscribeMessage struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}) (string, error)
}

func NewSubscribeMessage(request *request.Request, conbineURI func(url string, req interface{}) (string, error)) *SubscribeMessage {
	sm := SubscribeMessage{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
