package updatablemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiSetUpdatableMsg = "/cgi-bin/message/wxopen/updatablemsg/send"

// UpdatableMsgTargetState 动态消息修改后的状态
type UpdatableMsgTargetState = uint8

// 动态消息状态
const (
	UpdatableMsgJoining UpdatableMsgTargetState = iota // 未开始
	UpdatableMsgStarted                                // 已开始
)

type SetUpdatableMsgRequest struct {
	ActivityID   string                  `json:"activity_id"`   // 动态消息的 ID，通过 updatableMessage.createActivityId 接口获取
	TargetState  UpdatableMsgTargetState `json:"target_state"`  // 动态消息修改后的状态（具体含义见后文）
	TemplateInfo UpdatableMsgTempInfo    `json:"template_info"` // 动态消息对应的模板信息
}

// UpdatableMsgParamName 参数 name 的合法值
type UpdatableMsgParamName = string

// 动态消息状态
const (
	UpdatableMsgParamMemberCount UpdatableMsgParamName = "member_count" // target_state = 0 时必填，文字内容模板中 member_count 的值
	UpdatableMsgParamRoomLimit   UpdatableMsgParamName = "room_limit"   // target_state = 0 时必填，文字内容模板中 room_limit 的值
	UpdatableMsgParamPath        UpdatableMsgParamName = "path"         // target_state = 1 时必填，点击「进入」启动小程序时使用的路径。 对于小游戏，没有页面的概念，可以用于传递查询字符串（query），如 "?foo=bar"
	UpdatableMsgParamVersionType UpdatableMsgParamName = "version_type" // target_state = 1 时必填，点击「进入」启动小程序时使用的版本。有效参数值为：develop（开发版），trial（体验版），release（正式版）
)

// UpdatableMsgTempInfo 动态消息对应的模板信息
type UpdatableMsgTempInfo struct {
	ParameterList []UpdatableMsgParameter `json:"parameter_list"` // 模板中需要修改的参数列表
}
type UpdatableMsgParameter struct {
	Name  UpdatableMsgParamName `json:"name"`  // 要修改的参数名
	Value string                `json:"value"` // 修改后的参数值
}

// 修改被分享的动态消息。
func (cli *UpdatableMessage) SetUpdatableMsg(req *SetUpdatableMsgRequest) (*request.CommonError, error) {

	api, err := cli.combineURI(apiSetUpdatableMsg, nil, true)
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
