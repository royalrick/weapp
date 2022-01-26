package weapp

import (
	"fmt"

	"github.com/medivhzhan/weapp/v3/request"
)

const (
	apiLogin          = "/sns/jscode2session"
	apiGetAccessToken = "/cgi-bin/token"
	apiGetPaidUnionID = "/wxa/getpaidunionid"
)

// LoginResponse 返回给用户的数据
type LoginResponse struct {
	request.CommonError
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	// 用户在开放平台的唯一标识符
	// 只在满足一定条件的情况下返回
	UnionID string `json:"unionid"`
}

// Login 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
//
// appID 小程序 appID
// secret 小程序的 app secret
// code 小程序登录时获取的 code
func (cli *Client) Login(code string) (*LoginResponse, error) {
	api := baseURL + apiLogin
	return cli.login(code, api)
}

func (cli *Client) login(code, api string) (*LoginResponse, error) {
	queries := requestQueries{
		"appid":      cli.appid,
		"secret":     cli.secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(LoginResponse)
	if err := cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// TokenResponse 获取 access_token 成功返回数据
type TokenResponse struct {
	request.CommonError
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   uint   `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
}

// access_token 缓存 KEY
func (cli *Client) tokenCacheKey() string {
	return fmt.Sprintf("weapp.%s.access.token", cli.appid)
}

func (cli *Client) GetAccessToken() (*TokenResponse, error) {

	queries := requestQueries{
		"appid":      cli.appid,
		"secret":     cli.secret,
		"grant_type": "client_credential",
	}

	api := baseURL + apiGetAccessToken
	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(TokenResponse)
	if err := cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPaidUnionIDResponse response data
type GetPaidUnionIDResponse struct {
	request.CommonError
	UnionID string `json:"unionid"`
}

// GetPaidUnionID 用户支付完成后，通过微信支付订单号（transaction_id）获取该用户的 UnionId
func (cli *Client) GetPaidUnionID(openID, transactionID string) (*GetPaidUnionIDResponse, error) {
	api := baseURL + apiGetPaidUnionID
	accessToken, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}
	return cli.getPaidUnionID(accessToken, openID, transactionID, api)
}

func (cli *Client) getPaidUnionID(accessToken, openID, transactionID, api string) (*GetPaidUnionIDResponse, error) {
	queries := requestQueries{
		"openid":         openID,
		"access_token":   accessToken,
		"transaction_id": transactionID,
	}

	return cli.getPaidUnionIDRequest(api, queries)
}

// GetPaidUnionIDWithMCH 用户支付完成后，通过微信支付商户订单号和微信支付商户号（out_trade_no 及 mch_id）获取该用户的 UnionId
func (cli *Client) GetPaidUnionIDWithMCH(openID, outTradeNo, mchID string) (*GetPaidUnionIDResponse, error) {
	api := baseURL + apiGetPaidUnionID

	accessToken, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getPaidUnionIDWithMCH(accessToken, openID, outTradeNo, mchID, api)
}

func (cli *Client) getPaidUnionIDWithMCH(accessToken, openID, outTradeNo, mchID, api string) (*GetPaidUnionIDResponse, error) {
	queries := requestQueries{
		"openid":       openID,
		"mch_id":       mchID,
		"out_trade_no": outTradeNo,
		"access_token": accessToken,
	}

	return cli.getPaidUnionIDRequest(api, queries)
}

func (cli *Client) getPaidUnionIDRequest(api string, queries requestQueries) (*GetPaidUnionIDResponse, error) {
	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(GetPaidUnionIDResponse)
	if err := cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}
