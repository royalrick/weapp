package weapp

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/medivhzhan/weapp/util"
)

const (
	// BaseURL 微信请求基础URL
	BaseURL = "https://api.weixin.qq.com"

	codeAPI = "/sns/jscode2session"
)

const (
	// WeChatServerError 微信服务器错误时返回返回消息
	WeChatServerError = "微信服务器发生错误"
)

// Response 请求微信返回基础数据
type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// PhoneNumber 解密后的用户手机号码信息
type PhoneNumber struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       watermark `json:"watermark"`
}

// Userinfo 解密后的用户信息
type Userinfo struct {
	OpenID    string    `json:"openId"`
	Nickname  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	Province  string    `json:"province"`
	Language  string    `json:"language"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Avatar    string    `json:"avatarUrl"`
	UnionID   string    `json:"unionId"`
	Watermark watermark `json:"watermark"`
}

// LoginResponse 返回给用户的数据
type LoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	// 用户在开放平台的唯一标识符
	// 只在满足一定条件的情况下返回
	UnionID string `json:"unionid"`
}

type loginResponse struct {
	Response
	LoginResponse
}

// Login 用户登录
// @appID 小程序 appID
// @secret 小程序的 app secret
// @code 小程序登录时获取的 code
func Login(appID, secret, code string) (lres LoginResponse, err error) {
	if code == "" {
		err = errors.New("code can not be null")
		return
	}

	api, err := code2url(appID, secret, code)
	if err != nil {
		return
	}

	res, err := http.Get(api)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(WeChatServerError)
		return
	}

	var data loginResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	if data.Errcode != 0 {
		err = errors.New(data.Errmsg)
		return
	}

	lres = data.LoginResponse
	return
}

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// DecryptPhoneNumber 解密手机号码
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @data 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
func DecryptPhoneNumber(ssk, data, iv string) (phone PhoneNumber, err error) {
	bts, err := util.CBCDecrypt(ssk, data, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &phone)
	return
}

type group struct {
	GID string `json:"openGId"`
}

// DecryptShareInfo 解密转发信息的加密数据
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @data 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
//
// @gid 小程序唯一群号
func DecryptShareInfo(ssk, data, iv string) (string, error) {

	bts, err := util.CBCDecrypt(ssk, data, iv)
	if err != nil {
		return "", err
	}

	var g group
	err = json.Unmarshal(bts, &g)
	return g.GID, err
}

// DecryptUserInfo 解密用户信息
//
// @rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// @encryptedData 包括敏感数据在内的完整用户信息的加密数据
// @signature 使用 sha1( rawData + session_key ) 得到字符串，用于校验用户信息
// @iv 加密算法的初始向量
// @ssk 微信 session_key
func DecryptUserInfo(rawData, encryptedData, signature, iv, ssk string) (ui Userinfo, err error) {

	if ok := util.Validate(rawData, ssk, signature); !ok {
		err = errors.New("数据校验失败")
		return
	}

	bts, err := util.CBCDecrypt(ssk, encryptedData, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(bts, &ui)
	return
}

// 拼接 获取 session_key 的 URL
func code2url(appID, secret, code string) (string, error) {

	url, err := url.Parse(BaseURL + codeAPI)
	if err != nil {
		return "", err
	}

	query := url.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
	query.Set("js_code", code)
	query.Set("grant_type", "authorization_code")

	url.RawQuery = query.Encode()

	return url.String(), nil
}
