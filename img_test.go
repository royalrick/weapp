package weapp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func TestAICrop(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/cv/img/aicrop" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"results": [
			{
				"crop_left": 112,
				"crop_top": 0,
				"crop_right": 839,
				"crop_bottom": 727
			},
			{
				"crop_left": 0,
				"crop_top": 205,
				"crop_right": 965,
				"crop_bottom": 615
			}
			],
			"img_size": {
				"w": 966,
				"h": 728
			}
		 }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := aiCrop(ts.URL+apiAICrop, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAICropByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc("/cv/img/aicrop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/cv/img/aicrop" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"results": [
			{
				"crop_left": 112,
				"crop_top": 0,
				"crop_right": 839,
				"crop_bottom": 727
			},
			{
				"crop_left": 0,
				"crop_top": 205,
				"crop_right": 965,
				"crop_bottom": 615
			}
			],
			"img_size": {
				"w": 966,
				"h": 728
			}
		 }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
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

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := aiCropByURL(ts.URL+apiAICrop, "mock-access-token", ts.URL+"/mediaurl")
	if err != nil {
		t.Fatal(err)
	}
}
func TestScanQRCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiScanQRCode {
			t.Fatalf("Except to path '%s',get '%s'", apiScanQRCode, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
				"errcode": 0, 
				"errmsg": "ok", 
				"code_results": [
					{
						"type_name": "QR_CODE", 
						"data": "http://www.qq.com", 
						"pos": {
							"left_top": {
								"x": 585, 
								"y": 378
							}, 
							"right_top": {
								"x": 828, 
								"y": 378
							}, 
							"right_bottom": {
								"x": 828, 
								"y": 618
							}, 
							"left_bottom": {
								"x": 585, 
								"y": 618
							}
						}
					}, 
					{
						"type_name": "QR_CODE", 
						"data": "https://mp.weixin.qq.com", 
						"pos": {
							"left_top": {
								"x": 185, 
								"y": 142
							}, 
							"right_top": {
								"x": 396, 
								"y": 142
							}, 
							"right_bottom": {
								"x": 396, 
								"y": 353
							}, 
							"left_bottom": {
								"x": 185, 
								"y": 353
							}
						}
					}, 
					{
						"type_name": "EAN_13", 
						"data": "5906789678957"
					}, 
					{
						"type_name": "CODE_128", 
						"data": "50090500019191"
					}
				], 
				"img_size": {
					"w": 1000, 
					"h": 900
				}
			}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := scanQRCode(ts.URL+apiScanQRCode, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestScanQRCodeByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiScanQRCode, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiScanQRCode {
			t.Fatalf("Except to path '%s',get '%s'", apiScanQRCode, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0, 
			"errmsg": "ok", 
			"code_results": [
				{
					"type_name": "QR_CODE", 
					"data": "http://www.qq.com", 
					"pos": {
						"left_top": {
							"x": 585, 
							"y": 378
						}, 
						"right_top": {
							"x": 828, 
							"y": 378
						}, 
						"right_bottom": {
							"x": 828, 
							"y": 618
						}, 
						"left_bottom": {
							"x": 585, 
							"y": 618
						}
					}
				}, 
				{
					"type_name": "QR_CODE", 
					"data": "https://mp.weixin.qq.com", 
					"pos": {
						"left_top": {
							"x": 185, 
							"y": 142
						}, 
						"right_top": {
							"x": 396, 
							"y": 142
						}, 
						"right_bottom": {
							"x": 396, 
							"y": 353
						}, 
						"left_bottom": {
							"x": 185, 
							"y": 353
						}
					}
				}, 
				{
					"type_name": "EAN_13", 
					"data": "5906789678957"
				}, 
				{
					"type_name": "CODE_128", 
					"data": "50090500019191"
				}
			], 
			"img_size": {
				"w": 1000, 
				"h": 900
			}
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
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

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := scanQRCodeByURL(ts.URL+apiScanQRCode, "mock-access-token", ts.URL+"/mediaurl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSuperResolution(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSuperResolution {
			t.Fatalf("Except to path '%s',get '%s'", apiSuperResolution, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		if _, _, err := r.FormFile("img"); err != nil {
			t.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0, 
			"errmsg": "ok", 
			"media_id": "6WXsIXkG7lXuDLspD9xfm5dsvHzb0EFl0li6ySxi92ap8Vl3zZoD9DpOyNudeJGB"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := superResolution(ts.URL+apiSuperResolution, "mock-access-token", testIMGName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSuperResolutionByURL(t *testing.T) {
	server := http.NewServeMux()
	server.HandleFunc(apiSuperResolution, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSuperResolution {
			t.Fatalf("Except to path '%s',get '%s'", apiSuperResolution, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		queries := []string{"access_token", "img_url"}
		for _, v := range queries {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("%v can not be empty", v)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0, 
			"errmsg": "ok", 
			"media_id": "6WXsIXkG7lXuDLspD9xfm5dsvHzb0EFl0li6ySxi92ap8Vl3zZoD9DpOyNudeJGB"
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	})

	server.HandleFunc("/mediaurl", func(w http.ResponseWriter, r *http.Request) {
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

	ts := httptest.NewServer(server)
	defer ts.Close()

	_, err := superResolutionByURL(ts.URL+apiSuperResolution, "mock-access-token", ts.URL+"/mediaurl")
	if err != nil {
		t.Fatal(err)
	}
}
