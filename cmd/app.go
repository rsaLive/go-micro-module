package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps/strategy"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	rtcv2 "github.com/alibabacloud-go/rtc-20180111/v2/client"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis"
	"github.com/oa-meeting/internal/controller"
	_ "github.com/oa-meeting/internal/handler"
	"github.com/oa-meeting/pkg/app"
	appConfig "github.com/oa-meeting/pkg/config"
	common "github.com/oa-meeting/pkg/init"
	"github.com/oa-meeting/pkg/tracing"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewApp(Lg *zap.Logger, RedisClient *redis.Client, JaegerTracer *tracing.JaegerProvider, DbMeeting *gorm.DB, SfNode *snowflake.Node, RtcClient *rtcv2.Client) *app.App {
	return &app.App{
		Lg:           Lg,
		RedisClient:  RedisClient,
		JaegerTracer: JaegerTracer,
		DbMeeting:    DbMeeting,
		SfNode:       SfNode,
		RtcClient:    RtcClient,
	}
}

func main() {
	var err error
	appConfig.GetOptions()
	app.ModuleClients, err = InitApp()
	if err != nil {
		panic(err)
	}
	//注册服务
	config.SetProviderService(&controller.MeetingProvider{})
	common.Init()
	if err = config.Load(); err != nil {
		panic(err)
	}
	select {}
}
