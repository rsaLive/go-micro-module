// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/oa-meeting/pkg/aliyun"
	"github.com/oa-meeting/pkg/app"
	"github.com/oa-meeting/pkg/cache"
	"github.com/oa-meeting/pkg/db"
	"github.com/oa-meeting/pkg/logger"
	"github.com/oa-meeting/pkg/snowf"
	"github.com/oa-meeting/pkg/tracing"
)

func InitApp() (*app.App, error) {
	wire.Build(logger.Provider, cache.RedisProvider, tracing.Provider, db.MeetingProvider, snowf.Provider, aliyun.RtcProvider, NewApp)
	return &app.App{}, nil
}
