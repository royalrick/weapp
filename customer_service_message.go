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

// csMsgType 消息类型
type csMsgType string

// 所有消息类型
const (
	csMsgTypeText   csMsgType = "text"            // 文本消息类型
	csMsgTypeLink             = "link"            // 图文链接消息类型
	csMsgTypeImage            = "image"           // 图片消息类型
	csMsgTypeMPCard           = "miniprogrampage" // 小程序卡片消息类型
)

// csMessage 消息体
type csMessage struct {
	Receiver string      `json:"touser"`  // user openID
	Type     csMsgType   `json:"msgtype"` // text | image | link | miniprogrampage
	Text     CSMsgText   `json:"text,omitempty"`
	Image    CSMsgImage  `json:"image,omitempty"`
	Link     CSMsgLink   `json:"link,omitempty"`
	MPCard   CSMsgMPCard `json:"miniprogrampage,omitempty"`
}

// CSMsgText 接收的文本消息
type CSMsgText struct {
	Content string `json:"content"`
}

// SendTo 发送文本消息
//
// openID 用户openID
// token 微信 access_token
func (msg CSMsgText) SendTo(openID, token string) (*CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     csMsgTypeText,
		Text:     msg,
	}

	return sendMessage(token, params)
}

// CSMsgImage 客服图片消息
type CSMsgImage struct {
	MediaID string `json:"media_id"` // 发送的图片的媒体ID，通过 新增素材接口 上传图片文件获得。
}

// SendTo 发送图片消息
//
// openID 用户openID
// token 微信 access_token
func (msg CSMsgImage) SendTo(openID, token string) (*CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     csMsgTypeImage,
		Image:    msg,
	}

	return sendMessage(token, params)
}

// CSMsgLink 图文链接消息
type CSMsgLink struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}

// SendTo 发送图文链接消息
//
// openID 用户openID
// token 微信 access_token
func (msg CSMsgLink) SendTo(openID, token string) (*CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     csMsgTypeLink,
		Link:     msg,
	}

	return sendMessage(token, params)
}

// CSMsgMPCard 接收的卡片消息
type CSMsgMPCard struct {
	Title        string `json:"title"`          // 标题
	PagePath     string `json:"pagepath"`       // 小程序页面路径
	ThumbMediaID string `json:"thumb_media_id"` // 小程序消息卡片的封面， image 类型的 media_id，通过 新增素材接口 上传图片文件获得，建议大小为 520*416
}

// SendTo 发送卡片消息
//
// openID 用户openID
// token 微信 access_token
func (msg CSMsgMPCard) SendTo(openID, token string) (*CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     "miniprogrampage",
		MPCard:   msg,
	}

	return sendMessage(token, params)
}

// send 发送消息
//
// token 微信 access_token
func sendMessage(token string, params interface{}) (*CommonError, error) {
	api := baseURL + apiSendMessage
	return doSendMessage(token, params, api)
}

func doSendMessage(token string, params interface{}, api string) (*CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postJSON(url, params, res); err != nil {
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
	api := baseURL + apiSetTyping
	return setTyping(token, openID, cmd, api)
}

func setTyping(token, openID string, cmd SetTypingCommand, api string) (*CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"touser":  openID,
		"command": cmd,
	}

	res := new(CommonError)
	if err := postJSON(url, params, res); err != nil {
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
// medianame 媒体文件名
func UploadTempMedia(token string, mediaType TempMediaType, medianame string) (*UploadTempMediaResponse, error) {
	api := baseURL + apiUploadTemplateMedia
	return uploadTempMedia(token, mediaType, medianame, api)
}

func uploadTempMedia(token string, mediaType TempMediaType, medianame, api string) (*UploadTempMediaResponse, error) {
	queries := requestQueries{
		"type":         mediaType,
		"access_token": token,
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(UploadTempMediaResponse)
	if err := postFormByFile(url, "media", medianame, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetTempMedia 获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
//
// token 接口调用凭证
// mediaID 媒体文件 ID
func GetTempMedia(token, mediaID string) (*http.Response, *CommonError, error) {
	api := baseURL + apiGetTemplateMedia
	return getTempMedia(token, mediaID, api)
}

func getTempMedia(token, mediaID, api string) (*http.Response, *CommonError, error) {
	queries := requestQueries{
		"access_token": token,
		"media_id":     mediaID,
	}

	url, err := encodeURL(api, queries)
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

	case strings.HasPrefix(header, "image"): // 返回文件 TODO: 应该确认一下
		return res, response, nil

	default:
		res.Body.Close()
		return nil, nil, errors.New("invalid response header: " + header)
	}
}
