package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiURLLink = "/wxa/generate_urllink"
)

// URLLink 有效期类型
type URLLinkExpireType = uint8

const (
	URLLinkExpireTypeTime URLLinkExpireType = 0
	URLLinkExpireTypeDate URLLinkExpireType = 1
)

type URLLinkRequest struct {
	// 非必填	通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query 。path 为空时会跳转小程序主页
	Path string `json:"path"`
	// 非必填	通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%
	Query string `json:"query"`
	// 非必填	生成的 URL Link 类型，到期失效：true，永久有效：false
	IsExpired bool `json:"is_expire"`
	// 非必填	小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1
	ExpireType URLLinkExpireType `json:"expire_type"`
	// 非必填	到期失效的 URL Link 的失效时间，为 Unix 时间戳。生成的到期失效 URL Link 在该时间前有效。最长有效期为1年。expire_type 为 0 必填
	ExpireTime int64 `json:"expire_time"`
	// 非必填	到期失效的URL Link的失效间隔天数。生成的到期失效URL Link在该间隔时间到达前有效。最长间隔天数为365天。expire_type 为 1 必填
	ExpireInterval int64 `json:"expire_interval"`
	// 非必填	云开发静态网站自定义 H5 配置参数，可配置中转的云开发 H5 页面。不填默认用官方 H5 页面
	CloudBase *URLLinkCloudBase `json:"cloud_base"`
	// 非必填       要打开的小程序版本。正式版为 "release"，体验版为"trial"，开发版为"develop"。默认值"release"
	EnvVersion string `json:"env_version,omitempty"` 
}

type URLLinkCloudBase struct {
	// 必填	云开发环境
	Env string `json:"env"`
	// 非必填	静态网站自定义域名，不填则使用默认域名
	Domain string `json:"domain"`
	// 非必填	云开发静态网站 H5 页面路径，不可携带 query
	Path string `json:"path"`
	// 非必填	云开发静态网站 H5 页面 query 参数，最大 1024 个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%
	Query string `json:"query"`
	// 非必填	第三方批量代云开发时必填，表示创建该 env 的 appid （小程序/第三方平台）
	ResourceAppid string `json:"resource_appid"`
}

type URLLinkResponse struct {
	request.CommonError
	// 生成的小程序 URL Link
	URLLink string `json:"url_link"`
}

// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景。通过该接口，可以选择生成到期失效和永久有效的小程序链接，目前仅针对国内非个人主体的小程序开放
func (cli *Client) GenerateURLLink(link *URLLinkRequest) (*URLLinkResponse, error) {
	api := baseURL + apiURLLink

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.generateURLLink(api, token, link)
}

func (cli *Client) generateURLLink(api, token string, link *URLLinkRequest) (*URLLinkResponse, error) {
	uri, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(URLLinkResponse)
	err = cli.request.Post(uri, link, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
