package weapp

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
)

// TokenAPI 获取带 token 的 API
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

//Signature sha1签名
func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, P := range params {
		io.WriteString(h, P)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
