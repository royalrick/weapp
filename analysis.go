package weapp

import "github.com/medivhzhan/weapp/v3/request"

type dateRange struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

const (
	apiGetUserPortrait      = "/datacube/getweanalysisappiduserportrait"
	apiGetVisitDistribution = "/datacube/getweanalysisappidvisitdistribution"
	apiGetVisitPage         = "/datacube/getweanalysisappidvisitpage"
	apiGetDailySummary      = "/datacube/getweanalysisappiddailysummarytrend"
)

// UserPortrait response data of get user portrait
type UserPortrait struct {
	request.CommonError
	RefDate    string   `json:"ref_date"`
	VisitUV    Portrait `json:"visit_uv"`     // 活跃用户画像
	VisitUVNew Portrait `json:"visit_uv_new"` // 新用户画像
}

// Portrait 肖像
type Portrait struct {
	Index     uint        `json:"index"`     // 分布类型
	Province  []Attribute `json:"province"`  // 省份，如北京、广东等
	City      []Attribute `json:"city"`      // 城市，如北京、广州等
	Genders   []Attribute `json:"genders"`   // 性别，包括男、女、未知
	Platforms []Attribute `json:"platforms"` // 终端类型，包括 iPhone，android，其他
	Devices   []Attribute `json:"devices"`   // 机型，如苹果 iPhone 6，OPPO R9 等
	Ages      []Attribute `json:"ages"`      // 年龄，包括17岁以下、18-24岁等区间
}

// Attribute 描述内容
type Attribute struct {
	ID    uint   `json:"id"`   // 属性值id
	Name  string `json:"name"` // 属性值名称，与id对应。如属性为 province 时，返回的属性值名称包括「广东」等。
	Value uint   `json:"value"`
	// TODO: 确认后删除该字段
	AccessSourceVisitUV uint `json:"access_source_visit_uv"` // 该场景访问uv
}

// GetUserPortrait 获取小程序新增或活跃用户的画像分布数据。
// 时间范围支持昨天、最近7天、最近30天。
// 其中，新增用户数为时间范围内首次访问小程序的去重用户数，活跃用户数为时间范围内访问过小程序的去重用户数。
// begin 开始日期。格式为 yyyymmdd
// end 结束日期，开始日期与结束日期相差的天数限定为0/6/29，分别表示查询最近1/7/30天数据，允许设置的最大值为昨日。格式为 yyyymmdd
func (cli *Client) GetUserPortrait(begin, end string) (*UserPortrait, error) {
	api := baseURL + apiGetUserPortrait

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getUserPortrait(token, begin, end, api)
}

func (cli *Client) getUserPortrait(token, begin, end, api string) (*UserPortrait, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := dateRange{
		BeginDate: begin,
		EndDate:   end,
	}

	res := new(UserPortrait)
	if err := cli.request.Post(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// VisitDistribution 用户小程序访问分布数据
type VisitDistribution struct {
	request.CommonError
	RefDate string         `json:"ref_date"`
	List    []Distribution `json:"list"`
}

// Distribution 分布数据
type Distribution struct {
	// 分布类型
	// 	index 的合法值
	// access_source_session_cnt	访问来源分布
	// access_staytime_info	访问时长分布
	// access_depth_info	访问深度的分布
	Index    string             `json:"index"`
	ItemList []DistributionItem `json:"item_list"` // 分布数据列表
}

// DistributionItem 分布数据项
type DistributionItem struct {
	Key   uint `json:"key"`   // 场景 id，定义在各个 index 下不同，具体参见下方表格
	Value uint `json:"value"` // 该场景 id 访问 pv
	// TODO: 确认后删除该字段
	AccessSourceVisitUV uint `json:"access_source_visit_uv"` // 该场景 id 访问 uv
}

// GetVisitDistribution 获取用户小程序访问分布数据
// begin 开始日期。格式为 yyyymmdd
// end 结束日期，限定查询 1 天数据，允许设置的最大值为昨日。格式为 yyyymmdd
func (cli *Client) GetVisitDistribution(begin, end string) (*VisitDistribution, error) {
	api := baseURL + apiGetVisitDistribution

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getVisitDistribution(token, begin, end, api)
}

func (cli *Client) getVisitDistribution(token, begin, end, api string) (*VisitDistribution, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := dateRange{
		BeginDate: begin,
		EndDate:   end,
	}

	res := new(VisitDistribution)
	if err := cli.request.Post(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// VisitPage 页面访问数据
type VisitPage struct {
	request.CommonError
	RefDate string `json:"ref_date"`
	List    []Page `json:"list"`
}

// Page 页面
type Page struct {
	PagePath       string  `json:"Page_path"`        // 页面路径
	PageVisitPV    uint    `json:"Page_visit_pv"`    // 访问次数
	PageVisitUV    uint    `json:"Page_visit_uv"`    // 访问人数
	PageStaytimePV float64 `json:"page_staytime_pv"` // 次均停留时长
	EntrypagePV    uint    `json:"entrypage_pv"`     // 进入页次数
	ExitpagePV     uint    `json:"exitpage_pv"`      // 退出页次数
	PageSharePV    uint    `json:"page_share_pv"`    // 转发次数
	PageShareUV    uint    `json:"page_share_uv"`    // 转发人数

}

// GetVisitPage 访问页面。
// 目前只提供按 page_visit_pv 排序的 top200。
// begin 开始日期。格式为 yyyymmdd
// end 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
func (cli *Client) GetVisitPage(begin, end string) (*VisitPage, error) {
	api := baseURL + apiGetVisitPage

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getVisitPage(token, begin, end, api)
}

func (cli *Client) getVisitPage(token, begin, end, api string) (*VisitPage, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := dateRange{
		BeginDate: begin,
		EndDate:   end,
	}

	res := new(VisitPage)
	if err := cli.request.Post(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DailySummary 用户访问小程序数据概况
type DailySummary struct {
	request.CommonError
	List []Summary `json:"list"`
}

// Summary 概况
type Summary struct {
	RefDate    string `json:"ref_date"`    // 	日期，格式为 yyyymmdd
	VisitTotal uint   `json:"visit_total"` //	累计用户数
	SharePV    uint   `json:"share_pv"`    //	转发次数
	ShareUV    uint   `json:"share_uv"`    //	转发人数
}

// GetDailySummary 获取用户访问小程序数据概况
// begin 开始日期。格式为 yyyymmdd
// end 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
func (cli *Client) GetDailySummary(begin, end string) (*DailySummary, error) {
	api := baseURL + apiGetDailySummary

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getDailySummary(token, begin, end, api)
}

func (cli *Client) getDailySummary(token, begin, end, api string) (*DailySummary, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := dateRange{
		BeginDate: begin,
		EndDate:   end,
	}

	res := new(DailySummary)
	if err := cli.request.Post(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
