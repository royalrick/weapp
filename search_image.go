package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiImageSearch = "/wxa/imagesearch"
)

type SearchImageResponse struct {
	request.CommonError
	// 生成的小程序 URL Link
	URLLink string `json:"url_link"`
}

// 本接口提供基于小程序的站内搜商品图片搜索能力
func (cli *Client) SearchImage(filename string) (*SearchImageResponse, error) {
	api := baseURL + apiImageSearch

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.searchImage(api, token, filename)
}

func (cli *Client) searchImage(api, token, filename string) (*SearchImageResponse, error) {
	uri, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(SearchImageResponse)
	if err := cli.request.FormPostWithFile(uri, "img", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}
