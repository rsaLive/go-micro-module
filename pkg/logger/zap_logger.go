package logger

import (
	"github.com/google/wire"
	"os"

	appConfig "github.com/oa-meeting/pkg/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Provider = wire.NewSet(ZapInit)

// ZapInit 初始化lg
func ZapInit() *zap.Logger {
	var err error
	writeSyncer := getLogWriter(appConfig.Data.ZapLog.Filename, appConfig.Data.ZapLog.MaxSize, appConfig.Data.ZapLog.MaxBackups, appConfig.Data.ZapLog.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(appConfig.Data.ZapLog.Level))
	if err != nil {
		return nil
	}
	var core zapcore.Core
	if appConfig.Data.System.Mode == "dev" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	return lg
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
