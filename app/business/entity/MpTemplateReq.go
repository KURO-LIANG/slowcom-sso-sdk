package entity

type MpTemplateReq struct {
	TemplateDataMap map[string]TemplateDataItem `json:"templateDataMap" v:"required#消息内容不能为空" description:"消息内容"`
	UserId          uint64                      `json:"userId" description:"通知用户ID，与unionId二选一"`
	PagePath        string                      `json:"pagePath" description:"小程序页面地址"`
	MiniAppId       string                      `json:"miniAppId" description:"小程序APPID"`
	UnionId         string                      `json:"unionId" description:"开放平台unionId，与用户ID二选一"`
	TemplateId      string                      `json:"templateId" v:"required#模板ID不能为空" description:"公众号模板消息ID"`
}
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}
