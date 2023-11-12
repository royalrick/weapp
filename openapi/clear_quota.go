package openapi

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiClearQuota = "/cgi-bin/clear_quota"

type ClearQuotaRequest struct {
	// 必填 要被清空的账号的appid
	Appid string `json:"appid"`
}

type ClearQuotaResponse struct {
	request.CommonError
}

// 重置API调用次数
// doc: https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/clearQuota.html
func (cli *OpenApi) ClearQuota(req *ClearQuotaRequest) (*ClearQuotaResponse, error) {

	uri, err := cli.combineURI(apiClearQuota, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(ClearQuotaResponse)
	if err := cli.request.Post(uri, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
