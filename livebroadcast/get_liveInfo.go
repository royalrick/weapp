package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetLiveInfo = "/wxa/business/getliveinfo"

type PriceType uint8

const (
	_                 PriceType = iota
	PriceTypeNormal             // 一口价
	PriceTypeRange              // 区间价格
	PriceTypeDiscount           // 折扣价格
)

type LiveRoomGoods struct {
	// 商品封面图链接
	CoverIMG string `json:"cover_img"`
	// 商品小程序路径
	URL string `json:"url"`
	// 商品名称
	Name string `json:"name"`
	// 商品价格（分）
	Price int64 `json:"price"`
	// 商品价格（参考价格类型）
	Price2 int64 `json:"price2"`
	// 价格类型，1：一口价（只需要传入price，price2不传） 2：价格区间（price字段为左边界，price2字段为右边界，price和price2必填） 3：显示折扣价（price字段为原价，price2字段为现价， price和price2必填）
	Price_type PriceType `json:"price_type"`
	// 商品id
	GoodsID int `json:"goods_id"`
	//第三方商品appid ,当前小程序商品则为空
	ThirdPartyAppID string `json:"third_party_appid"`
}

type LiveType uint8

const (
	// 手机直播
	LiveTypePhone LiveType = iota
	// 推流
	LiveTypePushFlow
)

// 直播间状态
type LiveStatus uint8

const (
	_                   LiveStatus = iota + 100
	LiveStatusLiving               // 直播中
	LiveStatusNotStated            // 未开始
	LiveStatusFinished             // 已结束
	LiveStatusBan                  // 禁播
	LiveStatusPause                // 暂停
	LiveStatusException            // 异常
	LiveStatusExpired              // 已过期
)

type GetLiveInfoRequest struct {
	// 必填 起始房间，0表示从第1个房间开始拉取
	Start int `json:"start"`
	// 必填 每次拉取的房间数量，建议100以内
	Limit int `json:"limit"`
}

type GetLiveInfoResponse struct {
	request.CommonError
	// 拉取房间总数
	Total int         `json:"total"`
	List  []*RoomInfo `json:"room_info"`
}

type RoomInfo struct {
	// 直播间名称
	Name string `json:"name"`
	// 直播间ID
	RoomID int `json:"roomid"`
	// 直播间背景图链接
	CoverIMG string `json:"cover_img"`
	// 直播间分享图链接
	ShareIMG string `json:"share_img"`
	// 直播间状态
	LiveStatus LiveStatus `json:"live_status"`
	// 直播间开始时间
	StartTime int `json:"start_time"`
	// 直播计划结束时间
	EndTime int `json:"end_time"`
	// 主播名
	AnchorName string           `json:"anchor_name"`
	Goods      []*LiveRoomGoods `json:"goods"`
	// 直播类型
	LiveType LiveType `json:"live_type"`
	// 是否关闭点赞
	CloseLike int `json:"close_like"`
	// 是否关闭货架
	CloseGoods int `json:"close_goods"`
	// 是否关闭评论
	CloseComment int `json:"close_comment"`
	// 是否关闭客服
	CloseKF int `json:"close_kf"`
	// 是否关闭回放
	CloseReplay int `json:"close_replay"`
	// 是否开启官方收录
	IsFeedsPublic int `json:"is_feeds_public"`
	// 创建者openid
	CreatorOpenID string `json:"creater_openid"`
	// 官方收录封面
	FeedsIMG string `json:"feeds_img"`
}

// 获取直播间列表及直播间信息
func (cli *LiveBroadcast) GetLiveInfo(req *GetLiveInfoRequest) (*GetLiveInfoResponse, error) {

	api, err := cli.combineURI(apiGetLiveInfo, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetLiveInfoResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
