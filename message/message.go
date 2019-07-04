// Package message 消息
package message

import (
	"github.com/medivhzhan/weapp"
)

const (
	sendAPI = "/cgi-bin/message/custom/send"
)

// MsgType 消息类型
type MsgType string

const (
	// TextMsg 文本消息类型
	TextMsg MsgType = "text"

	// ImgMsg 图片消息类型
	ImgMsg MsgType = "image"

	// LinkMsg 图文链接消息类型
	LinkMsg MsgType = "link"

	// CardMsg 小程序卡片消息类型
	CardMsg MsgType = "miniprogrampage"

	// Event 事件类型
	Event MsgType = "event"
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

// Text 文本消息
// 支持添加可跳转小程序的文字链
type Text struct {
	Content string `json:"content"`
}

// Image 图片消息
type Image struct {
	MediaID string `json:"media_id"`
}

// Card 卡片消息
type Card struct {
	Title        string `json:"title"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// Link 图文链接消息
type Link struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}

// SendTo 发送文本消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Text) SendTo(openID, token string) (*weapp.BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "text",
		Text:     msg,
	}

	return send(token, params)
}

// SendTo 发送图片消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Image) SendTo(openID, token string) (*weapp.BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "image",
		Image:    msg,
	}

	return send(token, params)
}

// SendTo 发送图文链接消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Link) SendTo(openID, token string) (*weapp.BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "link",
		Link:     msg,
	}

	return send(token, params)
}

// SendTo 发送卡片消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Card) SendTo(openID, token string) (*weapp.BaseResponse, error) {

	params := message{
		Receiver: openID,
		Type:     "miniprogrampage",
		Card:     msg,
	}

	return send(token, params)
}

// send 发送消息
//
// @token 微信 access_token
func send(token string, params interface{}) (*weapp.BaseResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+sendAPI, token)
	if err != nil {
		return nil, err
	}

	res := new(weapp.BaseResponse)
	if err := weapp.PostJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
