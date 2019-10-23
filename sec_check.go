package weapp

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
func IMGSecCheck(token, filename string) (*CommonError, error) {
	api := baseURL + apiIMGSecCheck
	return imgSecCheck(api, filename, token)
}

func imgSecCheck(api, token, filename string) (*CommonError, error) {

	url, err := tokenAPI(api, token)
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
func MSGSecCheck(token, content string) (*CommonError, error) {
	api := baseURL + apiMSGSecCheck
	return msgSecCheck(api, token, content)
}

func msgSecCheck(api, token, content string) (*CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"content": content,
	}

	res := new(CommonError)
	if err = postJSON(url, params, res); err != nil {
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
	TraceID string `json:"trace_id"`
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容。
//
// mediaURL 要检测的多媒体url
// mediaType 接口调用凭证(access_token)
func MediaCheckAsync(token, mediaURL string, mediaType MediaType) (*CheckMediaResponse, error) {
	api := baseURL + apiMediaCheckAsync
	return mediaCheckAsync(api, token, mediaURL, mediaType)
}

func mediaCheckAsync(api, token, mediaURL string, mediaType MediaType) (*CheckMediaResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"media_url":  mediaURL,
		"media_type": mediaType,
	}

	res := new(CheckMediaResponse)
	if err = postJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
