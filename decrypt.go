package weapp

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/medivhzhan/weapp/v3/encrypt"
)

// DecryptUserData 解密用户数据
func (cli *Client) DecryptUserData(ssk, ciphertext, iv string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return nil, err
	}

	cipher, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	return encrypt.NewCBC(rawIV, key, cipher).Decrypt()
}

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// Mobile 解密后的用户手机号码信息
type Mobile struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       watermark `json:"watermark"`
}

// DecryptMobile 解密手机号码
//
// sessionKey 通过 Login 向微信服务端请求得到的 session_key
// encryptedData 小程序通过 api 得到的加密数据(encryptedData)
// iv 小程序通过 api 得到的初始向量(iv)
func (cli *Client) DecryptMobile(sessionKey, encryptedData, iv string) (*Mobile, error) {
	raw, err := cli.DecryptUserData(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	mobile := new(Mobile)
	if err := json.Unmarshal(raw, mobile); err != nil {
		return nil, err
	}

	return mobile, nil
}

// ShareInfo 解密后的分享信息
type ShareInfo struct {
	GID string `json:"openGId"`
}

// DecryptShareInfo 解密转发信息的加密数据
//
// sessionKey 通过 Login 向微信服务端请求得到的 session_key
// encryptedData 小程序通过 api 得到的加密数据(encryptedData)
// iv 小程序通过 api 得到的初始向量(iv)
//
// gid 小程序唯一群号
func (cli *Client) DecryptShareInfo(sessionKey, encryptedData, iv string) (*ShareInfo, error) {

	raw, err := cli.DecryptUserData(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	info := new(ShareInfo)
	if err = json.Unmarshal(raw, info); err != nil {
		return nil, err
	}

	return info, nil
}

// UserInfo 解密后的用户信息
type UserInfo struct {
	// 用户昵称
	// 用户头像图片的 URL。URL 最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640x640 的正方形头像，46 表示 46x46 的正方形头像，剩余数值以此类推。默认132），用户没有头像时该项为空。若用户更换头像，原有头像 URL 将失效。
	Avatar string `json:"avatarUrl"`
	// 用户性别
	// 0	未知
	// 1	男性
	// 2	女性
	Gender int `json:"gender"`
	// 用户所在国家
	Country string `json:"country"`
	// 用户所在城市
	City string `json:"city"`
	// 显示 country，province，city 所用的语言
	// 	en	英文
	// zh_CN	简体中文
	// zh_TW	繁体中文
	Language string `json:"language"`
	// 用户昵称
	Nickname string `json:"nickName"`
	// 用户所在省份
	Province string `json:"province"`
}

// DecryptUserInfo 解密用户信息
//
// sessionKey 微信 session_key
// rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// encryptedData 包括敏感数据在内的完整用户信息的加密数据
// signature 使用 sha1( rawData + session_key ) 得到字符串，用于校验用户信息
// iv 加密算法的初始向量
func (cli *Client) DecryptUserInfo(sessionKey, rawData, encryptedData, signature, iv string) (*UserInfo, error) {

	if !encrypt.NewSignable(false, rawData, sessionKey).IsEqual(signature) {
		return nil, errors.New("failed to validate signature")
	}

	raw, err := cli.DecryptUserData(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	info := new(UserInfo)
	if err := json.Unmarshal(raw, info); err != nil {
		return nil, err
	}

	return info, nil
}

// RunData 解密后的最近30天微信运动步数
type RunData struct {
	StepInfoList []SetpInfo `json:"stepInfoList"`
}

// SetpInfo 运动步数
type SetpInfo struct {
	Step      int   `json:"step"`
	Timestamp int64 `json:"timestamp"`
}

// DecryptRunData 解密微信运动的加密数据
//
// sessionKey 通过 Login 向微信服务端请求得到的 session_key
// encryptedData 小程序通过 api 得到的加密数据(encryptedData)
// iv 小程序通过 api 得到的初始向量(iv)
func (cli *Client) DecryptRunData(sessionKey, encryptedData, iv string) (*RunData, error) {
	raw, err := cli.DecryptUserData(sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	info := new(RunData)
	if err := json.Unmarshal(raw, info); err != nil {
		return nil, err
	}

	return info, nil
}
