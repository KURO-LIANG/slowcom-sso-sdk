package entity

type MpTemplateReq struct {
	TemplateDataMap map[string]TemplateDataItem `json:"templateDataMap" v:"required#消息内容不能为空" description:"消息内容"`
	UserId          uint64                      `json:"userId" description:"通知用户ID，与unionId二选一"`
	PagePath        string                      `json:"pagePath" description:"小程序页面地址"`
	MiniAppId       string                      `json:"miniAppId" description:"小程序APPID"`
	UnionId         string                      `json:"unionId" description:"开放平台unionId，与用户ID二选一"`
	MsgType         int8                        `json:"msgType" description:"消息类型 0-设备报警，1-入住成功，2-预订成功，3-订单来了(通知管家),4-取消预订"`
}
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}
