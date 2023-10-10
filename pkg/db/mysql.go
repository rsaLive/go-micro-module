package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	appConfig "github.com/oa-meeting/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"strconv"
	"strings"
	"time"
)

var DbAccount, DbGoods, DbOrder, DbMysql *gorm.DB

var Provider = wire.NewSet(NewDb)

func NewDb() *gorm.DB {
	connOrder := strings.Join([]string{appConfig.Data.MealOrder.User, ":", appConfig.Data.MealOrder.Password, "@tcp(", appConfig.Data.MealOrder.Host, ":", strconv.Itoa(int(appConfig.Data.MealOrder.Port)), ")/", appConfig.Data.MealOrder.DbName, "?charset=utf8mb4&parseTime=true"}, "")
	DbOrder = loadMysqlConn(connOrder)
	migration()
	return DbMysql
}

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

// 执行数据迁移
func migration() {
	//addColumn(&model_account.Users{}, "action_code")
	errOrder := DbOrder.AutoMigrate()

	if errOrder != nil {
		zap.L().Error("register table fail--", zap.Error(errOrder))
		os.Exit(0)
	}
}

func addColumn(dst interface{}, column string) {
	exist := DbAccount.Migrator().HasColumn(dst, column)
	if !exist {
		err := DbAccount.Migrator().AddColumn(dst, column)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return
}
