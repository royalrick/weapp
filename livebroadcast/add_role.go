package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiAddRole = "/wxaapi/broadcast/role/addrole"

type Role = int8

const (
	RoleALl Role = iota - 1
	// 超级管理员
	RoleRoot
	// 管理员
	RoleAdministrator
	// 主播
	RoleBroadcaster
	// 运营者
	RoleOperator
)

type AddRoleRequest struct {
	// 必填	设置用户的角色
	// 取值[1-管理员，2-主播，3-运营者]，设置超级管理员将无效
	Role Role `json:"role"`
	// 必填	用户昵称
	Nickname string `json:"nickname"`
}

// 调用此接口设置小程序直播成员的管理员、运营者和主播角色
func (cli *LiveBroadcast) AddRole(req *AddRoleRequest) (*request.CommonError, error) {

	api, err := cli.combineURI(apiAddRole, nil, true)
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
