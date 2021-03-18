package gorm

import (
	"fmt"
	"gin_basic/pkg/setting"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"time"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.UserName,
		setting.MysqlSetting.PassWord,
		setting.MysqlSetting.IPHost,
		setting.MysqlSetting.Port,
		setting.MysqlSetting.DbName))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.MysqlSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)                                                         //gorm会在创建表的时候去掉”s“的后缀
	db.DB().SetMaxIdleConns(setting.MysqlSetting.MaxIdleConns)                     // 设置空闲链接
	db.DB().SetMaxOpenConns(setting.MysqlSetting.MaxOpenConns)                     // 最大连接数 0为不限制
	db.DB().SetConnMaxLifetime(setting.MysqlSetting.ConnMaxLifetime * time.Minute) // 可重用链接得最大时间长度
}

// Close closes database connection (unnecessary)
func Close() {
	defer db.Close()
}
