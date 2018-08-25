# 微信小程序 SDK （for Golang）

## 目录

- [拉取代码](#拉取代码)
- [AccessToken](#AccessToken)
- [用户登录](#用户登录)
- [二维码](#二维码)
- [模板消息](#模板消息)
- [客服消息](#客服消息)
- [支付](#支付)
    - [付款](#付款)
    - [处理支付结果通知](#处理支付结果通知)
    - [退款](#退款)
    - [处理退款结果通知](#处理退款结果通知)
    - [转账(企业付款)](#转账(企业付款))
    - [查询转账](#查询转账)
- [微信通知](#微信通知)
- [解密](#解密)
    - [解密手机号码](#解密手机号码)
    - [解密分享内容](#解密分享内容)
    - [解密用户信息](#解密用户信息)

## 拉取代码

```sh

go get -u github.com/medivhzhan/weapp

```

## AccessToken

```go

import "github.com/medivhzhan/weapp/token"

// 获取次数有限制 获取后请缓存
tok, exp, err := token.AccessToken(appID, secret string)

```

## 用户登录

UnionID 只在满足一定条件的情况下返回。具体参看 [UnionID机制说明](https://developers.weixin.qq.com/miniprogram/dev/api/unionID.html)

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

### 获取小程序码: 适用于需要的码数量较少的业务场景

```go

import "github.com/medivhzhan/weapp/code"

coder := code.QRCoder {
    Path: "pages/index?query=1", // 识别二维码后进入小程序的页面链接
    With: 430, // 图片宽度
    IsHyaline: true, // 是否需要透明底色
    AutoColor: true, // 自动配置线条颜色, 如果颜色依然是黑色, 则说明不建议配置主色调
    LineColor: code.Color{ //  AutoColor 为 false 时生效, 使用 rgb 设置颜色 十进制表示
        R: "50",
        G: "50",
        B: "50",
    },
}

// 获取小程序码
// token: 微信access_token
res, err := coder.AppCode(token string)
if err != nil {
    // handle error
}
defer res.Body.Close()

```

### 获取小程序码: 适用于需要的码数量极多的业务场景

```go

import "github.com/medivhzhan/weapp/code"

coder := code.QRCoder {
    Scene: "...", // 参数数据
    Page: "pages/index", // 识别二维码后进入小程序的页面链接
    With: 430, // 图片宽度
    IsHyaline: true, // 是否需要透明底色
    AutoColor: true, // 自动配置线条颜色, 如果颜色依然是黑色, 则说明不建议配置主色调
    LineColor: code.Color{ //  AutoColor 为 false 时生效, 使用 rgb 设置颜色 十进制表示
        R: "50",
        G: "50",
        B: "50",
    },
}

// 获取小程序码
// token: 微信access_token
res, err := coder.UnlimitedAppCode(token string)
if err != nil {
    // handle error
}
defer res.Body.Close()

```

### 获取小程序二维码: 适用于需要的码数量较少的业务场景

```go

import "github.com/medivhzhan/weapp/code"

coder := code.QRCoder {
    Path: "pages/index?query=1", // 识别二维码后进入小程序的页面链接
    With: 430, // 图片宽度
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

## 客服消息

### 回复

```go

import "github.com/medivhzhan/weapp/message"

// 文本消息
msg := message.Text{
    Content: "消息内容",
}

// 图片消息
msg := message.Image{
    MediaID: "微信media_id"
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

---

## 微信通知

### 消息

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

---

## 解密

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

fmt.Printf("用户数据: %#v", ui)

```
