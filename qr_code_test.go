package weapp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestCreateQRCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		ePath := r.URL.EscapedPath()
		if ePath != apiCreateQRCode {
			t.Fatalf("Except to path '%s',get '%s'", apiCreateQRCode, ePath)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Path  string `json:"path"`
			Width int    `json:"width,omitempty"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Path == "" {
			t.Error("Response column path can not be empty")
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

	creator := QRCodeCreator{
		Path:  "mock/path",
		Width: 430,
	}
	resp, _, err := creator.create(ts.URL+apiCreateQRCode, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	resp.Body.Close()
}

func TestGetQRCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		ePath := r.URL.EscapedPath()
		if ePath != apiGetQrCode {
			t.Fatalf("Except to path '%s',get '%s'", apiGetQrCode, ePath)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Path      string `json:"path"`
			Width     int    `json:"width,omitempty"`
			AutoColor bool   `json:"auto_color,omitempty"`
			LineColor Color  `json:"line_color,omitempty"`
			IsHyaline bool   `json:"is_hyaline,omitempty"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Path == "" {
			t.Error("Response column path can not be empty")
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

	qr := QRCode{
		Path:      "mock/path",
		Width:     430,
		AutoColor: true,
		LineColor: Color{"r", "g", "b"},
		IsHyaline: true,
	}
	resp, _, err := qr.get(ts.URL+apiGetQrCode, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	resp.Body.Close()
}

func TestGetUnlimitedQRCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		ePath := r.URL.EscapedPath()
		if ePath != apiGetUnlimitedQRCode {
			t.Fatalf("Except to path '%s',get '%s'", apiGetUnlimitedQRCode, ePath)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Scene     string `json:"scene"`
			Page      string `json:"page,omitempty"`
			Width     int    `json:"width,omitempty"`
			AutoColor bool   `json:"auto_color,omitempty"`
			LineColor Color  `json:"line_color,omitempty"`
			IsHyaline bool   `json:"is_hyaline,omitempty"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Scene == "" {
			t.Error("Response column scene can not be empty")
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

	qr := UnlimitedQRCode{
		Scene:     "mock-scene-data",
		Page:      "mock/page",
		Width:     430,
		AutoColor: true,
		LineColor: Color{"r", "g", "b"},
		IsHyaline: true,
	}
	resp, _, err := qr.get(ts.URL+apiGetUnlimitedQRCode, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	resp.Body.Close()
}
