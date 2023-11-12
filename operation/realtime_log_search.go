package operation

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiRealtimelogSearch = "/wxaapi/userlog/userlog_search"

type RealtimelogSearchRequest struct {
	// 必填 YYYYMMDD格式的日期，仅支持最近7天
	Date string `query:"date"`
	// 必填 开始时间，必须是date指定日期的时间
	Begintime int64 `query:"begintime"`
	// 必填 结束时间，必须是date指定日期的时间
	Endtime int64 `query:"endtime"`
	// 非必填 开始返回的数据下标，用作分页，默认为0
	Start int `query:"start"`
	// 非必填 返回的数据条数，用作分页，默认为20
	Limit int `query:"limit"`
	// 非必填 小程序启动的唯一ID，按TraceId查询会展示该次小程序启动过程的所有页面的日志。
	TraceId string `query:"traceId"`
	// 非必填 小程序页面路径，例如pages/index/index
	Url string `query:"url"`
	// 非必填 用户微信号或者OpenId
	Id string `query:"id"`
	// 非必填 开发者通过setFileterMsg/addFilterMsg指定的filterMsg字段
	FilterMsg string `query:"filterMsg"`
	// 非必填 日志等级，返回大于等于level等级的日志，level的定义为2（Info）、4（Warn）、8（Error），如果指定为4，则返回大于等于4的日志，即返回Warn和Error日志。
	Level uint8 `query:"level"`
}

type RealtimelogSearchResponse struct {
	request.CommonError
	// 返回的日志数据列表
	List []struct {
		// 日志等级，是msg数组里面的所有level字段的或操作得到的结果。例如msg数组里有两条日志，Info（值为2）和Warn（值为4），则level值为6
		Level uint8 `json:"level"`
		// 基础库版本
		LibraryVersion string `json:"libraryVersion"`
		// 客户端版本
		ClientVersion string `json:"clientVersion"`
		// 微信用户OpenID
		Id string `json:"id"`
		// 打日志的Unix时间戳
		Timestamp int64 `json:"timestamp"`
		// 1 安卓 2 IOS
		Platform uint8 `json:"platform"`
		// 小程序页面链接
		Url string `json:"url"`
		// 日志内容数组，log.info等的内容存在这里
		Msg []struct {
			// log.info调用的时间
			Time int64 `json:"time"`
			// log.info调用的内容，每个参数分别是数组的一项
			Msg []string `json:"msg"`
			// log.info调用的日志等级
			Level uint8 `json:"level"`
		} `json:"msg"`
		// 小程序启动的唯一ID，按TraceId查询会展示该次小程序启动过程的所有页面的日志。
		TraceId string `json:"traceid"`
		// 微信用户OpenID
		FilterMsg string `json:"filterMsg"`
	} `json:"list"`
}

// 实时日志查询
func (cli *Operation) RealtimelogSearch(req *RealtimelogSearchRequest) (*RealtimelogSearchResponse, error) {

	uri, err := cli.combineURI(apiRealtimelogSearch, req, true)
	if err != nil {
		return nil, err
	}

	res := new(RealtimelogSearchResponse)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
