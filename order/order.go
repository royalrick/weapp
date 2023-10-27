package order

import "github.com/medivhzhan/weapp/v3/request"

type Order struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewOrder(request *request.Request, conbineURI func(url string, req interface{}, withToken bool) (string, error)) *Order {
	return &Order{
		request:    request,
		conbineURI: conbineURI,
	}
}
