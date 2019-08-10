package weapp

import (
	"strconv"
	"strings"
)

const (
	apiGetTemplateLibraryList     = "/cgi-bin/wxopen/template/library/list"
	apiGetTemplateLibraryByID     = "/cgi-bin/wxopen/template/library/get"
	apiAddTemplateMessage         = "/cgi-bin/wxopen/template/add"
	apiGetTemplateList            = "/cgi-bin/wxopen/template/list"
	apiDeleteTemplateMessage      = "/cgi-bin/wxopen/template/del"
	apiSendTemplateMessage        = "/cgi-bin/message/wxopen/template/send"
	apiUniformSendTemplateMessage = "/cgi-bin/message/wxopen/template/uniform_send"
)

// KeywordItem 关键字
type KeywordItem struct {
	KeywordID uint   `json:"keyword_id"`
	Name      string `json:"name"`
	Example   string `json:"example"`
}

// GetTemplateLibraryListResponse 获取模板列表返回的数据
type GetTemplateLibraryListResponse struct {
	CommonError
	List []struct {
		ID    uint `json:"id"`
		Title uint `json:"title"`
	} `json:"list"`
	TotalCount uint `json:"total_count"`
}

// GetTemplateLibraryList 获取小程序模板库标题列表
//
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func GetTemplateLibraryList(offset uint, count uint, token string) (*GetTemplateLibraryListResponse, error) {
	api := baseURL + apiGetTemplateLibraryList

	response := new(GetTemplateLibraryListResponse)
	err := getTemplateList(api, offset, count, token, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetTemplateListResponse 获取模板列表返回的数据
type GetTemplateListResponse struct {
	CommonError
	List []struct {
		ID      uint `json:"template_id"`
		Title   uint `json:"title"`
		Content uint `json:"content"`
		Example uint `json:"example"`
	} `json:"list"`
}

// GetTemplateList 获取帐号下已存在的模板列表
//
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func GetTemplateList(offset uint, count uint, token string) (*GetTemplateListResponse, error) {
	api := baseURL + apiGetTemplateList

	response := new(GetTemplateListResponse)
	err := getTemplateList(api, offset, count, token, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// 获取模板列表
//
// api 开始获取位置 从0开始
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func getTemplateList(api string, offset, count uint, token string, response interface{}) error {

	url, err := tokenAPI(api, token)
	if err != nil {
		return err
	}

	params := requestParams{
		"offset": offset,
		"count":  count,
	}

	if err := postJSON(url, params, response); err != nil {
		return err
	}

	return nil
}

// GetTemplateLibraryByIDResponse 消息模板
type GetTemplateLibraryByIDResponse struct {
	CommonError
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	KeywordList []KeywordItem `json:"keyword_list"`
}

// GetTemplateLibraryByID 获取模板库某个模板标题下关键词库
//
// id 模板ID
// token 微信 access_token
func GetTemplateLibraryByID(id, token string) (*GetTemplateLibraryByIDResponse, error) {
	api := baseURL + apiGetTemplateLibraryByID
	return getTemplateLibraryByID(id, token, api)
}

func getTemplateLibraryByID(id, token, api string) (*GetTemplateLibraryByIDResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"id": id,
	}

	res := new(GetTemplateLibraryByIDResponse)
	if err = postJSON(url, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// AddTemplateResponse 添加模版消息返回数据
type AddTemplateResponse struct {
	CommonError
	ID string `json:"id"`
}

// AddTemplate 组合模板并添加至帐号下的个人模板库
//
// id 模板ID
// token 微信 aceess_token
// keywordIDList 关键词 ID 列表
// 返回新建模板ID和错误信息
func AddTemplate(id, token string, keywordIDList []uint) (*AddTemplateResponse, error) {
	api := baseURL + apiAddTemplateMessage
	return addTemplate(id, token, keywordIDList, api)
}

func addTemplate(id, token string, keywordIDList []uint, api string) (*AddTemplateResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	var list []string
	for _, v := range keywordIDList {
		list = append(list, strconv.Itoa(int(v)))
	}

	params := requestParams{
		"id":              id,
		"keyword_id_list": "[" + strings.Join(list, ",") + "]",
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
// id 模板ID
// token 微信 aceess_token
func DeleteTemplate(id, token string) (*CommonError, error) {
	api := baseURL + apiDeleteTemplateMessage
	return deleteTemplate(id, token, api)
}

func deleteTemplate(id, token, api string) (*CommonError, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"template_id": id,
	}

	res := new(CommonError)
	err = postJSON(api, params, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// TemplateMessage 模版消息体
type TemplateMessage map[string]interface{}

// SendTemplateMessage 发送模板消息
//
// openID 接收者（用户）的 openid
// template 所需下发的模板消息的id
// page 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
// formID 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
// data 模板内容，不填则下发空模板
// emphasisKeyword 模板需要放大的关键词，不填则默认无放大
func SendTemplateMessage(openID, template, page, formID string, msg TemplateMessage, emphasisKeyword, token string) error {
	api := baseURL + apiSendTemplateMessage
	return sendTemplateMessage(openID, template, page, formID, msg, emphasisKeyword, token, api)
}

func sendTemplateMessage(openID, template, page, formID string, msg TemplateMessage, emphasisKeyword, token, api string) error {
	url, err := tokenAPI(api, token)
	if err != nil {
		return err
	}

	for key := range msg {
		msg[key] = TemplateMessage{"value": msg[key]}
	}

	params := requestParams{
		"touser":           openID,
		"template_id":      template,
		"page":             page,
		"form_id":          formID,
		"data":             msg,
		"emphasis_keyword": emphasisKeyword,
	}

	res := new(CommonError)
	err = postJSON(url, params, res)
	if err != nil {
		return err
	}

	return nil
}
