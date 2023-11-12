package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiUpdateComment = "/wxaapi/broadcast/room/updatecomment"

type UpdateCommentRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
	// 必填	1-禁言，0-取消禁言
	BanComment uint8 `json:"banComment"`
}

type UpdateCommentResponse struct {
	request.CommonError
}

// 开启/关闭直播间全局禁言
func (cli *LiveBroadcast) UpdateComment(req *UpdateCommentRequest) (*UpdateCommentResponse, error) {

	api, err := cli.combineURI(apiUpdateComment, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(UpdateCommentResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
