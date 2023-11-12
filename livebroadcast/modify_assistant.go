package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiModifyAssistant = "/wxaapi/broadcast/room/modifyassistant"

type ModifyAssistantRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
	// 必填	用户微信号
	Username string `json:"username"`
	// 必填	用户微信昵称
	Nickname string `json:"nickname"`
}

type ModifyAssistantResponse struct {
	request.CommonError
}

// 修改管理直播间小助手
func (cli *LiveBroadcast) ModifyAssistant(req *ModifyAssistantRequest) (*ModifyAssistantResponse, error) {

	api, err := cli.combineURI(apiModifyAssistant, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(ModifyAssistantResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
