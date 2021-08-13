package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetAssistantList = "/wxaapi/broadcast/room/getassistantlist"

type GetAssistantListRequest struct {
	RoomId int64 `query:"roomId"` // 必填	直播间id
}

type GetAssistantListResponse struct {
	request.CommonError
	List []struct {
		Timestamp int    `json:"timestamp"` // 修改时间
		Headimg   string `json:"headimg"`   // 头像
		Nickname  string `json:"nickname"`  // 昵称
		Alias     string `json:"alias"`     // 微信号
		Openid    string `json:"openid"`    // openid
	} `json:"list"` //	小助手列表
	Count    int `json:"count"`    // 小助手个数
	MaxCount int `json:"maxCount"` // 小助手最大个数

}

// 查询管理直播间小助手
func (cli *LiveBroadcast) GetAssistantList(req *GetAssistantListRequest) (*GetAssistantListResponse, error) {

	api, err := cli.conbineURI(apiGetAssistantList, req)
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
