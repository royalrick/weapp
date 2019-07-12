// Package weapp 接收并处理微信通知
package weapp

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// MsgType 消息类型
type MsgType = string

// 所有消息类型
const (
	MsgText  MsgType = "text"            // 文本消息类型
	MsgImg           = "image"           // 图片消息类型
	MsgLink          = "link"            // 图文链接消息类型
	MsgVideo         = "video"           // 视频消息类型
	MsgCard          = "miniprogrampage" // 小程序卡片消息类型
	MsgEvent         = "event"           // 事件类型
)

// EventType 事件类型
type EventType = string

// 所有事件类型
const (
	EventUserEntry EventType = "user_enter_tempsession" // 用户进入临时会话状态
	EventGetQuota            = "get_quota"              // 查询商户余额
)

// EncryptedMsgResponse 接收的的加密消息格式
type EncryptedMsgResponse struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	ToUserName string   `json:"ToUserName" xml:"ToUserName"` // 接收者 为公众号的原始ID
	Encrypt    string   `json:"Encrypt" xml:"Encrypt"`       // 加密消息
}

// EncryptedMsgRequest 发送的加密消息格式
type EncryptedMsgRequest struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `json:"Encrypt" xml:"Encrypt"`                               // 加密消息
	TimeStamp    string   `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`       // 时间戳
	Nonce        string   `json:"Nonce,omitempty" xml:"Nonce,omitempty"`               // 随机数
	MsgSignature string   `json:"MsgSignature,omitempty" xml:"MsgSignature,omitempty"` // 签名
}

// Mixture 混合消息体
type Mixture struct {
	XMLName      xml.Name  `xml:"xml" json:"-"`
	ID           int       `json:"MsgId" xml:"MsgId"`                     // 消息 ID
	Type         MsgType   `json:"MsgType" xml:"MsgType"`                 // 消息类型
	Event        EventType `json:"event,omitempty" xml:"event,omitempty"` // 事件类型
	FromUserName string    `json:"FromUserName" xml:"FromUserName"`       // 发送者的 openID
	ToUserName   string    `json:"ToUserName" xml:"ToUserName"`           // 小程序的原始ID
	CreateTime   int64     `json:"CreateTime" xml:"CreateTime"`           // 消息创建时间(整型）

	Text
	Card
	Image
}

// Text 接收的文本消息
type Text struct {
	Content string `json:"Content,omitempty" xml:"Content,omitempty"`
}

