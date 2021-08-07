package ocr

import "github.com/medivhzhan/weapp/v3/request"

const apiBusinessLicense = "/cv/ocr/bizlicense"

type BusinessLicenseResponse struct {
	request.CommonError
	RegNum              string `json:"reg_num"`              //	注册号
	Serial              string `json:"serial"`               //	编号
	LegalRepresentative string `json:"legal_representative"` //	法定代表人姓名
	EnterpriseName      string `json:"enterprise_name"`      //	企业名称
	TypeOfOrganization  string `json:"type_of_organization"` //	组成形式
	Address             string `json:"address"`              //	经营场所/企业住所
	TypeOfEnterprise    string `json:"type_of_enterprise"`   //	公司类型
	BusinessScope       string `json:"business_scope"`       //	经营范围
	RegisteredCapital   string `json:"registered_capital"`   //	注册资本
	PaidInCapital       string `json:"paid_in_capital"`      //	实收资本
	ValidPeriod         string `json:"valid_period"`         //	营业期限
	RegisteredDate      string `json:"registered_date"`      //	注册日期/成立日期
	CertPosition        struct {
		Position LicensePosition `json:"pos"`
	} `json:"cert_position"` //	营业执照位置
	ImgSize Size `json:"img_size"` //	图片大小
}

// 本接口提供基于小程序的营业执照 OCR 识别
// 通过图片链接识别
func (cli *OCR) BusinessLicenseByURL(cardURL string, mode RecognizeMode) (*BusinessLicenseResponse, error) {
	res := new(BusinessLicenseResponse)
	err := cli.ocrByURL(apiBusinessLicense, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 本接口提供基于小程序的营业执照 OCR 识别
// 通过图片文件识别
func (cli *OCR) BusinessLicenseByFile(filename string, mode RecognizeMode) (*BusinessLicenseResponse, error) {
	res := new(BusinessLicenseResponse)
	err := cli.ocrByFile(apiBusinessLicense, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}
