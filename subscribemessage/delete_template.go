package subscribemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiDeleteTemplate = "/wxaapi/newtmpl/deltemplate"

type DeleteTemplateRequest struct {
	// 必填	要删除的模板id
	PriTmplId string `json:"priTmplId"`
}

// 删除帐号下的某个模板
func (cli *SubscribeMessage) DeleteTemplate(req *DeleteTemplateRequest) (*request.CommonError, error) {
	api, err := cli.combineURI(apiDeleteTemplate, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
