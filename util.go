package weapp

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// TokenAPI 获取带 token 的 API 地址
func TokenAPI(api, token string) (string, error) {
	u, err := url.Parse(api)
	if err != nil {
		return "", err
	}
	query := u.Query()
	query.Set("access_token", token)
	u.RawQuery = query.Encode()

	return u.String(), nil
}

// EncodeURL add and encode parameters.
func EncodeURL(api string, params map[string]string) (string, error) {
	url, err := url.Parse(api)
	if err != nil {
		return "", err
	}

	query := url.Query()

	for k, v := range params {
		query.Set(k, v)
	}

	url.RawQuery = query.Encode()

	return url.String(), nil
}

// GetQuery returns url query value
func GetQuery(req *http.Request, key string) string {
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0]
	}

	return ""
}

// RandomString random string generator
//
// @ln length of return string
func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

// PostJSON perform a HTTP/POST request with json body
func PostJSON(api string, params interface{}, response interface{}) error {
	resp, err := PostJSONWithBody(api, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}

// PostJSONWithBody return with http body.
func PostJSONWithBody(api string, params interface{}) (*http.Response, error) {
	var reader *bytes.Reader
	if params != nil {
		raw, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		reader = bytes.NewReader(raw)
	}

	return http.Post(api, "application/json; charset=utf-8", reader)
}
