package weapp

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAbnormalConfirm(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/confirm_return" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ShopID       string `json:"shopid"`        // 商家id， 由配送公司分配的appkey
			ShopOrderID  string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
			ShopNo       string `json:"shop_no"`       // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
			DeliverySign string `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串
			WaybillID    string `json:"waybill_id"`    // 配送单id
			Remark       string `json:"remark"`        // 备注
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.ShopID == "" {
			t.Error("Response column shopid can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Response column shop_order_id can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Response column shop_no can not be empty")
		}

		if params.DeliverySign == "" {
			t.Error("Response column delivery_sign can not be empty")
		}

		if params.WaybillID == "" {
			t.Error("Response column waybill_id can not be empty")
		}
		if params.Remark == "" {
			t.Error("Response column remark can not be empty")
		}
		w.WriteHeader(http.StatusOK)

		raw := `{
			"errcode": 1,
			"errmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
			"shopid": "123456",
			"shop_order_id": "123456",
			"shop_no": "shop_no_111",
			"waybill_id": "123456",
			"remark": "remark",
			"delivery_sign": "123456"
		}`

	confirmer := new(AbnormalConfirmer)
	err := json.Unmarshal([]byte(raw), confirmer)
	if err != nil {
		t.Fatal(err)
	}

	_, err = confirmer.confirm(ts.URL+apiAbnormalConfirm, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/add" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			DeliveryToken string `json:"delivery_token"` // 预下单接口返回的参数，配送公司可保证在一段时间内运费不变
			ShopID        string `json:"shopid"`         // 商家id， 由配送公司分配的appkey
			ShopOrderID   string `json:"shop_order_id"`  // 唯一标识订单的 ID，由商户生成
			ShopNo        string `json:"shop_no"`        // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
			DeliverySign  string `json:"delivery_sign"`  // 用配送公司提供的appSecret加密的校验串
			DeliveryID    string `json:"delivery_id"`    // 配送公司ID
			OpenID        string `json:"openid"`         // 下单用户的openid
			SubBizID      string `json:"sub_biz_id"`     // 子商户id，区分小程序内部多个子商户
			Sender        struct {
				Name           string  `json:"name"`            // 姓名，最长不超过256个字符
				City           string  `json:"city"`            // 城市名称，如广州市
				Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
				AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
				Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
				Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
				Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
				CoordinateType uint8   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
			} `json:"sender"` // 发件人信息，闪送、顺丰同城急送必须填写，美团配送、达达，若传了shop_no的值可不填该字段
			Receiver struct {
				Name           string  `json:"name"`            // 姓名，最长不超过256个字符
				City           string  `json:"city"`            // 城市名称，如广州市
				Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
				AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
				Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
				Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
				Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
				CoordinateType uint8   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
			} `json:"receiver"` // 收件人信息
			Cargo struct {
				GoodsValue  float64 `json:"goods_value"`  // 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-5000]
				GoodsHeight float64 `json:"goods_height"` // 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-45]
				GoodsLength float64 `json:"goods_length"` // 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-65]
				GoodsWidth  float64 `json:"goods_width"`  // 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
				GoodsWeight float64 `json:"goods_weight"` // 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
				GoodsDetail struct {
					Goods []struct {
						Count uint    `json:"good_count"` // 货物数量
						Name  string  `json:"good_name"`  // 货品名称
						Price float32 `json:"good_price"` // 货品单价，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）
						Unit  string  `json:"good_unit"`  // 货品单位，最长不超过20个字符
					} `json:"goods"` // 货物交付信息，最长不超过100个字符
				} `json:"goods_detail"` // 货物详情，最长不超过10240个字符
				GoodsPickupInfo   string `json:"goods_pickup_info"`   // 货物取货信息，用于骑手到店取货，最长不超过100个字符
				GoodsDeliveryInfo string `json:"goods_delivery_info"` // 货物交付信息，最长不超过100个字符
				CargoFirstClass   string `json:"cargo_first_class"`   // 品类一级类目
				CargoSecondClass  string `json:"cargo_second_class"`  // 品类二级类目
			} `json:"cargo"` // 货物信息
			OrderInfo struct {
				DeliveryServiceCode  string  `json:"delivery_service_code"`  // 配送服务代码 不同配送公司自定义,微信侧不理解
				OrderType            uint8   `json:"order_type"`             // 订单类型, 0: 即时单 1 预约单，如预约单，需要设置expected_delivery_time或expected_finish_time或expected_pick_time
				ExpectedDeliveryTime uint    `json:"expected_delivery_time"` // 期望派单时间(顺丰同城急送、达达、支持)，unix-timestamp
				ExpectedFinishTime   uint    `json:"expected_finish_time"`   // 期望送达时间(顺丰同城急送、美团配送支持），unix-timestamp
				ExpectedPickTime     uint    `json:"expected_pick_time"`     // 期望取件时间（闪送支持），unix-timestamp
				PoiSeq               string  `json:"poi_seq"`                // 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
				Note                 string  `json:"note"`                   // 备注，最长不超过200个字符
				OrderTime            uint    `json:"order_time"`             // 用户下单付款时间
				IsInsured            uint8   `json:"is_insured"`             // 是否保价，0，非保价，1.保价
				DeclaredValue        float64 `json:"declared_value"`         // 保价金额，单位为元，精确到分
				Tips                 float64 `json:"tips"`                   // 小费，单位为元, 下单一般不加小费
				IsDirectDelivery     uint    `json:"is_direct_delivery"`     // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑手只能配送此订单至完成，配送费用也相应高一些，闪送必须选1，达达可选0或1，其余配送公司不支持直拿直送）
				CashOnDelivery       uint    `json:"cash_on_delivery"`       // 骑手应付金额，单位为元，精确到分
				CashOnPickup         uint    `json:"cash_on_pickup"`         // 骑手应收金额，单位为元，精确到分
				RiderPickMethod      uint8   `json:"rider_pick_method"`      // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
				IsFinishCodeNeeded   uint8   `json:"is_finish_code_needed"`  // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
				IsPickupCodeNeeded   uint8   `json:"is_pickup_code_needed"`  // 取货码（0：不需要；1：需要。取货码的作用是：骑手必须输入取货码才能从商家取货）
			} `json:"order_info"` // 订单信息
			Shop struct {
				WxaPath    string `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
				ImgURL     string `json:"img_url"`     // 商品缩略图 url
				GoodsName  string `json:"goods_name"`  // 商品名称
				GoodsCount uint   `json:"goods_count"` // 商品数量
			} `json:"shop"` // 商品信息，会展示到物流通知消息中
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.DeliveryToken == "" {
			t.Error("Response column 'delivery_token' can not be empty")
		}
		if params.ShopID == "" {
			t.Error("Response column 'shopid' can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Response column 'shop_order_id' can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Response column 'shop_no' can not be empty")
		}
		if params.DeliverySign == "" {
			t.Error("Response column 'delivery_sign' can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("Response column 'delivery_id' can not be empty")
		}
		if params.OpenID == "" {
			t.Error("Response column 'openid' can not be empty")
		}
		if params.SubBizID == "" {
			t.Error("Response column 'sub_biz_id' can not be empty")
		}

		if params.Sender.Name == "" {
			t.Error("Param 'sender.name' can not be empty")
		}
		if params.Sender.City == "" {
			t.Error("Param 'sender.city' can not be empty")
		}
		if params.Sender.Address == "" {
			t.Error("Param 'sender.address' can not be empty")
		}
		if params.Sender.AddressDetail == "" {
			t.Error("Param 'sender.address_detail' can not be empty")
		}
		if params.Sender.Phone == "" {
			t.Error("Param 'sender.phone' can not be empty")
		}
		if params.Sender.Lng == 0 {
			t.Error("Param 'sender.lng' can not be empty")
		}
		if params.Sender.Lat == 0 {
			t.Error("Param 'sender.lat' can not be empty")
		}
		if params.Sender.CoordinateType == 0 {
			t.Error("Param 'sender.coordinate_type' can not be empty")
		}

		if params.Receiver.Name == "" {
			t.Error("Param 'receiver.name' can not be empty")
		}
		if params.Receiver.City == "" {
			t.Error("Param 'receiver.city' can not be empty")
		}
		if params.Receiver.Address == "" {
			t.Error("Param 'receiver.address' can not be empty")
		}
		if params.Receiver.AddressDetail == "" {
			t.Error("Param 'receiver.address_detail' can not be empty")
		}
		if params.Receiver.Phone == "" {
			t.Error("Param 'receiver.phone' can not be empty")
		}
		if params.Receiver.Lng == 0 {
			t.Error("Param 'receiver.lng' can not be empty")
		}
		if params.Receiver.Lat == 0 {
			t.Error("Param 'receiver.lat' can not be empty")
		}
		if params.Receiver.CoordinateType == 0 {
			t.Error("Param 'receiver.coordinate_type' can not be empty")
		}
		if params.Cargo.GoodsValue == 0 {
			t.Error("Param 'cargo.goods_value' can not be empty")
		}
		if params.Cargo.GoodsHeight == 0 {
			t.Error("Param 'cargo.goods_height' can not be empty")
		}
		if params.Cargo.GoodsLength == 0 {
			t.Error("Param 'cargo.goods_length' can not be empty")
		}
		if params.Cargo.GoodsWidth == 0 {
			t.Error("Param 'cargo.goods_width' can not be empty")
		}
		if params.Cargo.GoodsWeight == 0 {
			t.Error("Param 'cargo.goods_weight' can not be empty")
		}
		if params.Cargo.CargoFirstClass == "" {
			t.Error("Param 'cargo.cargo_first_class' can not be empty")
		}
		if params.Cargo.CargoSecondClass == "" {
			t.Error("Param 'cargo.cargo_second_class' can not be empty")
		}
		if len(params.Cargo.GoodsDetail.Goods) > 0 {
			if params.Cargo.GoodsDetail.Goods[0].Count == 0 {
				t.Error("Param 'cargo.goods_detail.goods.good_count' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Name == "" {
				t.Error("Param 'cargo.goods_detail.goods.good_name' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Price == 0 {
				t.Error("Param 'cargo.goods_detail.goods.good_price' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Unit == "" {
				t.Error("Param 'cargo.goods_detail.goods.good_unit' can not be empty")
			}
		}
		if params.OrderInfo.DeliveryServiceCode == "" {
			t.Error("Param 'order_info.delivery_service_code' can not be empty")
		}
		if params.Shop.WxaPath == "" {
			t.Error("Param 'shop.wxa_path' can not be empty")
		}
		if params.Shop.ImgURL == "" {
			t.Error("Param 'shop.img_url' can not be empty")
		}
		if params.Shop.GoodsName == "" {
			t.Error("Param 'shop.goods_name' can not be empty")
		}
		if params.Shop.GoodsCount == 0 {
			t.Error("Param 'shop.goods_count' can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"fee": 11,
			"deliverfee": 11,
			"couponfee": 1,
			"tips": 1,
			"insurancefee": 1000,
			"insurancfee": 1,
			"distance": 1001,
			"waybill_id": "123456789",
			"order_status": 101,
			"finish_code": 1024,
			"pickup_code": 2048,
			"dispatch_duration": 300
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
		"cargo": {
			"cargo_first_class": "美食宵夜",
			"cargo_second_class": "零食小吃",
			"goods_detail": {
				"goods": [
					{
						"good_count": 1,
						"good_name": "水果",
						"good_price": 11,
						"good_unit": "元"
					},
					{
						"good_count": 2,
						"good_name": "蔬菜",
						"good_price": 21,
						"good_unit": "元"
					}
				]
			},
			"goods_height": 1,
			"goods_length": 3,
			"goods_value": 5,
			"goods_weight": 1,
			"goods_width": 2
		},
		"delivery_id": "SFTC",
		"delivery_sign": "01234567890123456789",
		"openid": "oABC123456",
		"order_info": {
			"delivery_service_code": "xxx",
			"expected_delivery_time": 1,
			"is_direct_delivery": 1,
			"is_finish_code_needed": 1,
			"is_insured": 1,
			"is_pickup_code_needed": 1,
			"note": "test_note",
			"order_time": 1555220757,
			"order_type": 1,
			"poi_seq": "1111",
			"tips": 0
		},
		"receiver": {
			"address": "xxx地铁站",
			"address_detail": "2号楼202",
			"city": "北京市",
			"coordinate_type": 1,
			"lat": 40.1529600001,
			"lng": 116.5060300001,
			"name": "老王",
			"phone": "18512345678"
		},
		"sender": {
			"address": "xx大厦",
			"address_detail": "1号楼101",
			"city": "北京市",
			"coordinate_type": 1,
			"lat": 40.4486120001,
			"lng": 116.3830750001,
			"name": "刘一",
			"phone": "13712345678"
		},
		"shop": {
			"goods_count": 2,
			"goods_name": "宝贝",
			"img_url": "https://mmbiz.qpic.cn/mmbiz_png/xxxxxxxxx/0?wx_fmt=png",
			"wxa_path": "/page/index/index"
		},
		"shop_no": "12345678",
		"sub_biz_id": "sub_biz_id_1",
		"shop_order_id": "SFTC_001",
		"shopid": "122222222",
		"delivery_token": "xxxxxxxx"
	 }`

	creator := new(DeliveryOrderCreator)
	err := json.Unmarshal([]byte(raw), creator)
	if err != nil {
		t.Fatal(err)
	}

	res, err := creator.create(ts.URL+apiAddDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.Fee == 0 {
		t.Error("Response 'fee' can not be empty")
	}
	if res.Deliverfee == 0 {
		t.Error("Response 'deliverfee' can not be empty")
	}
	if res.Couponfee == 0 {
		t.Error("Response 'couponfee' can not be empty")
	}
	if res.Tips == 0 {
		t.Error("Response 'tips' can not be empty")
	}
	if res.Insurancefee == 0 {
		t.Error("Response 'insurancefee' can not be empty")
	}
	if res.Distance == 0 {
		t.Error("Response 'distance' can not be empty")
	}
	if res.WaybillID == "" {
		t.Error("Response 'waybill_id' can not be empty")
	}
	if res.OrderStatus == 0 {
		t.Error("Response 'order_status' can not be empty")
	}
	if res.FinishCode == 0 {
		t.Error("Response 'finish_code' can not be empty")
	}
	if res.PickupCode == 0 {
		t.Error("Response 'pickup_code' can not be empty")
	}
	if res.DispatchDuration == 0 {
		t.Error("Response 'dispatch_duration' can not be empty")
	}
	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
}

func TestAddDeliveryTip(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/addtips" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}

		params := struct {
			ShopID       string  `json:"shopid"`
			ShopOrderID  string  `json:"shop_order_id"`
			ShopNo       string  `json:"shop_no"`
			DeliverySign string  `json:"delivery_sign"`
			WaybillID    string  `json:"waybill_id"`
			OpenID       string  `json:"openid"`
			Tips         float64 `json:"tips"`
			Remark       string  `json:"remark"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ShopID == "" {
			t.Error("Param 'shopid' can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Param 'shop_order_id' can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Param 'shop_no' can not be empty")
		}
		if params.DeliverySign == "" {
			t.Error("Param 'delivery_sign' can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("Param 'waybill_id' can not be empty")
		}
		if params.OpenID == "" {
			t.Error("Param 'openid' can not be empty")
		}
		if params.Tips == 0 {
			t.Error("Param 'tips' can not be empty")
		}
		if params.Remark == "" {
			t.Error("Param 'remark' can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
		"shopid": "123456",
		"shop_order_id": "123456",
		"waybill_id": "123456",
		"tips": 5,
		"openid": "mock-open-id",
		"remark": "gogogo",
		"delivery_sign": "123456",
		"shop_no": "shop_no_111"
	 }`

	adder := new(DeliveryTipAdder)
	err := json.Unmarshal([]byte(raw), adder)
	if err != nil {
		t.Fatal(err)
	}

	res, err := adder.add(ts.URL+apiAddDeliveryTip, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
}

func TestCancelDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/cancel" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}

		params := struct {
			ShopID       string `json:"shopid"`           // 商家id， 由配送公司分配的appkey
			ShopOrderID  string `json:"shop_order_id"`    // 唯一标识订单的 ID，由商户生成
			ShopNo       string `json:"shop_no"`          // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
			DeliverySign string `json:"delivery_sign"`    // 用配送公司提供的appSecret加密的校验串
			DeliveryID   string `json:"delivery_id"`      // 快递公司ID
			WaybillID    string `json:"waybill_id"`       // 配送单id
			ReasonID     uint8  `json:"cancel_reason_id"` // 取消原因Id
			Reason       string `json:"cancel_reason"`    // 取消原因
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ShopID == "" {
			t.Error("Param 'shopid' can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Param 'shop_order_id' can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Param 'shop_no' can not be empty")
		}
		if params.DeliverySign == "" {
			t.Error("Param 'delivery_sign' can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("Param 'waybill_id' can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("Param 'delivery_id' can not be empty")
		}
		if params.ReasonID == 0 {
			t.Error("Param 'cancel_reason_id' can not be empty")
		}
		if params.Reason == "" {
			t.Error("Param 'cancel_reason' can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"deduct_fee": 5,
			"desc": "blabla"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
		"shopid": "123456",
		"shop_order_id": "123456",
		"waybill_id": "123456",
		"delivery_id": "123456",
		"cancel_reason_id": 1,
		"cancel_reason": "mock-cancel-reson",
		"delivery_sign": "123456",
		"shop_no": "shop_no_111"
	 }`

	canceler := new(DeliveryOrderCanceler)
	err := json.Unmarshal([]byte(raw), canceler)
	if err != nil {
		t.Fatal(err)
	}

	res, err := canceler.cancel(ts.URL+apiCancelDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
	if res.DeductFee == 0 {
		t.Error("Response 'deduct_fee' can not be empty")
	}
	if res.Desc == "" {
		t.Error("Response 'desc' can not be empty")
	}
}

func TestGetAllImmediateDelivery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/delivery/getall" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"list": [
			   {
				   "delivery_id": "SFTC",
				   "delivery_name": "顺发同城"
			   },
			   {
				   "delivery_id": "MTPS",
				   "delivery_name": "美团配送"
			   }
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	res, err := getAllImmediateDelivery(ts.URL+apiGetAllImmediateDelivery, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
	if len(res.List) == 0 {
		t.Error("Response 'list' can not be empty")
	} else {
		for _, item := range res.List {
			if item.ID == "" {
				t.Error("Response 'list.delivery_id' can not be empty")
			}
			if item.Name == "" {
				t.Error("Response 'list.delivery_name' can not be empty")
			}
		}
	}
}

func TestGetBindAccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/shop/get" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}
		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"shop_list": [
				{
				 "delivery_id": "SFTC",
				 "shopid": "123456",
				 "audit_result": 1
				},
				{
				 "delivery_id": "MTPS",
				 "shopid": "123456",
				 "audit_result": 1
				}
			]
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	res, err := getBindAccount(ts.URL+apiGetDeliveryBindAccount, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
	if len(res.ShopList) == 0 {
		t.Error("Response 'shop_list' can not be empty")
	} else {
		for _, item := range res.ShopList {
			if item.DeliveryID == "" {
				t.Error("Response 'shop_list.delivery_id' can not be empty")
			}
			if item.ShopID == "" {
				t.Error("Response 'shop_list.shopid' can not be empty")
			}
			if item.AuditResult == 0 {
				t.Error("Response 'audit_result.shopid' can not be empty")
			}
		}
	}
}

func TestGetDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/get" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}

		params := struct {
			ShopID       string `json:"shopid"`
			ShopOrderID  string `json:"shop_order_id"`
			ShopNo       string `json:"shop_no"`
			DeliverySign string `json:"delivery_sign"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.ShopID == "" {
			t.Error("Response column shopid can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Response column shop_order_id can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Response column shop_no can not be empty")
		}
		if params.DeliverySign == "" {
			t.Error("Response column delivery_sign can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"order_status":	1,	
			"waybill_id":	"string",
			"rider_name":	"string",
			"rider_phone":	"string",	
			"rider_lng": 3.14,	
			"rider_lat": 3.14
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	raw := `{
		"shopid": "xxxxxx",
		"shop_order_id": "xxxxxx",
		"shop_no": "xxxxxx",
		"delivery_sign": "xxxxxx"
	 }`

	getter := new(DeliveryOrderGetter)
	err := json.Unmarshal([]byte(raw), getter)
	if err != nil {
		t.Fatal(err)
	}

	res, err := getter.get(ts.URL+apiGetDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
	if res.WaybillID == "" {
		t.Error("Response 'waybill_id' can not be empty")
	}
	if res.OrderStatus == 0 {
		t.Error("Response 'order_status' can not be empty")
	}

	if res.RiderName == "" {
		t.Error("Response 'rider_name' can not be empty")
	}
	if res.RiderPhone == "" {
		t.Error("Response 'rider_phone' can not be empty")
	}
	if res.RiderLng == 0 {
		t.Error("Response 'rider_lng' can not be empty")
	}
	if res.RiderLat == 0 {
		t.Error("Response 'rider_lat' can not be empty")
	}
}

func TestMockUpdateDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/test_update_order" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}

		params := struct {
			ShopID      string `json:"shopid"`
			ShopOrderID string `json:"shop_order_id"`
			ActionTime  uint   `json:"action_time"`
			OrderStatus int    `json:"order_status"`
			ActionMsg   string `json:"action_msg"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.ShopID == "" {
			t.Error("Response column shopid can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Response column shop_order_id can not be empty")
		}
		if params.ActionTime == 0 {
			t.Error("Response column action_time can not be empty")
		}
		if params.OrderStatus == 0 {
			t.Error("Response column order_status can not be empty")
		}
		if params.ActionMsg == "" {
			t.Error("Response column action_msg can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	raw := `{
		"shopid": "test_shop_id",
		"shop_order_id": "xxxxxxxxxxx",
		"waybill_id": "xxxxxxxxxxxxx",
		"action_time": 12345678,
		"order_status": 101,
		"action_msg": "xxxxxx"
	 }`

	mocker := new(UpdateDeliveryOrderMocker)
	err := json.Unmarshal([]byte(raw), mocker)
	if err != nil {
		t.Fatal(err)
	}

	res, err := mocker.mock(ts.URL+apiMockUpdateDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
}

func TestPreAddDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/pre_add" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			ShopID       string `json:"shopid"`        // 商家id， 由配送公司分配的appkey
			ShopOrderID  string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
			ShopNo       string `json:"shop_no"`       // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
			DeliverySign string `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串
			DeliveryID   string `json:"delivery_id"`   // 配送公司ID
			OpenID       string `json:"openid"`        // 下单用户的openid
			SubBizID     string `json:"sub_biz_id"`    // 子商户id，区分小程序内部多个子商户
			Sender       struct {
				Name           string  `json:"name"`            // 姓名，最长不超过256个字符
				City           string  `json:"city"`            // 城市名称，如广州市
				Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
				AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
				Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
				Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
				Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
				CoordinateType uint8   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
			} `json:"sender"` // 发件人信息，闪送、顺丰同城急送必须填写，美团配送、达达，若传了shop_no的值可不填该字段
			Receiver struct {
				Name           string  `json:"name"`            // 姓名，最长不超过256个字符
				City           string  `json:"city"`            // 城市名称，如广州市
				Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
				AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
				Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
				Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
				Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
				CoordinateType uint8   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
			} `json:"receiver"` // 收件人信息
			Cargo struct {
				GoodsValue  float64 `json:"goods_value"`  // 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-5000]
				GoodsHeight float64 `json:"goods_height"` // 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-45]
				GoodsLength float64 `json:"goods_length"` // 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-65]
				GoodsWidth  float64 `json:"goods_width"`  // 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
				GoodsWeight float64 `json:"goods_weight"` // 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
				GoodsDetail struct {
					Goods []struct {
						Count uint    `json:"good_count"` // 货物数量
						Name  string  `json:"good_name"`  // 货品名称
						Price float32 `json:"good_price"` // 货品单价，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）
						Unit  string  `json:"good_unit"`  // 货品单位，最长不超过20个字符
					} `json:"goods"` // 货物交付信息，最长不超过100个字符
				} `json:"goods_detail"` // 货物详情，最长不超过10240个字符
				GoodsPickupInfo   string `json:"goods_pickup_info"`   // 货物取货信息，用于骑手到店取货，最长不超过100个字符
				GoodsDeliveryInfo string `json:"goods_delivery_info"` // 货物交付信息，最长不超过100个字符
				CargoFirstClass   string `json:"cargo_first_class"`   // 品类一级类目
				CargoSecondClass  string `json:"cargo_second_class"`  // 品类二级类目
			} `json:"cargo"` // 货物信息
			OrderInfo struct {
				DeliveryServiceCode  string  `json:"delivery_service_code"`  // 配送服务代码 不同配送公司自定义,微信侧不理解
				OrderType            uint8   `json:"order_type"`             // 订单类型, 0: 即时单 1 预约单，如预约单，需要设置expected_delivery_time或expected_finish_time或expected_pick_time
				ExpectedDeliveryTime uint    `json:"expected_delivery_time"` // 期望派单时间(顺丰同城急送、达达、支持)，unix-timestamp
				ExpectedFinishTime   uint    `json:"expected_finish_time"`   // 期望送达时间(顺丰同城急送、美团配送支持），unix-timestamp
				ExpectedPickTime     uint    `json:"expected_pick_time"`     // 期望取件时间（闪送支持），unix-timestamp
				PoiSeq               string  `json:"poi_seq"`                // 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
				Note                 string  `json:"note"`                   // 备注，最长不超过200个字符
				OrderTime            uint    `json:"order_time"`             // 用户下单付款时间
				IsInsured            uint8   `json:"is_insured"`             // 是否保价，0，非保价，1.保价
				DeclaredValue        float64 `json:"declared_value"`         // 保价金额，单位为元，精确到分
				Tips                 float64 `json:"tips"`                   // 小费，单位为元, 下单一般不加小费
				IsDirectDelivery     uint    `json:"is_direct_delivery"`     // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑手只能配送此订单至完成，配送费用也相应高一些，闪送必须选1，达达可选0或1，其余配送公司不支持直拿直送）
				CashOnDelivery       uint    `json:"cash_on_delivery"`       // 骑手应付金额，单位为元，精确到分
				CashOnPickup         uint    `json:"cash_on_pickup"`         // 骑手应收金额，单位为元，精确到分
				RiderPickMethod      uint8   `json:"rider_pick_method"`      // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
				IsFinishCodeNeeded   uint8   `json:"is_finish_code_needed"`  // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
				IsPickupCodeNeeded   uint8   `json:"is_pickup_code_needed"`  // 取货码（0：不需要；1：需要。取货码的作用是：骑手必须输入取货码才能从商家取货）
			} `json:"order_info"` // 订单信息
			Shop struct {
				WxaPath    string `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
				ImgURL     string `json:"img_url"`     // 商品缩略图 url
				GoodsName  string `json:"goods_name"`  // 商品名称
				GoodsCount uint   `json:"goods_count"` // 商品数量
			} `json:"shop"` // 商品信息，会展示到物流通知消息中
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ShopID == "" {
			t.Error("Response column 'shopid' can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Response column 'shop_order_id' can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Response column 'shop_no' can not be empty")
		}
		if params.DeliverySign == "" {
			t.Error("Response column 'delivery_sign' can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("Response column 'delivery_id' can not be empty")
		}
		if params.OpenID == "" {
			t.Error("Response column 'openid' can not be empty")
		}
		if params.SubBizID == "" {
			t.Error("Response column 'sub_biz_id' can not be empty")
		}

		if params.Sender.Name == "" {
			t.Error("Param 'sender.name' can not be empty")
		}
		if params.Sender.City == "" {
			t.Error("Param 'sender.city' can not be empty")
		}
		if params.Sender.Address == "" {
			t.Error("Param 'sender.address' can not be empty")
		}
		if params.Sender.AddressDetail == "" {
			t.Error("Param 'sender.address_detail' can not be empty")
		}
		if params.Sender.Phone == "" {
			t.Error("Param 'sender.phone' can not be empty")
		}
		if params.Sender.Lng == 0 {
			t.Error("Param 'sender.lng' can not be empty")
		}
		if params.Sender.Lat == 0 {
			t.Error("Param 'sender.lat' can not be empty")
		}
		if params.Sender.CoordinateType == 0 {
			t.Error("Param 'sender.coordinate_type' can not be empty")
		}

		if params.Receiver.Name == "" {
			t.Error("Param 'receiver.name' can not be empty")
		}
		if params.Receiver.City == "" {
			t.Error("Param 'receiver.city' can not be empty")
		}
		if params.Receiver.Address == "" {
			t.Error("Param 'receiver.address' can not be empty")
		}
		if params.Receiver.AddressDetail == "" {
			t.Error("Param 'receiver.address_detail' can not be empty")
		}
		if params.Receiver.Phone == "" {
			t.Error("Param 'receiver.phone' can not be empty")
		}
		if params.Receiver.Lng == 0 {
			t.Error("Param 'receiver.lng' can not be empty")
		}
		if params.Receiver.Lat == 0 {
			t.Error("Param 'receiver.lat' can not be empty")
		}
		if params.Receiver.CoordinateType == 0 {
			t.Error("Param 'receiver.coordinate_type' can not be empty")
		}
		if params.Cargo.GoodsValue == 0 {
			t.Error("Param 'cargo.goods_value' can not be empty")
		}
		if params.Cargo.GoodsHeight == 0 {
			t.Error("Param 'cargo.goods_height' can not be empty")
		}
		if params.Cargo.GoodsLength == 0 {
			t.Error("Param 'cargo.goods_length' can not be empty")
		}
		if params.Cargo.GoodsWidth == 0 {
			t.Error("Param 'cargo.goods_width' can not be empty")
		}
		if params.Cargo.GoodsWeight == 0 {
			t.Error("Param 'cargo.goods_weight' can not be empty")
		}
		if params.Cargo.CargoFirstClass == "" {
			t.Error("Param 'cargo.cargo_first_class' can not be empty")
		}
		if params.Cargo.CargoSecondClass == "" {
			t.Error("Param 'cargo.cargo_second_class' can not be empty")
		}
		if len(params.Cargo.GoodsDetail.Goods) > 0 {
			if params.Cargo.GoodsDetail.Goods[0].Count == 0 {
				t.Error("Param 'cargo.goods_detail.goods.good_count' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Name == "" {
				t.Error("Param 'cargo.goods_detail.goods.good_name' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Price == 0 {
				t.Error("Param 'cargo.goods_detail.goods.good_price' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Unit == "" {
				t.Error("Param 'cargo.goods_detail.goods.good_unit' can not be empty")
			}
		}
		if params.OrderInfo.DeliveryServiceCode == "" {
			t.Error("Param 'order_info.delivery_service_code' can not be empty")
		}
		if params.Shop.WxaPath == "" {
			t.Error("Param 'shop.wxa_path' can not be empty")
		}
		if params.Shop.ImgURL == "" {
			t.Error("Param 'shop.img_url' can not be empty")
		}
		if params.Shop.GoodsName == "" {
			t.Error("Param 'shop.goods_name' can not be empty")
		}
		if params.Shop.GoodsCount == 0 {
			t.Error("Param 'shop.goods_count' can not be empty")
		}

		w.WriteHeader(http.StatusOK)
		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"fee": 11,
			"deliverfee": 11,
			"couponfee": 1,
			"insurancefee": 1000,
			"tips": 1,
			"insurancfee": 1,
			"distance": 1001,
			"dispatch_duration": 301,
			"delivery_token": "1111111"
			}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
		"cargo": {
			"cargo_first_class": "美食宵夜",
			"cargo_second_class": "零食小吃",
			"goods_detail": {
				"goods": [
					{
						"good_count": 1,
						"good_name": "水果",
						"good_price": 11,
						"good_unit": "元"
					},
					{
						"good_count": 2,
						"good_name": "蔬菜",
						"good_price": 21,
						"good_unit": "元"
					}
				]
			},
			"goods_height": 1,
			"goods_length": 3,
			"goods_value": 5,
			"goods_weight": 1,
			"goods_width": 2
		},
		"delivery_id": "SFTC",
		"delivery_sign": "01234567890123456789",
		"openid": "oABC123456",
		"order_info": {
			"delivery_service_code": "xxx",
			"expected_delivery_time": 1,
			"is_direct_delivery": 1,
			"is_finish_code_needed": 1,
			"is_insured": 1,
			"is_pickup_code_needed": 1,
			"note": "test_note",
			"order_time": 1555220757,
			"order_type": 1,
			"poi_seq": "1111",
			"tips": 0
		},
		"receiver": {
			"address": "xxx地铁站",
			"address_detail": "2号楼202",
			"city": "北京市",
			"coordinate_type": 1,
			"lat": 40.1529600001,
			"lng": 116.5060300001,
			"name": "老王",
			"phone": "18512345678"
		},
		"sender": {
			"address": "xx大厦",
			"address_detail": "1号楼101",
			"city": "北京市",
			"coordinate_type": 1,
			"lat": 40.4486120001,
			"lng": 116.3830750001,
			"name": "刘一",
			"phone": "13712345678"
		},
		"shop": {
			"goods_count": 2,
			"goods_name": "宝贝",
			"img_url": "https://mmbiz.qpic.cn/mmbiz_png/xxxxxxxxx/0?wx_fmt=png",
			"wxa_path": "/page/index/index"
		},
		"shop_no": "12345678",
		"sub_biz_id": "sub_biz_id_1",
		"shop_order_id": "SFTC_001",
		"shopid": "122222222"
	 }`

	creator := new(DeliveryOrderCreator)
	err := json.Unmarshal([]byte(raw), creator)
	if err != nil {
		t.Fatal(err)
	}

	res, err := creator.prepare(ts.URL+apiPreAddDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.Fee == 0 {
		t.Error("Response 'fee' can not be empty")
	}
	if res.Deliverfee == 0 {
		t.Error("Response 'deliverfee' can not be empty")
	}
	if res.Couponfee == 0 {
		t.Error("Response 'couponfee' can not be empty")
	}
	if res.Tips == 0 {
		t.Error("Response 'tips' can not be empty")
	}
	if res.Insurancefee == 0 {
		t.Error("Response 'insurancefee' can not be empty")
	}
	if res.Distance == 0 {
		t.Error("Response 'distance' can not be empty")
	}
	if res.DispatchDuration == 0 {
		t.Error("Response 'dispatch_duration' can not be empty")
	}
}

func TestPreCancelDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/precancel" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}

		params := struct {
			ShopID       string `json:"shopid"`           // 商家id， 由配送公司分配的appkey
			ShopOrderID  string `json:"shop_order_id"`    // 唯一标识订单的 ID，由商户生成
			ShopNo       string `json:"shop_no"`          // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
			DeliverySign string `json:"delivery_sign"`    // 用配送公司提供的appSecret加密的校验串
			DeliveryID   string `json:"delivery_id"`      // 快递公司ID
			WaybillID    string `json:"waybill_id"`       // 配送单id
			ReasonID     uint8  `json:"cancel_reason_id"` // 取消原因Id
			Reason       string `json:"cancel_reason"`    // 取消原因
		}{}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.ShopID == "" {
			t.Error("Param 'shopid' can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Param 'shop_order_id' can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Param 'shop_no' can not be empty")
		}
		if params.DeliverySign == "" {
			t.Error("Param 'delivery_sign' can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("Param 'waybill_id' can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("Param 'delivery_id' can not be empty")
		}
		if params.ReasonID == 0 {
			t.Error("Param 'cancel_reason_id' can not be empty")
		}
		if params.Reason == "" {
			t.Error("Param 'cancel_reason' can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"deduct_fee": 5,
			"desc": "blabla"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
		"shopid": "123456",
		"shop_order_id": "123456",
		"waybill_id": "123456",
		"delivery_id": "123456",
		"cancel_reason_id": 1,
		"cancel_reason": "xxxxxx",
		"delivery_sign": "123456",
		"shop_no": "shop_no_111"
	 }`

	canceler := new(DeliveryOrderCanceler)
	err := json.Unmarshal([]byte(raw), canceler)
	if err != nil {
		t.Fatal(err)
	}

	res, err := canceler.prepare(ts.URL+apiPreCancelDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
	if res.DeductFee == 0 {
		t.Error("Response 'deduct_fee' can not be empty")
	}
	if res.Desc == "" {
		t.Error("Response 'desc' can not be empty")
	}
}

func TestReAddDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/business/order/readd" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("access_token can not be empty")
		}

		params := struct {
			DeliveryToken string `json:"delivery_token"` // 预下单接口返回的参数，配送公司可保证在一段时间内运费不变
			ShopID        string `json:"shopid"`         // 商家id， 由配送公司分配的appkey
			ShopOrderID   string `json:"shop_order_id"`  // 唯一标识订单的 ID，由商户生成
			ShopNo        string `json:"shop_no"`        // 商家门店编号， 在配送公司登记，如果只有一个门店，可以不填
			DeliverySign  string `json:"delivery_sign"`  // 用配送公司提供的appSecret加密的校验串
			DeliveryID    string `json:"delivery_id"`    // 配送公司ID
			OpenID        string `json:"openid"`         // 下单用户的openid
			SubBizID      string `json:"sub_biz_id"`     // 子商户id，区分小程序内部多个子商户
			Sender        struct {
				Name           string  `json:"name"`            // 姓名，最长不超过256个字符
				City           string  `json:"city"`            // 城市名称，如广州市
				Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
				AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
				Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
				Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
				Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
				CoordinateType uint8   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
			} `json:"sender"` // 发件人信息，闪送、顺丰同城急送必须填写，美团配送、达达，若传了shop_no的值可不填该字段
			Receiver struct {
				Name           string  `json:"name"`            // 姓名，最长不超过256个字符
				City           string  `json:"city"`            // 城市名称，如广州市
				Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
				AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
				Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
				Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，确到小数点后6位
				Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
				CoordinateType uint8   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标
			} `json:"receiver"` // 收件人信息
			Cargo struct {
				GoodsValue  float64 `json:"goods_value"`  // 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-5000]
				GoodsHeight float64 `json:"goods_height"` // 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-45]
				GoodsLength float64 `json:"goods_length"` // 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-65]
				GoodsWidth  float64 `json:"goods_width"`  // 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
				GoodsWeight float64 `json:"goods_weight"` // 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为(0-50]
				GoodsDetail struct {
					Goods []struct {
						Count uint    `json:"good_count"` // 货物数量
						Name  string  `json:"good_name"`  // 货品名称
						Price float32 `json:"good_price"` // 货品单价，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）
						Unit  string  `json:"good_unit"`  // 货品单位，最长不超过20个字符
					} `json:"goods"` // 货物交付信息，最长不超过100个字符
				} `json:"goods_detail"` // 货物详情，最长不超过10240个字符
				GoodsPickupInfo   string `json:"goods_pickup_info"`   // 货物取货信息，用于骑手到店取货，最长不超过100个字符
				GoodsDeliveryInfo string `json:"goods_delivery_info"` // 货物交付信息，最长不超过100个字符
				CargoFirstClass   string `json:"cargo_first_class"`   // 品类一级类目
				CargoSecondClass  string `json:"cargo_second_class"`  // 品类二级类目
			} `json:"cargo"` // 货物信息
			OrderInfo struct {
				DeliveryServiceCode  string  `json:"delivery_service_code"`  // 配送服务代码 不同配送公司自定义,微信侧不理解
				OrderType            uint8   `json:"order_type"`             // 订单类型, 0: 即时单 1 预约单，如预约单，需要设置expected_delivery_time或expected_finish_time或expected_pick_time
				ExpectedDeliveryTime uint    `json:"expected_delivery_time"` // 期望派单时间(顺丰同城急送、达达、支持)，unix-timestamp
				ExpectedFinishTime   uint    `json:"expected_finish_time"`   // 期望送达时间(顺丰同城急送、美团配送支持），unix-timestamp
				ExpectedPickTime     uint    `json:"expected_pick_time"`     // 期望取件时间（闪送支持），unix-timestamp
				PoiSeq               string  `json:"poi_seq"`                // 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
				Note                 string  `json:"note"`                   // 备注，最长不超过200个字符
				OrderTime            uint    `json:"order_time"`             // 用户下单付款时间
				IsInsured            uint8   `json:"is_insured"`             // 是否保价，0，非保价，1.保价
				DeclaredValue        float64 `json:"declared_value"`         // 保价金额，单位为元，精确到分
				Tips                 float64 `json:"tips"`                   // 小费，单位为元, 下单一般不加小费
				IsDirectDelivery     uint    `json:"is_direct_delivery"`     // 是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑手只能配送此订单至完成，配送费用也相应高一些，闪送必须选1，达达可选0或1，其余配送公司不支持直拿直送）
				CashOnDelivery       uint    `json:"cash_on_delivery"`       // 骑手应付金额，单位为元，精确到分
				CashOnPickup         uint    `json:"cash_on_pickup"`         // 骑手应收金额，单位为元，精确到分
				RiderPickMethod      uint8   `json:"rider_pick_method"`      // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
				IsFinishCodeNeeded   uint8   `json:"is_finish_code_needed"`  // 收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
				IsPickupCodeNeeded   uint8   `json:"is_pickup_code_needed"`  // 取货码（0：不需要；1：需要。取货码的作用是：骑手必须输入取货码才能从商家取货）
			} `json:"order_info"` // 订单信息
			Shop struct {
				WxaPath    string `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
				ImgURL     string `json:"img_url"`     // 商品缩略图 url
				GoodsName  string `json:"goods_name"`  // 商品名称
				GoodsCount uint   `json:"goods_count"` // 商品数量
			} `json:"shop"` // 商品信息，会展示到物流通知消息中
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}

		if params.DeliveryToken == "" {
			t.Error("Response column 'delivery_token' can not be empty")
		}
		if params.ShopID == "" {
			t.Error("Response column 'shopid' can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Response column 'shop_order_id' can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Response column 'shop_no' can not be empty")
		}
		if params.DeliverySign == "" {
			t.Error("Response column 'delivery_sign' can not be empty")
		}
		if params.DeliveryID == "" {
			t.Error("Response column 'delivery_id' can not be empty")
		}
		if params.OpenID == "" {
			t.Error("Response column 'openid' can not be empty")
		}
		if params.SubBizID == "" {
			t.Error("Response column 'sub_biz_id' can not be empty")
		}

		if params.Sender.Name == "" {
			t.Error("Param 'sender.name' can not be empty")
		}
		if params.Sender.City == "" {
			t.Error("Param 'sender.city' can not be empty")
		}
		if params.Sender.Address == "" {
			t.Error("Param 'sender.address' can not be empty")
		}
		if params.Sender.AddressDetail == "" {
			t.Error("Param 'sender.address_detail' can not be empty")
		}
		if params.Sender.Phone == "" {
			t.Error("Param 'sender.phone' can not be empty")
		}
		if params.Sender.Lng == 0 {
			t.Error("Param 'sender.lng' can not be empty")
		}
		if params.Sender.Lat == 0 {
			t.Error("Param 'sender.lat' can not be empty")
		}
		if params.Sender.CoordinateType == 0 {
			t.Error("Param 'sender.coordinate_type' can not be empty")
		}

		if params.Receiver.Name == "" {
			t.Error("Param 'receiver.name' can not be empty")
		}
		if params.Receiver.City == "" {
			t.Error("Param 'receiver.city' can not be empty")
		}
		if params.Receiver.Address == "" {
			t.Error("Param 'receiver.address' can not be empty")
		}
		if params.Receiver.AddressDetail == "" {
			t.Error("Param 'receiver.address_detail' can not be empty")
		}
		if params.Receiver.Phone == "" {
			t.Error("Param 'receiver.phone' can not be empty")
		}
		if params.Receiver.Lng == 0 {
			t.Error("Param 'receiver.lng' can not be empty")
		}
		if params.Receiver.Lat == 0 {
			t.Error("Param 'receiver.lat' can not be empty")
		}
		if params.Receiver.CoordinateType == 0 {
			t.Error("Param 'receiver.coordinate_type' can not be empty")
		}
		if params.Cargo.GoodsValue == 0 {
			t.Error("Param 'cargo.goods_value' can not be empty")
		}
		if params.Cargo.GoodsHeight == 0 {
			t.Error("Param 'cargo.goods_height' can not be empty")
		}
		if params.Cargo.GoodsLength == 0 {
			t.Error("Param 'cargo.goods_length' can not be empty")
		}
		if params.Cargo.GoodsWidth == 0 {
			t.Error("Param 'cargo.goods_width' can not be empty")
		}
		if params.Cargo.GoodsWeight == 0 {
			t.Error("Param 'cargo.goods_weight' can not be empty")
		}
		if params.Cargo.CargoFirstClass == "" {
			t.Error("Param 'cargo.cargo_first_class' can not be empty")
		}
		if params.Cargo.CargoSecondClass == "" {
			t.Error("Param 'cargo.cargo_second_class' can not be empty")
		}
		if len(params.Cargo.GoodsDetail.Goods) > 0 {
			if params.Cargo.GoodsDetail.Goods[0].Count == 0 {
				t.Error("Param 'cargo.goods_detail.goods.good_count' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Name == "" {
				t.Error("Param 'cargo.goods_detail.goods.good_name' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Price == 0 {
				t.Error("Param 'cargo.goods_detail.goods.good_price' can not be empty")
			}
			if params.Cargo.GoodsDetail.Goods[0].Unit == "" {
				t.Error("Param 'cargo.goods_detail.goods.good_unit' can not be empty")
			}
		}
		if params.OrderInfo.DeliveryServiceCode == "" {
			t.Error("Param 'order_info.delivery_service_code' can not be empty")
		}
		if params.Shop.WxaPath == "" {
			t.Error("Param 'shop.wxa_path' can not be empty")
		}
		if params.Shop.ImgURL == "" {
			t.Error("Param 'shop.img_url' can not be empty")
		}
		if params.Shop.GoodsName == "" {
			t.Error("Param 'shop.goods_name' can not be empty")
		}
		if params.Shop.GoodsCount == 0 {
			t.Error("Param 'shop.goods_count' can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok",
			"fee": 11,
			"deliverfee": 11,
			"couponfee": 1,
			"tips": 1,
			"insurancefee": 1000,
			"insurancfee": 1,
			"distance": 1001,
			"waybill_id": "123456789",
			"order_status": 101,
			"finish_code": 1024,
			"pickup_code": 2048,
			"dispatch_duration": 300
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	raw := `{
		"cargo": {
			"cargo_first_class": "美食宵夜",
			"cargo_second_class": "零食小吃",
			"goods_detail": {
				"goods": [
					{
						"good_count": 1,
						"good_name": "水果",
						"good_price": 11,
						"good_unit": "元"
					},
					{
						"good_count": 2,
						"good_name": "蔬菜",
						"good_price": 21,
						"good_unit": "元"
					}
				]
			},
			"goods_height": 1,
			"goods_length": 3,
			"goods_value": 5,
			"goods_weight": 1,
			"goods_width": 2
		},
		"delivery_id": "SFTC",
		"delivery_sign": "01234567890123456789",
		"openid": "oABC123456",
		"order_info": {
			"delivery_service_code": "xxx",
			"expected_delivery_time": 1,
			"is_direct_delivery": 1,
			"is_finish_code_needed": 1,
			"is_insured": 1,
			"is_pickup_code_needed": 1,
			"note": "test_note",
			"order_time": 1555220757,
			"order_type": 1,
			"poi_seq": "1111",
			"tips": 0
		},
		"receiver": {
			"address": "xxx地铁站",
			"address_detail": "2号楼202",
			"city": "北京市",
			"coordinate_type": 1,
			"lat": 40.1529600001,
			"lng": 116.5060300001,
			"name": "老王",
			"phone": "18512345678"
		},
		"sender": {
			"address": "xx大厦",
			"address_detail": "1号楼101",
			"city": "北京市",
			"coordinate_type": 1,
			"lat": 40.4486120001,
			"lng": 116.3830750001,
			"name": "刘一",
			"phone": "13712345678"
		},
		"shop": {
			"goods_count": 2,
			"goods_name": "宝贝",
			"img_url": "https://mmbiz.qpic.cn/mmbiz_png/xxxxxxxxx/0?wx_fmt=png",
			"wxa_path": "/page/index/index"
		},
		"shop_no": "12345678",
		"sub_biz_id": "sub_biz_id_1",
		"shop_order_id": "SFTC_001",
		"shopid": "122222222",
		"delivery_token": "xxxxxxxx"
	 }`

	creator := new(DeliveryOrderCreator)
	err := json.Unmarshal([]byte(raw), creator)
	if err != nil {
		t.Fatal(err)
	}

	res, err := creator.recreate(ts.URL+apiReAddDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}

	if res.Fee == 0 {
		t.Error("Response 'fee' can not be empty")
	}
	if res.Deliverfee == 0 {
		t.Error("Response 'deliverfee' can not be empty")
	}
	if res.Couponfee == 0 {
		t.Error("Response 'couponfee' can not be empty")
	}
	if res.Tips == 0 {
		t.Error("Response 'tips' can not be empty")
	}
	if res.Insurancefee == 0 {
		t.Error("Response 'insurancefee' can not be empty")
	}
	if res.Distance == 0 {
		t.Error("Response 'distance' can not be empty")
	}
	if res.WaybillID == "" {
		t.Error("Response 'waybill_id' can not be empty")
	}
	if res.OrderStatus == 0 {
		t.Error("Response 'order_status' can not be empty")
	}
	if res.FinishCode == 0 {
		t.Error("Response 'finish_code' can not be empty")
	}
	if res.PickupCode == 0 {
		t.Error("Response 'pickup_code' can not be empty")
	}
	if res.DispatchDuration == 0 {
		t.Error("Response 'dispatch_duration' can not be empty")
	}
	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
}

func TestUpdateDeliveryOrder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Error("UnExpect request method")
		}

		if r.URL.EscapedPath() != "/cgi-bin/express/local/delivery/update_order" {
			t.Error("Unexpected path")
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		if r.Form.Get("access_token") == "" {
			t.Fatalf("Query 'access_token' can not be empty")
		}

		params := struct {
			WxToken     string `json:"wx_token"`
			ShopID      string `json:"shopid"`
			ShopOrderID string `json:"shop_order_id"`
			ShopNo      string `json:"shop_no"`
			WaybillID   string `json:"waybill_id"`
			ActionTime  uint   `json:"action_time"`
			OrderStatus int    `json:"order_status"`
			ActionMsg   string `json:"action_msg"`
			WxaPath     string `json:"wxa_path"`
			Agent       struct {
				Name             string `json:"name"`
				Phone            string `json:"phone"`
				IsPhoneEncrypted uint8  `json:"is_phone_encrypted"`
			} `json:"agent"`
			ExpectedDeliveryTime uint `json:"expected_delivery_time"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatal(err)
		}
		if params.WxToken == "" {
			t.Error("Response column wx_token can not be empty")
		}
		if params.Agent.Name == "" {
			t.Error("Response column agent.name can not be empty")
		}
		if params.Agent.Phone == "" {
			t.Error("Response column agent.phone can not be empty")
		}
		if params.Agent.IsPhoneEncrypted == 0 {
			t.Error("Response column agent.is_phone_encrypted can not be empty")
		}
		if params.ShopID == "" {
			t.Error("Response column shopid can not be empty")
		}
		if params.ShopNo == "" {
			t.Error("Response column shop_no can not be empty")
		}
		if params.WaybillID == "" {
			t.Error("Response column waybill_id can not be empty")
		}
		if params.ShopOrderID == "" {
			t.Error("Response column expected_delivery_time can not be empty")
		}
		if params.ExpectedDeliveryTime == 0 {
			t.Error("Response column action_time can not be empty")
		}
		if params.ActionTime == 0 {
			t.Error("Response column action_time can not be empty")
		}
		if params.OrderStatus == 0 {
			t.Error("Response column order_status can not be empty")
		}
		if params.ActionMsg == "" {
			t.Error("Response column action_msg can not be empty")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{
			"resultcode": 1,
			"resultmsg": "ok"
		  }`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	raw := `{
		"wx_token": "xxxxxx",
		"shopid": "test_shop_id",
		"shop_no": "test_shop_id",
		"shop_order_id": "xxxxxxxxxxx",
		"waybill_id": "xxxxxxxxxxxxx",
		"action_time": 12345678,
		"order_status": 101,
		"action_msg": "xxxxxx",
		"wxa_path": "xxxxxx",
		"expected_delivery_time": 123456,
		"agent": {
			"name": "xxxxxx",
			"phone": "xxxxxx",
			"is_phone_encrypted": 1
		}
	 }`

	updater := new(DeliveryOrderUpdater)
	err := json.Unmarshal([]byte(raw), updater)
	if err != nil {
		t.Fatal(err)
	}

	res, err := updater.update(ts.URL+apiUpdateDeliveryOrder, "mock-access-token")
	if err != nil {
		t.Fatal(err)
	}
	if res.ResultCode == 0 {
		t.Error("Response 'resultcode' can not be empty")
	}
	if res.ResultMsg == "" {
		t.Error("Response 'resultmsg' can not be empty")
	}
}

func TestOnDeliveryOrderStatusUpdate(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		aesKey := base64.StdEncoding.EncodeToString([]byte("mock-aes-key"))
		srv, err := NewServer("mock-app-id", "mock-access-token", aesKey, "mock-mch-id", "mock-api-key", false)
		if err != nil {
			t.Fatal(err)
		}

		srv.OnDeliveryOrderStatusUpdate(func(mix *DeliveryOrderStatusUpdateResult) *DeliveryOrderStatusUpdateReturn {
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

			if mix.Event != "update_waybill_status" {
				t.Error("Unexpected message event")
			}

			if mix.ShopID == "" {
				t.Error("Result 'shopid' can not be zero")
			}

			if mix.ShopOrderID == "" {
				t.Error("Result 'shop_order_id' can not be zero")
			}

			if mix.ShopNo == "" {
				t.Error("Result 'shop_no' can not be zero")
			}

			if mix.WaybillID == "" {
				t.Error("Result 'waybill_id' can not be zero")
			}

			if mix.ActionTime == 0 {
				t.Error("Result 'action_time' can not be zero")
			}

			if mix.OrderStatus == 0 {
				t.Error("Result 'order_status' can not be zero")
			}

			if mix.ActionMsg == "" {
				t.Error("Result 'action_msg' can not be zero")
			}

			if mix.Agent.Name == "" {
				t.Error("Result 'agent.name' can not be zero")
			}

			if mix.Agent.Phone == "" {
				t.Error("Result 'agent.phone' can not be zero")
			}

			return &DeliveryOrderStatusUpdateReturn{
				"mock-to-user-name",
				"mock-from-user-name",
				20060102150405,
				"mock-message-type",
				"mock-event",
				0,
				"mock-result-message",
			}
		})

		if err := srv.Serve(w, r); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	jsonData := `{
		"ToUserName": "toUser",
		"FromUserName": "fromUser",
		"CreateTime": 1546924844,
		"MsgType": "event",
		"Event": "update_waybill_status",
		"shopid": "123456",
		"shop_order_id": "123456",
		"waybill_id": "123456",
		"action_time": 1546924844,
		"order_status": 102,
		"action_msg": "xxx",
		"shop_no": "123456",
		"agent": {
		   "name": "xxx",
		   "phone": "1234567"
		}
	  }`

	res, err := http.Post(ts.URL, "application/json", strings.NewReader(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
}
