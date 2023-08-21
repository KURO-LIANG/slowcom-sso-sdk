package service

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

type SSORequest struct {
	SSOClient *http.SSOClient
}

// GetUserInfo 获取用户登录的信息
func (s *SSORequest) GetUserInfo(req *entity.UserInfoReq) (resData *entity.BaseUserInfo, err error) {
	url := "/api/wx/user/info"
	if req != nil {
		values, _ := query.Values(req)
		queryStr := values.Encode()
		url = fmt.Sprint(url, "?", queryStr)
	}
	res, err := s.SSOClient.Get(url)
	if err == nil {
		resData = res.Data.(*entity.BaseUserInfo)
	}
	return
}

// UserChannelInfo 获取用户渠道信息
func (s *SSORequest) UserChannelInfo(ids []uint64) (list []*entity.BaseUserChannelInfo, err error) {
	url := fmt.Sprint("/api/wx/user/channel/info?ids=", ids)
	res, err := s.SSOClient.Get(url)
	if err == nil {
		list = res.Data.([]*entity.BaseUserChannelInfo)
	}
	return
}
