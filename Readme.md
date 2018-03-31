# 微信小程序 SDK （for Golang）

## 使用

- 拉取代码

```sh
go get github.com/medivhzhan/weapp
```

- 引入

```go
import "github.com/medivhzhan/weapp"
```

## 获取 access_token 及其有效期

```go
// 获取次数有限制 获取后请缓存

weapp.AccessToken(appID, secret string) (accessToken string, expire uint, err error)

```

## 用户登录

```go
// 需要从小程序客户端获取到的code

weapp.Login(appID, secret, code string) (openID string, sessionKey uint, err error)

```

## 获取小程序码: 适用于需要的码数量较少的业务场景

```go
// AppCode 获取小程序码
// 可接受path参数较长 生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// @ path 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// @ lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// @ token 微信access_token
// @ filename 文件储存路径

weapp.AppCode(path string, width int, autoColor bool, lineColor, accessToken, filename string) error

```

## 获取小程序码: 适用于需要的码数量极多的业务场景

```go
// UnlimitedAppCode 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
// @ scene 需要使用 decodeURIComponent 才能获取到生成二维码时传入的 scene
// @ page 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// @ lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// @ token 微信access_token
// @ filename 文件储存路径

weapp.UnlimitedAppCode(scene, path string, width int, autoColor bool, lineColor, accessToken string) error

```

## 获取小程序二维码: 适用于需要的码数量较少的业务场景

```go
// QRCode 获取小程序二维码
// 可接受path参数较长，生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// @ path 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ token 微信access_token
// @ filename 文件储存路径

weapp.QRCode(path string, width int, token string) error

```

## 未实现功能

1. 保存图片
1. 客服消息
1. 模版消息
1. 微信支付
1. 数据统计
1. 临时素材
