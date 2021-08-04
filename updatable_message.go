package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiCreateActivityID = "/cgi-bin/message/wxopen/activityid/create"
	apiSetUpdatableMsg  = "/cgi-bin/message/wxopen/updatablemsg/send"
)

// CreateActivityIDResponse 动态消息
type CreateActivityIDResponse struct {
	request.CommonError
	ActivityID     string `json:"activity_id"`     //	动态消息的 ID
	ExpirationTime uint   `json:"expiration_time"` //	activity_id 的过期时间戳。默认24小时后过期。
}

// CreateActivityID 创建被分享动态消息的 activity_id。
// token 接口调用凭证
func (cli *Client) CreateActivityID() (*CreateActivityIDResponse, error) {
	api := baseURL + apiCreateActivityID

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.createActivityID(api, token)
}

func (cli *Client) createActivityID(api, token string) (*CreateActivityIDResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CreateActivityIDResponse)
	if err := cli.request.Get(api, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdatableMsgTempInfo 动态消息对应的模板信息
type UpdatableMsgTempInfo struct {
	ParameterList []UpdatableMsgParameter `json:"parameter_list"` // 模板中需要修改的参数列表
}

// UpdatableMsgParameter 参数
// member_count	target_state = 0 时必填，文字内容模板中 member_count 的值
// room_limit	target_state = 0 时必填，文字内容模板中 room_limit 的值
// path	target_state = 1 时必填，点击「进入」启动小程序时使用的路径。
// 对于小游戏，没有页面的概念，可以用于传递查询字符串（query），如 "?foo=bar"
// version_type	target_state = 1 时必填，点击「进入」启动小程序时使用的版本。
// 有效参数值为：develop（开发版），trial（体验版），release（正式版）
type UpdatableMsgParameter struct {
	Name  UpdatableMsgParamName `json:"name"`  // 要修改的参数名
	Value string                `json:"value"` // 修改后的参数值
}

// UpdatableMsgTargetState 动态消息修改后的状态
type UpdatableMsgTargetState = uint8

// 动态消息状态
const (
	UpdatableMsgJoining UpdatableMsgTargetState = iota // 未开始
	UpdatableMsgStarted                                // 已开始
)

// UpdatableMsgParamName 参数 name 的合法值
type UpdatableMsgParamName = string

// 动态消息状态
const (
	UpdatableMsgParamMemberCount UpdatableMsgParamName = "member_count" // target_state = 0 时必填，文字内容模板中 member_count 的值
	UpdatableMsgParamRoomLimit                         = "room_limit"   // target_state = 0 时必填，文字内容模板中 room_limit 的值
	UpdatableMsgParamPath                              = "path"         // target_state = 1 时必填，点击「进入」启动小程序时使用的路径。 对于小游戏，没有页面的概念，可以用于传递查询字符串（query），如 "?foo=bar"
	UpdatableMsgParamVersionType                       = "version_type" // target_state = 1 时必填，点击「进入」启动小程序时使用的版本。有效参数值为：develop（开发版），trial（体验版），release（正式版）
)

// UpdatableMsg 动态消息
type UpdatableMsg struct {
	ActivityID   string                  `json:"activity_id"`   // 动态消息的 ID，通过 updatableMessage.createActivityId 接口获取
	TargetState  UpdatableMsgTargetState `json:"target_state"`  // 动态消息修改后的状态（具体含义见后文）
	TemplateInfo UpdatableMsgTempInfo    `json:"template_info"` // 动态消息对应的模板信息
}

// 修改被分享的动态消息。
// accessToken 接口调用凭证
func (cli *Client) SetUpdatableMsg(msg *UpdatableMsg) (*request.CommonError, error) {
	api := baseURL + apiSetUpdatableMsg

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.setUpdatableMsg(api, token, msg)
}

func (cli *Client) setUpdatableMsg(api, token string, setter *UpdatableMsg) (*request.CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(api, setter, res); err != nil {
		return nil, err
	}

	return res, nil
}
