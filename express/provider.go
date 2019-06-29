package express

import "github.com/medivhzhan/weapp"

const (
	apiGetContact      = "/cgi-bin/express/delivery/contact/get"
	apiPreviewTemplate = "/cgi-bin/express/delivery/template/preview"
	apiUpdateBusiness  = "/cgi-bin/express/delivery/service/business/update"
)

// ContactGetter 面单联系人信息获取器
type ContactGetter struct {
	Token     string `json:"token"`      // 商户侧下单事件中推送的 Token 字段
	WaybillID string `json:"waybill_id"` // 运单 ID

}

// GetContactResponse 获取面单联系人信息返回数据
type GetContactResponse struct {
	weapp.Response
	WaybillID string        `json:"waybill_id"` // 运单 ID
	Sender    []ContactUser `json:"sender"`     // 发件人信息
	Receiver  []ContactUser `json:"receiver"`   // 收件人信息
}

// ContactUser 联系人
type ContactUser struct {
	Address string `json:"address"` //地址，已经将省市区信息合并
	Name    string `json:"name"`    //用户姓名
	Tel     string `json:"tel"`     //座机号码
	Mobile  string `json:"mobile"`  //手机号码
}

// Get 获取面单联系人信息
// @accessToken 接口调用凭证
func (cg *ContactGetter) Get(accessToken string) (*GetContactResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiGetContact, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(GetContactResponse)
	if err := weapp.PostJSON(api, cg, res); err != nil {
		return nil, err
	}

	return res, nil
}

// TemplateViewer 面单模板预览器
type TemplateViewer struct {
	WaybillID       string `json:"waybill_id"`       // 运单 ID
	WaybillTemplate string `json:"waybill_template"` // 面单 HTML 模板内容（需经 Base64 编码）
	WaybillData     string `json:"waybill_data"`     // 面单数据。详情参考下单事件返回值中的 WaybillData
	Custom          Order  `json:"custom"`           // 商户下单数据，格式是商户侧下单 API 中的请求体

}

// PreviewTemplateResponse 预览面单模板返回数据
type PreviewTemplateResponse struct {
	weapp.Response
	WaybillID               string `json:"waybill_id"`                // 运单 ID
	RenderedWaybillTemplate string `json:"rendered_waybill_template"` // 渲染后的面单 HTML 文件（已经过 Base64 编码）
}

// Preview 预览面单模板。用于调试面单模板使用。
// @accessToken 接口调用凭证
func (tv *OrderCreator) Preview(accessToken string) (*PreviewTemplateResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiPreviewTemplate, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(PreviewTemplateResponse)
	if err := weapp.PostJSON(api, tv, res); err != nil {
		return nil, err
	}

	return res, nil
}

// BusinessResultCode 商户审核结果状态码
type BusinessResultCode = int8

// 所有商户审核结果状态码
const (
	ResultSuccess BusinessResultCode = 0 // 审核通过
	ResultFailed                     = 1 // 审核失败
)

// BusinnessUpdater 商户审核结果更新器
type BusinnessUpdater struct {
	ShopAppID  string             `json:"shop_app_id"`          // 商户的小程序AppID，即审核商户事件中的 ShopAppID
	BizID      string             `json:"biz_id"`               // 商户账户
	ResultCode BusinessResultCode `json:"result_code"`          // 审核结果，0 表示审核通过，其他表示审核失败
	ResultMsg  string             `json:"result_msg,omitempty"` // 审核错误原因，仅 result_code 不等于 0 时需要设置
}

// Update 更新商户审核结果
// @accessToken 接口调用凭证
func (bu *BusinnessUpdater) Update(accessToken string) (*weapp.Response, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiPreviewTemplate, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(weapp.Response)
	if err := weapp.PostJSON(api, bu, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PathUpdater 运单轨迹更新器
type PathUpdater struct {
	Token      string `json:"token"`       // 商户侧下单事件中推送的 Token 字段
	WaybillID  string `json:"waybill_id"`  // 运单 ID
	ActionTime uint   `json:"action_time"` // 轨迹变化 Unix 时间戳
	ActionType int    `json:"action_type"` // 轨迹变化类型
	ActionMsg  string `json:"action_msg"`  // 轨迹变化具体信息说明，展示在快递轨迹详情页中。若有手机号码，则直接写11位手机号码。使用UTF-8编码。
}

// Update 更新运单轨迹
// @accessToken 接口调用凭证
func (pu *PathUpdater) Update(accessToken string) (*weapp.Response, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+apiPreviewTemplate, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(weapp.Response)
	if err := weapp.PostJSON(api, pu, res); err != nil {
		return nil, err
	}

	return res, nil
}
