package openapi

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetRidInfo = "/cgi-bin/openapi/rid/get"

type GetRidInfoRequest struct {
	// @必填
	// 调用接口报错返回的rid
	Rid string `json:"rid"`
}

type GetRidInfoResponse struct {
	request.CommonError
	Request struct {
		InvokeTime   int64  `json:"invoke_time"`   // 发起请求的时间戳
		CostInMs     int64  `json:"cost_in_ms"`    // 请求毫秒级耗时
		RequestURL   string `json:"request_url"`   // 请求的URL参数
		RequestBody  string `json:"request_body"`  // post请求的请求参数
		ResponseBody string `json:"response_body"` // 接口请求返回参数
		ClientIP     string `json:"client_ip"`     // 接口请求的客户端ip
	} `json:"request"` // 该rid对应的请求详情
}

// 查询rid信息
// doc: https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/getRidInfo.html
func (cli *OpenApi) GetRidInfo(req *GetRidInfoRequest) (*GetRidInfoResponse, error) {

	uri, err := cli.combineURI(apiGetRidInfo, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetRidInfoResponse)
	if err := cli.request.Post(uri, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
