package ocr

import "github.com/medivhzhan/weapp/v3/request"

const apiPrintedText = "/cv/ocr/comm"

type PrintedTextResponse struct {
	request.CommonError
	Items []struct {
		Text     string          `json:"text"`
		Position LicensePosition `json:"pos"`
	} `json:"items"` //	识别结果
	ImgSize Size `json:"img_size"` //	图片大小
}

// 本接口提供基于小程序的通用印刷体 OCR 识别
// 通过图片链接识别
func (cli *OCR) PrintedTextByURL(cardURL string, mode RecognizeMode) (*PrintedTextResponse, error) {
	res := new(PrintedTextResponse)
	err := cli.ocrByURL(apiPrintedText, cardURL, mode, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 本接口提供基于小程序的通用印刷体 OCR 识别
// 通过图片文件识别
func (cli *OCR) PrintedTextByFile(filename string, mode RecognizeMode) (*PrintedTextResponse, error) {
	res := new(PrintedTextResponse)
	err := cli.ocrByFile(apiPrintedText, filename, mode, res)
	if err != nil {
		return nil, err
	}

	return res, err
}
