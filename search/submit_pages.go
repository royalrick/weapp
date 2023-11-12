package search

import "github.com/medivhzhan/weapp/v3/request"

const apiSubmitPages = "/wxa/search/wxaapi_submitpages"

type SubmitPagesRequest struct {
	Pages []SubmitPage `json:"pages"`
}

type SubmitPage struct {
	Path  string `json:"path"`
	Query string `json:"query"`
}

// 小程序开发者可以通过本接口提交小程序页面url及参数信息(不要推送webview页面)，让微信可以更及时的收录到小程序的页面信息，开发者提交的页面信息将可能被用于小程序搜索结果展示。
func (cli *Search) SubmitPages(req *SubmitPagesRequest) (*request.CommonError, error) {

	api, err := cli.combineURI(apiSubmitPages, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
