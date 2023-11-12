package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsAudit = "/wxaapi/broadcast/goods/audit"

type GoodsAuditRequest struct {
	// 商品ID
	GoodsId int64 `json:"goodsId"`
}

type GoodsAuditResponse struct {
	request.CommonError
	// 审核单ID
	AuditId int64 `json:"auditId"`
}

// 重新提交审核
func (cli *LiveBroadcast) GoodsAudit(req *GoodsAuditRequest) (*GoodsAuditResponse, error) {

	api, err := cli.combineURI(apiGoodsAudit, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsAuditResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
