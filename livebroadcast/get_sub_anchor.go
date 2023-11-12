package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetSubAnchor = "/wxaapi/broadcast/room/GetSubAnchor"

type GetSubAnchorRequest struct {
	//	房间ID
	RoomId int64 `query:"roomId"`
}

type GetSubAnchorResponse struct {
	request.CommonError
	Username string `json:"username"`
}

// 获取主播副号
func (cli *LiveBroadcast) GetSubAnchor(req *GetSubAnchorRequest) (*GetSubAnchorResponse, error) {

	api, err := cli.combineURI(apiGetSubAnchor, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetSubAnchorResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
