package weapp

import (
	"encoding/json"
	"errors"
)

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
// ssk 通过 Login 向微信服务端请求得到的 session_key
// data 小程序通过 api 得到的加密数据(encryptedData)
// iv 小程序通过 api 得到的初始向量(iv)
func DecryptMobile(ssk, data, iv string) (*Mobile, error) {
	raw, err := decryptUserData(ssk, data, iv)
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
// ssk 通过 Login 向微信服务端请求得到的 session_key
// data 小程序通过 api 得到的加密数据(encryptedData)
// iv 小程序通过 api 得到的初始向量(iv)
//
// gid 小程序唯一群号
func DecryptShareInfo(ssk, data, iv string) (*ShareInfo, error) {

	raw, err := decryptUserData(ssk, data, iv)
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

// DecryptUserInfo 解密用户信息
//
// rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// encryptedData 包括敏感数据在内的完整用户信息的加密数据
// signature 使用 sha1( rawData + session_key ) 得到字符串，用于校验用户信息
// iv 加密算法的初始向量
// ssk 微信 session_key
func DecryptUserInfo(rawData, encryptedData, signature, iv, ssk string) (*UserInfo, error) {

	if ok := validateSignature(signature, rawData, ssk); !ok {
		return nil, errors.New("failed to validate signature")
	}

	raw, err := decryptUserData(ssk, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	info := new(UserInfo)
	if err := json.Unmarshal(raw, info); err != nil {
		return nil, err
	}

	return info, nil
}
