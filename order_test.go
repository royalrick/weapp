package weapp

import (
	"github.com/medivhzhan/weapp/v3/order"
	"testing"
	"time"
)

var appid = "wx0417444aae7355f7"
var accessToken = "74_7-AVnmRJZz5MpDoeaygGfuSZL5TbSjJQmJJ0lXSJhuH0z0IHCGNe3-Uw_VE42xCiJAVgj29fidOITvQvXwu1luKhIm0fbXZvqTWzBGVo98vYJff-DjSDtocqrFy8dYUoCV8sJCTzPtiKJV20CZFeADDGDJ"

func getClient() *Client {
	tokenGetter := func(appid, secret string) (token string, expireIn uint) {
		return accessToken, 10
	}

	sdk := NewClient(
		appid,
		"",
		WithAccessTokenSetter(tokenGetter),
	)

	return sdk
}

func TestClient_NewOrderIsTradeManaged(t *testing.T) {
	orderServe := getClient().NewOrder()
	isTradeManaged, err := orderServe.IsTradeManaged(&order.IsTradeManagedRequest{
		Appid: "wx0417444aae7355f7",
	})
	if err != nil {
		t.Errorf("isTradeManaged err: %+v", err)
		return
	}

	t.Logf("isTradeManaged: %#v", isTradeManaged)
}

func TestClient_NewOrderUploadShippingInfo(t *testing.T) {
	orderServe := getClient().NewOrder()
	isTradeManaged, err := orderServe.UploadShippingInfo(&order.UploadShippingInfoRequest{
		OrderKey: order.OrderKey{
			OrderNumberType: 2,
			TransactionId:   "4200002027202310272131934449",
			Mchid:           "",
			OutTradeNo:      "",
		},
		DeliveryMode:   "UNIFIED_DELIVERY",
		LogisticsType:  4,
		IsAllDelivered: false,
		ShippingList: []order.ShippingList{
			{
				ItemDesc: "这是一个商品",
			},
		},
		UploadTime: time.Now().Format(time.RFC3339),
		Payer: order.Payer{
			Openid: "omlPt4v2t9G40JnX4uXjlA9vsfK0",
		},
	})
	if err != nil {
		t.Errorf("UploadShippingInfo err: %+v", err)
		return
	}

	t.Logf("UploadShippingInfo: %#v", isTradeManaged)
}

func TestClient_NewOrderGetOrder(t *testing.T) {
	orderServe := getClient().NewOrder()
	resp, err := orderServe.GetOrder(&order.GetOrderRequest{
		TransactionId: "4200002027202310272131934449",
	})
	if err != nil {
		t.Errorf("GetOrder err: %+v", err)
		return
	}

	t.Logf("GetOrder: %#v", resp.Order)
}

func TestClient_NewOrderGetOrderList(t *testing.T) {
	orderServe := getClient().NewOrder()
	resp, err := orderServe.GetOrderList(&order.GetOrderListRequest{})
	if err != nil {
		t.Errorf("GetOrder err: %+v", err)
		return
	}

	for _, orderStruct := range resp.OrderList {
		t.Logf("GetOrder: %#v", orderStruct)
	}
}
