package utils

import (
	"errors"
	"github.com/oa-meeting/pkg/msg"
	"go.uber.org/zap"
	"time"
)

const (
	DateTimeFormat      = "2006-01-02 15:04:05"
	DateFormat          = "2006-01-02"
	DateFormatYmdHisDot = "2006.01.02 15:04"
	DateFormatMonth     = "2006年01月02日 15:04"
)

func DatetimeToTimes(datetime string, dateFormat string) (times int32, err error) {
	if datetime == "" {
		times = 0
		return
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(dateFormat, datetime, loc)
	if err != nil {
		zap.L().Error("DatetimeToTimes err:"+datetime+":", zap.Error(err))
		err = errors.New(msg.ErrorDatetime)
		return
	}
	times = int32(t.Unix())
	return
}

func TimesToDatetime(times int32, format string) string {
	if times == 0 {
		return ""
	}
	t := time.Unix(int64(times), 0)
	if format == "" {
		format = DateTimeFormat
	}
	return t.Format(format)
}

func MonthInfo() (startTime int64, endTime int64) {
	timeNow := time.Now()
	timeToday := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location()) // 获取当天0点时间 time类型
	startTime = timeToday.AddDate(0, 0, -timeToday.Day()+1).Unix()                                         // 获取本月第一天0点 时间戳类型
	endTime = timeToday.AddDate(0, 1, -timeToday.Day()+1).Unix()
	return
}
