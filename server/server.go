// Package server 接收微信服务
package server

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/message"
)

// Server 微信服务接收器
type Server struct {
	appID          string // 小程序 ID
	mchID          string // 商户号
	apiKey         string // 商户签名密钥
	token          string // 微信服务器验证令牌
	EncodingAESKey string // 消息加密密钥
	messageHandler func(message.MixMessage)
}

// HandleMessage 新建 Server 并设置消息处理器
func HandleMessage(fuck func(message.MixMessage)) *Server {
	return &Server{
	// messageHandler: fuck,
	}
}

// Serve 启动服务
func (srv *Server) Serve(res http.ResponseWriter, req *http.Request) error {
	switch req.Method {
	case "POST":

		// 处理加密消息
		if encrypted(req) {
			return errors.New("SDK 暂时还不支持加密消息")
			// dev: handle encrypted message
		}

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		var msg message.MixMessage
		switch t := req.Header.Get("Content-Type"); t {
		case "application/json":
			if err := json.Unmarshal(body, &msg); err != nil {
				return err
			}
		case "application/xml":
			if err := xml.Unmarshal(body, &msg); err != nil {
				return err
			}
		default:
			return errors.New("unknown content type: " + t)
		}

		switch msg.MsgType {
		case message.Text, message.Img, message.Card: // 消息

			if srv.messageHandler != nil {
				srv.messageHandler(msg)
			}
		case message.Event: // 事件

			switch msg.Event {
			case message.Enter:
				// dev: handle events
			}
		default:
			return errors.New("unknown message type： " + string(msg.MsgType))
		}

		return nil
	case "GET":
		_, err := io.WriteString(res, weapp.GetQuery(req, "echostr"))
		return err
	default:
		return errors.New("unexpected HTTP Method: " + req.Method)
	}
}

// 判断是否加密消息
func encrypted(req *http.Request) bool {
	return weapp.GetQuery(req, "encrypt_type") == "aes"
}
