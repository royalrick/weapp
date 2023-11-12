package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsList = "/wxaapi/broadcast/goods/getapproved"

type GoodsListRequest struct {
	// 必填	分页条数起点
	Offset int `query:"offset"`
	// 非必填 分页大小，默认30，不超过100
	Limit int `query:"limit"`
	// 必填	商品状态，0：未审核。1：审核中，2：审核通过，3：审核驳回
	Status uint8 `query:"status"`
}

type GoodsListResponse struct {
	request.CommonError
	// 商品个数
	Total int64 `json:"total"`
	// 商品
	Goods []struct {
		// 商品ID
		GoodsId int64 `json:"goodsId"`
		// 商品图片url
		CoverImgUrl string `json:"coverImgUrl"`
		// 商品名称
		Name string `json:"name"`
		// 价格左区间，单位“元”
		Price float64 `json:"price"`
		// 价格右区间，单位“元”
		Price2 float64 `json:"price2"`
		// 商品详情页的小程序路径
		Url string `json:"url"`
		// 1:一口价，此时读price字段; 2:价格区间，此时price字段为左边界，price2字段为右边界; 3:折扣价，此时price字段为原价，price2字段为现价；
		PriceType uint8 `json:"priceType"`
		// 1、2：表示是为 API 添加商品，否则是直播控制台添加的商品
		ThirdPartyTag uint8 `json:"thirdPartyTag"`
		// 当商品为第三方小程序的商品则为对应第三方小程序的appid，自身小程序商品	则为''
		ThirdPartyAppid string `json:"thirdPartyAppid"`
	} `json:"goods"`
}

// 获取商品列表
func (cli *LiveBroadcast) GoodsList(req *GoodsListRequest) (*GoodsListResponse, error) {

	api, err := cli.combineURI(apiGoodsList, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsListResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
