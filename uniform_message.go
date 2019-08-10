package weapp

// TemplateMsg 小程序模板消息
type TemplateMsg struct {
	TemplateID      string `json:"template_id"`
	Page            string `json:"page"`
	FormID          string `json:"form_id"`
	Data            Data   `json:"data"`
	EmphasisKeyword string `json:"emphasis_keyword,omitempty"`
}

// Data 模板消息内容
type Data = map[string]Keyword

// Keyword 关键字
type Keyword struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// MPTemplateMsg 公众号模板消息
type MPTemplateMsg struct {
	AppID       string      `json:"appid"`
	TemplateID  string      `json:"template_id"`
	URL         string      `json:"url"`
	Miniprogram Miniprogram `json:"miniprogram"`
	Data        Data        `json:"data"`
}

// Miniprogram 小程序
type Miniprogram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

// UniformMsg 统一服务消息
type UniformMsg struct {
	ToUser           string        `json:"touser"` // 用户 openid
	MPTemplateMsg    MPTemplateMsg `json:"mp_template_msg"`
	WeappTemplateMsg TemplateMsg   `json:"weapp_template_msg"`
}

// Send 统一服务消息
//
// token access_token
func (msg UniformMsg) Send(token string) error {
	api, err := tokenAPI(baseURL+apiUniformSendTemplateMessage, token)
	if err != nil {
		return err
	}

	res := new(CommonError)
	if err := postJSON(api, msg, res); err != nil {
		return err
	}

	return nil
}
