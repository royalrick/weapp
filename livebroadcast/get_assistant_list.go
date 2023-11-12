package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetAssistantList = "/wxaapi/broadcast/room/getassistantlist"

type GetAssistantListRequest struct {
	// 必填	直播间id
	RoomId int64 `query:"roomId"`
}

type GetAssistantListResponse struct {
	request.CommonError
	// 小助手列表
	List []struct {
		// 修改时间// 修改时间
		Timestamp int `json:"timestamp"`
		// 头像// 头像
		Headimg string `json:"headimg"`
		// 昵称// 昵称
		Nickname string `json:"nickname"`
		// 微信号// 微信号
		Alias string `json:"alias"`
		// openid// openid
		Openid string `json:"openid"`
	} `json:"list"`
	// 小助手个数// 小助手个数
	Count int `json:"count"`
	// 小助手最大个数// 小助手最大个数
	MaxCount int `json:"maxCount"`
}

// 查询管理直播间小助手
func (cli *LiveBroadcast) GetAssistantList(req *GetAssistantListRequest) (*GetAssistantListResponse, error) {

	api, err := cli.combineURI(apiGetAssistantList, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetAssistantListResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
