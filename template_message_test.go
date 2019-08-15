package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddTemplate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiAddTemplate {
			t.Fatalf("Except to path '%s',get '%s'", apiAddTemplate, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ID            string `json:"id"`
			KeywordIDList []uint `json:"keyword_id_list"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ID == "" {
			t.Error("Param template id can not be empty")
		}

		if len(params.KeywordIDList) == 0 {
			t.Error("Param template id list can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"template_id": "wDYzYZVxobJivW9oMpSCpuvACOfJXQIoKUm0PY397Tc"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	res, err := addTemplate("mock-template-id", "mock-access-token", []uint{1, 2, 3}, ts.URL+apiAddTemplate)
	if err != nil {
		t.Fatal(err)
	}

	if res.TemplateID == "" {
		t.Fatal("response did not contain column template_id")
	}
}

func TestDeleteTemplate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDeleteTemplate {
			t.Fatalf("Except to path '%s',get '%s'", apiDeleteTemplate, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			TemplateID string `json:"template_id"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.TemplateID == "" {
			t.Error("Param template id can not be empty")
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

	_, err := deleteTemplate("mock-template-id", "mock-access-token", ts.URL+apiDeleteTemplate)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTemplateLibraryByID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetTemplateLibraryByID {
			t.Fatalf("Except to path '%s',get '%s'", apiGetTemplateLibraryByID, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ID string `json:"id"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ID == "" {
			t.Error("Param template id can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"id": "AT0002",
			"title": "购买成功通知",
			"keyword_list": [
				{
					"keyword_id": 3,
					"name": "购买地点",
					"example": "TIT造舰厂"
				},
				{
					"keyword_id": 4,
					"name": "购买时间",
					"example": "2016年6月6日"
				},
				{
					"keyword_id": 5,
					"name": "物品名称",
					"example": "咖啡"
				}
			]
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	res, err := getTemplateLibraryByID("mock-template-id", "mock-access-token", ts.URL+apiGetTemplateLibraryByID)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID == "" {
		t.Error("response column ID can not be empty")
	}

	if res.Title == "" {
		t.Error("response column title can not be empty")
	}
	if len(res.KeywordList) == 0 {
		t.Error("response column keyword_list can not be empty")
	}

	keyword := res.KeywordList[0]
	if keyword.KeywordID == 0 {
		t.Error("response column keyword_list.keyword_id can not be zero")
	}

	if keyword.Name == "" {
		t.Error("response column keyword_list.name can not be empty")
	}

	if keyword.Example == "" {
		t.Error("response column keyword_list.example can not be empty")
	}
}

func TestGetTemplateLibraryList(t *testing.T) {
	const offset, count = 5, 10
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetTemplateLibraryList {
			t.Fatalf("Except to path '%s',get '%s'", apiGetTemplateLibraryList, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Offset uint `json:"offset"`
			Count  uint `json:"count"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Offset != offset {
			t.Errorf("expected %v get %v", offset, params.Offset)
		}

		if params.Count != count {
			t.Errorf("expected %v get %v", count, params.Count)
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode":0,
			"errmsg":"ok",
			"list":[
			{"id":"AT0002","title":"购买成功通知"},
			{"id":"AT0003","title":"购买失败通知"},
			{"id":"AT0004","title":"交易提醒"},
			{"id":"AT0005","title":"付款成功通知"},
			{"id":"AT0006","title":"付款失败通知"}
			],
			"total_count":599
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	err := getTemplateList(ts.URL+apiGetTemplateLibraryList, "mock-access-token", offset, count, new(GetTemplateLibraryListResponse))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTemplateList(t *testing.T) {
	const offset, count = 5, 10
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetTemplateList {
			t.Fatalf("Except to path '%s',get '%s'", apiGetTemplateList, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Offset uint `json:"offset"`
			Count  uint `json:"count"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Offset != offset {
			t.Errorf("expected %v get %v", offset, params.Offset)
		}

		if params.Count != count {
			t.Errorf("expected %v get %v", count, params.Count)
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"list": [
			  {
				"template_id": "wDYzYZVxobJivW9oMpSCpuvACOfJXQIoKUm0PY397Tc",
				"title": "购买成功通知",
				"content": "购买地点{{keyword1.DATA}}\n购买时间{{keyword2.DATA}}\n物品名称{{keyword3.DATA}}\n",
				"example": "购买地点：TIT造舰厂\n购买时间：2016年6月6日\n物品名称：咖啡\n"
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	err := getTemplateList(ts.URL+apiGetTemplateList, "mock-access-token", offset, count, new(GetTemplateListResponse))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendTemplateMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSendTemplateMessage {
			t.Fatalf("Except to path '%s',get '%s'", apiSendTemplateMessage, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ToUser     string `json:"touser"`
			TemplateID string `json:"template_id"`
			Page       string `json:"page"`
			FormID     string `json:"form_id"`
			Data       map[string]struct {
				Value string `json:"value"`
			} `json:"data"`
			EmphasisKeyword string `json:"emphasis_keyword"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ToUser == "" {
			t.Error("Response column to_user can not be empty")
		}

		if params.TemplateID == "" {
			t.Error("Response column template_id can not be empty")
		}

		if params.Page == "" {
			t.Error("Response column page can not be empty")
		}

		if params.FormID == "" {
			t.Error("Response column form_id can not be empty")
		}

		if len(params.Data) == 0 {
			t.Error("Response column data can not be empty")
		}

		for k, v := range params.Data {
			if v.Value == "" {
				t.Errorf("Response column data.%s can not be empty", k)
			}
		}

		if params.EmphasisKeyword == "" {
			t.Error("Response column emphasis_keyword can not be empty")
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

	sender := TempMsgSender{
		ToUser:     "mock-open-id",
		TemplateID: "mock-template-id",
		Page:       "mock/page?foo=bar",
		FormID:     "mock-form-id",
		Data: TempMsgData{
			"mock-key-word1": TempMsgKeyword{"mock-template-value1"},
			"mock-key-word2": TempMsgKeyword{"mock-template-value2"},
		},
		EmphasisKeyword: "mock-open-id",
	}

	_, err := sender.send(ts.URL+apiSendTemplateMessage, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}
