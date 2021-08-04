package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiSearchSubmitPages = "/wxa/search/wxaapi_submitpages"
)

// SearchSubmitPagesRequest 小程序页面收录请求
type SearchSubmitPagesRequest struct {
	Pages []SearchSubmitPage `json:"pages"`
}

// SearchSubmitPage 请求收录的页面
type SearchSubmitPage struct {
	Path  string `json:"path"`
	Query string `json:"query"`
}

// Send 提交收录请求
func (cli *Client) SendSearchSubmitPages(smp *SearchSubmitPagesRequest) (*request.CommonError, error) {
	api := baseURL + apiSearchSubmitPages
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.sendSearchSubmitPages(api, token, smp)
}

func (cli *Client) sendSearchSubmitPages(api, token string, smp *SearchSubmitPagesRequest) (*request.CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(api, smp, res); err != nil {
		return nil, err
	}

	return res, nil
}
