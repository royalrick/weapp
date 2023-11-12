package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiAddGoods = "/wxaapi/broadcast/room/addgoods"

type AddGoodsRequest struct {
	// 必填	数组列表，可传入多个，里面填写 商品 ID
	Ids []int64 `json:"ids"`
	// 必填	房间ID
	RoomId int64 `json:"roomId"`
}

// 调用接口往指定直播间导入已入库的商品
func (cli *LiveBroadcast) AddGoods(req *AddGoodsRequest) (*request.CommonError, error) {

	api, err := cli.combineURI(apiAddGoods, nil, true)
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
