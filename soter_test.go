package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiVerifySignature {
			t.Fatalf("Except to path '%s',get '%s'", apiVerifySignature, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			OpenID        string `json:"openid"`
			JSONString    string `json:"json_string"`
			JSONSignature string `json:"json_signature"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.OpenID == "" {
			t.Error("Response column openid can not be empty")
		}

		if params.JSONString == "" {
			t.Error("Response column json_string can not be empty")
		}
		if params.JSONSignature == "" {
			t.Error("Response column json_signature can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"is_ok": true
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := verifySignature(ts.URL+apiVerifySignature, "mock-access-token", "mock-open-id", "mock-data", "mock-signature")
	if err != nil {
		t.Fatal(err)
	}
}
