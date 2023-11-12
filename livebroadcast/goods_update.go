package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsUpdate = "/wxaapi/broadcast/goods/update"

type GoodsUpdateRequest struct {
	//	商品信息
	GoodsInfo *GoodsUpdateInfo `json:"goodsInfo"`
}

type GoodsUpdateInfo struct {
	// 必填 商品ID
	GoodsId int64 `json:"goodsId"`
	// 必填 填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html；图片规则：图片尺寸最大300像素*300像素；
	CoverImgUrl string `json:"coverImgUrl"`
	// 必填 商品名称，最长14个汉字，1个汉字相当于2个字符
	Name string `json:"name"`
	// 必填 价格类型，1：一口价（只需要传入price，price2不传） 2：价格区间（price字段为左边界，price2字段为右边界，price和price2必填） 3：显示折扣价（price字段为原价，price2字段为现价， price和price2必填）
	PriceType PriceType `json:"priceType"`
	// 必填 数字，最多保留两位小数，单位元
	Price float64 `json:"price"`
	// 非必填	数字，最多保留两位小数，单位元
	Price2 float64 `json:"price2"`
	// 必填 商品详情页的小程序路径，路径参数存在 url 的，该参数的值需要进行 encode 处理再填入
	Url string `json:"url"`
	// 非必填	当商品为第三方小程序的商品则填写为对应第三方小程序的appid，自身小程序商品则为''
	ThirdPartyAppid string `json:"thirdPartyAppid"`
}

type GoodsUpdateResponse struct {
	request.CommonError
}

// 更新商品
func (cli *LiveBroadcast) GoodsUpdate(req *GoodsUpdateRequest) (*GoodsUpdateResponse, error) {

	api, err := cli.combineURI(apiGoodsUpdate, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsUpdateResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
