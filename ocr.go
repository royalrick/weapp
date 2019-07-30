package weapp

const (
	apiBankcard       = "/cv/ocr/bankcard"
	apiDrivingLicense = "/cv/ocr/driving"
	apiIDCard         = "/cv/ocr/idcard"
)

// RecognizeMode 图片识别模式
type RecognizeMode = string

// 所有图片识别模式
const (
	RecognizeModePhoto RecognizeMode = "photo" // 拍照模式
	RecognizeModeScan  RecognizeMode = "scan"  // 扫描模式
)

// BankCardResponse 识别银行卡返回数据
type BankCardResponse struct {
	CommonError
	Number string `json:"number"` // 银行卡号
}

// BankCardByURL 通过URL识别银行卡
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// url 要检测的图片 url，传这个则不用传 img 参数。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func BankCardByURL(token, url string, mode RecognizeMode) (*BankCardResponse, error) {
	params := map[string]string{
		"access_token": token,
		"type":         mode,
		"img_url":      url,
	}

	url, err := encodeURL(baseURL+apiBankcard, params)
	if err != nil {
		return nil, err
	}

	res := new(BankCardResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// BankCardByFile 通过文件识别银行卡
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// img form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传递 img_url。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func BankCardByFile(token, img string, mode RecognizeMode) (*BankCardResponse, error) {
	params := map[string]string{
		"access_token": token,
		"type":         mode,
		"img":          img,
	}

	url, err := encodeURL(baseURL+apiDrivingLicense, params)
	if err != nil {
		return nil, err
	}

	res := new(BankCardResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CardType 卡片方向
type CardType = string

// 所有卡片方向
const (
	CardTypeFront = "Front" // 正面
	CardTypeBack  = "Back"  // 背面
)

// CardResponse 识别卡片返回数据
type CardResponse struct {
	CommonError
	Type      CardType `json:"type"`       // 正面或背面，Front / Back
	ValidDate string   `json:"valid_date"` // 有效期
}

// DrivingLicenseResponse 识别行驶证返回数据
type DrivingLicenseResponse = CardResponse

// DrivingLicenseByURL 通过URL识别行驶证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// url 要检测的图片 url，传这个则不用传 img 参数。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func DrivingLicenseByURL(token, url string, mode RecognizeMode) (*DrivingLicenseResponse, error) {
	params := map[string]string{
		"access_token": token,
		"type":         mode,
		"img_url":      url,
	}

	url, err := encodeURL(baseURL+apiDrivingLicense, params)
	if err != nil {
		return nil, err
	}

	res := new(DrivingLicenseResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// DrivingLicenseByFile 通过文件识别行驶证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// img form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传递 img_url。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func DrivingLicenseByFile(token, img string, mode RecognizeMode) (*DrivingLicenseResponse, error) {
	params := map[string]string{
		"access_token": token,
		"type":         mode,
		"img":          img,
	}

	url, err := encodeURL(baseURL+apiDrivingLicense, params)
	if err != nil {
		return nil, err
	}

	res := new(DrivingLicenseResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// IDCardResponse 识别身份证返回数据
type IDCardResponse = CardResponse

// IDCardByURL 通过URL识别身份证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// url 要检测的图片 url，传这个则不用传 img 参数。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func IDCardByURL(token, url string, mode RecognizeMode) (*IDCardResponse, error) {
	params := map[string]string{
		"access_token": token,
		"type":         mode,
		"img_url":      url,
	}

	url, err := encodeURL(baseURL+apiIDCard, params)
	if err != nil {
		return nil, err
	}

	res := new(IDCardResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}

// IDCardByFile 通过文件识别身份证
// 接口限制: 此接口需要提供对应小程序/公众号 appid，开通权限后方可调用。
//
// token 接口调用凭证
// img form-data 中媒体文件标识，有filename、filelength、content-type等信息，传这个则不用传递 img_url。
// mode 图片识别模式，photo（拍照模式）或 scan（扫描模式）
func IDCardByFile(token, img string, mode RecognizeMode) (*IDCardResponse, error) {
	params := map[string]string{
		"access_token": token,
		"type":         mode,
		"img":          img,
	}

	url, err := encodeURL(baseURL+apiIDCard, params)
	if err != nil {
		return nil, err
	}

	res := new(IDCardResponse)
	if err := postJSON(url, nil, res); err != nil {
		return nil, err
	}

	return res, nil
}
