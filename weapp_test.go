package weapp

import (
	"testing"
)

func TestDecryptUserInfo(t *testing.T) {
	ssk := "hA4M0cvG4/XLllgSF2uIDA=="
	encryptedData := "ITIEO7DAkGz/yyKkl2T9jGSBYiVQCqdjMYXuiDOVcLqXNkBwajRl4CHVd0NeKYks+NaLBk4b1VvHieJ/VYDnxTS6kXcAzyRfAR7WFfwZvRMfQwlo9yXrvPgFcM30YI46XXYCix+A6174KRX1+GJN77V90YJF+KFYIZsmMaMAqNrzes4DvK7u2162lGxzaQ+2VQmKAT2iIIwHANuTqnZNVK5h6e5kymwA0CuhCy2VGI4gFu6Q8WAmPCMbsZKtXVYIP8oFORZHbXda63jfGQZVtoWBMRVba7omui0vl6lUVyFmqiucXYhQEuLfWyWCNpDuzHsVWdobjgUWAWXhw6MNm3ZewSx7ynWsu+S0ODzU44H+iA/tnZQ5QsnF1AG9gF8i3QczB1d/g8egb0VETOvspXU7CYYj10QfepWxJeSF2JbYfauFMQku/iNoPlDjeYbdsfSSZjFs1UphhEIaA+J6o6XLUqZ86GtdTHXiPJ5iy5o="
	rawData := `{"nickName":"克罗地亚幻想","gender":1,"language":"zh_CN","city":"Chengdu","province":"Sichuan","country":"China","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/NvfKP2XE4RK6Jyn4bsN8oD5dBd6Ker0KbS7sgibicQOoTJp4XtwbBeBxrqTtpHeHib7KiaRl1qpe3jflvoyJg9N8Lw/132"}`
	signature := "3b28beb55152b17f85f513d53153d8ac37563a19"
	iv := "qWE6I82SlAJR9wBK4i2jqg=="

	ui, err := DecryptUserInfo(rawData, encryptedData, signature, iv, ssk)
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
	_, err = DecryptUserInfo(rawData, encryptedData, signature, iv, ssk)
	if err == nil {
		t.Error("错误的数据得到了正确的结果")
	}
}

type encryptedPhoneNumber struct {
	IV         string
	Assert     bool
	SessionKey string
	Ciphertext string
}

var encryptedPhoneNumbers = []encryptedPhoneNumber{
	encryptedPhoneNumber{
		SessionKey: "hA4M0cvG4/XLllgSF2uIDA==",
		Ciphertext: "MYOAFL9fs9wjc/39xw+qfUgGadRSdbwqNFqVOt0v2UZhJjM5Csrvt0uF4GBuPTfBzSZeDkmSVZEWw7Uk5h3Re/igz6tXHrRgbepZj5eoBdsAZNESR/1SIuf936wogXGYlMGOqL+gWWwazFPR7aG6aZYgOLB28pMeOBpVIU0Kv5sI1C6Ot8iOrxIrmY04leO989Sdz33WOdL7eM2hnl4DsQ==",
		IV:         "F1T12AC9QN965KEG12qbmg==",
		Assert:     true,
	},
	encryptedPhoneNumber{
		SessionKey: "k18A8hHN236qkAlTV+SrQQ==",
		Ciphertext: "uTVeF3fPEItGvzAf6TLHiqIHzGztS/MjeF0HndOSGDWqsc5t4R6HDN2tUF+4aCzVYRgJwIWNeGKeHSjQ85jHNjdQHOfYu60l1Eoq/lL3vd31NT4bMVo2wFqoQ2jOdDi/0w+mfTvmsxk1WcdECS3uLeZEJ3N+9sFnyeWwoS5qyEDqMjEVGX1Rflp+SeBkYuo+gnWzBLPIdEJFbAe/uPM91w==",
		IV:         "/WDoLmVaW89+zltIsBxNCQ==",
		Assert:     true,
	},

	encryptedPhoneNumber{
		SessionKey: "23TSo+wX/eDxwLMFKD18Dg==",
		Ciphertext: "1i5uYNkUOU1F8GQFtSUWE5WdFRpJXt7YUrcbPeYtbYo1shfbFXOBcLFMQiQY4QHKsl79GFJluTRnCiVgvXIBLM5a/itdI7tOp7x4vZMf2BJ9kqcgab0URU21l9102IpVxs9p9l79m3ThQQABHVIPmrze7djGm8mZOwjFGjkHQPzBFybbDZTgQ4KWGYymgbdTKgmx3mBi8hosMPI3skJ6Wg==",
		IV:         "+e0+EmDwziw2rHSqTx/NhA==",
		Assert:     true,
	},

	// error: invalid character
	encryptedPhoneNumber{
		SessionKey: "4dcDdFhjmcWRJVlW8cxJMQ==",
		Ciphertext: "opjR1AiWT2JpOnsvp/mU453nYc4ptCYJeH5iQ2QuOcO6aqYlsaAS61DxPOFLWNdfIE4o71tIyDuRuLlO1tS+jk3TBJJl18d6vuC9q/qgz/dyxl6+8xsuQo8S9O55IOoYzO105QgDMPyH84fCXTPjCiM/+0xCcjjVHeqlN/f8oLXpUZ2hUNlqAmk1cm9ab3/RGeg8JF0IazYraLjSeUrzgg==",
		IV:         "q8PDCUo7st3j/qGcio8Ppw==",
		Assert:     false,
	},
}

func TestDecryptPhoneNumber(t *testing.T) {

	for _, data := range encryptedPhoneNumbers {

		if data.Assert {
			phone, err := DecryptPhoneNumber(data.SessionKey, data.Ciphertext, data.IV)
			if err != nil {
				t.Error(err)
			}
			if phone.PhoneNumber != "18048574657" ||
				phone.PurePhoneNumber != "18048574657" ||
				phone.CountryCode != "86" {
				t.Error("结果数据不一致")
				return
			}
			ssk := data.SessionKey + "..."
			_, err = DecryptPhoneNumber(ssk, data.Ciphertext, data.IV)
			if err == nil {
				t.Error("改变数据后得到了正确的结果")
			}
		} else {
			_, err := DecryptPhoneNumber(data.SessionKey, data.Ciphertext, data.IV)
			if err == nil {
				t.Error("错误的数据得到了正确的结果")
			}
		}
	}

}
