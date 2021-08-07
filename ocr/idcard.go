package ocr

import "github.com/medivhzhan/weapp/v3/request"

const apiIDCard = "/cv/ocr/idcard"

// CardType 卡片方向
type CardType = string

// 所有卡片方向
const (
	CardTypeFront = "Front" // 正面
	CardTypeBack  = "Back"  // 背面
)

type IDCardResponse struct {
	request.CommonError
	Type      CardType `json:"type"`       // 正面或背面，Front / Back
	ValidDate string   `json:"valid_date"` // 有效期
}

// 本接口提供基于小程序的身份证 OCR 识别
// 通过图片链接识别
func (cli *OCR) IDCardByURL(cardURL string, mode RecognizeMode) (*IDCardResponse, error) {
	res := new(IDCardResponse)
	err := cli.ocrByURL(apiIDCard, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 本接口提供基于小程序的身份证 OCR 识别
// 通过图片文件识别
func (cli *OCR) IDCardByFile(filename string, mode RecognizeMode) (*IDCardResponse, error) {
	res := new(IDCardResponse)
	err := cli.ocrByFile(apiIDCard, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}
