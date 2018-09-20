package template

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
)

const uniformSendAPI = "/cgi-bin/message/wxopen/template/uniform_send"

// WeappTemplateMsg 小程序模板消息
type WeappTemplateMsg struct {
	TemplateID      string `json:"template_id"`
	Page            string `json:"page"`
	FormID          string `json:"form_id"`
	Data            Data   `json:"data"`
	EmphasisKeyword string `json:"emphasis_keyword,omitempty"`
}

// Data 模板消息内容
type Data map[string]Keyword

// Keyword 关键字
type Keyword struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// MPTemplateMsg 公众号模板消息
type MPTemplateMsg struct {
	AppID       string             `json:"appid"`
	TemplateID  string             `json:"template_id"`
	URL         string             `json:"url"`
	Miniprogram Miniprogram        `json:"miniprogram"`
	Data        map[string]Keyword `json:"data"`
}

// Miniprogram 小程序
type Miniprogram struct {
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

// UniformMsg 统一服务消息
type UniformMsg struct {
	ToUser           string           `json:"touser"` // 用户 openid
	MPTemplateMsg    MPTemplateMsg    `json:"mp_template_msg"`
	WeappTemplateMsg WeappTemplateMsg `json:"weapp_template_msg"`
}

// Send 统一服务消息
//
// @token access_token
func (msg UniformMsg) Send(token string) error {
	api, err := util.TokenAPI(weapp.BaseURL+uniformSendAPI, token)
	if err != nil {
		return err
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := http.Post(api, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(weapp.WeChatServerError)
		return err
	}

	var resp weapp.Response
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	if resp.Errcode != 0 {
		return errors.New(resp.Errmsg)
	}

	return nil
}
