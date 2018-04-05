// Package service 客服消息
package service

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

// Msg 发送的消息体
type Msg struct {
	ToUser          string  `json:"touser"`  // user openid
	MsgType         string  `json:"msgtype"` // text | image | link | miniprogrampage
	Text            MsgText `json:"text,omitempty"`
	Image           MsgImg  `json:"image,omitempty"`
	Link            MsgLink `json:"link,omitempty"`
	Miniprogrampage MsgCard `json:"miniprogrampage,omitempty"`
}

// MsgText 文本消息
// 发送文本消息时，支持添加可跳转小程序的文字链
type MsgText struct {
	Content string `json:"content"`
}

// MsgImg 图片消息
type MsgImg struct {
	MediaID string `json:"media_id"`
}

// MsgLink 图文链接消息
type MsgLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}

// MsgCard 卡片消息
type MsgCard struct {
	Title        string `json:"title"`
	Pagepath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// Send 发送消息
// @ token 微信 access_token
func (msg Msg) Send(token string) error {
	api, err := weapp.TokenAPI(weapp.BaseURL+sendAPI, token)
	if err != nil {
		return err
	}

	body, err := json.Marshal(msg)

	res, err := http.Post(api, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New(weapp.WechatServerError)
	}

	var data weapp.Response
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	if data.Errcode != 0 {
		return errors.New(data.Errmsg)
	}

	return nil
}
