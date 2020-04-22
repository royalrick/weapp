package weapp

const (
	apiSubscribeMessage = "/cgi-bin/message/subscribe/send"
)

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	ToUser           string               `json:"touser"`
	TemplateID       string               `json:"template_id"`
	Page             string               `json:"page,omitempty"`
	MiniprogramState MiniprogramState     `json:"miniprogram_state,omitempty"`
	Data             SubscribeMessageData `json:"data"`
}

// MiniprogramState 跳转小程序类型
type MiniprogramState = string

// developer为开发版；trial为体验版；formal为正式版；默认为正式版
const (
	MiniprogramStateDeveloper = "developer"
	MiniprogramStateTrial     = "trial"
	MiniprogramStateFormal    = "formal"
)

// SubscribeMessageData 订阅消息模板数据
type SubscribeMessageData map[string]struct {
	Value string `json:"value"`
}

// Send 发送订阅消息
//
// token access_token
func (sm *SubscribeMessage) Send(token string) (*CommonError, error) {
	api := baseURL + apiSubscribeMessage
	return sm.send(api, token)
}

func (sm *SubscribeMessage) send(api, token string) (*CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := &CommonError{}
	if err := postJSON(api, sm, res); err != nil {
		return nil, err
	}

	return res, nil
}
