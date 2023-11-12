package operation

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetSceneList = "/wxaapi/log/get_scene"

type GetSceneListResponse struct {
	request.CommonError
	// 访问来源
	Scene []struct {
		// 来源中文名
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"scene"`
}

// 获取访问来源
func (cli *Operation) GetSceneList() (*GetSceneListResponse, error) {

	uri, err := cli.combineURI(apiGetSceneList, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetSceneListResponse)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
