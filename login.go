package weapp

import (
	"encoding/json"
	"net/http"
)

const (
	apiLogin = "/sns/jscode2session"
)

// LoginResponse 返回给用户的数据
type LoginResponse struct {
	commonError
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	// 用户在开放平台的唯一标识符
	// 只在满足一定条件的情况下返回
	UnionID string `json:"unionid"`
}

// Login 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
//
// @appID 小程序 appID
// @secret 小程序的 app secret
// @code 小程序登录时获取的 code
func Login(appID, secret, code string) (*LoginResponse, error) {
	params := map[string]string{
		"appid":      appID,
		"secret":     secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	api, err := encodeURL(baseURL+apiLogin, params)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := new(LoginResponse)
	if err = json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
