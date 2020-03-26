package weapp

const (
	apiSearchSubmitPages = "/wxa/search/wxaapi_submitpages"
)

// SearchSubmitPages 小程序页面收录请求
type SearchSubmitPages struct {
	Pages []SearchSubmitPage `json:"pages"`
}

// SearchSubmitPage 请求收录的页面
type SearchSubmitPage struct {
	Path  string `json:"path"`
	Query string `json:"query"`
}

// Send 提交收录请求
func (s *SearchSubmitPages) Send(token string) (*CommonError, error) {
	return s.send(baseURL+apiSearchSubmitPages, token)
}

func (s *SearchSubmitPages) send(api, token string) (*CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postJSON(api, s, res); err != nil {
		return nil, err
	}

	return res, nil
}
