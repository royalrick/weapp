package updatablemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiCreateActivityId = "/cgi-bin/message/wxopen/updatablemsg/send"

type CreateActivityIdRequest struct {
	// 非必填	为私密消息创建activity_id时，指定分享者为unionid用户。其余用户不能用此activity_id分享私密消息。openid与unionid填一个即可。私密消息暂不支持云函数生成activity id。
	Unionid string `query:"unionid"`
	// 非必填	为私密消息创建activity_id时，指定分享者为openid用户。其余用户不能用此activity_id分享私密消息。openid与unionid填一个即可。私密消息暂不支持云函数生成activity id。
	Openid string `query:"openid"`
}

type CreateActivityIDResponse struct {
	request.CommonError
	ActivityID     string `json:"activity_id"`     //	动态消息的 ID
	ExpirationTime uint   `json:"expiration_time"` //	activity_id 的过期时间戳。默认24小时后过期。
}

// 组合模板并添加至帐号下的个人模板库
func (cli *UpdatableMessage) CreateActivityId(req *CreateActivityIdRequest) (*CreateActivityIDResponse, error) {

	api, err := cli.conbineURI(apiCreateActivityId, req)
	if err != nil {
		return nil, err
	}

	res := new(CreateActivityIDResponse)
	err = cli.request.Get(api, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
