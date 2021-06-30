package weapp

import (
	"net/http"

	"github.com/medivhzhan/weapp/v3/cache"
)

type Client interface {
	Login(code string) (*LoginResponse, error)
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
	// HTTP请求客户端
	httpClient *http.Client

	// 数据缓存器
	cache cache.Cache
	// 缓存前缀
	cachePrefix string

	// 微信账号信息
	account AccountInfo
}

// 初始化客户端
func newClient(info AccountInfo) *client {
	cli := client{
		account:     info,
		cache:       cache.NewMemoryCache(),
		cachePrefix: "weapp",
		httpClient:  http.DefaultClient,
	}

	return &cli
}

// 初始化客户端并用自定义配置替换默认配置
func NewClient(info AccountInfo, fns ...func(Client)) Client {

	cli := newClient(info)

	// 执行额外的配置函数
	for _, fn := range fns {
		fn(cli)
	}

	return cli
}

func WithHttpClient(c *http.Client) func(client) {
	return func(cli client) {
		cli.httpClient = c
	}
}

func WithCache(c cache.Cache) func(client) {
	return func(cli client) {
		cli.cache = c
	}
}

// 配置缓存器前缀
func WithCachePrefix(s string) func(client) {
	return func(cli client) {
		cli.cachePrefix = s
	}
}
