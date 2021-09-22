package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apigetSubAnchor = "/wxaapi/broadcast/room/getsubanchor"

type getSubAnchorRequest struct {
	//	房间ID
	RoomId int64 `json:"roomId"`
}

type getSubAnchorResponse struct {
	request.CommonError
	Username string `json:"username"`
}

// 获取主播副号
func (cli *LiveBroadcast) GetSubAnchor(req *getSubAnchorRequest) (*getSubAnchorResponse, error) {

	api, err := cli.conbineURI(apigetSubAnchor, req)
	if err != nil {
		return nil, err
	}

	res := new(getSubAnchorResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
