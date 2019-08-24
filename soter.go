package weapp

const (
	apiVerifySignature = "/cgi-bin/soter/verify_signature"
)

// VerifySignatureResponse 生物认证秘钥签名验证请求返回数据
type VerifySignatureResponse struct {
	CommonError
	IsOk bool `json:"is_ok"`
}

// VerifySignature 生物认证秘钥签名验证
// accessToken 接口调用凭证
// openID 用户 openid
// data 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
// signature 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
func VerifySignature(token, openID, data, signature string) (*VerifySignatureResponse, error) {
	api := baseURL + apiVerifySignature
	return verifySignature(api, token, openID, data, signature)
}

func verifySignature(api, token, openID, data, signature string) (*VerifySignatureResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"openid":         openID,
		"json_string":    data,
		"json_signature": signature,
	}

	res := new(VerifySignatureResponse)
	if err := postJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
