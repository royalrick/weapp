package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetPushUrl = "/wxaapi/broadcast/room/getpushurl"

type GetPushUrlRequest struct {
	// 必填	直播间id
	RoomId int64 `query:"roomId"`
}

type GetPushUrlResponse struct {
	request.CommonError
	// 直播间推流地址
	PushAddr string `json:"pushAddr"`
}

// 获取直播间推流地址
func (cli *LiveBroadcast) GetPushUrl(req *GetPushUrlRequest) (*GetPushUrlResponse, error) {

	api, err := cli.combineURI(apiGetPushUrl, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetPushUrlResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
