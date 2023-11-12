package subscribemessage

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiSend = "/cgi-bin/message/subscribe/send"

// SubscribeMessage 订阅消息
type SendRequest struct {
	ToUser           string           `json:"touser"`
	TemplateID       string           `json:"template_id"`
	Page             string           `json:"page,omitempty"`
	MiniprogramState MiniprogramState `json:"miniprogram_state,omitempty"`
	Data             SendData         `json:"data"`
}

// 模板数据内容
type SendData map[string]SendValue

// 模板变量值
type SendValue struct {
	Value string `json:"value"`
}

// MiniprogramState 跳转小程序类型
type MiniprogramState = string

// developer为开发版；trial为体验版；formal为正式版；默认为正式版
const (
	MiniprogramStateDeveloper MiniprogramState = "developer"
	MiniprogramStateTrial     MiniprogramState = "trial"
	MiniprogramStateFormal    MiniprogramState = "formal"
)

// 发送订阅消息
func (cli *SubscribeMessage) Send(msg *SendRequest) (*request.CommonError, error) {
	api, err := cli.combineURI(apiSend, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(api, msg, res); err != nil {
		return nil, err
	}

	return res, nil
}
