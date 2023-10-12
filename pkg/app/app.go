package app

import (
	rtcv2 "github.com/alibabacloud-go/rtc-20180111/v2/client"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis"
	"github.com/oa-meeting/pkg/tracing"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ModuleClients *App

type App struct {
	Lg           *zap.Logger
	RedisClient  *redis.Client
	JaegerTracer *tracing.JaegerProvider
	DbMeeting    *gorm.DB
	SfNode       *snowflake.Node
	RtcClient    *rtcv2.Client
}
