package message

import (
	"encoding/xml"
	"time"
)

//Encrypter 经过加密的消息体
type Encrypter struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	EncryptedMsg string   `xml:"Encrypt"    json:"Encrypt"`
}

// RMsgText 接收的文本消息
type RMsgText struct {
	Content string `json:"Content,omitempty" xml:"Content,omitempty"`
}

// RMsgImg 接收的图片消息
type RMsgImg struct {
	PicURL  string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	MediaID string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

// RMsgCard 接收的卡片消息
type RMsgCard struct {
	Title        string `json:"Title,omitempty" xml:"Title,omitempty"`               // 标题
	AppID        string `json:"AppId,omitempty" xml:"AppId,omitempty"`               // 小程序 appid
	PagePath     string `json:"PagePath,omitempty" xml:"PagePath,omitempty"`         // 小程序页面路径
	ThumbURL     string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`         // 封面图片的临时cdn链接
	ThumbMediaID string `json:"ThumbMediaId,omitempty" xml:"ThumbMediaId,omitempty"` // 封面图片的临时素材id
}

// MixMessage 从微信服务器接收的混合消息体
type MixMessage struct {
	XMLName      xml.Name      `xml:"xml" json:"-"`
	ToUserName   string        `json:"ToUserName" xml:"ToUserName"`     // 小程序的原始ID
	FromUserName string        `json:"FromUserName" xml:"FromUserName"` // 发送者的 openID
	CreateTime   time.Duration `json:"CreateTime" xml:"CreateTime"`     // 消息创建时间(整型）
	MsgType      MsgType       `json:"MsgType" xml:"MsgType"`
	MsgID        int64         `json:"MsgId" xml:"MsgId"` // 消息 ID

	Event EventType `json:"event,omitempty" xml:"event,omitempty"` // 事件类型

	RMsgText
	RMsgImg
	RMsgCard
}
