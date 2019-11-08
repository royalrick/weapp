package weapp

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddNearbyPoi(t *testing.T) {

	localServer := http.NewServeMux()
	localServer.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnAddNearbyPoi(func(mix *AddNearbyPoiResult) {
			if mix.ToUserName == "" {
				t.Error("ToUserName can not be empty")
			}

			if mix.FromUserName == "" {
				t.Error("FromUserName can not be empty")
			}
			if mix.CreateTime == 0 {
				t.Error("CreateTime can not be zero")
			}
			if mix.MsgType != "event" {
				t.Error("Unexpected message type")
			}

			if mix.Event != "add_nearby_poi_audit_info" {
				t.Error("Unexpected message event")
			}

			if mix.AuditID == 0 {
				t.Error("audit_id can not be zero")
			}
			if mix.Status == 0 {
				t.Error("status can not be zero")
			}
			if mix.Reason == "" {
				t.Error("reason can not be empty")
			}
			if mix.PoiID == 0 {
				t.Error("poi_id can not be zero")
			}

		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	})

	tls := httptest.NewServer(localServer)
	defer tls.Close()

	remoteServer := http.NewServeMux()
	remoteServer.HandleFunc(apiAddNearbyPoi, func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiAddNearbyPoi {
			t.Fatalf("Except to path '%s',get '%s'", apiAddNearbyPoi, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			IsCommNearby      string `json:"is_comm_nearby"`
			PicList           string `json:"pic_list"`           // 门店图片，最多9张，最少1张，上传门店图片如门店外景、环境设施、商品服务等，图片将展示在微信客户端的门店页。图片链接通过文档https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1444738729中的《上传图文消息内的图片获取URL》接口获取。必填，文件格式为bmp、png、jpeg、jpg或gif，大小不超过5M pic_list是字符串，内容是一个json！
			ServiceInfos      string `json:"service_infos"`      // 必服务标签列表 选填，需要填写服务标签ID、APPID、对应服务落地页的path路径，详细字段格式见下方示例
			StoreName         string `json:"store_name"`         // 门店名字 必填，门店名称需按照所选地理位置自动拉取腾讯地图门店名称，不可修改，如需修改请重现选择地图地点或重新创建地点
			Hour              string `json:"hour"`               // 营业时间，格式11:11-12:12 必填
			Credential        string `json:"credential"`         // 资质号 必填, 15位营业执照注册号或9位组织机构代码
			Address           string `json:"address"`            // 地址 必填
			CompanyName       string `json:"company_name"`       // 主体名字 必填
			QualificationList string `json:"qualification_list"` // 证明材料 必填 如果company_name和该小程序主体不一致，需要填qualification_list，详细规则见附近的小程序使用指南-如何证明门店的经营主体跟公众号或小程序帐号主体相关http://kf.qq.com/faq/170401MbUnim17040122m2qY.html
			KFInfo            string `json:"kf_info"`            // 客服信息 选填，可自定义服务头像与昵称，具体填写字段见下方示例kf_info pic_list是字符串，内容是一个json！
			PoiID             string `json:"poi_id"`             // 如果创建新的门店，poi_id字段为空 如果更新门店，poi_id参数则填对应门店的poi_id 选填
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.IsCommNearby != "1" {
			t.Error("param pic_list is invalid")
		}
		if params.PicList == "" {
			t.Error("param pic_list can not be empty")
		}
		if params.ServiceInfos == "" {
			t.Error("param service_infos can not be empty")
		}
		if params.StoreName == "" {
			t.Error("param store_name can not be empty")
		}
		if params.Hour == "" {
			t.Error("param hour can not be empty")
		}
		if params.Credential == "" {
			t.Error("param credential can not be empty")
		}
		if params.Address == "" {
			t.Error("param address can not be empty")
		}
		if params.CompanyName == "" {
			t.Error("param company_name can not be empty")
		}
		if params.QualificationList == "" {
			t.Error("param qualification_list can not be empty")
		}
		if params.KFInfo == "" {
			t.Error("param kf_info can not be empty")
		}
		if params.PoiID == "" {
			t.Error("param poi_id can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "ok",
			"data":{
				"audit_id": "xxxxx",
				"poi_id": "xxxxx",
				"related_credential":"xxxxx"
			}
		}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}

		raw = `<xml>
					<ToUserName><![CDATA[gh_4346ac1514d8]]></ToUserName>
					<FromUserName><![CDATA[od1P50M-fNQI5Gcq-trm4a7apsU8]]></FromUserName>
					<CreateTime>1488856741</CreateTime>
					<MsgType><![CDATA[event]]></MsgType>
					<Event><![CDATA[add_nearby_poi_audit_info]]></Event>
					<audit_id>11111</audit_id>
					<status>3</status>
					<reason><![CDATA[xxx]]></reason>
					<poi_id>111111</poi_id>
				</xml>`
		reader := strings.NewReader(raw)
		http.Post(tls.URL+"/notify", "application/xml", reader)
	})
	trs := httptest.NewServer(remoteServer)
	defer trs.Close()

	poi := NearbyPoi{
		PicList: PicList{[]string{"first-mock-picture-url", "second-mock-picture-url", "third-mock-picture-url"}},
		ServiceInfos: ServiceInfos{[]ServiceInfo{
			{1, 1, "mock-name", "mock-app-id", "mock-path"},
		}},
		StoreName:         "mock-store-name",
		Hour:              "11:11-12:12",
		Credential:        "mock-credential",
		Address:           "mock-address",                         // 地址 必填
		CompanyName:       "mock-company-name",                    // 主体名字 必填
		QualificationList: "mock-qualification-list",              // 证明材料 必填 如果company_name和该小程序主体不一致，需要填qualification_list，详细规则见附近的小程序使用指南-如何证明门店的经营主体跟公众号或小程序帐号主体相关http://kf.qq.com/faq/170401MbUnim17040122m2qY.html
		KFInfo:            KFInfo{true, "kf-head-img", "kf-name"}, // 客服信息 选填，可自定义服务头像与昵称，具体填写字段见下方示例kf_info pic_list是字符串，内容是一个json！
		PoiID:             "mock-poi-id",                          // 如果创建新的门店，poi_id字段为空 如果更新门店，poi_id参数则填对应门店的poi_id 选填
	}
	_, err := poi.add(trs.URL+apiAddNearbyPoi, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteNearbyPoi(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiDeleteNearbyPoi {
			t.Fatalf("Except to path '%s',get '%s'", apiDeleteNearbyPoi, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			PoiID string `json:"poi_id"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.PoiID == "" {
			t.Error("Response column poi_id can not be empty")
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

	_, err := deleteNearbyPoi(ts.URL+apiDeleteNearbyPoi, "mock-access-token", "mock-poi-id")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetNearbyPoiList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetNearbyPoiList {
			t.Fatalf("Except to path '%s',get '%s'", apiGetNearbyPoiList, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Page uint `json:"page"`
			Rows uint `json:"page_rows"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Page == 0 {
			t.Error("Response column page can not be empty")
		}

		if params.Rows == 0 {
			t.Error("Response column page_rows can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 0,
			"errmsg": "",
			"data": {
			   "left_apply_num": 9,
			   "max_apply_num": 10,
			   "data": "{\"poi_list\": [{\"poi_id\": \"123456\",\"qualification_address\": \"广东省广州市海珠区新港中路123号\",\"qualification_num\": \"123456789-1\",\"audit_status\": 3,\"display_status\": 0,\"refuse_reason\": \"\"}]}"
			}
		 }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getNearbyPoiList(ts.URL+apiGetNearbyPoiList, "mock-access-token", 1, 10)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSetNearbyPoiShowStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiSetNearbyPoiShowStatus {
			t.Fatalf("Except to path '%s',get '%s'", apiSetNearbyPoiShowStatus, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			PoiID  string `json:"poi_id"`
			Status uint8  `json:"status"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.PoiID == "" {
			t.Error("Response column poi_id can not be empty")
		}

		if params.Status == 0 {
			t.Error("Response column status can not be zero")
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

	_, err := setNearbyPoiShowStatus(ts.URL+apiSetNearbyPoiShowStatus, "mock-access-token", "mock-poi-id", ShowNearbyPoi)
	if err != nil {
		t.Fatal(err)
	}
}
