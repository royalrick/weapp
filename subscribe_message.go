package weapp

const (
	apiSubscribeMessage = "/cgi-bin/message/subscribe/send"
)

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	ToUser     string               `json:"touser"`
	TemplateID string               `json:"template_id"`
	Page       string               `json:"page,omitempty"`
	Data       SubscribeMessageData `json:"data"`
}

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
