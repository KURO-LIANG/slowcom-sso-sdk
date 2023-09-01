package service

import (
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

type SSOSmsRequest struct {
	SSOSmsClient *http.SSOClient
}

// SendVoiceNotice 发送设备报警通知
func (s *SSOSmsRequest) SendVoiceNotice(req *entity.SendDeviceNoticeReq) (res *http.BaseRes, err error) {
	res, err = s.SSOSmsClient.PostJson("/api/v1/sms/notice/sendDeviceNotice", req)
	return
}
