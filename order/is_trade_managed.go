package order

import "github.com/medivhzhan/weapp/v3/request"

type IsTradeManagedRequest struct {
	Appid string `json:"appid"` //待查询小程序的 appid，非服务商调用时仅能查询本账号
}

type IsTradeManagedResponse struct {
	request.CommonError
	IsTradeManaged bool `json:"is_trade_managed"`
}

// IsTradeManaged 查询小程序是否已开通发货信息管理服务
func (cli *Order) IsTradeManaged(req *IsTradeManagedRequest) (*IsTradeManagedResponse, error) {

	url, err := cli.combineURI("/wxa/sec/order/is_trade_managed", nil, true)
	if err != nil {
		return nil, err
	}

	rsp := new(IsTradeManagedResponse)
	if err := cli.request.Post(url, req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
