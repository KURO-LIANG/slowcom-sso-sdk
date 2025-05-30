package entity

type BaseUserInfo struct {
	UserId   uint64 `json:"userId" description:"用户ID"`   // 用户ID
	NickName string `json:"nickName" description:"用户昵称"` // 用户昵称
	Avatar   string `json:"avatar" description:"用户头像"`   // 用户头像
	Phone    string `json:"phone" description:"手机号"`     // 手机号
	Email    string `json:"email" description:"邮箱"`
	MpOpenId string `json:"mpOpenId" description:"用户的公众号openid"`
	MaOpenId string `json:"maOpenId" description:"小程序openid"` // 小程序openid
	UnionId  string `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
}
type UserInfoReq struct {
	Phone string `json:"phone" description:"手机号"` // id
	Id    uint64 `json:"id" description:"用户ID"`   // 用户ID
}

// UserInfoListReq 获取用户列表请求
type UserInfoListReq struct {
	Ids []uint64 `json:"ids" v:"required#用户id列表不能为空" description:"用户ID列表"`
}

type BaseUserChannelInfoReq struct {
	Ids []uint64 `json:"ids" description:"用户Ids"` // 用户ID
}
type BaseUserChannelInfo struct {
	UserId       uint64 `json:"userId" description:"用户ID"`   // 用户ID
	NickName     string `json:"nickName" description:"用户昵称"` // 用户昵称
	MpOpenId     string `json:"mpOpenId" description:"用户的公众号openid"`
	AppID        string `json:"appId" description:"用户所在渠道的公众号id"`
	AppSecret    string `json:"appSecret" description:"公众号appSecret"` // 公众号appSecret
	TemplateId   string `json:"templateId" description:"用户所在渠道的公众号模板消息ID"`
	MiniAppId    string `json:"miniAppId" description:"用户所在渠道的小程序id"`
	SmsNum       int    `json:"smsNum" description:"短信数量"`
	VmsNum       int    `json:"vmsNum" description:"语音数量"`
	PlatformName string `json:"platformName" description:"平台名称"`
}

type RefreshTokenRes struct {
	UserInfo   *BaseUserInfo `json:"userInfo" description:"用户信息"`
	Token      string        `json:"token" description:"token"`
	ExpireTime int64         `json:"expireTime" description:"过期时间"`
	MaOpenId   string        `json:"maOpenId" description:"小程序openid"` // 小程序openid
	UnionId    string        `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
}

type CreateUserAccountReq struct {
	UserId              uint64 `json:"userId" description:"用户ID"`        // 用户ID
	MpOpenId            string `json:"mpOpenId" description:"公众号openid"` // 小程序openid
	UnionId             string `json:"unionId" description:"微信开放平台id"`   // 微信开放平台id
	SubscribeNum        int    `json:"subscribeNum" description:"关注次数"`
	SubscribeScene      string `json:"subscribeScene" description:"返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENEPROFILE LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他"`
	SubscribeTime       string `json:"subscribeTime" description:"关注时间"`
	CancelSubscribeTime string `json:"cancelSubscribeTime" description:"取消关注时间"`
	QRSceneStr          string `json:"qrSceneStr" description:"二维码扫码场景"`
	Language            string `json:"language" description:"语言"`
	Subscribe           int    `json:"subscribe" description:"是否关注公众号 0-否；1-是；"`
	LongAndLati         string `json:"longAndLati" description:"经纬度  经度,纬度"`
	ChannelId           uint64 `json:"channelId" description:"渠道ID"`
}
