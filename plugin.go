package weapp

const (
	apiAddPlugin    = "/wxa/plugin"
	apiGetPluginDev = "/wxa/devplugin"
)

// ApplyPlugin 向插件开发者发起使用插件的申请
// @accessToken 接口调用凭证
// action	string		是	此接口下填写 "apply"
// @appID	string		是	插件 appId
// @reason	string		否	申请使用理由
func ApplyPlugin(accessToken, appID, reason string) (*commonError, error) {
	api, err := tokenAPI(baseURL+apiAddPlugin, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"action":       "apply",
		"reason":       reason,
		"plugin_appid": appID,
	}

	res := new(commonError)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPluginDevApplyList 获取当前所有插件使用方
// @accessToken 接口调用凭证
// @page	number		是	要拉取第几页的数据
// @num		是	每页的记录数
func GetPluginDevApplyList(accessToken string, page, num uint) (*commonError, error) {
	api, err := tokenAPI(baseURL+apiGetPluginDev, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"num":    num,
		"page":   page,
		"action": "dev_apply_list",
	}

	res := new(commonError)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPluginListResponse 查询已添加的插件返回数据
type GetPluginListResponse struct {
	commonError
	PluginList []struct {
		AppID      string `json:"appid"`      // 插件 appId
		Status     int8   `json:"status"`     // 插件状态
		Nickname   string `json:"nickname"`   // 插件昵称
		HeadImgURL string `json:"headimgurl"` // 插件头像
	} `json:"plugin_list"` // 申请或使用中的插件列表
}

// GetPluginList 查询已添加的插件
// @accessToken 接口调用凭证
func GetPluginList(accessToken string) (*GetPluginListResponse, error) {
	api, err := tokenAPI(baseURL+apiAddPlugin, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"action": "list",
	}

	res := new(GetPluginListResponse)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DevAction 修改操作
type DevAction = string

// 所有修改操作
const (
	DevAgree  DevAction = "dev_agree"  // 同意申请
	DevRefuse DevAction = "dev_refuse" // 拒绝申请
	DevDelete DevAction = "dev_refuse" // 删除已拒绝的申请者
)

// SetDevPluginApplyStatus 修改插件使用申请的状态
// @accessToken 接口调用凭证
// @appID 使用者的 appid。同意申请时填写。
// @reason 拒绝理由。拒绝申请时填写。
// @action 修改操作
func SetDevPluginApplyStatus(accessToken, appID, reason string, action DevAction) (*commonError, error) {
	api, err := tokenAPI(baseURL+apiGetPluginDev, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"action": action,
		"appid":  appID,
		"reason": reason,
	}

	res := new(commonError)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// UnbindPlugin 查询已添加的插件
// @accessToken 接口调用凭证
// @appID 插件 appId
func UnbindPlugin(accessToken, appID string) (*commonError, error) {
	api, err := tokenAPI(baseURL+apiAddPlugin, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"action":       "unbind",
		"plugin_appid": appID,
	}

	res := new(commonError)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
