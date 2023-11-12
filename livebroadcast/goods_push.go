package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsPush = "/wxaapi/broadcast/goods/push"

type GoodsPushRequest struct {
	// 必填 商品ID
	GoodsId int64 `json:"goodsId"`
	// 必填 房间ID
	RoomId int64 `json:"roomId"`
}

type GoodsPushResponse struct {
	request.CommonError
}

// 推送商品
func (cli *LiveBroadcast) GoodsPush(req *GoodsPushRequest) (*GoodsPushResponse, error) {

	api, err := cli.combineURI(apiGoodsPush, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsPushResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
