package order

import "github.com/medivhzhan/weapp/v3/request"

type GetOrderListPayTimeRange struct {
	BeginTime int `json:"begin_time"` //起始时间，时间戳形式，不填则视为从0开始。
	EndTime   int `json:"end_time"`   //结束时间（含），时间戳形式，不填则视为32位无符号整型的最大值。
}

type GetOrderListRequest struct {
	PayTimeRange *GetOrderListPayTimeRange `json:"pay_time_range,omitempty"` //支付时间所属范围。
	PageSize     int                       `json:"page_size,omitempty"`      //翻页时使用，返回列表的长度，默认为100。
	OrderState   int                       `json:"order_state,omitempty"`    //订单状态枚举：(1) 待发货；(2) 已发货；(3) 确认收货；(4) 交易完成；(5) 已退款。
	Openid       int                       `json:"openid,omitempty"`         //支付者openid。
	LastIndex    int                       `json:"last_index,omitempty"`     //翻页时使用，获取第一页时不用传入，如果查询结果中 has_more 字段为 true，则传入该次查询结果中返回的 last_index 字段可获取下一页。
}

type GetOrderListResponse struct {
	request.CommonError
	LastIndex string         `json:"last_index"` //翻页时使用。
	HasMore   bool           `json:"has_more"`   //是否还有更多支付单。
	OrderList []*OrderStruct `json:"order_list"`
}

// GetOrderList 查询订单列表
func (cli *Order) GetOrderList(req *GetOrderListRequest) (*GetOrderListResponse, error) {

	url, err := cli.combineURI("/wxa/sec/order/get_order_list", nil, true)
	if err != nil {
		return nil, err
	}

	rsp := new(GetOrderListResponse)
	if err := cli.request.Post(url, req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
