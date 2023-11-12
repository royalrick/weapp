package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiDeleteRoom = "wxaapi/broadcast/room/deleteroom"

type DeleteRoomRequest struct {
	//	房间ID
	Id int64 `json:"id"`
}

// 删除直播间
func (cli *LiveBroadcast) DeleteRoom(req *DeleteRoomRequest) (*request.CommonError, error) {

	api, err := cli.combineURI(apiDeleteRoom, nil, true)
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
