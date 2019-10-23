# ![title](title.png)

## `注意⚠️`

- [1.x 版本入口](https://github.com/medivhzhan/weapp/tree/v1)
- 2.0 版本开始支付相关内容会分离到一个单独的包。
- 目前 v2 还在测试阶段，请勿用于生产环境。

## 获取代码

```sh

go get -u github.com/medivhzhan/weapp

```

## `目录`

> 文档按照[小程序服务端官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)排版，方便您一一对照查找相关内容。

✅：代表已经通过线上测试
⚠️：代表还没有或者未完成

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
- [模板消息](#模板消息)(腾讯将于 2020 年 1 月 10 日下线该接口，请使用`订阅消息`))
- [统一服务消息](#统一服务消息)
  - [sendUniformMessage](#sendUniformMessage)
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
- [内容安全](#内容安全)
  - [imgSecCheck](#imgSecCheck) ✅
  - [mediaCheckAsync](#mediaCheckAsync)
  - [msgSecCheck](#msgSecCheck) ✅
- [图像处理](#图像处理)
  - [aiCrop](#aiCrop)⚠️
  - [scanQRCode](#scanQRCode)
  - [superResolution](#superResolution)
- [及时配送](#及时配送)⚠️
- [物流助手](#物流助手)⚠️
- [OCR](#OCR)
  - [bankcard](#bankcard)
  - [businessLicense](#businessLicense)⚠️
  - [driverLicense](#driverLicense)
  - [idcard](#idcard)
  - [printedText](#printedText)⚠️
  - [vehicleLicense](#vehicleLicense)
- [生物认证](#生物认证)
  - [verifySignature](#verifySignature)
- [订阅消息](#订阅消息)⚠️

---

## 登录

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

### 访问趋势

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

---

## 客服消息

### getTempMedia

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.getTempMedia.html)

```go

import "github.com/medivhzhan/weapp"

resp, res, err := weapp.GetTempMedia("mock-access-token", "mock-media-id")
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

import "github.com/medivhzhan/weapp"


// 文本消息
msg := weapp.CSMsgText{
    Content: "mock-content",
}
// 或者
// 图片消息
msg := weapp.CSMsgImage{
    MediaID: "mock-media-id",
}
// 或者
// 链接消息
msg := weapp.CSMsgLink{
    Title:       "mock-title",
    Description: "mock-description",
    URL:         "mock-url",
    ThumbURL:    "mock-thumb-url",
}
// 或者
// 小程序卡片消息
msg := weapp.CSMsgMPCard{
    Title:        "mock-title",
    PagePath:     "mock-page-path",
    ThumbMediaID: "mock-thumb-media-id",
}

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

### setTyping

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.SetTyping("mock-access-token", "mock-open-id", weapp.SetTypingCommandTyping)
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

res, err := weapp.UploadTempMedia("mock-access-token", weapp.TempMediaTypeImage, "mock-media-filename")
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

import "github.com/medivhzhan/weapp"

sender := weapp.UniformMsgSender{
    ToUser: "mock-open-id",
    UniformMPMsg: UniformMPMsg{
        TemplateID: "mock-template-id",
        Page:       "mock-page",
        FormID:     "mock-form-id",
        Data: UniformMsgData{
            "mock-keyword": UniformMsgKeyword{Value: "mock-value"},
        },
        EmphasisKeyword: "mock-keyword.DATA",
    },
    UniformOAMsg: UniformOAMsg{
        AppID:       "mock-app-id",
        TemplateID:  "mock-template-id",
        URL:         "mock-url",
        Miniprogram: UniformMsgMiniprogram{"mock-miniprogram-app-id", "mock-page-path"},
        Data: UniformMsgData{
            "mock-keyword": UniformMsgKeyword{"mock-value", "mock-color"},
        },
    },
}

_, err := sender.Send("mock-access-token")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.CreateActivityId("mock-access-token")
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

import "github.com/medivhzhan/weapp"


setter := weapp.UpdatableMsgSetter{
    "mock-activity-id",
    UpdatableMsgJoining,
    UpdatableMsgTempInfo{
        []UpdatableMsgParameter{
            {UpdatableMsgParamMemberCount, "mock-parameter-value-number"},
            {UpdatableMsgParamRoomLimit, "mock-parameter-value-number"},
        },
    },
}

res, err := setter.Set("mock-access-token")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.ApplyPlugin("access-token", "plugin-app-id", "reason")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.GetPluginDevApplyList("mock-access-token", 1, 2)
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

import "github.com/medivhzhan/weapp"

res, err := weapp.GetPluginList("access-token")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.SetDevPluginApplyStatus("access-token", "plugin-app-id", "reason", weapp.DevAgree)
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

import "github.com/medivhzhan/weapp"

res, err := weapp.UnbindPlugin("access-token", "plugin-app-id")
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

import "github.com/medivhzhan/weapp"

poi := NearbyPoi{
    PicList: PicList{[]string{"first-picture-url", "second-picture-url", "third-picture-url"}},
    ServiceInfos: weapp.ServiceInfos{[]weapp.ServiceInfo{
        {1, 1, "name", "app-id", "path"},
    }},
    StoreName:         "store-name",
    Hour:              "11:11-12:12",
    Credential:        "credential",
    Address:           "address",                         // 地址 必填
    CompanyName:       "company-name",                    // 主体名字 必填
    QualificationList: "qualification-list",              // 证明材料 必填 如果company_name和该小程序主体不一致，需要填qualification_list，详细规则见附近的小程序使用指南-如何证明门店的经营主体跟公众号或小程序帐号主体相关http://kf.qq.com/faq/170401MbUnim17040122m2qY.html
    KFInfo:            weapp.KFInfo{true, "kf-head-img", "kf-name"}, // 客服信息 选填，可自定义服务头像与昵称，具体填写字段见下方示例kf_info pic_list是字符串，内容是一个json！
    PoiID:             "poi-id",                          // 如果创建新的门店，poi_id字段为空 如果更新门店，poi_id参数则填对应门店的poi_id 选填
}

res, err := poi.Add("access-token")
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
srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false, func(mix *Mixture) bool {
    if mix.MsgType != weapp.MsgEvent {
        if mix.Event != weapp.EventNearbyPoiAuditInfoAdd {
            if mix.AuditID == res.Data.AuditID {

                fmt.Printf("返回结果: %#v", mix)
                return true
            }
        }
    }

    return false
})
if err != nil {
     // 处理错误
    return
}

if err := srv.HandleRequest(w, r); err != nil {
     // 处理错误
    return
}

```

### deleteNearbyPoi

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.DeleteNearbyPoi("access-token", "poi-id")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.GetNearbyPoiList("access-token", 1, 10)
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

import "github.com/medivhzhan/weapp"

res, err := weapp.SetNearbyPoiShowStatus("access-token", "poi-id", weapp.ShowNearbyPoi)
// 或者
res, err := weapp.SetNearbyPoiShowStatus("access-token", "poi-id", weapp.HideNearbyPoi)
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
    "github.com/medivhzhan/weapp"
)


creator := weapp.QRCodeCreator{
    Path:  "mock/path",
    Width: 430,
}

resp, res, err := creator.Create("mock-access-token")
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
    "github.com/medivhzhan/weapp"
)


getter := weapp.QRCode{
    Path:      "mock/path",
    Width:     430,
    AutoColor: true,
    LineColor: weapp.Color{"r", "g", "b"},
    IsHyaline: true,
}

resp, res, err := getter.Get("mock-access-token")
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
    "github.com/medivhzhan/weapp"
)


getter :=  weapp.UnlimitedQRCode{
    Scene:     "mock-scene-data",
    Page:      "mock/page",
    Width:     430,
    AutoColor: true,
    LineColor: weapp.Color{"r", "g", "b"},
    IsHyaline: true,
}

resp, res, err := getter.Get("mock-access-token")
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

## 内容安全

### imgSecCheck

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.IMGSecCheck("mock-access-token", "local-filename")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.MediaCheckAsync("access-token", "image-url", weapp.MediaTypeImage)
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
srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false, func(mix *Mixture) bool {
    if mix.MsgType != weapp.MsgEvent {
        if mix.Event != weapp.EventAsyncMediaCheck {
            if mix.TraceID == res.TraceID {

                fmt.Printf("返回结果: %#v", mix)
                return true
            }
        }
    }

    return false
})
if err != nil {
     // 处理错误
    return
}

if err := srv.HandleRequest(w, r); err != nil {
     // 处理错误
    return
}

```

### msgSecCheck

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.MSGSecCheck("access-token", "message-content")
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

### scanQRCode

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html)

```go

import "github.com/medivhzhan/weapp"

res, err := weapp.scanQRCode("access-token", "file-path")
// 或者
res, err := weapp.scanQRCodeByURL("access-token", "qr-code-url")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.SuperResolution("access-token", "file-path")
// 或者
res, err := weapp.SuperResolutionByURL("access-token", "img-url")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.BankCard("access-token", "file-path", weapp.RecognizeModeScan)
// 或者
res, err := weapp.BankCardByURL("access-token", "card-url", weapp.RecognizeModePhoto)
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

import "github.com/medivhzhan/weapp"

res, err := weapp.DriverLicense("access-token", "file-path")
// 或者
res, err := weapp.DriverLicenseByURL("access-token", "card-url")
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

import "github.com/medivhzhan/weapp"

res, err := weapp.IDCardByURL("access-token", "card-url", weapp.RecognizeModePhoto)
// 或者
res, err := weapp.IDCard("access-token", "file-path", weapp.RecognizeModeScan)
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

import "github.com/medivhzhan/weapp"

res, err := weapp.VehicleLicense("access-token", "file-path", weapp.RecognizeModeScan)
// 或者
res, err := weapp.VehicleLicenseByURL("access-token", "card-url", weapp.RecognizeModePhoto)
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

import "github.com/medivhzhan/weapp"

res, err := weapp.VerifySignature("access-token", "open-id", "data", "signature")
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
