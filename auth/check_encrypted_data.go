package auth

import "github.com/medivhzhan/weapp/v3/request"

const apiCheckEncryptedData = "/wxa/business/checkencryptedmsg"

type CheckEncryptedDataRequest struct {
	// 必填 加密数据的sha256，通过Hex（Base16）编码后的字符串
	EncryptedMsgHash string `json:"encrypted_msg_hash"`
}

type CheckEncryptedDataResponse struct {
	request.CommonError
	// 是否是合法的数据
	Valid bool `json:"vaild"`
	// 加密数据生成的时间戳
	CreateTime int64 `json:"create_time"`
}

// 检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (cli *Auth) CheckEncryptedData(req *CheckEncryptedDataRequest) (*CheckEncryptedDataResponse, error) {

	api, err := cli.combineURI(apiCheckEncryptedData, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(CheckEncryptedDataResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
