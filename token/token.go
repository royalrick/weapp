// Package token 微信 access_token
package token

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
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

	params := map[string]string{
		"appid":      appID,
		"secret":     secret,
		"grant_type": "client_credential",
	}
	api, err := util.EncodeURL(weapp.BaseURL+tokenAPI, params)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.Get(api)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	res := new(response)
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return "", 0, err
	}

	if res.HasError() {
		return "", 0, res.CreateError("failed to get access token")
	}

	return res.AccessToken, time.Second * res.ExpireIn, nil
}
