package request

import (
	"fmt"
	"net/url"
)

func EncodeURL(api string, queries map[string]interface{}) (string, error) {
	url, err := url.Parse(api)
	if err != nil {
		return "", err
	}

	query := url.Query()

	for k, v := range queries {
		query.Set(k, fmt.Sprintf("%v", v))
	}

	url.RawQuery = query.Encode()

	return url.String(), nil
}
