package server

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/medivhzhan/weapp/v3/encrypt"
	"github.com/medivhzhan/weapp/v3/request"
)

// Server 微信通知服务处理器
type Server struct {
	appID    string // 小程序 ID
	mchID    string // 商户号
	apiKey   string // 商户签名密钥
	token    string // 微信服务器验证令牌
	aesKey   []byte // base64 解码后的消息加密密钥
	validate bool   // 是否验证请求来自微信服务器

	// 默认处理器
	// 没有自定义处理器的情况下调用该处理器
	handler func(map[string]interface{}) map[string]interface{}

	// 自定义处理器
	textMessageHandler                func(*TextMessageResult) *TransferCustomerMessage
	imageMessageHandler               func(*ImageMessageResult) *TransferCustomerMessage
	cardMessageHandler                func(*CardMessageResult) *TransferCustomerMessage
	userTempsessionEnterHandler       func(*UserTempsessionEnterResult)
	mediaCheckAsyncHandler            func(*MediaCheckAsyncResult)
	expressPathUpdateHandler          func(*ExpressPathUpdateResult)
	addNearbyPoiAuditHandler          func(*AddNearbyPoiResult)
	addExpressOrderHandler            func(*AddExpressOrderResult) *AddExpressOrderReturn
	expressOrderCancelHandler         func(*CancelExpressOrderResult) *CancelExpressOrderReturn
	checkExpressBusinessHandler       func(*CheckExpressBusinessResult) *CheckExpressBusinessReturn
	quotaGetHandler                   func(*GetExpressQuotaResult) *GetExpressQuotaReturn
	deliveryOrderStatusUpdateHandler  func(*DeliveryOrderStatusUpdateResult) *DeliveryOrderStatusUpdateReturn
	agentPosQueryHandler              func(*AgentPosQueryResult) *AgentPosQueryReturn
	authInfoGetHandler                func(*AuthInfoGetResult) *AuthInfoGetReturn
	authAccountCancelHandler          func(*CancelAuthResult) *CancelAuthReturn
	deliveryOrderAddHandler           func(*DeliveryOrderAddResult) *DeliveryOrderAddReturn
	deliveryOrderTipsAddHandler       func(*DeliveryOrderAddTipsResult) *DeliveryOrderAddTipsReturn
	deliveryOrderCancelHandler        func(*DeliveryOrderCancelResult) *DeliveryOrderCancelReturn
	deliveryOrderReturnConfirmHandler func(*DeliveryOrderReturnConfirmResult) *DeliveryOrderReturnConfirmReturn
	deliveryOrderPreAddHandler        func(*DeliveryOrderPreAddResult) *DeliveryOrderPreAddReturn
	deliveryOrderPreCancelHandler     func(*DeliveryOrderPreCancelResult) *DeliveryOrderPreCancelReturn
	deliveryOrderQueryHandler         func(*DeliveryOrderQueryResult) *DeliveryOrderQueryReturn
	deliveryOrderReaddHandler         func(*DeliveryOrderReaddResult) *DeliveryOrderReaddReturn
	preAuthCodeGetHandler             func(*PreAuthCodeGetResult) *PreAuthCodeGetReturn
	riderScoreSetHandler              func(*RiderScoreSetResult) *RiderScoreSetReturn
	subscribeMsgPopupHandler          func(*SubscribeMsgPopupEvent)
	subscribeMsgSentHandler           func(*SubscribeMsgSentEvent)
	subscribeMsgChangeHandler         func(*SubscribeMsgChangeEvent)
}

// OnCustomerServiceTextMessage add handler to handle customer text service message.
func (srv *Server) OnCustomerServiceTextMessage(fn func(*TextMessageResult) *TransferCustomerMessage) {
	srv.textMessageHandler = fn
}

// OnCustomerServiceImageMessage add handler to handle customer image service message.
func (srv *Server) OnCustomerServiceImageMessage(fn func(*ImageMessageResult) *TransferCustomerMessage) {
	srv.imageMessageHandler = fn
}

// OnCustomerServiceCardMessage add handler to handle customer card service message.
func (srv *Server) OnCustomerServiceCardMessage(fn func(*CardMessageResult) *TransferCustomerMessage) {
	srv.cardMessageHandler = fn
}

// OnUserTempsessionEnter add handler to handle customer service message.
func (srv *Server) OnUserTempsessionEnter(fn func(*UserTempsessionEnterResult)) {
	srv.userTempsessionEnterHandler = fn
}

