package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiUpdateFeedPublic = "/wxaapi/broadcast/room/updatefeedpublic"

type UpdateFeedPublicRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
	// 必填	是否开启官方收录 【1: 开启，0：关闭】
	IsFeedsPublic uint8 `json:"isFeedsPublic"`
}

type UpdateFeedPublicResponse struct {
	request.CommonError
}

// 开启/关闭直播间官方收录
func (cli *LiveBroadcast) UpdateFeedPublic(req *UpdateFeedPublicRequest) (*UpdateFeedPublicResponse, error) {

	api, err := cli.combineURI(apiUpdateFeedPublic, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(UpdateFeedPublicResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
