package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGetRoleList = "/wxaapi/broadcast/role/getrolelist"

type GetRoleListRequest struct {
	// 非必填	查询的用户角色，取值 [-1-所有成员， 0-超级管理员，1-管理员，2-主播，3-运营者]，默认-1
	Role Role `query:"role"`
	// 非必填	起始偏移量, 默认0
	Offset int `query:"offset"`
	// 非必填	查询个数，最大30，默认10
	Limit int `query:"limit"`
	// 非必填	搜索的微信号或昵称，不传则返回全部
	Keyword string `query:"keyword"`
}

type GetRoleListResponse struct {
	request.CommonError
	// 角色列表
	List []struct {
		// 微信用户头像url
		Headingimg string `json:"headingimg"`
		// 微信用户昵称
		Nickname string `json:"nickname"`
		// 微信openid
		Openid string `json:"openid"`
		// 具有的身份，[0-超级管理员，1-管理员，2-主播，3-运营者]
		RoleList []Role `json:"roleList"`
		// 更新时间戳
		UpdateTimestamp string `json:"updateTimestamp"`
		// 脱敏微信号
		Username string `json:"username"`
	} `json:"list"`
}

// 查询小程序直播成员列表
func (cli *LiveBroadcast) GetRoleList(req *GetRoleListRequest) (*GetRoleListResponse, error) {

	api, err := cli.combineURI(apiGetRoleList, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetRoleListResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
