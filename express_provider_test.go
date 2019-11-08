package weapp

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetContact(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			t.Fatalf("Expect 'POST' get '%s'", r.Method)
		}

		path := r.URL.EscapedPath()
		if path != "/cgi-bin/express/delivery/contact/get" {
			t.Error("Invalid request path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Token     string `json:"token"`
			WaybillID string `json:"waybill_id"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.Token == "" {
			t.Error("param token can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("param waybill_id can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		raw := `{
			"waybill_id": "12345678901234567890",
			"sender": {
			  "address": "广东省广州市海珠区XX路XX号XX大厦XX栋XX",
			  "name": "张三",
			  "tel": "020-88888888",
			  "mobile": "18666666666"
			},
			"receiver": {
			  "address": "广东省广州市天河区XX路XX号XX大厦XX栋XX",
			  "name": "王小蒙",
			  "tel": "029-77777777",
			  "mobile": "18610000000"
			}
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	_, err := getContact(ts.URL+apiGetContact, "mock-access-token", "mock-token", "mock-wat-bill-id")
	if err != nil {
		t.Fatal(err)
	}
}

func TestOnAddExpressOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnAddExpressOrder(func(result *AddExpressOrderResult) *AddExpressOrderReturn {
			if result.ToUserName == "" {
				t.Error("ToUserName can not be empty")
			}
			if result.FromUserName == "" {
				t.Error("FromUserName can not be empty")
			}
			if result.CreateTime == 0 {
				t.Error("CreateTime can not be zero")
			}
			if result.MsgType != "event" {
				t.Error("Unexpected message type")
			}
			if result.Event != "add_waybill" {
				t.Error("Unexpected message event")
			}
			if result.Token == "" {
				t.Error("Result column 'Token' can not be empty")
			}
			if result.OrderID == "" {
				t.Error("Result column 'OrderID' can not be empty")
			}
			if result.BizID == "" {
				t.Error("Result column 'BizID' can not be empty")
			}
			if result.BizPwd == "" {
				t.Error("Result column 'BizPwd' can not be empty")
			}
			if result.ShopAppID == "" {
				t.Error("Result column 'ShopAppID' can not be empty")
			}
			if result.WayBillID == "" {
				t.Error("Result column 'WayBillID' can not be empty")
			}
			if result.Remark == "" {
				t.Error("Result column 'Remark' can not be empty")
			}

			if result.Sender.Name == "" {
				t.Error("Result column 'Sender.Name' can not be empty")
			}
			if result.Sender.Tel == "" {
				t.Error("Result column 'Sender.Tel' can not be empty")
			}
			if result.Sender.Mobile == "" {
				t.Error("Result column 'Sender.Mobile' can not be empty")
			}
			if result.Sender.Company == "" {
				t.Error("Result column 'Sender.Company' can not be empty")
			}
			if result.Sender.PostCode == "" {
				t.Error("Result column 'Sender.PostCode' can not be empty")
			}
			if result.Sender.Country == "" {
				t.Error("Result column 'Sender.Country' can not be empty")
			}
			if result.Sender.Province == "" {
				t.Error("Result column 'Sender.Province' can not be empty")
			}
			if result.Sender.City == "" {
				t.Error("Result column 'Sender.City' can not be empty")
			}
			if result.Sender.Area == "" {
				t.Error("Result column 'Sender.Area' can not be empty")
			}
			if result.Sender.Address == "" {
				t.Error("Result column 'Sender.Address' can not be empty")
			}
			if result.Receiver.Name == "" {
				t.Error("Result column 'Receiver.Name' can not be empty")
			}
			if result.Receiver.Tel == "" {
				t.Error("Result column 'Receiver.Tel' can not be empty")
			}
			if result.Receiver.Mobile == "" {
				t.Error("Result column 'Receiver.Mobile' can not be empty")
			}
			if result.Receiver.Company == "" {
				t.Error("Result column 'Receiver.Company' can not be empty")
			}
			if result.Receiver.PostCode == "" {
				t.Error("Result column 'Receiver.PostCode' can not be empty")
			}
			if result.Receiver.Country == "" {
				t.Error("Result column 'Receiver.Country' can not be empty")
			}
			if result.Receiver.Province == "" {
				t.Error("Result column 'Receiver.Province' can not be empty")
			}
			if result.Receiver.City == "" {
				t.Error("Result column 'Receiver.City' can not be empty")
			}
			if result.Receiver.Area == "" {
				t.Error("Result column 'Receiver.Area' can not be empty")
			}
			if result.Receiver.Address == "" {
				t.Error("Result column 'Receiver.Address' can not be empty")
			}
			if result.Cargo.Weight == 0 {
				t.Error("Result column 'Cargo.Weight' can not be zero")
			}
			if result.Cargo.SpaceX == 0 {
				t.Error("Result column 'Cargo.SpaceX' can not be zero")
			}
			if result.Cargo.SpaceY == 0 {
				t.Error("Result column 'Cargo.SpaceY' can not be zero")
			}
			if result.Cargo.SpaceZ == 0 {
				t.Error("Result column 'Cargo.SpaceZ' can not be zero")
			}
			if result.Cargo.Count == 0 {
				t.Error("Result column 'Cargo.Count' can not be zero")
			}
			if result.Insured.Used == 0 {
				t.Error("Result column 'Insured.Used' can not be zero")
			}
			if result.Insured.Value == 0 {
				t.Error("Result column 'Insured.Value' can not be zero")
			}
			if result.Service.Type == 0 {
				t.Error("Result column 'Service.Type' can not be zero")
			}
			if result.Service.Name == "" {
				t.Error("Result column 'Service.Name' can not be empty")
			}

			res := AddExpressOrderReturn{
				CommonServerReturn: CommonServerReturn{
					"oABCD", "gh_abcdefg", 1533042556, "event", "add_waybill", 1, "success",
				},
				Token:       "1234ABC234523451",
				OrderID:     "012345678901234567890123456789",
				BizID:       "xyz",
				WayBillID:   "123456789",
				WaybillData: "##ZTO_bagAddr##广州##ZTO_mark##888-666-666##",
			}

			return &res
		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	xmlData := `<xml>
	<ToUserName><![CDATA[gh_abcdefg]]></ToUserName>
	<FromUserName><![CDATA[oABCD]]></FromUserName>
	<CreateTime>1533042556</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[add_waybill]]></Event>
	<Token>1234ABC234523451</Token>
	<OrderID><![CDATA[012345678901234567890123456789]]></OrderID>
	<BizID><![CDATA[xyz]]></BizID>
	<BizPwd><![CDATA[xyz123]]></BizPwd>
	<ShopAppID><![CDATA[wxABCD]]></ShopAppID>
	<WayBillID><![CDATA[123456789]]></WayBillID>
	<Remark><![CDATA[易碎物品]]></Remark>
	<Sender>
		<Name><![CDATA[张三]]></Name>
		<Tel><![CDATA[020-88888888]]></Tel>
		<Mobile><![CDATA[18666666666]]></Mobile>
		<Company><![CDATA[公司名]]></Company>
		<PostCode><![CDATA[123456]]></PostCode>
		<Country><![CDATA[中国]]></Country>
		<Province><![CDATA[广东省]]></Province>
		<City><![CDATA[广州市]]></City>
		<Area><![CDATA[海珠区]]></Area>
		<Address><![CDATA[XX路XX号XX大厦XX栋XX]]></Address>
	</Sender>
	<Receiver>
		<Name><![CDATA[王小蒙]]></Name>
		<Tel><![CDATA[029-77777777]]></Tel>
		<Mobile><![CDATA[18610000000]]></Mobile>
		<Company><![CDATA[公司名]]></Company>
		<PostCode><![CDATA[654321]]></PostCode>
		<Country><![CDATA[中国]]></Country>
		<Province><![CDATA[广东省]]></Province>
		<City><![CDATA[广州市]]></City>
		<Area><![CDATA[天河区]]></Area>
		<Address><![CDATA[XX路XX号XX大厦XX栋XX]]></Address>
	</Receiver>
	<Cargo>
		<Weight>1.2</Weight>
		<Space_X>20.5</Space_X>
		<Space_Y>15.0</Space_Y>
		<Space_Z>10.0</Space_Z>
		<Count>2</Count>
		<DetailList>
			<Name><![CDATA[一千零一夜钻石包]]></Name>
			<Count>1</Count>
		</DetailList>
		<DetailList>
			<Name><![CDATA[爱马仕柏金钻石包]]></Name>
			<Count>1</Count>
		</DetailList>
	</Cargo>
	<Insured>
		<UseInsured>1</UseInsured>
		<InsuredValue>10000</InsuredValue>
	</Insured>
	<Service>
		<ServiceType>123</ServiceType>
		<ServiceName><![CDATA[标准快递]]></ServiceName>
	</Service>
  </xml>`
	xmlResp, err := http.Post(ts.URL, "application/xml", strings.NewReader(xmlData))
	if err != nil {
		t.Error(err)
	}
	defer xmlResp.Body.Close()
	res := new(AddExpressOrderReturn)
	if err := xml.NewDecoder(xmlResp.Body).Decode(res); err != nil {
		t.Error(err)
	}

	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event == "" {
		t.Error("Response column 'Event' can not be empty")
	}
	if res.Token == "" {
		t.Error("Response column 'Token' can not be empty")
	}
	if res.OrderID == "" {
		t.Error("Response column 'OrderID' can not be empty")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.WayBillID == "" {
		t.Error("Response column 'WayBillID' can not be empty")
	}

	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}
	if res.WaybillData == "" {
		t.Error("Response column 'WaybillData' can not be empty")
	}

	jsonData := `{
		"ToUserName": "gh_abcdefg",
		"FromUserName": "oABCD",
		"CreateTime": 1533042556,
		"MsgType": "event",
		"Event": "add_waybill",
		"Token": "1234ABC234523451",
		"OrderID": "012345678901234567890123456789",
		"BizID": "xyz",
		"BizPwd": "xyz123",
		"ShopAppID": "wxABCD",
		"WayBillID": "123456789",
		"Remark": "易碎物品",
		"Sender": {
		  "Name": "张三",
		  "Tel": "020-88888888",
		  "Mobile": "18666666666",
		  "Company": "公司名",
		  "PostCode": "123456",
		  "Country": "中国",
		  "Province": "广东省",
		  "City": "广州市",
		  "Area": "海珠区",
		  "Address": "XX路XX号XX大厦XX栋XX"
		},
		"Receiver": {
		  "Name": "王小蒙",
		  "Tel": "029-77777777",
		  "Mobile": "18610000000",
		  "Company": "公司名",
		  "PostCode": "654321",
		  "Country": "中国",
		  "Province": "广东省",
		  "City": "广州市",
		  "Area": "天河区",
		  "Address": "XX路XX号XX大厦XX栋XX"
		},
		"Cargo": {
		  "Weight": 1.2,
		  "Space_X": 20.5,
		  "Space_Y": 15,
		  "Space_Z": 10,
		  "Count": 2,
		  "DetailList": [
			{
			  "Name": "一千零一夜钻石包",
			  "Count": 1
			},
			{
			  "Name": "爱马仕柏金钻石包",
			  "Count": 1
			}
		  ]
		},
		"Insured": {
		  "UseInsured": 1,
		  "InsuredValue": 10000
		},
		"Service": {
		  "ServiceType": 123,
		  "ServiceName": "标准快递"
		}
	  }`

	jsonResp, err := http.Post(ts.URL, "application/json", strings.NewReader(jsonData))
	if err != nil {
		t.Error(err)
	}
	defer jsonResp.Body.Close()
	res = new(AddExpressOrderReturn)
	if err := json.NewDecoder(jsonResp.Body).Decode(res); err != nil {
		t.Error(err)
	}

	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event == "" {
		t.Error("Response column 'Event' can not be empty")
	}
	if res.Token == "" {
		t.Error("Response column 'Token' can not be empty")
	}
	if res.OrderID == "" {
		t.Error("Response column 'OrderID' can not be empty")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.WayBillID == "" {
		t.Error("Response column 'WayBillID' can not be empty")
	}

	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}
	if res.WaybillData == "" {
		t.Error("Response column 'WaybillData' can not be empty")
	}
}

func TestOnCancelExpressOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnCancelExpressOrder(func(result *CancelExpressOrderResult) *CancelExpressOrderReturn {
			if result.ToUserName == "" {
				t.Error("ToUserName can not be empty")
			}
			if result.FromUserName == "" {
				t.Error("FromUserName can not be empty")
			}
			if result.CreateTime == 0 {
				t.Error("CreateTime can not be zero")
			}
			if result.MsgType != "event" {
				t.Error("Unexpected message type")
			}
			if result.Event != "cancel_waybill" {
				t.Error("Unexpected message event")
			}

			if result.OrderID == "" {
				t.Error("Result column 'OrderID' can not be empty")
			}
			if result.BizID == "" {
				t.Error("Result column 'BizID' can not be empty")
			}
			if result.BizPwd == "" {
				t.Error("Result column 'BizPwd' can not be empty")
			}
			if result.ShopAppID == "" {
				t.Error("Result column 'ShopAppID' can not be empty")
			}
			if result.WayBillID == "" {
				t.Error("Result column 'WayBillID' can not be empty")
			}

			res := CancelExpressOrderReturn{
				CommonServerReturn: CommonServerReturn{
					"oABCD", "gh_abcdefg", 1533042556, "event", "cancel_waybill", 1, "success",
				},
				OrderID:   "012345678901234567890123456789",
				BizID:     "xyz",
				WayBillID: "123456789",
			}

			return &res
		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	xmlData := `<xml>
    <ToUserName><![CDATA[gh_abcdefg]]></ToUserName>
    <FromUserName><![CDATA[oABCD]]></FromUserName>
    <CreateTime>1533042556</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[cancel_waybill]]></Event>
    <BizID><![CDATA[xyz]]></BizID>
    <BizPwd><![CDATA[xyz123]]></BizPwd>
    <ShopAppID><![CDATA[wxABCD]]></ShopAppID>
    <OrderID><![CDATA[012345678901234567890123456789]]></OrderID>
    <WayBillID><![CDATA[123456789]]></WayBillID>
</xml>`
	xmlResp, err := http.Post(ts.URL, "application/xml", strings.NewReader(xmlData))
	if err != nil {
		t.Error(err)
	}
	defer xmlResp.Body.Close()
	res := new(CancelExpressOrderReturn)
	if err := xml.NewDecoder(xmlResp.Body).Decode(res); err != nil {
		t.Error(err)
	}

	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event == "" {
		t.Error("Response column 'Event' can not be empty")
	}
	if res.OrderID == "" {
		t.Error("Response column 'OrderID' can not be empty")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.WayBillID == "" {
		t.Error("Response column 'WayBillID' can not be empty")
	}

	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}

	jsonData := `{
		"ToUserName": "gh_abcdefg",
		"FromUserName": "oABCD",
		"CreateTime": 1533042556,
		"MsgType": "event",
		"Event": "cancel_waybill",
		"BizID": "xyz",
		"BizPwd": "xyz123",
		"ShopAppID": "wxABCD",
		"OrderID": "012345678901234567890123456789",
		"WayBillID": "123456789"
	  }`

	jsonResp, err := http.Post(ts.URL, "application/json", strings.NewReader(jsonData))
	if err != nil {
		t.Error(err)
	}
	defer jsonResp.Body.Close()
	res = new(CancelExpressOrderReturn)
	if err := json.NewDecoder(jsonResp.Body).Decode(res); err != nil {
		t.Error(err)
	}

	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event == "" {
		t.Error("Response column 'Event' can not be empty")
	}
	if res.OrderID == "" {
		t.Error("Response column 'OrderID' can not be empty")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.WayBillID == "" {
		t.Error("Response column 'WayBillID' can not be empty")
	}

	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}
}

func TestOnCheckBusiness(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnCheckExpressBusiness(func(result *CheckExpressBusinessResult) *CheckExpressBusinessReturn {
			if result.ToUserName == "" {
				t.Error("ToUserName can not be empty")
			}
			if result.FromUserName == "" {
				t.Error("FromUserName can not be empty")
			}
			if result.CreateTime == 0 {
				t.Error("CreateTime can not be zero")
			}
			if result.MsgType != "event" {
				t.Error("Unexpected message type")
			}
			if result.Event != "check_biz" {
				t.Error("Unexpected message event")
			}

			if result.BizID == "" {
				t.Error("Result column 'BizID' can not be empty")
			}
			if result.BizPwd == "" {
				t.Error("Result column 'BizPwd' can not be empty")
			}
			if result.ShopAppID == "" {
				t.Error("Result column 'ShopAppID' can not be empty")
			}
			if result.ShopName == "" {
				t.Error("Result column 'ShopName' can not be empty")
			}

			if result.ShopTelphone == "" {
				t.Error("Result column 'ShopTelphone' can not be empty")
			}
			if result.SenderAddress == "" {
				t.Error("Result column 'SenderAddress' can not be empty")
			}
			if result.ShopContact == "" {
				t.Error("Result column 'ShopContact' can not be empty")
			}
			if result.ServiceName == "" {
				t.Error("Result column 'ServiceName' can not be empty")
			}

			res := CheckExpressBusinessReturn{
				CommonServerReturn: CommonServerReturn{
					"oABCD", "gh_abcdefg", 1533042556, "event", "check_biz", 1, "success",
				},
				BizID: "xyz",
				Quota: 3.14159265358,
			}
			return &res
		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	xmlData := `<xml>
    <ToUserName><![CDATA[gh_abcdefg]]></ToUserName>
    <FromUserName><![CDATA[oABCD]]></FromUserName>
    <CreateTime>1533042556</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[check_biz]]></Event>
    <BizID><![CDATA[xyz]]></BizID>
    <BizPwd><![CDATA[xyz123]]></BizPwd>
    <ShopAppID><![CDATA[wxABCD]]></ShopAppID>
    <ShopName><![CDATA[商户名称]]></ShopName>
    <ShopTelphone><![CDATA[18677778888]]></ShopTelphone>
    <ShopContact><![CDATA[村正]]></ShopContact>
    <ServiceName><![CDATA[标准快递]]></ServiceName>
    <SenderAddress><![CDATA[广东省广州市海珠区新港中路397号]]></SenderAddress>
</xml>`
	xmlResp, err := http.Post(ts.URL, "application/xml", strings.NewReader(xmlData))
	if err != nil {
		t.Error(err)
	}
	defer xmlResp.Body.Close()
	res := new(CheckExpressBusinessReturn)
	if err := xml.NewDecoder(xmlResp.Body).Decode(res); err != nil {
		t.Error(err)
	}

	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event == "" {
		t.Error("Response column 'Event' can not be empty")
	}
	if res.Quota == 0 {
		t.Error("Response column 'Quota' can not be zero")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}

	jsonData := `{
		"ToUserName": "gh_abcdefg",
		"FromUserName": "oABCD",
		"CreateTime": 1533042556,
		"MsgType": "event",
		"Event": "check_biz",
		"BizID": "xyz",
		"BizPwd": "xyz123",
		"ShopAppID": "wxABCD",
		"ShopName": "商户名称",
		"ShopTelphone": "18677778888",
		"ShopContact": "村正",
		"ServiceName": "标准快递",
		"SenderAddress": "广东省广州市海珠区新港中路397号"
	  }`

	jsonResp, err := http.Post(ts.URL, "application/json", strings.NewReader(jsonData))
	if err != nil {
		t.Error(err)
	}
	defer jsonResp.Body.Close()
	res = new(CheckExpressBusinessReturn)
	if err := json.NewDecoder(jsonResp.Body).Decode(res); err != nil {
		t.Error(err)
	}

	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event == "" {
		t.Error("Response column 'Event' can not be empty")
	}
	if res.Quota == 0 {
		t.Error("Response column 'Quota' can not be zero")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}
}

func TestOnGetExpressQuota(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnGetExpressQuota(func(result *GetExpressQuotaResult) *GetExpressQuotaReturn {
			if result.ToUserName == "" {
				t.Error("ToUserName can not be empty")
			}
			if result.FromUserName == "" {
				t.Error("FromUserName can not be empty")
			}
			if result.CreateTime == 0 {
				t.Error("CreateTime can not be zero")
			}
			if result.MsgType != "event" {
				t.Error("Unexpected message type")
			}
			if result.Event != "get_quota" {
				t.Error("Unexpected message event")
			}

			if result.BizID == "" {
				t.Error("Result column 'BizID' can not be empty")
			}
			if result.BizPwd == "" {
				t.Error("Result column 'BizPwd' can not be empty")
			}
			if result.ShopAppID == "" {
				t.Error("Result column 'ShopAppID' can not be empty")
			}

			res := GetExpressQuotaReturn{
				CommonServerReturn: CommonServerReturn{
					"oABCD", "gh_abcdefg", 1533042556, "event", "get_quota", 1, "success",
				},
				BizID: "xyz",
				Quota: 3.14159265358,
			}
			return &res
		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	xmlData := `<xml>
    <ToUserName><![CDATA[gh_abcdefg]]></ToUserName>
    <FromUserName><![CDATA[oABCD]]></FromUserName>
    <CreateTime>1533042556</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[get_quota]]></Event>
    <BizID><![CDATA[xyz]]></BizID>
    <BizPwd><![CDATA[xyz123]]></BizPwd>
    <ShopAppID><![CDATA[wxABCD]]></ShopAppID>
</xml>`
	xmlResp, err := http.Post(ts.URL, "application/xml", strings.NewReader(xmlData))
	if err != nil {
		t.Error(err)
	}
	defer xmlResp.Body.Close()
	res := new(GetExpressQuotaReturn)
	if err := xml.NewDecoder(xmlResp.Body).Decode(res); err != nil {
		t.Error(err)
	}

	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event != "get_quota" {
		t.Error("Invalid event")
	}
	if res.Quota == 0 {
		t.Error("Response column 'Quota' can not be zero")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}

	jsonData := `{
		"ToUserName": "gh_abcdefg",
		"FromUserName": "oABCD",
		"CreateTime": 1533042556,
		"MsgType": "event",
		"Event": "get_quota",
		"BizID": "xyz",
		"BizPwd": "xyz123",
		"ShopAppID": "wxABCD"
	  }`

	jsonResp, err := http.Post(ts.URL, "application/json", strings.NewReader(jsonData))
	if err != nil {
		t.Error(err)
	}
	defer jsonResp.Body.Close()
	res = new(GetExpressQuotaReturn)
	if err := json.NewDecoder(jsonResp.Body).Decode(res); err != nil {
		t.Error(err)
	}
	if res.ToUserName == "" {
		t.Error("Response column 'ToUserName' can not be empty")
	}
	if res.FromUserName == "" {
		t.Error("Response column 'FromUserName' can not be empty")
	}
	if res.CreateTime == 0 {
		t.Error("Response column 'CreateTime' can not be zero")
	}
	if res.MsgType == "" {
		t.Error("Response column 'MsgType' can not be empty")
	}
	if res.Event != "get_quota" {
		t.Error("Invalid event")
	}
	if res.Quota == 0 {
		t.Error("Response column 'Quota' can not be zero")
	}
	if res.BizID == "" {
		t.Error("Response column 'BizID' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response column 'ResultMsg' can not be empty")
	}
}

func TestPreviewTemplate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/delivery/template/preview" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			WaybillTemplate string `json:"waybill_template"`
			WaybillData     string `json:"waybill_data"`
			Custom          struct {
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
			} `json:"custom"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.WaybillTemplate == "" {
			t.Error("Response column waybill_template can not be empty")
		}
		if params.WaybillData == "" {
			t.Error("Response column waybill_data can not be empty")
		}
		if params.Custom.OrderID == "" {
			t.Error("param custom.order_id can not be empty")
		}
		if params.Custom.DeliveryID == "" {
			t.Error("param custom.delivery_id can not be empty")
		}

		if params.Custom.BizID == "" {
			t.Error("param custom.biz_id can not be empty")
		}

		if params.Custom.Sender.Name == "" {
			t.Error("param custom.sender.name can not be empty")
		}
		if params.Custom.Sender.Province == "" {
			t.Error("param custom.sender.province can not be empty")
		}
		if params.Custom.Sender.City == "" {
			t.Error("param custom.sender.city can not be empty")
		}
		if params.Custom.Sender.Area == "" {
			t.Error("param custom.sender.area can not be empty")
		}
		if params.Custom.Sender.Address == "" {
			t.Error("param custom.sender.address can not be empty")
		}
		if params.Custom.Receiver.Name == "" {
			t.Error("param custom.receiver.name can not be empty")
		}
		if params.Custom.Receiver.Province == "" {
			t.Error("param custom.receiver.province can not be empty")
		}
		if params.Custom.Receiver.City == "" {
			t.Error("param custom.receiver.city can not be empty")
		}
		if params.Custom.Receiver.Area == "" {
			t.Error("param custom.receiver.area can not be empty")
		}
		if params.Custom.Receiver.Address == "" {
			t.Error("param custom.receiver.address can not be empty")
		}

		if params.Custom.Cargo.Count == 0 {
			t.Error("param custom.cargo.count can not be zero")
		}
		if params.Custom.Cargo.Weight == 0 {
			t.Error("param custom.cargo.weight can not be zero")
		}
		if params.Custom.Cargo.SpaceX == 0 {
			t.Error("param custom.cargo.spaceX can not be zero")
		}
		if params.Custom.Cargo.SpaceY == 0 {
			t.Error("param custom.cargo.spaceY can not be zero")
		}
		if params.Custom.Cargo.SpaceZ == 0 {
			t.Error("param custom.cargo.spaceZ can not be zero")
		}
		if len(params.Custom.Cargo.DetailList) == 0 {
			t.Error("param cargo.custom.detailList can not be empty")
		} else {
			if (params.Custom.Cargo.DetailList[0].Name) == "" {
				t.Error("param custom.cargo.detailList.name can not be empty")
			}
			if (params.Custom.Cargo.DetailList[0].Count) == 0 {
				t.Error("param custom.cargo.detailList.count can not be zero")
			}
		}
		if params.Custom.Shop.WXAPath == "" {
			t.Error("param custom.shop.wxa_path can not be empty")
		}
		if params.Custom.Shop.IMGUrl == "" {
			t.Error("param custom.shop.img_url can not be empty")
		}
		if params.Custom.Shop.GoodsName == "" {
			t.Error("param custom.shop.goods_name can not be empty")
		}
		if params.Custom.Shop.GoodsCount == 0 {
			t.Error("param custom.shop.goods_count can not be zero")
		}
		if params.Custom.Insured.Used == 0 {
			t.Error("param custom.insured.use_insured can not be zero")
		}
		if params.Custom.Service.Name == "" {
			t.Error("param custom.service.service_name can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"waybill_id": "1234567890123",
			"rendered_waybill_template": "PGh0bWw+dGVzdDwvaHRtbD4="
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
		"waybill_id": "1234567890123",
		"waybill_data": "##ZTO_mark##11-22-33##ZTO_bagAddr##广州##",
		"waybill_template": "PGh0bWw+dGVzdDwvaHRtbD4=",
		"custom": {
		  "order_id": "012345678901234567890123456789",
		  "openid": "oABC123456",
		  "delivery_id": "ZTO",
		  "biz_id": "xyz",
		  "custom_remark": "易碎物品",
		  "sender": {
			"name": "张三",
			"tel": "18666666666",
			"mobile": "020-88888888",
			"company": "公司名",
			"post_code": "123456",
			"country": "中国",
			"province": "广东省",
			"city": "广州市",
			"area": "海珠区",
			"address": "XX路XX号XX大厦XX栋XX"
		  },
		  "receiver": {
			"name": "王小蒙",
			"tel": "18610000000",
			"mobile": "020-77777777",
			"company": "公司名",
			"post_code": "654321",
			"country": "中国",
			"province": "广东省",
			"city": "广州市",
			"area": "天河区",
			"address": "XX路XX号XX大厦XX栋XX"
		  },
		  "shop": {
			"wxa_path": "/index/index?from=waybill",
			"img_url": "https://mmbiz.qpic.cn/mmbiz_png/KfrZwACMrmwbPGicysN6kibW0ibXwzmA3mtTwgSsdw4Uicabduu2pfbfwdKicQ8n0v91kRAUX6SDESQypl5tlRwHUPA/640",
			"goods_name": "一千零一夜钻石包&爱马仕柏金钻石包",
			"goods_count": 2
		  },
		  "cargo": {
			"count": 2,
			"weight": 5.5,
			"space_x": 30.5,
			"space_y": 20,
			"space_z": 20,
			"detail_list": [
			  {
				"name": "一千零一夜钻石包",
				"count": 1
			  },
			  {
				"name": "爱马仕柏金钻石包",
				"count": 1
			  }
			]
		  },
		  "insured": {
			"use_insured": 1,
			"insured_value": 10000
		  },
		  "service": {
			"service_type": 0,
			"service_name": "标准快递"
		  }
		}
	  }`

	previewer := new(ExpressTemplatePreviewer)
	err := json.Unmarshal([]byte(raw), previewer)
	if err != nil {
		t.Fatal(err)
	}

	_, err = previewer.preview(ts.URL+apiPreviewTemplate, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateBusiness(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/delivery/service/business/update" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ShopAppID  string `json:"shop_app_id"` // 商户的小程序AppID，即审核商户事件中的 ShopAppID
			BizID      string `json:"biz_id"`      // 商户账户
			ResultCode int    `json:"result_code"` // 审核结果，0 表示审核通过，其他表示审核失败
			ResultMsg  string `json:"result_msg"`  // 审核错误原因，仅 result_code 不等于 0 时需要设置
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.ShopAppID == "" {
			t.Error("Response column shop_app_id can not be empty")
		}
		if params.BizID == "" {
			t.Error("Response column biz_id can not be empty")
		}
		if params.ResultCode == 0 {
			t.Error("Response column result_code can not be zero")
		}
		if params.ResultMsg == "" {
			t.Error("Response column result_msg can not be empty")
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

	raw := `{
		"shop_app_id": "wxABCD",
		"biz_id": "xyz",
		"result_code": 1,
		"result_msg": "审核通过"
	  }`

	updater := new(BusinessUpdater)
	err := json.Unmarshal([]byte(raw), updater)
	if err != nil {
		t.Fatal(err)
	}

	_, err = updater.update(ts.URL+apiUpdateBusiness, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdatePath(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/delivery/path/update" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			Token      string `json:"token"`       // 商户侧下单事件中推送的 Token 字段
			WaybillID  string `json:"waybill_id"`  // 运单 ID
			ActionTime uint   `json:"action_time"` // 轨迹变化 Unix 时间戳
			ActionType int    `json:"action_type"` // 轨迹变化类型
			ActionMsg  string `json:"action_msg"`  // 轨迹变化具体信息说明，展示在快递轨迹详情页中。若有手机号码，则直接写11位手机号码。使用UTF-8编码。
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.Token == "" {
			t.Error("Response column token can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("Response column waybill_id can not be empty")
		}
		if params.ActionMsg == "" {
			t.Error("Response column action_msg can not be empty")
		}
		if params.ActionTime == 0 {
			t.Error("Response column action_time can not be zero")
		}
		if params.ActionType == 0 {
			t.Error("Response column action_type can not be zero")
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

	raw := `{
		"token": "TOKEN",
		"waybill_id": "12345678901234567890",
		"action_time": 1533052800,
		"action_type": 300002,
		"action_msg": "丽影邓丽君【18666666666】正在派件"
	  }`

	updater := new(ExpressPathUpdater)
	err := json.Unmarshal([]byte(raw), updater)
	if err != nil {
		t.Fatal(err)
	}

	_, err = updater.update(ts.URL+apiUpdatePath, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}
