package ocr

import "github.com/medivhzhan/weapp/v3/request"

const apiBankcard = "/cv/ocr/bankcard"

type BankcardResponse struct {
	request.CommonError
	Number string `json:"number"` // 银行卡号
}

// 本接口提供基于小程序的银行卡 OCR 识别
// 通过图片链接识别
func (cli *OCR) BankcardByURL(cardURL string, mode RecognizeMode) (*BankcardResponse, error) {
	res := new(BankcardResponse)
	err := cli.ocrByURL(apiBankcard, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 本接口提供基于小程序的银行卡 OCR 识别
// 通过图片文件识别
func (cli *OCR) BankcardByFile(filename string, mode RecognizeMode) (*BankcardResponse, error) {
	res := new(BankcardResponse)
	err := cli.ocrByFile(apiBankcard, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}
