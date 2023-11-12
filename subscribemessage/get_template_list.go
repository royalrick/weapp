package subscribemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiGetTemplateList = "/wxaapi/newtmpl/gettemplate"

type GetTemplateListResponse struct {
	request.CommonError
	Data []struct {
		Pid     string `json:"priTmplId"` // 添加至帐号下的模板 id，发送小程序订阅消息时所需
		Title   string `json:"title"`     // 模版标题
		Content string `json:"content"`   // 模版内容
		Example string `json:"example"`   // 模板内容示例
		Type    int32  `json:"type"`      // 模版类型，2 为一次性订阅，3 为长期订阅
	} `json:"data"` // 个人模板列表
}

// 获取帐号下已存在的模板列表
func (cli *SubscribeMessage) GetTemplateList() (*GetTemplateListResponse, error) {
	uri, err := cli.combineURI(apiGetTemplateList, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetTemplateListResponse)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
