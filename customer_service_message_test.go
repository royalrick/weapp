package weapp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestSetTyping(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSetTyping {
			t.Fatalf("Except to path '%s',get '%s'", apiSetTyping, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		param, ok := params["touser"]
		if !ok || param == "" {
			t.Error("param touser can not be empty")
		}
		param, ok = params["command"]
		if !ok || param == "" {
			t.Error("param command can not be empty")
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

	_, err := setTyping("mock-access-token", "mock-open-id", "mock-command", ts.URL+apiSetTyping)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadTempMedia(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiUploadTemplateMedia {
			t.Fatalf("Except to path '%s',get '%s'", apiUploadTemplateMedia, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "type"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		if _, _, err := r.FormFile("media"); err != nil {
			t.Fatal(err)

		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"type": "image",
			"media_id": "MEDIA_ID",
			"created_at": 1234567890
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := uploadTempMedia("mock-access-token", "mock-media-type", testIMGName, ts.URL+apiUploadTemplateMedia)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTempMedia(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("Expect 'GET' get '%s'", r.Method)
		}

		ePath := r.URL.EscapedPath()
		if ePath != apiGetTemplateMedia {
			t.Fatalf("Except to path '%s',get '%s'", apiGetTemplateMedia, ePath)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "media_id"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

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
	}))
	defer ts.Close()

	resp, _, err := getTempMedia("mock-access-token", "mock-media-id", ts.URL+apiGetTemplateMedia)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
}

func TestSendCustomerServiceMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSendMessage {
			t.Fatalf("Except to path '%s',get '%s'", apiSendMessage, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Error("access_token can not be empty")
		}

		params := make(object)
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		param, ok := params["touser"]
		if !ok || param == "" {
			t.Error("param touser can not be empty")
		}
		param, ok = params["msgtype"]
		if !ok || param == "" {
			t.Error("param command can not be empty")
		}
		switch param {

		case "text":
			param, ok := params["text"]
			if !ok {
				t.Error("param text can not be empty")
			}

			obj, ok := param.(object)
			if !ok {
				t.Errorf("unexpected value type of tex: %s", reflect.TypeOf(param))
			}

			param, ok = obj["content"]
			if !ok {
				t.Error("param text.content can not be empty")
			}

		case "image":
			param, ok := params["image"]
			if !ok {
				t.Error("param command can not be empty")
			}

			obj, ok := param.(object)
			if !ok {
				t.Error("unexpected value type of image")
			}

			param, ok = obj["media_id"]
			if !ok {
				t.Error("param image.media_id can not be empty")
			}

		case "link":
			param, ok := params["link"]
			if !ok {
				t.Error("param link can not be empty")
			}

			obj, ok := param.(object)
			if !ok {
				t.Error("unexpected value type of link")
			}

			param, ok = obj["title"]
			if !ok {
				t.Error("param link.title can not be empty")
			}

			param, ok = obj["description"]
			if !ok {
				t.Error("param link.description can not be empty")
			}

			param, ok = obj["url"]
			if !ok {
				t.Error("param link.url can not be empty")
			}

			param, ok = obj["thumb_url"]
			if !ok {
				t.Error("param link.thumb_url can not be empty")
			}

		case "miniprogrampage":
			param, ok := params["miniprogrampage"]
			if !ok {
				t.Error("param miniprogrampage can not be empty")
			}

			obj, ok := param.(object)
			if !ok {
				t.Error("unexpected value type of miniprogrampage")
			}

			param, ok = obj["title"]
			if !ok {
				t.Error("param miniprogrampage.title can not be empty")
			}

			param, ok = obj["pagepath"]
			if !ok {
				t.Error("param miniprogrampage.pagepath can not be empty")
			}

			param, ok = obj["thumb_media_id"]
			if !ok {
				t.Error("param miniprogrampage.thumb_media_id can not be empty")
			}

		default:
			t.Fatalf("unexpected msgtype: %s", param)
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

	messages := []csMessage{
		csMessage{
			Receiver: "mock-open-id",
			Type:     csMsgTypeText,
			Image: CSMsgImage{
				MediaID: "mock-media-id",
			},
		},
		csMessage{
			Receiver: "mock-open-id",
			Type:     csMsgTypeLink,
			Link: CSMsgLink{
				Title:       "mock-title",
				Description: "mock-description",
				URL:         "mock-url",
				ThumbURL:    "mock-thumb-url",
			},
		},
		csMessage{
			Receiver: "mock-open-id",
			Type:     csMsgTypeMPCard,
			MPCard: CSMsgMPCard{
				Title:        "mock-title",
				PagePath:     "mock-page-path",
				ThumbMediaID: "mock-thumb-media-id",
			},
		},
		csMessage{
			Receiver: "mock-open-id",
			Type:     csMsgTypeText,
			Text: CSMsgText{
				Content: "mock-content",
			},
		},
	}
	for _, msg := range messages {
		_, err := doSendMessage("mock-access-token", msg, ts.URL+apiSendMessage)
		if err != nil {
			t.Error(err)
		}
	}
}
