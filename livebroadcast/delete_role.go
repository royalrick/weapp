package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiDeleteRole = "/wxaapi/broadcast/role/deleterole"

type DeleteRoleRequest struct {
	// 必填	设置用户的角色
	// 取值[1-管理员，2-主播，3-运营者]，设置超级管理员将无效
	Role Role `json:"role"`
	// 必填	用户昵称
	Nickname string `json:"nickname"`
}

// 解除成员角色
func (cli *LiveBroadcast) DeleteRole(req *DeleteRoleRequest) (*request.CommonError, error) {

	api, err := cli.combineURI(apiDeleteRole, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
