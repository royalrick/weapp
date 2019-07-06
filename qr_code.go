package weapp

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	apiGetAppCode          = "/wxa/getwxacode"
	apiGetUnlimitedAppCode = "/wxa/getwxacodeunlimit"
	apiCreateQRCode        = "/cgi-bin/wxaapp/createwxaqrcode"
)

// Coder 小程序码参数
type Coder struct {
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

// GetAppCode 获取小程序码
// 可接受path参数较长 生成个数受限 永久有效 适用于需要的码数量较少的业务场景
//
// @token 微信access_token
func (code Coder) GetAppCode(token string) (*http.Response, *BaseResponse, error) {
	return fetchCode(BaseURL+apiGetAppCode, token, code)
}

// GetUnlimitedAppCode 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
//
// @token 微信access_token
func (code Coder) GetUnlimitedAppCode(token string) (*http.Response, *BaseResponse, error) {
	return fetchCode(BaseURL+apiGetUnlimitedAppCode, token, code)
}

// CreateQRCode 获取小程序二维码，适用于需要的码数量较少的业务场景。
// 通过该接口生成的小程序码，永久有效，有数量限制
//
// @token 微信access_token
func (code Coder) CreateQRCode(token string) (*http.Response, *BaseResponse, error) {
	return fetchCode(BaseURL+apiCreateQRCode, token, code)
}

// 向微信服务器获取二维码
// 返回 HTTP 请求实例
func fetchCode(api, token string, params interface{}) (*http.Response, *BaseResponse, error) {

	api, err := TokenAPI(BaseURL+api, token)
	if err != nil {
		return nil, nil, err
	}

	res, err := PostJSONWithBody(api, params)
	if err != nil {
		return nil, nil, err
	}

	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		response := new(BaseResponse)
		if err := json.NewDecoder(res.Body).Decode(response); err != nil {
			res.Body.Close()
			return nil, nil, err
		}
		return res, response, nil

	case header == "image/jpeg": // 返回文件
		return res, nil, nil

	default:
		res.Body.Close()
		return nil, nil, errors.New("invalid response header: " + header)
	}
}
