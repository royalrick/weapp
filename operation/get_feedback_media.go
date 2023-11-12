package operation

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/medivhzhan/weapp/v3/request"
)

const apiGetFeedbackMedia = "/cgi-bin/media/getfeedbackmedia"

type GetFeedbackMediaRequest struct {
	// 必填	用户反馈信息的 record_id, 可通过 getFeedback 获取
	RecordID int `query:"record_id"`
	// 必填	图片的 mediaId
	MediaId string `query:"media_id"`
}

// 获取 mediaId 图片
func (cli *Operation) GetFeedbackMedia(req *GetFeedbackMediaRequest) (*http.Response, *request.CommonError, error) {

	uri, err := cli.combineURI(apiGetFeedbackMedia, req, true)
	if err != nil {
		return nil, nil, err
	}

	res, err := cli.request.GetWithBody(uri)
	if err != nil {
		return nil, nil, err
	}

	response := new(request.CommonError)
	switch header := res.Header.Get("Content-Type"); {
	case strings.HasPrefix(header, "application/json"): // 返回错误信息
		if err := json.NewDecoder(res.Body).Decode(response); err != nil {
			res.Body.Close()
			return nil, nil, err
		}
		return res, response, nil

	case strings.HasPrefix(header, "image"): // 返回文件
		return res, response, nil

	default:
		res.Body.Close()
		return nil, nil, errors.New("invalid response header: " + header)
	}
}
