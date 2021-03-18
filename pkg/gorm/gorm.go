package gorm

import (
	"fmt"
	"gin_basic/pkg/logger"
	"gin_basic/pkg/setting"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

// Setup initializes the database instance
func Setup() {
	database, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.UserName,
		setting.MysqlSetting.PassWord,
		setting.MysqlSetting.IPHost,
		setting.MysqlSetting.Port,
		setting.MysqlSetting.DbName))

	if err != nil {
		logger.Fatal("gorm.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(database *gorm.DB, defaultTableName string) string {
		return setting.MysqlSetting.TablePrefix + defaultTableName
	}

	database.SingularTable(true)                                                         //gorm会在创建表的时候去掉”s“的后缀
	database.DB().SetMaxIdleConns(setting.MysqlSetting.MaxIdleConns)                     // 设置空闲链接
	database.DB().SetMaxOpenConns(setting.MysqlSetting.MaxOpenConns)                     // 最大连接数 0为不限制
	database.DB().SetConnMaxLifetime(setting.MysqlSetting.ConnMaxLifetime * time.Minute) // 可重用链接得最大时间长度
	Conn = database
}

// Close closes database connection (unnecessary)
func Close() {
	defer Conn.Close()
}
