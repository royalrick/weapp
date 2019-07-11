// Package weapp 接收并处理微信通知
package weapp

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
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

// Server 微信服务接收器
// TODO: 删除不必要的字段
type Server struct {
	appID          string // 小程序 ID
	mchID          string // 商户号
	apiKey         string // 商户签名密钥
	token          string // 微信服务器验证令牌
	EncodingAESKey string // 消息加密密钥
	ValidateServer bool   // 是否验证请求来自微信服务器

	TextMessageHandler  func(Text) bool     // 文本消息处理器
	CardMessageHandler  func(Card) bool     // 卡片消息处理器
	ImageMessageHandler func(Image) bool    // 图片消息处理器
	EventHandler        func(*Mixture) bool // 事件处理器
}

func NewServer(AppID, Token, AESKeyBase64, MerchantID, MerchantAPIKey string) *Server {
	return &Server{appID: AppID, mchID: MerchantID, apiKey: MerchantAPIKey, token: Token, EncodingAESKey: AESKeyBase64}
}

func (srv *Server) getAESKey() ([]byte, error) {
	str := srv.EncodingAESKey + "="
	return base64.StdEncoding.DecodeString(str)
}

type dataType = string

const (
	dataJSON dataType = "application/json"
	dataXML           = "application/xml"
)

func unmarshal(data []byte, v interface{}, ct dataType) error {
	switch ct {
	case dataJSON:
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
	case dataXML:
		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}
	default:
		return errors.New("invalid content type: " + ct)
	}

	return nil
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		mix := new(Mixture)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		if isEncrypted(r) { // 处理加密消息
			encrypted := new(EncryptedMsgResponse)

			if err := unmarshal(body, encrypted, r.Header.Get("Content-Type")); err != nil {
				panic(err)
			}

			nonce := GetQuery(r, "nonce")
			signature := GetQuery(r, "msg_signature")
			timestamp := GetQuery(r, "timestamp")

			// 检验消息的真实性
			if !validateSignature(signature, srv.token, timestamp, nonce, encrypted.Encrypt) {
				panic(errors.New("invalid signature"))
			}
			err := srv.decryptMsg(encrypted.Encrypt, mix, r.Header.Get("Content-Type"))
			if err != nil {
				panic(err)
			}
		} else {
			if err := unmarshal(body, mix, r.Header.Get("Content-Type")); err != nil {
				panic(err)
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

		case MsgVideo: // 图文链接消息
			// TODO: ...

		case MsgEvent: // 事件
			if srv.EventHandler != nil {
				ok = srv.EventHandler(mix)
			}

		default:
			panic(errors.New("invalid message type: " + mix.Type))
		}

		if ok {
			_, err := io.WriteString(w, "SUCCESS")
			if err != nil {
				panic(err)
			}
		}

		return
	case "GET":
		echostr := GetQuery(r, "echostr")
		if srv.ValidateServer {

			// 请求来自微信验证成功后原样返回 echostr 参数内容
			if srv.validateServer(r) {
				_, err := io.WriteString(w, echostr)
				if err != nil {
					panic(err)
				}
				return
			}

			panic(errors.New("request server is invalid"))
		}

		_, err := io.WriteString(w, echostr)
		if err != nil {
			panic(err)
		}
	default:
		panic(errors.New("invalid request method: " + r.Method))
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

func validateSignature(signature string, parts ...string) bool {
	sort.Strings(parts)
	raw := sha1.Sum([]byte(strings.Join(parts, "")))
	return signature == hex.EncodeToString(raw[:])
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
	slice := sort.StringSlice{srv.token, timestr, nonce, encrypt}

	slice.Sort()
	signature := sha1.Sum([]byte(strings.Join(slice, "")))

	return &EncryptedMsgRequest{
		Nonce:        nonce,
		Encrypt:      encrypt,
		TimeStamp:    timestr,
		MsgSignature: string(signature[:]),
	}, nil
}

// 检验消息的真实性，并且获取解密后的明文.
func (srv *Server) decryptMsg(encryptedBase64 string, msg *Mixture, contentType string) error {
	// EncodingAESKey长度固定为43个字符，从a-z,A-Z,0-9共62个字符中选取。
	if len(srv.EncodingAESKey) != 43 {
		return errors.New("invalid aes key")
	}

	key, err := srv.getAESKey()
	if err != nil {
		return err
	}

	encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
	data, err := cbcDecrypt(key, encrypted, key)
	if err != nil {
		return err
	}

	length := binary.BigEndian.Uint32(data[16:20])
	if err := unmarshal(data[20:20+length], msg, contentType); err != nil {
		return err
	}

	return nil
}
