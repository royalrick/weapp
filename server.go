package weapp

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// MsgType 消息类型
type MsgType = string

// 所有消息类型
const (
	MsgText  MsgType = "text"                      // 文本消息类型
	MsgImg           = "image"                     // 图片消息类型
	MsgCard          = "miniprogrampage"           // 小程序卡片消息类型
	MsgEvent         = "event"                     // 事件类型
	MsgTrans         = "transfer_customer_service" // 转发客服消息
)

// EventType 事件类型
type EventType string

// 所有事件类型
const (
	EventQuotaGet                   EventType = "get_quota"                       // 查询商户余额
	EventCheckBusiness                        = "check_biz"                       // 取消订单事件
	EventMediaCheckAsync                      = "wxa_media_check"                 // 异步校验图片/音频
	EventAddExpressOrder                      = "add_waybill"                     // 请求下单事件
	EventExpressPathUpdate                    = "add_express_path"                // 运单轨迹更新事件
	EventExpressOrderCancel                   = "cancel_waybill"                  // 审核商户事件
	EventUserTempsessionEnter                 = "user_enter_tempsession"          // 用户进入临时会话状态
	EventNearbyPoiAuditInfoAdd                = "add_nearby_poi_audit_info"       // 附近小程序添加地点审核状态通知
	EventDeliveryOrderStatusUpdate            = "update_waybill_status"           // 配送单配送状态更新通知
	EventAgentPosQuery                        = "transport_get_agent_pos"         // 查询骑手当前位置信息
	EventAuthInfoGet                          = "get_auth_info"                   // 使用授权码拉取授权信息
	EventAuthAccountCancel                    = "cancel_auth_account"             // 取消授权帐号
	EventDeliveryOrderAdd                     = "transport_add_order"             // 真实发起下单任务
	EventDeliveryOrderTipsAdd                 = "transport_add_tips"              // 对待接单状态的订单增加小费
	EventDeliveryOrderCancel                  = "transport_cancel_order"          // 取消订单操作
	EventDeliveryOrderReturnConfirm           = "transport_confirm_return_to_biz" // 异常妥投商户收货确认
	EventDeliveryOrderPreAdd                  = "transport_precreate_order"       // 预下单
	EventDeliveryOrderPreCancel               = "transport_precancel_order"       // 预取消订单
	EventDeliveryOrderQuery                   = "transport_query_order_status"    // 查询订单状态
	EventDeliveryOrderReadd                   = "transport_readd_order"           // 下单
	EventPreAuthCodeGet                       = "get_pre_auth_code"               // 获取预授权码
	EventRiderScoreSet                        = "transport_set_rider_score"       // 给骑手评分
)

// Server 微信通知服务处理器
type Server struct {
	appID    string // 小程序 ID
	mchID    string // 商户号
	apiKey   string // 商户签名密钥
	token    string // 微信服务器验证令牌
	aesKey   []byte // base64 解码后的消息加密密钥
	validate bool   // 是否验证请求来自微信服务器

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

type dataType = string

const (
	dataTypeJSON dataType = "application/json"
	dataTypeXML           = "text/xml"
)

// NewServer 返回经过初始化的Server
func NewServer(appID, token, aesKey, mchID, apiKey string, validate bool) (*Server, error) {

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
	}

	return &server, nil
}

func getDataType(req *http.Request) dataType {
	content := req.Header.Get("Content-Type")

	switch {
	case strings.Contains(content, dataTypeJSON):
		return dataTypeJSON
	case strings.Contains(content, dataTypeXML):
		return dataTypeXML
	default:
		return content
	}
}

func unmarshal(data []byte, tp dataType, v interface{}) error {
	switch tp {
	case dataTypeJSON:
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
	case dataTypeXML:
		if err := xml.Unmarshal(data, v); err != nil {
			return err
		}
	default:
		return errors.New("invalid content type: " + tp)
	}

	return nil
}

func marshal(data interface{}, tp dataType) ([]byte, error) {
	switch tp {
	case dataTypeJSON:
		return json.Marshal(data)
	case dataTypeXML:
		return xml.Marshal(data)
	default:
		return nil, errors.New("invalid content type: " + tp)
	}
}

