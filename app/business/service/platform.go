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

// ChannelList 获取平台列表信息
//func (s *SSOPlatformRequest) ChannelList() (list interface{}, err error) {
//	res, err := s.SSOPlatformClient.Get("/api/v1/channel/getAll")
//	if err == nil {
//		if res.Code != 0 {
//			err = errors.New(res.Message)
//			return
//		}
//		list = res.Data
//	}
//	return
//}

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
func (s *SSOPlatformRequest) UserChannelInfo(ids []uint64) (list []entity.BaseUserChannelInfo, err error) {
	res, err := s.SSOPlatformClient.PostJson("/platform/v1/userChannelInfo", map[string]interface{}{
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
