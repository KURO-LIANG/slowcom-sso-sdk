package entity

type MpTemplateReq struct {
	TemplateDataMap map[string]TemplateDataItem `json:"templateDataMap" v:"required#消息内容不能为空" description:"消息内容"`
	UserId          uint64                      `json:"userId" v:"required#通知用户Id不能为空" description:"通知用户ID"`
	PagePath        string                      `json:"pagePath" description:"小程序页面地址"`
}
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}
