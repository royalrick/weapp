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
	EventUserEntry       EventType = "user_enter_tempsession" // 用户进入临时会话状态
	EventGetQuota                  = "get_quota"              // 查询商户余额
	EventAsyncMediaCheck           = "wxa_media_check"        // 异步校验图片/音频
)

// EncryptedMsgResponse 接收的的加密消息格式
type EncryptedMsgResponse struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	ToUserName string   `json:"ToUserName" xml:"ToUserName"` // 接收者 为小程序 AppID
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
	MsgID        int       `json:"MsgId" xml:"MsgId"`               // 消息 ID
	ToUserName   string    `json:"ToUserName" xml:"ToUserName"`     // 小程序的原始ID
	FromUserName string    `json:"FromUserName" xml:"FromUserName"` // 发送者的 openID | 平台推送服务UserName
	CreateTime   uint64    `json:"CreateTime" xml:"CreateTime"`     // 消息创建时间(整型）
	MsgType      MsgType   `json:"MsgType" xml:"MsgType"`           // 消息类型
	Event        EventType `json:"Event" xml:"Event"`               // 事件类型
	SessionFrom  string    `json:"SessionFrom" xml:"SessionFrom"`
	Text
	Card
	Image
	AsyncMedia

	RawData map[string]interface{} `json:"-" xml:"-"` // 原始数据
}

// Serverhandler 服务处理器
type Serverhandler = func(*Mixture) bool

// Server 微信通知服务处理器
type Server struct {
	appID    string // 小程序 ID
	mchID    string // 商户号
	apiKey   string // 商户签名密钥
	token    string // 微信服务器验证令牌
	aesKey   []byte // base64 解码后的消息加密密钥
	validate bool   // 是否验证请求来自微信服务器
	Handler  Serverhandler
}

type dataType = string

const (
	dataTypeJSON dataType = "application/json"
	dataTypeXML           = "application/xml"
)

// NewServer 返回经过初始化的Server
func NewServer(appID, token, aesKey, mchID, apiKey string, validate bool, handler Serverhandler) (*Server, error) {

	key, err := base64.RawStdEncoding.DecodeString(aesKey)
	if err != nil {
		return nil, err
	}

	server := Server{
		appID:    appID,
		mchID:    mchID,
		apiKey:   apiKey,
		token:    token,
		aesKey:   key,
		validate: validate,
		Handler:  handler,
	}

	return &server, nil
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

// HandleRequest 接收并处理微信通知服务
func (srv *Server) HandleRequest(w http.ResponseWriter, r *http.Request) error {
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

			nonce := getQuery(r, "nonce")
			signature := getQuery(r, "msg_signature")
			timestamp := getQuery(r, "timestamp")

			// 检验消息的真实性
			if !validateSignature(signature, srv.token, timestamp, nonce, res.Encrypt) {
				return errors.New("invalid signature")
			}
			body, err := srv.decryptMsg(res.Encrypt)
			if err != nil {
				return err
			}
			length := binary.BigEndian.Uint32(body[16:20])
			raw = body[20 : 20+length]
		}
		if err := unmarshal(raw, tp, mix); err != nil {
			return err
		}

		if err := unmarshal(raw, tp, &mix.RawData); err != nil {
			return err
		}

		ok := srv.Handler(mix)
		if ok {
			_, err := io.WriteString(w, "SUCCESS")
			if err != nil {
				return err
			}
		}

		return nil
	case "GET":
		echostr := getQuery(r, "echostr")
		if srv.validate {

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
	return getQuery(req, "encrypt_type") == "aes"
}

// 验证消息的确来自微信服务器
// 1.将token、timestamp、nonce三个参数进行字典序排序
// 2.将三个参数字符串拼接成一个字符串进行sha1加密
// 3.开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
func (srv *Server) validateServer(req *http.Request) bool {
	nonce := getQuery(req, "nonce")
	signature := getQuery(req, "signature")
	timestamp := getQuery(req, "timestamp")

	return validateSignature(signature, nonce, timestamp, srv.token)
}

// 加密消息
func (srv *Server) encryptMsg(message, nonce string, timestamp int) (*EncryptedMsgRequest, error) {

	key := srv.aesKey

	//获得16位随机字符串，填充到明文之前
	random := randomString(16)
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
func (srv *Server) decryptMsg(encrypted string) ([]byte, error) {

	key := srv.aesKey

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	data, err := cbcDecrypt(key, ciphertext, key)
	if err != nil {
		return nil, err
	}

	return data, nil
}
