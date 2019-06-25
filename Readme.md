# 微信小程序 SDK （for Golang）

## 目录

- [拉取代码](#拉取代码)
- [AccessToken](#AccessToken)
- [用户登录](#用户登录)
- [二维码](#二维码)
  - [获取小程序码](#获取小程序码)
  - [获取小程序二维码](#获取小程序二维码)
- [模板消息](#模板消息)
  - [获取小程序模板库标题列表](#获取小程序模板库标题列表)
  - [获取帐号下已存在的模板列表](#获取帐号下已存在的模板列表)
  - [获取模板库某个模板标题下关键词库](#获取模板库某个模板标题下关键词库)
  - [组合模板并添加至帐号下的个人模板库](#组合模板并添加至帐号下的个人模板库)
  - [删除帐号下的某个模板](#删除帐号下的某个模板)
  - [发送模板消息](#发送模板消息)
- [统一服务消息](#统一服务消息)
- [客服消息](#客服消息)
  - [接收客服消息](#接收客服消息)
  - [发送客服消息](#发送客服消息)
- [支付](#支付)
  - [付款](#付款)
  - [处理支付结果通知](#处理支付结果通知)
  - [退款](#退款)
  - [处理退款结果通知](#处理退款结果通知)
  - [转账(企业付款)](#转账(企业付款))
  - [查询转账](#查询转账)
  - [订单查询](#订单查询)
  - [支付后获取UnionID](#支付后获取UnionID)
- [解密](#解密)
  - [解密手机号码](#解密手机号码)
  - [解密分享内容](#解密分享内容)
  - [解密用户信息](#解密用户信息)
- [内容检测](#内容检测)
  - [检测图片](#检测图片)
  - [检测文本](#检测文本)
- [生物认证秘钥签名验证](#生物认证秘钥签名验证)

## 拉取代码

```sh

go get -u github.com/medivhzhan/weapp

```

## AccessToken

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html)

```go

import "github.com/medivhzhan/weapp/token"

// 获取次数有限制 获取后请缓存
tok, exp, err := token.AccessToken(appID, secret string)

```

## 用户登录

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

```go

import "github.com/medivhzhan/weapp"

// @appID 小程序 appID
// @secret 小程序的 app secret
// @code 小程序登录时获取的 code
res, err := weapp.Login(appID, secret, code)
if err != nil {
    // handle error
}

// res.OpenID
// res.SessionKey
// res.UnionID
fmt.Printf("返回结果: %#v", res)


```

---

## 二维码

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html)

### 获取小程序码

需要二维码数量较少的业务场景

```go

import "github.com/medivhzhan/weapp/code"

coder := code.QRCoder {
    Path: "pages/index?query=1", // 识别二维码后进入小程序的页面链接
    Width: 430, // 图片宽度
    IsHyaline: true, // 是否需要透明底色
    AutoColor: true, // 自动配置线条颜色, 如果颜色依然是黑色, 则说明不建议配置主色调
    LineColor: code.Color{ //  AutoColor 为 false 时生效, 使用 rgb 设置颜色 十进制表示
        R: "50",
        G: "50",
        B: "50",
    },
}

// token: 微信 access_token
res, err := coder.AppCode(token string)
if err != nil {
    // handle error
}
defer res.Body.Close()
```

需要二维码数量极多的业务场景

```go
coder := code.QRCoder {
    Scene: "...", // 参数数据
    Page: "pages/index", // 识别二维码后进入小程序的页面链接
    Width: 430, // 图片宽度
    IsHyaline: true, // 是否需要透明底色
    AutoColor: true, // 自动配置线条颜色, 如果颜色依然是黑色, 则说明不建议配置主色调
    LineColor: code.Color{ //  AutoColor 为 false 时生效, 使用 rgb 设置颜色 十进制表示
        R: "50",
        G: "50",
        B: "50",
    },
}

// token: 微信 access_token
res, err := coder.UnlimitedAppCode(token string)
if err != nil {
    // handle error
}
defer res.Body.Close()

```

### 获取小程序二维码

适用于需要的码数量较少的业务场景

```go

import "github.com/medivhzhan/weapp/code"

coder := code.QRCoder {
    Path: "pages/index?query=1", // 识别二维码后进入小程序的页面链接
    Width: 430, // 图片宽度
}

// 获取小程序二维码
// token: 微信access_token
res, err := coder.QRCode(token string)
if err != nil {
    // handle error
}
defer res.Body.Close()

```

---

## 模板消息

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.addTemplate.html)

### 获取小程序模板库标题列表

```go

import "github.com/medivhzhan/weapp/message/template"

// 获取小程序模板库标题列表
// offset: 开始获取位置 从0开始
// count: 获取记录条数 最大为20
// token: 微信 access_token
list, total, err := template.List(offset uint, count uint, token string)

```

### 获取帐号下已存在的模板列表

```go

import "github.com/medivhzhan/weapp/message/template"

// 获取帐号下已存在的模板列表
// offset: 开始获取位置 从0开始
// count: 获取记录条数 最大为20
// token: 微信 access_token
list, total, err := template.Selves(offset uint, count uint, token string)
```

### 获取模板库某个模板标题下关键词库

```go

import "github.com/medivhzhan/weapp/message/template"

// 获取模板库某个模板标题下关键词库
// id: 模板ID
// token: 微信 access_token
keywords, err := template.Get(id, token string)

```

### 组合模板并添加至帐号下的个人模板库

```go

import "github.com/medivhzhan/weapp/message/template"

// 组合模板并添加至帐号下的个人模板库
// id: 模板ID
// token: 微信 aceess_token
// keywordIDList: 关键词 ID 列表
// 返回新建模板ID和错误信息
tid, err := template.Add(id, token string, keywordIDList []uint)

```

### 删除帐号下的某个模板

```go

import "github.com/medivhzhan/weapp/message/template"

// 删除帐号下的某个模板
// id: 模板ID
// token: 微信 aceess_token
err := template.Delete(id, token string)

```

### 发送模板消息

```go

import "github.com/medivhzhan/weapp/message/template"

msg := template.Message{
    "keyword1": "content ...",
    "keyword2": "content ...",
}

// 发送模板消息
// openid: 接收者（用户）的 openid
// template: 所需下发的模板消息的id
// page: 点击模板卡片后的跳转页面, 仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
// formID: 表单提交场景下, 为 submit 事件带上的 formId；支付场景下, 为本次支付的 prepay_id
// data: 模板内容, 不填则下发空模板
// emphasisKeyword: 模板需要放大的关键词, 不填则默认无放大
err := template.Send(openid, template, page, formID string, msg template.Message, emphasisKeyword, token string)

```

---

## 统一服务消息

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html)

```go

// 消息体
msg := template.UniformMsg{
    ToUser: "用户 openid",
    // 小程序模板消息
    WeappTemplateMsg: template.WeappTemplateMsg{
        TemplateID:      schedule.ActivityWillStartTemplateID,
        Page:            "pages/messages/main",
        FormID:          "1537411865951",
        EmphasisKeyword: "keyword1.DATA",
        Data: template.Data{
            "keyword1": template.Keyword{
                Value: "恭喜你购买成功！",
                Color: "#173177",
            },
            "keyword2": template.Keyword{
                Value: "巧克力",
                Color: "#173177",
            },
            "keyword3": template.Keyword{
                Value: "39.8元",
                Color: "#173177",
            },
        },
    },
    // 公众号模板消息
    MPTemplateMsg: template.MPTemplateMsg{
        AppID:      "wx2c5a33d31b4ee88f",
        TemplateID: "UmuX15eBoonYkLy-7Xle1rA6xHhv4bsbie1Viidg2Cs",
        URL:        "https://medivhzhan.me",
        Miniprogram: template.Miniprogram{
            AppID:    "wx7ad9cfdc85a2fdb2",
            Pagepath: "pages/me/main",
        },
        Data: template.Data{
            "first": template.Keyword{
                Value: "恭喜你购买成功！",
                Color: "#173177",
            },
            "keyword1": template.Keyword{
                Value: "巧克力",
                Color: "#173177",
            },
            "remark": template.Keyword{
                Value: "remark content ...",
                Color: "#173177",
            },
        },
    },
}

err := msg.Send(token)

```

---

## 客服消息

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html)

### 接收客服消息

```go

import "github.com/medivhzhan/weapp/notify"

// 新建服务
srv := notify.NewServer(http.ResponseWriter, *http.Request)

srv.HandleTextMessage(func(msg notify.Text)) {
    // 处理文本消息
})

srv.HandleCardMessage(func(msg notify.Card)) {
    // 处理卡片消息
})

srv.HandleImageMessage(func(msg notify.Image)) {
    // 处理图片消息
})

srv.HandleEvent(func(msg notify.Event)) {
    // 处理事件
})

// 执行服务
err := srv.Serve()

```

### 发送客服消息

```go

import "github.com/medivhzhan/weapp/message"

// 文本消息
msg := message.Text{
    Content: "消息内容",
}

// 图片消息
msg := message.Image{
    MediaID: "微信 media_id"
}

// 图文链接消息
msg := message.Link{
    Title: "标题"
    Description: "描述"
    URL: "点击跳转链接"
    ThumbURL: "图片链接"
}

// 卡片消息
msg := message.Card{
    Title: "标题"
    PagePath: "小程序页面路径"
    ThumbMediaID: "卡封面图片 media_id"
}

// 发送消息
// openid: 用户 openid
// token: 微信 access_token
res, err := msg.SendTo(openid, token string)

```

---

## 支付

### 付款

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1)

```go

import "github.com/medivhzhan/weapp/payment"

// 新建支付订单
form := payment.Order{
    // 必填
    AppID:      "APPID",
    MchID:      "商户号",
    Body:       "商品描述",
    NotifyURL:  "通知地址",
    OpenID:     "通知用户的 openid",
    OutTradeNo: "商户订单号",
    TotalFee:   "总金额(分)",

    // 选填 ...
    IP:        "发起支付终端IP",
    NoCredit:  "是否允许使用信用卡",
    StartedAt: "交易起始时间",
    ExpiredAt: "交易结束时间",
    Tag:       "订单优惠标记",
    Detail:    "商品详情",
    Attach:    "附加数据",
}

res, err := form.Unify("支付密钥")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

// 获取小程序前点调用支付接口所需参数
params, err := payment.GetParams(res.AppID, "微信支付密钥", res.NonceStr, res.PrePayID)
if err != nil {
    // handle error
    return
}

```

### 处理支付结果通知

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8)

```go

import "github.com/medivhzhan/weapp/payment"

// 必须在下单时指定的 notify_url 的路由处理器下
err := payment.HandlePaidNotify(w http.ResponseWriter, req *http.Request,  func(ntf payment.PaidNotify) (bool, string) {
    // 处理通知
    fmt.Printf("%#v", ntf)

    // 处理成功 return true, ""
    // or
    // 处理失败 return false, "失败原因..."
})

```

### 退款

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_4)

```go

import "github.com/medivhzhan/weapp/payment"

// 新建退款订单
form := payment.Refunder{
    // 必填
    AppID:       "APPID",
    MchID:       "商户号",
    TotalFee:    "总金额(分)",
    RefundFee:   "退款金额(分)",
    OutRefundNo: "商户退款单号",
    // 二选一
    OutTradeNo: "商户订单号", // or TransactionID: "微信订单号",

    // 选填 ...
    RefundDesc: "退款原因",   // 若商户传入, 会在下发给用户的退款消息中体现退款原因
    NotifyURL:  "结果通知地址", // 覆盖商户平台上配置的回调地址
}

// 需要证书
res, err := form.Refund("支付密钥",  "cert 证书路径", "key 证书路径")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 处理退款结果通知

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_5)

```go

import "github.com/medivhzhan/weapp/payment"

// 必须在商户平台上配置的回调地址或者发起退款时指定的 notify_url 的路由处理器下
err := payment.HandleRefundedNotify(w http.ResponseWriter, req *http.Request,  "支付密钥", func(ntf payment.RefundedNotify) (bool,         // 处理通知
    fmt.Printf("%#v", ntf)

    // 处理成功 return true, ""
    // or
    // 处理失败 return false, "失败原因..."
})

```

### 转账(企业付款到零钱)

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2)

```go

import "github.com/medivhzhan/weapp/payment"

// 新建退款订单
form := payment.Transferer{
    // 必填 ...
    AppID:       "APPID",
    MchID:       "商户号",
    Amount:      "总金额(分)",
    OutRefundNo: "商户退款单号",
    OutTradeNo:  "商户订单号", // or TransactionID: "微信订单号",
    ToUser:      "转账目标用户的 openid",
    Desc:        "转账描述", // 若商户传入, 会在下发给用户的退款消息中体现退款原因

    // 选填 ...
    IP: "发起转账端 IP 地址", // 若商户传入, 会在下发给用户的退款消息中体现退款原因
    CheckName: "校验用户姓名选项 true/false",
    RealName: "收款用户真实姓名", // 如果 CheckName 设置为 true 则必填用户真实姓名
    Device:   "发起转账设备信息",
}

// 需要证书
res, err := form.Transfer("支付密钥",  "cert 证书路径", "key 证书路径")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 查询转账

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3)

```go

import "github.com/medivhzhan/weapp/payment"

// 新建退款订单
form := payment.TransferInfo{
    AppID:       "APPID",
    MchID:       "商户号",
    OutTradeNo:  "商户订单号", // or TransactionID: "微信订单号",
}

// 需要证书
res, err := form.GetInfo("支付密钥",  "cert 证书路径", "key 证书路径")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 订单查询

[官方文档](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_2)

```go

import "github.com/medivhzhan/weapp/payment"

q := payment.Query{
  AppID:       "APPID",
  MchID:       "商户号",
  OutTradeNo:  "商户订单号",//商户订单号和微信订单号 至少填一个
  TransactionID: "微信订单号",
}
res, err := q.Query("支付密钥")  //只有当res.TradeState == "SUCCESS" 才是支付成功了
if err != nil {
   fmt.Println(err)
   return
}
fmt.Printf("返回结果: %#v", res)

```

### 支付后获取UnionID

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html)

```go

import "github.com/medivhzhan/weapp/payment"


    res, err := payment.GetPaidUnionID("your-weapp-access-token", "user-open-id", "transaction-id")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("返回结果: %#v", res)

    res, err := payment.GetPaidUnionIDWithMCH("your-weapp-access-token", "user-open-id","out-trade-no", "mch-id")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("返回结果: %#v", res)

```

---

## 解密

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#%E5%8A%A0%E5%AF%86%E6%95%B0%E6%8D%AE%E8%A7%A3%E5%AF%86%E7%AE%97%E6%B3%95)

> 请注意: 前端应当先完成**登录流程**再调用获取**加密数据**的相关接口。

### 解密手机号码

```go

import "github.com/medivhzhan/weapp"

// 解密手机号码
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @data 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
phone , err := weapp.DecryptPhoneNumber(ssk, data, iv string)

fmt.Printf("手机数据: %#v", phone)

```

### 解密分享内容

```go

import "github.com/medivhzhan/weapp"

// 解密转发信息的加密数据
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @data 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
//
// @gid 小程序唯一群号
openGid , err := weapp.DecryptShareInfo(ssk, data, iv string)

```

### 解密用户信息

```go

import "github.com/medivhzhan/weapp"

// 解密用户信息
//
// @rawData 不包括敏感信息的原始数据字符串, 用于计算签名。
// @encryptedData 包括敏感数据在内的完整用户信息的加密数据
// @signature 使用 sha1( rawData + session_key ) 得到字符串, 用于校验用户信息
// @iv 加密算法的初始向量
// @ssk 微信 session_key
ui, err := weapp.DecryptUserInfo(rawData, encryptedData, signature, iv, ssk string)
if err != nil {
    return
}

fmt.Printf("用户数据: %#v", ui)

```

---

## 内容检测

### 检测图片

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html)

```go

import "github.com/medivhzhan/weapp"

// 本地图片检测
//
// @filename 要检测的图片本地路径
// @token 接口调用凭证(access_token)
res, err := IMGSecCheck(filename, token string)
if err != nil {
    return
}

fmt.Printf("返回结果: %#v", res)

// 网络图片检测
//
// @url 要检测的图片网络路径
// @token 接口调用凭证(access_token)
res, err := IMGSecCheckFromNet(url, token string)
if err != nil {
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 检测文本

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api/msgSecCheck.html)

```go

import "github.com/medivhzhan/weapp"

// 文本检测
//
// @content 要检测的文本内容，长度不超过 500KB，编码格式为utf-8
// @token 接口调用凭证(access_token)
res, err := MSGSecCheck(content, token string)
if err != nil {
    return
}

fmt.Printf("返回结果: %#v", res)

```

---

## 生物认证秘钥签名验证

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html)

```go

import "github.com/medivhzhan/weapp"

res, err := VerifySignature("access-token", "user_openid", "result-json", "result-json-signature")
if err != nil {
    // handle error
    return
}

// res.IsOk
// res.Errcode
// res.Errmsg
fmt.Printf("返回结果: %#v", res)

```
