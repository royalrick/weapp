package subscribemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiGetPubTemplateKeyWordsById = "/wxaapi/newtmpl/getpubtemplatekeywords"

type GetPubTemplateKeyWordsByIdRequest struct {
	// tid	string		必填	模板标题 id，可通过接口获取
	Tid string `query:"tid"`
}
type GetPubTemplateKeyWordsByIdResponse struct {
	request.CommonError
	Count int32 `json:"count"` // 模版标题列表总数
	Data  []struct {
		Kid     int    `json:"kid"`     // 关键词 id，选用模板时需要
		Name    string `json:"name"`    // 关键词内容
		Example string `json:"example"` // 关键词内容对应的示例
		Rule    string `json:"rule"`    // 参数类型
	} `json:"data"` // 关键词列表
}

// 获取模板标题下的关键词列表
func (cli *SubscribeMessage) GetPubTemplateKeyWordsById(req *GetPubTemplateKeyWordsByIdRequest) (*GetPubTemplateKeyWordsByIdResponse, error) {

	uri, err := cli.combineURI(apiGetPubTemplateKeyWordsById, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetPubTemplateKeyWordsByIdResponse)
	if err = cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