// 处理消息体
func (srv *Server) handleRequest(w http.ResponseWriter, r *http.Request, isEncrpt bool, tp dataType) (interface{}, error) {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if isEncrpt { // 处理加密消息

		query := r.URL.Query()
		nonce, signature, timestamp := query.Get("nonce"), query.Get("signature"), query.Get("timestamp")

		// 检验消息是否来自微信服务器
		if !validateSignature(signature, srv.token, timestamp, nonce) {
			return nil, errors.New("failed to validate signature")
		}

		res := new(EncryptedResult)
		if err := unmarshal(raw, tp, res); err != nil {
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
	if err := unmarshal(raw, tp, res); err != nil {
		return nil, err
	}

	switch res.MsgType {
	case MsgText:
		msg := new(TextMessageResult)
		if err := unmarshal(raw, tp, msg); err != nil {
			return nil, err
		}
		if srv.textMessageHandler != nil {
			return srv.textMessageHandler(msg), nil
		}

	case MsgImg:
		msg := new(ImageMessageResult)
		if err := unmarshal(raw, tp, msg); err != nil {
			return nil, err
		}
		if srv.imageMessageHandler != nil {
			return srv.imageMessageHandler(msg), nil
		}

	case MsgCard:
		msg := new(CardMessageResult)
		if err := unmarshal(raw, tp, msg); err != nil {
			return nil, err
		}
		if srv.cardMessageHandler != nil {
			return srv.cardMessageHandler(msg), nil
		}
	case MsgEvent:
		switch res.Event {
		case EventUserTempsessionEnter:
			msg := new(UserTempsessionEnterResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.userTempsessionEnterHandler != nil {
				srv.userTempsessionEnterHandler(msg)
			}
		case EventQuotaGet:
			msg := new(GetExpressQuotaResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.quotaGetHandler != nil {
				return srv.quotaGetHandler(msg), nil
			}
		case EventMediaCheckAsync:
			msg := new(MediaCheckAsyncResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.mediaCheckAsyncHandler != nil {
				srv.mediaCheckAsyncHandler(msg)
			}
		case EventAddExpressOrder:
			msg := new(AddExpressOrderResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.addExpressOrderHandler != nil {
				return srv.addExpressOrderHandler(msg), nil
			}
		case EventExpressOrderCancel:
			msg := new(CancelExpressOrderResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.expressOrderCancelHandler != nil {
				return srv.expressOrderCancelHandler(msg), nil
			}
		case EventCheckBusiness:
			msg := new(CheckExpressBusinessResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.checkExpressBusinessHandler != nil {
				return srv.checkExpressBusinessHandler(msg), nil
			}
		case EventDeliveryOrderStatusUpdate:
			msg := new(DeliveryOrderStatusUpdateResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderStatusUpdateHandler != nil {
				return srv.deliveryOrderStatusUpdateHandler(msg), nil
			}
		case EventAgentPosQuery:
			msg := new(AgentPosQueryResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.agentPosQueryHandler != nil {
				return srv.agentPosQueryHandler(msg), nil
			}
		case EventAuthInfoGet:
			msg := new(AuthInfoGetResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.authInfoGetHandler != nil {
				return srv.authInfoGetHandler(msg), nil
			}
		case EventAuthAccountCancel:
			msg := new(CancelAuthResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.authAccountCancelHandler != nil {
				return srv.authAccountCancelHandler(msg), nil
			}
		case EventDeliveryOrderAdd:
			msg := new(DeliveryOrderAddResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderAddHandler != nil {
				return srv.deliveryOrderAddHandler(msg), nil
			}
		case EventDeliveryOrderTipsAdd:
			msg := new(DeliveryOrderAddTipsResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderTipsAddHandler != nil {
				return srv.deliveryOrderTipsAddHandler(msg), nil
			}
		case EventDeliveryOrderCancel:
			msg := new(DeliveryOrderCancelResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderCancelHandler != nil {
				return srv.deliveryOrderCancelHandler(msg), nil
			}
		case EventDeliveryOrderReturnConfirm:
			msg := new(DeliveryOrderReturnConfirmResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderReturnConfirmHandler != nil {
				return srv.deliveryOrderReturnConfirmHandler(msg), nil
			}
		case EventDeliveryOrderPreAdd:
			msg := new(DeliveryOrderPreAddResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderPreAddHandler != nil {
				return srv.deliveryOrderPreAddHandler(msg), nil
			}
		case EventDeliveryOrderPreCancel:
			msg := new(DeliveryOrderPreCancelResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderPreCancelHandler != nil {
				return srv.deliveryOrderPreCancelHandler(msg), nil
			}
		case EventDeliveryOrderQuery:
			msg := new(DeliveryOrderQueryResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderQueryHandler != nil {
				return srv.deliveryOrderQueryHandler(msg), nil
			}
		case EventDeliveryOrderReadd:
			msg := new(DeliveryOrderReaddResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.deliveryOrderReaddHandler != nil {
				return srv.deliveryOrderReaddHandler(msg), nil
			}
		case EventPreAuthCodeGet:
			msg := new(PreAuthCodeGetResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.preAuthCodeGetHandler != nil {
				return srv.preAuthCodeGetHandler(msg), nil
			}
		case EventRiderScoreSet:
			msg := new(RiderScoreSetResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.riderScoreSetHandler != nil {
				return srv.riderScoreSetHandler(msg), nil
			}
		case EventExpressPathUpdate:
			msg := new(ExpressPathUpdateResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.expressPathUpdateHandler != nil {
				srv.expressPathUpdateHandler(msg)
			}
		case EventNearbyPoiAuditInfoAdd:
			msg := new(AddNearbyPoiResult)
			if err := unmarshal(raw, tp, msg); err != nil {
				return nil, err
			}
			if srv.addNearbyPoiAuditHandler != nil {
				srv.addNearbyPoiAuditHandler(msg)
			}
		default:
			return nil, fmt.Errorf("unexpected message type '%s'", res.MsgType)
		}
	default:
		return nil, fmt.Errorf("unexpected message type '%s'", res.MsgType)
	}
	return nil, nil
}

// Serve 接收并处理微信通知服务
func (srv *Server) Serve(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		tp := getDataType(r)
		isEncrpt := isEncrypted(r)
		res, err := srv.handleRequest(w, r, isEncrpt, tp)
		if err != nil {
			return fmt.Errorf("handle request content error: %s", err)
		}

		if res != nil {
			raw, err := marshal(res, tp)
			if err != nil {
				return err
			}
			if isEncrpt {
				res, err := srv.encryptMsg(string(raw), time.Now().Unix())
				if err != nil {
					return err
				}
				raw, err = marshal(res, tp)
				if err != nil {
					return err
				}
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", tp)
			if _, err := w.Write(raw); err != nil {
				return err
			}
		} else {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
			w.WriteHeader(http.StatusOK)
			_, err = io.WriteString(w, "success")
			if err != nil {
				return err
			}
		}

		return nil
	case "GET":
		if srv.validate { // 请求来自微信验证成功后原样返回 echostr 参数内容
			if !srv.validateServer(r) {
				return errors.New("验证消息来自微信服务器失败")
			}

			w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
			w.WriteHeader(http.StatusOK)

			echostr := r.URL.Query().Get("echostr")
			_, err := io.WriteString(w, echostr)
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return errors.New("invalid request method: " + r.Method)
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

	return validateSignature(signature, nonce, timestamp, srv.token)
}

// 加密消息
func (srv *Server) encryptMsg(message string, timestamp int64) (*EncryptedMsgRequest, error) {

	key := srv.aesKey

	//获得16位随机字符串，填充到明文之前
	nonce := randomString(16)
	text := nonce + strconv.Itoa(len(message)) + message + srv.appID
	plaintext := pkcs7encode([]byte(text))

	cipher, err := cbcEncrypt(key, plaintext, key)
	if err != nil {
		return nil, err
	}

	encrypt := base64.StdEncoding.EncodeToString(cipher)
	timestr := strconv.FormatInt(timestamp, 10)

	//生成安全签名
	signature := createSignature(srv.token, timestr, nonce, encrypt)

	request := EncryptedMsgRequest{
		Nonce:        nonce,
		Encrypt:      encrypt,
		TimeStamp:    timestr,
		MsgSignature: signature,
	}

	return &request, nil
}

// 检验消息的真实性，并且获取解密后的明文.
func (srv *Server) decryptMsg(encrypted string) ([]byte, error) {

	key := srv.aesKey

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	data, err := cbcDecrypt(key, ciphertext, key)
	if err != nil {
		return nil, err
	}

	return data, nil
}
