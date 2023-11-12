package operation

import "github.com/medivhzhan/weapp/v3/request"

const apiGetJsErrDetail = "/wxaapi/log/jserr_detail"

type GetJsErrDetailRequest struct {
	// 必填	开始时间， 格式 "xxxx-xx-xx"
	StartTime string `json:"startTime"`
	// 必填	结束时间，格式 “xxxx-xx-xx”
	EndTime string `json:"endTime"`
	// 必填	错误列表查询 接口 返回的 errorMsgMd5 字段
	ErrorMsgMd5 string `json:"errorMsgMd5"`
	// 必填	错误列表查询 接口 返回的 errorStackMd5 字段
	ErrorStackMd5 string `json:"errorStackMd5"`
	// 必填	小程序版本 "0"代表全部， 例如：“2.0.18”
	AppVersion string `json:"appVersion"`
	// 必填	基础库版本 "0"表示所有版本，例如 "2.14.1"
	SdkVersion string `json:"sdkVersion"`
	// 必填	系统类型 "0"【全部】，"1" 【安卓】，"2" 【IOS】，"3"【其他】
	OsName string `json:"osName"`
	// 必填	客户端版本 "0"表示所有版本， 例如 "7.0.22"
	ClientVersion string `json:"clientVersion"`
	// 必填	发生错误的用户 openId
	Openid string `json:"openid"`
	// 必填	排序规则 "0" 升序, "1" 降序
	Desc string `json:"desc"`
	// 必填	分页起始值
	Offset int `json:"offset"`
	// 必填	一次拉取最大值
	Limit int `json:"limit"`
}

type GetJsErrDetailResponse struct {
	request.CommonError

	Openid  string `json:"openid"`
	Success bool   `json:"success"`
	// 总条数
	TotalCount int `json:"totalCount"`
	// 错误列表
	Data []struct {
		Count         int    `json:"Count"`
		SdkVersion    string `json:"sdkVersion"`
		ClientVersion string `json:"clientVersion"`
		ErrorStackMd5 string `json:"errorStackMd5"`
		TimeStamp     string `json:"TimeStamp"`
		AppVersion    string `json:"appVersion"`
		ErrorMsgMd5   string `json:"errorMsgMd5"`
		ErrorMsg      string `json:"errorMsg"`
		ErrorStack    string `json:"errorStack"`
		Ds            string `json:"Ds"`
		OsName        string `json:"osName"`
		Openid        string `json:"openid"`
		Pluginversion string `json:"pluginversion"`
		AppId         string `json:"appId"`
		DeviceModel   string `json:"DeviceModel"`
		Source        string `json:"source"`
		Route         string `json:"route"`
		Uin           string `json:"Uin"`
		Nickname      string `json:"nickname"`
	} `json:"data"`
}

// 错误查询详情
func (cli *Operation) GetJsErrDetail(req *GetJsErrDetailRequest) (*GetJsErrDetailResponse, error) {

	uri, err := cli.combineURI(apiGetJsErrDetail, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetJsErrDetailResponse)
	if err := cli.request.Post(uri, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
