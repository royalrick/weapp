package order

import "github.com/medivhzhan/weapp/v3/request"

type SetMsgJumpPathRequest struct {
	Path string `json:"path"` //商户自定义跳转路径。
}

// SetMsgJumpPath 消息跳转路径设置接口
func (cli *Order) SetMsgJumpPath(req *SetMsgJumpPathRequest) (*request.CommonError, error) {

	url, err := cli.combineURI("/wxa/sec/order/set_msg_jump_path", nil, true)
	if err != nil {
		return nil, err
	}

	rsp := new(request.CommonError)
	if err := cli.request.Post(url, req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
