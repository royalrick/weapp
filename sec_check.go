package weapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/medivhzhan/weapp/util"
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
func IMGSecCheckFromNet(url, token string) (res Response, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bts, err := ioutil.ReadAll(resp.Body)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	str := strings.Split(url, "/")
	fmt.Println(str)
	fmt.Println(str[len(str)-1])
	fileWriter, err := writer.CreateFormFile("media", str[len(str)-1])
	if err != nil {
		return
	}
	_, err = fileWriter.Write(bts)
	if err != nil {
		return
	}
	contentType := writer.FormDataContentType()
	writer.Close()

	return imgSecCheck(body, contentType, token)
}

// IMGSecCheck 本地图片检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html
//
// @filename 要检测的图片本地路径
// @token 接口调用凭证(access_token)
func IMGSecCheck(filename, token string) (res Response, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileWriter, err := writer.CreateFormFile("media", filename)
	if err != nil {
		return
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return
	}
	contentType := writer.FormDataContentType()
	writer.Close()

	return imgSecCheck(body, contentType, token)
}

func imgSecCheck(body io.Reader, contentType, token string) (res Response, err error) {

	api, err := util.TokenAPI(BaseURL+IMGSecCheckURL, token)
	if err != nil {
		return
	}

	resp, err := http.Post(api, contentType, body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	return
}

// MSGSecCheck 文本检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/msgSecCheck.html
//
// @content 要检测的文本内容，长度不超过 500KB，编码格式为utf-8
// @token 接口调用凭证(access_token)
func MSGSecCheck(content, token string) (res Response, err error) {
	api, err := util.TokenAPI(BaseURL+MSGSecCheckURL, token)
	if err != nil {
		return
	}

	body := fmt.Sprintf(`{"content": "%s"}`, content)

	resp, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	return
}
