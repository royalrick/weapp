package weapp

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	baseURL          = "https://api.weixin.qq.com"
	codeToSessionAPI = "/sns/jscode2session"
)

// Response 请求微信返回基础数据
type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// WeApp 小程序实例
type WeApp struct {
	AppID  string
	Secret string
	Token  string
	AesKey string
}

var app WeApp

// Init 初始化小程序
func Init(appID, secret, token, aesKey string) {
	app.AppID = appID
	app.Secret = secret
	app.Token = token
	app.AesKey = aesKey
}

// code2url 拼接 获取 session_key 的 URL
func (app WeApp) code2url(code string) (string, error) {

	url, err := url.Parse(baseURL + codeToSessionAPI)
	if err != nil {
		return "", err
	}

	query := url.Query()

	query.Set("appid", app.AppID)
	query.Set("secret", app.Secret)
	query.Set("js_code", code)
	query.Set("grant_type", "authorization_code")

	url.RawQuery = query.Encode()

	return url.String(), nil
}

type loginForm struct {
	Response
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// Login 小程序登陆
// 返回 微信端 openid 和 session_key
func Login(code string) (string, string, error) {
	if code == "" {
		return "", "", errors.New("code can not be null")
	}

	url, err := app.code2url(code)
	if err != nil {
		return "", "", err
	}

	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", "", errors.New("error when request wechat server")
	}

	var data loginForm

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", "", err
	}

	if data.Errcode != 0 {
		return "", "", errors.New(data.Errmsg)
	}

	return data.Openid, data.SessionKey, nil
}
