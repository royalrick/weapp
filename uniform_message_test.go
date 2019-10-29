package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendUniformMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSendUniformMessage {
			t.Fatalf("Except to path '%s',get '%s'", apiSendUniformMessage, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ToUser string `json:"touser"` // 用户 openid

			UniformWeappTmpMsg struct {
				TemplateID string `json:"template_id"`
				Page       string `json:"page"`
				FormID     string `json:"form_id"`
				Data       map[string]struct {
					Value string `json:"value"`
				} `json:"data"`
				EmphasisKeyword string `json:"emphasis_keyword"`
			} `json:"weapp_template_msg"`
			UniformMpTmpMsg struct {
				AppID       string `json:"appid"`
				TemplateID  string `json:"template_id"`
				URL         string `json:"url"`
				Miniprogram struct {
					AppID    string `json:"appid"`
					PagePath string `json:"pagepath"`
				} `json:"miniprogram"`
				Data map[string]struct {
					Value string `json:"value"`
					Color string `json:"color,omitempty"`
				} `json:"data"`
			} `json:"mp_template_msg"`
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

	sender := UniformMsgSender{
		ToUser: "mock-open-id",
		UniformWeappTmpMsg: UniformWeappTmpMsg{
			TemplateID: "mock-template-id",
			Page:       "mock-page",
			FormID:     "mock-form-id",
			Data: UniformMsgData{
				"mock-keyword": UniformMsgKeyword{Value: "mock-value"},
			},
			EmphasisKeyword: "mock-keyword.DATA",
		},
		UniformMpTmpMsg: UniformMpTmpMsg{
			AppID:       "mock-app-id",
			TemplateID:  "mock-template-id",
			URL:         "mock-url",
			Miniprogram: UniformMsgMiniprogram{"mock-miniprogram-app-id", "mock-page-path"},
			Data: UniformMsgData{
				"mock-keyword": UniformMsgKeyword{"mock-value", "mock-color"},
			},
		},
	}

	_, err := sender.send(ts.URL+apiSendUniformMessage, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}
