package auth

import "github.com/medivhzhan/weapp/v3/request"

const apiGetStableAccessToken = "/cgi-bin/stable_token"

type GetStableAccessTokenRequest struct {
	// 必填 填写 client_credential
	GrantType string `json:"grant_type"`
	// 必填 小程序唯一凭证，即 AppID，可在「微信公众平台 - 设置 - 开发设置」页中获得。（需要已经成为开发者，且帐号没有异常状态）
	Appid string `json:"appid"`
	// 必填 小程序唯一凭证密钥，即 AppSecret，获取方式同 appid
	Secret string `json:"secret"`
	// 默认使用 false。1. force_refresh = false 时为普通调用模式，access_token 有效期内重复调用该接口不会更新 access_token；2. 当force_refresh = true 时为强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token
	ForceRefresh bool `json:"force_refresh"`
}

type GetStableAccessTokenResponse struct {
	request.CommonError
	// 获取到的凭证
	AccessToken string `json:"access_token"`
	// 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ExpiresIn int64 `json:"expires_in"`
}

// 获取稳定版接口调用凭据
// doc: https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html
func (cli *Auth) GetStableAccessToken(req *GetStableAccessTokenRequest) (*GetStableAccessTokenResponse, error) {

	api, err := cli.combineURI(apiGetStableAccessToken, nil, false)
	if err != nil {
		return nil, err
	}

	res := new(GetStableAccessTokenResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
