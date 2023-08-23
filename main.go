package main

import (
	"fmt"
	"github.com/kuro-liang/slowcom-sso-sdk/app/business/service"
	"github.com/kuro-liang/slowcom-sso-sdk/http"
)

func main() {
	client := service.SSORequest{SSOClient: &http.SSOClient{
		BaseUrl: "https://wx.api.sso.slowcom.cn",
	}}
	list, err := client.UserChannelInfo([]uint64{11, 9})
	if err != nil {
		fmt.Printf("错误：%s", err.Error())
		return
	}
	fmt.Printf("数据条数：%d", len(list))
}
