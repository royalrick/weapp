package weapp

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddExpressOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiAddExpressOrder {
			t.Fatalf("Except to path '%s',get '%s'", apiAddExpressOrder, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			AddSource    uint8  `json:"add_source"`
			WXAppID      string `json:"wx_appid"`
			OrderID      string `json:"order_id"`
			OpenID       string `json:"openid"`
			DeliveryID   string `json:"delivery_id"`
			BizID        string `json:"biz_id"`
			CustomRemark string `json:"custom_remark"`
			Sender       struct {
				Name     string `json:"name"`
				Tel      string `json:"tel"`
				Mobile   string `json:"mobile"`
				Company  string `json:"company"`
				PostCode string `json:"post_code"`
				Country  string `json:"country"`
				Province string `json:"province"`
				City     string `json:"city"`
				Area     string `json:"area"`
				Address  string `json:"address"`
			} `json:"sender"`
			Receiver struct {
				Name     string `json:"name"`
				Tel      string `json:"tel"`
				Mobile   string `json:"mobile"`
				Company  string `json:"company"`
				PostCode string `json:"post_code"`
				Country  string `json:"country"`
				Province string `json:"province"`
				City     string `json:"city"`
				Area     string `json:"area"`
				Address  string `json:"address"`
			} `json:"receiver"`
			Cargo struct {
				Count      uint    `json:"count"`
				Weight     float64 `json:"weight"`
				SpaceX     float64 `json:"space_x"`
				SpaceY     float64 `json:"space_y"`
				SpaceZ     float64 `json:"space_z"`
				DetailList []struct {
					Name  string `json:"name"`
					Count uint   `json:"count"`
				} `json:"detail_list"`
			} `json:"cargo"`
			Shop struct {
				WXAPath    string `json:"wxa_path"`
				IMGUrl     string `json:"img_url"`
				GoodsName  string `json:"goods_name"`
				GoodsCount uint   `json:"goods_count"`
			} `json:"shop"`
			Insured struct {
				Used  InsureStatus `json:"use_insured"`
				Value uint         `json:"insured_value"`
			} `json:"insured"`
			Service struct {
				Type uint8  `json:"service_type"`
				Name string `json:"service_name"`
			} `json:"service"`
			ExpectTime uint `json:"expect_time"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.AddSource == 2 && params.WXAppID == "" {
			t.Error("param wx_appid can not be empty")
		}
		if params.AddSource != 2 && params.OpenID == "" {
			t.Error("param openid can not be empty")
		}
		if params.OrderID == "" {
			t.Error("param order_id can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("param delivery_id can not be empty")
		}

		if params.BizID == "" {
			t.Error("param biz_id can not be empty")
		}

		if params.Sender.Name == "" {
			t.Error("param sender.name can not be empty")
		}
		if params.Sender.Province == "" {
			t.Error("param sender.province can not be empty")
		}
		if params.Sender.City == "" {
			t.Error("param sender.city can not be empty")
		}
		if params.Sender.Area == "" {
			t.Error("param sender.area can not be empty")
		}
		if params.Sender.Address == "" {
			t.Error("param sender.address can not be empty")
		}
		if params.Receiver.Name == "" {
			t.Error("param receiver.name can not be empty")
		}
		if params.Receiver.Province == "" {
			t.Error("param receiver.province can not be empty")
		}
		if params.Receiver.City == "" {
			t.Error("param receiver.city can not be empty")
		}
		if params.Receiver.Area == "" {
			t.Error("param receiver.area can not be empty")
		}
		if params.Receiver.Address == "" {
			t.Error("param receiver.address can not be empty")
		}

		if params.Cargo.Count == 0 {
			t.Error("param cargo.count can not be zero")
		}
		if params.Cargo.Weight == 0 {
			t.Error("param cargo.weight can not be zero")
		}
		if params.Cargo.SpaceX == 0 {
			t.Error("param cargo.spaceX can not be zero")
		}
		if params.Cargo.SpaceY == 0 {
			t.Error("param cargo.spaceY can not be zero")
		}
		if params.Cargo.SpaceZ == 0 {
			t.Error("param cargo.spaceZ can not be zero")
		}
		if len(params.Cargo.DetailList) == 0 {
			t.Error("param cargo.detailList can not be empty")
		} else {
			if (params.Cargo.DetailList[0].Name) == "" {
				t.Error("param cargo.detailList.name can not be empty")
			}
			if (params.Cargo.DetailList[0].Count) == 0 {
				t.Error("param cargo.detailList.count can not be zero")
			}
		}
		if params.Shop.WXAPath == "" {
			t.Error("param shop.wxa_path can not be empty")
		}
		if params.Shop.IMGUrl == "" {
			t.Error("param shop.img_url can not be empty")
		}
		if params.Shop.GoodsName == "" {
			t.Error("param shop.goods_name can not be empty")
		}
		if params.Shop.GoodsCount == 0 {
			t.Error("param shop.goods_count can not be zero")
		}
		if params.Insured.Used == 0 {
			t.Error("param insured.use_insured can not be zero")
		}
		if params.Service.Name == "" {
			t.Error("param Service.service_name can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 9300501,
			"errmsg": "delivery logic fail",
			"delivery_resultcode": 10002,
			"delivery_resultmsg": "客户密码不正确",
			"order_id": "01234567890123456789",
			"waybill_id": "123456789",
			"waybill_data": [
			  {
				"key": "SF_bagAddr",
				"value": "广州"
			  },
			  {
				"key": "SF_mark",
				"value": "101- 07-03 509"
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	creator := ExpressOrderCreator{
		AddSource: 0,
		ExpressOrder: ExpressOrder{
			OrderID:      "01234567890123456789",
			OpenID:       "oABC123456",
			DeliveryID:   "SF",
			BizID:        "xyz",
			CustomRemark: "易碎物品",
			Sender: ExpreseeUserInfo{
				"张三",
				"020-88888888",
				"18666666666",
				"公司名",
				"123456",
				"中国",
				"广东省",
				"广州市",
				"海珠区",
				"XX路XX号XX大厦XX栋XX",
			},
			Receiver: ExpreseeUserInfo{
				"王小蒙",
				"020-77777777",
				"18610000000",
				"公司名",
				"654321",
				"中国",
				"广东省",
				"广州市",
				"天河区",
				"XX路XX号XX大厦XX栋XX",
			},
			Shop: ExpressShop{
				"/index/index?from=waybill&id=01234567890123456789",
				"https://mmbiz.qpic.cn/mmbiz_png/OiaFLUqewuIDNQnTiaCInIG8ibdosYHhQHPbXJUrqYSNIcBL60vo4LIjlcoNG1QPkeH5GWWEB41Ny895CokeAah8A/640",
				"一千零一夜钻石包&爱马仕铂金包",
				2,
			},
			Cargo: ExpressCargo{
				2,
				5.5,
				30.5,
				20,
				20,
				[]CargoDetail{
					{
						"一千零一夜钻石包",
						1,
					},
					{
						"爱马仕铂金包",
						1,
					},
				},
			},
			Insured: ExpressInsure{
				1,
				10000,
			},
			Service: ExpressService{
				0,
				"标准快递",
			},
		},
	}
	_, err := creator.create(ts.URL+apiAddExpressOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCancelExpressOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiCancelExpressOrder {
			t.Fatalf("Except to path '%s',get '%s'", apiCancelExpressOrder, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			OrderID    string `json:"order_id"`
			OpenID     string `json:"openid"`
			DeliveryID string `json:"delivery_id"`
			WaybillID  string `json:"waybill_id"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.OrderID == "" {
			t.Error("param order_id can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("param delivery_id can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("param waybill_id can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	canceler := ExpressOrderCanceler{
		OrderID:    "01234567890123456789",
		OpenID:     "oABC123456",
		DeliveryID: "SF",
		WaybillID:  "123456789",
	}
	_, err := canceler.cancel(ts.URL+apiCancelExpressOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllDelivery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			t.Fatalf("Expect 'GET' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetAllDelivery {
			t.Fatalf("Except to path '%s',get '%s'", apiGetAllDelivery, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"count": 8,
			"data": [
			  {
				"delivery_id": "BEST",
				"delivery_name": "百世快递"
			  },
			  {
				"delivery_id": "EMS",
				"delivery_name": "中国邮政速递物流"
			  },
			  {
				"delivery_id": "OTP",
				"delivery_name": "承诺达特快"
			  },
			  {
				"delivery_id": "PJ",
				"delivery_name": "品骏物流"
			  },
			  {
				"delivery_id": "SF",
				"delivery_name": "顺丰速运"
			  },
			  {
				"delivery_id": "YTO",
				"delivery_name": "圆通速递"
			  },
			  {
				"delivery_id": "YUNDA",
				"delivery_name": "韵达快递"
			  },
			  {
				"delivery_id": "ZTO",
				"delivery_name": "中通快递"
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getAllDelivery(ts.URL+apiGetAllDelivery, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetExpressOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != apiGetExpressOrder {
			t.Fatalf("Except to path '%s',get '%s'", apiGetExpressOrder, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			OrderID    string `json:"order_id"`
			OpenID     string `json:"openid"`
			DeliveryID string `json:"delivery_id"`
			WaybillID  string `json:"waybill_id"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.OrderID == "" {
			t.Error("param order_id can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("param delivery_id can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("param waybill_id can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	canceler := ExpressOrderGetter{
		OrderID:    "01234567890123456789",
		OpenID:     "oABC123456",
		DeliveryID: "SF",
		WaybillID:  "123456789",
	}
	_, err := canceler.get(ts.URL+apiGetExpressOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetExpressPath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		expectedPath := "/cgi-bin/express/business/path/get"
		if path != expectedPath {
			t.Fatalf("Except to path '%s',get '%s'", expectedPath, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			OrderID    string `json:"order_id"`
			OpenID     string `json:"openid"`
			DeliveryID string `json:"delivery_id"`
			WaybillID  string `json:"waybill_id"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.OrderID == "" {
			t.Error("param order_id can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("param delivery_id can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("param waybill_id can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"openid": "OPENID",
			"delivery_id": "SF",
			"waybill_id": "12345678901234567890",
			"path_item_num": 3,
			"path_item_list": [
			  {
				"action_time": 1533052800,
				"action_type": 100001,
				"action_msg": "快递员已成功取件"
			  },
			  {
				"action_time": 1533062800,
				"action_type": 200001,
				"action_msg": "快件已到达xxx集散中心，准备发往xxx"
			  },
			  {
				"action_time": 1533072800,
				"action_type": 300001,
				"action_msg": "快递员已出发，联系电话xxxxxx"
			  }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	getter := ExpressPathGetter{
		OrderID:    "01234567890123456789",
		OpenID:     "oABC123456",
		DeliveryID: "SF",
		WaybillID:  "123456789",
	}
	_, err := getter.get(ts.URL+apiGetExpressPath, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPrinter(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			t.Fatalf("Expect 'GET' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		expectedPath := "/cgi-bin/express/business/printer/getall"
		if path != expectedPath {
			t.Fatalf("Except to path '%s',get '%s'", expectedPath, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"count": 2,
			"openid": [
			  "oABC",
			  "oXYZ"
			],
			"tagid_list": [
			  "123",
			  "456"
			]
		   }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getPrinter(ts.URL+apiGetPrinter, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetQuota(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		expectedPath := "/cgi-bin/express/business/quota/get"
		if path != expectedPath {
			t.Fatalf("Except to path '%s',get '%s'", expectedPath, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			DeliveryID string `json:"delivery_id"`
			BizID      string `json:"biz_id"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.DeliveryID == "" {
			t.Error("param delivery_id can not be empty")
		}
		if params.BizID == "" {
			t.Error("param biz_id can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"quota_num": 210
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	getter := QuotaGetter{
		DeliveryID: "YTO",
		BizID:      "xyz",
	}

	_, err := getter.get(ts.URL+apiGetQuota, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestOnPathUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnExpressPathUpdate(func(mix *ExpressPathUpdateResult) {
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

			if mix.Event != "add_express_path" {
				t.Error("Unexpected message event")
			}

			if mix.DeliveryID == "" {
				t.Error("DeliveryID can not be empty")
			}
			if mix.WayBillID == "" {
				t.Error("WayBillID can not be empty")
			}
			if mix.Version == 0 {
				t.Error("Version can not be zero")
			}
			if mix.Count == 0 {
				t.Error("Count can not be zero")
			}

			if len(mix.Actions) > 0 {
				if mix.Actions[0].ActionTime == 0 {
					t.Error("Actions.ActionTime can not be zero")
				}
				if mix.Actions[0].ActionType == 0 {
					t.Error("Actions.ActionType can not be zero")
				}
				if mix.Actions[0].ActionMsg == "" {
					t.Error("Actions.ActionMsg can not be empty")
				}
			}
		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	xmlData := `<xml>
		<ToUserName><![CDATA[toUser]]></ToUserName>
		<FromUserName><![CDATA[fromUser]]></FromUserName>
		<CreateTime>1546924844</CreateTime>
		<MsgType><![CDATA[event]]></MsgType>
		<Event><![CDATA[add_express_path]]></Event>
		<DeliveryID><![CDATA[SF]]></DeliveryID>
		<WayBillId><![CDATA[123456789]]></WayBillId>
		<Version>3</Version>
		<Count>3</Count>
		<Actions>
		<ActionTime>1546924840</ActionTime>
		<ActionType>100001</ActionType>
		<ActionMsg><![CDATA[小哥A揽件成功]]></ActionMsg>
		</Actions>
		<Actions>
		<ActionTime>1546924841</ActionTime>
		<ActionType>200001</ActionType>
		<ActionMsg><![CDATA[到达广州集包地]]></ActionMsg>
		</Actions>
		<Actions>
		<ActionTime>1546924842</ActionTime>
		<ActionType>200001</ActionType>
		<ActionMsg><![CDATA[运往目的地]]></ActionMsg>
		</Actions>
	</xml>`
	res, err := http.Post(ts.URL, "application/xml", strings.NewReader(xmlData))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	jsonData := `{
		"ToUserName": "toUser",
		"FromUserName": "fromUser",
		"CreateTime": 1546924844,
		"MsgType": "event",
		"Event": "add_express_path",
		"DeliveryID": "SF",
		"WayBillId": "123456789",
		"Version": 2,
		"Count": 3,
		"Actions": [
		  {
			"ActionTime": 1546924840,
			"ActionType": 100001,
			"ActionMsg": "小哥A揽件成功"
		  },
		  {
			"ActionTime": 1546924841,
			"ActionType": 200001,
			"ActionMsg": "到达广州集包地"
		  },
		  {
			"ActionTime": 1546924842,
			"ActionType": 200001,
			"ActionMsg": "运往目的地"
		  }
		]
	  }`
	res, err = http.Post(ts.URL, "application/json", strings.NewReader(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
}

func TestTestUpdateOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		expectedPath := "/cgi-bin/express/business/test_update_order"
		if path != expectedPath {
			t.Fatalf("Except to path '%s',get '%s'", expectedPath, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			BizID      string `json:"biz_id"`      // 商户id,需填test_biz_id
			OrderID    string `json:"order_id"`    //	订单ID，下单成功时返回
			WaybillID  string `json:"waybill_id"`  // 运单 ID
			DeliveryID string `json:"delivery_id"` // 快递公司 ID
			ActionTime uint   `json:"action_time"` // 轨迹变化 Unix 时间戳
			ActionType int    `json:"action_type"` // 轨迹变化类型
			ActionMsg  string `json:"action_msg"`  // 轨迹变化具体信息说明，展示在快递轨迹详情页中。若有手机号码，则直接写11位手机号码。使用UTF-8编码。
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.DeliveryID == "" {
			t.Error("param delivery_id can not be empty")
		}
		if params.OrderID == "" {
			t.Error("param order_id can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("param waybill_id can not be empty")
		}

		if params.BizID == "" {
			t.Error("param biz_id can not be empty")
		}
		if params.ActionMsg == "" {
			t.Error("param action_msg can not be empty")
		}
		if params.ActionTime == 0 {
			t.Error("param action_time can not be empty")
		}
		if params.ActionType == 0 {
			t.Error("param action_type can not be empty")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	params := `{
		"biz_id": "test_biz_id",
		"order_id": "xxxxxxxxxxxx",
		"delivery_id": "TEST",
		"waybill_id": "xxxxxxxxxx",
		"action_time": 123456789,
		"action_type": 100001,
		"action_msg": "揽件阶段"
	  }`

	tester := new(UpdateExpressOrderTester)
	err := json.Unmarshal([]byte(params), tester)
	if err != nil {
		t.Error(err)
	}

	_, err = tester.test(ts.URL+apiTestUpdateOrder, "mock-access-token")
	if err != nil {
		t.Error(err)
	}
}

func TestUpdatePrinter(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != "/cgi-bin/express/business/printer/update" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			OpenID    string   `json:"openid"`      // 打印员 openid
			Type      BindType `json:"update_type"` // 更新类型
			TagIDList string   `json:"tagid_list"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.OpenID == "" {
			t.Error("param openid can not be empty")
		}
		if params.Type == "" {
			t.Error("param update_type can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"errcode": 0,
			"errmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()
	params := `{
		"openid": "oJ4v0wRAfiXcnIbM3SgGEUkTw3Qw",
		"update_type": "bind",
		"tagid_list": "123,456"
	  }`
	updater := new(PrinterUpdater)
	err := json.Unmarshal([]byte(params), updater)
	if err != nil {
		t.Fatal(err)
	}

	_, err = updater.update(ts.URL+apiUpdatePrinter, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}
