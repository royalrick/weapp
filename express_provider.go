package weapp

const (
	apiGetContact      = "/cgi-bin/express/delivery/contact/get"
	apiPreviewTemplate = "/cgi-bin/express/delivery/template/preview"
	apiUpdateBusiness  = "/cgi-bin/express/delivery/service/business/update"
	apiUpdatePath      = "/cgi-bin/express/delivery/path/update"
)

// GetContactResponse 获取面单联系人信息返回数据
type GetContactResponse struct {
	CommonError
	WaybillID string      `json:"waybill_id"` // 运单 ID
	Sender    ContactUser `json:"sender"`     // 发件人信息
	Receiver  ContactUser `json:"receiver"`   // 收件人信息
}

// ContactUser 联系人
type ContactUser struct {
	Address string `json:"address"` //地址，已经将省市区信息合并
	Name    string `json:"name"`    //用户姓名
	Tel     string `json:"tel"`     //座机号码
	Mobile  string `json:"mobile"`  //手机号码
}

// GetContact 获取面单联系人信息
// accessToken, token, watBillID 接口调用凭证
func GetContact(accessToken, token, watBillID string) (*GetContactResponse, error) {
	api := baseURL + apiGetContact
	return getContact(api, accessToken, token, watBillID)
}

func getContact(api, accessToken, token, watBillID string) (*GetContactResponse, error) {
	url, err := tokenAPI(api, accessToken)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"token":      token,
		"waybill_id": watBillID,
	}

	res := new(GetContactResponse)
	if err := postJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ExpressTemplatePreviewer 面单模板预览器
type ExpressTemplatePreviewer struct {
	WaybillID       string       `json:"waybill_id"`       // 运单 ID
	WaybillTemplate string       `json:"waybill_template"` // 面单 HTML 模板内容（需经 Base64 编码）
	WaybillData     string       `json:"waybill_data"`     // 面单数据。详情参考下单事件返回值中的 WaybillData
	Custom          ExpressOrder `json:"custom"`           // 商户下单数据，格式是商户侧下单 API 中的请求体
}

// PreviewTemplateResponse 预览面单模板返回数据
type PreviewTemplateResponse struct {
	CommonError
	WaybillID               string `json:"waybill_id"`                // 运单 ID
	RenderedWaybillTemplate string `json:"rendered_waybill_template"` // 渲染后的面单 HTML 文件（已经过 Base64 编码）
}

// Preview 预览面单模板。用于调试面单模板使用。
// token 接口调用凭证
func (previewer *ExpressTemplatePreviewer) Preview(token string) (*PreviewTemplateResponse, error) {
	api := baseURL + apiPreviewTemplate
	return previewer.preview(api, token)
}

func (previewer *ExpressTemplatePreviewer) preview(api, token string) (*PreviewTemplateResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(PreviewTemplateResponse)
	if err := postJSON(url, previewer, res); err != nil {
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

// BusinessUpdater 商户审核结果更新器
type BusinessUpdater struct {
	ShopAppID  string             `json:"shop_app_id"`          // 商户的小程序AppID，即审核商户事件中的 ShopAppID
	BizID      string             `json:"biz_id"`               // 商户账户
	ResultCode BusinessResultCode `json:"result_code"`          // 审核结果，0 表示审核通过，其他表示审核失败
	ResultMsg  string             `json:"result_msg,omitempty"` // 审核错误原因，仅 result_code 不等于 0 时需要设置
}

// Update 更新商户审核结果
// token 接口调用凭证
func (updater *BusinessUpdater) Update(token string) (*CommonError, error) {
	api := baseURL + apiUpdateBusiness
	return updater.update(api, token)
}

func (updater *BusinessUpdater) update(api, token string) (*CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postJSON(url, updater, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ExpressPathUpdater 运单轨迹更新器
type ExpressPathUpdater struct {
	Token      string `json:"token"`       // 商户侧下单事件中推送的 Token 字段
	WaybillID  string `json:"waybill_id"`  // 运单 ID
	ActionTime uint   `json:"action_time"` // 轨迹变化 Unix 时间戳
	ActionType uint   `json:"action_type"` // 轨迹变化类型
	ActionMsg  string `json:"action_msg"`  // 轨迹变化具体信息说明，展示在快递轨迹详情页中。若有手机号码，则直接写11位手机号码。使用UTF-8编码。
}

// Update 更新运单轨迹
// token 接口调用凭证
func (updater *ExpressPathUpdater) Update(token string) (*CommonError, error) {
	api := baseURL + apiUpdatePath
	return updater.update(api, token)
}

func (updater *ExpressPathUpdater) update(api, token string) (*CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	if err := postJSON(url, updater, res); err != nil {
		return nil, err
	}

	return res, nil
}
