package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiRemoveAssistant = "/wxaapi/broadcast/room/removeassistant"

type RemoveAssistantRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"room_id"`
	// 必填	用户微信号
	Username string `json:"username"`
}

type RemoveAssistantResponse struct {
	request.CommonError
}

// 删除管理直播间小助手
func (cli *LiveBroadcast) RemoveAssistant(req *RemoveAssistantRequest) (*RemoveAssistantResponse, error) {

	api, err := cli.combineURI(apiRemoveAssistant, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(RemoveAssistantResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
