package weapp

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp/v3/request"
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
	csMsgTypeLink   csMsgType = "link"            // 图文链接消息类型
	csMsgTypeImage  csMsgType = "image"           // 图片消息类型
	csMsgTypeMPCard csMsgType = "miniprogrampage" // 小程序卡片消息类型
)

// csMessage 消息体
type csMessage struct {
	Receiver string       `json:"touser"`  // user openID
	Type     csMsgType    `json:"msgtype"` // text | image | link | miniprogrampage
	Text     *CSMsgText   `json:"text,omitempty"`
	Image    *CSMsgImage  `json:"image,omitempty"`
	Link     *CSMsgLink   `json:"link,omitempty"`
	MPCard   *CSMsgMPCard `json:"miniprogrampage,omitempty"`
}

// CSMsgText 接收的文本消息
type CSMsgText struct {
	Content string `json:"content"`
}

// SendTo 发送文本消息
//
// openID 用户openID
func (cli *Client) SendTextMsg(openID string, msg *CSMsgText) (*request.CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     csMsgTypeText,
		Text:     msg,
	}

	return cli.sendMessage(params)
}

// CSMsgImage 客服图片消息
type CSMsgImage struct {
	MediaID string `json:"media_id"` // 发送的图片的媒体ID，通过 新增素材接口 上传图片文件获得。
}

// SendTo 发送图片消息
//
// openID 用户openID
func (cli *Client) SendImageMsg(openID string, msg *CSMsgImage) (*request.CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     csMsgTypeImage,
		Image:    msg,
	}

	return cli.sendMessage(params)
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
func (cli *Client) SendLinkMsg(openID string, msg *CSMsgLink) (*request.CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     csMsgTypeLink,
		Link:     msg,
	}

	return cli.sendMessage(params)
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
func (cli *Client) SendCardMsg(openID string, msg *CSMsgMPCard) (*request.CommonError, error) {

	params := csMessage{
		Receiver: openID,
		Type:     csMsgTypeMPCard,
		MPCard:   msg,
	}

	return cli.sendMessage(params)
}

// send 发送消息
func (cli *Client) sendMessage(params interface{}) (*request.CommonError, error) {
	api := baseURL + apiSendMessage

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.doSendMessage(api, token, params)
}

func (cli *Client) doSendMessage(api, token string, params interface{}) (*request.CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SetTypingCommand 下发客服当前输入状态命令
type SetTypingCommand = string

// 所有下发客服当前输入状态命令
const (
	SetTypingCommandTyping       SetTypingCommand = "Typing"       // 对用户下发"正在输入"状态
	SetTypingCommandCancelTyping SetTypingCommand = "CancelTyping" // 取消对用户的"正在输入"状态
)

// SetTyping 下发客服当前输入状态给用户。
//
// token 接口调用凭证
// openID 用户的 OpenID
// cmd 命令
func (cli *Client) SetTyping(openID string, cmd SetTypingCommand) (*request.CommonError, error) {
	api := baseURL + apiSetTyping

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.setTyping(token, openID, cmd, api)
}

func (cli *Client) setTyping(token, openID string, cmd SetTypingCommand, api string) (*request.CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"touser":  openID,
		"command": cmd,
	}

	res := new(request.CommonError)
	if err := cli.request.Post(url, params, res); err != nil {
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
	request.CommonError
	Type      string `json:"type"`       // 文件类型
	MediaID   string `json:"media_id"`   // 媒体文件上传后，获取标识，3天内有效。
	CreatedAt uint   `json:"created_at"` // 媒体文件上传时间戳
}

// UploadTempMedia 把媒体文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
//
// token 接口调用凭证
// mediaType 文件类型
// medianame 媒体文件名
func (cli *Client) UploadTempMedia(mediaType TempMediaType, medianame string) (*UploadTempMediaResponse, error) {
	api := baseURL + apiUploadTemplateMedia

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.uploadTempMedia(token, mediaType, medianame, api)
}

func (cli *Client) uploadTempMedia(token string, mediaType TempMediaType, medianame, api string) (*UploadTempMediaResponse, error) {
	queries := requestQueries{
		"type":         mediaType,
		"access_token": token,
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(UploadTempMediaResponse)
	if err := cli.request.FormPostWithFile(url, "media", medianame, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetTempMedia 获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
//
// token 接口调用凭证
// mediaID 媒体文件 ID
func (cli *Client) GetTempMedia(mediaID string) (*http.Response, *request.CommonError, error) {
	api := baseURL + apiGetTemplateMedia

	token, err := cli.AccessToken()
	if err != nil {
		return nil, nil, err
	}

	return cli.getTempMedia(token, mediaID, api)
}

func (cli *Client) getTempMedia(token, mediaID, api string) (*http.Response, *request.CommonError, error) {
	queries := requestQueries{
		"access_token": token,
		"media_id":     mediaID,
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, nil, err
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	response := new(request.CommonError)
	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		if err := json.NewDecoder(res.Body).Decode(response); err != nil {
			res.Body.Close()
			return nil, nil, err
		}
		return res, response, nil

	case strings.HasPrefix(header, "image"):
		return res, response, nil

	default:
		res.Body.Close()
		return nil, nil, errors.New("invalid response header: " + header)
	}
}
