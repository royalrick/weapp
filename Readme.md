# ![title](title.png)

## `注意`

- [1.x 版本入口](https://github.com/medivhzhan/weapp/tree/v1)
- 2.0 版本开始支付相关内容会分离到一个单独的包。
- 文档按照[小程序服务端官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)排版，方便您一一对照查找相关内容。
- 目前 v2 还在测试阶段，请勿用于生产环境。

## 获取代码

```sh

go get -u github.com/medivhzhan/weapp

```

## `目录`

- [用户登录](#用户登录)
  - [code2Session](#code2Session)
- [用户信息](#用户信息)
  - [getPaidUnionId](#getPaidUnionId)
- [接口调用凭证](#接口调用凭证)
  - [getAccessToken](#getAccessToken)
- [数据分析](#数据分析)
  - [访问留存](#访问留存)
    - [getMonthlyRetain](#getMonthlyRetain)
    - [getWeeklyRetain](#getWeeklyRetain)
    - [getDailyRetain](#getDailyRetain)
  - [访问趋势](#访问趋势)
    - [getMonthlyVisitTrend](#getMonthlyVisitTrend)
    - [getWeeklyVisitTrend](#getWeeklyVisitTrend)
    - [getDailyVisitTrend](#getDailyVisitTrend)
  - [getUserPortrait](#getUserPortrait)
  - [getVisitDistribution](#getVisitDistribution)
  - [getVisitPage](#getVisitPage)
  - [getDailySummary](#getDailySummary)
- [客服消息](#客服消息)
  - [setTyping](#setTyping)
  - [uploadTempMedia](#uploadTempMedia)
  - [getTempMedia](#getTempMedia)
  - [customerServiceMessage.send](#customerServiceMessage.send)
- [模板消息](#模板消息)
  - [addTemplate](#addTemplate)
  - [deleteTemplate](#deleteTemplate)
  - [getTemplateLibraryById](#getTemplateLibraryById)
  - [getTemplateLibraryList](#getTemplateLibraryList)
  - [getTemplateList](#getTemplateList)
  - [templateMessage.send](#templateMessage.send)

---

## 用户登录

### code2Session

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

## 用户信息

### getPaidUnionId

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetPaidUnionID("mock-access-token", "mock-open-id", "mock-transaction-id")
// 或者
res, err := weapp.GetPaidUnionIDWithMCH("mock-access-token", "mock-open-id", "mock-out-trade-number", "mock-mch-id")

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

## 接口调用凭证

### getAccessToken

> 调用次数有限制 请注意缓存

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetAccessToken("mock-appid", "mock-secret")
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

## 数据分析

### 访问留存

#### getMonthlyRetain

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetMonthlyRetain("mock-access-token", "mock-begin-date", "mock-end-date")
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

#### getWeeklyRetain

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetWeeklyRetain("mock-access-token", "mock-begin-date", "mock-end-date")
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

#### getDailyRetain

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetDailyRetain("mock-access-token", "mock-begin-date", "mock-end-date")
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

### 访问趋势

#### getMonthlyVisitTrend

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetMonthlyVisitTrend("mock-access-token", "mock-begin-date", "mock-end-date")
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

#### getWeeklyVisitTrend

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetWeeklyVisitTrend("mock-access-token", "mock-begin-date", "mock-end-date")
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

#### getDailyVisitTrend

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetDailyVisitTrend("mock-access-token", "mock-begin-date", "mock-end-date")
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

### getUserPortrait

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetUserPortrait("mock-access-token", "mock-begin-date", "mock-end-date")
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

### getVisitDistribution

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetVisitDistribution("mock-access-token", "mock-begin-date", "mock-end-date")
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

### getVisitPage

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetVisitPage("mock-access-token", "mock-begin-date", "mock-end-date")
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

### getDailySummary

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetDailySummary("mock-access-token", "mock-begin-date", "mock-end-date")
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

## 客服消息

### setTyping

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.SetTyping("mock-access-token", "mock-open-id", "mock-command")
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

### uploadTempMedia

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.UploadTempMedia("mock-access-token", "mock-media-type", "mock-media-filename")
// 或者
res, err := weapp.uploadTempMediaByURL("mock-access-token", "mock-media-type", "mock-media-url") // 通过 URL 上传
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

### getTempMedia

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetTempMedia("mock-access-token", "mock-media-id")
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

### customerServiceMessage.send

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html)

```go

import "github.com/medivhzhan/weapp"


// 文本消息
msg := CSMsgText{
    Content: "mock-content",
},
// 或者
// 图片消息
msg := CSMsgImage{
    MediaID: "mock-media-id",
},
// 或者
// 链接消息
msg := CSMsgLink{
    Title:       "mock-title",
    Description: "mock-description",
    URL:         "mock-url",
    ThumbURL:    "mock-thumb-url",
},
// 或者
// 小程序卡片消息
msg := CSMsgMPCard{
    Title:        "mock-title",
    PagePath:     "mock-page-path",
    ThumbMediaID: "mock-thumb-media-id",
},

res, err := msg.SendTo("mock-open-id", "mock-access-token")
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

## 模板消息

### addTemplate

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.addTemplate.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.AddTemplate("mock-template-id", "mock-access-token")
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

### deleteTemplate

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.deleteTemplate.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.DeleteTemplate("mock-template-id", "mock-access-token")
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

### getTemplateLibraryById

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateLibraryById.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetTemplateLibraryByID("mock-template-id", "mock-access-token")
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

### getTemplateLibraryList

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateLibraryList.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetTemplateLibraryList("mock-template-token", 0, 10)
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

### getTemplateList

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.getTemplateList.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.GetTemplateList("mock-template-token", 0, 10)
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

### send

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/template-message/templateMessage.send.html)

```go

import "github.com/medivhzhan/weapp"

sender := TempMsgSender{
    ToUser:     "mock-open-id",
    TemplateID: "mock-template-id",
    Page:       "mock/page?foo=bar",
    FormID:     "mock-form-id",
    Data: TempMsgData{
        "mock-key-word1": TempMsgKeyword{"mock-template-value1"},
        "mock-key-word2": TempMsgKeyword{"mock-template-value2"},
    },
    EmphasisKeyword: "mock-open-id",
}

res, err := sender.Send("mock-template-token")
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
