package analysis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
)

const (
	getMonthlyVisitTrendAPI = "/datacube/getweanalysisappidmonthlyvisittrend"
	getWeeklyVisitTrendAPI  = "/datacube/getweanalysisappidweeklyvisittrend"
	getDailyVisitTrendAPI   = "/datacube/getweanalysisappiddailyvisittrend"
)

// VisitTrend 用户趋势
type VisitTrend struct {
	RefDate         string  `json:"ref_date"`          // 时间，月格式为 yyyymm | 周格式为 yyyymmdd-yyyymmdd | 天格式为 yyyymmdd
	SessionCNT      uint    `json:"session_cnt"`       // 打开次数（自然月内汇总）
	VisitPV         uint    `json:"visit_pv"`          // 访问次数（自然月内汇总）
	VisitUV         uint    `json:"visit_uv"`          // 访问人数（自然月内去重）
	VisitUVNew      uint    `json:"visit_uv_new"`      // 新用户数（自然月内去重）
	StayTimeUV      float64 `json:"stay_time_uv"`      // 人均停留时长 (浮点型，单位：秒)
	StayTimeSession float64 `json:"stay_time_session"` // 次均停留时长 (浮点型，单位：秒)
	VisitDepth      float64 `json:"visit_depth"`       // 平均访问深度 (浮点型)
}

// GetVisitTrendResponse 生物认证秘钥签名验证请求返回数据
type GetVisitTrendResponse struct {
	weapp.Response
	List []VisitTrend `json:"list"`
}

// GetMonthlyVisitTrend 获取用户访问小程序数据月趋势
// @accessToken 接口调用凭证
// @start 开始日期，为自然月第一天。格式为 yyyymmdd
// @end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
func GetMonthlyVisitTrend(accessToken, start, end string) (*GetVisitTrendResponse, error) {
	return getVisitTrend(weapp.BaseURL+getMonthlyVisitTrendAPI, accessToken, start, end)
}

// GetWeeklyVisitTrend 获取用户访问小程序数据周趋势
// @accessToken 接口调用凭证
// @start 开始日期，为自然月第一天。格式为 yyyymmdd
// @end 结束日期，为周日日期，限定查询一周数据。格式为 yyyymmdd
func GetWeeklyVisitTrend(accessToken, start, end string) (*GetVisitTrendResponse, error) {
	return getVisitTrend(weapp.BaseURL+getWeeklyVisitTrendAPI, accessToken, start, end)
}

// GetDailyVisitTrendAPI 获取用户访问小程序数据日趋势
// @accessToken 接口调用凭证
// @start 开始日期，为自然月第一天。格式为 yyyymmdd
// @end 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
func GetDailyVisitTrendAPI(accessToken, start, end string) (*GetVisitTrendResponse, error) {
	return getVisitTrend(weapp.BaseURL+getDailyVisitTrendAPI, accessToken, start, end)
}

func getVisitTrend(api, accessToken, start, end string) (*GetVisitTrendResponse, error) {
	api, err := util.TokenAPI(api, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"begin_date": start,
		"end_date":   end,
	}

	raw, err := json.Marshal(params)
	fmt.Println(string(raw))
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(api, "application/json", bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := new(GetVisitTrendResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
