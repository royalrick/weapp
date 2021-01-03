package weapp

const (
	// 创建直播间
	apiCreateLiveRoom = "/wxaapi/broadcast/room/create"
)

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
	CommonError
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
func (creator *LiveRoomCreator) CreateLiveRoom(accessToken string) (*CreateLiveRoomResponse, error) {
	api := baseURL + apiCreateLiveRoom

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

	return createLiveRoom(api, accessToken, &req)
}

func createLiveRoom(api, token string, info *createLiveRoomRequest) (*CreateLiveRoomResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	rsp := new(CreateLiveRoomResponse)
	if err := postJSON(api, nil, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
