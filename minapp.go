package weapp

import (
	"net/http"
	"time"
)

type Client interface {
}

// 消息返回数据类型
type ContentType string

const (
	ContentTypeXML  = "XML"
	ContentTypeJSON = "JSON"
)

// 小程序账号信息
type AccountInfo struct {
	// 小程序后台配置: 小程序ID
	AppID string
	// 小程序后台配置: 小程序密钥
	AppSecret string
	// 小程序后台配置: 消息推送令牌
	Token string
	// 小程序后台配置: 消息加密密钥
	EncodingAESKey string
	// 小程序后台配置: 消息返回数据类型
	ContentType ContentType
	// 小程序后台配置: 消息是否加密
	IsEncrypted string
}

type client struct {
	httpClient *http.Client
	cache      Cache

	// 微信账号信息
	account AccountInfo
}

func (c *client) initialize() {
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
}

func NewClient(info AccountInfo, fns ...func(Client)) Client {
	c := &client{
		account: info,
	}
	for _, fn := range fns {
		fn(c)
	}

	c.initialize()

	return c
}

func WithHttpClient(c *http.Client) func(client) {
	return func(cli client) {
		cli.httpClient = c
	}
}

type Cache interface {
	// 获取数据
	Get(key string) (string, error)
	// @exp为0时表示不设置有效期
	Set(key string, val interface{}, exp time.Duration) error
	// 判断数据是否存在
	Exists(key string) (bool, error)
}

func WithCache(c Cache) func(client) {
	return func(cli client) {
		cli.cache = c
	}
}
