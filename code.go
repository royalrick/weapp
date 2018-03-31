package weapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	appCodeAPI          = "/wxa/getwxacode"
	unlimitedAppCodeAPI = "/wxa/getwxacodeunlimit"
	qrCodeAPI           = "/cgi-bin/wxaapp/createwxaqrcode"
)

// AppCode 获取小程序码
// 可接受path参数较长 生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// @ path 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// @ lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// @ token 微信access_token
// @ filename 文件储存路径
func AppCode(path string, width int, autoColor bool, lineColor, token, filename string) error {

	body := fmt.Sprintf(`{"path":"%s","width": %v,"auto_color": %v,"line_color": %s}`, path, width, autoColor, lineColor)

	res, err := requestCode(appCodeAPI, body, token)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return saveCode(res, filename)
}

// UnlimitedAppCode 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
// @ scene 需要使用 decodeURIComponent 才能获取到生成二维码时传入的 scene
// @ page 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// @ lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// @ token 微信access_token
// @ filename 文件储存路径
func UnlimitedAppCode(scene, page string, width int, autoColor bool, lineColor, token, filename string) error {

	body := fmt.Sprintf(`{"scene": "%s","page":"%s","width": %v,"auto_color": %v,"line_color": %s}`, scene, page, width, autoColor, lineColor)

	res, err := requestCode(unlimitedAppCodeAPI, body, token)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	return saveCode(res, filename)
}

// QRCode 获取小程序二维码
// 可接受path参数较长，生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// @ path 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ token 微信access_token
// @ filename 文件储存路径
func QRCode(path string, width int, token, filename string) error {

	body := fmt.Sprintf(`{"path":"%s","width": %v}`, path, width)

	res, err := requestCode(qrCodeAPI, body, token)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	return saveCode(res, filename)
}

// 向微信服务器获取二维码
// 返回 HTTP 请求实例
func requestCode(path, body, token string) (res *http.Response, err error) {

	api, err := TokenAPI(BaseURL+path, token)
	if err != nil {
		return
	}

	res, err = http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
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

// 保存二维码文件
func saveCode(res *http.Response, filename string) error {

	dir := filepath.Dir(filename)

	// 查看文件夹是否存在
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		// 文件夹不存在就创建文件夹
		if err = os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0666)
}
