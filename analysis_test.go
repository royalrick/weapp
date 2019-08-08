package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserPortrait(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		raw := `{
			"refDate": "20170611",
			"visitUvNew": {
			  "province": [
				{
				  "id": 31,
				  "name": "广东省",
				  "value": 215
				}
			  ],
			  "city": [
				{
				  "id": 3102,
				  "name": "广州",
				  "value": 78
				}
			  ],
			  "genders": [
				{
				  "id": 1,
				  "name": "男",
				  "value": 2146
				}
			  ],
			  "platforms": [
				{
				  "id": 1,
				  "name": "iPhone",
				  "value": 27642
				}
			  ],
			  "devices": [
				{
				  "name": "OPPO R9",
				  "value": 61
				}
			  ],
			  "ages": [
				{
				  "id": 1,
				  "name": "17岁以下",
				  "value": 151
				}
			  ]
			},
			"visitUv": {
			  "province": [
				{
				  "id": 31,
				  "name": "广东省",
				  "value": 1341
				}
			  ],
			  "city": [
				{
				  "id": 3102,
				  "name": "广州",
				  "value": 234
				}
			  ],
			  "genders": [
				{
				  "id": 1,
				  "name": "男",
				  "value": 14534
				}
			  ],
			  "platforms": [
				{
				  "id": 1,
				  "name": "iPhone",
				  "value": 21750
				}
			  ],
			  "devices": [
				{
				  "name": "OPPO R9",
				  "value": 617
				}
			  ],
			  "ages": [
				{
				  "id": 1,
				  "name": "17岁以下",
				  "value": 3156
				}
			  ]
			},
			"errMsg": "openapi.analysis.getUserPortrait:ok"
		  }`
		w.Write([]byte(raw))

		if r.Method != "POST" {
			t.Fatalf("Except 'POST' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != getUserPortraitAPI {
			t.Fatalf("Except to path '%s',got '%s'", getUserPortraitAPI, path)
		}

		r.ParseForm()
		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		param, ok := params["begin_date"]
		if !ok || param == "" {
			t.Log("param begin_date can not be empty")
			t.Fail()
		}
		param, ok = params["end_date"]
		if !ok || param == "" {
			t.Log("param end_date can not be empty")
			t.Fail()
		}

	}))
	defer ts.Close()

	_, err := getUserPortrait("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+getUserPortraitAPI)
	if err != nil {
		t.Fatal(err)
	}
}
