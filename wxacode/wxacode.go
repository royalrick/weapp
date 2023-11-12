package wxacode

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp/v3/request"
)

type WXACode struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	combineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewWXACode(request *request.Request, combineURI func(url string, req interface{}, withToken bool) (string, error)) *WXACode {
	sm := WXACode{
		request:    request,
		combineURI: combineURI,
	}

	return &sm
}

// 生成二维码
func (cli *WXACode) generate(api string, params interface{}) (*http.Response, *request.CommonError, error) {

	url, err := cli.combineURI(api, nil, true)
	if err != nil {
		return nil, nil, err
	}

	res, err := cli.request.PostWithBody(url, params)
	if err != nil {
		return nil, nil, err
	}

	response := new(request.CommonError)
	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		if err := json.NewDecoder(res.Body).Decode(response); err != nil {
			res.Body.Close()
			return nil, nil, err
		}
		return res, response, nil

	case strings.HasPrefix(header, "image"): // 返回文件
		return res, response, nil

	default:
		res.Body.Close()
		return nil, nil, errors.New("invalid response header: " + header)
	}
}
