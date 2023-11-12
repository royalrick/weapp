package operation

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetDomainInfo = "/wxa/getwxadevinfo"

type DomainAction = string

const (
	DomainActionAll    DomainAction = "getbizdomain"
	DomainActionBiz    DomainAction = "getbizdomain"
	DomainActionServer DomainAction = "getserverdomain"
)

type GetDomainInfoRequest struct {
	// 查询配置域名的类型, 可选值如下：
	// 1. getbizdomain 返回业务域名
	// 2. getserverdomain 返回服务器域名
	// 3. 不指明返回全部
	Action DomainAction `query:"action"`
}

type GetDomainInfoResponse struct {
	request.CommonError
	// request合法域名列表
	RequestDomain []string `json:"requestdomain"`
	// socket合法域名列表
	WSRequestDomain []string `json:"wsrequestdomain"`
	// uploadFile合法域名列表
	UploadDomain []string `json:"uploaddomain"`
	// downloadFile合法域名列表
	DownloadDomain []string `json:"downloaddomain"`
	// udp合法域名列表
	UDPDomain []string `json:"udpdomain"`
	// 业务域名列表
	BizDomain []string `json:"bizdomain"`
}

// 查询域名配置
func (cli *Operation) GetDomainInfo(req *GetDomainInfoRequest) (*GetDomainInfoResponse, error) {

	uri, err := cli.combineURI(apiGetDomainInfo, req, true)
	if err != nil {
		return nil, err
	}

	res := new(GetDomainInfoResponse)
	if err := cli.request.Get(uri, res); err != nil {
		return nil, err
	}

	return res, nil
}
