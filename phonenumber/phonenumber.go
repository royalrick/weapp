package phonenumber

import "github.com/medivhzhan/weapp/v3/request"

type Phonenumber struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	conbineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewPhonenumber(request *request.Request, conbineURI func(url string, req interface{}, withToken bool) (string, error)) *Phonenumber {
	sm := Phonenumber{
		request:    request,
		conbineURI: conbineURI,
	}

	return &sm
}
