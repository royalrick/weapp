package wxacode

import (
	"net/http"

	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetUnlimited = "/wxa/getwxacodeunlimit"

// UnlimitedQRCode 小程序码参数
type UnlimitedQRCode struct {
	Scene     string `json:"scene"`
	Page      string `json:"page,omitempty"`
	Width     int    `json:"width,omitempty"`
	AutoColor bool   `json:"auto_color,omitempty"`
	LineColor Color  `json:"line_color,omitempty"`
	IsHyaline bool   `json:"is_hyaline,omitempty"`
	// 要打开的小程序版本。正式版为 release，体验版为 trial，开发版为 develop
	EnvVersion string `json:"env_version,omitempty"`
	// 检查 page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但 page 有数量上限（60000个）请勿滥用
	CheckPath bool `json:"check_path"`
}

// 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。
func (cli *WXACode) GetUnlimited(req *UnlimitedQRCode) (*http.Response, *request.CommonError, error) {
	return cli.generate(apiGetUnlimited, req)
}
