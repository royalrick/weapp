package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

type LiveBroadcast struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewLiveBroadcast(request *request.Request, conbineURI func(url string, req interface{}, withToken bool) (string, error)) *LiveBroadcast {
	sm := LiveBroadcast{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
