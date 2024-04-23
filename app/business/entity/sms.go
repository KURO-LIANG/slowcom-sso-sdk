package entity

type SendDeviceNoticeReq struct {
	MsgType         int8   `json:"msgType" v:"required#消息类型不能为空" description:"消息类型 0-短信，1-语音短信"` // 消息类型 0-短信，1-语音短信
	Phone           string `json:"phone" v:"required#手机号不能为空" description:"手机号"`
	Content         string `json:"content" v:"required#通知内容不能为空" description:"通知内容"`
	UserName        string `json:"userName" v:"required#通知用户称呼不能为空" description:"通知用户称呼"`
	UserId          uint64 `json:"userId" v:"required#通知用户Id不能为空" description:"通知用户ID"`
	GroupId         string `json:"groupId" v:"required#报警家庭组ID不能为空" description:"报警家庭组ID"`
	DeviceAlarmType string `json:"deviceAlarmType" description:"报警类型"`
	AlarmTime       string `json:"alarmTime" description:"报警发生时间"`
	FamilyName      string `json:"familyName" description:"家庭名称"`
	MsgAlarmType    int8   `json:"msgAlarmType" description:"消息类型"`
}
