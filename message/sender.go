package message

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin/json"
	"github.com/medivhzhan/weapp"
)

const (
	sendAPI = "/cgi-bin/message/custom/send"
)

// Code 微信服务器返回状态码
type Code int

const (
	// SystemBusy 系统繁忙 稍候再试
	SystemBusy weapp.Code = -1

	// SendSuccess 发送消息成功
	SendSuccess weapp.Code = 0

	// ErrInvalidToken 获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的小程序调用接口
	ErrInvalidToken weapp.Code = 40001

	// ErrInvalidClaims 不合法的凭证类型
	ErrInvalidClaims weapp.Code = 40002

	// ErrInvalidOpenid 不合法的 OpenID，请开发者确认OpenID否是其他小程序的 OpenID
	ErrInvalidOpenid weapp.Code = 40003

	// ErrOutOfTime 回复时间超过限制
	ErrOutOfTime weapp.Code = 45015

	// ErrOutOfLimit 客服接口下行条数超过上限
	ErrOutOfLimit weapp.Code = 45047

	// ErrUnauthorized api功能未授权，请确认小程序已获得该接口
	ErrUnauthorized weapp.Code = 48001
)

// 消息体
type message struct {
	ToUser          string  `json:"touser"`  // user openid
	MsgType         MsgType `json:"msgtype"` // text | image | link | miniprogrampage
	Text            MsgText `json:"text,omitempty"`
	Image           MsgImg  `json:"image,omitempty"`
	Link            MsgLink `json:"link,omitempty"`
	Miniprogrampage MsgCard `json:"miniprogrampage,omitempty"`
}

// MsgText 文本消息
// 支持添加可跳转小程序的文字链
type MsgText struct {
	Content string `json:"content"`
}

// MsgImg 图片消息
type MsgImg struct {
	MediaID string `json:"media_id"`
}

// MsgCard 卡片消息
type MsgCard struct {
	Title        string `json:"title"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// MsgLink 图文链接消息
type MsgLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}

// SendTo 发送文本消息
// @ openid 用户openid
// @ token 微信 access_token
func (msg MsgText) SendTo(openid, token string) (wres weapp.Response, err error) {

	m := message{
		ToUser:  openid,
		MsgType: "text",
		Text:    msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// SendTo 发送图片消息
// @ openid 用户openid
// @ token 微信 access_token
func (msg MsgImg) SendTo(openid, token string) (wres weapp.Response, err error) {

	m := message{
		ToUser:  openid,
		MsgType: "image",
		Image:   msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// SendTo 发送图文链接消息
// @ openid 用户openid
// @ token 微信 access_token
func (msg MsgLink) SendTo(openid, token string) (wres weapp.Response, err error) {

	m := message{
		ToUser:  openid,
		MsgType: "link",
		Link:    msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// SendTo 发送卡片消息
// @ openid 用户openid
// @ token 微信 access_token
func (msg MsgCard) SendTo(openid, token string) (wres weapp.Response, err error) {

	m := message{
		ToUser:          openid,
		MsgType:         "miniprogrampage",
		Miniprogrampage: msg,
	}

	body, err := json.Marshal(m)
	if err != nil {
		return
	}

	return send(token, string(body))
}

// send 发送消息
// @ token 微信 access_token
func send(token, body string) (wres weapp.Response, err error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+sendAPI, token)
	if err != nil {
		return
	}

	res, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(weapp.WechatServerError)
		return
	}

	err = json.NewDecoder(res.Body).Decode(&wres)
	return
}
