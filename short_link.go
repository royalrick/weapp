package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiShortLink = "/wxa/genwxashortlink"
)

type ShortLinkRequest struct {
	// 必填	通过 Short Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，可携带 query，最大1024个字符
	PageURL string `json:"page_url"`
	// 必填	页面标题，不能包含违法信息，超过20字符会用... 截断代替
	PageTitle string `json:"page_title"`
	// 非必填	生成的 Short Link 类型，短期有效：false，永久有效：true
	IsPermanent bool `json:"is_permanent"`
}

type ShortLinkResponse struct {
	request.CommonError
	// 生成的小程序 Short Link
	Link string `json:"link"`
}

// 获取小程序 Short Link，适用于微信内拉起小程序的业务场景。目前只开放给电商类目(具体包含以下一级类目：电商平台、商家自营、跨境电商)。通过该接口，可以选择生成到期失效和永久有效的小程序短链
func (cli *Client) GenerateShortLink(link *ShortLinkRequest) (*ShortLinkResponse, error) {
	api := baseURL + apiShortLink

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.generateShortLink(api, token, link)
}

func (cli *Client) generateShortLink(api, token string, link *ShortLinkRequest) (*ShortLinkResponse, error) {
	uri, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(ShortLinkResponse)
	err = cli.request.Post(uri, link, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
