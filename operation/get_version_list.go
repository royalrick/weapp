package operation

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetVersionList = "/wxaapi/log/get_client_version"

type GetVersionListResponse struct {
	request.CommonError
	// 版本列表
	CVList []struct {
		// 查询类型 1 代表客户端，2 代表服务直达
		Type uint8 `json:"type"`
		// 版本列表
		ClientVersionList []string `json:"client_version_list"`
	} `json:"cvlist"`
}

// 获取客户端版本
func (cli *Operation) GetVersionList() (*GetVersionListResponse, error) {

	uri, err := cli.combineURI(apiGetVersionList, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetVersionListResponse)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
