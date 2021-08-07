package ocr

import "github.com/medivhzhan/weapp/v3/request"

const apiVehicleLicense = "/cv/ocr/driving"

type VehicleLicenseResponse struct {
	request.CommonError
	VehicleType                string `json:"vehicle_type"`
	Owner                      string `json:"owner"`
	Addr                       string `json:"addr"`
	UseCharacter               string `json:"use_character"`
	Model                      string `json:"model"`
	Vin                        string `json:"vin"`
	EngineNum                  string `json:"engine_num"`
	RegisterDate               string `json:"register_date"`
	IssueDate                  string `json:"issue_date"`
	PlateNumB                  string `json:"plate_num_b"`
	Record                     string `json:"record"`
	PassengersNum              string `json:"passengers_num"`
	TotalQuality               string `json:"total_quality"`
	TotalprepareQualityQuality string `json:"totalprepare_quality_quality"`
}

// 本接口提供基于小程序的行驶证 OCR 识别
// 通过图片链接识别
func (cli *OCR) VehicleLicenseByURL(cardURL string, mode RecognizeMode) (*VehicleLicenseResponse, error) {
	res := new(VehicleLicenseResponse)
	err := cli.ocrByURL(apiVehicleLicense, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 本接口提供基于小程序的行驶证 OCR 识别
// 通过图片文件识别
func (cli *OCR) VehicleLicenseByFile(filename string, mode RecognizeMode) (*VehicleLicenseResponse, error) {
	res := new(VehicleLicenseResponse)
	err := cli.ocrByFile(apiVehicleLicense, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}
