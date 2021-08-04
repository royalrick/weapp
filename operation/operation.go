package operation

import (
	"github.com/medivhzhan/weapp/v3/request"
)

type Operation struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}) (string, error)
}

func NewOperation(request *request.Request, conbineURI func(url string, req interface{}) (string, error)) *Operation {
	sm := Operation{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
