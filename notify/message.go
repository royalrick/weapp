package notify

import "encoding/xml"

// EncryptedMsg 经过加密的消息体
type EncryptedMsg struct {
	XMLName  xml.Name `xml:"xml" json:"-"`
	Receiver string   `xml:"ToUserName" json:"ToUserName"`
	Message  string   `xml:"Encrypt" json:"Encrypt"`
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
