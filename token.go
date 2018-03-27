package weapp

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

var (
	tokenAPI = "/cgi-bin/token"
)

// 获取 access_token 成功返回数据
type accessTokenResponse struct {
	Response
	AccessToken string `json:"access_token"`
	ExpireIn    time.Duration    `json:"expires_in"`
}

// 通过微信服务器获取 access_token 以及其有效期
func (app WeApp) AccessToken() (string, time.Duration, error) {
	url, err := url.Parse(baseURL + tokenAPI)
	if err != nil {
		return "", 0, err
	}

	query := url.Query()

	query.Set("appid", app.AppID)
	query.Set("secret", app.Secret)
	query.Set("grant_type", "client_credential")

	url.RawQuery = query.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		// dev: log
		return "", 0, errors.New("fetch access_token failed")
	}

	var data accessTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", 0, err
	}

	if data.Errcode != 0 {
		// dev: log
		return "", 0, errors.New(data.Errmsg)
	}

	return data.AccessToken, data.ExpireIn, nil
}