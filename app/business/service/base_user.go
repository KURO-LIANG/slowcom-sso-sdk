package service

import (
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

type SSORequest struct {
	SSOClient *http.SSOClient
}

// GetUserInfo 获取用户登录的信息
func (s *SSORequest) GetUserInfo(phone string) (resData *entity.BaseUserInfo, err error) {
	res, err := s.SSOClient.Get("/api/wx/user/info?phone=" + phone)
	if err == nil {
		resData = res.Data.(*entity.BaseUserInfo)
	}
	return
}
