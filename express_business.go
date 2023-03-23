package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiBindAccount            = "/cgi-bin/express/business/account/bind"
	apiGetAllLogisticsAccount = "/cgi-bin/express/business/account/getall"
	apiGetExpressPath         = "/cgi-bin/express/business/path/get"
	apiAddExpressOrder        = "/cgi-bin/express/business/order/add"
	apiCancelExpressOrder     = "/cgi-bin/express/business/order/cancel"
	apiGetAllDelivery         = "/cgi-bin/express/business/delivery/getall"
	apiGetExpressOrder        = "/cgi-bin/express/business/order/get"
	apiGetPrinter             = "/cgi-bin/express/business/printer/getall"
	apiGetExpressQuota        = "/cgi-bin/express/business/quota/get"
	apiUpdatePrinter          = "/cgi-bin/express/business/printer/update"
	apiTestUpdateOrder        = "/cgi-bin/express/business/test_update_order"
)

// ExpressAccount 物流账号
type ExpressAccount struct {
	Type          BindType `json:"type"`           // bind表示绑定，unbind表示解除绑定
	BizID         string   `json:"biz_id"`         // 快递公司客户编码
	DeliveryID    string   `json:"delivery_id"`    // 快递公司 ID
	Password      string   `json:"password"`       // 快递公司客户密码
	RemarkContent string   `json:"remark_content"` // 备注内容（提交EMS审核需要）
}

// BindType 绑定动作类型
type BindType = string

// 所有绑定动作类型
const (
	Bind   = "bind"   // 绑定
	Unbind = "unbind" // 解除绑定
)

// Bind 绑定、解绑物流账号
// token 接口调用凭证
func (cli *Client) BindLogisticsAccount(ea *ExpressAccount) (*request.CommonError, error) {
	api := baseURL + apiBindAccount

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.bindLogisticsAccount(api, token, ea)
}

func (cli *Client) bindLogisticsAccount(api, token string, ea *ExpressAccount) (*request.CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(url, ea, res); err != nil {
		return nil, err
	}

	return res, nil
}

// AccountList 所有绑定的物流账号
type AccountList struct {
	request.CommonError
	Count uint `json:"count"` // 账号数量
	List  []struct {
		BizID           string     `json:"biz_id"`            // 	快递公司客户编码
		DeliveryID      string     `json:"delivery_id"`       // 	快递公司 ID
		CreateTime      uint       `json:"create_time"`       // 	账号绑定时间
		UpdateTime      uint       `json:"update_time"`       // 	账号更新时间
		StatusCode      BindStatus `json:"status_code"`       // 	绑定状态
		Alias           string     `json:"alias"`             // 	账号别名
		RemarkWrongMsg  string     `json:"remark_wrong_msg"`  // 	账号绑定失败的错误信息（EMS审核结果）
		RemarkContent   string     `json:"remark_content"`    // 	账号绑定时的备注内容（提交EMS审核需要)）
		QuotaNum        uint       `json:"quota_num"`         // 	电子面单余额
		QuotaUpdateTime uint       `json:"quota_update_time"` // 	电子面单余额更新时间
	} `json:"list"` // 账号列表
}

// BindStatus 账号绑定状态
type BindStatus = int8

// 所有账号绑定状态
const (
	BindSuccess = 0  // 成功
	BindFailed  = -1 // 系统失败
)

// GetAllLogisticsAccount 获取所有绑定的物流账号
// token 接口调用凭证
func (cli *Client) GetAllLogisticsAccount() (*AccountList, error) {
	api := baseURL + apiGetAllLogisticsAccount
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getAllLogisticsAccount(api, token)
}

