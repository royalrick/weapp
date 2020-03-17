package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchSubmitPages(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSearchSubmitPages {
			t.Fatalf("Except to path '%s',get '%s'", apiSearchSubmitPages, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Pages []struct {
				Path  string `json:"path"`
				Query string `json:"query"`
			} `json:"pages"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if len(params.Pages) != 1 {
			t.Fatal("param pages can not be empty")
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

	sender := SearchSubmitPages{
		Pages: []struct {
			Path  string `json:"path"`
			Query string `json:"query"`
		}{{
			Path:  "/pages/index/index",
			Query: "id=test",
		}},
	}

	_, err := sender.send(ts.URL+apiSearchSubmitPages, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}
