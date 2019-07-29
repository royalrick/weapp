package weapp

import "errors"

// commonError 微信返回错误信息
type commonError struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}

// HasError 判断是否包含错误
func (err *commonError) HasError() bool {
	return err.ErrCode != 0
}

// HasError 获取错误信息
func (err *commonError) GetError() error {
	return errors.New(err.ErrMSG)
}
