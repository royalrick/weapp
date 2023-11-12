package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetSharedCode = "/wxaapi/broadcast/room/GetSharedCode"

type GetSharedCodeRequest struct {
	//	房间ID
	RoomId int64 `query:"roomId"`
	// 自定义参数
	Params int64 `query:"params"`
}

type GetSharedCodeResponse struct {
	request.CommonError
	// 分享二维码地址
	CdnUrl string `json:"cdnUrl"`
	// 分享路径
	PagePath string `json:"pagePath"`
	// 分享海报地址
	PosterUrl string `json:"posterUrl"`
}

// 获取直播间分享二维码
func (cli *LiveBroadcast) GetSharedCode(req *GetSharedCodeRequest) (*GetSharedCodeResponse, error) {

	api, err := cli.combineURI(apiGetSharedCode, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetSharedCodeResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