// OnMediaCheckAsync add handler to handle MediaCheckAsync.
func (srv *Server) OnMediaCheckAsync(fn func(*MediaCheckAsyncResult)) {
	srv.mediaCheckAsyncHandler = fn
}

// OnExpressPathUpdate add handler to handle ExpressPathUpdate.
func (srv *Server) OnExpressPathUpdate(fn func(*ExpressPathUpdateResult)) {
	srv.expressPathUpdateHandler = fn
}

// OnAddNearbyPoi add handler to handle AddNearbyPoiAudit.
func (srv *Server) OnAddNearbyPoi(fn func(*AddNearbyPoiResult)) {
	srv.addNearbyPoiAuditHandler = fn
}

// OnAddExpressOrder add handler to handle AddExpressOrder.
func (srv *Server) OnAddExpressOrder(fn func(*AddExpressOrderResult) *AddExpressOrderReturn) {
	srv.addExpressOrderHandler = fn
}

// OnCheckExpressBusiness add handler to handle CheckBusiness.
func (srv *Server) OnCheckExpressBusiness(fn func(*CheckExpressBusinessResult) *CheckExpressBusinessReturn) {
	srv.checkExpressBusinessHandler = fn
}

// OnCancelExpressOrder add handler to handle ExpressOrderCancel.
func (srv *Server) OnCancelExpressOrder(fn func(*CancelExpressOrderResult) *CancelExpressOrderReturn) {
	srv.expressOrderCancelHandler = fn
}

// OnGetExpressQuota add handler to handle QuotaGet.
func (srv *Server) OnGetExpressQuota(fn func(*GetExpressQuotaResult) *GetExpressQuotaReturn) {
	srv.quotaGetHandler = fn
}

// OnDeliveryOrderStatusUpdate add handler to handle DeliveryOrderStatusUpdate.
// OnDeliveryOrderStatusUpdate add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderStatusUpdate(fn func(*DeliveryOrderStatusUpdateResult) *DeliveryOrderStatusUpdateReturn) {
	srv.deliveryOrderStatusUpdateHandler = fn
}

// OnAgentPosQuery add handler to handle AgentPosQuery.
func (srv *Server) OnAgentPosQuery(fn func(*AgentPosQueryResult) *AgentPosQueryReturn) {
	srv.agentPosQueryHandler = fn
}

// OnAuthInfoGet add handler to handle AuthInfoGet.
func (srv *Server) OnAuthInfoGet(fn func(*AuthInfoGetResult) *AuthInfoGetReturn) {
	srv.authInfoGetHandler = fn
}

// OnCancelAuth add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnCancelAuth(fn func(*CancelAuthResult) *CancelAuthReturn) {
	srv.authAccountCancelHandler = fn
}

// OnDeliveryOrderAdd add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderAdd(fn func(*DeliveryOrderAddResult) *DeliveryOrderAddReturn) {
	srv.deliveryOrderAddHandler = fn
}

// OnDeliveryOrderAddTips add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderAddTips(fn func(*DeliveryOrderAddTipsResult) *DeliveryOrderAddTipsReturn) {
	srv.deliveryOrderTipsAddHandler = fn
}

// OnDeliveryOrderCancel add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderCancel(fn func(*DeliveryOrderCancelResult) *DeliveryOrderCancelReturn) {
	srv.deliveryOrderCancelHandler = fn
}

// OnDeliveryOrderReturnConfirm add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderReturnConfirm(fn func(*DeliveryOrderReturnConfirmResult) *DeliveryOrderReturnConfirmReturn) {
	srv.deliveryOrderReturnConfirmHandler = fn
}

// OnDeliveryOrderPreAdd add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderPreAdd(fn func(*DeliveryOrderPreAddResult) *DeliveryOrderPreAddReturn) {
	srv.deliveryOrderPreAddHandler = fn
}

// OnDeliveryOrderPreCancel add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderPreCancel(fn func(*DeliveryOrderPreCancelResult) *DeliveryOrderPreCancelReturn) {
	srv.deliveryOrderPreCancelHandler = fn
}

// OnDeliveryOrderQuery add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderQuery(fn func(*DeliveryOrderQueryResult) *DeliveryOrderQueryReturn) {
	srv.deliveryOrderQueryHandler = fn
}

// OnDeliveryOrderReadd add handler to handle deliveryOrderStatusUpdate.
func (srv *Server) OnDeliveryOrderReadd(fn func(*DeliveryOrderReaddResult) *DeliveryOrderReaddReturn) {
	srv.deliveryOrderReaddHandler = fn
}

