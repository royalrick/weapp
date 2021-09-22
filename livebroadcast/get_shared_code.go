package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apigetSharedCode = "/wxaapi/broadcast/room/getsharedcode"

type getSharedCodeRequest struct {
	//	房间ID
	RoomId int64 `json:"roomId"`
	// 自定义参数
	Params int64 `json:"params"`
}

type getSharedCodeResponse struct {
	request.CommonError
	// 分享二维码地址
	CdnUrl string `json:"cdnUrl"`
	// 分享路径
	PagePath string `json:"pagePath"`
	// 分享海报地址
	PosterUrl string `json:"posterUrl"`
}

// 获取直播间分享二维码
func (cli *LiveBroadcast) GetSharedCode(req *getSharedCodeRequest) (*getSharedCodeResponse, error) {

	api, err := cli.conbineURI(apigetSharedCode, req)
	if err != nil {
		return nil, err
	}

	res := new(getSharedCodeResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
