package weapp

import (
	"testing"
)

func TestDecodeUserinfo(t *testing.T) {
	ssk := "hA4M0cvG4/XLllgSF2uIDA=="
	encryptedData := "ITIEO7DAkGz/yyKkl2T9jGSBYiVQCqdjMYXuiDOVcLqXNkBwajRl4CHVd0NeKYks+NaLBk4b1VvHieJ/VYDnxTS6kXcAzyRfAR7WFfwZvRMfQwlo9yXrvPgFcM30YI46XXYCix+A6174KRX1+GJN77V90YJF+KFYIZsmMaMAqNrzes4DvK7u2162lGxzaQ+2VQmKAT2iIIwHANuTqnZNVK5h6e5kymwA0CuhCy2VGI4gFu6Q8WAmPCMbsZKtXVYIP8oFORZHbXda63jfGQZVtoWBMRVba7omui0vl6lUVyFmqiucXYhQEuLfWyWCNpDuzHsVWdobjgUWAWXhw6MNm3ZewSx7ynWsu+S0ODzU44H+iA/tnZQ5QsnF1AG9gF8i3QczB1d/g8egb0VETOvspXU7CYYj10QfepWxJeSF2JbYfauFMQku/iNoPlDjeYbdsfSSZjFs1UphhEIaA+J6o6XLUqZ86GtdTHXiPJ5iy5o="
	rawData := `{"nickName":"克罗地亚幻想","gender":1,"language":"zh_CN","city":"Chengdu","province":"Sichuan","country":"China","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/NvfKP2XE4RK6Jyn4bsN8oD5dBd6Ker0KbS7sgibicQOoTJp4XtwbBeBxrqTtpHeHib7KiaRl1qpe3jflvoyJg9N8Lw/132"}`
	signature := "3b28beb55152b17f85f513d53153d8ac37563a19"
	iv := "qWE6I82SlAJR9wBK4i2jqg=="

	ui, err := DecodeUserInfo(rawData, encryptedData, signature, iv, ssk)
	if err != nil {
		t.Error(err)
	}

	if ui.Gender != 1 ||
		ui.City != "Chengdu" ||
		ui.Province != "Sichuan" ||
		ui.Country != "China" ||
		ui.Avatar != "https://wx.qlogo.cn/mmopen/vi_32/NvfKP2XE4RK6Jyn4bsN8oD5dBd6Ker0KbS7sgibicQOoTJp4XtwbBeBxrqTtpHeHib7KiaRl1qpe3jflvoyJg9N8Lw/132" ||
		ui.Language != "zh_CN" ||
		ui.Nickname != "克罗地亚幻想" {

		t.Error("结果数据不一致")
	}

	ssk = ssk + "..."
	_, err = DecodeUserInfo(rawData, encryptedData, signature, iv, ssk)
	if err == nil {
		t.Error("错误的数据得到了正确的结果")
	}
}

func TestDecodePhoneNumber(t *testing.T) {
	iv := "F1T12AC9QN965KEG12qbmg=="
	ssk := "hA4M0cvG4/XLllgSF2uIDA=="
	encryptedData := "MYOAFL9fs9wjc/39xw+qfUgGadRSdbwqNFqVOt0v2UZhJjM5Csrvt0uF4GBuPTfBzSZeDkmSVZEWw7Uk5h3Re/igz6tXHrRgbepZj5eoBdsAZNESR/1SIuf936wogXGYlMGOqL+gWWwazFPR7aG6aZYgOLB28pMeOBpVIU0Kv5sI1C6Ot8iOrxIrmY04leO989Sdz33WOdL7eM2hnl4DsQ=="

	phone, err := DecodePhoneNumber(ssk, encryptedData, iv)
	if err != nil {
		t.Error(err)
	}
	if phone.PhoneNumber != "18048574657" ||
		phone.PurePhoneNumber != "18048574657" ||
		phone.CountryCode != "86" {
		t.Error("结果数据不一致")
	}

	ssk = ssk + "..."
	phone, err = DecodePhoneNumber(ssk, encryptedData, iv)
	if err == nil {
		t.Error("错误的数据得到了正确的结果")
	}
}
