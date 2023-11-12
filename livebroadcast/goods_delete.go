package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsDelete = "/wxaapi/broadcast/goods/delete"

type GoodsDeleteRequest struct {
	// 商品ID
	GoodsId int64 `json:"goodsId"`
}

type GoodsDeleteResponse struct {
	request.CommonError
}

// 删除商品
func (cli *LiveBroadcast) GoodsDelete(req *GoodsDeleteRequest) (*GoodsDeleteResponse, error) {

	api, err := cli.combineURI(apiGoodsDelete, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsDeleteResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
