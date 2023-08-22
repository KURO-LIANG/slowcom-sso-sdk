package entity

type BaseUserInfo struct {
	Id        uint64 `json:"id" description:"id"`         // id
	NickName  string `json:"nickName" description:"用户昵称"` // 用户昵称
	Avatar    string `json:"avatar" description:"用户头像"`   // 用户头像
	Phone     string `json:"phone" description:"手机号"`     // 手机号
	ChannelId uint64 `json:"channelId" description:"渠道ID"`
	MaOpenId  string `json:"maOpenId" description:"小程序openid"` // 小程序openid
	UnionId   string `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
}
type UserInfoReq struct {
	Phone string `json:"phone" description:"手机号"` // id
	Id    uint64 `json:"id" description:"用户ID"`   // 用户ID
}

type BaseUserChannelInfoReq struct {
	Ids []uint64 `json:"ids" description:"用户Ids"` // 用户ID
}
type BaseUserChannelInfo struct {
	UserId     uint64 `json:"userId" description:"用户ID"`   // 用户ID
	NickName   string `json:"nickName" description:"用户昵称"` // 用户昵称
	MpOpenId   string `json:"mpOpenId" description:"用户的公众号openid"`
	AppID      string `json:"appId" description:"用户所在渠道的公众号id"`
	AppSecret  string `json:"appSecret" description:"公众号appSecret"` // 公众号appSecret
	TemplateId string `json:"templateId" description:"用户所在渠道的公众号模板消息ID"`
	MiniAppId  string `json:"miniAppId" description:"用户所在渠道的小程序id"`
}
