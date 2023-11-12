package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiModifySubAnchor = "/wxaapi/broadcast/room/modifysubanchor"

type ModifySubAnchorRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
	// 必填	用户微信号
	Username string `json:"username"`
}

type ModifySubAnchorResponse struct {
	request.CommonError
}

// 修改主播副号
func (cli *LiveBroadcast) ModifySubAnchor(req *ModifySubAnchorRequest) (*ModifySubAnchorResponse, error) {

	api, err := cli.combineURI(apiModifySubAnchor, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(ModifySubAnchorResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
