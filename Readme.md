# ![title](title.png)

## 注意

- 2.0 版本开始支付相关内容会分离到一个单独的包。
- 文档按照[小程序服务端官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/)排版，方便您一一对照查找相关内容。
- 目前 v2 还在测试阶段，请勿用于生产环境。

## 目录

- [获取代码](#获取代码)
- [用户登录](#用户登录)
- [用户信息](#用户信息)
- [接口调用凭证](#接口调用凭证)
- [数据分析](#数据分析)
  - [访问留存](#访问留存)
  - [访问趋势](#访问趋势)
  - [用户画像](#用户画像)
  - [分布数据](#分布数据)
  - [访问页面](#访问页面)
  - [访问概况](#访问概况)
- [客服消息](#客服消息)
  - [接收客服消息](#接收客服消息)
  - [发送客服消息](#发送客服消息)
- [模板消息](#模板消息)
  - [获取小程序模板库标题列表](#获取小程序模板库标题列表)
  - [获取帐号下已存在的模板列表](#获取帐号下已存在的模板列表)
  - [获取模板库某个模板标题下关键词库](#获取模板库某个模板标题下关键词库)
  - [组合模板并添加至帐号下的个人模板库](#组合模板并添加至帐号下的个人模板库)
  - [删除帐号下的某个模板](#删除帐号下的某个模板)
  - [发送模板消息](#发送模板消息)
- [统一服务消息](#统一服务消息)
- [动态消息](#动态消息)
  - [创建被分享动态消息](#创建被分享动态消息)
  - [修改被分享的动态消息](#修改被分享的动态消息)
- [附近的小程序](#附近的小程序)
  - [添加地点](#添加地点)
  - [删除地点](#删除地点)
  - [查看地点列表](#查看地点列表)
  - [展示/取消展示附近小程序](#展示/取消展示附近小程序)
- [二维码](#二维码)
  - [获取小程序码](#获取小程序码)
  - [获取小程序二维码](#获取小程序二维码)
- [内容检测](#内容检测)
  - [检测图片](#检测图片)
  - [检测文本](#检测文本)
- [生物认证](#生物认证)
  - [秘钥签名验证](#秘钥签名验证)
- [其他](#其他)
  - [解密手机号码](#解密手机号码)
  - [解密分享内容](#解密分享内容)
  - [解密用户信息](#解密用户信息)

## 获取代码

```sh

go get -u github.com/medivhzhan/weapp

```

---

## 用户登录

### 登录凭证校验

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

```go

import "github.com/medivhzhan/weapp"

// appID 小程序 appID
// secret 小程序的 app secret
// code 小程序登录时获取的 code
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

## 用户信息

### 支付后获取 UnionID

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/user-info/auth.getPaidUnionId.html)

```go

import "github.com/medivhzhan/weapp/payment"

res, err := payment.GetPaidUnionID("access-token", "user-openid", "transaction-id")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Printf("返回结果: %#v", res)

res, err := payment.GetPaidUnionIDWithMCH("access-token", "user-openid","out-trade-no", "mch-id")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Printf("返回结果: %#v", res)

```

---

## 接口调用凭证

### 获取 AccessToken

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html)

```go

import "github.com/medivhzhan/weapp/token"

// 获取次数有限制 获取后请缓存
tok, exp, err := token.AccessToken(appID, secret string)

```

---

## 数据分析

```go

// 引入子包
import "github.com/medivhzhan/weapp/analysis"

```

### 访问留存

#### 月留存

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html)

```go

// GetMonthlyRetain 获取用户访问小程序月留存
// accessToken 接口调用凭证
// start 开始日期，为自然月第一天。格式为 yyyymmdd
// end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
res, err := analysis.GetMonthlyRetain("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

#### 周留存

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html)

```go

// GetWeeklyRetain 获取用户访问小程序周留存
// accessToken 接口调用凭证
// start 开始日期，为自然月第一天。格式为 yyyymmdd
// end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
res, err := analysis.GetWeeklyRetain("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

#### 日留存

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html)

```go

// GetDailyRetainAPI 获取用户访问小程序日留存
// accessToken 接口调用凭证
// start 开始日期，为自然月第一天。格式为 yyyymmdd
// end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
res, err := analysis.GetDailyRetainAPI("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 访问趋势

#### 月趋势

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html)

```go

// GetMonthlyVisitTrend 获取用户访问小程序数据月趋势
// accessToken 接口调用凭证
// start 开始日期，为自然月第一天。格式为 yyyymmdd
// end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
res, err := analysis.GetMonthlyVisitTrend("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

#### 周趋势

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html)

```go

// GetWeeklyVisitTrend 获取用户访问小程序数据周趋势
// accessToken 接口调用凭证
// start 开始日期，为自然月第一天。格式为 yyyymmdd
// end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
res, err := analysis.GetWeeklyVisitTrend("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

#### 日趋势

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html)

```go

// GetDailyVisitTrendAPI 获取用户访问小程序数据日趋势
// accessToken 接口调用凭证
// start 开始日期，为自然月第一天。格式为 yyyymmdd
// end 结束日期，为自然月最后一天，限定查询一个月数据。格式为 yyyymmdd
res, err := analysis.GetDailyVisitTrendAPI("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 用户画像

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html)

```go

// GetUserPortrait 获取小程序新增或活跃用户的画像分布数据。
// 时间范围支持昨天、最近7天、最近30天。
// 其中，新增用户数为时间范围内首次访问小程序的去重用户数，活跃用户数为时间范围内访问过小程序的去重用户数。
// start 开始日期。格式为 yyyymmdd
// end 结束日期，开始日期与结束日期相差的天数限定为0/6/29，分别表示查询最近1/7/30天数据，允许设置的最大值为昨日。格式为 yyyymmdd
res, err := analysis.GetUserPortrait("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 分布数据

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html)

```go

// GetVisitDistribution 获取用户小程序访问分布数据
// start 开始日期。格式为 yyyymmdd
// end 结束日期，限定查询 1 天数据，允许设置的最大值为昨日。格式为 yyyymmdd
res, err := analysis.GetVisitDistribution("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 访问页面

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitPage.html)

```go

// GetVisitPage 访问页面。
// 目前只提供按 page_visit_pv 排序的 top200。
// start 开始日期。格式为 yyyymmdd
// end 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
res, err := analysis.GetVisitPage("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 访问概况

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html)

```go

// getDailySummary 获取用户访问小程序数据概况
// start 开始日期。格式为 yyyymmdd
// end 结束日期，限定查询1天数据，允许设置的最大值为昨日。格式为 yyyymmdd
res, err := analysis.getDailySummary("access-token", "start-date-string", "end-date-string")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

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
if err := srv.Serve(); err != nil{
    // handle error and do something
}

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
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

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

### 发送统一服务消息

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html)

```go

// 引入子包
import "github.com/medivhzhan/weapp/message/template"

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

## 动态消息

### 创建被分享动态消息

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html)

```go

// 引入子包
import "github.com/medivhzhan/weapp/message/updatable"

// accessToken 接口调用凭证
res, err :=  CreateActivityID("access-token")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 修改被分享的动态消息

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html)

```go

// 引入子包
import "github.com/medivhzhan/weapp/message/updatable"

msg := updatable.Message{
    ID:    "activity-id",
    State: updatable.Started, // or updatable. Unstarted
    Template: updatable.Template{
        Params: []updatable.Param{
            updatable.Param{
                Name:  "param-name-you-want-change",
                Value: "value-you-want-change",
            },
        },
    },
}
// accessToken 接口调用凭证
res, err := msg.SetUpdatableMsg("access-token")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

---

## 附近的小程序

### 添加地点

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.add.html)

```go

// 引入子包
import "github.com/medivhzhan/weapp/message/nearby"

picList := nearby.PicList{
    List: []string{
        "your",
        "picture",
        "url",
        "list",
    },
}
serviceInfos := nearby.ServiceInfos{
    ServiceInfos: []nearby.ServiceInfo{
        nearby.ServiceInfo{
            ID:    1,
            Type:  2,
            Name:  "name",
            AppID: "appid",
            Path:  "path",
        },
    },
}

kfInfo := nearby.KFInfo{
    OpenKF:    true,
    KDHeading: "kf-head-img-url",
    KFName:    "kf-name",
}

point := nearby.Position{
    PicList:           picList,              // 门店图片，最多9张，最少1张，上传门店图片如门店外景、环境设施、商品服务等，图片将展示在微信客户端的门店页。图片链接通过文档https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1444738729中的《上传图文消息内的图片获取URL》接口获取。必填，文件格式为bmp、png、jpeg、jpg或gif，大小不超过5M pic_list是字符串，内容是一个json！
    ServiceInfos:      serviceInfos,         // 必服务标签列表 选填，需要填写服务标签ID、APPID、对应服务落地页的path路径，详细字段格式见下方示例
    StoreName:         "store_name",         // 门店名字 必填，门店名称需按照所选地理位置自动拉取腾讯地图门店名称，不可修改，如需修改请重现选择地图地点或重新创建地点
    Hour:              "hour",               // 营业时间，格式11:11-12:12 必填
    Credential:        "credential",         // 资质号 必填, 15位营业执照注册号或9位组织机构代码
    Address:           "address",            // 地址 必填
    CompanyName:       "company-name",       // 主体名字 必填
    QualificationList: "qualification-list", // 证明材料 必填 如果company_name和该小程序主体不一致，需要填qualification_list，详细规则见附近的小程序使用指南-如何证明门店的经营主体跟公众号或小程序帐号主体相关http://kf.qq.com/faq/170401MbUnim17040122m2qY.html
    KFInfo:            kfInfo,               // 客服信息 选填，可自定义服务头像与昵称，具体填写字段见下方示例kf_info pic_list是字符串，内容是一个json！
    PoiID:             "poi-id",             // 如果创建新的门店，poi_id字段为空 如果更新门店，poi_id参数则填对应门店的poi_id 选填
}

// accessToken  接口调用凭证
res, err := point.Add("access-token")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 删除地点

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.delete.html)

```go

// 引入子包
import "github.com/medivhzhan/weapp/message/nearby"

// accessToken  接口调用凭证
// id  附近地点 ID
res, err := point.Delete("access-token", "poi-id")
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 查看地点列表

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.getList.html)

```go

// 引入子包
import "github.com/medivhzhan/weapp/message/nearby"

// accessToken  接口调用凭证
// page  起始页id（从1开始计数）
// pageRows  每页展示个数（最多1000个）
res, err := point.GetList("access-token", 1, 10)
if err != nil {
    // handle error
    return
}

fmt.Printf("返回结果: %#v", res)

```

### 展示/取消展示附近小程序

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/nearby-poi/nearbyPoi.setShowStatus.html)

```go

// 引入子包
import "github.com/medivhzhan/weapp/message/nearby"

// accessToken  接口调用凭证
// poiID  附近地点 ID
// status  是否展示
res, err := point.SetShowStatus("access-token", "poi-id", nearby.Show)
if err != nil {
    // handle error
    return
}

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

## 内容检测

### 检测图片

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api/imgSecCheck.html)

```go

import "github.com/medivhzhan/weapp"

// 本地图片检测
//
// filename 要检测的图片本地路径
// token 接口调用凭证(access_token)
res, err := IMGSecCheck(filename, token string)
if err != nil {
    return
}

fmt.Printf("返回结果: %#v", res)

// 网络图片检测
//
// url 要检测的图片网络路径
// token 接口调用凭证(access_token)
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
// content 要检测的文本内容，长度不超过 500KB，编码格式为utf-8
// token 接口调用凭证(access_token)
res, err := MSGSecCheck(content, token string)
if err != nil {
    return
}

fmt.Printf("返回结果: %#v", res)

```

---

## 生物认证

### 秘钥签名验证

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html)

```go

import "github.com/medivhzhan/weapp"

res, err := VerifySignature("access-token", "user-openid", "result-json", "result-json-signature")
if err != nil {
    // handle error
    return
}

// res.IsOk
// res.Errcode
// res.Errmsg
fmt.Printf("返回结果: %#v", res)

```

---

## 数据解密

> **注意:**
> 在回调中调用 wx.login 登录，可能会刷新登录态。此时服务器使用 code 换取的 sessionKey 不是加密时使用的 sessionKey，导致解密失败。建议开发者提前进行 login；或者在回调中先使用 checkSession 进行登录态检查，避免 login 刷新登录态。

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#%E5%8A%A0%E5%AF%86%E6%95%B0%E6%8D%AE%E8%A7%A3%E5%AF%86%E7%AE%97%E6%B3%95)

### 解密手机号码

```go

import "github.com/medivhzhan/weapp"

// 解密手机号码
//
// ssk 通过 Login 向微信服务端请求得到的 session_key
// data 小程序通过 api 得到的加密数据(encryptedData)
// iv 小程序通过 api 得到的初始向量(iv)
phone , err := weapp.DecryptPhoneNumber(ssk, data, iv string)

fmt.Printf("手机数据: %#v", phone)

```

### 解密分享内容

```go

import "github.com/medivhzhan/weapp"

// 解密转发信息的加密数据
//
// ssk 通过 Login 向微信服务端请求得到的 session_key
// data 小程序通过 api 得到的加密数据(encryptedData)
// iv 小程序通过 api 得到的初始向量(iv)
//
// gid 小程序唯一群号
openGid , err := weapp.DecryptShareInfo(ssk, data, iv string)

```

### 解密用户信息

```go

import "github.com/medivhzhan/weapp"

// 解密用户信息
//
// rawData 不包括敏感信息的原始数据字符串, 用于计算签名。
// encryptedData 包括敏感数据在内的完整用户信息的加密数据
// signature 使用 sha1( rawData + session_key ) 得到字符串, 用于校验用户信息
// iv 加密算法的初始向量
// ssk 微信 session_key
ui, err := weapp.DecryptUserInfo(rawData, encryptedData, signature, iv, ssk string)
if err != nil {
    return
}

fmt.Printf("用户数据: %#v", ui)

```
