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

## 获取小程序码: 适用于需要的码数量较少的业务场景

```go
// 可接受path参数较长 生成个数受限 永久有效
// path 识别二维码后进入小程序的页面链接
// width 图片宽度
// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// token access_token
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
res, err := weapp.AppCode(path string, width int, autoColor bool, lineColor, accessToken string)
if err != nil {
    panic(err)
}
defer res.Body.Close()
```

## 获取小程序码: 适用于需要的码数量极多的业务场景

```go
// 可接受页面参数较短 生成个数不受限
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
// scene 需要使用 decodeURIComponent 才能获取到生成二维码时传入的 scene
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
res, err := weapp.UnlimitedAppCode(scene, path string, width int, autoColor bool, lineColor, accessToken string)
if err != nil {
    panic(err)
}
defer res.Body.Close()
```

## 获取小程序二维码: 适用于需要的码数量较少的业务场景

```go
// QRCode 获取小程序二维码
// 可接受path参数较长，生成个数受限 永久有效
// 返回小程序码HTTP请求
// 请记得关闭资源
// 获取后请注意保存到本地以减少请求次数
res, err := weapp.QRCode(path string, width int, token string)
if err != nil {
    panic(err)
}
defer res.Body.Close()
```

## 获取 access_token 及其有效期

```go
// 获取次数有限制 获取后请缓存
accessToken, expire, err := weapp.AccessToken()
```