package openapi

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetApiQuota = "/cgi-bin/openapi/quota/get"

type GetApiQuotaRequest struct {
	// @必填
	// api的请求地址，例如"/cgi-bin/message/custom/send";不要前缀“https://api.weixin.qq.com” ，也不要漏了"/",否则都会76003的报错
	CgiPath string `json:"cgi_path"`
}

type GetApiQuotaResponse struct {
	request.CommonError
	Quota struct {
		DailyLimit int64 `json:"daily_limit"` // 当天该账号可调用该接口的次数
		Used       int64 `json:"used"`        // 当天已经调用的次数
		Remain     int64 `json:"remain"`      // 当天剩余调用次数
		RateLimit  struct {
			CallCount     int64 `json:"call_count"`     // 周期内可调用数量，单位 次
			RefreshSecond int64 `json:"refresh_second"` // 更新周期，单位 秒
		} `json:"rate_limit"` // 普通调用频率限制
		ComponentRateLimit struct {
			CallCount     int64 `json:"call_count"`     // 周期内可调用数量，单位 次
			RefreshSecond int64 `json:"refresh_second"` // 更新周期，单位 秒
		} `json:"component_rate_limit"` // 代调用频率限制
	} `json:"quota"` // quota详情
}

// 查询API调用额度
// doc: https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/getApiQuota.html
func (cli *OpenApi) GetApiQuota(req *GetApiQuotaRequest) (*GetApiQuotaResponse, error) {

	uri, err := cli.combineURI(apiGetApiQuota, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetApiQuotaResponse)
	if err := cli.request.Post(uri, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
