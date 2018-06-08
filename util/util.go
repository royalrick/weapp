package util

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// TokenAPI 获取带 token 的 API 地址
func TokenAPI(api, token string) (string, error) {
	u, err := url.Parse(api)
	if err != nil {
		return "", nil
	}
	query := u.Query()
	query.Set("access_token", token)
	u.RawQuery = query.Encode()

	return u.String(), nil
}

// GetQuery returns url query value
func GetQuery(req *http.Request, key string) string {
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0]
	}

	return ""
}

// SignByMD5 多参数通过MD5签名
func SignByMD5(data map[string]string) (string, error) {

	var group []string
	for k, v := range data {
		group = append(group, k+"="+v)
	}

	sort.Strings(group)
	str := strings.Join(group, "&")

	str, err := MD5(str)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(str), nil
}

// MD5 加密
func MD5(str string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(str)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
