package weapp

// Text 接收的文本消息
type Text struct {
	Content string `json:"Content,omitempty" xml:"Content,omitempty"`
}

// Image 接收的图片消息
type Image struct {
	PicURL  string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	MediaID string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

// Card 接收的卡片消息
type Card struct {
	Title        string `json:"Title,omitempty" xml:"Title,omitempty"`               // 标题
	AppID        string `json:"AppId,omitempty" xml:"AppId,omitempty"`               // 小程序 appid
	PagePath     string `json:"PagePath,omitempty" xml:"PagePath,omitempty"`         // 小程序页面路径
	ThumbURL     string `json:"ThumbUrl,omitempty" xml:"ThumbUrl,omitempty"`         // 封面图片的临时cdn链接
	ThumbMediaID string `json:"ThumbMediaId,omitempty" xml:"ThumbMediaId,omitempty"` // 封面图片的临时素材id
}

// Link 图文链接消息
type Link struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ThumbURL    string `json:"thumb_url"`
}

// AsyncMedia 异步校验的图片/音频
type AsyncMedia struct {
	IsRisky       uint8  `json:"isrisky"`         // 检测结果，0：暂未检测到风险，1：风险
	ExtraInfoJSON string `json:"extra_info_json"` // 附加信息，默认为空
	AppID         string `json:"appid"`           // 小程序的appid
	TraceID       string `json:"trace_id"`        // 任务id
	StatusCode    int    `json:"status_code"`     // 默认为：0，4294966288(-1008)为链接无法下载
}
