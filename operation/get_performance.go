package operation

import "github.com/medivhzhan/weapp/v3/request"

const apiGetPerformance = "/wxaapi/log/get_performance"

type GetPerformanceRequest struct {
	// 必填 可选值 1（启动总耗时）， 2（下载耗时），3（初次渲染耗时）
	CostTimeType int `json:"cost_time_type"`
	// 必填 查询开始时间
	DefaultStartTime int `json:"default_start_time"`
	// 必填 查询结束时间
	DefaultEndTime int `json:"default_end_time"`
	// 必填 系统平台，可选值 "@_all:"（全部），1（IOS）， 2（android）
	Device string `json:"device"`
	// 必填 是否下载代码包，当 type 为 1 的时候才生效，可选值 "@_all:"（全部），1（是）， 2（否）
	IsDownloadCode string `json:"is_download_code"`
	// 必填 访问来源，当 type 为 1 或者 2 的时候才生效，通过 getSceneList 接口获取
	Scene string `json:"scene"`
	// 必填 网络环境, 当 type 为 2 的时候才生效，可选值 "@_all:"，wifi, 4g, 3g, 2g
	Networktype string `json:"networktype"`
}

type GetPerformanceResponse struct {
	request.CommonError
	//	错误查询数据(json字符串，结构如下所述的 strbody)
	DefaultTimeData string `json:"default_time_data"`
	//	比较数据
	CompareTimeData string `json:"compare_time_data"`
}

// 性能监控
func (cli *Operation) GetPerformance(req *GetPerformanceRequest) (*GetPerformanceResponse, error) {

	uri, err := cli.combineURI(apiGetPerformance, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetPerformanceResponse)
	if err := cli.request.Post(uri, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
