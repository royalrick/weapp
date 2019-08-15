package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMonthlyVisitTrend(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetMonthlyVisitTrend {
			t.Fatalf("Except to path '%s',get '%s'", apiGetMonthlyVisitTrend, path)
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

		w.WriteHeader(http.StatusOK)

		raw := `{
			"list": [
			  {
				"ref_date": "201703",
				"session_cnt": 126513,
				"visit_pv": 426113,
				"visit_uv": 48659,
				"visit_uv_new": 6726,
				"stay_time_session": 56.4112,
				"visit_depth": 2.0189
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getVisitTrend("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetMonthlyVisitTrend)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetWeeklyVisitTrend(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetWeeklyVisitTrend {
			t.Fatalf("Except to path '%s',get '%s'", apiGetWeeklyVisitTrend, path)
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

		w.WriteHeader(http.StatusOK)

		raw := `{
			"list": [
			  {
				"ref_date": "20170306-20170312",
				"session_cnt": 986780,
				"visit_pv": 3251840,
				"visit_uv": 189405,
				"visit_uv_new": 45592,
				"stay_time_session": 54.5346,
				"visit_depth": 1.9735
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getVisitTrend("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetWeeklyVisitTrend)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDailyVisitTrend(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetDailyVisitTrend {
			t.Fatalf("Except to path '%s',get '%s'", apiGetDailyVisitTrend, path)
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

		w.WriteHeader(http.StatusOK)

		raw := `{
			"list": [
			  {
				"ref_date": "20170313",
				"session_cnt": 142549,
				"visit_pv": 472351,
				"visit_uv": 55500,
				"visit_uv_new": 5464,
				"stay_time_session": 0,
				"visit_depth": 1.9838
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getVisitTrend("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetDailyVisitTrend)
	if err != nil {
		t.Fatal(err)
	}
}
