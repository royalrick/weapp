package wxacode

import (
	"net/http"

	"github.com/medivhzhan/weapp/v3/request"
)

const apiGet = "/wxa/getwxacode"

// Color QRCode color
type Color struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

type GetRequest struct {
	Path      string `json:"path"`
	Width     int    `json:"width,omitempty"`
	AutoColor bool   `json:"auto_color,omitempty"`
	LineColor Color  `json:"line_color,omitempty"`
	IsHyaline bool   `json:"is_hyaline,omitempty"`
}

// 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
func (cli *WXACode) QRCode(req *GetRequest) (*http.Response, *request.CommonError, error) {
	return cli.generate(apiGet, req)
}
