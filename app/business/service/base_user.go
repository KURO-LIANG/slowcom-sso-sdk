package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

type SSOUserRequest struct {
	SSOUserClient *http.SSOClient
}

// GetUserInfo 获取用户登录的信息
func (s *SSOUserRequest) GetUserInfo(req *entity.UserInfoReq) (resData *entity.BaseUserInfo, err error) {
	url := "/api/v1/wx/mini/user/info"
	if req != nil {
		values, _ := query.Values(req)
		queryStr := values.Encode()
		url = fmt.Sprint(url, "?", queryStr)
	}
	res, err := s.SSOUserClient.Get(url)
	if err == nil {
		var mapData = res.Data.(map[string]interface{})
		user := mapData["userInfo"].(map[string]interface{})
		d, _ := json.Marshal(user)
		_ = json.Unmarshal(d, &resData)
	}
	return
}

// UserChannelInfo 获取用户渠道信息
func (s *SSOUserRequest) UserChannelInfo(ids []uint64) (list []entity.BaseUserChannelInfo, err error) {
	res, err := s.SSOUserClient.PostJson("/api/v1/wx/mini/user/channel/info", map[string]interface{}{
		"ids": ids,
	})
	if err == nil {
		data := res.Data.(map[string]interface{})
		dataList := data["list"].([]interface{})
		d, _ := json.Marshal(dataList)
		_ = json.Unmarshal(d, &list)
	}
	return
}
