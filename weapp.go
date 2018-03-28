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

// 拼接 获取 session_key 的 URL
func code2url(appID, secret, code string) (string, error) {

	url, err := url.Parse(baseURL + codeToSessionAPI)
	if err != nil {
		return "", err
	}

	query := url.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
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

// Login 用户登录
// 返回 微信端 openid 和 session_key
func Login(appID, secret, code string) (string, string, error) {
	if code == "" {
		return "", "", errors.New("code can not be null")
	}

	api, err := code2url(appID, secret, code)
	if err != nil {
		return "", "", err
	}

	res, err := http.Get(api)
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
