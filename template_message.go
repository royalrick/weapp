package weapp

const (
	apiAddTemplate            = "/cgi-bin/wxopen/template/add"
	apiDeleteTemplate         = "/cgi-bin/wxopen/template/del"
	apiGetTemplateLibraryByID = "/cgi-bin/wxopen/template/library/get"
	apiGetTemplateLibraryList = "/cgi-bin/wxopen/template/library/list"
	apiGetTemplateList        = "/cgi-bin/wxopen/template/list"
	apiSendTemplateMessage    = "/cgi-bin/message/wxopen/template/send"
)

// KeywordItem 关键字
type KeywordItem struct {
	KeywordID uint   `json:"keyword_id"` // 关键词 id，添加模板时需要
	Name      string `json:"name"`       // 关键词内容
	Example   string `json:"example"`    // 关键词内容对应的示例
}

// GetTemplateLibraryListResponse 获取模板列表返回的数据
type GetTemplateLibraryListResponse struct {
	CommonError
	List []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"list"`
	TotalCount uint `json:"total_count"`
}

// GetTemplateLibraryList 获取小程序模板库标题列表
//
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func GetTemplateLibraryList(token string, offset uint, count uint) (*GetTemplateLibraryListResponse, error) {
	api := baseURL + apiGetTemplateLibraryList

	response := new(GetTemplateLibraryListResponse)
	err := getTemplateList(api, token, offset, count, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetTemplateListResponse 获取模板列表返回的数据
type GetTemplateListResponse struct {
	CommonError
	List []struct {
		ID      string `json:"template_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
		Example string `json:"example"`
	} `json:"list"`
}

// GetTemplateList 获取帐号下已存在的模板列表
//
// offset 开始获取位置 从0开始
// count 获取记录条数 最大为20
// token 微信 access_token
func GetTemplateList(token string, offset uint, count uint) (*GetTemplateListResponse, error) {
	api := baseURL + apiGetTemplateList

	response := new(GetTemplateListResponse)
	err := getTemplateList(api, token, offset, count, response)
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
func getTemplateList(api string, token string, offset, count uint, response interface{}) error {

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
	TemplateID string `json:"template_id"`
}

// AddTemplate 组合模板并添加至帐号下的个人模板库
//
// id 模板ID
// token 微信 aceess_token
// keywordIDList 关键词 ID 列表
// 返回新建模板ID和错误信息
func AddTemplate(id, token string, keywordIDList []uint) (*AddTemplateResponse, error) {
	api := baseURL + apiAddTemplate
	return addTemplate(id, token, keywordIDList, api)
}

func addTemplate(id, token string, keywordIDList []uint, api string) (*AddTemplateResponse, error) {
	api, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	params := requestParams{
		"id":              id,
		"keyword_id_list": keywordIDList,
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
	api := baseURL + apiDeleteTemplate
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

// TempMsgKeyword 模板内容关键字
type TempMsgKeyword struct {
	Value string `json:"value"`
}

// TempMsgData 模板内容
type TempMsgData map[string]TempMsgKeyword

// TempMsgSender 模版消息发送器
type TempMsgSender struct {
	ToUser          string      `json:"touser"`           // 必填	接收者（用户）的 openid
	TemplateID      string      `json:"template_id"`      //必填	所需下发的模板消息的id
	Page            string      `json:"page"`             // 选填	点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	FormID          string      `json:"form_id"`          // 必填	表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	Data            TempMsgData `json:"data"`             // 选填	模板内容，不填则下发空模板。具体格式请参考示例。
	EmphasisKeyword string      `json:"emphasis_keyword"` // 选填	模板需要放大的关键词，不填则默认无放大
}

// Send 发送模版消息
func (sender *TempMsgSender) Send(token string) (*CommonError, error) {
	api := baseURL + apiSendTemplateMessage
	return sender.send(api, token)
}

func (sender *TempMsgSender) send(api, token string) (*CommonError, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(CommonError)
	err = postJSON(url, sender, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
