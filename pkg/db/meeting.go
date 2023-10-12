package db

import (
	"github.com/google/wire"
	appConfig "github.com/oa-meeting/pkg/config"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var MeetingProvider = wire.NewSet(NewMeeting)

func NewMeeting() *gorm.DB {
	connOaMeeting := strings.Join([]string{appConfig.Data.OaMeeting.User, ":", appConfig.Data.OaMeeting.Password,
		"@tcp(", appConfig.Data.OaMeeting.Host, ":", strconv.Itoa(int(appConfig.Data.OaMeeting.Port)), ")/",
		appConfig.Data.OaMeeting.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	DbOaMeeting := loadMysqlConn(connOaMeeting)
	return DbOaMeeting
}
