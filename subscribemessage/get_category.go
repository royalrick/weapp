package subscribemessage

import "github.com/medivhzhan/weapp/v3/request"

const apiGetCategory = "/wxaapi/newtmpl/getcategory"

type GetCategoryResponse struct {
	request.CommonError
	Data []struct {
		ID   int    `json:"id"`   // 类目id，查询公共库模版时需要
		Name string `json:"name"` // 类目的中文名
	} `json:"data"` // 类目列表
}

// 删除帐号下的某个模板
func (cli *SubscribeMessage) GetCategory() (*GetCategoryResponse, error) {
	api, err := cli.combineURI(apiGetCategory, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetCategoryResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
