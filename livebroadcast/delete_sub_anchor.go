package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiDeleteSubAnchor = "/wxaapi/broadcast/room/deletesubanchor"

type DeleteSubAnchorRequest struct {
	// 必填	房间ID
	RoomId int64 `json:"roomId"`
}

// 删除主播副号
func (cli *LiveBroadcast) DeleteSubAnchor(req *DeleteSubAnchorRequest) (*request.CommonError, error) {

	api, err := cli.combineURI(apiDeleteSubAnchor, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
