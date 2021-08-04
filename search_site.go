package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiSiteSearch = "/wxa/sitesearch"
)

type SearchSiteRequest struct {
	// 必填	关键词
	Keyword string `json:"keyword"`
	// 必填	请求下一页的参数，开发者无需理解。为空时查询的是第一页内容，如需查询下一页，把返回参数的next_page_info填充到这里即可
	NextPageInfo string `json:"next_page_info"`
}

type SearchSiteResponse struct {
	request.CommonError
	// 生成的小程序 URL Link
	URLLink string `json:"url_link"`
}

// 小程序内部搜索API提供针对页面的查询能力，小程序开发者输入搜索词后，将返回自身小程序和搜索词相关的页面。因此，利用该接口，开发者可以查看指定内容的页面被微信平台的收录情况；同时，该接口也可供开发者在小程序内应用，给小程序用户提供搜索能力。
func (cli *Client) SearchSite(req *SearchSiteRequest) (*SearchSiteResponse, error) {
	api := baseURL + apiSiteSearch

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.searchSite(api, token, req)
}

func (cli *Client) searchSite(api, token string, req *SearchSiteRequest) (*SearchSiteResponse, error) {
	uri, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(SearchSiteResponse)
	if err := cli.request.Post(uri, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
