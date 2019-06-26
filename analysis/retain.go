package analysis

import (
	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
)

const (
	getMonthlyRetainAPI = "/datacube/getweanalysisappidmonthlyretaininfo"
	getWeeklyRetainAPI  = "/datacube/getweanalysisappidweeklyretaininfo"
	getDailyRetainAPI   = "/datacube/getweanalysisappiddailyretaininfo"
)

// Retain 用户留存
type Retain struct {
	Key   uint8 `json:"key"`   // 标识，0开始，表示当月，1表示1月后。key取值分别是：0,1
	Value uint  `json:"value"` // key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
}

// RetainResponse 生物认证秘钥签名验证请求返回数据
type RetainResponse struct {
	weapp.Response
	RefDate    string   `json:"ref_date"`     // 时间，月格式为 yyyymm | 周格式为 yyyymmdd-yyyymmdd | 天格式为 yyyymmdd
	VisitUV    []Retain `json:"visit_uv"`     // 活跃用户留存
	VisitUVNew []Retain `json:"visit_uv_new"` // 新增用户留存
}

// GetMonthlyRetain 获取用户访问小程序月留存
// @accessToken 接口调用凭证
// @start 开始日期，为自然月第一天。格式为 yyyymmdd
// @end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
func GetMonthlyRetain(accessToken, start, end string) (*RetainResponse, error) {
	return getRetain(weapp.BaseURL+getMonthlyRetainAPI, accessToken, start, end)
}

// GetWeeklyRetain 获取用户访问小程序周留存
// @accessToken 接口调用凭证
// @start 开始日期，为自然月第一天。格式为 yyyymmdd
// @end 结束日期，为周日日期，限定查询一周数据。格式为 yyyymmdd
func GetWeeklyRetain(accessToken, start, end string) (*RetainResponse, error) {
	return getRetain(weapp.BaseURL+getWeeklyRetainAPI, accessToken, start, end)
}

// GetDailyRetainAPI 获取用户访问小程序日留存
// @accessToken 接口调用凭证
// @start 开始日期，为自然月第一天。格式为 yyyymmdd
// @end 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
func GetDailyRetainAPI(accessToken, start, end string) (*RetainResponse, error) {
	return getRetain(weapp.BaseURL+getDailyRetainAPI, accessToken, start, end)
}

func getRetain(api, accessToken, start, end string) (*RetainResponse, error) {
	api, err := util.TokenAPI(api, accessToken)
	if err != nil {
		return nil, err
	}

	params := dateRange{
		StartDate: start,
		EndDate:   end,
	}

	res := new(RetainResponse)
	if err := util.PostJSON(api, params, res); err != nil {
		return nil, err
	}

	if res.HasError() {
		return nil, res.CreateError("failed to get retain")
	}

	return res, nil
}
