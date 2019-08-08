package weapp

import "errors"

// CommonError 微信返回错误信息
type CommonError struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}

// GetResponseError 获取微信服务器错返回误信息
func (err *CommonError) GetResponseError() error {
	if err.ErrCode == 0 {
		return nil
	}

	return errors.New(err.ErrMSG)
}
