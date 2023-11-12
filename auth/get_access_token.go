package auth

import "github.com/medivhzhan/weapp/v3/request"

const apiGetAccessToken = "/cgi-bin/token"

type GetAccessTokenRequest struct {
	// 必填 填写 client_credential
	GrantType string `query:"grant_type"`
	// 必填 小程序唯一凭证，即 AppID，可在「微信公众平台 - 设置 - 开发设置」页中获得。（需要已经成为开发者，且帐号没有异常状态）
	Appid string `query:"appid"`
	// 必填 小程序唯一凭证密钥，即 AppSecret，获取方式同 appid
	Secret string `query:"secret"`
}

type GetAccessTokenResponse struct {
	request.CommonError
	// 获取到的凭证
	AccessToken string `json:"access_token"`
	// 凭证有效时间，单位：秒。目前是7200秒之内的值。
	ExpiresIn int64 `json:"expires_in"`
}

// 获取小程序全局唯一后台接口调用凭据
// 通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
func (cli *Auth) GetAccessToken(req *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {

	api, err := cli.combineURI(apiGetAccessToken, req, false)
	if err != nil {
		return nil, err
	}

	res := new(GetAccessTokenResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
