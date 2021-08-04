package weapp

import "github.com/medivhzhan/weapp/v3/request"

// 检测地址
const (
	apiIMGSecCheck     = "/wxa/img_sec_check"
	apiMSGSecCheck     = "/wxa/msg_sec_check"
	apiMediaCheckAsync = "/wxa/media_check_async"
)

// IMGSecCheck 本地图片检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html
//
// filename 要检测的图片本地路径
// token 接口调用凭证(access_token)
func (cli *Client) IMGSecCheck(filename string) (*request.CommonError, error) {
	api := baseURL + apiIMGSecCheck

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.imgSecCheck(api, token, filename)
}

func (cli *Client) imgSecCheck(api, token, filename string) (*request.CommonError, error) {

	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.FormPostWithFile(url, "media", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}

// MSGSecCheck 文本检测
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api/msgSecCheck.html
//
// content 要检测的文本内容，长度不超过 500KB，编码格式为utf-8
// token 接口调用凭证(access_token)
func (cli *Client) MSGSecCheck(content string) (*request.CommonError, error) {
	api := baseURL + apiMSGSecCheck

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.msgSecCheck(api, token, content)
}

func (cli *Client) msgSecCheck(api, token, content string) (*request.CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"content": content,
	}

	res := new(request.CommonError)
	if err = cli.request.Post(url, params, res); err != nil {
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
	request.CommonError
	TraceID string `json:"trace_id"`
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容。
//
// mediaURL 要检测的多媒体url
// mediaType 接口调用凭证(access_token)
func (cli *Client) MediaCheckAsync(mediaURL string, mediaType MediaType) (*CheckMediaResponse, error) {
	api := baseURL + apiMediaCheckAsync

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.mediaCheckAsync(api, token, mediaURL, mediaType)
}

func (cli *Client) mediaCheckAsync(api, token, mediaURL string, mediaType MediaType) (*CheckMediaResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"media_url":  mediaURL,
		"media_type": mediaType,
	}

	res := new(CheckMediaResponse)
	if err = cli.request.Post(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
