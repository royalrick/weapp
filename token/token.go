// Package token 微信 access_token
package token

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/medivhzhan/weapp"
)

const tokenAPI = "/cgi-bin/token"

// 获取 access_token 成功返回数据
type response struct {
	weapp.Response
	AccessToken string        `json:"access_token"`
	ExpireIn    time.Duration `json:"expires_in"`
}

// AccessToken 通过微信服务器获取 access_token 以及其有效期
func AccessToken(appID, secret string) (string, time.Duration, error) {
	url, err := url.Parse(weapp.BaseURL + tokenAPI)
	if err != nil {
		return "", 0, err
	}

	query := url.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
	query.Set("grant_type", "client_credential")

	url.RawQuery = query.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", 0, errors.New(weapp.WeChatServerError)
	}

	var data response
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", 0, err
	}

	if data.Errcode != 0 {
		return "", 0, errors.New(data.Errmsg)
	}

	return data.AccessToken, time.Second * data.ExpireIn, nil
}
