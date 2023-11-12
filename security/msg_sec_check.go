package security

import "github.com/medivhzhan/weapp/v3/request"

const apiMsgSecCheck = "/wxa/msg_sec_check"

type MsgSecCheckRequest struct {
	// 必填	接口版本号，2.0版本为固定值2
	Version uint8 `json:"version"`
	// 必填	用户的openid（用户需在近两小时访问过小程序）
	Openid string `json:"openid"`
	// 必填	场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
	Scene uint8 `json:"scene"`
	// 必填	需检测的文本内容，文本字数的上限为2500字
	Content string `json:"content"`
	// 非必填	用户昵称
	Nickname string `json:"nickname"`
	// 非必填	文本标题
	Title string `json:"title"`
	// 非必填	个性签名，该参数仅在资料类场景有效(scene=1)
	Signature string `json:"signature"`
}

type MsgSecCheckResponse struct {
	request.CommonError
	// 唯一请求标识，标记单次请求
	TraceId string `json:"trace_id"`
	// 综合结果
	Result struct {
		// 建议，有risky、pass、review三种值
		Suggest string `json:"suggest"`
		// 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
		Label int `json:"label"`
	} `json:"result"`
	// 详细检测结果
	Detail []struct {
		// 策略类型
		Strategy string `json:"strategy"`
		// 错误码，仅当该值为0时，该项结果有效
		Errcode int `json:"errcode"`
		// 建议，有risky、pass、review三种值
		Suggest string `json:"suggest"`
		// 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
		Label int `json:"label"`
		// 0-100，代表置信度，越高代表越有可能属于当前返回的标签（label）
		Prob int `json:"prob"`
		// 命中的自定义关键词
		Keyword string `json:"keyword"`
	} `json:"detail"`
}

// 检查一段文本是否含有违法违规内容。
//
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html
func (cli *Security) MsgSecCheck(req *MsgSecCheckRequest) (*MsgSecCheckResponse, error) {
	url, err := cli.combineURI(apiMsgSecCheck, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(MsgSecCheckResponse)
	if err := cli.request.Post(url, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
