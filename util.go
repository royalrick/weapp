package weapp

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// tokenAPI 获取带 token 的 API 地址
func tokenAPI(api, token string) (string, error) {
	queries := requestQueries{
		"access_token": token,
	}

	return encodeURL(api, queries)
}

// encodeURL add and encode parameters.
func encodeURL(api string, params requestQueries) (string, error) {
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

// getQuery returns url query value
func getQuery(req *http.Request, key string) string {
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0]
	}

	return ""
}

// randomString random string generator
//
// ln length of return string
func randomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

// postJSON perform a HTTP/POST request with json body
func postJSON(api string, params interface{}, response interface{}) error {
	resp, err := postJSONWithBody(api, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}

// postJSONWithBody return with http body.
func postJSONWithBody(api string, params interface{}) (*http.Response, error) {
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
