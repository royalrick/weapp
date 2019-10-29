package weapp

const (
	apiSendUniformMessage = "/cgi-bin/message/wxopen/template/uniform_send"
)

// UniformMsgData 模板消息内容
type UniformMsgData map[string]UniformMsgKeyword

// UniformMsgKeyword 关键字
type UniformMsgKeyword struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// UniformWeappTmpMsg 小程序模板消息
type UniformWeappTmpMsg struct {
	TemplateID      string         `json:"template_id"`
	Page            string         `json:"page"`
	FormID          string         `json:"form_id"`
	Data            UniformMsgData `json:"data"`
	EmphasisKeyword string         `json:"emphasis_keyword,omitempty"`
}

// UniformMsgMiniprogram 小程序
type UniformMsgMiniprogram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

// UniformMpTmpMsg 公众号模板消息
type UniformMpTmpMsg struct {
	AppID       string                `json:"appid"`
	TemplateID  string                `json:"template_id"`
	URL         string                `json:"url"`
	Miniprogram UniformMsgMiniprogram `json:"miniprogram"`
	Data        UniformMsgData        `json:"data"`
}

// Miniprogram 小程序
type Miniprogram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

// UniformMsgSender 统一服务消息
type UniformMsgSender struct {
	ToUser             string             `json:"touser"` // 用户 openid
	UniformWeappTmpMsg UniformWeappTmpMsg `json:"weapp_template_msg,omitempty"`
	UniformMpTmpMsg    UniformMpTmpMsg    `json:"mp_template_msg,omitempty"`
}

// Send 统一服务消息
//
// token access_token
func (msg *UniformMsgSender) Send(token string) (*CommonError, error) {
	api := baseURL + apiSendUniformMessage
	return msg.send(api, token)
}

func (msg *UniformMsgSender) send(api, token string) (*CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postJSON(api, msg, res); err != nil {
		return nil, err
	}

	return res, nil
}
