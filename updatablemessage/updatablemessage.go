package updatablemessage

import "github.com/medivhzhan/weapp/v3/request"

type UpdatableMessage struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	combineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewUpdatableMessage(request *request.Request, combineURI func(url string, req interface{}, withToken bool) (string, error)) *UpdatableMessage {
	sm := UpdatableMessage{
		request:    request,
		combineURI: combineURI,
	}

	return &sm
}
