package weapp

const (
	apiCreateActivityID = "/cgi-bin/message/wxopen/activityid/create"
	apiSetUpdatableMsg  = "/cgi-bin/message/wxopen/updatablemsg/send"
)

// ActivityID 动态消息
type ActivityID struct {
	CommonError
	ID             string `json:"activity_id"`     //	动态消息的 ID
	ExpirationTime uint   `json:"expiration_time"` //	activity_id 的过期时间戳。默认24小时后过期。
}

// CreateActivityID 创建被分享动态消息的 activity_id。
// @accessToken 接口调用凭证
func CreateActivityID(accessToken string) (*ActivityID, error) {
	api, err := tokenAPI(baseURL+apiCreateActivityID, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(ActivityID)
	if err := postJSON(api, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ActMsg 动态消息
type ActMsg struct {
	ID       string      `json:"activity_id"`   // 动态消息的 ID，通过 updatableMessage.createActivityId 接口获取
	State    TargetState `json:"target_state"`  // 动态消息修改后的状态（具体含义见后文）
	Template Template    `json:"template_info"` // 动态消息对应的模板信息
}

// ActMsgTemplate 动态消息对应的模板信息
type ActMsgTemplate struct {
	Params []ActMsgParam `json:"parameter_list"` // 模板中需要修改的参数
}

// ActMsgParam 模板中需要修改的参数
type ActMsgParam struct {
	// name 的合法值
	// 	member_count	target_state = 0 时必填，文字内容模板中 member_count 的值
	// room_limit	target_state = 0 时必填，文字内容模板中 room_limit 的值
	// path	target_state = 1 时必填，点击「进入」启动小程序时使用的路径。
	// 对于小游戏，没有页面的概念，可以用于传递查询字符串（query），如 "?foo=bar"
	// version_type	target_state = 1 时必填，点击「进入」启动小程序时使用的版本。
	// 有效参数值为：develop（开发版），trial（体验版），release（正式版）
	Name  string `json:"name"`  // 要修改的参数名
	Value string `json:"value"` // 修改后的参数值
}

// TargetState 动态消息修改后的状态
type TargetState = int8

// 动态消息状态
const (
	Unstarted TargetState = 0 // 未开始
	Started               = 1 // 已开始
)

// SetUpdatableMsg 修改被分享的动态消息。
// @accessToken 接口调用凭证
func (msg *ActMsg) SetUpdatableMsg(accessToken string) (*CommonError, error) {
	api, err := tokenAPI(baseURL+apiSetUpdatableMsg, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postJSON(api, msg, res); err != nil {
		return nil, err
	}

	return res, nil
}
