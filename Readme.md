# ![微信小程序服务端 SDK (for Golang)](weapp.png)

## `注意`

- v3 版本为测试版本
- [v1 版本入口](https://github.com/medivhzhan/weapp/tree/v1)
- [v2 版本入口](https://github.com/medivhzhan/weapp/tree/v2)
- 新版本不包含支付相关内容, 已有很多优秀的支付相关模块;
- 请使用经过`线上测试` ✅ 的接口。
- 未完成的接口将在经过线上测试后在新版本中提供。
- 大部分接口需要去线上测试, 欢迎一起完善 :)

## 获取最新版本代码

```sh

go get -u github.com/medivhzhan/weapp/v3

```

## `目录`

> 文档按照[小程序服务端官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)排版，方便您一一对照查找相关内容。

✅: 代表已经经过线上测试
❌: 代表还没有经过线上测试或者未完成

- [初始化](#初始化)
- [登录](#登录)
  - [code2Session](#code2Session) ✅
- [用户信息](#用户信息)
  - [getPaidUnionId](#getPaidUnionId) ✅
- [接口调用凭证](#接口调用凭证)
  - [getAccessToken](#getAccessToken) ✅
- [数据分析](#数据分析)
  - [访问留存](#访问留存)
    - [getDailyRetain](#getDailyRetain) ✅
    - [getWeeklyRetain](#getWeeklyRetain) ✅
    - [getMonthlyRetain](#getMonthlyRetain) ✅
  - [getDailySummary](#getDailySummary) ✅
  - [访问趋势](#访问趋势)
    - [getDailyVisitTrend](#getDailyVisitTrend) ✅
    - [getWeeklyVisitTrend](#getWeeklyVisitTrend) ✅
    - [getMonthlyVisitTrend](#getMonthlyVisitTrend) ✅
  - [getUserPortrait](#getUserPortrait) ✅
  - [getVisitDistribution](#getVisitDistribution) ✅
  - [getVisitPage](#getVisitPage) ✅
- [客服消息](#客服消息)
  - [getTempMedia](#getTempMedia) ✅
  - [sendCustomerServiceMessage](#sendCustomerServiceMessage) ✅
  - [setTyping](#setTyping) ✅
  - [uploadTempMedia](#uploadTempMedia) ✅
- [统一服务消息](#统一服务消息)
  - [sendUniformMessage](#sendUniformMessage) ✅
- [动态消息](#动态消息)
  - [createActivityId](#createActivityId)
  - [setUpdatableMsg](#setUpdatableMsg)
- [插件管理](#插件管理)
  - [applyPlugin](#applyPlugin)
  - [getPluginDevApplyList](#getPluginDevApplyList)
  - [getPluginList](#getPluginList)
  - [setDevPluginApplyStatus](#setDevPluginApplyStatus)
  - [unbindPlugin](#unbindPlugin)
- [附近的小程序](#附近的小程序)
  - [addNearbyPoi](#addNearbyPoi)
  - [deleteNearbyPoi](#deleteNearbyPoi)
  - [getNearbyPoiList](#getNearbyPoiList)
  - [setNearbyPoiShowStatus](#setNearbyPoiShowStatus)
- [小程序码](#小程序码) ✅
  - [createQRCode](#createQRCode) ✅
  - [get](#get) ✅
  - [getUnlimited](#getUnlimited) ✅
- [URL-Scheme](#URL-Scheme)
  - [generate](#generate) ✅
- [内容安全](#内容安全)
  - [imgSecCheck](#imgSecCheck) ✅
  - [mediaCheckAsync](#mediaCheckAsync)✅
  - [msgSecCheck](#msgSecCheck) ✅
- [图像处理](#图像处理)
  - [aiCrop](#aiCrop) ✅
  - [scanQRCode](#scanQRCode) ✅
  - [superResolution](#superResolution)
- [及时配送](#及时配送)❌
  - [小程序使用](#小程序使用)
    - [abnormalConfirm](#abnormalConfirm)
    - [addDeliveryOrder](#addDeliveryOrder)
    - [addDeliveryTip](#addDeliveryTip)
    - [cancelDeliveryOrder](#cancelDeliveryOrder)
    - [getAllImmediateDelivery](#getAllImmediateDelivery)
    - [getBindAccount](#getBindAccount)
    - [getDeliveryOrder](#getDeliveryOrder)
    - [mockUpdateDeliveryOrder](#mockUpdateDeliveryOrder)
    - [onDeliveryOrderStatus](#onDeliveryOrderStatus)
    - [preAddDeliveryOrder](#preAddDeliveryOrder)
    - [preCancelDeliveryOrder](#preCancelDeliveryOrder)
    - [reDeliveryOrder](#reDeliveryOrder)
  - [服务提供方使用](#服务提供方使用)
    - [updateDeliveryOrder](#updateDeliveryOrder)
    - [onAgentPosQuery](#onAgentPosQuery)
    - [onAuthInfoGet](#onAuthInfoGet)
    - [onCancelAuth](#onCancelAuth)
    - [onDeliveryOrderAdd](#onDeliveryOrderAdd)
    - [onDeliveryOrderAddTips](#onDeliveryOrderAddTips)
    - [onDeliveryOrderCancel](#onDeliveryOrderCancel)
    - [onDeliveryOrderConfirmReturn](#onDeliveryOrderConfirmReturn)
    - [onDeliveryOrderPreAdd](#onDeliveryOrderPreAdd)
    - [onDeliveryOrderPreCancel](#onDeliveryOrderPreCancel)
    - [onDeliveryOrderQuery](#onDeliveryOrderQuery)
    - [onDeliveryOrderReAdd](#onDeliveryOrderReAdd)
    - [onPreAuthCodeGet](#onPreAuthCodeGet)
    - [onRiderScoreSet](#onRiderScoreSet)
- [物流助手](#物流助手)❌
  - [小程序使用](#小程序使用)
    - [addExpressOrder](#addExpressOrder)
    - [cancelExpressOrder](#cancelExpressOrder)
    - [getAllDelivery](#getAllDelivery)
    - [getExpressOrder](#getExpressOrder)
    - [getExpressPath](#getExpressPath)
    - [getExpressPrinter](#getExpressPrinter)
    - [getExpressQuota](#getExpressQuota)
    - [onExpressPathUpdate](#onExpressPathUpdate)
    - [testUpdateExpressOrder](#testUpdateExpressOrder)
    - [updateExpressPrinter](#updateExpressPrinter)
  - [服务提供方使用](#服务提供方使用)
    - [getExpressContact](#getExpressContact)
    - [onAddExpressOrder](#onAddExpressOrder)
    - [onCancelExpressOrder](#onCancelExpressOrder)
    - [onCheckExpressBusiness](#onCheckExpressBusiness)
    - [onGetExpressQuota](#onGetExpressQuota)
    - [previewExpressTemplate](#previewExpressTemplate)
    - [updateExpressBusiness](#updateExpressBusiness)
    - [updateExpressPath](#updateExpressPath)
- [OCR](#OCR)
  - [bankcard](#bankcard) ✅
  - [businessLicense](#businessLicense) ✅
  - [driverLicense](#driverLicense) ✅
  - [idcard](#idcard) ✅
  - [printedText](#printedText) ✅
  - [vehicleLicense](#vehicleLicense) ✅
- [运维中心](#运维中心)❌
  - [realTimeLogSearch](#realTimeLogSearch)
- [小程序搜索](#小程序搜索)❌
  - [siteSearch](#siteSearch)
  - [submitPages](#submitPages)
- [生物认证](#生物认证)
  - [verifySignature](#verifySignature)
- [订阅消息](#订阅消息) ✅
  - [addTemplate](#addTemplate) ✅
  - [deleteTemplate](#deleteTemplate) ✅
  - [getCategory](#getCategory) ✅
  - [getPubTemplateKeyWordsById](#getPubTemplateKeyWordsById)✅
  - [getPubTemplateTitleList](#getPubTemplateTitleList) ✅
  - [getTemplateList](#getTemplateList) ✅
  - [sendSubscribeMessage](#sendSubscribeMessage) ✅
- [解密](#解密)
  - [解密手机号码](#解密手机号码) ✅
  - [解密分享内容](#解密分享内容)
  - [解密用户信息](#解密用户信息) ✅
  - [解密微信运动](#解密微信运动)
- [人脸识别](#人脸识别)

---

## 初始化

1. 初始化接口客户端

```go
import "github.com/medivhzhan/weapp/v3"

cli := weapp.NewClient("your-appid", "your-secret")

// 自定义缓存
// 默认缓存在内存中
// 用于缓存 AccessToken
// 只需要实现 cache 模块下 的 Cache 接口即可使用
cc := MyCache{
    //
}

cli := weapp.NewClient(
    "your-appid",
    "your-secret",
    weapp.WithCache(cc),
)

// 自定义 HTTP 请求客户端
// 默认为 http.DefaultClient
httpCli := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{
        // 跳过校验
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    },
}

cli := weapp.NewClient(
    "your-appid",
    "your-secret",
    weapp.WithHttpClient(httpCli),
)

```

1. 初始化微信通知服务

```go
import "github.com/medivhzhan/weapp/v3/server"

//  通用处理器
handler := func(req map[string]interface{}) map[string]interface{}{


    switch req["MsgType"] {

        case "text":
        // Do something cool ...
    }

    return nil
}

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, handler)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

err:= srv.Serve()
if err != nil {
    lof.Fatalf("serving error: %s", err)
}

```

## 登录

### code2Session

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.Login("code")
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetPaidUnionID("open-id", "transaction-id")
// 或者
res, err := cli.GetPaidUnionIDWithMCH("open-id", "out-trade-number", "mch-id")

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

> SDK 默认会自动获取并缓存 AccessToken 不建议直接调用

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetAccessToken()
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

#### getDailyRetain

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetDailyRetain("begin-date", "end-date")
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetWeeklyRetain("begin-date", "end-date")
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

#### getMonthlyRetain

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetMonthlyRetain("begin-date", "end-date")
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetDailySummary("begin-date", "end-date")
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

#### getDailyVisitTrend

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetDailyVisitTrend("begin-date", "end-date")
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetWeeklyVisitTrend("begin-date", "end-date")
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

#### getMonthlyVisitTrend

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetMonthlyVisitTrend("begin-date", "end-date")
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetUserPortrait("begin-date", "end-date")
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetVisitDistribution("begin-date", "end-date")
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetVisitPage("begin-date", "end-date")
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

### getTempMedia

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html)

```go

import "github.com/medivhzhan/weapp/v3"

resp, res, err := cli.GetTempMedia("media-id")
if err != nil {
    // 处理一般错误信息
    return
}
defer resp.Close()

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

```

### sendCustomerServiceMessage

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

// 文本消息
srv.OnCustomerServiceTextMessage(func(msg *weapp.TextMessageResult) *weapp.TransferCustomerMessage {

    reply := cli.CSMsgText{
        Content: "content",
    }

    res, err := cli.SendTextMsg(msg.FromUserName, &reply)
    if err != nil {
        // 处理一般错误信息
        return nil
    }

    if err := res.GetResponseError(); err !=nil {
        // 处理微信返回错误信息
        return nil
    }

    return nil
})

// 图片消息
srv.OnCustomerServiceImageMessage(func(msg *weapp.TextMessageResult) *weapp.TransferCustomerMessage {

    reply := cli.CSMsgImage{
        MediaID: "media-id",
    }

    res, err := cli.SendImageMsg(msg.FromUserName, &reply)
    if err != nil {
        // 处理一般错误信息
        return nil
    }

    if err := res.GetResponseError(); err !=nil {
        // 处理微信返回错误信息
        return nil
    }

    return nil
})

// 小程序卡片消息
srv.OnCustomerServiceCardMessage(func(msg *weapp.TextMessageResult) *weapp.TransferCustomerMessage {

    reply := cli.CSMsgMPCard{
        Title:        "title",
        PagePath:     "page-path",
        ThumbMediaID: "thumb-media-id",
    }
    res, err := cli.SendCardMsg(msg.FromUserName, &reply)
    if err != nil {
        // 处理一般错误信息
        return nil
    }

    if err := res.GetResponseError(); err !=nil {
        // 处理微信返回错误信息
        return nil
    }

    return nil
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

### setTyping

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.SetTyping("open-id",cli.SetTypingCommandTyping)
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

import "github.com/medivhzhan/weapp/v3"

res, err := cli.UploadTempMedia(cli.TempMediaTypeImage, "media-filename")
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

## 统一服务消息

### sendUniformMessage

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html)

```go

import "github.com/medivhzhan/weapp/v3"

sender := cli.UniformMsgSender{
    ToUser: "open-id",
    UniformWeappTmpMsg:cli.UniformWeappTmpMsg{
        TemplateID: "template-id",
        Page:       "page",
        FormID:     "form-id",
        Data:cli.UniformMsgData{
            "keyword": {Value: "value"},
        },
        EmphasisKeyword: "keyword.DATA",
    },
    UniformMpTmpMsg:cli.UniformMpTmpMsg{
        AppID:       "app-id",
        TemplateID:  "template-id",
        URL:         "url",
        Miniprogram:cli.UniformMsgMiniprogram{"miniprogram-app-id", "page-path"},
        Data:cli.UniformMsgData{
            "keyword": {"value", "color"},
        },
    },
}

res, err := cli.SendUniformMsg(&sender)
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

## 动态消息

### createActivityId

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.CreateActivityId()
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

### setUpdatableMsg

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html)

```go

import "github.com/medivhzhan/weapp/v3"


setter := cli.UpdatableMsg{
    "activity-id",
    UpdatableMsgJoining,
    UpdatableMsgTempInfo{
        []UpdatableMsgParameter{
            {UpdatableMsgParamMemberCount, "parameter-value-number"},
            {UpdatableMsgParamRoomLimit, "parameter-value-number"},
        },
    },
}

res, err := setter.SetUpdatableMsg(&setter)
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

## 插件管理

### applyPlugin

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.ApplyPlugin("plugin-app-id", "reason")
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

### getPluginDevApplyList

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetPluginDevApplyList(1, 2)
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

### getPluginList

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginList.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetPluginList()
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

### setDevPluginApplyStatus

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.setDevPluginApplyStatus.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.SetDevPluginApplyStatus("plugin-app-id", "reason",cli.DevAgree)
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

### unbindPlugin

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.unbindPlugin.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.UnbindPlugin("plugin-app-id")
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

## 附近的小程序

### addNearbyPoi

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html)

```go

import "github.com/medivhzhan/weapp/v3"

poi := NearbyPoi{
    PicList: PicList{[]string{"first-picture-url", "second-picture-url", "third-picture-url"}},
    ServiceInfos:cli.ServiceInfos{[]weapp.ServiceInfo{
        {1, 1, "name", "app-id", "path"},
    }},
    StoreName:         "store-name",
    Hour:              "11:11-12:12",
    Credential:        "credential",
    Address:           "address",                         // 地址 必填
    CompanyName:       "company-name",                    // 主体名字 必填
    QualificationList: "qualification-list",              // 证明材料 必填 如果company_name和该小程序主体不一致，需要填qualification_list，详细规则见附近的小程序使用指南-如何证明门店的经营主体跟公众号或小程序帐号主体相关http://kf.qq.com/faq/170401MbUnim17040122m2qY.html
    KFInfo:           cli.KFInfo{true, "kf-head-img", "kf-name"}, // 客服信息 选填，可自定义服务头像与昵称，具体填写字段见下方示例kf_info pic_list是字符串，内容是一个json！
    PoiID:             "poi-id",                          // 如果创建新的门店，poi_id字段为空 如果更新门店，poi_id参数则填对应门店的poi_id 选填
}

res, err := cli.AddNearByPoi(&poi)
if err != nil {
    // 处理一般错误信息
    return
}

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnAddNearbyPoi(func(mix *weapp.AddNearbyPoiResult) {
    // 处理接收的数据
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 服务出错
    return
}

```

### deleteNearbyPoi

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.DeleteNearbyPoi("poi-id")
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

### getNearbyPoiList

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.getList.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetNearbyPoiList(1, 10)
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

### setNearbyPoiShowStatus

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.setShowStatus.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.SetNearbyPoiShowStatus("poi-id",cli.ShowNearbyPoi)
// 或者
res, err := cli.SetNearbyPoiShowStatus("poi-id",cli.HideNearbyPoi)
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

## 小程序码

### createQRCode

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html)

```go

import (
    "ioutil"
    "github.com/medivhzhan/weapp/v3"
)


creator := cli.QRCodeCreator{
    Path:  "mock/path",
    Width: 430,
}

resp, res, err := cli.CreateQRCode(&creator)
if err != nil {
    // 处理一般错误信息
    return
}

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}
defer resp.Body.Close()

content, err := ioutil.ReadAll(resp.Body)
// 处理图片内容

```

### get

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html)

```go

import (
    "ioutil"
    "github.com/medivhzhan/weapp/v3"
)


getter := cli.QRCode{
    Path:      "mock/path",
    Width:     430,
    AutoColor: true,
    LineColor:cli.Color{"r", "g", "b"},
    IsHyaline: true,
}

resp, res, err := cli.GetQRCode(&getter)
if err != nil {
    // 处理一般错误信息
    return
}

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}
defer resp.Body.Close()

content, err := ioutil.ReadAll(resp.Body)
// 处理图片内容

```

### getUnlimited

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html)

```go

import (
    "ioutil"
    "github.com/medivhzhan/weapp/v3"
)


getter := cli.UnlimitedQRCode{
    Scene:     "scene-data",
    Page:      "mock/page",
    Width:     430,
    AutoColor: true,
    LineColor:cli.Color{"r", "g", "b"},
    IsHyaline: true,
}

resp, res, err := cli.GetUnlimitedQRCode(&getter)
if err != nil {
    // 处理一般错误信息
    return
}

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}
defer resp.Body.Close()

content, err := ioutil.ReadAll(resp.Body)
// 处理图片内容

```

---

## URL-Scheme

### generate

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/url-scheme.html)

```go

import (
    "ioutil"
    "github.com/medivhzhan/weapp/v3"
)


scheme := cli.URLScheme{
    SchemedInfo:  &weapp.SchemedInfo{
        Path:  "mock/path",
        Query:  "",
    },
    IsExpire: true,
    ExpireTime: 1234567890,
}

resp, res, err := cli.GenerateURLSchema(&scheme)
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

## 内容安全

### imgSecCheck

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.IMGSecCheck("local-filename")
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

### mediaCheckAsync

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.MediaCheckAsync("image-url",cli.MediaTypeImage)
if err != nil {
    // 处理一般错误信息
    return
}

if err := res.GetResponseError(); err !=nil {
    // 处理微信返回错误信息
    return
}

fmt.Printf("返回结果: %#v", res)

// 接收并处理异步结果
srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnMediaCheckAsync(func(mix *weapp.MediaCheckAsyncResult) {
    // 处理返回结果
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

### msgSecCheck

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.MSGSecCheck("message-content")
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

## 图像处理

### aiCrop

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.AICrop("filename")
// 或者
res, err := cli.AICropByURL("url")
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

### scanQRCode

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.ScanQRCode("file-path")
// 或者
res, err := cli.ScanQRCodeByURL("qr-code-url")
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

### superResolution

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.SuperResolution("file-path")
// 或者
res, err := cli.SuperResolutionByURL("img-url")
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

## 及时配送

### 服务提供方使用

#### updateDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.updateOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

mocker := cli.DeliveryOrderUpdater{
   // ...
}

res, err := cli.UpdateImmediateDeliveryOrder(&mocker)
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

#### onAgentPosQuery

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onAgentPosQuery.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnAgentPosQuery(func(mix *weapp.AgentPosQueryResult) *weapp.AgentPosQueryReturn {
    // 处理返回结果

    return &weapp.AgentPosQueryReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onAuthInfoGet

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onAuthInfoGet.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnAuthInfoGet(func(mix *weapp.AuthInfoGetResult) *weapp.AuthInfoGetReturn {
    // 处理返回结果

    return &weapp.AuthInfoGetReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onCancelAuth

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onCancelAuth.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnCancelAuth(func(mix *weapp.CancelAuthResult) *weapp.CancelAuthReturn {
    // 处理返回结果

    return &weapp.CancelAuthReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderAdd

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderAdd.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderAdd(func(mix *weapp.DeliveryOrderAddResult) *weapp.DeliveryOrderAddReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderAddReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderAddTips

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderAddTips.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderAddTips(func(mix *weapp.DeliveryOrderAddTipsResult) *weapp.DeliveryOrderAddTipsReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderAddTipsReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderCancel

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderCancel.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderCancel(func(mix *weapp.DeliveryOrderCancelResult) *weapp.DeliveryOrderCancelReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderCancelReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderConfirmReturn

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderConfirmReturn.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderReturnConfirm(func(mix *weapp.DeliveryOrderReturnConfirmResult) *weapp.DeliveryOrderReturnConfirmReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderReturnConfirmReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderPreAdd

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderPreAdd.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderPreAdd(func(mix *weapp.DeliveryOrderPreAddResult) *weapp.DeliveryOrderPreAddReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderPreAddReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderPreCancel

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderPreCancel.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderPreCancel(func(mix *weapp.DeliveryOrderPreCancelResult) *weapp.DeliveryOrderPreCancelReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderPreCancelReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderQuery

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderQuery.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderQuery(func(mix *weapp.DeliveryOrderQueryResult) *weapp.DeliveryOrderQueryReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderQueryReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onDeliveryOrderReAdd

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onOrderReAdd.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderReadd(func(mix *weapp.DeliveryOrderReaddResult) *weapp.DeliveryOrderReaddReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderReaddReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onPreAuthCodeGet

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onPreAuthCodeGet.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnPreAuthCodeGet(func(mix *weapp.PreAuthCodeGetResult) *weapp.PreAuthCodeGetReturn {
    // 处理返回结果

    return &weapp.PreAuthCodeGetReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onRiderScoreSet

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-provider/immediateDelivery.onRiderScoreSet.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnRiderScoreSet(func(mix *weapp.RiderScoreSetResult) *weapp.RiderScoreSetReturn {
    // 处理返回结果

    return &weapp.PreAuthCodeGetReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

### 小程序使用

#### abnormalConfirm

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.abnormalConfirm.html)

```go

import "github.com/medivhzhan/weapp/v3"

confirmer := cli.AbnormalConfirmer{
    ShopID:       "123456",
    ShopOrderID:  "123456",
    ShopNo:       "shop_no_111",
    WaybillID:    "123456",
    Remark:       "remark",
    DeliverySign: "123456",
}

res, err := cli.AbnormalImmediateDeliveryConfirm(&confirmer)
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

#### addDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

creator := cli.DeliveryOrderCreator{
   // ...
}

res, err := cli.AddImmediateDeliveryOrder(&creator)
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

#### addDeliveryTip

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.addTip.html)

```go

import "github.com/medivhzhan/weapp/v3"

adder := cli.DeliveryTipAdder{
   // ...
}

res, err := cli.AddImmediateDeliveryTip(&addr)
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

#### cancelDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.cancelOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

canceler := cli.DeliveryOrderCanceler{
   // ...
}

res, err := cli.PreCancelImmediateDeliveryOrder(&canceler)
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

#### getAllImmediateDelivery

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getAllImmeDelivery.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetAllImmediateDelivery()
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

#### getBindAccount

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getBindAccount.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetBindAccount()
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

#### getDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.getOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

getter := cli.DeliveryOrderGetter{
   // ...
}

res, err := cli.GetImmediateDeliveryOrder(getter)
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

#### mockUpdateDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.mockUpdateOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

mocker := cli.UpdateDeliveryOrderMocker{
   // ...
}

res, err := cli.MockUpdateImmediateDeliveryOrder(mocker)
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

#### onDeliveryOrderStatus

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.onOrderStatus.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnDeliveryOrderStatusUpdate(func(mix *weapp.DeliveryOrderStatusUpdateResult) *weapp.DeliveryOrderStatusUpdateReturn {
    // 处理返回结果

    return &weapp.DeliveryOrderStatusUpdateReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### preAddDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preAddOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

creator := cli.DeliveryOrderCreator{
   // ...
}

res, err := cli.PreAddImmediateDeliveryOrder(&creator)
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

#### preCancelDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.preCancelOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

canceler := cli.DeliveryOrderCanceler{
   // ...
}

res, err := cli.PreCancelImmediateDeliveryOrder(&canceler)
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

#### reDeliveryOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/immediate-delivery/by-business/immediateDelivery.reOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

creator := cli.DeliveryOrderCreator{
   // ...
}

res, err := cli.ReImmediateDeliveryOrder(&creator)
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

## 物流助手

### 小程序使用

#### addExpressOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.addOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

creator := cli.ExpressOrderCreator{
   // ...
}

res, err := cli.AddLogisticOrder(&creator)
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

#### cancelExpressOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.cancelOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

canceler := cli.ExpressOrderCanceler{
   // ...
}

res, err := cli.CancelLogisticsOrder(&canceler)
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

#### getAllDelivery

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getAllDelivery.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.getAllDelivery()
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

#### getExpressOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

getter := cli.ExpressOrderGetter{
   // ...
}

res, err := cli.GetLogisticsOrder(&getter)
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

#### getExpressPath

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPath.html)

```go

import "github.com/medivhzhan/weapp/v3"

getter := cli.ExpressPathGetter{
   // ...
}

res, err := cli.GetLogisticsPath(&getter)
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

#### getExpressPrinter

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getPrinter.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetPrinter()
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

#### getExpressQuota

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.getQuota.html)

```go

import "github.com/medivhzhan/weapp/v3"

getter := cli.QuotaGetter{
   // ...
}

res, err := cli.GetExpressQuota(&getter)
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

#### onExpressPathUpdate

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.onPathUpdate.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnExpressPathUpdate(func(mix *weapp.ExpressPathUpdateResult) {
    // 处理返回结果
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### testUpdateExpressOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.testUpdateOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

tester := cli.UpdateExpressOrderTester{
   // ...
}

res, err := cli.TestUpdateExpressOrder(&tester)
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

#### updateExpressPrinter

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-business/logistics.updatePrinter.html)

```go

import "github.com/medivhzhan/weapp/v3"

updater := cli.PrinterUpdater{
   // ...
}

res, err := cli.updateExpressOrder(&updater)
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

### 服务提供方使用

#### getExpressContact

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.getContact.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.GetContact("token", "wat-bill-id")
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

#### onAddExpressOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.onAddOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnAddExpressOrder(func(mix *weapp.AddExpressOrderResult) *weapp.AddExpressOrderReturn {
    // 处理返回结果

    return &weapp.AddExpressOrderReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onCancelExpressOrder

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.onCancelOrder.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnCancelExpressOrder(func(mix *weapp.CancelExpressOrderResult) *weapp.CancelExpressOrderReturn {
    // 处理返回结果

    return &weapp.CancelExpressOrderReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onCheckExpressBusiness

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.onCheckBusiness.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnCheckExpressBusiness(func(mix *weapp.CheckExpressBusinessResult) *weapp.CheckExpressBusinessReturn {
    // 处理返回结果

    return &weapp.CheckExpressBusinessReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### onGetExpressQuota

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.onGetQuota.html)

```go

import "github.com/medivhzhan/weapp/v3"

srv, err := server.NewServer("appid", "token", "aesKey", "mchID", "apiKey", false, nil)
if err != nil {
    lof.Fatalf("init server error: %s", err)
}

srv.OnGetExpressQuota(func(mix *weapp.GetExpressQuotaResult) *weapp.GetExpressQuotaReturn {
    // 处理返回结果

    return &weapp.GetExpressQuotaReturn{
        // ...
    }
})

if err := srv.Serve(http.ResponseWriter, *http.Request); err != nil {
    // 处理微信返回错误信息
    return
}

```

#### previewExpressTemplate

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.previewTemplate.html)

```go

import "github.com/medivhzhan/weapp/v3"

previewer := cli.ExpressTemplatePreviewer{
   // ...
}

res, err := cli.PreviewLogisticsTemplate(&previewer)
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

#### updateExpressBusiness

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updateBusiness.html)

```go

import "github.com/medivhzhan/weapp/v3"

updater := cli.BusinessUpdater{
   // ...
}

res, err := cli.UpdateLogisticsBusiness(&updater)
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

#### updateExpressPath

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/express/by-provider/logistics.updatePath.html)

```go

import "github.com/medivhzhan/weapp/v3"

updater := cli.ExpressPathUpdater{
   // ...
}

res, err := cli.UpdateLogisticsPath(&updater)
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

## OCR

### bankcard

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.BankCard("file-path",cli.RecognizeModeScan)
// 或者
res, err := cli.BankCardByURL("card-url",cli.RecognizeModePhoto)
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

### businessLicense

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.BusinessLicense("file-path")
// 或者
res, err := cli.BusinessLicenseByURL("card-url")
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

### driverLicense

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.DriverLicense("file-path")
// 或者
res, err := cli.DriverLicenseByURL("card-url")
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

### idcard

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.IDCardByURL("card-url",cli.RecognizeModePhoto)
// 或者
res, err := cli.IDCard("file-path",cli.RecognizeModeScan)
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

### printedText

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.PrintedText("file-path")
// 或者
res, err := cli.PrintedTextByURL("card-url")
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

### vehicleLicense

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.VehicleLicense("file-path",cli.RecognizeModeScan)
// 或者
res, err := cli.VehicleLicenseByURL("card-url",cli.RecognizeModePhoto)
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

## 小程序搜索

### siteSearch

### submitPages

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.submitPages.html)

```go

import "github.com/medivhzhan/weapp/v3"

sender := weapp.SearchSubmitPages{
    []weapp.SearchSubmitPage{
        {
            Path:  "pages/index/index",
            Query: "id=test",
        },
    },
}

res, err := cli.SendSearchSubmitPages(&sender)
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

## 生物认证

### verifySignature

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html)

```go

import "github.com/medivhzhan/weapp/v3"

res, err := cli.VerifySignature("open-id", "data", "signature")
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

## 订阅消息

### addTemplate

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html)

```go
import "github.com/medivhzhan/weapp/v3"

// AddTemplate 组合模板并添加至帐号下的个人模板库
//
// tid 模板ID
// desc 服务场景描述，15个字以内
// keywordIDList 关键词 ID 列表
res, err := cli.AddTemplate("tid", "desc", []int32{1, 2, 3})
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

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html)

```go
import "github.com/medivhzhan/weapp/v3"

// DeleteTemplate 删除帐号下的某个模板
//
// pid 模板ID
res, err := cli.DeleteTemplate("pid")
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

### getCategory

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html)

```go
import "github.com/medivhzhan/weapp/v3"

// GetTemplateCategory 删除帐号下的某个模板
//
res, err := cli.GetTemplateCategory()
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

### getPubTemplateKeyWordsById

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html)

```go
import "github.com/medivhzhan/weapp/v3"

// GetPubTemplateKeyWordsById 获取模板标题下的关键词列表
//
// tid 模板ID
res, err := cli.GetPubTemplateKeyWordsById("tid")
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

### getPubTemplateTitleList

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html)

```go
import "github.com/medivhzhan/weapp/v3"

// GetPubTemplateTitleList 获取帐号所属类目下的公共模板标题
//
// ids 类目 id，多个用逗号隔开
// start 用于分页，表示从 start 开始。从 0 开始计数。
// limit 用于分页，表示拉取 limit 条记录。最大为 30
res, err := cli.GetPubTemplateTitleList("1,2,3", 0, 10)
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

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html)

```go
import "github.com/medivhzhan/weapp/v3"

// GetTemplateList 获取帐号下已存在的模板列表
//
res, err := cli.GetTemplateList()
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

### sendSubscribeMessage

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html)

```go

import "github.com/medivhzhan/weapp/v3"

sender := weapp.SubscribeMessage{
    ToUser:     mpOpenID,
    TemplateID: "template-id",
    Page:       "mock/page/path",
    MiniprogramState:cli.MiniprogramStateDeveloper, // 或者: "developer"
    Data:cli.SubscribeMessageData{
        "first-key": {
            Value: "value",
        },
        "second-key": {
            Value: "value",
        },
    },
}

res, err := cli.SendSubscribeMsg(&sender)
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

## 解密

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)

> ❌ 前端应当先完成[登录](#登录)流程再调用获取加密数据的相关接口。

### 解密手机号码

```go
import "github.com/medivhzhan/weapp/v3"

res, err := cli.DecryptMobile("session-key", "encrypted-data", "iv" )
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)
```

### 解密分享内容

```go
import "github.com/medivhzhan/weapp/v3"

res, err := cli.DecryptShareInfo("session-key", "encrypted-data", "iv" )
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)
```

### 解密用户信息

```go
import "github.com/medivhzhan/weapp/v3"

res, err := cli.DecryptUserInfo( "session-key", "raw-data", "encrypted-data", "signature", "iv")
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)
```

### 解密微信运动

```go
import "github.com/medivhzhan/weapp/v3"

res, err := cli.DecryptRunData("session-key", "encrypted-data", "iv" )
if err != nil {
    // 处理一般错误信息
    return
}

fmt.Printf("返回结果: %#v", res)
```

---

## 人脸识别

```go
import "github.com/medivhzhan/weapp/v3"

// FaceIdentify 获取人脸识别结果
//
// key 小程序 verify_result
res, err := cli.FaceIdentify("verify_result")
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
