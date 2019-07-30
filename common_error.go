package weapp

import "errors"

// CommonError 微信返回错误信息
type CommonError struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}

// HasError 判断是否包含错误
func (err *CommonError) HasError() bool {
	return err.ErrCode != 0
}

// GetError 获取错误信息
func (err *CommonError) GetError() error {
	return errors.New(err.ErrMSG)
}
