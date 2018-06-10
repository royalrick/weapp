package weapp

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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

// Code 微信服务器返回状态码
type Code int

// Response 请求微信返回基础数据
type Response struct {
	Errcode Code   `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
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

type loginForm struct {
	Response
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// Login 用户登录
// 返回 微信端 openid 和 session_key
func Login(appID, secret, code string) (string, string, error) {
	if code == "" {
		return "", "", errors.New("code can not be null")
	}

	api, err := code2url(appID, secret, code)
	if err != nil {
		return "", "", err
	}

	res, err := http.Get(api)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", "", errors.New(WeChatServerError)
	}

	var data loginForm

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", "", err
	}

	if data.Errcode != 0 {
		return "", "", errors.New(data.Errmsg)
	}

	return data.Openid, data.SessionKey, nil
}

// PhoneNumber 解密后的用户手机号码信息
type PhoneNumber struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       watermark `json:"watermark"`
}

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// DecodePhoneNumber 解密手机号码
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @data 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
func DecodePhoneNumber(ssk, data, iv string) (phone PhoneNumber, err error) {

	dSsk, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	dIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(dSsk)
	if err != nil {
		return
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("cipher too short")
		return
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		err = errors.New("cipher is not a multiple of the block size")
		return
	}

	mode := cipher.NewCBCDecrypter(block, []byte(dIv))
	mode.CryptBlocks(ciphertext, ciphertext)

	bts := util.PKCS5UnPadding([]byte(ciphertext))

	err = json.Unmarshal(bts, &phone)

	return
}
