package nearby

import (
	"encoding/json"
	"fmt"

	"github.com/medivhzhan/weapp"
)

// apis
const (
	addAPI           = "/wxa/addnearbypoi"
	deleteAPI        = "/wxa/delnearbypoi"
	getListAPI       = "/wxa/getnearbypoilist"
	setShowStatusAPI = "/wxa/setnearbypoishowstatus"
)

// Position 地点
type Position struct {
	PicList           PicList      `json:"pic_list"`           // 门店图片，最多9张，最少1张，上传门店图片如门店外景、环境设施、商品服务等，图片将展示在微信客户端的门店页。图片链接通过文档https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1444738729中的《上传图文消息内的图片获取URL》接口获取。必填，文件格式为bmp、png、jpeg、jpg或gif，大小不超过5M pic_list是字符串，内容是一个json！
	ServiceInfos      ServiceInfos `json:"service_infos"`      // 必服务标签列表 选填，需要填写服务标签ID、APPID、对应服务落地页的path路径，详细字段格式见下方示例
	StoreName         string       `json:"store_name"`         // 门店名字 必填，门店名称需按照所选地理位置自动拉取腾讯地图门店名称，不可修改，如需修改请重现选择地图地点或重新创建地点
	Hour              string       `json:"hour"`               // 营业时间，格式11:11-12:12 必填
	Credential        string       `json:"credential"`         // 资质号 必填, 15位营业执照注册号或9位组织机构代码
	Address           string       `json:"address"`            // 地址 必填
	CompanyName       string       `json:"company_name"`       // 主体名字 必填
	QualificationList string       `json:"qualification_list"` // 证明材料 必填 如果company_name和该小程序主体不一致，需要填qualification_list，详细规则见附近的小程序使用指南-如何证明门店的经营主体跟公众号或小程序帐号主体相关http://kf.qq.com/faq/170401MbUnim17040122m2qY.html
	KFInfo            KFInfo       `json:"kf_info"`            // 客服信息 选填，可自定义服务头像与昵称，具体填写字段见下方示例kf_info pic_list是字符串，内容是一个json！
	PoiID             string       `json:"poi_id"`             // 如果创建新的门店，poi_id字段为空 如果更新门店，poi_id参数则填对应门店的poi_id 选填
}

// PicList 门店图片
type PicList struct {
	List []string `json:"list"`
}

// ServiceInfos 必服务标签列表
type ServiceInfos struct {
	ServiceInfos []ServiceInfo `json:"service_infos"`
}

// ServiceInfo 必服务标签
type ServiceInfo struct {
	ID    uint   `json:"id"`
	Type  uint8  `json:"type"`
	Name  string `json:"name"`
	AppID string `json:"appid"`
	Path  string `json:"path"`
}

// KFInfo // 客服信息
type KFInfo struct {
	OpenKF    bool   `json:"open_kf"`
	KDHeading string `json:"kf_headimg"`
	KFName    string `json:"kf_name"`
}

// PositionResponse response of add position.
type PositionResponse struct {
	weapp.BaseResponse
	Data struct {
		AuditID           string `json:"audit_id"`           //	审核单 ID
		PoiID             string `json:"poi_id"`             //	附近地点 ID
		RelatedCredential string `json:"related_credential"` //	经营资质证件号
	} `json:"data"`
}

// Add 添加地点
// @accessToken  接口调用凭证
func (p *Position) Add(accessToken string) (*PositionResponse, error) {

	pisList, err := json.Marshal(p.PicList)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal picture list to json: %v", err)
	}

	serviceInfos, err := json.Marshal(p.ServiceInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal service info list to json: %v", err)
	}

	kfInfo, err := json.Marshal(p.KFInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal customer service staff info list to json: %v", err)
	}

	params := map[string]string{
		"is_comm_nearby":     "1",
		"pic_list":           string(pisList),
		"service_infos":      string(serviceInfos),
		"store_name":         p.StoreName,
		"hour":               p.Hour,
		"credential":         p.Credential,
		"address":            p.Address,
		"company_name":       p.CompanyName,
		"qualification_list": p.QualificationList,
		"kf_info":            string(kfInfo),
		"poi_id":             p.PoiID,
	}

	api, err := weapp.TokenAPI(weapp.BaseURL+addAPI, accessToken)
	if err != nil {
		return nil, err
	}

	res := new(PositionResponse)
	if err := weapp.PostJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Delete 删除地点
// @accessToken  接口调用凭证
// @id  附近地点 ID
func Delete(accessToken, id string) (*weapp.BaseResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+deleteAPI, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"poi_id": id,
	}

	res := new(weapp.BaseResponse)
	if err := weapp.PostJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PositionList 地点列表
type PositionList struct {
	weapp.BaseResponse
	Data struct {
		LeftApplyNum uint `json:"left_apply_num"` // 剩余可添加地点个数
		MaxApplyNum  uint `json:"max_apply_num"`  // 最大可添加地点个数
		Data         struct {
			List []struct {
				PoiID                string `json:"poi_id"`                // 附近地点 ID
				QualificationAddress string `json:"qualification_address"` // 资质证件地址
				QualificationNum     string `json:"qualification_num"`     // 资质证件证件号
				AuditStatus          int    `json:"audit_status"`          // 地点审核状态
				DisplayStatus        int    `json:"display_status"`        // 地点展示在附近状态
				RefuseReason         string `json:"refuse_reason"`         // 审核失败原因，audit_status=4 时返回
			} `json:"poi_list"` // 地址列表
		} `json:"data"` // 地址列表的 JSON 格式字符串
	} `json:"data"` // 返回数据
}

// GetList 查看地点列表
// @accessToken  接口调用凭证
// @page  起始页id（从1开始计数）
// @pageRows  每页展示个数（最多1000个）
func GetList(accessToken string, page, pageRows uint) (*PositionList, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+getListAPI, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"page":      page,
		"page_rows": pageRows,
	}

	res := new(PositionList)
	if err := weapp.PostJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ShowStatus 展示状态
type ShowStatus = int8

// 所有展示状态
const (
	Hide ShowStatus = 0 // 不展示
	Show            = 1 // 展示
)

// SetShowStatus 展示/取消展示附近小程序
// @accessToken  接口调用凭证
// @poiID  附近地点 ID
// @status  是否展示
func SetShowStatus(accessToken, poiID string, status ShowStatus) (*weapp.BaseResponse, error) {
	api, err := weapp.TokenAPI(weapp.BaseURL+getListAPI, accessToken)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"poi_id": poiID,
		"status": status,
	}

	res := new(weapp.BaseResponse)
	if err := weapp.PostJSON(api, params, res); err != nil {
		return nil, err
	}

	return res, nil
}
