package security

import "github.com/medivhzhan/weapp/v3/request"

const apiMediaCheckAsync = "/wxa/media_check_async"

type MediaCheckAsyncRequest struct {
	// 必填 要检测的图片或音频的url，支持图片格式包括jpg, jepg, png, bmp, gif（取首帧），支持的音频格式包括mp3, aac, ac3, wma, flac, vorbis, opus, wav
	MediaUrl string `json:"media_url"`
	// 必填 1:音频;2:图片
	MediaType uint8 `json:"media_type"`
	// 必填 接口版本号，2.0版本为固定值2
	Version uint8 `json:"version"`
	// 必填 用户的openid（用户需在近两小时访问过小程序）
	Openid string `json:"openid"`
	// 必填 场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
	Scene uint8 `json:"scene"`
}

type MediaCheckAsyncResponse struct {
	request.CommonError
	// 唯一请求标识，标记单次请求，用于匹配异步推送结果
	TraceId string `json:"trace_id"`
}

// 异步校验图片/音频是否含有违法违规内容。
//
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html
func (cli *Security) MediaCheckAsync(req *MediaCheckAsyncRequest) (*MediaCheckAsyncResponse, error) {
	url, err := cli.combineURI(apiMediaCheckAsync, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(MediaCheckAsyncResponse)
	if err := cli.request.Post(url, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
