package weapp

const (
	apiLogin          = "/sns/jscode2session"
	apiGetAccessToken = "/cgi-bin/token"
	apiGetPaidUnionID = "/wxa/getpaidunionid"
)

// LoginResponse 返回给用户的数据
type LoginResponse struct {
	CommonError
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
func Login(appID, secret, code string) (*LoginResponse, error) {
	api := baseURL + apiLogin

	return login(appID, secret, code, api)
}

func login(appID, secret, code, api string) (*LoginResponse, error) {
	queries := requestQueries{
		"appid":      appID,
		"secret":     secret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(LoginResponse)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// TokenResponse 获取 access_token 成功返回数据
type TokenResponse struct {
	CommonError
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   uint   `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值。
}

// GetAccessToken 获取小程序全局唯一后台接口调用凭据（access_token）。
// 调调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存，注意缓存。
func GetAccessToken(appID, secret string) (*TokenResponse, error) {
	api := baseURL + apiGetAccessToken
	return getAccessToken(appID, secret, api)
}

func getAccessToken(appID, secret, api string) (*TokenResponse, error) {

	queries := requestQueries{
		"appid":      appID,
		"secret":     secret,
		"grant_type": "client_credential",
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(TokenResponse)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPaidUnionIDResponse response data
type GetPaidUnionIDResponse struct {
	CommonError
	UnionID string `json:"unionid"`
}

// GetPaidUnionID 用户支付完成后，通过微信支付订单号（transaction_id）获取该用户的 UnionId，
func GetPaidUnionID(accessToken, openID, transactionID string) (*GetPaidUnionIDResponse, error) {
	api := baseURL + apiGetPaidUnionID
	return getPaidUnionID(accessToken, openID, transactionID, api)
}

func getPaidUnionID(accessToken, openID, transactionID, api string) (*GetPaidUnionIDResponse, error) {
	queries := requestQueries{
		"openid":         openID,
		"access_token":   accessToken,
		"transaction_id": transactionID,
	}

	return getPaidUnionIDRequest(api, queries)
}

// GetPaidUnionIDWithMCH 用户支付完成后，通过微信支付商户订单号和微信支付商户号（out_trade_no 及 mch_id）获取该用户的 UnionId，
func GetPaidUnionIDWithMCH(accessToken, openID, outTradeNo, mchID string) (*GetPaidUnionIDResponse, error) {
	api := baseURL + apiGetPaidUnionID
	return getPaidUnionIDWithMCH(accessToken, openID, outTradeNo, mchID, api)
}

func getPaidUnionIDWithMCH(accessToken, openID, outTradeNo, mchID, api string) (*GetPaidUnionIDResponse, error) {
	queries := requestQueries{
		"openid":       openID,
		"mch_id":       mchID,
		"out_trade_no": outTradeNo,
		"access_token": accessToken,
	}

	return getPaidUnionIDRequest(api, queries)
}

func getPaidUnionIDRequest(api string, queries requestQueries) (*GetPaidUnionIDResponse, error) {
	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(GetPaidUnionIDResponse)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}
