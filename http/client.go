package http

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	netUrl "net/url"
	"sync"
)

type SSOClient struct {
	BaseUrl      string
	Token  string
	rwLock       sync.RWMutex
}

const (
	USERAGENT       = "slowcom_agent"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 10
)

// 生成一个http请求客户端
func buildHttpClient() *httpclient.HttpClient {
	return httpclient.NewHttpClient().Defaults(httpclient.Map{
		"opt_useragent":      USERAGENT,
		"opt_timeout":        TIMEOUT,
		"Accept-Encoding":    "gzip,deflate,sdch",
		"opt_connecttimeout": CONNECT_TIMEOUT,
		"OPT_DEBUG":          true,
	})
}

// Get get请求
func (s *SSOClient) Get(url string) (response *BaseRes, err error) {
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+s.Token).
		Get(fmt.Sprintf("%s%s", s.BaseUrl, url), netUrl.Values{})
	response, err = checkResponse(res, err)
	return
}
