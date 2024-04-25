package http

import (
	"fmt"
	netUrl "net/url"
	"sync"
)

type SSOPlatformClient struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	rwLock       sync.RWMutex
}

// Get get请求
func (s *SSOPlatformClient) Get(url string) (response *BaseRes, err error) {
	res, err := buildHttpClient().WithHeader("Client-Id", s.ClientId).WithHeader("Client-Secret", s.ClientSecret).
		Get(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	response, err = checkResponse(res, err)
	return
}

// PostJson json请求
func (s *SSOPlatformClient) PostJson(url string, data interface{}) (response *BaseRes, err error) {
	res, err := buildHttpClient().WithHeader("Client-Id", s.ClientId).WithHeader("Client-Secret", s.ClientSecret).
		PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	response, err = checkResponse(res, err)
	return
}

// PutJson json请求
func (s *SSOPlatformClient) PutJson(url string, data interface{}) (response *BaseRes, err error) {
	res, err := buildHttpClient().WithHeader("Client-Id", s.ClientId).WithHeader("Client-Secret", s.ClientSecret).
		PutJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	response, err = checkResponse(res, err)
	return
}
