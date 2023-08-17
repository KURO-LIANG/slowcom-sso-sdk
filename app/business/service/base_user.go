package service

import (
	"fmt"
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

type SSORequest struct {
	SSOClient *http.SSOClient
}

// GetUserInfo 获取用户登录的信息
func (s *SSORequest) GetUserInfo() (resData *entity.BaseUserInfo, err error) {
	res, err := s.SSOClient.Get("/api/wx/user/info")
	if err == nil {
		resData = res.Data.(*entity.BaseUserInfo)
	}
	return
}

// GetUserInfoByPhone 通过手机号查询用户信息
func (s *SSORequest) GetUserInfoByPhone(phone string) (resData *entity.BaseUserInfo, err error) {
	res, err := s.SSOClient.Get("/api/wx/user/info?phone=" + phone)
	if err == nil {
		resData = res.Data.(*entity.BaseUserInfo)
	}
	return
}

// GetUserInfoById 通过用户id查询用户信息
func (s *SSORequest) GetUserInfoById(id uint64) (resData *entity.BaseUserInfo, err error) {
	res, err := s.SSOClient.Get(fmt.Sprint("/api/wx/user/info?id=", id))
	if err == nil {
		resData = res.Data.(*entity.BaseUserInfo)
	}
	return
}
