// Package notify 接收并处理微信通知
package notify

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/medivhzhan/weapp/message"
	"github.com/medivhzhan/weapp/util"
)

// Server 微信服务接收器
// dev: 删除不必要的字段
type Server struct {
	appID          string // 小程序 ID
	mchID          string // 商户号
	apiKey         string // 商户签名密钥
	token          string // 微信服务器验证令牌
	EncodingAESKey string // 消息加密密钥

	Writer  http.ResponseWriter
	Request *http.Request

	textMessageHandler  func(Text)    // 文本消息处理器
	cardMessageHandler  func(Card)    // 卡片消息处理器
	imageMessageHandler func(Image)   // 图片消息处理器
	eventHandler        func(Mixture) // 事件处理器
}

// Mixture 从微信服务器接收的混合消息体
type Mixture struct {
	XMLName  xml.Name        `xml:"xml" json:"-"`
	ID       int64           `json:"MsgId" xml:"MsgId"`                     // 消息 ID
	Type     message.MsgType `json:"MsgType" xml:"MsgType"`                 // 消息类型
	Event    EventType       `json:"event,omitempty" xml:"event,omitempty"` // 事件类型
	Sender   string          `json:"FromUserName" xml:"FromUserName"`       // 发送者的 openID
	Receiver string          `json:"ToUserName" xml:"ToUserName"`           // 小程序的原始ID
	Datetime int64           `json:"CreateTime" xml:"CreateTime"`           // 消息创建时间(整型）

	Text
	Card
	Image
}

// NewServer Create new Server
func NewServer(res http.ResponseWriter, req *http.Request) *Server {
	return &Server{
		Request: req,
		Writer:  res,
	}
}

// HandleTextMessage 新建 Server 并设置文本消息处理器
func (srv *Server) HandleTextMessage(fuck func(Text)) {
	srv.textMessageHandler = fuck
}

// HandleCardMessage 新建 Server 并设置卡片消息处理器
func (srv *Server) HandleCardMessage(fuck func(Card)) {
	srv.cardMessageHandler = fuck
}

// HandleImageMessage 新建 Server 并设置图片消息处理器
func (srv *Server) HandleImageMessage(fuck func(Image)) {
	srv.imageMessageHandler = fuck
}

// HandleEvent 新建 Server 并设置事件处理器
func (srv *Server) HandleEvent(fuck func(Mixture)) {
	srv.eventHandler = fuck
}

// Serve 启动服务
func (srv *Server) Serve() error {
	switch srv.Request.Method {
	case "POST":

		// 处理加密消息
		if encrypted(srv.Request) {
			return errors.New("SDK 暂时还不支持加密消息")
			// dev: handle encrypted message
		}

		body, err := ioutil.ReadAll(srv.Request.Body)
		if err != nil {
			return err
		}

		var mix Mixture
		switch t := srv.Request.Header.Get("Content-Type"); t {
		case "application/json":
			if err := json.Unmarshal(body, &mix); err != nil {
				return err
			}
		case "application/xml":
			if err := xml.Unmarshal(body, &mix); err != nil {
				return err
			}
		default:
			return errors.New("unknown content type: " + t)
		}

		switch mix.Type {
		case message.TextMsg: // 文本消息
			if srv.textMessageHandler != nil {
				msg := mix.Text
				srv.textMessageHandler(msg)
			}
		case message.ImgMsg: // 图片消息
			if srv.imageMessageHandler != nil {
				msg := mix.Image
				srv.imageMessageHandler(msg)
			}
		case message.CardMsg: // 卡片消息
			if srv.cardMessageHandler != nil {
				msg := mix.Card
				srv.cardMessageHandler(msg)
			}
		case message.Event: // 事件
			if srv.eventHandler != nil {
				srv.eventHandler(mix)
			}
		default:
			return errors.New("无效的消息类型: " + string(mix.Type))
		}

		srv.Writer.WriteHeader(http.StatusOK)
		_, err = io.WriteString(srv.Writer, "SUCCESS")
		return err
	case "GET":
		srv.Writer.WriteHeader(http.StatusOK)
		_, err := io.WriteString(srv.Writer, util.GetQuery(srv.Request, "echostr"))
		return err
	default:
		return errors.New("无效的请求方法: " + srv.Request.Method)
	}
}

// 判断是否加密消息
func encrypted(req *http.Request) bool {
	return util.GetQuery(req, "encrypt_type") == "aes"
}

// 检验消息的真实性，并且获取解密后的明文
func decrypt(ciphertext string) error {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(ciphertext) < BlockSize {
		return errors.New("cipher too short")
	}
	// ECB mode always works in whole blocks.
	if len(ciphertext)%BlockSize != 0 {
		return errors.New("cipher is not a multiple of the block size")
	}

	return nil
}

// 将公众号回复用户的消息加密打包
func encrypt(token, aesKey, appID, msg, nonce string, timestamp int64) error {
	// key, err := base64.StdEncoding.DecodeString(aesKey + "=")
	// if err != nil {
	// 	return err
	// }

	// if len(key) != 32 {
	// 	return errors.New("invalid encoding AES key")
	// }

	// str := util.RandomString(16) + msg + appID

	// bts := util.PKCS5Padding([]byte(str), aes.BlockSize)

	// ....

	// pc = Prpcrypt(self.key)
	//     ret,encrypt = pc.encrypt(sReplyMsg, self.appid)
	//     if ret != 0:
	//         return ret,None
	//     if timestamp is None:
	//         timestamp = str(int(time.time()))
	//     # 生成安全签名
	//     sha1 = SHA1()
	//     ret,signature = sha1.getSHA1(self.token, timestamp, sNonce, encrypt)
	//     if ret != 0:
	//         return ret,None
	//     xmlParse = XMLParse()
	//     return ret,xmlParse.generate(encrypt, signature, timestamp, sNonce)

	return nil
}
