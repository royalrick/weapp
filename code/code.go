// Package code 小程序二维码
package code

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
)

const (
	appCodeAPI          = "/wxa/getwxacode"
	unlimitedAppCodeAPI = "/wxa/getwxacodeunlimit"
	QRCodeAPI           = "/cgi-bin/wxaapp/createwxaqrcode"
)

// QRCoder 小程序码参数
type QRCoder struct {
	Page string `json:"page,omitempty"`
	// path 识别二维码后进入小程序的页面链接
	Path string `json:"path,omitempty"`
	// width 图片宽度
	Width int `json:"width,omitempty"`
	// scene 参数数据
	Scene string `json:"scene,omitempty"`
	// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	AutoColor bool `json:"auto_color,omitempty"`
	// lineColor AutoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
	LineColor Color `json:"line_color,omitempty"`
	// isHyaline 是否需要透明底色
	IsHyaline bool `json:"is_hyaline,omitempty"`
}

// Color QRCode color
type Color struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

// AppCode 获取小程序码
// 可接受path参数较长 生成个数受限 永久有效 适用于需要的码数量较少的业务场景
//
// @token 微信access_token
func (code QRCoder) AppCode(token string) (*http.Response, error) {
	body, err := json.Marshal(code)
	if err != nil {
		return nil, err
	}

	return fetchCode(appCodeAPI, string(body), token)
}

// UnlimitedAppCode 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
//
// @token 微信access_token
func (code QRCoder) UnlimitedAppCode(token string) (*http.Response, error) {

	body, err := json.Marshal(code)
	if err != nil {
		return nil, err
	}

	return fetchCode(unlimitedAppCodeAPI, string(body), token)
}

// QRCode 获取小程序二维码
// 可接受path参数较长，生成个数受限 永久有效 适用于需要的码数量较少的业务场景
//
// @token 微信access_token
func (code QRCoder) QRCode(token string) (*http.Response, error) {

	body, err := json.Marshal(code)
	if err != nil {
		return nil, err
	}

	return fetchCode(QRCodeAPI, string(body), token)
}

// 向微信服务器获取二维码
// 返回 HTTP 请求实例
func fetchCode(path, body, token string) (res *http.Response, err error) {

	api, err := util.TokenAPI(weapp.BaseURL+path, token)
	if err != nil {
		return
	}

	res, err = http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}

	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		var data weapp.Response
		if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
			return
		}

		return res, errors.New(data.Errmsg)
	case header == "image/jpeg": // 返回文件
		return res, nil
	default:
		err = errors.New("unknown response header: " + header)
		return
	}
}
