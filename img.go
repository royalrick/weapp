package weapp

const (
	apiScanQRCode      = "/cv/img/qrcode"
	apiSuperResolution = "/cv/img/superResolution"
)

// QRCodePoint 二维码角的位置
type QRCodePoint struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}

// ScanQRCodeResponse 小程序的条码/二维码识别后的返回数据
type ScanQRCodeResponse struct {
	CommonError
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
func ScanQRCode(token, filename string) (*ScanQRCodeResponse, error) {
	api := baseURL + apiScanQRCode
	return scanQRCode(api, token, filename)
}

func scanQRCode(api, token, filename string) (*ScanQRCodeResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(ScanQRCodeResponse)
	if err := postFormByFile(url, "img", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ScanQRCodeByURL 把网络文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
func ScanQRCodeByURL(token, imgURL string) (*ScanQRCodeResponse, error) {
	api := baseURL + apiScanQRCode
	return scanQRCodeByURL(api, token, imgURL)
}

func scanQRCodeByURL(api, token, imgURL string) (*ScanQRCodeResponse, error) {
	queries := requestQueries{
		"access_token": token,
		"img_url":      imgURL,
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(ScanQRCodeResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SuperResolutionResponse 图片高清化后的返回数据
type SuperResolutionResponse struct {
	CommonError
	MediaID string `json:"media_id"`
}

// SuperResolution 本接口提供基于小程序的图片高清化能力。
func SuperResolution(token, filename string) (*SuperResolutionResponse, error) {
	api := baseURL + apiSuperResolution
	return superResolution(api, token, filename)
}

func superResolution(api, token, filename string) (*SuperResolutionResponse, error) {
	url, err := tokenAPI(api, token)
	if err != nil {
		return nil, err
	}

	res := new(SuperResolutionResponse)
	if err := postFormByFile(url, "img", filename, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SuperResolutionByURL 把网络文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
func SuperResolutionByURL(token, imgURL string) (*SuperResolutionResponse, error) {
	api := baseURL + apiSuperResolution
	return superResolutionByURL(api, token, imgURL)
}

func superResolutionByURL(api, token, imgURL string) (*SuperResolutionResponse, error) {
	queries := requestQueries{
		"access_token": token,
		"img_url":      imgURL,
	}

	url, err := encodeURL(api, queries)
	if err != nil {
		return nil, err
	}

	res := new(SuperResolutionResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}
