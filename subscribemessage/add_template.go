package subscribemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiAddTemplate = "/wxaapi/newtmpl/addtemplate"

type AddTemplateRequest struct {
	// 必填	模板标题 id，可通过接口获取，也可登录小程序后台查看获取
	Tid string `json:"tid"`
	// 必填	开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合
	KidList []int `json:"kidList"`
	// 非必填	服务场景描述，15个字以内
	SceneDesc string `json:"sceneDesc"`
}

type AddTemplateResponse struct {
	request.CommonError
	Pid string `json:"priTmplId"` // 添加至帐号下的模板id，发送小程序订阅消息时所需
}

// 组合模板并添加至帐号下的个人模板库
func (cli *SubscribeMessage) AddTemplate(req *AddTemplateRequest) (*AddTemplateResponse, error) {

	api, err := cli.combineURI(apiAddTemplate, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(AddTemplateResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
