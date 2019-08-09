package weapp

import (
	"strconv"
	"strings"
)

const (
	apiGetTemplateMessageList     = "/cgi-bin/wxopen/template/library/list"
	apiGetTempalteMessageDetail   = "/cgi-bin/wxopen/template/library/get"
	apiAddTemplateMessage         = "/cgi-bin/wxopen/template/add"
	apiSelvesTemplateMessageList  = "/cgi-bin/wxopen/template/list"
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

// Template 消息模板
type Template struct {
	CommonError
	ID         string `json:"id,omitempty"`
	TemplateID string `json:"template_id,omitempty"`
	Title      string `json:"title"`
	Content    string `json:"content,omitempty"`
	Example    string `json:"example,omitempty"`

	KeywordList []KeywordItem `json:"keyword_list,omitempty"`
}

// GetTemplateListResponse 获取模板列表返回的数据
type GetTemplateListResponse struct {
	CommonError
	List       []Template `json:"list"`
	TotalCount uint       `json:"total_count"`
}

// List 获取小程序模板库标题列表
//
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func List(offset uint, count uint, token string) (*GetTemplateListResponse, error) {
	return templates(baseURL+apiGetTemplateMessageList, offset, count, token)
}

// Selves 获取帐号下已存在的模板列表
//
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func Selves(offset uint, count uint, token string) (*GetTemplateListResponse, error) {
	return templates(baseURL+apiSelvesTemplateMessageList, offset, count, token)
}

// 获取模板列表
//
// api 开始获取位置 从0开始
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func templates(api string, offset, count uint, token string) (*GetTemplateListResponse, error) {

	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"offset": offset,
		"count":  count,
	}

	res := new(GetTemplateListResponse)
	if err := postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Get 获取模板库某个模板标题下关键词库
//
// id 模板ID
// token 微信 access_token
func Get(id, token string) ([]KeywordItem, error) {
	api, err := tokenAPI(baseURL+apiGetTempalteMessageDetail, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"id": id,
	}

	res := new(Template)
	if err = postJSON(api, params, res); err != nil {
		return nil, err
	}

	return res.KeywordList, nil
}

// Add 组合模板并添加至帐号下的个人模板库
//
// id 模板ID
// token 微信 aceess_token
// keywordIDList 关键词 ID 列表
// 返回新建模板ID和错误信息
func Add(id, token string, keywordIDList []uint) (string, error) {
	api, err := tokenAPI(baseURL+apiAddTemplateMessage, token)
	if err != nil {
		return "", err
	}

	var list []string
	for _, v := range keywordIDList {
		list = append(list, strconv.Itoa(int(v)))
	}

	params := requestParams{
		"id":              id,
		"keyword_id_list": "[" + strings.Join(list, ",") + "]",
	}

	res := new(Template)
	err = postJSON(api, params, res)
	if err != nil {
		return "", err
	}

	return res.TemplateID, nil
}

// DeleteTempalteMessage 删除帐号下的某个模板
//
// id 模板ID
// token 微信 aceess_token
func DeleteTempalteMessage(id, token string) error {
	api, err := tokenAPI(baseURL+apiDeleteTemplateMessage, token)
	if err != nil {
		return err
	}

	params := requestParams{
		"template_id": id,
	}

	res := new(CommonError)
	err = postJSON(api, params, res)
	if err != nil {
		return err
	}

	return nil
}

// Message 模版消息体
type Message map[string]interface{}

// SendTemplateMessage 发送模板消息
//
// openid 接收者（用户）的 openid
// template 所需下发的模板消息的id
// page 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
// formID 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
// data 模板内容，不填则下发空模板
// emphasisKeyword 模板需要放大的关键词，不填则默认无放大
func SendTemplateMessage(openid, template, page, formID string, data Message, emphasisKeyword, token string) error {
	api, err := tokenAPI(baseURL+apiSendTemplateMessage, token)
	if err != nil {
		return err
	}

	for key := range data {
		data[key] = Message{"value": data[key]}
	}

	params := requestParams{
		"touser":           openid,
		"template_id":      template,
		"page":             page,
		"form_id":          formID,
		"data":             data,
		"emphasis_keyword": emphasisKeyword,
	}

	res := new(CommonError)
	err = postJSON(api, params, res)
	if err != nil {
		return err
	}

	return nil
}
