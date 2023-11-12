package openapi

import (
	"github.com/medivhzhan/weapp/v3/request"
)

type OpenApi struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	combineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewOpenApi(request *request.Request, combineURI func(url string, req interface{}, withToken bool) (string, error)) *OpenApi {
	sm := OpenApi{
		request:    request,
		combineURI: combineURI,
	}

	return &sm
}
