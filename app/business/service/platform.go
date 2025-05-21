package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

type SSOPlatformRequest struct {
	SSOPlatformClient *http.SSOPlatformClient
}

// GetPlatformInfo 获取平台信息
func (s *SSOPlatformRequest) GetPlatformInfo(id uint64) (resData entity.ChannelInfo, err error) {
	res, err := s.SSOPlatformClient.Get(fmt.Sprint("/platform/v1/info?id=", id))
	if err == nil {
		if res.Code != 0 {
			err = errors.New(res.Message)
			return
		}
		var mapData = res.Data.(map[string]interface{})
		platformInfo := mapData["platformInfo"].(map[string]interface{})
		d, _ := json.Marshal(platformInfo)
		_ = json.Unmarshal(d, &resData)
	}
	return
}

// UserChannelInfo 获取用户渠道信息
func (s *SSOPlatformRequest) UserChannelInfo(ids []uint64, clientId string) (list []entity.BaseUserChannelInfo, err error) {
	res, err := s.SSOPlatformClient.PostJson("/platform/v1/userChannelInfo", map[string]interface{}{
		"ids":      ids,
		"clientId": clientId,
	})
	if err == nil {
		data := res.Data.(map[string]interface{})
		if data["list"] != nil {
			dataList := data["list"].([]interface{})
			d, _ := json.Marshal(dataList)
			_ = json.Unmarshal(d, &list)
		}
	}
	return
}

// SendMpTemplateNotice 发送模板消息通知
func (s *SSOPlatformRequest) SendMpTemplateNotice(req *entity.MpTemplateReq) (res *http.BaseRes, err error) {
	res, err = s.SSOPlatformClient.PostJson("/platform/v1/mp/template/notice", req)
	return
}

// GetChannelList 获取渠道列表
func (s *SSOPlatformRequest) GetChannelList() (list []entity.ChannelList, err error) {
	res, err := s.SSOPlatformClient.Get("/platform/v1/channel/list")
	if err == nil {
		data := res.Data.(map[string]interface{})
		if data["list"] != nil {
			dataList := data["list"].([]interface{})
			d, _ := json.Marshal(dataList)
			_ = json.Unmarshal(d, &list)
		}
	}
	return
}

// UserInfoByMpOpenId 获取公众号用户信息
func (s *SSOPlatformRequest) UserInfoByMpOpenId(mpOpenId string, unionId string) (resData *entity.BaseUserInfo, err error) {
	res, err := s.SSOPlatformClient.Get(fmt.Sprint("/platform/v1/getUserInfo?mpOpenId=", mpOpenId, "&unionId=", unionId))
	if err == nil {
		var mapData = res.Data.(map[string]interface{})
		user := mapData["userInfo"].(map[string]interface{})
		d, _ := json.Marshal(user)
		_ = json.Unmarshal(d, &resData)
	}
	return
}

// SaveUserAccount 创建公众号用户信息
func (s *SSOPlatformRequest) SaveUserAccount(req entity.CreateUserAccountReq) (resData *entity.BaseUserInfo, err error) {
	res, err := s.SSOPlatformClient.PostJson("/platform/v1/createUserAccount", req)
	if err == nil {
		var mapData = res.Data.(map[string]interface{})
		user := mapData["userInfo"].(map[string]interface{})
		d, _ := json.Marshal(user)
		_ = json.Unmarshal(d, &resData)
	}
	return
}
