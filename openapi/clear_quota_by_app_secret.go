package openapi

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiClearQuotaByAppSecret = "/cgi-bin/clear_quota/v2"

type ClearQuotaByAppSecretRequest struct {
	// 必填 要被清空的账号的appid
	Appid string `query:"appid"`
	// 唯一凭证密钥，即 AppSecret，获取方式同 appid
	AppSecret string `query:"appsecret"`
}

type ClearQuotaByAppSecretResponse struct {
	request.CommonError
}

// 使用AppSecret重置 API 调用次数
// doc: https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/clearQuotaByAppSecret.html
func (cli *OpenApi) ClearQuotaByAppSecret(req *ClearQuotaByAppSecretRequest) (*ClearQuotaByAppSecretResponse, error) {

	uri, err := cli.combineURI(apiClearQuotaByAppSecret, req, true)
	if err != nil {
		return nil, err
	}

	res := new(ClearQuotaByAppSecretResponse)
	if err := cli.request.Post(uri, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}
