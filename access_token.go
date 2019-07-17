package weapp

import (
	"encoding/json"
	"net/http"
)

const apiGetAccessToken = "/cgi-bin/token"

// TokenResponse 获取 access_token 成功返回数据
type TokenResponse struct {
	baseResponse
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   uint   `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
}

// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）。
// 调调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存，注意缓存。
func GetAccessToken(appID, secret string) (*TokenResponse, error) {

	params := map[string]string{
		"appid":      appID,
		"secret":     secret,
		"grant_type": "client_credential",
	}

	api, err := encodeURL(baseURL+apiGetAccessToken, params)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := new(TokenResponse)
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