// OnPreAuthCodeGet add handler to handle preAuthCodeGet.
func (srv *Server) OnPreAuthCodeGet(fn func(*PreAuthCodeGetResult) *PreAuthCodeGetReturn) {
	srv.preAuthCodeGetHandler = fn
}

// OnRiderScoreSet add handler to handle riderScoreSet.
func (srv *Server) OnRiderScoreSet(fn func(*RiderScoreSetResult) *RiderScoreSetReturn) {
	srv.riderScoreSetHandler = fn
}

// 当用户触发订阅消息弹框后
func (srv *Server) OnSubscribeMsgPopup(fn func(*SubscribeMsgPopupEvent)) {
	srv.subscribeMsgPopupHandler = fn
}

// 当订阅消息发送结束后触发
func (srv *Server) OnSubscribeMsgSent(fn func(*SubscribeMsgSentEvent)) {
	srv.subscribeMsgSentHandler = fn
}

// 当用户通过设置界面改变订阅消息事件内容
func (srv *Server) OnSubscribeMsgChange(fn func(*SubscribeMsgChangeEvent)) {
	srv.subscribeMsgChangeHandler = fn
}

// NewServer 返回经过初始化的Server
func NewServer(appID, token, aesKey, mchID, apiKey string, validate bool, handler func(map[string]interface{}) map[string]interface{}) (*Server, error) {

	key, err := base64.RawStdEncoding.DecodeString(aesKey)
	if err != nil {
		return nil, err
	}

	server := Server{
		appID:    appID,
		mchID:    mchID,
		apiKey:   apiKey,
		token:    token,
		aesKey:   key,
		validate: validate,
		handler:  handler,
	}

	return &server, nil
}

func contentType(req *http.Request) request.ContentType {
	ctp := req.Header.Get("Content-Type")

	switch {
	case strings.Contains(ctp, request.ContentTypeJSON.String()):
		return request.ContentTypeJSON
	case strings.Contains(ctp, request.ContentTypeXML.String()):
		return request.ContentTypeXML
	default:
		return request.ContentTypePlain
	}
}

func unmarshal(data []byte, ctp request.ContentType, v interface{}) error {
	switch ctp {
	case request.ContentTypeJSON:
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
	case request.ContentTypeXML:
		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}
	default:
		return errors.New("invalid content type")
	}

	return nil
}

func marshal(data interface{}, ctp request.ContentType) ([]byte, error) {
	switch ctp {
	case request.ContentTypeJSON:
		return json.Marshal(data)
	case request.ContentTypeXML:
		return xml.Marshal(data)
	default:
		return nil, errors.New("invalid content type")
	}
}

