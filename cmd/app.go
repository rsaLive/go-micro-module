package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps/strategy"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis"
	_ "github.com/oa-meeting/internal/handler"
	appConfig "github.com/oa-meeting/pkg/config"
	"github.com/oa-meeting/pkg/tracing"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	lg           *zap.Logger
	RedisClient  *redis.Client
	JaegerTracer *tracing.JaegerProvider
	MysqlDb      *gorm.DB
	sfNode       *snowflake.Node
}

func NewApp(lg *zap.Logger, RedisClient *redis.Client, JaegerTracer *tracing.JaegerProvider, MysqlDb *gorm.DB, sfNode *snowflake.Node) *App {
	return &App{
		lg:           lg,
		RedisClient:  RedisClient,
		JaegerTracer: JaegerTracer,
		MysqlDb:      MysqlDb,
		sfNode:       sfNode,
	}
}

func main() {
	var err error
	appConfig.GetOptions()
	_, err = InitApp()
	if err != nil {
		panic(err)
	}
	//注册服务
	//config.SetProviderService(&controller.OrderProvider{})
	if err = config.Load(); err != nil {
		panic(err)
	}
	select {}
}
