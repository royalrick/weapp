package operation

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetGrayReleasePlan = "/wxa/getgrayreleaseplan"

type GetGrayReleasePlanResponse struct {
	request.CommonError
	// 分阶段发布计划详情
	GrayReleaseplan struct {
		//	0:初始状态 1:执行中 2:暂停中 3:执行完毕 4:被删除
		Status int `json:"status"`
		//	分阶段发布计划的创建事件
		CreateTimestamp int `json:"create_timestamp"`
		//	当前的灰度比例
		GrayPercentage int `json:"gray_percentage"`
		//	预计全量时间
		DefaultFinishTimestamp int `json:"default_finish_timestamp"`
	} `json:"gray_release_plan"`
}

// 查询当前分阶段发布详情
func (cli *Operation) GetGrayReleasePlan() (*GetGrayReleasePlanResponse, error) {

	uri, err := cli.combineURI(apiGetGrayReleasePlan, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetGrayReleasePlanResponse)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
