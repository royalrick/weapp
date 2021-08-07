package ocr

import "github.com/medivhzhan/weapp/v3/request"

const apiDriverLicense = "/cv/ocr/drivinglicense"

type DriverLicenseResponse struct {
	request.CommonError
	IDNum        string `json:"id_num"`        // 证号
	Name         string `json:"name"`          // 姓名
	Nationality  string `json:"nationality"`   // 国家
	Sex          string `json:"sex"`           // 性别
	Address      string `json:"address"`       // 地址
	BirthDate    string `json:"birth_date"`    // 出生日期
	IssueDate    string `json:"issue_date"`    // 初次领证日期
	CarClass     string `json:"car_class"`     // 准驾车型
	ValidFrom    string `json:"valid_from"`    // 有效期限起始日
	ValidTo      string `json:"valid_to"`      // 有效期限终止日
	OfficialSeal string `json:"official_seal"` // 印章文构
}

// 本接口提供基于小程序的驾驶证 OCR 识别
// 通过图片链接识别
func (cli *OCR) DriverLicenseByURL(cardURL string, mode RecognizeMode) (*DriverLicenseResponse, error) {
	res := new(DriverLicenseResponse)
	err := cli.ocrByURL(apiDriverLicense, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 本接口提供基于小程序的驾驶证 OCR 识别
// 通过图片文件识别
func (cli *OCR) DriverLicenseByFile(filename string, mode RecognizeMode) (*DriverLicenseResponse, error) {
	res := new(DriverLicenseResponse)
	err := cli.ocrByFile(apiDriverLicense, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}