// Image 接收的图片消息
type Image struct {
	PicURL  string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	MediaID string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

// Card 接收的卡片消息
type Card struct {
	Title        string `json:"Title,omitempty" xml:"Title,omitempty"`               // 标题
	AppID        string `json:"AppId,omitempty" xml:"AppId,omitempty"`               // 小程序 appid
	PagePath     string `json:"PagePath,omitempty" xml:"PagePath,omitempty"`         // 小程序页面路径
	ThumbURL     string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`         // 封面图片的临时cdn链接
	ThumbMediaID string `json:"ThumbMediaId,omitempty" xml:"ThumbMediaId,omitempty"` // 封面图片的临时素材id
}

// Server 微信通知服务处理器
type Server struct {
	appID          string // 小程序 ID
	mchID          string // 商户号
	apiKey         string // 商户签名密钥
	token          string // 微信服务器验证令牌
	aesKey         string // 消息加密密钥 长度固定为43个字符，从a-z,A-Z,0-9共62个字符中选取。
	ValidateServer bool   // 是否验证请求来自微信服务器

	TextMessageHandler  func(Text) bool     // 文本消息处理器
	CardMessageHandler  func(Card) bool     // 卡片消息处理器
	ImageMessageHandler func(Image) bool    // 图片消息处理器
	EventHandler        func(*Mixture) bool // 事件处理器
}

type dataType = string

const (
	dataTypeJSON dataType = "application/json"
	dataTypeXML           = "application/xml"
)

// NewServer 返回经过初始化的Server
func NewServer(appID, token, aesKey, mchID, apiKey string) *Server {
	return &Server{appID: appID, mchID: mchID, apiKey: apiKey, token: token, aesKey: aesKey}
}

func getDataType(req *http.Request) dataType {
	content := req.Header.Get("Content-Type")

	switch {
	case strings.Contains(content, dataTypeJSON):
		return dataTypeJSON
	case strings.Contains(content, dataTypeXML):
		return dataTypeXML
	default:
		return content
	}
}

func (srv *Server) getAESKey() ([]byte, error) {

	if len(srv.aesKey) != 43 {
		return nil, errors.New("invalid aes key")
	}

	str := srv.aesKey + "="
	return base64.StdEncoding.DecodeString(str)
}

func unmarshal(data []byte, tp dataType, v interface{}) error {
	switch tp {
	case dataTypeJSON:
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
	case dataTypeXML:
		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}
	default:
		return errors.New("invalid content type: " + tp)
	}

	return nil
}

// ServeHTTP 接收并处理微信通知服务
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		raw, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		mix := new(Mixture)
		tp := getDataType(r)
		if isEncrypted(r) { // 处理加密消息
			res := new(EncryptedMsgResponse)
			if err := unmarshal(raw, tp, res); err != nil {
				return err
			}

			nonce := GetQuery(r, "nonce")
			signature := GetQuery(r, "msg_signature")
			timestamp := GetQuery(r, "timestamp")

			// 检验消息的真实性
			if !validateSignature(signature, srv.token, timestamp, nonce, res.Encrypt) {
				return errors.New("invalid signature")
			}
			err := srv.decryptMsg(res.Encrypt, tp, mix)
			if err != nil {
				return err
			}
		} else {
			if err := unmarshal(raw, tp, mix); err != nil {
				return err
			}
		}

		ok := false // 是否已经收到消息
		switch mix.Type {

		case MsgText: // 文本消息
			if srv.TextMessageHandler != nil {
				msg := mix.Text
				ok = srv.TextMessageHandler(msg)
			}

		case MsgImg: // 图片消息
			if srv.ImageMessageHandler != nil {
				msg := mix.Image
				ok = srv.ImageMessageHandler(msg)
			}

		case MsgCard: // 卡片消息
			if srv.CardMessageHandler != nil {
				msg := mix.Card
				ok = srv.CardMessageHandler(msg)
			}

		case MsgLink: // 图文链接消息
			// TODO: ...

		case MsgVideo: // 视频消息
			// TODO: ...

		case MsgEvent: // 事件
			if srv.EventHandler != nil {
				ok = srv.EventHandler(mix)
			}

		default:
			return errors.New("invalid message type: " + mix.Type)
		}

		if ok {
			_, err := io.WriteString(w, "SUCCESS")
			if err != nil {
				return err
			}
		}

		return nil
	case "GET":
		echostr := GetQuery(r, "echostr")
		if srv.ValidateServer {

			// 请求来自微信验证成功后原样返回 echostr 参数内容
			if srv.validateServer(r) {
				_, err := io.WriteString(w, echostr)
				if err != nil {
					return err
				}

				return nil
			}

			return errors.New("request server is invalid")
		}

		_, err := io.WriteString(w, echostr)
		if err != nil {
			return err
		}

		return nil
	default:
		return errors.New("invalid request method: " + r.Method)
	}
}

// 判断消息是否加密
func isEncrypted(req *http.Request) bool {
	return GetQuery(req, "encrypt_type") == "aes"
}

// 验证消息的确来自微信服务器
// 1.将token、timestamp、nonce三个参数进行字典序排序
// 2.将三个参数字符串拼接成一个字符串进行sha1加密
// 3.开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
func (srv *Server) validateServer(req *http.Request) bool {
	nonce := GetQuery(req, "nonce")
	signature := GetQuery(req, "signature")
	timestamp := GetQuery(req, "timestamp")

	return validateSignature(signature, nonce, timestamp, srv.token)
}

// 将公众号回复用户的消息加密打包
func (srv *Server) encryptMsg(message, nonce string, timestamp int) (*EncryptedMsgRequest, error) {

	key, err := srv.getAESKey()
	if err != nil {
		return nil, err
	}

	//获得16位随机字符串，填充到明文之前
	random := RandomString(16)
	text := random + string(len(message)) + message + srv.appID
	plaintext := pkcs7encode([]byte(text))

	cipher, err := cbcEncrypt(key, plaintext, key)
	if err != nil {
		return nil, err
	}

	encrypt := base64.StdEncoding.EncodeToString(cipher)
	timestr := strconv.Itoa(timestamp)

	//生成安全签名
	signature := createSignature(srv.token, timestr, nonce, encrypt)

	request := EncryptedMsgRequest{
		Nonce:        nonce,
		Encrypt:      encrypt,
		TimeStamp:    timestr,
		MsgSignature: signature,
	}

	return &request, nil
}

// 检验消息的真实性，并且获取解密后的明文.
func (srv *Server) decryptMsg(encrypted string, tp dataType, msg *Mixture) error {

	key, err := srv.getAESKey()
	if err != nil {
		return err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return err
	}

	data, err := cbcDecrypt(key, ciphertext, key)
	if err != nil {
		return err
	}

	// TODO: old
	// 	if err := unmarshal(data[:], msg, srv.dataType()); err != nil {
	// 		return err
	// 	}
	length := binary.BigEndian.Uint32(data[16:20])
	if err := unmarshal(data[20:20+length], tp, msg); err != nil {
		return err
	}

	return nil
}
