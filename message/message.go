// Package message 消息
package message

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
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
func (msg Text) SendTo(openID, token string) (wres weapp.Response, err error) {

	m := message{
		Receiver: openID,
		Type:     "text",
		Text:     msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// SendTo 发送图片消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Image) SendTo(openID, token string) (wres weapp.Response, err error) {

	m := message{
		Receiver: openID,
		Type:     "image",
		Image:    msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// SendTo 发送图文链接消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Link) SendTo(openID, token string) (wres weapp.Response, err error) {

	m := message{
		Receiver: openID,
		Type:     "link",
		Link:     msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// SendTo 发送卡片消息
//
// @openID 用户openID
// @token 微信 access_token
func (msg Card) SendTo(openID, token string) (wres weapp.Response, err error) {

	m := message{
		Receiver: openID,
		Type:     "miniprogrampage",
		Card:     msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// send 发送消息
//
// @token 微信 access_token
func send(token, body string) (wres weapp.Response, err error) {
	api, err := util.TokenAPI(weapp.BaseURL+sendAPI, token)
	if err != nil {
		return
	}

	res, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(weapp.WeChatServerError)
		return
	}

	if err = json.NewDecoder(res.Body).Decode(&wres); err != nil {
		return
	}

	err = errors.New(wres.Errmsg)
	return
}
