package http

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/kuro-liang/slowcom-sso-sdk/serror"
)

type BaseRes struct {
	Code    int    `json:"code" description:"响应码 0-成功"`
	Message string `json:"message" description:"响应消息"`
	Data interface{} `json:"data" description:"数据"`
}

// checkResponse 校验请求
func checkResponse(res *httpclient.Response, requestError error) (baseRes *BaseRes, err error) {
	if requestError != nil {
		return nil, serror.New(405, "请求服务异常：请求超时")
	}
	if res.Response.StatusCode != 200 {
		return nil, serror.New(res.Response.StatusCode, fmt.Sprint(res.Response.StatusCode, " 请求服务异常"))
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &baseRes)
	if err != nil {
		return nil, serror.ErrIs数据解析异常
	}
	if baseRes.Code == 0 {
		return
	} else {
		return baseRes, serror.New(baseRes.Code, baseRes.Message)
	}
}
