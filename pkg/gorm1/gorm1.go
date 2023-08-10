package gorm1

//gorm v1弃用
import (
	"fmt"
	"gin_basic/pkg/logger"
	"gin_basic/pkg/setting"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Conn *gorm.DB

// GormLogger struct
type GormLogger struct{}

// Print - Log Formatter
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logger.Infow("gin_gorm",
			"sql", v[3],
			"src", v[1],
			"duration", v[2],
			"values", v[4],
			"rows_returned", v[5],
		)
	}
}

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
	// TODO
	// https://github.com/jinzhu/gorm/blob/master/logger.go
	// http://www.manongjc.com/detail/17-gjclhsxhxgsgelp.html 日志
	gorm.DefaultTableNameHandler = func(database *gorm.DB, defaultTableName string) string {
		return setting.MysqlSetting.TablePrefix + defaultTableName
	}

	database.Debug()
	database.LogMode(true)
	database.SetLogger(&GormLogger{})

	database.SingularTable(true)                                                         // gorm会在创建表的时候去掉”s“的后缀
	database.DB().SetMaxIdleConns(setting.MysqlSetting.MaxIdleConns)                     // 设置空闲链接
	database.DB().SetMaxOpenConns(setting.MysqlSetting.MaxOpenConns)                     // 最大连接数 0为不限制
	database.DB().SetConnMaxLifetime(setting.MysqlSetting.ConnMaxLifetime * time.Minute) // 可重用链接得最大时间长度
	Conn = database
}

// Close closes database connection (unnecessary)
func Close() {
	defer Conn.Close()
}
