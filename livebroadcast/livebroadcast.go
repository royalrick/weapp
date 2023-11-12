package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

type LiveBroadcast struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	combineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewLiveBroadcast(request *request.Request, combineURI func(url string, req interface{}, withToken bool) (string, error)) *LiveBroadcast {
	sm := LiveBroadcast{
		request:    request,
		combineURI: combineURI,
	}

	return &sm
}
