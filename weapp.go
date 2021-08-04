package weapp

import (
	"net/http"

	"github.com/medivhzhan/weapp/v3/cache"
	"github.com/medivhzhan/weapp/v3/request"
	"github.com/medivhzhan/weapp/v3/subscribemessage"
)

const (
	// baseURL 微信请求基础URL
	baseURL = "https://api.weixin.qq.com"
)

type Client struct {
	// HTTP请求客户端
	request *request.Request
	// 数据缓存器
	cache cache.Cache
	// 小程序后台配置: 小程序ID
	appid string
	// 小程序后台配置: 小程序密钥
	secret string
}

// 初始化客户端
func newClient(appid, secret string) *Client {
	cli := Client{
		appid:   appid,
		secret:  secret,
		cache:   cache.NewMemoryCache(),
		request: request.NewRequest(http.DefaultClient, request.ContentTypeJSON),
	}

	return &cli
}

// 初始化客户端并用自定义配置替换默认配置
func NewClient(appid, secret string, opts ...func(*Client)) *Client {
	cli := newClient(appid, secret)

	// 执行额外的配置函数
	for _, fn := range opts {
		fn(cli)
	}

	return cli
}

func WithHttpClient(c *http.Client) func(*Client) {
	return func(cli *Client) {
		cli.request = request.NewRequest(c, request.ContentTypeJSON)
	}
}

func WithCache(c cache.Cache) func(*Client) {
	return func(cli *Client) {
		cli.cache = c
	}
}

// POST 参数
type requestParams map[string]interface{}

// URL 参数
type requestQueries map[string]string

// tokenAPI 获取带 token 的 API 地址
func tokenAPI(api, token string) (string, error) {
	queries := requestQueries{
		"access_token": token,
	}

	return request.EncodeURL(api, queries)
}

// convert bool to int
func bool2int(ok bool) uint8 {

	if ok {
		return 1
	}

	return 0
}

// 拼凑完整的 URI
func (cli *Client) conbineURI(url string, queries map[string]string) (string, error) {
	token, err := cli.AccessToken()
	if err != nil {
		return "", err
	}

	if queries == nil {
		queries = map[string]string{"access_token": token}
	} else {
		queries["access_token"] = token
	}

	return request.EncodeURL(baseURL+url, queries)
}

// 订阅消息
func (cli *Client) NewSubscribeMessage() *subscribemessage.SubscribeMessage {
	return subscribemessage.NewSubscribeMessage(cli.request, cli.conbineURI)
}
