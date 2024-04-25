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
