// Package template 模版消息
package template

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/util"
)

const (
	listAPI   = "/cgi-bin/wxopen/template/library/list"
	detailAPI = "/cgi-bin/wxopen/template/library/get"
	addAPI    = "/cgi-bin/wxopen/template/add"
	selvesAPI = "/cgi-bin/wxopen/template/list"
	deleteAPI = "/cgi-bin/wxopen/template/del"
	sendAPI   = "/cgi-bin/message/wxopen/template/send"
)

// KeywordItem 关键字
type KeywordItem struct {
	KeywordID uint   `json:"keyword_id"`
	Name      string `json:"name"`
	Example   string `json:"example"`
}

// Template 消息模板
type Template struct {
	weapp.Response
	ID         string `json:"id,omitempty"`
	TemplateID string `json:"template_id,omitempty"`
	Title      string `json:"title"`
	Content    string `json:"content,omitempty"`
	Example    string `json:"example,omitempty"`

	KeywordList []KeywordItem `json:"keyword_list,omitempty"`
}

// Templates 获取模板列表返回的数据
type Templates struct {
	weapp.Response
	List       []Template `json:"list"`
	TotalCount uint       `json:"total_count"`
}

// List 获取小程序模板库标题列表
//
// @offset 开始获取位置 从0开始
// @count 获取记录条数 最大为20
// @token 微信 access_token
func List(offset uint, count uint, token string) (list []Template, total uint, err error) {
	return templates(weapp.BaseURL+listAPI, offset, count, token)
}

// Selves 获取帐号下已存在的模板列表
//
// @offset 开始获取位置 从0开始
// @count 获取记录条数 最大为20
// @token 微信 access_token
func Selves(offset uint, count uint, token string) (list []Template, total uint, err error) {
	return templates(weapp.BaseURL+selvesAPI, offset, count, token)
}

// 获取模板列表
//
// @api 开始获取位置 从0开始
// @offset 开始获取位置 从0开始
// @count 获取记录条数 最大为20
// @token 微信 access_token
func templates(api string, offset, count uint, token string) (list []Template, total uint, err error) {

	if count > 20 {
		err = errors.New("'count' cannot be great than 20")
		return
	}

	api, err = util.TokenAPI(api, token)
	if err != nil {
		return
	}

	body := fmt.Sprintf(`{"offset": "%v","count":"%v"}`, offset, count)

	res, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(weapp.WeChatServerError)
		return
	}

	var data Templates
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return
	}

	if data.Errcode != 0 {
		err = errors.New(data.Errmsg)
		return
	}

	list = data.List
	total = data.TotalCount
	return
}

// Get 获取模板库某个模板标题下关键词库
//
// @id 模板ID
// @token 微信 access_token
func Get(id, token string) (keywords []KeywordItem, err error) {
	api, err := util.TokenAPI(weapp.BaseURL+detailAPI, token)
	if err != nil {
		return
	}

	body := fmt.Sprintf(`{"id": "%s"}`, id)

	res, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(weapp.WeChatServerError)
		return
	}

	var data Template
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return
	}

	if data.Errcode != 0 {
		err = errors.New(data.Errmsg)
		return
	}

	keywords = data.KeywordList
	return
}

// Add 组合模板并添加至帐号下的个人模板库
//
// @id 模板ID
// @token 微信 aceess_token
// @keywordIDList 关键词 ID 列表
// 返回新建模板ID和错误信息
func Add(id, token string, keywordIDList []uint) (string, error) {
	api, err := util.TokenAPI(weapp.BaseURL+addAPI, token)
	if err != nil {
		return "", err
	}

	var list []string
	for _, v := range keywordIDList {
		list = append(list, strconv.Itoa(int(v)))
	}

	body := fmt.Sprintf(`{"id": "%s", "keyword_id_list": [%s]}`, id, strings.Join(list, ","))

	res, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(weapp.WeChatServerError)
	}

	var tmp Template
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		return "", err
	}

	if tmp.Errcode != 0 {
		return "", errors.New(tmp.Errmsg)
	}

	return tmp.TemplateID, nil
}

// Delete 删除帐号下的某个模板
//
// @id 模板ID
// @token 微信 aceess_token
func Delete(id, token string) error {
	api, err := util.TokenAPI(weapp.BaseURL+deleteAPI, token)
	if err != nil {
		return err
	}

	body := fmt.Sprintf(`{"template_id": "%s"}`, id)

	res, err := http.Post(api, "application/json", strings.NewReader(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(weapp.WeChatServerError)
		return err
	}

	var data weapp.Response
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	if data.Errcode != 0 {
		return errors.New(data.Errmsg)
	}

	return nil
}

// Message 模版消息体
type Message map[string]interface{}

// Send 发送模板消息
//
// @openid 接收者（用户）的 openid
// @template 所需下发的模板消息的id
// @page 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
// @formID 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
// @data 模板内容，不填则下发空模板
// @emphasisKeyword 模板需要放大的关键词，不填则默认无放大
func Send(openid, template, page, formID string, data Message, emphasisKeyword, token string) error {
	api, err := util.TokenAPI(weapp.BaseURL+sendAPI, token)
	if err != nil {
		return err
	}

	for key := range data {
		data[key] = Message{"value": data[key]}
	}

	body := map[string]interface{}{
		"touser":           openid,
		"template_id":      template,
		"page":             page,
		"form_id":          formID,
		"data":             data,
		"emphasis_keyword": emphasisKeyword,
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res, err := http.Post(api, "application/json", strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(weapp.WeChatServerError)
		return err
	}

	var resp weapp.Response
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	if resp.Errcode != 0 {
		return errors.New(resp.Errmsg)
	}

	return nil
}
