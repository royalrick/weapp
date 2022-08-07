# ![微信小程序服务端 SDK (for Golang)](https://repository-images.githubusercontent.com/126961623/e561e692-7eac-4cd1-955b-f4fe3ff6f7b7)

## 说明

- [v1 版本入口](https://github.com/medivhzhan/weapp/tree/v1)
- [v2 版本入口](https://github.com/medivhzhan/weapp/tree/v2)
- [查看完整文档](https://pkg.go.dev/github.com/medivhzhan/weapp/v3)
- SDK 暂不包含支付相关内容 已有很多优秀的支付相关模块;
- 微信小程序的功能和接口一直在持续更新迭代,如果遇到没有的接口或者不符合当前实际情况的接口请提交 [issue](https://github.com/royalrick/weapp/issues/new) 或者发起 pull request;

## 获取代码

```sh

go get -u github.com/medivhzhan/weapp/v3

```

## 初始化

- 初始化 SDK

```go
package main

import (
 "github.com/medivhzhan/weapp/v3"
)

func main() {
 sdk := weapp.NewClient("your-appid", "your-secret")
}
```

- 自定义 HTTP 客户端

```go
package main

import (
 "crypto/tls"
 "net/http"
 "time"

 "github.com/medivhzhan/weapp/v3"
)

func main() {
 cli := &http.Client{
  Timeout: 10 * time.Second,
  Transport: &http.Transport{
   // 跳过安全校验
   TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  },
 }

 sdk := weapp.NewClient(
  "your-appid",
  "your-secret",
  weapp.WithHttpClient(cli),
 )
}

```

- 自定义日志

```go
package main

import (
 "log"
 "os"

 "github.com/medivhzhan/weapp/v3"
 "github.com/medivhzhan/weapp/v3/logger"
)

func main() {
 lgr := logger.NewLogger(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Info, true)

 sdk := weapp.NewClient(
  "your-appid",
  "your-secret",
  weapp.WithLogger(lgr),
 )

 // 任意切换日志等级
 sdk.SetLogLevel(logger.Silent)
}

```

- 自定义缓存

```go
package main

import (
 "time"

 "github.com/medivhzhan/weapp/v3"
)

type MyCache struct{}

func (cc *MyCache) Set(key string, val interface{}, timeout time.Duration) {
 // ...
}

func (cc *MyCache) Get(key string) (interface{}, bool) {
 return "your-access-token", true
}

func main() {
 cc := new(MyCache)

 sdk := weapp.NewClient(
  "your-appid",
  "your-secret",
  weapp.WithCache(cc),
 )
}

```

- 自定义 token 获取方法

```go
package main

import (
 "github.com/medivhzhan/weapp/v3"
)

func main() {
 tokenGetter := func() (token string, expireIn uint) {

  expireIn = 1000
  token = "your-custom-token"

  return token, expireIn
 }

 sdk := weapp.NewClient(
  "your-appid",
  "your-secret",
  weapp.WithAccessTokenSetter(tokenGetter),
 )
}

```

---

## 调用接口示例

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)

```go
package main

import (
 "fmt"
 "log"

 "github.com/medivhzhan/weapp/v3"
 "github.com/medivhzhan/weapp/v3/auth"
)

func main() {
 sdk := weapp.NewClient("your-appid", "your-secret")

 cli := sdk.NewAuth()

    // 用户支付完成后获取该用户的 UnionId
 rsp, err := cli.GetPaidUnionId(&auth.GetPaidUnionIdRequest{})
 if err != nil {
  log.Fatal(err)
 }

    // 检查加密信息是否由微信生成
 rsp, err := cli.CheckEncryptedData(&auth.CheckEncryptedDataRequest{})
 if err != nil {
  log.Fatal(err)
 }

    // 登录凭证校验
 rsp, err := cli.Code2Session(&auth.Code2SessionRequest{})
 if err != nil {
  log.Fatal(err)
 }

    // 获取小程序全局唯一后台接口调用凭据
 rsp, err := cli.GetAccessToken(&auth.GetAccessTokenRequest{})
 if err != nil {
  log.Fatal(err)
 }

    // 检查微信是否返回错误
 if err := rsp.GetResponseError(); err != nil {
  log.Println(err)
 }

 fmt.Println(rsp)
}

```

---

## 接收微信通知

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html#option-url)

```go
package main

import (
 "log"
 "net/http"

 "github.com/medivhzhan/weapp/v3"
 "github.com/medivhzhan/weapp/v3/server"
)

func main() {
 sdk := weapp.NewClient("your-appid", "your-secret")

 //  通用处理器
 handler := func(req map[string]interface{}) map[string]interface{} {
  switch req["MsgType"] {
  case "text":
   // Do something cool ...
  }

  return nil
 }

    // HTTP handler
 http.HandleFunc("/wechat/notify", func(w http.ResponseWriter, r *http.Request) {
  srv, err := sdk.NewServer("token", "aesKey", "mchID", "apiKey", false, handler)
  if err != nil {
   log.Fatalf("init server error: %s", err)
  }

  // 调用事件处理器后 通用处理器不再处理该事件
  srv.OnCustomerServiceTextMessage(func(tmr *server.TextMessageResult) *server.TransferCustomerMessage {

   return &server.TransferCustomerMessage{}
  })

  if err := srv.Serve(w, r); err != nil {
   log.Fatalf("serving error: %s", err)
  }
 })
}

```
