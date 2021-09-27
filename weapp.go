package weapp

import (
	"log"
	"net/http"
	"os"

	"github.com/medivhzhan/weapp/v3/auth"
	"github.com/medivhzhan/weapp/v3/cache"
	"github.com/medivhzhan/weapp/v3/livebroadcast"
	"github.com/medivhzhan/weapp/v3/logger"
	"github.com/medivhzhan/weapp/v3/ocr"
	"github.com/medivhzhan/weapp/v3/operation"
	"github.com/medivhzhan/weapp/v3/request"
	"github.com/medivhzhan/weapp/v3/search"
	"github.com/medivhzhan/weapp/v3/security"
	"github.com/medivhzhan/weapp/v3/server"
	"github.com/medivhzhan/weapp/v3/subscribemessage"
	"github.com/medivhzhan/weapp/v3/updatablemessage"
	"github.com/medivhzhan/weapp/v3/wxacode"
	"github.com/mitchellh/mapstructure"
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
	// 日志记录器
	logger logger.Logger
	// 小程序后台配置: 小程序ID
	appid string
	// 小程序后台配置: 小程序密钥
	secret string
	// 用户自定义获取access_token的方法
	accessTokenGetter AccessTokenGetter
}

// 用户自定义获取access_token的方法
type AccessTokenGetter func() (token string, expireIn uint)

// 初始化客户端并用自定义配置替换默认配置
func NewClient(appid, secret string, opts ...func(*Client)) *Client {
	cli := &Client{
		appid:  appid,
		secret: secret,
	}

	// 执行额外的配置函数
	for _, fn := range opts {
		fn(cli)
	}

	if cli.cache == nil {
		cli.cache = cache.NewMemoryCache()
	}

	if cli.request == nil {
		cli.request = request.NewRequest(http.DefaultClient, request.ContentTypeJSON, cli.Logger)
	}

	if cli.logger == nil {
		cli.logger = logger.NewLogger(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Info, true)
	}

	return cli
}

// 自定义 HTTP Client
func WithHttpClient(hc *http.Client) func(*Client) {
	return func(cli *Client) {
		cli.request = request.NewRequest(hc, request.ContentTypeJSON, cli.Logger)
	}
}

// 自定义缓存
func WithCache(cc cache.Cache) func(*Client) {
	return func(cli *Client) {
		cli.cache = cc
	}
}

// 自定义获取access_token的方法
func WithAccessTokenSetter(getter AccessTokenGetter) func(*Client) {
	return func(cli *Client) {
		cli.accessTokenGetter = getter
	}
}

// 自定义日志
func WithLogger(logger logger.Logger) func(*Client) {
	return func(cli *Client) {
		cli.logger = logger
	}
}

// POST 参数
type requestParams map[string]interface{}

// URL 参数
type requestQueries map[string]interface{}

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

// 获取日志记录器
func (cli *Client) Logger() logger.Logger { return cli.logger }

// 设置日志等级
func (cli *Client) SetLogLevel(lv logger.Level) {
	if cli.logger != nil {
		cli.logger.SetLevel(lv)
	}
}

// 拼凑完整的 URI
func (cli *Client) conbineURI(url string, req interface{}) (string, error) {

	output := make(map[string]interface{})

	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &output,
		TagName:  "query",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return "", err
	}

	err = decoder.Decode(req)
	if err != nil {
		return "", err
	}

	token, err := cli.AccessToken()
	if err != nil {
		return "", err
	}

	output["access_token"] = token

	return request.EncodeURL(baseURL+url, output)
}

// 用户信息
func (cli *Client) NewAuth() *auth.Auth {
	return auth.NewAuth(cli.request, cli.conbineURI)
}

// 微信通知监听服务
func (cli *Client) NewServer(token, aesKey, mchID, apiKey string, validate bool, handler func(map[string]interface{}) map[string]interface{}) (*server.Server, error) {
	return server.NewServer(cli.appid, token, aesKey, mchID, apiKey, validate, handler)
}

// 订阅消息
func (cli *Client) NewSubscribeMessage() *subscribemessage.SubscribeMessage {
	return subscribemessage.NewSubscribeMessage(cli.request, cli.conbineURI)
}

// 运维中心
func (cli *Client) NewOperation() *operation.Operation {
	return operation.NewOperation(cli.request, cli.conbineURI)
}

// 小程序码
func (cli *Client) NewWXACode() *wxacode.WXACode {
	return wxacode.NewWXACode(cli.request, cli.conbineURI)
}

// OCR
func (cli *Client) NewOCR() *ocr.OCR {
	return ocr.NewOCR(cli.request, cli.conbineURI)
}

// 动态消息
func (cli *Client) NewUpdatableMessage() *updatablemessage.UpdatableMessage {
	return updatablemessage.NewUpdatableMessage(cli.request, cli.conbineURI)
}

// 小程序搜索
func (cli *Client) NewSearch() *search.Search {
	return search.NewSearch(cli.request, cli.conbineURI)
}

// 直播
func (cli *Client) NewLiveBroadcast() *livebroadcast.LiveBroadcast {
	return livebroadcast.NewLiveBroadcast(cli.request, cli.conbineURI)
}

// 内容安全
func (cli *Client) NewSecurity() *security.Security {
	return security.NewSecurity(cli.request, cli.conbineURI)
}
