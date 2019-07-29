package weapp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// 检测地址
const (
	IMGSecCheckURL = "/wxa/img_sec_check"
	MSGSecCheckURL = "/wxa/msg_sec_check"
)

// IMGSecCheckFromNet 网络图片检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html
//
// @url 要检测的图片网络路径
// @token 接口调用凭证(access_token)
func IMGSecCheckFromNet(url, token string) (*commonError, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()
	str := strings.Split(url, "/")
	fileWriter, err := writer.CreateFormFile("media", str[len(str)-1])
	if err != nil {
		return nil, err
	}
	_, err = fileWriter.Write(bts)
	if err != nil {
		return nil, err
	}
	contentType := writer.FormDataContentType()

	return imgSecCheck(body, contentType, token)
}

// IMGSecCheck 本地图片检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html
//
// @filename 要检测的图片本地路径
// @token 接口调用凭证(access_token)
func IMGSecCheck(filename, token string) (*commonError, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{} // TODO: 优化
	writer := multipart.NewWriter(body)
	defer writer.Close()
	fileWriter, err := writer.CreateFormFile("media", filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}
	contentType := writer.FormDataContentType()

	return imgSecCheck(body, contentType, token)
}

func imgSecCheck(body io.Reader, contentType, token string) (*commonError, error) {

	api, err := tokenAPI(baseURL+IMGSecCheckURL, token)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(api, contentType, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := new(commonError)
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}

// MSGSecCheck 文本检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/msgSecCheck.html
//
// @content 要检测的文本内容，长度不超过 500KB，编码格式为utf-8
// @token 接口调用凭证(access_token)
func MSGSecCheck(content, token string) (*commonError, error) {
	api, err := tokenAPI(baseURL+MSGSecCheckURL, token)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"content": content,
	}

	res := new(commonError)
	if err = postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
