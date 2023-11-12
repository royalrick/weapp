package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsVideo = "/wxaapi/broadcast/goods/getVideo"

type GoodsVideoRequest struct {
	// 必填 商品ID
	GoodsId int64 `json:"goodsId"`
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
}

type GoodsVideoResponse struct {
	request.CommonError
	// 必填 讲解链接
	Url int64 `json:"url"`
}

// 更新商品
func (cli *LiveBroadcast) GoodsVideo(req *GoodsVideoRequest) (*GoodsVideoResponse, error) {

	api, err := cli.combineURI(apiGoodsVideo, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsVideoResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
