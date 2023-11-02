package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/entity"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

type SSOChannelRequest struct {
	SSOChannelClient *http.SSOPlatformClient
}

// ChannelList 获取平台列表信息
//func (s *SSOChannelRequest) ChannelList() (list interface{}, err error) {
//	res, err := s.SSOChannelClient.Get("/api/v1/channel/getAll")
//	if err == nil {
//		if res.Code != 0 {
//			err = errors.New(res.Message)
//			return
//		}
//		list = res.Data
//	}
//	return
//}

// GetChannelById 获取平台信息
func (s *SSOChannelRequest) GetChannelById(id uint64) (resData entity.ChannelInfo, err error) {
	res, err := s.SSOChannelClient.Get(fmt.Sprint("/platform/v1/channel/info?id=", id))
	if err == nil {
		if res.Code != 0 {
			err = errors.New(res.Message)
			return
		}
		var mapData = res.Data.(map[string]interface{})
		channel := mapData["channel"].(map[string]interface{})
		d, _ := json.Marshal(channel)
		_ = json.Unmarshal(d, &resData)
	}
	return
}
