package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiGoodsResetAudit = "/wxaapi/broadcast/goods/resetaudit"

type GoodsResetAuditRequest struct {
	// 必填 商品ID
	GoodsId int64 `json:"goodsId"`
	// 必填 审核单ID
	AuditId int64 `json:"auditId"`
}

type GoodsResetAuditResponse struct {
	request.CommonError
}

// 撤回商品审核
func (cli *LiveBroadcast) GoodsResetAudit(req *GoodsResetAuditRequest) (*GoodsResetAuditResponse, error) {

	api, err := cli.combineURI(apiGoodsResetAudit, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GoodsResetAuditResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
