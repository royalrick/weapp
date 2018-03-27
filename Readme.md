# GO 微信小程序 SDK

## 使用

- 拉取代码

```sh
go get github.com/medivhzhan/weapp
```

- 引入

```go
import "github.com/medivhzhan/weapp"
```

- 新建实例

需要把下列变量配置到全局
    WEAPP_APPID
    WEAPP_SECRET
    WEAPP_TOKEN
    WEAPP_AES_KEY

```go
app := weapp.New()
```

- OR

```go
app := &weapp.WeApp{
        AppID:  "Your APPID",
        Secret: "Your SECRET",
        Token:  "Your TOKEN",
        AesKey: "Your AES_KEY",
    }
```

## 登录

```go
app := weapp.New()

// 需要从小程序客户端获取到的code
openid, session_key, err := app.Login(code string)
```

## 获取小程序码

```go
app := weapp.New()

// AppCode 获取小程序码
// path 识别二维码后进入小程序的页面链接
// width 图片宽度 目前最小只能 280
// autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// 返回小程序码HTTP请求
res, err := app.AppCode(path string, width int, autoColor bool, lineColor string)
```

## 获取 access_token

```go
app := weapp.New()

// 获取次数有限制 获取后请缓存
access_token, err := app.AccessToken()
```