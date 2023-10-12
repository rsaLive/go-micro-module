package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	rtcv2 "github.com/alibabacloud-go/rtc-20180111/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/google/wire"
	appConfig "github.com/oa-meeting/pkg/config"
)

var RtcProvider = wire.NewSet(NewRtcClient)

func NewRtcClient() *rtcv2.Client {
	config := &openapi.Config{
		AccessKeyId:     &appConfig.Data.AliYun.AccessKeyId,
		AccessKeySecret: &appConfig.Data.AliYun.AccessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/rtc
	config.Endpoint = tea.String("rtc.aliyuncs.com")
	client, err := rtcv2.NewClient(config)
	if err != nil {
		panic(client)
	}
	return client
}
