package weapp

import (
	"strconv"

	"github.com/medivhzhan/weapp/v3/request"
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
	request.CommonError
	Pid string `json:"priTmplId"` // 添加至帐号下的模板id，发送小程序订阅消息时所需
}

// AddTemplate 组合模板并添加至帐号下的个人模板库
//
// token 微信 access_token
// tid 模板ID
// desc 服务场景描述，15个字以内
// keywordIDList 关键词 ID 列表
func (cli *Client) AddTemplate(tid, desc string, keywordIDList []int32) (*AddTemplateResponse, error) {
	api := baseURL + apiAddTemplate

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.addTemplate(api, token, tid, desc, keywordIDList)
}

func (cli *Client) addTemplate(api, token, tid, desc string, keywordIDList []int32) (*AddTemplateResponse, error) {
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
	err = cli.request.Post(api, params, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteTemplate 删除帐号下的某个模板
//
// token 微信 access_token
// pid 模板ID
func (cli *Client) DeleteTemplate(pid string) (*request.CommonError, error) {
	api := baseURL + apiDeleteTemplate
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.deleteTemplate(api, token, pid)
}

func (cli *Client) deleteTemplate(api, token, pid string) (*request.CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"priTmplId": pid,
	}

	res := new(request.CommonError)
	err = cli.request.Post(api, params, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTemplateCategoryResponse 删除帐号下的某个模板返回数据
type GetTemplateCategoryResponse struct {
	request.CommonError
	Data []struct {
		ID   int    `json:"id"`   // 类目id，查询公共库模版时需要
		Name string `json:"name"` // 类目的中文名
	} `json:"data"` // 类目列表
}

// GetTemplateCategory 删除帐号下的某个模板
//
// token 微信 access_token
func (cli *Client) GetTemplateCategory() (*GetTemplateCategoryResponse, error) {
	api := baseURL + apiGetTemplateCategory
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getTemplateCategory(token, api)
}

func (cli *Client) getTemplateCategory(token, api string) (*GetTemplateCategoryResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetTemplateCategoryResponse)
	err = cli.request.Get(api, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetPubTemplateKeyWordsByIdResponse 模板标题下的关键词列表
type GetPubTemplateKeyWordsByIdResponse struct {
	request.CommonError
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
func (cli *Client) GetPubTemplateKeyWordsById(tid string) (*GetPubTemplateKeyWordsByIdResponse, error) {
	api := baseURL + apiGetPubTemplateKeyWordsById
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getPubTemplateKeyWordsById(api, token, tid)
}

func (cli *Client) getPubTemplateKeyWordsById(api, token, tid string) (*GetPubTemplateKeyWordsByIdResponse, error) {
	queries := requestQueries{
		"access_token": token,
		"tid":          tid,
	}
	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(GetPubTemplateKeyWordsByIdResponse)
	if err = cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPubTemplateTitleListResponse 帐号所属类目下的公共模板标题
type GetPubTemplateTitleListResponse struct {
	request.CommonError
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
func (cli *Client) GetPubTemplateTitleList(ids string, start, limit int) (*GetPubTemplateTitleListResponse, error) {
	api := baseURL + apiGetPubTemplateTitleList
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}
	return cli.getPubTemplateTitleList(api, token, ids, start, limit)
}

func (cli *Client) getPubTemplateTitleList(api, token, ids string, start, limit int) (*GetPubTemplateTitleListResponse, error) {

	queries := requestQueries{
		"access_token": token,
		"ids":          ids,
		"start":        strconv.Itoa(start),
		"limit":        strconv.Itoa(limit),
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(GetPubTemplateTitleListResponse)
	if err := cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetTemplateListResponse 获取模板列表返回的数据
type GetTemplateListResponse struct {
	request.CommonError
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
func (cli *Client) GetTemplateList() (*GetTemplateListResponse, error) {
	api := baseURL + apiGetTemplateList
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.getTemplateList(api, token)
}

func (cli *Client) getTemplateList(api, token string) (*GetTemplateListResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(GetTemplateListResponse)
	if err := cli.request.Get(url, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	ToUser           string           `json:"touser"`
	TemplateID       string           `json:"template_id"`
	Page             string           `json:"page,omitempty"`
	MiniprogramState MiniprogramState `json:"miniprogram_state,omitempty"`
	Data             string           `json:"data"`
}

// MiniprogramState 跳转小程序类型
type MiniprogramState = string

// developer为开发版；trial为体验版；formal为正式版；默认为正式版
const (
	MiniprogramStateDeveloper MiniprogramState = "developer"
	MiniprogramStateTrial     MiniprogramState = "trial"
	MiniprogramStateFormal    MiniprogramState = "formal"
)

// Send 发送订阅消息
//
// token access_token
func (cli *Client) SendSubscribeMsg(msg *SubscribeMessage) (*request.CommonError, error) {
	api := baseURL + apiSendSubscribeMessage
	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.sendSubscribeMsg(api, token, msg)
}

func (cli *Client) sendSubscribeMsg(api, token string, msg *SubscribeMessage) (*request.CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(api, msg, res); err != nil {
		return nil, err
	}

	return res, nil
}
