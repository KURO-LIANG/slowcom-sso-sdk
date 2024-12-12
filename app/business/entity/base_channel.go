package entity

type ChannelInfo struct {
	Id            uint64 `json:"id" description:"渠道ID"`                    // 渠道ID
	ChannelName   string `json:"channelName" description:"渠道名称"`           // 渠道名称
	MiniAppId     string `json:"miniAppId" description:"小程序appid"`         // 小程序appid
	MiniAppSecret string `json:"miniAppSecret" description:"小程序appSecret"` // 小程序appSecret
	AppId         string `json:"appId" description:"公众号appid"`             // 公众号appid
	AppSecret     string `json:"appSecret" description:"公众号appSecret"`     // 公众号appSecret
	TemplateId    string `json:"templateId" description:"公众号模板消息ID"`       // 公众号模板消息ID
}

type ChannelList struct {
	Id          uint64 `json:"id" description:"渠道ID"`          // 渠道ID
	ChannelName string `json:"channelName" description:"渠道名称"` // 渠道名称
}