// 处理消息体
func (srv *Server) handleRequest(w http.ResponseWriter, r *http.Request, isEncrpt bool, ctp request.ContentType) (interface{}, error) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if isEncrpt { // 处理加密消息

		query := r.URL.Query()
		nonce, signature, timestamp := query.Get("nonce"), query.Get("signature"), query.Get("timestamp")

		// 检验消息是否来自微信服务器
		if !encrypt.NewSignable(true, srv.token, timestamp, nonce).IsEqual(signature) {
			return nil, errors.New("failed to validate signature")
		}

		res := new(EncryptedResult)
		if err := unmarshal(raw, ctp, res); err != nil {
			return nil, err
		}

		body, err := srv.decryptMsg(res.Encrypt)
		if err != nil {
			return nil, err
		}
		length := binary.BigEndian.Uint32(body[16:20])
		raw = body[20 : 20+length]
	}

	res := new(CommonServerResult)
	if err := unmarshal(raw, ctp, res); err != nil {
		return nil, err
	}

	switch res.MsgType {
	case MsgText:
		msg := new(TextMessageResult)
		if err := unmarshal(raw, ctp, msg); err != nil {
			return nil, err
		}
		if srv.textMessageHandler != nil {
			return srv.textMessageHandler(msg), nil
		}

	case MsgImg:
		msg := new(ImageMessageResult)
		if err := unmarshal(raw, ctp, msg); err != nil {
			return nil, err
		}
		if srv.imageMessageHandler != nil {
			return srv.imageMessageHandler(msg), nil
		}

	case MsgCard:
		msg := new(CardMessageResult)
		if err := unmarshal(raw, ctp, msg); err != nil {
			return nil, err
		}
		if srv.cardMessageHandler != nil {
			return srv.cardMessageHandler(msg), nil
		}
	case MsgEvent:
		switch res.Event {
		case EventUserTempsessionEnter:
			msg := new(UserTempsessionEnterResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.userTempsessionEnterHandler != nil {
				srv.userTempsessionEnterHandler(msg)
			}

		case EventQuotaGet:
			msg := new(GetExpressQuotaResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.quotaGetHandler != nil {
				return srv.quotaGetHandler(msg), nil
			}

		case EventMediaCheckAsync:
			msg := new(MediaCheckAsyncResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.mediaCheckAsyncHandler != nil {
				srv.mediaCheckAsyncHandler(msg)
			}

		case EventAddExpressOrder:
			msg := new(AddExpressOrderResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.addExpressOrderHandler != nil {
				return srv.addExpressOrderHandler(msg), nil
			}

		case EventExpressOrderCancel:
			msg := new(CancelExpressOrderResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.expressOrderCancelHandler != nil {
				return srv.expressOrderCancelHandler(msg), nil
			}

		case EventCheckBusiness:
			msg := new(CheckExpressBusinessResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.checkExpressBusinessHandler != nil {
				return srv.checkExpressBusinessHandler(msg), nil
			}

		case EventDeliveryOrderStatusUpdate:
			msg := new(DeliveryOrderStatusUpdateResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderStatusUpdateHandler != nil {
				return srv.deliveryOrderStatusUpdateHandler(msg), nil
			}

		case EventAgentPosQuery:
			msg := new(AgentPosQueryResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.agentPosQueryHandler != nil {
				return srv.agentPosQueryHandler(msg), nil
			}

		case EventAuthInfoGet:
			msg := new(AuthInfoGetResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.authInfoGetHandler != nil {
				return srv.authInfoGetHandler(msg), nil
			}

		case EventAuthAccountCancel:
			msg := new(CancelAuthResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.authAccountCancelHandler != nil {
				return srv.authAccountCancelHandler(msg), nil
			}

		case EventDeliveryOrderAdd:
			msg := new(DeliveryOrderAddResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderAddHandler != nil {
				return srv.deliveryOrderAddHandler(msg), nil
			}

		case EventDeliveryOrderTipsAdd:
			msg := new(DeliveryOrderAddTipsResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderTipsAddHandler != nil {
				return srv.deliveryOrderTipsAddHandler(msg), nil
			}

		case EventDeliveryOrderCancel:
			msg := new(DeliveryOrderCancelResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderCancelHandler != nil {
				return srv.deliveryOrderCancelHandler(msg), nil
			}

		case EventDeliveryOrderReturnConfirm:
			msg := new(DeliveryOrderReturnConfirmResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderReturnConfirmHandler != nil {
				return srv.deliveryOrderReturnConfirmHandler(msg), nil
			}

		case EventDeliveryOrderPreAdd:
			msg := new(DeliveryOrderPreAddResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderPreAddHandler != nil {
				return srv.deliveryOrderPreAddHandler(msg), nil
			}

		case EventDeliveryOrderPreCancel:
			msg := new(DeliveryOrderPreCancelResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderPreCancelHandler != nil {
				return srv.deliveryOrderPreCancelHandler(msg), nil
			}

		case EventDeliveryOrderQuery:
			msg := new(DeliveryOrderQueryResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderQueryHandler != nil {
				return srv.deliveryOrderQueryHandler(msg), nil
			}

		case EventDeliveryOrderReadd:
			msg := new(DeliveryOrderReaddResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderReaddHandler != nil {
				return srv.deliveryOrderReaddHandler(msg), nil
			}

		case EventPreAuthCodeGet:
			msg := new(PreAuthCodeGetResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.preAuthCodeGetHandler != nil {
				return srv.preAuthCodeGetHandler(msg), nil
			}

		case EventRiderScoreSet:
			msg := new(RiderScoreSetResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.riderScoreSetHandler != nil {
				return srv.riderScoreSetHandler(msg), nil
			}

		case EventExpressPathUpdate:
			msg := new(ExpressPathUpdateResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.expressPathUpdateHandler != nil {
				srv.expressPathUpdateHandler(msg)
			}

		case EventNearbyPoiAuditInfoAdd:
			msg := new(AddNearbyPoiResult)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.addNearbyPoiAuditHandler != nil {
				srv.addNearbyPoiAuditHandler(msg)
			}

		case EventSubscribeMsgPopup:
			msg := new(SubscribeMsgPopupEvent)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.subscribeMsgPopupHandler != nil {
				srv.subscribeMsgPopupHandler(msg)
			}

		case EventSubscribeMsgSentEvent:
			msg := new(SubscribeMsgSentEvent)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.subscribeMsgPopupHandler != nil {
				srv.subscribeMsgSentHandler(msg)
			}

		case EventSubscribeMsgChange:
			msg := new(SubscribeMsgChangeEvent)
			if err := unmarshal(raw, ctp, msg); err != nil {
				return nil, err
			}
			if srv.subscribeMsgChangeHandler != nil {
				srv.subscribeMsgChangeHandler(msg)
			}

		default:
			msg := make(map[string]interface{})
			if err := unmarshal(raw, ctp, &msg); err != nil {
				return nil, err
			}
			if srv.handler != nil {
				return srv.handler(msg), nil
			}
		}

	default:
		msg := make(map[string]interface{})
		if err := unmarshal(raw, ctp, &msg); err != nil {
			return nil, err
		}
		if srv.handler != nil {
			return srv.handler(msg), nil
		}
	}

	return nil, nil
}

// 判断 interface{} 是否为空
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	return reflect.ValueOf(i).IsZero()
}

// Serve 接收并处理微信通知服务
func (srv *Server) Serve(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		ctp := contentType(r)
		isEncrypted := isEncrypted(r)
		res, err := srv.handleRequest(w, r, isEncrypted, ctp)
		if err != nil {
			return fmt.Errorf("handle request content error: %s", err)
		}

		var raw []byte
		if !isNil(res) {
			raw, err = marshal(res, ctp)
			if err != nil {
				return err
			}
		} else {
			raw = []byte("success")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", ctp.String())
		if _, err := w.Write(raw); err != nil {
			return err
		}

		return nil
	case "GET":
		if srv.validate { // 请求来自微信验证成功后原样返回 echostr 参数内容
			if !srv.validateServer(r) {
				return errors.New("验证消息来自微信服务器失败")
			}

			raw := []byte(r.URL.Query().Get("echostr"))

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", request.ContentTypePlain.String())
			if _, err := w.Write(raw); err != nil {
				return err
			}
		}

		return nil
	default:
		return errors.New("invalid request method")
	}
}

// 判断消息是否加密
func isEncrypted(req *http.Request) bool {
	return req.URL.Query().Get("encrypt_type") == "aes"
}

// 验证消息的确来自微信服务器
// 1.将token、timestamp、nonce三个参数进行字典序排序
// 2.将三个参数字符串拼接成一个字符串进行sha1加密
// 3.开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
func (srv *Server) validateServer(req *http.Request) bool {
	query := req.URL.Query()
	nonce := query.Get("nonce")
	signature := query.Get("signature")
	timestamp := query.Get("timestamp")

	return encrypt.NewSignable(true, nonce, timestamp, srv.token).IsEqual(signature)
}

// 加密消息
func (srv *Server) encryptMsg(message string, timestamp int64) (*EncryptedMsgRequest, error) {

	key := make([]byte, len(srv.aesKey))
	copy(key, srv.aesKey)

	//获得16位随机字符串，填充到明文之前
	nonce := randomString(16)
	text := nonce + strconv.Itoa(len(message)) + message + srv.appID

	data, err := encrypt.NewCBC(key, key, []byte(text)).Encrypt()
	if err != nil {
		return nil, err
	}

	cipher := base64.StdEncoding.EncodeToString(data)
	timestr := strconv.FormatInt(timestamp, 10)

	//生成安全签名
	signature := encrypt.NewSignable(true, srv.token, timestr, nonce, cipher).Sign()
	request := EncryptedMsgRequest{
		Nonce:        nonce,
		Encrypt:      cipher,
		TimeStamp:    timestr,
		MsgSignature: signature,
	}

	return &request, nil
}

// 生成随机字符串
//
// @ln 需要生成字符串长度
func randomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

// 检验消息的真实性，并且获取解密后的明文.
func (srv *Server) decryptMsg(encrypted string) ([]byte, error) {

	//不能直接赋值, 底层slice会被修改
	//key := srv.aesKey
	key := make([]byte, len(srv.aesKey))
	copy(key, srv.aesKey)

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	data, err := encrypt.NewCBC(key, key, ciphertext).Decrypt()
	if err != nil {
		return nil, err
	}

	return data, nil
}
