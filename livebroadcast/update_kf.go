package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiUpdateKF = "/wxaapi/broadcast/room/updatekf"

type UpdateKFRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
	// 必填	是否关闭客服 【0：开启，1：关闭】
	CloseKf uint8 `json:"closeKf"`
}

type UpdateKFResponse struct {
	request.CommonError
}

// 开启/关闭客服功能
func (cli *LiveBroadcast) UpdateKF(req *UpdateKFRequest) (*UpdateKFResponse, error) {

	api, err := cli.combineURI(apiUpdateKF, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(UpdateKFResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
