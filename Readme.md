# WeAPP

> Go 微信小程序 SDK

## 使用

- 拉取代码

```sh
go get github.com/medivhzhan/weapp
```

- 引入

```go
import "github.com/medivhzhan/weapp"
```

- 初始化小程序

```go
// 初始化小程序才能继续调用下面的接口
weapp.Init(appID, secret, token, aesKey string)
```

## 登录

```go
// 需要从小程序客户端获取到的code
openID, sessionKey, err := weapp.Login(code string)
```

## 获取小程序码(圆形)

```go
// AppCode 获取小程序码
// path 识别二维码后进入小程序的页面链接
// width 图片宽度 目前最小只能 280
// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"105","g":"195","b":"170"}, 十进制表示
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
res, err := weapp.AppCode(path string, width int, autoColor bool, lineColor, accessToken string)
```

## 获取 access_token 及其有效期

```go
// 获取次数有限制 获取后请缓存
accessToken, expire, err := weapp.AccessToken()
```