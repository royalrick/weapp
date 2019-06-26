package util

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"math/rand"
	"net"
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

// PostXML perform a HTTP/POST request with XML body.
func PostXML(api string, params interface{}, response interface{}) error {
	var reader *bytes.Reader
	if params != nil {

		raw, err := xml.Marshal(params)
		if err != nil {
			return err
		}

		reader = bytes.NewReader(raw)
	}

	resp, err := http.Post(api, "application/xml; charset=utf-8", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return xml.NewDecoder(resp.Body).Decode(response)
}

// TSLPostXML perform a HTTPS/POST request with XML body.
func TSLPostXML(uri string, params interface{}, certPath, keyPath string, response interface{}) error {

	data, err := xml.Marshal(params)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(data)

	cli, err := NewTLSClient(certPath, keyPath)
	if err != nil {
		return err
	}

	resp, err := cli.Post(uri, "application/xml; charset=utf-8", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return xml.NewDecoder(resp.Body).Decode(response)
}

// NewTLSClient 创建支持双向证书认证的 http.Client.
func NewTLSClient(certPath, keyPath string) (httpClient *http.Client, err error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return newTLSClient(tlsConfig)
}

func newTLSClient(tlsConfig *tls.Config) (*http.Client, error) {

	dialTLS := func(network, addr string) (net.Conn, error) {
		return tls.DialWithDialer(&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}, network, addr, tlsConfig)
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			DialTLS:               dialTLS,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}, nil
}

// FetchIP current IP address
func FetchIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for index := range addrs {

		// 检查ip地址判断是否回环地址
		if IPNet, ok := addrs[index].(*net.IPNet); ok && !IPNet.IP.IsLoopback() {
			if IPNet.IP.To4() != nil {
				return IPNet.IP, nil
			}

		}
	}

	return nil, errors.New("failed to found IP address")
}
