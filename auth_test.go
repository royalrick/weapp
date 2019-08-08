package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"openid":      "mock-openid",
			"session_key": "mock-session_key",
			"unionid":     "mock-unionid",
			"errcode":     0,
			"errmsg":      "mock-errmsg",
		}
		bts, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bts)
		if r.Method != "GET" {
			t.Fatalf("Except 'GET' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiLogin {
			t.Fatalf("Except to path '%s',got '%s'", apiLogin, path)
		}

		r.ParseForm()
		queries := []string{"appid", "secret", "js_code", "grant_type"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}

		}
	}))
	defer ts.Close()

	_, err := login("mock-appid", "mock-secret", "mock-code", ts.URL+apiLogin)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAccessToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"access_token": "mock-access-token",
			"expires_in":   1000,
			"errcode":      0,
			"errmsg":       "mock-errmsg",
		}
		bts, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bts)
		if r.Method != "GET" {
			t.Fatalf("Except 'GET' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetAccessToken {
			t.Fatalf("Except to path '%s',got '%s'", apiGetAccessToken, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("grant_type") != "client_credential" {
			t.Fatal("invalid client_credential")
		}

		queries := []string{"appid", "secret"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}

		}
	}))
	defer ts.Close()

	_, err := getAccessToken("mock-appid", "mock-secret", ts.URL+apiGetAccessToken)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPaidUnionID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"unionid": "mock-unionid",
			"errcode": 0,
			"errmsg":  "mock-errmsg",
		}
		bts, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bts)
		if r.Method != "GET" {
			t.Fatalf("Except 'GET' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetPaidUnionID {
			t.Fatalf("Except to path '%s',got '%s'", apiGetPaidUnionID, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"openid", "access_token", "transaction_id"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}

		}
	}))
	defer ts.Close()

	_, err := getPaidUnionID("mock-access-token", "mock-open-id", "mock-transaction-id", ts.URL+apiGetPaidUnionID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPaidUnionIDWithMCH(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"unionid": "mock-unionid",
			"errcode": 0,
			"errmsg":  "mock-errmsg",
		}
		bts, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(bts)
		if r.Method != "GET" {
			t.Fatalf("Except 'GET' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetPaidUnionID {
			t.Fatalf("Except to path '%s',got '%s'", apiGetPaidUnionID, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"openid", "access_token", "mch_id", "out_trade_no"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}

		}
	}))
	defer ts.Close()

	_, err := getPaidUnionIDWithMCH("mock-access-token", "mock-open-id", "mock-out-trade-number", "mock-mch-id", ts.URL+apiGetPaidUnionID)
	if err != nil {
		t.Fatal(err)
	}
}
