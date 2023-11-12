package ocr

import (
	"github.com/medivhzhan/weapp/v3/request"
)

type OCR struct {
	request *request.Request
	// 组成完整的 URL 地址
	// 默认包含 AccessToken
	combineURI func(url string, req interface{}, withToken bool) (string, error)
}

func NewOCR(request *request.Request, combineURI func(url string, req interface{}, withToken bool) (string, error)) *OCR {
	sm := OCR{
		request:    request,
		combineURI: combineURI,
	}

	return &sm
}

// 证件点位
type Point struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

// 证件尺寸
type Size struct {
	H uint `json:"h"`
	W uint `json:"w"`
}

// LicensePosition 证件位置
type LicensePosition struct {
	LeftTop     Point `json:"left_top"`
	RightTop    Point `json:"right_top"`
	RightBottom Point `json:"right_bottom"`
	LeftBottom  Point `json:"left_bottom"`
}

// RecognizeMode 图片识别模式
type RecognizeMode = string

// 所有图片识别模式
const (
	RecognizeModePhoto RecognizeMode = "photo" // 拍照模式
	RecognizeModeScan  RecognizeMode = "scan"  // 扫描模式
)

type ocrByFileRequest struct {
	Type RecognizeMode `query:"type"`
}

func (cli *OCR) ocrByFile(api, filename string, mode RecognizeMode, response interface{}) error {

	req := ocrByFileRequest{
		Type: mode,
	}

	url, err := cli.combineURI(api, &req, true)
	if err != nil {
		return err
	}

	if err := cli.request.FormPostWithFile(url, "img", filename, response); err != nil {
		return err
	}

	return nil
}

type ocrByURLRequest struct {
	Type   RecognizeMode `query:"type"`
	ImgUrl RecognizeMode `query:"img_url"`
}

func (cli *OCR) ocrByURL(api, cardURL string, mode RecognizeMode, response interface{}) error {

	req := ocrByURLRequest{
		Type:   mode,
		ImgUrl: cardURL,
	}

	url, err := cli.combineURI(api, &req, true)
	if err != nil {
		return err
	}

	if err := cli.request.Post(url, nil, response); err != nil {
		return err
	}

	return nil
}
