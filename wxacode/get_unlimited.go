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
}

// 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。
func (cli *WXACode) GetUnlimited(req *UnlimitedQRCode) (*http.Response, *request.CommonError, error) {
	return cli.generate(apiGetUnlimited, req)
}
