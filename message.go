package weapp

const (
	apiSendMessage = "/cgi-bin/message/custom/send"
)

// 消息体
type message struct {
	Receiver string  `json:"touser"`  // user openID
	Type     MsgType `json:"msgtype"` // text | image | link | miniprogrampage
	Text     Text    `json:"text,omitempty"`
	Image    Image   `json:"image,omitempty"`
	Link     Link    `json:"link,omitempty"`
	Card     Card    `json:"miniprogrampage,omitempty"`
}

// SendTo 发送文本消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Text) SendTo(openID, token string) (*BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "text",
		Text:     msg,
	}

	return sendMessage(token, params)
}

// SendTo 发送图片消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Image) SendTo(openID, token string) (*BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "image",
		Image:    msg,
	}

	return sendMessage(token, params)
}

// SendTo 发送图文链接消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Link) SendTo(openID, token string) (*BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "link",
		Link:     msg,
	}

	return sendMessage(token, params)
}

// SendTo 发送卡片消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Card) SendTo(openID, token string) (*BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "miniprogrampage",
		Card:     msg,
	}

	return sendMessage(token, params)
}

// send 发送消息
//
// @token 微信 access_token
func sendMessage(token string, params interface{}) (*BaseResponse, error) {
	api, err := TokenAPI(BaseURL+apiSendMessage, token)
	if err != nil {
		return nil, err
	}

	res := new(BaseResponse)
	if err := PostJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
