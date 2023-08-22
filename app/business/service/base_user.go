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
	url := "/api/v1/wx/mini/user/info"
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
	res, err := s.SSOClient.PostJson("/api/v1/wx/mini/user/channel/info", map[string]interface{}{
		"ids": ids,
	})
	if err == nil {
		list = res.Data.([]*entity.BaseUserChannelInfo)
	}
	return
}
