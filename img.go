package weapp

import "github.com/medivhzhan/weapp/v3/request"

const (
	apiAICrop          = "/cv/img/aicrop"
	apiScanQRCode      = "/cv/img/qrcode"
	apiSuperResolution = "/cv/img/superResolution"
)

// AICropResponse 图片智能裁剪后的返回数据
type AICropResponse struct {
	request.CommonError
	Results []struct {
		CropLeft   uint `json:"crop_left"`
		CropTop    uint `json:"crop_top"`
		CropRight  uint `json:"crop_right"`
		CropBottom uint `json:"crop_bottom"`
	} `json:"results"`
	IMGSize struct {
		Width  uint `json:"w"`
		Height uint `json:"h"`
	} `json:"img_size"`
}

// AICrop 本接口提供基于小程序的图片智能裁剪能力。
func (cli *Client) AICrop(filename string) (*AICropResponse, error) {
	api := baseURL + apiAICrop

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.aiCrop(api, token, filename)
}

func (cli *Client) aiCrop(api, token, filename string) (*AICropResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(AICropResponse)
	if err := cli.request.FormPostWithFile(url, "img", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}

// AICropByURL 本接口提供基于小程序的图片智能裁剪能力。
func (cli *Client) AICropByURL(url string) (*AICropResponse, error) {
	api := baseURL + apiAICrop

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.aiCropByURL(api, token, url)
}

func (cli *Client) aiCropByURL(api, token, imgURL string) (*AICropResponse, error) {
	queries := requestQueries{
		"access_token": token,
		"img_url":      imgURL,
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(AICropResponse)
	if err := cli.request.Post(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// QRCodePoint 二维码角的位置
type QRCodePoint struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

// ScanQRCodeResponse 小程序的条码/二维码识别后的返回数据
type ScanQRCodeResponse struct {
	request.CommonError
	CodeResults []struct {
		TypeName string `json:"type_name"`
		Data     string `json:"data"`
		Position struct {
			LeftTop     QRCodePoint `json:"left_top"`
			RightTop    QRCodePoint `json:"right_top"`
			RightBottom QRCodePoint `json:"right_bottom"`
			LeftBottom  QRCodePoint `json:"left_bottom"`
		} `json:"pos"`
	} `json:"code_results"`
	IMGSize struct {
		Width  uint `json:"w"`
		Height uint `json:"h"`
	} `json:"img_size"`
}

// ScanQRCode 本接口提供基于小程序的条码/二维码识别的API。
func (cli *Client) ScanQRCode(filename string) (*ScanQRCodeResponse, error) {
	api := baseURL + apiScanQRCode

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.scanQRCode(api, token, filename)
}

func (cli *Client) scanQRCode(api, token, filename string) (*ScanQRCodeResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(ScanQRCodeResponse)
	if err := cli.request.FormPostWithFile(url, "img", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ScanQRCodeByURL 把网络文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
func (cli *Client) ScanQRCodeByURL(imgURL string) (*ScanQRCodeResponse, error) {
	api := baseURL + apiScanQRCode

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.scanQRCodeByURL(api, token, imgURL)
}

func (cli *Client) scanQRCodeByURL(api, token, imgURL string) (*ScanQRCodeResponse, error) {
	queries := requestQueries{
		"access_token": token,
		"img_url":      imgURL,
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(ScanQRCodeResponse)
	if err := cli.request.Post(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SuperResolutionResponse 图片高清化后的返回数据
type SuperResolutionResponse struct {
	request.CommonError
	MediaID string `json:"media_id"`
}

// SuperResolution 本接口提供基于小程序的图片高清化能力。
func (cli *Client) SuperResolution(filename string) (*SuperResolutionResponse, error) {
	api := baseURL + apiSuperResolution

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.superResolution(api, token, filename)
}

func (cli *Client) superResolution(api, token, filename string) (*SuperResolutionResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(SuperResolutionResponse)
	if err := cli.request.FormPostWithFile(url, "img", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SuperResolutionByURL 把网络文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
func (cli *Client) SuperResolutionByURL(imgURL string) (*SuperResolutionResponse, error) {
	api := baseURL + apiSuperResolution

	token, err := cli.AccessToken()
	if err != nil {
		return nil, err
	}

	return cli.superResolutionByURL(api, token, imgURL)
}

func (cli *Client) superResolutionByURL(api, token, imgURL string) (*SuperResolutionResponse, error) {
	queries := requestQueries{
		"access_token": token,
		"img_url":      imgURL,
	}

	url, err := request.EncodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(SuperResolutionResponse)
	if err := cli.request.Post(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}
