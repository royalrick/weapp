package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsSale = "/wxaapi/broadcast/goods/onsale"

type GoodsSaleRequest struct {
	// 必填 商品ID
	GoodsId int64 `json:"goodsId"`
	// 必填 审核单ID
	AuditId int64 `json:"auditId"`
	// 必填 上下架 【0：下架，1：上架】
	OnSale uint8 `json:"onSale"`
}

type GoodsSaleResponse struct {
	request.CommonError
}

// 上下架商品
func (cli *LiveBroadcast) GoodsSale(req *GoodsSaleRequest) (*GoodsSaleResponse, error) {

	api, err := cli.combineURI(apiGoodsSale, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsSaleResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
