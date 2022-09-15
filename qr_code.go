package weapp

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp/v3/request"
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
func (cli *Client) GetQRCode(code *QRCode) (*http.Response, *request.CommonError, error) {
	api := baseURL + apiGetQrCode

	token, err := cli.AccessToken()
	if err != nil {
		return nil, nil, err
	}

	return cli.getQRCode(api, token, code)
}

func (cli *Client) getQRCode(api, token string, code *QRCode) (*http.Response, *request.CommonError, error) {
	return cli.qrCodeRequest(api, token, code)
}

// UnlimitedQRCode 小程序码参数
type UnlimitedQRCode struct {
	Scene      string `json:"scene"`
	Page       string `json:"page,omitempty"`
	CheckPath  bool   `json:"check_path"`
	EnvVersion string `json:"env_version,omitempty"`
	Width      int    `json:"width,omitempty"`
	AutoColor  bool   `json:"auto_color,omitempty"`
	LineColor  Color  `json:"line_color,omitempty"`
	IsHyaline  bool   `json:"is_hyaline,omitempty"`
}

// Get 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
func (cli *Client) GetUnlimitedQRCode(code *UnlimitedQRCode) (*http.Response, *request.CommonError, error) {
	api := baseURL + apiGetUnlimitedQRCode

	token, err := cli.AccessToken()
	if err != nil {
		return nil, nil, err
	}

	return cli.getUnlimitedQRCode(api, token, code)
}

func (cli *Client) getUnlimitedQRCode(api, token string, code *UnlimitedQRCode) (*http.Response, *request.CommonError, error) {
	return cli.qrCodeRequest(api, token, code)
}

// QRCodeCreator 二维码创建器
type QRCodeCreator struct {
	Path  string `json:"path"`            // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width int    `json:"width,omitempty"` // 二维码的宽度，单位 px。最小 280px，最大 1280px
}

// Create 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// 通过该接口生成的小程序码，永久有效，有数量限制
func (cli *Client) CreateQRCode(creator *QRCodeCreator) (*http.Response, *request.CommonError, error) {
	api := baseURL + apiCreateQRCode

	token, err := cli.AccessToken()
	if err != nil {
		return nil, nil, err
	}

	return cli.createQRCode(api, token, creator)
}

func (cli *Client) createQRCode(api, token string, creator *QRCodeCreator) (*http.Response, *request.CommonError, error) {
	return cli.qrCodeRequest(api, token, creator)
}

// 向微信服务器获取二维码
// 返回 HTTP 请求实例
func (cli *Client) qrCodeRequest(api, token string, params interface{}) (*http.Response, *request.CommonError, error) {

	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, nil, err
	}

	res, err := cli.request.PostWithBody(url, params)
	if err != nil {
		return nil, nil, err
	}

	response := new(request.CommonError)
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
