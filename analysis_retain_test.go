package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMonthlyRetain(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		raw := `{
			"ref_date": "201702",
			"visit_uv_new": [
			  {
				"key": 0,
				"value": 346249
			  }
			],
			"visit_uv": [
			  {
				"key": 0,
				"value": 346249
			  }
			]
		  }`
		w.Write([]byte(raw))

		if r.Method != "POST" {
			t.Fatalf("Except 'POST' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetMonthlyRetain {
			t.Fatalf("Except to path '%s',got '%s'", apiGetMonthlyRetain, path)
		}

		r.ParseForm()
		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		param, ok := params["begin_date"]
		if !ok || param == "" {
			t.Log("param begin_date can not be empty")
			t.Fail()
		}
		param, ok = params["end_date"]
		if !ok || param == "" {
			t.Log("param end_date can not be empty")
			t.Fail()
		}

	}))
	defer ts.Close()

	_, err := getRetain("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetMonthlyRetain)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetWeeklyRetain(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		raw := `{
			"ref_date": "20170306-20170312",
			"visit_uv_new": [
			  {
				"key": 0,
				"value": 0
			  },
			  {
				"key": 1,
				"value": 16853
			  }
			],
			"visit_uv": [
			  {
				"key": 0,
				"value": 0
			  },
			  {
				"key": 1,
				"value": 99310
			  }
			]
		  }`
		w.Write([]byte(raw))

		if r.Method != "POST" {
			t.Fatalf("Except 'POST' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetWeeklyRetain {
			t.Fatalf("Except to path '%s',got '%s'", apiGetWeeklyRetain, path)
		}

		r.ParseForm()
		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		param, ok := params["begin_date"]
		if !ok || param == "" {
			t.Log("param begin_date can not be empty")
			t.Fail()
		}
		param, ok = params["end_date"]
		if !ok || param == "" {
			t.Log("param end_date can not be empty")
			t.Fail()
		}

	}))
	defer ts.Close()

	_, err := getRetain("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetWeeklyRetain)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDailyRetainAPI(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		raw := `{
			"ref_date": "20170313",
			"visit_uv_new": [
			  {
				"key": 0,
				"value": 5464
			  }
			],
			"visit_uv": [
			  {
				"key": 0,
				"value": 55500
			  }
			]
		  }`
		w.Write([]byte(raw))

		if r.Method != "POST" {
			t.Fatalf("Except 'POST' got '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetDailyRetain {
			t.Fatalf("Except to path '%s',got '%s'", apiGetDailyRetain, path)
		}

		r.ParseForm()
		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		param, ok := params["begin_date"]
		if !ok || param == "" {
			t.Log("param begin_date can not be empty")
			t.Fail()
		}
		param, ok = params["end_date"]
		if !ok || param == "" {
			t.Log("param end_date can not be empty")
			t.Fail()
		}

	}))
	defer ts.Close()

	_, err := getRetain("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetDailyRetain)
	if err != nil {
		t.Fatal(err)
	}
}
