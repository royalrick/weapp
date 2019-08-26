package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplyPlugin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiPlugin {
			t.Fatalf("Except to path '%s',get '%s'", apiPlugin, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Action      string `json:"action"`
			PluginAppID string `json:"plugin_appid"`
			Reason      string `json:"reason"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Action != "apply" {
			t.Error("Unexpected action")
		}

		if params.PluginAppID == "" {
			t.Error("Response column plugin_appid can not be empty")
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

	_, err := applyPlugin(ts.URL+apiPlugin, "mock-access-token", "plugin-app-id", "mock-reason")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPluginDevApplyList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDevPlugin {
			t.Fatalf("Except to path '%s',get '%s'", apiDevPlugin, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Action string `json:"action"`
			Page   uint   `json:"page"`
			Number uint   `json:"num"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Action != "dev_apply_list" {
			t.Error("Unexpected action")
		}

		if params.Page == 0 {
			t.Error("Response column page can not be zero")
		}

		if params.Number == 0 {
			t.Error("Response column num can not be zero")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"apply_list": [{
			  "appid": "xxxxxxxxxxxxx",
			  "status": 1,
			  "nickname": "名称",
			  "headimgurl": "**********",
			  "reason": "polo has gone",
			  "apply_url": "*******",
			  "create_time": "1536305096",
			  "categories": [{
				"first": "IT科技",
				"second": "硬件与设备"
			  }]
			}]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getPluginDevApplyList(ts.URL+apiDevPlugin, "mock-access-token", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPluginList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiPlugin {
			t.Fatalf("Except to path '%s',get '%s'", apiPlugin, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Action string `json:"action"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Action != "list" {
			t.Error("Unexpected action")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"plugin_list": [{
			  "appid": "aaaa",
			  "status": 1,
			  "nickname": "插件昵称",
			  "headimgurl": "http://plugin.qq.com"
			}]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getPluginList(ts.URL+apiPlugin, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetDevPluginApplyStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDevPlugin {
			t.Fatalf("Except to path '%s',get '%s'", apiDevPlugin, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Action string `json:"action"`
			AppID  string `json:"appid"`
			Reason string `json:"reason"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Action == "" {
			t.Error("Response column action can not be empty")
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

	_, err := setDevPluginApplyStatus(ts.URL+apiDevPlugin, "mock-access-token", "mock-plugin-app-id", "mock-reason", DevAgree)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnbindPlugin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiPlugin {
			t.Fatalf("Except to path '%s',get '%s'", apiPlugin, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Action      string `json:"action"`
			PluginAppID string `json:"plugin_appid"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Action != "unbind" {
			t.Error("Unexpected action")
		}

		if params.PluginAppID == "" {
			t.Error("Response column plugin_appid can not be empty")
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

	_, err := unbindPlugin(ts.URL+apiPlugin, "mock-access-token", "mock-plugin-app-id")
	if err != nil {
		t.Fatal(err)
	}
}
