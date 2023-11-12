package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiPushMessage = "/wxa/business/push_message"

type PushMessageRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"room_id"`
	// 必填	接收该群发开播事件的订阅用户OpenId列表
	UserOpenid []string `json:"user_openid"`
}

type PushMessageResponse struct {
	request.CommonError
	// 此次群发消息的标识ID，用于对应【长期订阅群发结果回调】的message_id
	MessageId string `json:"message_id"`
}

// 向长期订阅用户群发直播间开始事件
func (cli *LiveBroadcast) PushMessage(req *PushMessageRequest) (*PushMessageResponse, error) {

	api, err := cli.combineURI(apiPushMessage, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(PushMessageResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
