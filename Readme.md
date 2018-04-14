# 微信小程序 SDK （for Golang）

## 使用

- 拉取代码

```sh
go get github.com/medivhzhan/weapp
```

- 引入

```go
import "github.com/medivhzhan/weapp"
```

## 获取 access_token 及其有效期

```go
// 获取次数有限制 获取后请缓存

weapp.AccessToken(appID, secret string) (accessToken string, expire uint, err error)

```

## 用户登录

```go
// 需要从小程序客户端获取到的code

weapp.Login(appID, secret, code string) (openID string, sessionKey uint, err error)

```

## 返回内容

```go
// Response 请求微信返回基础数据

type Response struct {
	Errcode int
	Errmsg  string
}

```

***

## 二维码

### 获取小程序码: 适用于需要的码数量较少的业务场景

```go
// AppCode 获取小程序码
// 可接受path参数较长 生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// @ path 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// @ lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// @ token 微信access_token
// @ filename 文件储存路径

weapp.AppCode(path string, width int, autoColor bool, lineColor, accessToken, filename string) error

```

### 获取小程序码: 适用于需要的码数量极多的业务场景

```go
// UnlimitedAppCode 获取小程序码
// 可接受页面参数较短 生成个数不受限 适用于需要的码数量极多的业务场景
// 根路径前不要填加'/' 不能携带参数（参数请放在scene字段里）
// @ scene 需要使用 decodeURIComponent 才能获取到生成二维码时传入的 scene
// @ page 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ autoColor 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
// @ lineColor autoColor 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"},十进制表示
// @ token 微信access_token
// @ filename 文件储存路径

weapp.UnlimitedAppCode(scene, path string, width int, autoColor bool, lineColor, accessToken string) error

```

### 获取小程序二维码: 适用于需要的码数量较少的业务场景

```go
// QRCode 获取小程序二维码
// 可接受path参数较长，生成个数受限 永久有效 适用于需要的码数量较少的业务场景
// @ path 识别二维码后进入小程序的页面链接
// @ width 图片宽度
// @ token 微信access_token
// @ filename 文件储存路径

weapp.QRCode(path string, width int, token string) error

```

***

## 模板消息

- 引入子包

```go
import "github.com/medivhzhan/weapp/message/template"
```

### 获取小程序模板库标题列表

```go
// List 获取小程序模板库标题列表
// @ offset 开始获取位置 从0开始
// @ count 获取记录条数 最大为20
// @ token 微信 access_token

message.List(offset uint, count uint, token string) (list []Template, total uint, err error)

```

### 获取帐号下已存在的模板列表

```go
// Selves 获取帐号下已存在的模板列表
// @ offset 开始获取位置 从0开始
// @ count 获取记录条数 最大为20
// @ token 微信 access_token

message.Selves(offset uint, count uint, token string) (list []Template, total uint, err error)

```

### 获取模板库某个模板标题下关键词库

```go
// Get 获取模板库某个模板标题下关键词库
// @ id 模板ID
// @ token 微信 access_token

message.Get(id, token string) (keywords []Keyword, err error)

```

### 组合模板并添加至帐号下的个人模板库

```go
// Add 组合模板并添加至帐号下的个人模板库
// @ id 模板ID
// @ token 微信 aceess_token
// @ keywordIDList 关键词 ID 列表
// 返回新建模板ID和错误信息

message.Add(id, token string, keywordIDList []uint) (string, error)

```

### 删除帐号下的某个模板

```go
// Delete 删除帐号下的某个模板
// @ id 模板ID
// @ token 微信 aceess_token

message.Delete(id, token string) error

```

### 发送模板消息

```go
// Send 发送模板消息
// @ openid 接收者（用户）的 openid
// @ template 所需下发的模板消息的id
// @ page 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
// @ formID 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
// @ data 模板内容，不填则下发空模板
// @ color 模板内容字体的颜色，不填默认黑色
// @ emphasisKeyword 模板需要放大的关键词，不填则默认无放大

message.Send(openid, template, page, formID, data, color, emphasisKeyword, token string) error

```

***

## 客服消息

### 回复

```go
import "github.com/medivhzhan/weapp/message/service"

// 文本消息
msg := service.Msg{
    Content: "消息内容",
}

// 图片消息
msg := service.Msg{
    MediaID: "微信media_id"
}

// 图文链接消息
msg := service.Msg{
    Title: "标题"
    Description: "描述"
    URL: "点击跳转链接"
    ThumbURL: "图片链接"
}

// 卡片消息
msg := service.Msg{
    Title: "标题"
    PagePath: "小程序页面路径"
    ThumbMediaID: "卡封面图片 media_id"
}

// SendTo 发送消息
// @ openid 用户openid
// @ token 微信 access_token

msg.SendTo(openid, token string) (weapp.Response, error)

// 返回参照码 (weapp.Response.Errcode)

// SystemBusy 系统繁忙 稍候再试
SystemBusy weapp.Code = -1

// SendSuccess 发送消息成功
SendSuccess weapp.Code = 0

// ErrInvalidToken 获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的小程序调用接口
ErrInvalidToken weapp.Code = 40001

// ErrInvalidClaims 不合法的凭证类型
ErrInvalidClaims weapp.Code = 40002

// ErrInvalidOpenid 不合法的 OpenID，请开发者确认OpenID否是其他小程序的 OpenID
ErrInvalidOpenid weapp.Code = 40003

// ErrOutOfTime 回复时间超过限制
ErrOutOfTime weapp.Code = 45015

// ErrOutOfLimit 客服接口下行条数超过上限
ErrOutOfLimit weapp.Code = 45047

// ErrUnauthorized api功能未授权，请确认小程序已获得该接口
ErrUnauthorized weapp.Code = 48001

```

***

## 未实现功能

1. 切换JSON接口
1. 接收客服消息
1. 微信支付
1. 数据统计
1. 临时素材
