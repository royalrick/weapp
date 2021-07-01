package weapp

import (
	"net/http"

	"github.com/medivhzhan/weapp/v3/cache"
	"github.com/medivhzhan/weapp/v3/request"
)

type Client interface {
	Login(code string) (*LoginResponse, error)
}

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
	ContentType request.ContentType
	// 小程序后台配置: 消息是否加密
	IsEncrypted string
}

type client struct {
	// HTTP请求客户端
	request *request.Request
	// 数据缓存器
	cache cache.Cache
	// 小程序后台配置: 消息返回数据类型
	contentType request.ContentType
	// 微信账号信息
	account AccountInfo
}

// 初始化客户端
func newClient(info AccountInfo) *client {
	cli := client{
		account:     info,
		cache:       cache.NewMemoryCache(),
		contentType: info.ContentType,
		request:     request.NewRequest(http.DefaultClient, info.ContentType),
	}

	return &cli
}

// 初始化客户端并用自定义配置替换默认配置
func NewClient(info AccountInfo, opts ...func(Client)) Client {
	cli := newClient(info)

	// 执行额外的配置函数
	for _, fn := range opts {
		fn(cli)
	}

	return cli
}

func WithHttpClient(c *http.Client) func(client) {
	return func(cli client) {
		cli.request = request.NewRequest(c, cli.contentType)
	}
}

func WithCache(c cache.Cache) func(client) {
	return func(cli client) {
		cli.cache = c
	}
}
