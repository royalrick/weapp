package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	// 创建直播间
	apiCreateLiveRoom    = "/wxaapi/broadcast/room/create"
	apiFetchLiveRoomList = "/wxa/business/getliveinfo"
)

// 分页器
type PageAble struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
}
type createLiveRoomRequest struct {
	Name            string `json:"name"`                      // 房间名字
	CoverImg        string `json:"coverImg"`                  // 通过 uploadfile 上传，填写 mediaID
	StartTime       int64  `json:"startTime"`                 // 开始时间
	EndTime         int64  `json:"endTime"`                   // 结束时间
	AnchorName      string `json:"anchorName"`                // 主播昵称
	AnchorWechat    string `json:"anchorWechat"`              // 主播微信号
	SubAnchorWechat string `json:"subAnchorWechat,omitempty"` // 主播副号微信号
	CreaterWechat   string `json:"createrWechat,omitempty"`   // 创建者微信号
	ShareImg        string `json:"shareImg"`                  //通过 uploadfile 上传，填写 mediaID
	FeedsImg        string `json:"feedsImg,omitempty"`        // 通过 uploadfile 上传，填写 mediaID
	IsFeedsPublic   uint8  `json:"isFeedsPublic,omitempty"`   // 是否开启官方收录，1 开启，0 关闭
	Type            uint8  `json:"type"`                      // 直播类型，1 推流 0 手机直播
	CloseLike       uint8  `json:"closeLike"`                 // 是否关闭点赞 1：关闭
	CloseGoods      uint8  `json:"closeGoods"`                // 是否关闭商品货架，1：关闭
	CloseComment    uint8  `json:"closeComment"`              //是否开启评论，1：关闭
	CloseReplay     uint8  `json:"closeReplay,omitempty"`     // 是否关闭回放 1 关闭
	CloseShare      uint8  `json:"closeShare,omitempty"`      // 是否关闭分享 1 关闭
	CloseKf         uint8  `json:"closeKf,omitempty"`         // 是否关闭客服，1 关闭
}

type CreateLiveRoomResponse struct {
	request.CommonError
	RoomID string `json:"roomId"` //房间ID
	// 当主播微信号没有在 “小程序直播“ 小程序实名认证 返回该字段
	QRCodeURL string `json:"qrcode_url"`
}

type LiveType uint8

const (
	// 手机直播
	LiveTypePhone LiveType = iota
	// 推流
	LiveTypePushFlow
)

// 直播间创建器
type LiveRoomCreator struct {
	Name            string   // 房间名字
	CoverImg        string   // 通过 uploadfile 上传，填写 mediaID
	StartTime       int64    // 开始时间
	EndTime         int64    // 结束时间
	AnchorName      string   // 主播昵称
	AnchorWechat    string   // 主播微信号
	SubAnchorWechat string   // 主播副号微信号
	CreatorWechat   string   // 创建者微信号
	ShareIMGMediaID string   //通过 uploadfile 上传，填写 mediaID
	FeedsIMGMediaID string   // 通过 uploadfile 上传，填写 mediaID
	IsFeedsPublic   bool     // 是否开启官方收录，1 开启，0 关闭
	Type            LiveType // 直播类型，1 推流 0 手机直播
	CloseLike       bool     // 是否关闭点赞
	CloseGoods      bool     // 是否关闭商品货架
	CloseComment    bool     //是否开启评论
	CloseReplay     bool     // 是否关闭回放
	CloseShare      bool     // 是否关闭分享
	CloseKf         bool     // 是否关闭客服
}

// 创建直播间
func (cli *Client) CreateLiveRoom(creator *LiveRoomCreator) (*CreateLiveRoomResponse, error) {
	api := baseURL + apiCreateLiveRoom

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	req := createLiveRoomRequest{
		Name:            creator.Name,
		CoverImg:        creator.CoverImg,
		StartTime:       creator.StartTime,
		EndTime:         creator.EndTime,
		AnchorName:      creator.AnchorName,
		AnchorWechat:    creator.AnchorWechat,
		SubAnchorWechat: creator.SubAnchorWechat,
		CreaterWechat:   creator.CreatorWechat,
		ShareImg:        creator.ShareIMGMediaID,
		FeedsImg:        creator.FeedsIMGMediaID,
		IsFeedsPublic:   bool2int(creator.IsFeedsPublic),
		Type:            uint8(creator.Type),
		CloseLike:       bool2int(creator.CloseLike),
		CloseGoods:      bool2int(creator.CloseGoods),
		CloseComment:    bool2int(creator.CloseComment),
		CloseReplay:     bool2int(creator.CloseReplay),
		CloseShare:      bool2int(creator.CloseShare),
		CloseKf:         bool2int(creator.CloseKf),
	}

	return cli.createLiveRoom(api, token, &req)
}

func (cli *Client) createLiveRoom(api, token string, info *createLiveRoomRequest) (*CreateLiveRoomResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	rsp := new(CreateLiveRoomResponse)
	if err := cli.request.Post(api, nil, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

type PriceType uint8

const (
	_                 PriceType = iota
	PriceTypeOne                // 一口价
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

type LiveRoomItem struct {
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
type FetchLiveRoomListResponse struct {
	request.CommonError
	// 拉取房间总数
	Total int             `json:"total"`
	List  []*LiveRoomItem `json:"room_info"`
}

// 获取直播间列表及直播间信息
//
// @start: 起始拉取房间，start = 0 表示从第 1 个房间开始拉取
// @limit: 每次拉取的个数上限，不要设置过大，建议 100 以内
func (cli *Client) FetchLiveRoomList(start, limit int) (*FetchLiveRoomListResponse, error) {
	api := baseURL + apiFetchLiveRoomList

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.fetchLiveRoomList(api, token, start, limit)
}

func (cli *Client) fetchLiveRoomList(api, token string, start, limit int) (*FetchLiveRoomListResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := PageAble{
		Start: start,
		Limit: limit,
	}

	rsp := new(FetchLiveRoomListResponse)
	if err := cli.request.Post(api, params, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
