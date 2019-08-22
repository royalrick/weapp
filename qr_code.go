package weapp

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	apiGetQrCode          = "/wxa/getwxacode"
	apiGetUnlimitedQRCode = "/wxa/getwxacodeunlimit"
	apiCreateQRCode       = "/cgi-bin/wxaapp/createwxaqrcode"
)

// Color QRCode color
type Color struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

// QRCode 小程序码参数
type QRCode struct {
	Path      string `json:"path"`
	Width     int    `json:"width,omitempty"`
	AutoColor bool   `json:"auto_color,omitempty"`
	LineColor Color  `json:"line_color,omitempty"`
	IsHyaline bool   `json:"is_hyaline,omitempty"`
}

// Get 获取小程序码
// 可接受path参数较长 生成个数受限 永久有效 适用于需要的码数量较少的业务场景
//
// token 微信access_token
func (code *QRCode) Get(token string) (*http.Response, *CommonError, error) {
	api := baseURL + apiGetQrCode
	return code.get(api, token)
}

func (code *QRCode) get(api, token string) (*http.Response, *CommonError, error) {
	return qrCodeRequest(api, token, code)
}

// UnlimitedQRCode 小程序码参数
type UnlimitedQRCode struct {
	Scene     string `json:"scene"`
	Page      string `json:"page,omitempty"`
	Width     int    `json:"width,omitempty"`
	AutoColor bool   `json:"auto_color,omitempty"`
	LineColor Color  `json:"line_color,omitempty"`
	IsHyaline bool   `json:"is_hyaline,omitempty"`
}

// Get 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
//
// token 微信access_token
func (code *UnlimitedQRCode) Get(token string) (*http.Response, *CommonError, error) {
	api := baseURL + apiGetUnlimitedQRCode
	return code.get(api, token)
}

func (code *UnlimitedQRCode) get(api, token string) (*http.Response, *CommonError, error) {
	return qrCodeRequest(api, token, code)
}

// QRCodeCreator 二维码创建器
type QRCodeCreator struct {
	Path  string `json:"path"`            // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width int    `json:"width,omitempty"` // 二维码的宽度，单位 px。最小 280px，最大 1280px
}

// Create 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// 通过该接口生成的小程序码，永久有效，有数量限制
//
// token 微信access_token
func (creator *QRCodeCreator) Create(token string) (*http.Response, *CommonError, error) {
	api := baseURL + apiCreateQRCode
	return creator.create(api, token)
}

func (creator *QRCodeCreator) create(api, token string) (*http.Response, *CommonError, error) {
	return qrCodeRequest(api, token, creator)
}

// 向微信服务器获取二维码
// 返回 HTTP 请求实例
func qrCodeRequest(api, token string, params interface{}) (*http.Response, *CommonError, error) {

	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, nil, err
	}

	res, err := postJSONWithBody(url, params)
	if err != nil {
		return nil, nil, err
	}

	response := new(CommonError)
	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		if err := json.NewDecoder(res.Body).Decode(response); err != nil {
			res.Body.Close()
			return nil, nil, err
		}
		return res, response, nil

	case strings.HasPrefix(header, "image"): // 返回文件
		return res, response, nil

	default:
		res.Body.Close()
		return nil, nil, errors.New("invalid response header: " + header)
	}
}
