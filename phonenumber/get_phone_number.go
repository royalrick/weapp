package phonenumber

import "github.com/medivhzhan/weapp/v3/request"

const apiGetPhoneNumber = "/wxa/business/getuserphonenumber"

type GetPhoneNumberRequest struct {
	Code string `json:"code"`
}

type GetPhoneNumberResponse struct {
	request.CommonError
	Data struct {
		PhoneNumber     string `json:"phoneNumber"`     //	用户绑定的手机号（国外手机号会有区号）
		PurePhoneNumber string `json:"purePhoneNumber"` //	没有区号的手机号
		CountryCode     string `json:"countryCode"`     //	区号
		Watermark       struct {
			Appid     string `json:"appid"`     // 小程序appid
			Timestamp int64  `json:"timestamp"` // 用户获取手机号操作的时间戳
		} `json:"watermark"` // 数据水印
	} `json:"phone_info"` // 类目列表
}

// code换取用户手机号。 每个code只能使用一次，code的有效期为5min
func (cli *Phonenumber) GetPhoneNumber(req *GetPhoneNumberRequest) (*GetPhoneNumberResponse, error) {

	api, err := cli.combineURI(apiGetPhoneNumber, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(GetPhoneNumberResponse)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
