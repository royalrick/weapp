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
	apiCheckImg   = "/wxa/img_sec_check"
	apiCheckMsg   = "/wxa/msg_sec_check"
	apiCheckMedia = "/wxa/media_check_async"
)

// IMGSecCheckByURL 网络图片检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html
//
// @url 要检测的图片网络路径
// @token 接口调用凭证(access_token)
func IMGSecCheckByURL(url, token string) (*CommonError, error) {
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
func IMGSecCheck(filename, token string) (*CommonError, error) {
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

func imgSecCheck(body io.Reader, contentType, token string) (*CommonError, error) {

	api, err := tokenAPI(baseURL+apiCheckImg, token)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(api, contentType, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := new(CommonError)
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
func MSGSecCheck(content, token string) (*CommonError, error) {
	api, err := tokenAPI(baseURL+apiCheckMsg, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"content": content,
	}

	res := new(CommonError)
	if err = postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// MediaType 检测内容类型
type MediaType = uint8

// 所有检测内容类型
const (
	_              MediaType = iota
	MediaTypeAudio           // 音频
	MediaTypeImage           // 图片
)

// CheckMediaResponse 异步校验图片/音频返回数据
type CheckMediaResponse struct {
	CommonError
	TranceID string `json:"trace_id"`
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容。
//
// @mediaURL 要检测的多媒体url
// @mediaType 接口调用凭证(access_token)
func MediaCheckAsync(mediaURL string, mediaType MediaType, token string) (*CheckMediaResponse, error) {
	api, err := tokenAPI(baseURL+apiCheckMedia, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"media_url":  mediaURL,
		"media_type": mediaType,
	}

	res := new(CheckMediaResponse)
	if err = postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
