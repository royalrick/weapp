package weapp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateActivityID(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			t.Fatalf("Expect 'GET' get '%s'", r.Method)
		}

		realPath := r.URL.EscapedPath()
		expectPath := "/cgi-bin/message/wxopen/activityid/create"
		if realPath != expectPath {
			t.Fatalf("Expect to path '%s',get '%s'", expectPath, realPath)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"expiration_time": 1000,
			"activity_id": "ok",
			"errcode": 0,
			"errmsg": "ok"
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	res, err := createActivityID(ts.URL+apiCreateActivityID, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.ActivityID == "" {
		t.Error("Response column activity_id can not be empty")
	}

	if res.ExpirationTime == 0 {
		t.Error("Response column expiration_time can not be zero")
	}
}

func TestSetUpdatableMsg(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		realPath := r.URL.EscapedPath()
		expectPath := "/cgi-bin/message/wxopen/updatablemsg/send"
		if realPath != expectPath {
			t.Fatalf("Expect to path '%s',get '%s'", expectPath, realPath)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ActivityID   string `json:"activity_id"`
			TargetState  uint8  `json:"target_state"`
			TemplateInfo struct {
				ParameterList []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"parameter_list"`
			} `json:"template_info"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ActivityID == "" {
			t.Fatal("param activity_id can not be empty")
		}

		if len(params.TemplateInfo.ParameterList) == 0 {
			t.Fatal("param template_info.parameter_list can not be empty")
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

	setter := UpdatableMsgSetter{
		"mock-activity-id",
		UpdatableMsgJoining,
		UpdatableMsgTempInfo{
			[]UpdatableMsgParameter{
				{UpdatableMsgParamMemberCount, "mock-parameter-value-number"},
				{UpdatableMsgParamRoomLimit, "mock-parameter-value-number"},
			},
		},
	}

	_, err := setter.set(ts.URL+apiSetUpdatableMsg, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}
