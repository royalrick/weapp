package security

import "github.com/medivhzhan/weapp/v3/request"

const apiImgSecCheck = "/wxa/img_sec_check"

type ImgSecCheckRequest struct {
	// 必填 要检测的图片文件，格式支持PNG、JPEG、JPG、GIF，图片尺寸不超过 750px x 1334px
	Media string
}

// 校验一张图片是否含有违法违规内容。
//
// 官方文档: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func (cli *Security) ImgSecCheck(req *ImgSecCheckRequest) (*request.CommonError, error) {
	url, err := cli.combineURI(apiImgSecCheck, nil, true)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.FormPostWithFile(url, "media", req.Media, res); err != nil {
		return nil, err
	}

	return res, nil
}
