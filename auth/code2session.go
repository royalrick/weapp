package auth

import "github.com/medivhzhan/weapp/v3/request"

const apiCode2Session = "/sns/jscode2session"

type Code2SessionRequest struct {
	// 必填 小程序 appId
	Appid string `query:"appid"`
	// 必填 小程序 appSecret
	Secret string `query:"secret"`
	// 必填 登录时获取的 code
	JsCode string `query:"js_code"`
	// 必填 授权类型，此处只需填写 authorization_code
	GrantType string `query:"grant_type"`
}

type Code2SessionResponse struct {
	request.CommonError
	// 用户唯一标识
	Openid string `json:"openid"`
	// 会话密钥
	SessionKey string `json:"session_key"`
	// 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回
	Unionid string `json:"unionid"`
}

// 登录凭证校验。
// 通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
func (cli *Auth) Code2Session(req *Code2SessionRequest) (*Code2SessionResponse, error) {

	api, err := cli.combineURI(apiCode2Session, req, false)
	if err != nil {
		return nil, err
	}

	res := new(Code2SessionResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
