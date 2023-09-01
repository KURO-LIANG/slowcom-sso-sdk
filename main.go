package main

import (
	"fmt"
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/service"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

func main() {
	client := service.SSOUserRequest{SSOUserClient: &http.SSOClient{
		BaseUrl: "https://wx.api.sso.slowcom.cn",
		Token:   "Bearer B5idb5D01+uHITLlexQMUTZjaveDxr8qgbGroLZJVqQpBKSXuVmDDr9/vaZ8vAK7DGLBCKFsef4+0k7TpIAvmF1UTXyO4Q0iruS3S1WH9sMIlfiNpeyVr3Wi/aSNKJ2UlxNskUy/WoJC8TNSkgCrsg==",
	}}
	res, err := client.UserChannelInfo([]uint64{10, 9})
	if err != nil {
		fmt.Printf("错误：%s", err.Error())
		return
	}
	fmt.Printf("数据：%+v", res)
}
