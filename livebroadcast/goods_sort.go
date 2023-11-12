package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsSort = "/wxaapi/broadcast/goods/sort"

type GoodsSortRequest struct {
	// 必填 房间ID
	RoomId int64 `json:"roomId"`

	// 必填 房间ID
	Goods []*SortGoods `json:"goods"`
}

type SortGoods struct {
	// 必填 商品ID
	GoodsId int64 `json:"goodsId"`
}

type GoodsSortResponse struct {
	request.CommonError
}

// 直播间商品排序
func (cli *LiveBroadcast) GoodsSort(req *GoodsSortRequest) (*GoodsSortResponse, error) {

	api, err := cli.combineURI(apiGoodsSort, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsSortResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
