# ![title](title.png)

## `注意`

- [1.x 版本入口](https://github.com/medivhzhan/weapp/tree/v1)
- 2.0 版本开始支付相关内容会分离到一个单独的包。
- 文档按照[小程序服务端官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)排版，方便您一一对照查找相关内容。
- 目前 v2 还在测试阶段，请勿用于生产环境。

## `目录`

- [获取代码](#获取代码)
- [用户登录](#用户登录)

## 获取代码

```sh

go get -u github.com/medivhzhan/weapp

```

---

## 用户登录

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.Login("mock-appid", "mock-secret", "mock-code")
if err != nil {
    // 处理一般错误信息
    return
}

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

---
