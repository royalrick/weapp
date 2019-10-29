package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendSubscribeMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSubscribeMessage {
			t.Fatalf("Except to path '%s',get '%s'", apiSubscribeMessage, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ToUser string `json:"touser"` // 用户 openid

			TemplateID string `json:"template_id"`
			Page       string `json:"page,omitempty"`
			Data       map[string]struct {
				Value string `json:"value"`
			} `json:"data"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ToUser == "" {
			t.Fatal("param touser can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	sender := SubscribeMessage{
		ToUser:     "mock-open-id",
		TemplateID: "mock-template-id",
		Page:       "mock-page",
		Data: SubscribeMessageData{
			"mock01.DATA": {
				Value: "mock-value",
			},
		},
	}

	_, err := sender.send(ts.URL+apiSubscribeMessage, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}
