package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiUpdateReplay = "/wxaapi/broadcast/room/updatereplay"

type UpdateReplayRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
	// 必填	是否关闭回放 【0：开启，1：关闭】
	CloseReplay uint8 `json:"closeReplay"`
}

type UpdateReplayResponse struct {
	request.CommonError
}

// 开开启/关闭回放功能
func (cli *LiveBroadcast) UpdateReplay(req *UpdateReplayRequest) (*UpdateReplayResponse, error) {

	api, err := cli.combineURI(apiUpdateReplay, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(UpdateReplayResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
