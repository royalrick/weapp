package server

// EventType 事件类型
type EventType string

// 所有事件类型
const (
	EventQuotaGet                   EventType = "get_quota"                       // 查询商户余额
	EventCheckBusiness              EventType = "check_biz"                       // 取消订单事件
	EventMediaCheckAsync            EventType = "wxa_media_check"                 // 异步校验图片/音频
	EventAddExpressOrder            EventType = "add_waybill"                     // 请求下单事件
	EventExpressPathUpdate          EventType = "add_express_path"                // 运单轨迹更新事件
	EventExpressOrderCancel         EventType = "cancel_waybill"                  // 审核商户事件
	EventUserTempsessionEnter       EventType = "user_enter_tempsession"          // 用户进入临时会话状态
	EventNearbyPoiAuditInfoAdd      EventType = "add_nearby_poi_audit_info"       // 附近小程序添加地点审核状态通知
	EventDeliveryOrderStatusUpdate  EventType = "update_waybill_status"           // 配送单配送状态更新通知
	EventAgentPosQuery              EventType = "transport_get_agent_pos"         // 查询骑手当前位置信息
	EventAuthInfoGet                EventType = "get_auth_info"                   // 使用授权码拉取授权信息
	EventAuthAccountCancel          EventType = "cancel_auth_account"             // 取消授权帐号
	EventDeliveryOrderAdd           EventType = "transport_add_order"             // 真实发起下单任务
	EventDeliveryOrderTipsAdd       EventType = "transport_add_tips"              // 对待接单状态的订单增加小费
	EventDeliveryOrderCancel        EventType = "transport_cancel_order"          // 取消订单操作
	EventDeliveryOrderReturnConfirm EventType = "transport_confirm_return_to_biz" // 异常妥投商户收货确认
	EventDeliveryOrderPreAdd        EventType = "transport_precreate_order"       // 预下单
	EventDeliveryOrderPreCancel     EventType = "transport_precancel_order"       // 预取消订单
	EventDeliveryOrderQuery         EventType = "transport_query_order_status"    // 查询订单状态
	EventDeliveryOrderReadd         EventType = "transport_readd_order"           // 下单
	EventPreAuthCodeGet             EventType = "get_pre_auth_code"               // 获取预授权码
	EventRiderScoreSet              EventType = "transport_set_rider_score"       // 给骑手评分
	EventSubscribeMsgSentEvent      EventType = "subscribe_msg_sent_event"        // 订阅消息发送结果通知
	EventSubscribeMsgPopup          EventType = "subscribe_msg_popup_event"       // 订阅消息弹框事件
	EventSubscribeMsgChange         EventType = "subscribe_msg_change_event"      // 用户改变订阅消息事件
)
