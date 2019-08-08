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

		attr := map[string]interface{}{
			"id":                     1,
			"name":                   "mock-name",
			"access_source_visit_uv": 100,
		}

		uv := map[string]interface{}{
			"index":     1,
			"province":  []interface{}{attr},
			"city":      []interface{}{attr},
			"genders":   []interface{}{attr},
			"platforms": []interface{}{attr},
			"devices":   []interface{}{attr},
			"ages":      []interface{}{attr},
		}

		data := map[string]interface{}{
			"ref_date":     "mock-ref-date",
			"visit_uv":     uv,
			"visit_uv_new": uv,
			"errcode":      0,
			"errmsg":       "mock-errmsg",
		}
		bts, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bts)
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
