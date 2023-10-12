package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oa-meeting/pkg/app"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

func loadMysqlConn(conn string) *gorm.DB {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,  // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	return db
}

func DBMigration() {
	//addColumn(&model_account.Users{}, "action_code")
	err := app.ModuleClients.DbMeeting.AutoMigrate()
	if err != nil {
		zap.L().Error("register table fail--", zap.Error(err))
		os.Exit(0)
	}
}

func addColumn(dst interface{}, column string) {
	exist := app.ModuleClients.DbMeeting.Migrator().HasColumn(dst, column)
	if !exist {
		err := app.ModuleClients.DbMeeting.Migrator().AddColumn(dst, column)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}
