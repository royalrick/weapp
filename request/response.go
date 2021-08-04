package request

import (
	"fmt"
)

// CommonError 微信返回错误信息
type CommonError struct {
	ErrCode int    `json:"errcode"` // 	错误码
	ErrMSG  string `json:"errmsg"`  // 	错误描述
}

// GetResponseError 获取微信服务器错返回误信息
func (err CommonError) GetResponseError() error {
	if err.ErrCode != 0 {
		return fmt.Errorf("wechat server error: code[%d] msg[%s]", err.ErrCode, err.ErrMSG)
	}

	return nil
}

// CommonResult 微信返回错误信息
type CommonResult struct {
	ResultCode int    `json:"resultcode"` // 	错误码
	ResultMsg  string `json:"resultmsg"`  // 	错误描述
}

// GetResponseError 获取微信服务器错返回误信息
func (err CommonResult) GetResponseError() error {

	if err.ResultCode != 0 {
		return fmt.Errorf("wechat server error: code[%d] msg[%s]", err.ResultCode, err.ResultMsg)
	}

	return nil
}
