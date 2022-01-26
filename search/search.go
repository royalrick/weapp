package search

import "github.com/medivhzhan/weapp/v3/request"

type Search struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewSearch(request *request.Request, conbineURI func(url string, req interface{}, withToken bool) (string, error)) *Search {
	sm := Search{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
