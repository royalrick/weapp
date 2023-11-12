package operation

import "github.com/medivhzhan/weapp/v3/request"

const apiGetJsErrList = "/wxaapi/log/jserr_list"

type GetJsErrListRequest struct {
	// 必填	小程序版本 "0"代表全部， 例如：“2.0.18”
	AppVersion string `json:"appVersion"`
	// 必填	错误类型 "0"【全部】，"1"【业务代码错误】，"2"【插件错误】，"3"【系统框架错误】
	ErrType string `json:"errType"`
	// 必填	开始时间， 格式 "xxxx-xx-xx"
	StartTime string `json:"startTime"`
	// 必填	结束时间，格式 “xxxx-xx-xx”
	EndTime string `json:"endTime"`
	// 必填	从错误中搜索关键词，关键词过滤
	Keyword string `json:"keyword"`
	// 必填	发生错误的用户 openId
	Openid string `json:"openid"`
	// 必填	排序字段 "uv", "pv" 二选一
	Orderby string `json:"orderby"`
	// 必填	排序规则 "1" orderby字段降序，"2" orderby字段升序
	Desc string `json:"desc"`
	// 必填	分页起始值
	Offset int64 `json:"offset"`
	// 必填	一次拉取最大值， 最大 30
	Limit int64 `json:"limit"`
}

// 错误查询列表
func (cli *Operation) GetJsErrList(req *GetJsErrListRequest) (*request.CommonError, error) {

	uri, err := cli.combineURI(apiGetJsErrList, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(uri, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
