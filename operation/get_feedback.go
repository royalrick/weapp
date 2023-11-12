package operation

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetFeedback = "/wxaapi/feedback/list"

type FeedbackType = uint8

const (
	FeedbackTypeAll          FeedbackType = iota // 全部类型
	FeedbackTypeCannotOpen                       //无法打开小程序
	FeedbackTypeQuit                             //小程序闪退
	FeedbackTypeLags                             //卡顿
	FeedbackTypeBlackOrWhite                     //黑屏白屏
	FeedbackTypeDead                             //死机
	FeedbackTypeUICross                          //界面错位
	FeedbackTypeLoadSlowly                       //界面加载慢
	FeedbackTypeOther                            //其他异常

)

type GetFeedbackRequest struct {
	// 非必填	反馈的类型，默认拉取全部类型，详细定义见下面
	Type FeedbackType `query:"type"`
	// 必填	分页的页数，从1开始
	Page int `query:"page"`
	// 必填	分页拉取的数据数量
	Num int `query:"num"`
}

// 获取用户反馈列表
func (cli *Operation) GetFeedback(req *GetFeedbackRequest) (*request.CommonError, error) {

	uri, err := cli.combineURI(apiGetFeedback, req, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
