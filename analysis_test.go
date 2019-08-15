package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserPortrait(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetUserPortrait {
			t.Fatalf("Except to path '%s',get '%s'", apiGetUserPortrait, path)
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
			"ref_date": "20170611",
			"visit_uv_new": {
			  "province": [
				{
				  "id": 31,
				  "name": "广东省",
				  "value": 215
				}
			  ],
			  "city": [
			   {
				  "id": 3102,
				  "name": "广州",
				  "value": 78
				}
			  ],
			  "genders": [
				{
				  "id": 1,
				  "name": "男",
				  "value": 2146
				}
			  ],
			  "platforms": [
				{
				  "id": 1,
				  "name": "iPhone",
				  "value": 27642
				}
			  ],
			  "devices": [
				{
				  "name": "OPPO R9",
				  "value": 61
				}
			  ],
			  "ages": [
				{
				  "id": 1,
				  "name": "17岁以下",
				  "value": 151
				}
			  ]
			},
			"visit_uv": {
			  "province": [
				{
				  "id": 31,
				  "name": "广东省",
				  "value": 1341
				}
			  ],
			  "city": [
			   {
				  "id": 3102,
				  "name": "广州",
				  "value": 234
				}
			  ],
			  "genders": [
				{
				  "id": 1,
				  "name": "男",
				  "value": 14534
				}
			  ],
			  "platforms": [
				{
				  "id": 1,
				  "name": "iPhone",
				  "value": 21750
				}
			  ],
			  "devices": [
				{
				  "name": "OPPO R9",
				  "value": 617
				}
			  ],
			  "ages": [
				{
				  "id": 1,
				  "name": "17岁以下",
				  "value": 3156
				}
			  ]
			}
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getUserPortrait("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetUserPortrait)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetVisitDistribution(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetVisitDistribution {
			t.Fatalf("Except to path '%s',get '%s'", apiGetVisitDistribution, path)
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
			"ref_date": "20170313",
			"list": [
			  {
				"index": "access_source_session_cnt",
				"item_list": [
				  {
					"key": 10,
					"value": 5
				  },
				  {
					"key": 8,
					"value": 687
				  },
				  {
					"key": 7,
					"value": 10740
				  },
				  {
					"key": 6,
					"value": 1961
				  },
				  {
					"key": 5,
					"value": 677
				  },
				  {
					"key": 4,
					"value": 653
				  },
				  {
					"key": 3,
					"value": 1120
				  },
				  {
					"key": 2,
					"value": 10243
				  },
				  {
					"key": 1,
					"value": 116578
				  }
				]
			  },
			  {
				"index": "access_staytime_info",
				"item_list": [
				  {
					"key": 8,
					"value": 16329
				  },
				  {
					"key": 7,
					"value": 19322
				  },
				  {
					"key": 6,
					"value": 21832
				  },
				  {
					"key": 5,
					"value": 19539
				  },
				  {
					"key": 4,
					"value": 29670
				  },
				  {
					"key": 3,
					"value": 19667
				  },
				  {
					"key": 2,
					"value": 11794
				  },
				  {
					"key": 1,
					"value": 4511
				  }
				]
			  },
			  {
				"index": "access_depth_info",
				"item_list": [
				  {
					"key": 5,
					"value": 217
				  },
				  {
					"key": 4,
					"value": 3259
				  },
				  {
					"key": 3,
					"value": 32445
				  },
				  {
					"key": 2,
					"value": 63542
				  },
				  {
					"key": 1,
					"value": 43201
				  }
				]
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getVisitDistribution("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetVisitDistribution)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetVisitPage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetVisitPage {
			t.Fatalf("Except to path '%s',get '%s'", apiGetVisitPage, path)
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
			"ref_date": "20170313",
			"list": [
			  {
				"page_path": "pages/main/main.html",
				"page_visit_pv": 213429,
				"page_visit_uv": 55423,
				"page_staytime_pv": 8.139198,
				"entrypage_pv": 117922,
				"exitpage_pv": 61304,
				"page_share_pv": 180,
				"page_share_uv": 166
			  },
			  {
				"page_path": "pages/linedetail/linedetail.html",
				"page_visit_pv": 155030,
				"page_visit_uv": 42195,
				"page_staytime_pv": 35.462395,
				"entrypage_pv": 21101,
				"exitpage_pv": 47051,
				"page_share_pv": 47,
				"page_share_uv": 42
			  },
			  {
				"page_path": "pages/search/search.html",
				"page_visit_pv": 65011,
				"page_visit_uv": 24716,
				"page_staytime_pv": 6.889634,
				"entrypage_pv": 1811,
				"exitpage_pv": 3198,
				"page_share_pv": 0,
				"page_share_uv": 0
			  },
			  {
				"page_path": "pages/stationdetail/stationdetail.html",
				"page_visit_pv": 29953,
				"page_visit_uv": 9695,
				"page_staytime_pv": 7.558508,
				"entrypage_pv": 1386,
				"exitpage_pv": 2285,
				"page_share_pv": 0,
				"page_share_uv": 0
			  },
			  {
				"page_path": "pages/switch-city/switch-city.html",
				"page_visit_pv": 8928,
				"page_visit_uv": 4017,
				"page_staytime_pv": 9.22659,
				"entrypage_pv": 748,
				"exitpage_pv": 1613,
				"page_share_pv": 0,
				"page_share_uv": 0
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getVisitPage("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetVisitPage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDailySummary(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetDailySummary {
			t.Fatalf("Except to path '%s',get '%s'", apiGetDailySummary, path)
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
				"visit_total": 391,
				"share_pv": 572,
				"share_uv": 383
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getDailySummary("mock-access-token", "mock-begin-date", "mock-end-date", ts.URL+apiGetDailySummary)
	if err != nil {
		t.Fatal(err)
	}
}
