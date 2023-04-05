package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiURLScheme = "/wxa/generatescheme"
)

type URLSchemeRequest struct {
	// 跳转到的目标小程序信息。
	SchemedInfo *SchemedInfo `json:"jump_wxa,omitempty"`
	// 成的scheme码类型，到期失效：true，永久有效：false。
	IsExpire bool `json:"is_expire,omitempty"`
	// 到期失效的scheme码的失效时间，为Unix时间戳。生成的到期失效scheme码在该时间前有效。最长有效期为1年。生成到期失效的scheme时必填。
	ExpireTime int64 `json:"expire_time,omitempty"`
}

type SchemedInfo struct {
	// 通过scheme码进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带query。path为空时会跳转小程序主页。
	Path string `json:"path"`
	// 通过scheme码进入小程序时的query，最大128个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~
	Query string `json:"query"`
	// 非必填       要打开的小程序版本。正式版为 "release"，体验版为"trial"，开发版为"develop"。默认值"release"
        EnvVersion string `json:"env_version,omitempty"`
}

type URLSchemeResponse struct {
	request.CommonError
	// 生成的小程序scheme码
	Openlink string `json:"openlink"`
}

// 获取小程序scheme码，适用于短信、邮件、外部网页等拉起小程序的业务场景。
func (cli *Client) GenerateURLScheme(scheme *URLSchemeRequest) (*URLSchemeResponse, error) {
	api := baseURL + apiURLScheme

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.generateURLScheme(api, token, scheme)
}

func (cli *Client) generateURLScheme(api, token string, scheme *URLSchemeRequest) (*URLSchemeResponse, error) {
	uri, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(URLSchemeResponse)
	err = cli.request.Post(uri, scheme, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
