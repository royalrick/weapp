package weapp

import (
	"net/http"
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
// imgURL 要检测的图片网络路径
// token 接口调用凭证(access_token)
func IMGSecCheckByURL(imgURL, token string) (*CommonError, error) {
	api := baseURL + apiCheckImg
	return imgSecCheckByURL(imgURL, token, api)
}

func imgSecCheckByURL(imgURL, token, api string) (*CommonError, error) {
	resp, err := http.Get(imgURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postForm(url, "media", "filename", resp.Body, res); err != nil {
		return nil, err
	}

	return res, nil
}

// IMGSecCheck 本地图片检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html
//
// filename 要检测的图片本地路径
// token 接口调用凭证(access_token)
func IMGSecCheck(filename, token string) (*CommonError, error) {
	api := baseURL + apiCheckImg
	return imgSecCheck(filename, token, api)
}

func imgSecCheck(filename, token, api string) (*CommonError, error) {

	url, err := tokenAPI(baseURL+apiCheckImg, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postFormByFile(url, "media", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}

// MSGSecCheck 文本检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/msgSecCheck.html
//
// content 要检测的文本内容，长度不超过 500KB，编码格式为utf-8
// token 接口调用凭证(access_token)
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
// mediaURL 要检测的多媒体url
// mediaType 接口调用凭证(access_token)
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
