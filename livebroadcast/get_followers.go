package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetFollowers = "/wxa/business/get_wxa_followers"

type GetFollowersRequest struct {
	Limit     int `json:"limit"`      // 非必填	获取长期订阅用户的个数限制，默认200，最大2000
	PageBreak int `json:"page_break"` // 非必填	翻页标记，获取第一页时不带，第二页开始需带上上一页返回结果中的page_break
}

type GetFollowersResponse struct {
	request.CommonError
	Followers []struct {
		SubscribeTime int    `json:"subscribe_time"` // 长期订阅用户订阅时间
		RoomId        int    `json:"room_id"`        // 用户订阅时所处房间
		RoomStatus    int    `json:"room_status"`    // 用户订阅时房间状态
		Openid        string `json:"openid"`         // 长期订阅用户OpenId
	} `json:"followers"` //	长期订阅用户列表
	PageBreak int `json:"page_break"` // 翻页标记，获取下一页时带上该值
}

// 获取长期订阅用户
func (cli *LiveBroadcast) GetFollowers(req *GetFollowersRequest) (*GetFollowersResponse, error) {

	api, err := cli.combineURI(apiGetFollowers, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetFollowersResponse)
	err = cli.request.Post(api, res, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
