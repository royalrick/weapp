package weapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	appCodeAPI          = "/wxa/getwxacode"
	unlimitedAppCodeAPI = "/wxa/getwxacodeunlimit"
	QRCodeAPI           = "/cgi-bin/wxaapp/createwxaqrcode"
)

// AppCode 获取小程序码
// 可接受path参数较长 生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// path 识别二维码后进入小程序的页面链接
// width 图片宽度
// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// token access_token
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
func AppCode(path string, width int, autoColor bool, lineColor, token string) (res *http.Response, err error) {
	body := fmt.Sprintf(`{"path":"%s","width": %v,"auto_color": %v,"line_color": %s}`, path, width, autoColor, lineColor)

	return requestCode(appCodeAPI, body, token)
}

// UnlimitedAppCode 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
// scene 需要使用 decodeURIComponent 才能获取到生成二维码时传入的 scene
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
func UnlimitedAppCode(scene, page string, width int, autoColor bool, lineColor, token string) (res *http.Response, err error) {
	body := fmt.Sprintf(`{"scene": "%s","page":"%s","width": %v,"auto_color": %v,"line_color": %s}`, scene, page, width, autoColor, lineColor)

	return requestCode(unlimitedAppCodeAPI, body, token)
}

// QRCode 获取小程序二维码
// 可接受path参数较长，生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
func QRCode(path string, width int, token string) (res *http.Response, err error) {

	body := fmt.Sprintf(`{"path":"%s","width": %v}`, path, width)

	return requestCode(QRCodeAPI, body, token)
}

// 向微信服务器获取二维码
func requestCode(path, body, token string) (res *http.Response, err error) {
	api, err := url.Parse(baseURL + path)
	if err != nil {
		return res, err
	}

	query := api.Query()
	query.Set("access_token", token)
	api.RawQuery = query.Encode()

	res, err = http.Post(api.String(), "application/json", strings.NewReader(body))
	if err != nil {
		return res, err
	}

	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		var data Response
		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			return res, err
		}
		return res, errors.New(data.Errmsg)
	case header == "image/jpeg": // 返回文件
		return res, nil
	}

	return res, errors.New("unknown error when fetch app code")
}
