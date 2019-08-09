package weapp

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	apiSendMessage         = "/cgi-bin/message/custom/send"
	apiSetTyping           = "/cgi-bin/message/custom/typing"
	apiUploadTemplateMedia = "/cgi-bin/media/upload"
	apiGetTemplateMedia    = "/cgi-bin/media/get"
)

// 消息体
type message struct {
	Receiver string  `json:"touser"`  // user openID
	Type     MsgType `json:"msgtype"` // text | image | link | miniprogrampage
	Text     Text    `json:"text,omitempty"`
	Image    Image   `json:"image,omitempty"`
	Link     Link    `json:"link,omitempty"`
	Card     Card    `json:"miniprogrampage,omitempty"`
}

// SendTo 发送文本消息
//
// openID 用户openID
// token 微信 access_token
func (msg Text) SendTo(openID, token string) (*CommonError, error) {

	params := message{
		Receiver: openID,
		Type:     "text",
		Text:     msg,
	}

	return sendMessage(token, params)
}

// SendTo 发送图片消息
//
// openID 用户openID
// token 微信 access_token
func (msg Image) SendTo(openID, token string) (*CommonError, error) {

	params := message{
		Receiver: openID,
		Type:     "image",
		Image:    msg,
	}

	return sendMessage(token, params)
}

// SendTo 发送图文链接消息
//
// openID 用户openID
// token 微信 access_token
func (msg Link) SendTo(openID, token string) (*CommonError, error) {

	params := message{
		Receiver: openID,
		Type:     "link",
		Link:     msg,
	}

	return sendMessage(token, params)
}

// SendTo 发送卡片消息
//
// openID 用户openID
// token 微信 access_token
func (msg Card) SendTo(openID, token string) (*CommonError, error) {

	params := message{
		Receiver: openID,
		Type:     "miniprogrampage",
		Card:     msg,
	}

	return sendMessage(token, params)
}

// send 发送消息
//
// token 微信 access_token
func sendMessage(token string, params interface{}) (*CommonError, error) {
	api, err := tokenAPI(baseURL+apiSendMessage, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SetTypingCommand 下发客服当前输入状态命令
type SetTypingCommand = string

// 所有下发客服当前输入状态命令
const (
	SetTypingCommandTyping       SetTypingCommand = "Typing"       // 对用户下发"正在输入"状态
	SetTypingCommandCancelTyping                  = "CancelTyping" // 取消对用户的"正在输入"状态
)

// SetTyping 下发客服当前输入状态给用户。
//
// token 接口调用凭证
// openID 用户的 OpenID
// cmd 命令
func SetTyping(token, openID string, cmd SetTypingCommand) (*CommonError, error) {
	api, err := tokenAPI(baseURL+apiSetTyping, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"touser":  openID,
		"command": cmd,
	}

	res := new(CommonError)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// TempMediaType 文件类型
type TempMediaType = string

// 所有文件类型
const (
	TempMediaTypeImage TempMediaType = "image" // 图片
)

// UploadTempMediaResponse 上传媒体文件返回
type UploadTempMediaResponse struct {
	CommonError
	Type      string `json:"type"`       // 文件类型
	MediaID   string `json:"media_id"`   // 媒体文件上传后，获取标识，3天内有效。
	CreatedAt uint   `json:"created_at"` // 媒体文件上传时间戳
}

// UploadTempMedia 把媒体文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
//
// token 接口调用凭证
// mediaType 文件类型
// media form-data 中媒体文件标识，有filename、filelength、content-type等信息
func UploadTempMedia(token string, mediaType TempMediaType, media string) (*UploadTempMediaResponse, error) {
	api, err := tokenAPI(baseURL+apiUploadTemplateMedia, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"type":  mediaType,
		"media": media,
	}

	res := new(UploadTempMediaResponse)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetTempMedia 获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
//
// token 接口调用凭证
// mediaID 媒体文件 ID
func GetTempMedia(token, mediaID string) (*http.Response, *CommonError, error) {
	params := requestQueries{
		"access_token": token,
		"media_id":     mediaID,
	}

	url, err := encodeURL(baseURL+apiGetTemplateMedia, params)
	if err != nil {
		return nil, nil, err
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	response := new(CommonError)
	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		if err := json.NewDecoder(res.Body).Decode(response); err != nil {
			res.Body.Close()
			return nil, nil, err
		}
		return res, response, nil

	case header == "image/jpeg": // 返回文件 TODO: 应该确认一下
		return res, response, nil

	default:
		res.Body.Close()
		return nil, nil, errors.New("invalid response header: " + header)
	}
}
