package weapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var (
	baseURL    = "https://api.weixin.qq.com"
	codeAPI    = "/sns/jscode2session"
	appCodeAPI = "/wxa/getwxacode"
)

// Response 请求微信返回基础数据
type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// WeApp 小程序实例
type WeApp struct {
	AppID  string
	Secret string
	Token  string
	AesKey string
}

var app WeApp

// Init 初始化小程序
func Init(appID, secret, token, aesKey string) {
	app.AppID = appID
	app.Secret = secret
	app.Token = token
	app.AesKey = aesKey
}

// code2url 拼接 获取 session_key 的 URL
func (app WeApp) code2url(code string) (string, error) {

	url, err := url.Parse(baseURL + codeAPI)
	if err != nil {
		return "", err
	}

	query := url.Query()

	query.Set("appid", app.AppID)
	query.Set("secret", app.Secret)
	query.Set("js_code", code)
	query.Set("grant_type", "authorization_code")

	url.RawQuery = query.Encode()

	return url.String(), nil
}

type loginForm struct {
	Response
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// Login 小程序登陆
// 返回 微信端 openid 和 session_key
func Login(code string) (string, string, error) {
	if code == "" {
		return "", "", errors.New("code can not be null")
	}

	url, err := app.code2url(code)
	if err != nil {
		return "", "", err
	}

	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", "", errors.New("error when request wechat server")
	}

	var data loginForm

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", "", err
	}

	if data.Errcode != 0 {
		return "", "", errors.New(data.Errmsg)
	}

	return data.Openid, data.SessionKey, nil
}

// AppCode 获取小程序码
// path 识别二维码后进入小程序的页面链接
// width 图片宽度
// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// token access_token
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
func AppCode(path string, width int, autoColor bool, lineColor, token string) (res *http.Response, err error) {

	api, err := url.Parse(baseURL + appCodeAPI)
	if err != nil {
		return res, err
	}

	query := api.Query()
	query.Set("access_token", token)
	api.RawQuery = query.Encode()

	body := fmt.Sprintf(`{"path":"%s","width": %v,"auto_color": %v,"line_color": %s}`, path, width, autoColor, lineColor)
	res, err = http.Post(api.String(), "application/json", strings.NewReader(body))
	if err != nil {
		return res, err
	}

	switch res.Header.Get("Content-Type") {
	case "application/json": // 返回错误信息
		var data Response
		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			return res, err
		}
		return res, errors.New(data.Errmsg)
	case "image/jpeg": // 返回文件
		return res, nil
	}

	return res, errors.New("unknown error when fetch app code")
}
