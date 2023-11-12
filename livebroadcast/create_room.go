package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiCreateRoom = "/wxaapi/broadcast/room/create"

type CreateRoomRequest struct {
	// 必填 直播间名字，最短3个汉字，最长17个汉字，1个汉字相当于2个字符
	Name string `json:"name"`
	// 必填 背景图，填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html；直播间背景图，图片规则：建议像素1080*1920，大小不超过2M
	CoverImg string `json:"coverImg"`
	// 必填 直播计划开始时间（开播时间需要在当前时间的10分钟后 并且 开始时间不能在 6 个月后）
	StartTime int64 `json:"startTime"`
	// 必填 直播计划结束时间（开播时间和结束时间间隔不得短于30分钟，不得超过24小时）
	EndTime int64 `json:"endTime"`
	// 必填 主播昵称，最短2个汉字，最长15个汉字，1个汉字相当于2个字符
	AnchorName string `json:"anchorName"`
	// 必填 主播微信号，如果未实名认证，需要先前往“小程序直播”小程序进行实名验证, 小程序二维码链接：https://res.wx.qq.com/op_res/9rSix1dhHfK4rR049JL0PHJ7TpOvkuZ3mE0z7Ou_Etvjf-w1J_jVX0rZqeStLfwh
	AnchorWechat string `json:"anchorWechat"`
	// 非必填 主播副号微信号，如果未实名认证，需要先前往“小程序直播”小程序进行实名验证, 小程序二维码链接：https://res.wx.qq.com/op_res/9rSix1dhHfK4rR049JL0PHJ7TpOvkuZ3mE0z7Ou_Etvjf-w1J_jVX0rZqeStLfwh
	SubAnchorWechat string `json:"subAnchorWechat"`
	// 非必填 创建者微信号，不传入则此直播间所有成员可见。传入则此房间仅创建者、管理员、超管、直播间主播可见
	CreaterWechat string `json:"createrWechat"`
	// 必填 分享图，填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html；直播间分享图，图片规则：建议像素800*640，大小不超过1M；
	ShareImg string `json:"shareImg"`
	// 必填 购物直播频道封面图，填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，请参考以下文档： https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html; 购物直播频道封面图，图片规则：建议像素800*800，大小不超过100KB；
	FeedsImg string `json:"feedsImg"`
	// 非必填 是否开启官方收录 【1: 开启，0：关闭】，默认开启收录
	IsFeedsPublic uint8 `json:"isFeedsPublic"`
	// 必填 直播间类型 【1: 推流，0：手机直播】
	Type uint8 `json:"type"`
	// 必填 是否关闭点赞 【0：开启，1：关闭】（若关闭，观众端将隐藏点赞按钮，直播开始后不允许开启）
	CloseLike uint8 `json:"closeLike"`
	// 必填 是否关闭货架 【0：开启，1：关闭】（若关闭，观众端将隐藏商品货架，直播开始后不允许开启）
	CloseGoods uint8 `json:"closeGoods"`
	// 必填 是否关闭评论 【0：开启，1：关闭】（若关闭，观众端将隐藏评论入口，直播开始后不允许开启）
	CloseComment uint8 `json:"closeComment"`
	// 非必填 是否关闭回放 【0：开启，1：关闭】默认关闭回放（直播开始后允许开启）
	CloseReplay uint8 `json:"closeReplay"`
	// 非必填 是否关闭分享 【0：开启，1：关闭】默认开启分享（直播开始后不允许修改）
	CloseShare uint8 `json:"closeShare"`
	// 非必填 是否关闭客服 【0：开启，1：关闭】 默认关闭客服（直播开始后允许开启）
	CloseKf uint8 `json:"closeKf"`
}
type CreateRoomResponse struct {
	request.CommonError
	// 当主播微信号没有在 “小程序直播“ 小程序实名认证 返回该字段
	// "小程序直播" 小程序码
	QRCodeURL string `json:"qrcode_url"`
	//	房间ID
	RoomId int64 `json:"roomId"`
}

// 创建直播间
func (cli *LiveBroadcast) CreateRoom(req *CreateRoomRequest) (*CreateRoomResponse, error) {

	api, err := cli.combineURI(apiCreateRoom, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(CreateRoomResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
