package search

import "github.com/medivhzhan/weapp/v3/request"

const apiImageSearch = "/wxa/imagesearch"

type CreateActivityIDResponse struct {
	request.CommonError
	// 生成的小程序 URL Link
	URLLink string `json:"url_link"`
}

// 本接口提供基于小程序的站内搜商品图片搜索能力
func (cli *Search) ImageSearch(filename string) (*CreateActivityIDResponse, error) {

	url, err := cli.combineURI(apiImageSearch, nil, true)
	if err != nil {
		return nil, err
	}

	rsp := new(CreateActivityIDResponse)
	if err := cli.request.FormPostWithFile(url, "img", filename, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
