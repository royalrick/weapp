package weapp

import (
	"strconv"
)

const (
	apiAddTemplate                = "/wxaapi/newtmpl/addtemplate"
	apiDeleteTemplate             = "/wxaapi/newtmpl/deltemplate"
	apiGetTemplateCategory        = "/wxaapi/newtmpl/getcategory"
	apiGetPubTemplateKeyWordsById = "/wxaapi/newtmpl/getpubtemplatekeywords"
	apiGetPubTemplateTitleList    = "/wxaapi/newtmpl/getpubtemplatetitles"
	apiGetTemplateList            = "/wxaapi/newtmpl/gettemplate"
	apiSendSubscribeMessage       = "/cgi-bin/message/subscribe/send"
)

// AddTemplateResponse 添加模版消息返回数据
type AddTemplateResponse struct {
	CommonError
	Pid string `json:"priTmplId"` // 添加至帐号下的模板id，发送小程序订阅消息时所需
}

// AddTemplate 组合模板并添加至帐号下的个人模板库
//
// token 微信 access_token
// tid 模板ID
// desc 服务场景描述，15个字以内
// keywordIDList 关键词 ID 列表
func AddTemplate(token, tid, desc string, keywordIDList []int32) (*AddTemplateResponse, error) {
	api := baseURL + apiAddTemplate
	return addTemplate(api, token, tid, desc, keywordIDList)
}

func addTemplate(api, token, tid, desc string, keywordIDList []int32) (*AddTemplateResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"tid":       tid,
		"kidList":   keywordIDList,
		"sceneDesc": desc,
	}

	res := new(AddTemplateResponse)
	err = postJSON(api, params, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteTemplate 删除帐号下的某个模板
//
// token 微信 access_token
// pid 模板ID
func DeleteTemplate(token, pid string) (*CommonError, error) {
	api := baseURL + apiDeleteTemplate
	return deleteTemplate(api, token, pid)
}

func deleteTemplate(api, token, pid string) (*CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"priTmplId": pid,
	}

	res := new(CommonError)
	err = postJSON(api, params, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTemplateCategoryResponse 删除帐号下的某个模板返回数据
type GetTemplateCategoryResponse struct {
	CommonError
	Data []struct {
		ID   int    `json:"id"`   // 类目id，查询公共库模版时需要
		Name string `json:"name"` // 类目的中文名
	} `json:"data"` // 类目列表
}

// GetTemplateCategory 删除帐号下的某个模板
//
// token 微信 access_token
func GetTemplateCategory(token string) (*GetTemplateCategoryResponse, error) {
	api := baseURL + apiGetTemplateCategory
	return getTemplateCategory(token, api)
}

func getTemplateCategory(token, api string) (*GetTemplateCategoryResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetTemplateCategoryResponse)
	err = getJSON(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetPubTemplateKeyWordsByIdResponse 模板标题下的关键词列表
type GetPubTemplateKeyWordsByIdResponse struct {
	CommonError
	Count int32 `json:"count"` // 模版标题列表总数
	Data  []struct {
		Kid     int    `json:"kid"`     // 关键词 id，选用模板时需要
		Name    string `json:"name"`    // 关键词内容
		Example string `json:"example"` // 关键词内容对应的示例
		Rule    string `json:"rule"`    // 参数类型
	} `json:"data"` // 关键词列表
}

// GetPubTemplateKeyWordsById 获取模板标题下的关键词列表
//
// token 微信 access_token
// tid 模板ID
func GetPubTemplateKeyWordsById(token, tid string) (*GetPubTemplateKeyWordsByIdResponse, error) {
	api := baseURL + apiGetPubTemplateKeyWordsById
	return getPubTemplateKeyWordsById(api, token, tid)
}

func getPubTemplateKeyWordsById(api, token, tid string) (*GetPubTemplateKeyWordsByIdResponse, error) {
	queries := requestQueries{
		"access_token": token,
		"tid":          tid,
	}
	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(GetPubTemplateKeyWordsByIdResponse)
	if err = getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPubTemplateTitleListResponse 帐号所属类目下的公共模板标题
type GetPubTemplateTitleListResponse struct {
	CommonError
	Count uint `json:"count"` // 模版标题列表总数
	Data  []struct {
		Tid        int    `json:"tid"`        // 模版标题 id
		Title      string `json:"title"`      // 模版标题
		Type       int32  `json:"type"`       // 模版类型，2 为一次性订阅，3 为长期订阅
		CategoryId string `json:"categoryId"` // 模版所属类目 id
	} `json:"data"` // 模板标题列表
}

// GetPubTemplateTitleList 获取帐号所属类目下的公共模板标题
//
// token 微信 access_token
// ids 类目 id，多个用逗号隔开
// start 用于分页，表示从 start 开始。从 0 开始计数。
// limit 用于分页，表示拉取 limit 条记录。最大为 30
func GetPubTemplateTitleList(token, ids string, start, limit int) (*GetPubTemplateTitleListResponse, error) {
	api := baseURL + apiGetPubTemplateTitleList
	return getPubTemplateTitleList(api, token, ids, start, limit)
}

func getPubTemplateTitleList(api, token, ids string, start, limit int) (*GetPubTemplateTitleListResponse, error) {

	queries := requestQueries{
		"access_token": token,
		"ids":          ids,
		"start":        strconv.Itoa(start),
		"limit":        strconv.Itoa(limit),
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(GetPubTemplateTitleListResponse)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetTemplateListResponse 获取模板列表返回的数据
type GetTemplateListResponse struct {
	CommonError
	Data []struct {
		Pid     string `json:"priTmplId"` // 添加至帐号下的模板 id，发送小程序订阅消息时所需
		Title   string `json:"title"`     // 模版标题
		Content string `json:"content"`   // 模版内容
		Example string `json:"example"`   // 模板内容示例
		Type    int32  `json:"type"`      // 模版类型，2 为一次性订阅，3 为长期订阅
	} `json:"data"` // 个人模板列表
}

// GetTemplateList 获取帐号下已存在的模板列表
//
// token 微信 access_token
func GetTemplateList(token string) (*GetTemplateListResponse, error) {
	api := baseURL + apiGetTemplateList
	return getTemplateList(api, token)
}

func getTemplateList(api, token string) (*GetTemplateListResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetTemplateListResponse)
	if err := getJSON(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	ToUser           string               `json:"touser"`
	TemplateID       string               `json:"template_id"`
	Page             string               `json:"page,omitempty"`
	MiniprogramState MiniprogramState     `json:"miniprogram_state,omitempty"`
	Data             SubscribeMessageData `json:"data"`
}

// MiniprogramState 跳转小程序类型
type MiniprogramState = string

// developer为开发版；trial为体验版；formal为正式版；默认为正式版
const (
	MiniprogramStateDeveloper = "developer"
	MiniprogramStateTrial     = "trial"
	MiniprogramStateFormal    = "formal"
)

// SubscribeMessageData 订阅消息模板数据
type SubscribeMessageData map[string]struct {
	Value string `json:"value"`
}

// Send 发送订阅消息
//
// token access_token
func (sm *SubscribeMessage) Send(token string) (*CommonError, error) {
	api := baseURL + apiSendSubscribeMessage
	return sm.send(api, token)
}

func (sm *SubscribeMessage) send(api, token string) (*CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := &CommonError{}
	if err := postJSON(api, sm, res); err != nil {
		return nil, err
	}

	return res, nil
}
