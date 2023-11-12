package subscribemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiGetPubTemplateTitleList = "/wxaapi/newtmpl/getpubtemplatetitles"

type GetPubTemplateTitleListRequest struct {
	// 必填	类目 id，多个用逗号隔开
	Ids string `query:"ids"`
	// 必填	用于分页，表示从 start 开始。从 0 开始计数。
	Start int `query:"start"`
	// 必填	用于分页，表示拉取 limit 条记录。最大为 30。
	Limit int `query:"limit"`
}

// 帐号所属类目下的公共模板标题
type GetPubTemplateTitleListResponse struct {
	request.CommonError
	Count uint `json:"count"` // 模版标题列表总数
	Data  []struct {
		Tid        int    `json:"tid"`        // 模版标题 id
		Title      string `json:"title"`      // 模版标题
		Type       int32  `json:"type"`       // 模版类型，2 为一次性订阅，3 为长期订阅
		CategoryId string `json:"categoryId"` // 模版所属类目 id
	} `json:"data"` // 模板标题列表
}

// 获取帐号所属类目下的公共模板标题
func (cli *SubscribeMessage) GetPubTemplateTitleList(req *GetPubTemplateTitleListRequest) (*GetPubTemplateTitleListResponse, error) {

	uri, err := cli.combineURI(apiGetPubTemplateTitleList, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetPubTemplateTitleListResponse)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
