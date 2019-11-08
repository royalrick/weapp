package weapp

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"
)

func TestIMGSecCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiIMGSecCheck {
			t.Fatalf("Except to path '%s',get '%s'", apiIMGSecCheck, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("media"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := imgSecCheck(ts.URL+apiIMGSecCheck, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMediaCheckAsync(t *testing.T) {

	localServer := http.NewServeMux()
	localServer.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnMediaCheckAsync(func(mix *MediaCheckAsyncResult) {
			if mix.ToUserName == "" {
				t.Error("ToUserName can not be empty")
			}

			if mix.FromUserName == "" {
				t.Error("FromUserName can not be empty")
			}
			if mix.CreateTime == 0 {
				t.Error("CreateTime can not be empty")
			}
			if mix.MsgType != "event" {
				t.Error("Unexpected message type")
			}
			if mix.Event != "wxa_media_check" {
				t.Error("Unexpected message event")
			}
			if mix.AppID == "" {
				t.Error("AppID can not be empty")
			}
			if mix.TraceID == "" {
				t.Error("TraceID can not be empty")
			}

		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	})

	tls := httptest.NewServer(localServer)
	defer tls.Close()

	remoteServer := http.NewServeMux()
	remoteServer.HandleFunc(apiMediaCheckAsync, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiMediaCheckAsync {
			t.Fatalf("Except to path '%s',get '%s'", apiMediaCheckAsync, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			MediaURL  string `json:"media_url"`
			MediaType uint8  `json:"media_type"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.MediaURL == "" {
			t.Error("Response column media_url can not be empty")
		}

		if params.MediaType == 0 {
			t.Error("Response column media_type can not be zero")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode"  : 0,
			"errmsg"   : "ok",
			"trace_id" : "967e945cd8a3e458f3c74dcb886068e9"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}

		raw = `{
			"ToUserName"      : "gh_38cc49f9733b",
			"FromUserName"    : "oH1fu0FdHqpToe2T6gBj0WyB8iS1",
			"CreateTime"      : 1552465698,
			"MsgType"         : "event",
			"Event"           : "wxa_media_check",
			"isrisky"         : 0,
			"extra_info_json" : "",
			"appid"           : "wxd8c59133dfcbfc71",
			"trace_id"        : "967e945cd8a3e458f3c74dcb886068e9",
			"status_code"     : 0
		 }`
		reader := strings.NewReader(raw)
		http.Post(tls.URL+"/notify", "application/json", reader)
	})

	remoteServer.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
		filename := testIMGName
		file, err := os.Open(filename)
		if err != nil {
			t.Fatal((err))
		}
		defer file.Close()

		ext := path.Ext(filename)
		ext = ext[1:len(ext)]
		w.Header().Set("Content-Type", "image/"+ext)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", path.Base(filename)))
		w.WriteHeader(http.StatusOK)

		if _, err := io.Copy(w, file); err != nil {
			t.Fatal(err)
		}
	})

	trs := httptest.NewServer(remoteServer)
	defer trs.Close()

	_, err := mediaCheckAsync(trs.URL+apiMediaCheckAsync, "mock-access-token", trs.URL+"/mediaurl", MediaTypeImage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMSGSecCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != "/wxa/img_sec_check" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Content string `json:"content"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Content == "" {
			t.Error("Response column content can not be empty")
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

	_, err := msgSecCheck(ts.URL+apiIMGSecCheck, "mock-access-token", "mock-content")
	if err != nil {
		t.Fatal(err)
	}
}
