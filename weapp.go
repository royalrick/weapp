package weapp

const (
	// baseURL 微信请求基础URL
	baseURL = "https://api.weixin.qq.com"
)

// POST 参数
type requestParams map[string]interface{}

// URL 参数
type requestQueries map[string]string
