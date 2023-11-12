package auth

import "github.com/medivhzhan/weapp/v3/request"

const apiGetPaidUnionId = "/wxa/getpaidunionid"

type GetPaidUnionIdRequest struct {
	// 必填	支付用户唯一标识
	Openid string `query:"openid"`
	// 非必填	微信支付订单号
	TransactionId string `query:"transaction_id"`
	// 非必填	微信支付分配的商户号，和商户订单号配合使用
	MchId string `query:"mch_id"`
	// 非必填	微信支付商户订单号，和商户号配合使用
	OutTradeNo string `query:"out_trade_no"`
}

type GetPaidUnionIdResponse struct {
	request.CommonError
	// 用户唯一标识，调用成功后返回
	UnionID string `json:"unionid"`
}

// 用户支付完成后，获取该用户的 UnionId，无需用户授权。
func (cli *Auth) GetPaidUnionId(req *GetPaidUnionIdRequest) (*GetPaidUnionIdResponse, error) {

	api, err := cli.combineURI(apiGetPaidUnionId, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetPaidUnionIdResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