func (cli *Client) getAllLogisticsAccount(api, token string) (*AccountList, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(AccountList)
	if err := cli.request.Post(url, requestParams{}, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ExpressPathGetter 查询运单轨迹所需参数
type ExpressPathGetter ExpressOrderGetter

// GetExpressPathResponse 运单轨迹
type GetExpressPathResponse struct {
	request.CommonError
	OpenID       string            `json:"openid"`         // 用户openid
	DeliveryID   string            `json:"delivery_id"`    // 快递公司 ID
	WaybillID    string            `json:"waybill_id"`     // 运单 ID
	PathItemNum  uint              `json:"path_item_num"`  // 轨迹节点数量
	PathItemList []ExpressPathNode `json:"path_item_list"` // 轨迹节点列表
}

// ExpressPathNode 运单轨迹节点
type ExpressPathNode struct {
	ActionTime uint   `json:"action_time"` // 轨迹节点 Unix 时间戳
	ActionType uint   `json:"action_type"` // 轨迹节点类型
	ActionMsg  string `json:"action_msg"`  // 轨迹节点详情
}

// Get 查询运单轨迹
// token 接口调用凭证
func (cli *Client) GetLogisticsPath(ep *ExpressPathGetter) (*GetExpressPathResponse, error) {
	api := baseURL + apiGetExpressPath

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getLogisticsPath(api, token, ep)
}

func (cli *Client) getLogisticsPath(api, token string, ep *ExpressPathGetter) (*GetExpressPathResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetExpressPathResponse)
	if err := cli.request.Post(url, ep, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ExpressOrderSource 订单来源
type ExpressOrderSource = uint8

// 所有订单来源
const (
	FromWeapp   ExpressOrderSource = 0 // 小程序订单
	FromAppOrH5                    = 2 // APP或H5订单
)

// ExpressOrderCreator 订单创建器
type ExpressOrderCreator struct {
	ExpressOrder
	AddSource  ExpressOrderSource `json:"add_source"`            // 订单来源，0为小程序订单，2为App或H5订单，填2则不发送物流服务通知
	WXAppID    string             `json:"wx_appid,omitempty"`    // App或H5的appid，add_source=2时必填，需和开通了物流助手的小程序绑定同一open帐号
	ExpectTime uint               `json:"expect_time,omitempty"` //	顺丰必须填写此字段。预期的上门揽件时间，0表示已事先约定取件时间；否则请传预期揽件时间戳，需大于当前时间，收件员会在预期时间附近上门。例如expect_time为“1557989929”，表示希望收件员将在2019年05月16日14:58:49-15:58:49内上门取货。
	TagID      uint               `json:"tagid,omitempty"`       //订单标签id，用于平台型小程序区分平台上的入驻方，tagid须与入驻方账号一一对应，非平台型小程序无需填写该字段
}

// InsureStatus 保价状态
type InsureStatus = uint8

// 所有保价状态
const (
	Uninsured = 0 // 不保价
	Insured   = 1 // 保价
)

// CreateExpressOrderResponse 创建订单返回数据
type CreateExpressOrderResponse struct {
	request.CommonError
	OrderID     string `json:"order_id"`   //	订单ID，下单成功时返回
	WaybillID   string `json:"waybill_id"` //	运单ID，下单成功时返回
	WaybillData []struct {
		Key   string `json:"key"`   // 运单信息 key
		Value string `json:"value"` // 运单信息 value
	} `json:"waybill_data"` //	运单信息，下单成功时返回
	DeliveryResultcode int    `json:"delivery_resultcode"` //	快递侧错误码，下单失败时返回
	DeliveryResultmsg  string `json:"delivery_resultmsg"`  //	快递侧错误信息，下单失败时返回
}

// Create 生成运单
// token 接口调用凭证
func (cli *Client) AddLogisticOrder(creator *ExpressOrderCreator) (*CreateExpressOrderResponse, error) {
	api := baseURL + apiAddExpressOrder
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.addLogisticOrder(api, token, creator)
}

func (cli *Client) addLogisticOrder(api, token string, creator *ExpressOrderCreator) (*CreateExpressOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CreateExpressOrderResponse)
	if err := cli.request.Post(url, creator, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CancelOrderResponse 取消订单返回数据
type CancelOrderResponse struct {
	request.CommonError
	Count uint `json:"count"` //快递公司数量
	Data  []struct {
		DeliveryID   string `json:"delivery_id"`   // 快递公司 ID
		DeliveryName string `json:"delivery_name"` // 快递公司名称

	} `json:"data"` //快递公司信息列表
}

// DeliveryList 支持的快递公司列表
type DeliveryList struct {
	request.CommonError
	Count uint `json:"count"` // 快递公司数量
	Data  []struct {
		ID   string `json:"delivery_id"`   // 快递公司 ID
		Name string `json:"delivery_name"` // 快递公司名称
	} `json:"data"` // 快递公司信息列表
}

// GetAllDelivery 获取支持的快递公司列表
// token 接口调用凭证
func (cli *Client) GetAllDelivery() (*DeliveryList, error) {
	api := baseURL + apiGetAllDelivery
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getAllDelivery(api, token)
}

func (cli *Client) getAllDelivery(api, token string) (*DeliveryList, error) {
	queries := requestQueries{
		"access_token": token,
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(DeliveryList)
	if err := cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ExpressOrderGetter 订单获取器
type ExpressOrderGetter struct {
	OrderID    string `json:"order_id"`         // 订单 ID，需保证全局唯一
	OpenID     string `json:"openid,omitempty"` // 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID string `json:"delivery_id"`      // 快递公司ID，参见getAllDelivery
	WaybillID  string `json:"waybill_id"`       // 运单ID
}

// GetExpressOrderResponse 获取运单返回数据
type GetExpressOrderResponse struct {
	request.CommonError
	PrintHTML   string `json:"print_html"` // 运单 html 的 BASE64 结果
	WaybillData []struct {
		Key   string `json:"key"`   // 运单信息 key
		Value string `json:"value"` // 运单信息 value
	} `json:"waybill_data"` // 运单信息
}

// Get 获取运单数据
// token 接口调用凭证
func (cli *Client) GetLogisticsOrder(getter *ExpressOrderGetter) (*GetExpressOrderResponse, error) {
	api := baseURL + apiGetExpressOrder
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getLogisticsOrder(api, token, getter)
}

func (cli *Client) getLogisticsOrder(api, token string, getter *ExpressOrderGetter) (*GetExpressOrderResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetExpressOrderResponse)
	if err := cli.request.Post(url, getter, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ExpressOrderCanceler 订单取消器
type ExpressOrderCanceler ExpressOrderGetter

// Cancel 取消运单
// token 接 口调用凭证
func (cli *Client) CancelLogisticsOrder(canceler *ExpressOrderCanceler) (*request.CommonError, error) {
	api := baseURL + apiCancelExpressOrder
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.cancelLogisticsOrder(api, token, canceler)
}

func (cli *Client) cancelLogisticsOrder(api, token string, canceler *ExpressOrderCanceler) (*request.CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(url, canceler, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPrinterResponse 获取打印员返回数据
type GetPrinterResponse struct {
	request.CommonError
	Count     uint     `json:"count"`  // 已经绑定的打印员数量
	OpenID    []string `json:"openid"` // 打印员 openid 列表
	TagIDList []string `json:"tagid_list"`
}

// GetPrinter 获取打印员。若需要使用微信打单 PC 软件，才需要调用。
// token 接口调用凭证
func (cli *Client) GetPrinter(token string) (*GetPrinterResponse, error) {
	api := baseURL + apiGetPrinter
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getPrinter(api, token)
}

func (cli *Client) getPrinter(api, token string) (*GetPrinterResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetPrinterResponse)
	if err := cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// QuotaGetter 电子面单余额获取器
type QuotaGetter struct {
	DeliveryID string `json:"delivery_id"` // 快递公司ID，参见getAllDelivery
	BizID      string `json:"biz_id"`      // 快递公司客户编码
}

// QuotaGetResponse 电子面单余额
type QuotaGetResponse struct {
	request.CommonError
	QuotaNum int `json:"quota_num"` // 电子面单余额
}

// Get 获取电子面单余额。仅在使用加盟类快递公司时，才可以调用。
func (cli *Client) GetExpressQuota(getter *QuotaGetter) (*QuotaGetResponse, error) {
	api := baseURL + apiGetExpressQuota
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getExpressQuota(api, token, getter)
}

func (cli *Client) getExpressQuota(api, token string, getter *QuotaGetter) (*QuotaGetResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(QuotaGetResponse)
	if err := cli.request.Post(url, getter, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateExpressOrderTester 模拟的快递公司更新订单
type UpdateExpressOrderTester struct {
	BizID      string `json:"biz_id"`      // 商户id,需填test_biz_id
	OrderID    string `json:"order_id"`    //	订单ID，下单成功时返回
	WaybillID  string `json:"waybill_id"`  // 运单 ID
	DeliveryID string `json:"delivery_id"` // 快递公司 ID
	ActionTime uint   `json:"action_time"` // 轨迹变化 Unix 时间戳
	ActionType uint   `json:"action_type"` // 轨迹变化类型
	ActionMsg  string `json:"action_msg"`  // 轨迹变化具体信息说明，展示在快递轨迹详情页中。若有手机号码，则直接写11位手机号码。使用UTF-8编码。
}

// Test 模拟快递公司更新订单状态, 该接口只能用户测试
func (cli *Client) TestUpdateExpressOrder(tester *UpdateExpressOrderTester) (*request.CommonError, error) {
	api := baseURL + apiTestUpdateOrder
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.testUpdateExpressOrder(api, token, tester)
}

func (cli *Client) testUpdateExpressOrder(api, token string, tester *UpdateExpressOrderTester) (*request.CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(url, tester, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PrinterUpdater 打印员更新器
type PrinterUpdater struct {
	OpenID    string   `json:"openid"`      // 打印员 openid
	Type      BindType `json:"update_type"` // 更新类型
	TagIDList string   `json:"tagid_list"`  // 用于平台型小程序设置入驻方的打印员面单打印权限，同一打印员最多支持10个tagid，使用逗号分隔，如填写123，456，表示该打印员可以拉取到tagid为123和456的下的单，非平台型小程序无需填写该字段
}

// Update 更新打印员。若需要使用微信打单 PC 软件，才需要调用。
func (cli *Client) UpdateExpressOrder(updater *PrinterUpdater) (*request.CommonError, error) {
	api := baseURL + apiUpdatePrinter
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.updateExpressOrder(api, token, updater)
}

func (cli *Client) updateExpressOrder(api, token string, updater *PrinterUpdater) (*request.CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(url, updater, res); err != nil {
		return nil, err
	}

	return res, nil
}
